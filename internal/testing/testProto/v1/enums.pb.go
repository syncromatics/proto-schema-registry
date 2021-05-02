// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.8
// source: v1/enums.proto

package v1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Enum1 int32

const (
	Enum1_ZERO  Enum1 = 0
	Enum1_ONE   Enum1 = 1
	Enum1_TWO   Enum1 = 2
	Enum1_THREE Enum1 = 3
	Enum1_FOUR  Enum1 = 4
)

// Enum value maps for Enum1.
var (
	Enum1_name = map[int32]string{
		0: "ZERO",
		1: "ONE",
		2: "TWO",
		3: "THREE",
		4: "FOUR",
	}
	Enum1_value = map[string]int32{
		"ZERO":  0,
		"ONE":   1,
		"TWO":   2,
		"THREE": 3,
		"FOUR":  4,
	}
)

func (x Enum1) Enum() *Enum1 {
	p := new(Enum1)
	*p = x
	return p
}

func (x Enum1) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Enum1) Descriptor() protoreflect.EnumDescriptor {
	return file_v1_enums_proto_enumTypes[0].Descriptor()
}

func (Enum1) Type() protoreflect.EnumType {
	return &file_v1_enums_proto_enumTypes[0]
}

func (x Enum1) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Enum1.Descriptor instead.
func (Enum1) EnumDescriptor() ([]byte, []int) {
	return file_v1_enums_proto_rawDescGZIP(), []int{0}
}

type EnumMessage_Enum2 int32

const (
	EnumMessage_STUFF EnumMessage_Enum2 = 0
	EnumMessage_PIE   EnumMessage_Enum2 = 1
)

// Enum value maps for EnumMessage_Enum2.
var (
	EnumMessage_Enum2_name = map[int32]string{
		0: "STUFF",
		1: "PIE",
	}
	EnumMessage_Enum2_value = map[string]int32{
		"STUFF": 0,
		"PIE":   1,
	}
)

func (x EnumMessage_Enum2) Enum() *EnumMessage_Enum2 {
	p := new(EnumMessage_Enum2)
	*p = x
	return p
}

func (x EnumMessage_Enum2) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (EnumMessage_Enum2) Descriptor() protoreflect.EnumDescriptor {
	return file_v1_enums_proto_enumTypes[1].Descriptor()
}

func (EnumMessage_Enum2) Type() protoreflect.EnumType {
	return &file_v1_enums_proto_enumTypes[1]
}

func (x EnumMessage_Enum2) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use EnumMessage_Enum2.Descriptor instead.
func (EnumMessage_Enum2) EnumDescriptor() ([]byte, []int) {
	return file_v1_enums_proto_rawDescGZIP(), []int{0, 0}
}

type EnumMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Enum1Message Enum1             `protobuf:"varint,1,opt,name=enum1_message,json=enum1Message,proto3,enum=v1.Enum1" json:"enum1_message,omitempty"`
	Enum2Message EnumMessage_Enum2 `protobuf:"varint,2,opt,name=enum2_message,json=enum2Message,proto3,enum=v1.EnumMessage_Enum2" json:"enum2_message,omitempty"`
}

func (x *EnumMessage) Reset() {
	*x = EnumMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_enums_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EnumMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EnumMessage) ProtoMessage() {}

func (x *EnumMessage) ProtoReflect() protoreflect.Message {
	mi := &file_v1_enums_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EnumMessage.ProtoReflect.Descriptor instead.
func (*EnumMessage) Descriptor() ([]byte, []int) {
	return file_v1_enums_proto_rawDescGZIP(), []int{0}
}

func (x *EnumMessage) GetEnum1Message() Enum1 {
	if x != nil {
		return x.Enum1Message
	}
	return Enum1_ZERO
}

func (x *EnumMessage) GetEnum2Message() EnumMessage_Enum2 {
	if x != nil {
		return x.Enum2Message
	}
	return EnumMessage_STUFF
}

var File_v1_enums_proto protoreflect.FileDescriptor

var file_v1_enums_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x76, 0x31, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x02, 0x76, 0x31, 0x22, 0x96, 0x01, 0x0a, 0x0b, 0x45, 0x6e, 0x75, 0x6d, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x12, 0x2e, 0x0a, 0x0d, 0x65, 0x6e, 0x75, 0x6d, 0x31, 0x5f, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x09, 0x2e, 0x76, 0x31,
	0x2e, 0x45, 0x6e, 0x75, 0x6d, 0x31, 0x52, 0x0c, 0x65, 0x6e, 0x75, 0x6d, 0x31, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x12, 0x3a, 0x0a, 0x0d, 0x65, 0x6e, 0x75, 0x6d, 0x32, 0x5f, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x15, 0x2e, 0x76, 0x31,
	0x2e, 0x45, 0x6e, 0x75, 0x6d, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x45, 0x6e, 0x75,
	0x6d, 0x32, 0x52, 0x0c, 0x65, 0x6e, 0x75, 0x6d, 0x32, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x22, 0x1b, 0x0a, 0x05, 0x45, 0x6e, 0x75, 0x6d, 0x32, 0x12, 0x09, 0x0a, 0x05, 0x53, 0x54, 0x55,
	0x46, 0x46, 0x10, 0x00, 0x12, 0x07, 0x0a, 0x03, 0x50, 0x49, 0x45, 0x10, 0x01, 0x2a, 0x38, 0x0a,
	0x05, 0x45, 0x6e, 0x75, 0x6d, 0x31, 0x12, 0x08, 0x0a, 0x04, 0x5a, 0x45, 0x52, 0x4f, 0x10, 0x00,
	0x12, 0x07, 0x0a, 0x03, 0x4f, 0x4e, 0x45, 0x10, 0x01, 0x12, 0x07, 0x0a, 0x03, 0x54, 0x57, 0x4f,
	0x10, 0x02, 0x12, 0x09, 0x0a, 0x05, 0x54, 0x48, 0x52, 0x45, 0x45, 0x10, 0x03, 0x12, 0x08, 0x0a,
	0x04, 0x46, 0x4f, 0x55, 0x52, 0x10, 0x04, 0x42, 0x4a, 0x5a, 0x48, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x79, 0x6e, 0x63, 0x72, 0x6f, 0x6d, 0x61, 0x74, 0x69,
	0x63, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2d, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2d,
	0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61,
	0x6c, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_v1_enums_proto_rawDescOnce sync.Once
	file_v1_enums_proto_rawDescData = file_v1_enums_proto_rawDesc
)

func file_v1_enums_proto_rawDescGZIP() []byte {
	file_v1_enums_proto_rawDescOnce.Do(func() {
		file_v1_enums_proto_rawDescData = protoimpl.X.CompressGZIP(file_v1_enums_proto_rawDescData)
	})
	return file_v1_enums_proto_rawDescData
}

var file_v1_enums_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_v1_enums_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_v1_enums_proto_goTypes = []interface{}{
	(Enum1)(0),             // 0: v1.Enum1
	(EnumMessage_Enum2)(0), // 1: v1.EnumMessage.Enum2
	(*EnumMessage)(nil),    // 2: v1.EnumMessage
}
var file_v1_enums_proto_depIdxs = []int32{
	0, // 0: v1.EnumMessage.enum1_message:type_name -> v1.Enum1
	1, // 1: v1.EnumMessage.enum2_message:type_name -> v1.EnumMessage.Enum2
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_v1_enums_proto_init() }
func file_v1_enums_proto_init() {
	if File_v1_enums_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_v1_enums_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EnumMessage); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_v1_enums_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_v1_enums_proto_goTypes,
		DependencyIndexes: file_v1_enums_proto_depIdxs,
		EnumInfos:         file_v1_enums_proto_enumTypes,
		MessageInfos:      file_v1_enums_proto_msgTypes,
	}.Build()
	File_v1_enums_proto = out.File
	file_v1_enums_proto_rawDesc = nil
	file_v1_enums_proto_goTypes = nil
	file_v1_enums_proto_depIdxs = nil
}
