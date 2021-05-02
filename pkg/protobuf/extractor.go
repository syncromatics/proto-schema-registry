package protobuf

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io/ioutil"
	"reflect"
	"strings"

	"github.com/golang/protobuf/descriptor"
	"github.com/golang/protobuf/proto"
	protobuf "github.com/golang/protobuf/protoc-gen-go/descriptor"
	"github.com/pkg/errors"
)

// ExtractSchema will extract a flattened out schema from a protobuf message
func ExtractSchema(message descriptor.Message) (string, error) {
	messageType := reflect.TypeOf(message).Elem().Name()

	d, _ := message.Descriptor()
	desc, err := extractDescriptor(d)
	if err != nil {
		return "", err
	}

	types := &typesGroup{
		messageTypes: map[string]*typeHolder{},
		enumTypes:    map[string]enumType{},
		seenPackages: map[string]struct{}{},
	}
	err = getTypes(desc, types)
	if err != nil {
		return "", errors.Wrap(err, "failed to get types")
	}

	var root *protobuf.DescriptorProto
	for _, m := range desc.MessageType {
		if m.Name != nil && *m.Name == messageType {
			root = m
			break
		}
	}

	if root == nil {
		return "", errors.Errorf("could not find root type '%s'", messageType)
	}

	name := fmt.Sprintf(".%s.%s", *desc.Package, messageType)

	tree := &node{
		Value:  name,
		Leaves: []*node{},
	}

	err = buildTree(tree, root, types)
	if err != nil {
		return "", errors.Wrap(err, "failed to build the tree")
	}

	builder := &strings.Builder{}

	builder.WriteString("syntax = \"proto3\";\n")
	builder.WriteString("package gen;\n")

	err = writeTree(tree, &writeTreeJob{
		builder:    builder,
		seenTypes:  map[string]struct{}{},
		types:      types,
		recordType: name})
	if err != nil {
		return "", err
	}

	return builder.String(), nil
}

type enumType struct {
	isEmbedded bool
	parentName string
	descriptor *protobuf.EnumDescriptorProto
	pkg        string
}

type typeHolder struct {
	messageType *protobuf.DescriptorProto
	visited     bool
}

type typesGroup struct {
	messageTypes map[string]*typeHolder
	enumTypes    map[string]enumType
	seenPackages map[string]struct{}
}

func getTypes(pkg *protobuf.FileDescriptorProto, types *typesGroup) error {
	if pkg.Package == nil {
		return errors.Errorf("package does not have package name")
	}

	p := fmt.Sprintf("%s_%s", *pkg.Package, *pkg.Name)

	_, ok := types.seenPackages[p]
	if ok {
		return nil
	}

	types.seenPackages[p] = struct{}{}

	for _, m := range pkg.MessageType {
		if m.Name == nil {
			continue
		}

		name := fmt.Sprintf(".%s.%s", *pkg.Package, *m.Name)

		types.messageTypes[name] = &typeHolder{messageType: m}

		for _, nt := range m.NestedType {
			name := fmt.Sprintf(".%s.%s.%s", *pkg.Package, *m.Name, *nt.Name)
			types.messageTypes[name] = &typeHolder{messageType: nt}
		}

		for _, enum := range m.EnumType {
			ename := fmt.Sprintf(".%s.%s.%s", *pkg.Package, *m.Name, *enum.Name)
			types.enumTypes[ename] = enumType{
				descriptor: enum,
				isEmbedded: true,
				parentName: name,
				pkg:        *pkg.Package,
			}
		}
	}

	for _, enum := range pkg.EnumType {
		name := fmt.Sprintf(".%s.%s", *pkg.Package, *enum.Name)
		types.enumTypes[name] = enumType{
			descriptor: enum,
			isEmbedded: false,
		}
	}

	for _, de := range pkg.Dependency {
		dedesc := proto.FileDescriptor(de)
		dd, err := extractDescriptor(dedesc)
		if err != nil {
			return err
		}

		err = getTypes(dd, types)
		if err != nil {
			return err
		}
	}

	return nil
}

func extractDescriptor(m []byte) (*protobuf.FileDescriptorProto, error) {
	r, err := gzip.NewReader(bytes.NewReader(m))
	if err != nil {
		return nil, errors.Wrap(err, "failed to create gzip ready")
	}
	defer r.Close()

	buf, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, errors.Wrap(err, "failed to read from gzip")
	}

	desc := &protobuf.FileDescriptorProto{}
	err = proto.Unmarshal(buf, desc)
	if err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal proto")
	}

	return desc, nil
}

func buildTree(tree *node, d *protobuf.DescriptorProto, types *typesGroup) error {
	for _, t := range d.Field {
		if *t.Type != protobuf.FieldDescriptorProto_TYPE_MESSAGE && *t.Type != protobuf.FieldDescriptorProto_TYPE_ENUM {
			continue
		}

		leaf := &node{
			Value:  *t.TypeName,
			Leaves: []*node{},
		}

		ty, ok := types.messageTypes[*t.TypeName]

		if !ok {
			et, ok := types.enumTypes[*t.TypeName]
			if !ok {
				return errors.Errorf("could not find type '%s'", *t.TypeName)
			}

			if et.isEmbedded {
				if fmt.Sprintf(".%s.%s", et.pkg, *d.Name) == et.parentName {
					return nil
				}
				parent, ok := types.messageTypes[et.parentName]
				if !ok {
					return errors.Errorf("enum parent '%s' not found", et.parentName)
				}

				leaf.Value = et.parentName
				err := buildTree(leaf, parent.messageType, types)
				if err != nil {
					return err
				}
				continue
			} else {
				tree.Leaves = append(tree.Leaves, leaf)
				continue
			}
		}

		if ty.visited {
			continue
		}
		ty.visited = true

		tree.Leaves = append(tree.Leaves, leaf)

		err := buildTree(leaf, ty.messageType, types)
		if err != nil {
			return err
		}
	}
	return nil
}

type writeTreeJob struct {
	recordType string
	builder    *strings.Builder
	seenTypes  map[string]struct{}
	types      *typesGroup
}

func writeTree(tree *node, job *writeTreeJob) error {
	var err error
	for _, n := range tree.Leaves {
		err = writeTree(n, job)
		if err != nil {
			return err
		}
	}

	if _, ok := job.seenTypes[tree.Value]; ok {
		return nil
	}

	job.seenTypes[tree.Value] = struct{}{}

	t, ok := job.types.messageTypes[tree.Value]
	if !ok {
		e, ok := job.types.enumTypes[tree.Value]
		if !ok {
			return errors.Errorf("failed to find type '%s'", tree.Value)
		}

		return writeEnum(tree.Value, e.descriptor, job)
	}

	var name string
	if tree.Value == job.recordType {
		name = "record"
	} else {
		name = tree.Value[1:]
		name = strings.Replace(name, ".", "_", -1)
	}

	job.builder.WriteString(fmt.Sprintf("message %s {\n", name))
	inOnOf := false
	for _, v := range t.messageType.Field {

		if v.OneofIndex != nil && !inOnOf {
			job.builder.WriteString(fmt.Sprintf("	oneof oneof_%d {\n", *v.OneofIndex))
			inOnOf = true
		}

		if v.OneofIndex != nil {
			job.builder.WriteString("	")
		}

		if v.OneofIndex == nil && inOnOf {
			job.builder.WriteString("	}\n")
			inOnOf = false
		}

		job.builder.WriteString("	")

		switch *v.Label {
		case protobuf.FieldDescriptorProto_LABEL_REPEATED:
			job.builder.WriteString("repeated ")
		}

		switch *v.Type {
		case protobuf.FieldDescriptorProto_TYPE_BOOL:
			job.builder.WriteString("bool")

		case protobuf.FieldDescriptorProto_TYPE_BYTES:
			job.builder.WriteString("bytes")

		case protobuf.FieldDescriptorProto_TYPE_DOUBLE:
			job.builder.WriteString("double")

		case protobuf.FieldDescriptorProto_TYPE_FIXED32:
			job.builder.WriteString("fixed32")

		case protobuf.FieldDescriptorProto_TYPE_FIXED64:
			job.builder.WriteString("fixed64")

		case protobuf.FieldDescriptorProto_TYPE_FLOAT:
			job.builder.WriteString("float")

		case protobuf.FieldDescriptorProto_TYPE_INT32:
			job.builder.WriteString("int32")

		case protobuf.FieldDescriptorProto_TYPE_INT64:
			job.builder.WriteString("int64")

		case protobuf.FieldDescriptorProto_TYPE_SFIXED32:
			job.builder.WriteString("sfixed32")

		case protobuf.FieldDescriptorProto_TYPE_SFIXED64:
			job.builder.WriteString("sfixed64")

		case protobuf.FieldDescriptorProto_TYPE_SINT32:
			job.builder.WriteString("sint32")

		case protobuf.FieldDescriptorProto_TYPE_SINT64:
			job.builder.WriteString("sint64")

		case protobuf.FieldDescriptorProto_TYPE_STRING:
			job.builder.WriteString("string")

		case protobuf.FieldDescriptorProto_TYPE_UINT32:
			job.builder.WriteString("uint32")

		case protobuf.FieldDescriptorProto_TYPE_UINT64:
			job.builder.WriteString("uint64")

		case protobuf.FieldDescriptorProto_TYPE_MESSAGE:
			n := (*v.TypeName)[1:]
			n = strings.Replace(n, ".", "_", -1)
			job.builder.WriteString(n)

		case protobuf.FieldDescriptorProto_TYPE_ENUM:
			enum, ok := job.types.enumTypes[*v.TypeName]
			if !ok {
				return errors.Errorf("enum of type '%s' not found", *v.TypeName)
			}

			n := (*v.TypeName)[1:]
			if enum.isEmbedded {
				if enum.parentName == tree.Value {
					job.builder.WriteString(*enum.descriptor.Name)
				} else {
					frags := strings.Split(n, ".")
					for i, s := range frags {
						job.builder.WriteString(s)
						switch {
						case i < len(frags)-2:
							job.builder.WriteString("_")
						case i == len(frags)-2:
							job.builder.WriteString(".")
						}
					}
				}
			} else {
				n = strings.Replace(n, ".", "_", -1)
				job.builder.WriteString(n)
			}
		}
		job.builder.WriteString(fmt.Sprintf(" %s = %d;\n", *v.Name, *v.Number))
	}
	if inOnOf {
		job.builder.WriteString("	}\n")
	}

	for _, e := range t.messageType.EnumType {
		job.builder.WriteString(fmt.Sprintf("	enum %s {\n", *e.Name))
		for _, o := range e.Value {
			job.builder.WriteString(fmt.Sprintf("		%s = %d;\n", *o.Name, *o.Number))
		}
		job.builder.WriteString("	}\n")
	}
	job.builder.WriteString("}\n")

	return nil
}

func writeEnum(name string, enum *protobuf.EnumDescriptorProto, job *writeTreeJob) error {
	n := strings.Replace(name[1:], ".", "_", -1)
	job.builder.WriteString(fmt.Sprintf("enum %s {\n", n))

	for _, o := range enum.Value {
		job.builder.WriteString(fmt.Sprintf("	%s = %d;\n", *o.Name, *o.Number))
	}
	job.builder.WriteString("}\n")
	return nil
}

type node struct {
	Value  string
	Leaves []*node
}
