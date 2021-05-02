// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.8
// source: v1/testObject3.proto

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

type TestObject3 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Request:
	//	*TestObject3_RequestTestObject
	//	*TestObject3_RequestTestObject_2
	//	*TestObject3_RequestString
	Request isTestObject3_Request `protobuf_oneof:"request"`
	Bla     string                `protobuf:"bytes,3,opt,name=bla,proto3" json:"bla,omitempty"`
	// Types that are assignable to Request2:
	//	*TestObject3_Request2String
	//	*TestObject3_Request2Int32
	Request2 isTestObject3_Request2 `protobuf_oneof:"request2"`
}

func (x *TestObject3) Reset() {
	*x = TestObject3{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_testObject3_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TestObject3) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TestObject3) ProtoMessage() {}

func (x *TestObject3) ProtoReflect() protoreflect.Message {
	mi := &file_v1_testObject3_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TestObject3.ProtoReflect.Descriptor instead.
func (*TestObject3) Descriptor() ([]byte, []int) {
	return file_v1_testObject3_proto_rawDescGZIP(), []int{0}
}

func (m *TestObject3) GetRequest() isTestObject3_Request {
	if m != nil {
		return m.Request
	}
	return nil
}

func (x *TestObject3) GetRequestTestObject() *TestObject {
	if x, ok := x.GetRequest().(*TestObject3_RequestTestObject); ok {
		return x.RequestTestObject
	}
	return nil
}

func (x *TestObject3) GetRequestTestObject_2() *TestObject2 {
	if x, ok := x.GetRequest().(*TestObject3_RequestTestObject_2); ok {
		return x.RequestTestObject_2
	}
	return nil
}

func (x *TestObject3) GetRequestString() string {
	if x, ok := x.GetRequest().(*TestObject3_RequestString); ok {
		return x.RequestString
	}
	return ""
}

func (x *TestObject3) GetBla() string {
	if x != nil {
		return x.Bla
	}
	return ""
}

func (m *TestObject3) GetRequest2() isTestObject3_Request2 {
	if m != nil {
		return m.Request2
	}
	return nil
}

func (x *TestObject3) GetRequest2String() string {
	if x, ok := x.GetRequest2().(*TestObject3_Request2String); ok {
		return x.Request2String
	}
	return ""
}

func (x *TestObject3) GetRequest2Int32() int32 {
	if x, ok := x.GetRequest2().(*TestObject3_Request2Int32); ok {
		return x.Request2Int32
	}
	return 0
}

type isTestObject3_Request interface {
	isTestObject3_Request()
}

type TestObject3_RequestTestObject struct {
	RequestTestObject *TestObject `protobuf:"bytes,1,opt,name=request_test_object,json=requestTestObject,proto3,oneof"`
}

type TestObject3_RequestTestObject_2 struct {
	RequestTestObject_2 *TestObject2 `protobuf:"bytes,2,opt,name=request_test_object_2,json=requestTestObject2,proto3,oneof"`
}

type TestObject3_RequestString struct {
	RequestString string `protobuf:"bytes,4,opt,name=request_string,json=requestString,proto3,oneof"`
}

func (*TestObject3_RequestTestObject) isTestObject3_Request() {}

func (*TestObject3_RequestTestObject_2) isTestObject3_Request() {}

func (*TestObject3_RequestString) isTestObject3_Request() {}

type isTestObject3_Request2 interface {
	isTestObject3_Request2()
}

type TestObject3_Request2String struct {
	Request2String string `protobuf:"bytes,5,opt,name=request2_string,json=request2String,proto3,oneof"`
}

type TestObject3_Request2Int32 struct {
	Request2Int32 int32 `protobuf:"varint,6,opt,name=request2_int32,json=request2Int32,proto3,oneof"`
}

func (*TestObject3_Request2String) isTestObject3_Request2() {}

func (*TestObject3_Request2Int32) isTestObject3_Request2() {}

var File_v1_testObject3_proto protoreflect.FileDescriptor

var file_v1_testObject3_proto_rawDesc = []byte{
	0x0a, 0x14, 0x76, 0x31, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x33,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x76, 0x31, 0x1a, 0x13, 0x76, 0x31, 0x2f, 0x74,
	0x65, 0x73, 0x74, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x14, 0x76, 0x31, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x32, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xbb, 0x02, 0x0a, 0x0b, 0x54, 0x65, 0x73, 0x74, 0x4f, 0x62,
	0x6a, 0x65, 0x63, 0x74, 0x33, 0x12, 0x40, 0x0a, 0x13, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x5f, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x4f, 0x62, 0x6a, 0x65,
	0x63, 0x74, 0x48, 0x00, 0x52, 0x11, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x54, 0x65, 0x73,
	0x74, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x44, 0x0a, 0x15, 0x72, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x5f, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x32,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x65, 0x73, 0x74,
	0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x32, 0x48, 0x00, 0x52, 0x12, 0x72, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x54, 0x65, 0x73, 0x74, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x32, 0x12, 0x27, 0x0a,
	0x0e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0d, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x12, 0x10, 0x0a, 0x03, 0x62, 0x6c, 0x61, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x62, 0x6c, 0x61, 0x12, 0x29, 0x0a, 0x0f, 0x72, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x32, 0x5f, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x48, 0x01, 0x52, 0x0e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x32, 0x53, 0x74, 0x72,
	0x69, 0x6e, 0x67, 0x12, 0x27, 0x0a, 0x0e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x32, 0x5f,
	0x69, 0x6e, 0x74, 0x33, 0x32, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x48, 0x01, 0x52, 0x0d, 0x72,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x32, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x42, 0x09, 0x0a, 0x07,
	0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x42, 0x0a, 0x0a, 0x08, 0x72, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x32, 0x42, 0x4a, 0x5a, 0x48, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x73, 0x79, 0x6e, 0x63, 0x72, 0x6f, 0x6d, 0x61, 0x74, 0x69, 0x63, 0x73, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2d, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2d, 0x72, 0x65, 0x67, 0x69,
	0x73, 0x74, 0x72, 0x79, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x74, 0x65,
	0x73, 0x74, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_v1_testObject3_proto_rawDescOnce sync.Once
	file_v1_testObject3_proto_rawDescData = file_v1_testObject3_proto_rawDesc
)

func file_v1_testObject3_proto_rawDescGZIP() []byte {
	file_v1_testObject3_proto_rawDescOnce.Do(func() {
		file_v1_testObject3_proto_rawDescData = protoimpl.X.CompressGZIP(file_v1_testObject3_proto_rawDescData)
	})
	return file_v1_testObject3_proto_rawDescData
}

var file_v1_testObject3_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_v1_testObject3_proto_goTypes = []interface{}{
	(*TestObject3)(nil), // 0: v1.TestObject3
	(*TestObject)(nil),  // 1: v1.TestObject
	(*TestObject2)(nil), // 2: v1.TestObject2
}
var file_v1_testObject3_proto_depIdxs = []int32{
	1, // 0: v1.TestObject3.request_test_object:type_name -> v1.TestObject
	2, // 1: v1.TestObject3.request_test_object_2:type_name -> v1.TestObject2
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_v1_testObject3_proto_init() }
func file_v1_testObject3_proto_init() {
	if File_v1_testObject3_proto != nil {
		return
	}
	file_v1_testObject_proto_init()
	file_v1_testObject2_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_v1_testObject3_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TestObject3); i {
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
	file_v1_testObject3_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*TestObject3_RequestTestObject)(nil),
		(*TestObject3_RequestTestObject_2)(nil),
		(*TestObject3_RequestString)(nil),
		(*TestObject3_Request2String)(nil),
		(*TestObject3_Request2Int32)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_v1_testObject3_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_v1_testObject3_proto_goTypes,
		DependencyIndexes: file_v1_testObject3_proto_depIdxs,
		MessageInfos:      file_v1_testObject3_proto_msgTypes,
	}.Build()
	File_v1_testObject3_proto = out.File
	file_v1_testObject3_proto_rawDesc = nil
	file_v1_testObject3_proto_goTypes = nil
	file_v1_testObject3_proto_depIdxs = nil
}
