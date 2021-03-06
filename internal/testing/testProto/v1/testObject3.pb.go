// Code generated by protoc-gen-go. DO NOT EDIT.
// source: v1/testObject3.proto

package v1

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type TestObject3 struct {
	// Types that are valid to be assigned to Request:
	//	*TestObject3_RequestTestObject
	//	*TestObject3_RequestTestObject_2
	//	*TestObject3_RequestString
	Request isTestObject3_Request `protobuf_oneof:"request"`
	Bla     string                `protobuf:"bytes,3,opt,name=bla,proto3" json:"bla,omitempty"`
	// Types that are valid to be assigned to Request2:
	//	*TestObject3_Request2String
	//	*TestObject3_Request2Int32
	Request2             isTestObject3_Request2 `protobuf_oneof:"request2"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *TestObject3) Reset()         { *m = TestObject3{} }
func (m *TestObject3) String() string { return proto.CompactTextString(m) }
func (*TestObject3) ProtoMessage()    {}
func (*TestObject3) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b3ae9f05d77f413, []int{0}
}

func (m *TestObject3) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TestObject3.Unmarshal(m, b)
}
func (m *TestObject3) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TestObject3.Marshal(b, m, deterministic)
}
func (m *TestObject3) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TestObject3.Merge(m, src)
}
func (m *TestObject3) XXX_Size() int {
	return xxx_messageInfo_TestObject3.Size(m)
}
func (m *TestObject3) XXX_DiscardUnknown() {
	xxx_messageInfo_TestObject3.DiscardUnknown(m)
}

var xxx_messageInfo_TestObject3 proto.InternalMessageInfo

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

func (m *TestObject3) GetRequest() isTestObject3_Request {
	if m != nil {
		return m.Request
	}
	return nil
}

func (m *TestObject3) GetRequestTestObject() *TestObject {
	if x, ok := m.GetRequest().(*TestObject3_RequestTestObject); ok {
		return x.RequestTestObject
	}
	return nil
}

func (m *TestObject3) GetRequestTestObject_2() *TestObject2 {
	if x, ok := m.GetRequest().(*TestObject3_RequestTestObject_2); ok {
		return x.RequestTestObject_2
	}
	return nil
}

func (m *TestObject3) GetRequestString() string {
	if x, ok := m.GetRequest().(*TestObject3_RequestString); ok {
		return x.RequestString
	}
	return ""
}

func (m *TestObject3) GetBla() string {
	if m != nil {
		return m.Bla
	}
	return ""
}

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

func (m *TestObject3) GetRequest2() isTestObject3_Request2 {
	if m != nil {
		return m.Request2
	}
	return nil
}

func (m *TestObject3) GetRequest2String() string {
	if x, ok := m.GetRequest2().(*TestObject3_Request2String); ok {
		return x.Request2String
	}
	return ""
}

func (m *TestObject3) GetRequest2Int32() int32 {
	if x, ok := m.GetRequest2().(*TestObject3_Request2Int32); ok {
		return x.Request2Int32
	}
	return 0
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*TestObject3) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*TestObject3_RequestTestObject)(nil),
		(*TestObject3_RequestTestObject_2)(nil),
		(*TestObject3_RequestString)(nil),
		(*TestObject3_Request2String)(nil),
		(*TestObject3_Request2Int32)(nil),
	}
}

func init() {
	proto.RegisterType((*TestObject3)(nil), "v1.TestObject3")
}

func init() { proto.RegisterFile("v1/testObject3.proto", fileDescriptor_9b3ae9f05d77f413) }

var fileDescriptor_9b3ae9f05d77f413 = []byte{
	// 234 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x29, 0x33, 0xd4, 0x2f,
	0x49, 0x2d, 0x2e, 0xf1, 0x4f, 0xca, 0x4a, 0x4d, 0x2e, 0x31, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9,
	0x17, 0x62, 0x2a, 0x33, 0x94, 0x12, 0x46, 0x91, 0x81, 0x48, 0x48, 0xa1, 0x2a, 0x37, 0x82, 0x88,
	0x2a, 0xed, 0x66, 0xe2, 0xe2, 0x0e, 0x41, 0x18, 0x22, 0xe4, 0xc0, 0x25, 0x5c, 0x94, 0x5a, 0x58,
	0x9a, 0x5a, 0x5c, 0x12, 0x0f, 0x52, 0x1c, 0x9f, 0x0f, 0x16, 0x97, 0x60, 0x54, 0x60, 0xd4, 0xe0,
	0x36, 0xe2, 0xd3, 0x2b, 0x33, 0xd4, 0x43, 0xa8, 0xf6, 0x60, 0x08, 0x12, 0x84, 0x2a, 0x46, 0x08,
	0x0a, 0xb9, 0x70, 0x89, 0x62, 0x31, 0x21, 0xde, 0x48, 0x82, 0x09, 0x6c, 0x06, 0x3f, 0xaa, 0x19,
	0x46, 0x1e, 0x0c, 0x41, 0x42, 0x18, 0x86, 0x18, 0x09, 0xa9, 0x73, 0xf1, 0xc1, 0x4c, 0x29, 0x2e,
	0x29, 0xca, 0xcc, 0x4b, 0x97, 0x60, 0x51, 0x60, 0xd4, 0xe0, 0xf4, 0x60, 0x08, 0xe2, 0x85, 0x8a,
	0x07, 0x83, 0x85, 0x85, 0x04, 0xb8, 0x98, 0x93, 0x72, 0x12, 0x25, 0x98, 0x41, 0xb2, 0x41, 0x20,
	0xa6, 0x90, 0x26, 0x17, 0x3f, 0x54, 0x89, 0x11, 0x4c, 0x2f, 0x2b, 0x58, 0x2f, 0x63, 0x10, 0xcc,
	0x4c, 0x23, 0xa8, 0x66, 0x84, 0x2d, 0x46, 0xf1, 0x99, 0x79, 0x25, 0xc6, 0x46, 0x12, 0x6c, 0x0a,
	0x8c, 0x1a, 0xac, 0x1e, 0x8c, 0x70, 0x5b, 0x8c, 0x3c, 0x41, 0xc2, 0x4e, 0x9c, 0x5c, 0xec, 0x50,
	0x01, 0x27, 0x2e, 0x2e, 0x0e, 0x98, 0x5c, 0x12, 0x1b, 0x38, 0x10, 0x8d, 0x01, 0x01, 0x00, 0x00,
	0xff, 0xff, 0x6c, 0x02, 0xd9, 0x07, 0x8b, 0x01, 0x00, 0x00,
}
