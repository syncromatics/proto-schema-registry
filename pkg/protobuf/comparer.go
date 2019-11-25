package protobuf

import (
	"bytes"
	"fmt"

	"github.com/emicklei/proto"
	"github.com/pkg/errors"
)

// CheckForBreakingChanges will compare the current protobuf definition with a new one
// and detect breaking changes.
func CheckForBreakingChanges(current []byte, new []byte) (bool, []string, error) {
	p := proto.NewParser(bytes.NewReader(current))
	parser, err := p.Parse()
	if err != nil {
		return false, nil, errors.Wrap(err, "failed to parse current protobuf file")
	}

	currentDefinition := newDefinition()
	proto.Walk(parser,
		proto.WithEnum(currentDefinition.HandleEnum),
		proto.WithMessage(currentDefinition.HandleMessage))

	p = proto.NewParser(bytes.NewReader(new))
	parser, err = p.Parse()
	if err != nil {
		return false, nil, errors.Wrap(err, "failed to parse new protobuf file")
	}
	newDefinition := newDefinition()
	proto.Walk(parser,
		proto.WithEnum(newDefinition.HandleEnum),
		proto.WithMessage(newDefinition.HandleMessage))

	errors := compareDefinitions(currentDefinition, newDefinition)
	if len(errors) > 0 {
		return false, errors, nil
	}

	return true, nil, nil
}

type enumValueDefinition struct {
	value string
}
type enumDefintion struct {
	values map[int]enumValueDefinition
}

func newEnumDefinition(enum *proto.Enum) enumDefintion {
	e := enumDefintion{
		values: map[int]enumValueDefinition{},
	}

	for _, ee := range enum.Elements {
		switch v := ee.(type) {
		case *proto.EnumField:
			e.values[v.Integer] = enumValueDefinition{
				value: v.Name,
			}
		}
	}

	return e
}

type fieldDefinition struct {
	fieldtype string
	repeated  bool
}
type oneOfDefintion struct {
	fields map[int]fieldDefinition
}

func newOneOfDefintion(oneOf *proto.Oneof) oneOfDefintion {
	o := oneOfDefintion{
		fields: map[int]fieldDefinition{},
	}

	for _, d := range oneOf.Elements {
		switch v := d.(type) {
		case *proto.OneOfField:
			o.fields[v.Sequence] = fieldDefinition{
				fieldtype: v.Type,
			}
		}
	}
	return o
}

type messageDefinition struct {
	fields   map[int]fieldDefinition
	enums    map[string]enumDefintion
	oneOfs   map[string]oneOfDefintion
	reserved map[int]struct{}
}

func newMessageDefinition() messageDefinition {
	return messageDefinition{
		fields:   map[int]fieldDefinition{},
		enums:    map[string]enumDefintion{},
		oneOfs:   map[string]oneOfDefintion{},
		reserved: map[int]struct{}{},
	}
}

type definition struct {
	enums    map[string]enumDefintion
	messages map[string]messageDefinition
}

func newDefinition() *definition {
	return &definition{
		enums:    map[string]enumDefintion{},
		messages: map[string]messageDefinition{},
	}
}

func (d *definition) HandleEnum(enum *proto.Enum) {
	if _, ok := enum.Parent.(*proto.Message); ok {
		return
	}

	e := newEnumDefinition(enum)
	d.enums[enum.Name] = e
}

func (d *definition) HandleMessage(message *proto.Message) {
	m := newMessageDefinition()

	for _, f := range message.Elements {
		switch v := f.(type) {
		case *proto.NormalField:
			m.fields[v.Sequence] = fieldDefinition{
				fieldtype: v.Type,
				repeated:  v.Repeated,
			}

		case *proto.Enum:
			m.enums[v.Name] = newEnumDefinition(v)

		case *proto.Oneof:
			m.oneOfs[v.Name] = newOneOfDefintion(v)

		case *proto.Reserved:
			for _, r := range v.Ranges {
				for i := r.From; i <= r.To; i++ {
					m.reserved[i] = struct{}{}
				}
			}
		}
	}

	d.messages[message.Name] = m
}

func compareDefinitions(current *definition, new *definition) []string {
	errors := []string{}
	for name, message := range current.messages {
		newMessage, ok := new.messages[name]
		if !ok {
			errors = append(errors, fmt.Sprintf("new definition does not contain the message '%s'", name))
			continue
		}

		for fieldSequence, field := range message.fields {
			_, isReserved := message.reserved[fieldSequence]
			if isReserved {
				continue
			}

			_, isReserved = newMessage.reserved[fieldSequence]
			if isReserved {
				continue
			}

			newfield, ok := newMessage.fields[fieldSequence]
			if !ok {
				errors = append(errors, fmt.Sprintf("message '%s' does not contain field sequence '%d' with field type '%s'", name, fieldSequence, field.fieldtype))
				continue
			}

			if newfield.fieldtype != field.fieldtype {
				errors = append(errors, fmt.Sprintf("message '%s' had field type '%s' changed to '%s' at sequence '%d'", name, field.fieldtype, newfield.fieldtype, fieldSequence))
				continue
			}

			if newfield.repeated != field.repeated {
				errors = append(errors, fmt.Sprintf("message '%s' had field repeated '%t' changed to '%t' at sequence '%d'", name, field.repeated, newfield.repeated, fieldSequence))
				continue
			}
		}

		for enumName, enum := range message.enums {
			newEnum, ok := newMessage.enums[enumName]
			if !ok {
				errors = append(errors, fmt.Sprintf("message '%s' is missing embedded enum '%s'", name, enumName))
				continue
			}

			for valueSequence, value := range enum.values {
				newvalue, ok := newEnum.values[valueSequence]
				if !ok {
					errors = append(errors, fmt.Sprintf("message '%s' embedded enum '%s' is missing value '%s' at sequence '%d'", name, enumName, value, valueSequence))
					continue
				}

				if newvalue != value {
					errors = append(errors, fmt.Sprintf("message '%s' embedded enum '%s' value changed from '%s' to '%s' at sequence '%d'", name, enumName, value, newvalue, valueSequence))
					continue
				}
			}
		}

		for oneOfName, oneOf := range message.oneOfs {

			newOneOf, ok := newMessage.oneOfs[oneOfName]
			if !ok {
				errors = append(errors, fmt.Sprintf("message '%s' does not contain oneOf '%s'", name, oneOfName))
				continue
			}

			for oneOfSequence, value := range oneOf.fields {
				_, isReserved := message.reserved[oneOfSequence]
				if isReserved {
					continue
				}

				_, isReserved = newMessage.reserved[oneOfSequence]
				if isReserved {
					continue
				}

				newValue, ok := newOneOf.fields[oneOfSequence]
				if !ok {
					errors = append(errors, fmt.Sprintf("message '%s' oneOf '%s' does not contain sequence '%d' with type '%s'", name, oneOfName, oneOfSequence, value.fieldtype))
					continue
				}

				if newValue.fieldtype != value.fieldtype {
					errors = append(errors, fmt.Sprintf("message '%s' oneOf '%s' type changed from '%s' to '%s' at sequence '%d'", name, oneOfName, value.fieldtype, newValue.fieldtype, oneOfSequence))
					continue
				}
			}
		}

		for i := range message.reserved {
			_, ok := newMessage.reserved[i]
			if !ok {
				errors = append(errors, fmt.Sprintf("message '%s' reserved '%d' was removed", name, i))
				continue
			}
		}
	}

	for name, enum := range current.enums {
		newEnum, ok := new.enums[name]
		if !ok {
			errors = append(errors, fmt.Sprintf("new definition does not contain enum '%s'", name))
			continue
		}

		for sequence, value := range enum.values {
			newvalue, ok := newEnum.values[sequence]
			if !ok {
				errors = append(errors, fmt.Sprintf("enum '%s' does not contain value sequence '%d' with name '%s'", name, sequence, value))
				continue
			}

			if newvalue != value {
				errors = append(errors, fmt.Sprintf("enum '%s' had value changed from '%s' to '%s' at sequence '%d'", name, value, newvalue, sequence))
				continue
			}
		}
	}

	return errors
}
