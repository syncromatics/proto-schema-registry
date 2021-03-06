// Code generated by protoc-gen-go. DO NOT EDIT.
// source: v1/testObject2.proto

package v1

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
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

type TestObject2 struct {
	Strings              []*wrappers.StringValue `protobuf:"bytes,1,rep,name=strings,proto3" json:"strings,omitempty"`
	MapObjects           map[string]*TestObject  `protobuf:"bytes,2,rep,name=map_objects,json=mapObjects,proto3" json:"map_objects,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	StringMaps           map[int32]string        `protobuf:"bytes,3,rep,name=string_maps,json=stringMaps,proto3" json:"string_maps,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *TestObject2) Reset()         { *m = TestObject2{} }
func (m *TestObject2) String() string { return proto.CompactTextString(m) }
func (*TestObject2) ProtoMessage()    {}
func (*TestObject2) Descriptor() ([]byte, []int) {
	return fileDescriptor_7fc4307b3585cf93, []int{0}
}

func (m *TestObject2) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TestObject2.Unmarshal(m, b)
}
func (m *TestObject2) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TestObject2.Marshal(b, m, deterministic)
}
func (m *TestObject2) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TestObject2.Merge(m, src)
}
func (m *TestObject2) XXX_Size() int {
	return xxx_messageInfo_TestObject2.Size(m)
}
func (m *TestObject2) XXX_DiscardUnknown() {
	xxx_messageInfo_TestObject2.DiscardUnknown(m)
}

var xxx_messageInfo_TestObject2 proto.InternalMessageInfo

func (m *TestObject2) GetStrings() []*wrappers.StringValue {
	if m != nil {
		return m.Strings
	}
	return nil
}

func (m *TestObject2) GetMapObjects() map[string]*TestObject {
	if m != nil {
		return m.MapObjects
	}
	return nil
}

func (m *TestObject2) GetStringMaps() map[int32]string {
	if m != nil {
		return m.StringMaps
	}
	return nil
}

func init() {
	proto.RegisterType((*TestObject2)(nil), "v1.TestObject2")
	proto.RegisterMapType((map[string]*TestObject)(nil), "v1.TestObject2.MapObjectsEntry")
	proto.RegisterMapType((map[int32]string)(nil), "v1.TestObject2.StringMapsEntry")
}

func init() { proto.RegisterFile("v1/testObject2.proto", fileDescriptor_7fc4307b3585cf93) }

var fileDescriptor_7fc4307b3585cf93 = []byte{
	// 258 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0xd0, 0x41, 0x4b, 0xc3, 0x30,
	0x14, 0x07, 0x70, 0x92, 0x32, 0x65, 0xaf, 0xe0, 0x24, 0xee, 0x10, 0x8a, 0x68, 0x11, 0x0f, 0x3d,
	0xa5, 0xb4, 0x82, 0x88, 0x20, 0x78, 0xf1, 0x58, 0x84, 0x2a, 0x5e, 0x47, 0x2a, 0xb1, 0xa8, 0xeb,
	0x12, 0x92, 0xac, 0xb2, 0x4f, 0xea, 0xd7, 0x91, 0x26, 0x74, 0x66, 0x63, 0xb7, 0xf0, 0x7f, 0xef,
	0xfd, 0xfe, 0xa5, 0x30, 0xef, 0x8b, 0xdc, 0x0a, 0x63, 0x9f, 0x9b, 0x2f, 0xf1, 0x6e, 0x4b, 0xa6,
	0xb4, 0xb4, 0x92, 0xe0, 0xbe, 0x48, 0x2e, 0x5a, 0x29, 0xdb, 0xa5, 0xc8, 0x5d, 0xd2, 0xac, 0x3f,
	0xf2, 0x1f, 0xcd, 0x95, 0x12, 0xda, 0xf8, 0x9d, 0xe4, 0x6c, 0xe7, 0xd2, 0x87, 0x57, 0xbf, 0x18,
	0xe2, 0xd7, 0x7f, 0x8e, 0xdc, 0xc2, 0xb1, 0xb1, 0xfa, 0x73, 0xd5, 0x1a, 0x8a, 0xd2, 0x28, 0x8b,
	0xcb, 0x73, 0xe6, 0x59, 0x36, 0xb2, 0xec, 0xc5, 0xcd, 0xdf, 0xf8, 0x72, 0x2d, 0xea, 0x71, 0x99,
	0x3c, 0x42, 0xdc, 0x71, 0xb5, 0x90, 0x8e, 0x31, 0x14, 0xbb, 0xdb, 0x4b, 0xd6, 0x17, 0x2c, 0xd0,
	0x59, 0xc5, 0x95, 0x7f, 0x9a, 0xa7, 0x95, 0xd5, 0x9b, 0x1a, 0xba, 0x6d, 0x30, 0x08, 0x1e, 0x5b,
	0x74, 0x5c, 0x19, 0x1a, 0x1d, 0x16, 0x7c, 0x79, 0xc5, 0xd5, 0x28, 0x98, 0x6d, 0x90, 0x54, 0x30,
	0xdb, 0x2b, 0x20, 0xa7, 0x10, 0x7d, 0x8b, 0x0d, 0x45, 0x29, 0xca, 0xa6, 0xf5, 0xf0, 0x24, 0xd7,
	0x30, 0xe9, 0x87, 0x4f, 0xa7, 0x38, 0x45, 0x59, 0x5c, 0x9e, 0xec, 0x16, 0xd4, 0x7e, 0x78, 0x8f,
	0xef, 0x50, 0xf2, 0x00, 0xb3, 0xbd, 0xb6, 0x90, 0x9b, 0x78, 0x6e, 0x1e, 0x72, 0xd3, 0xe0, 0xbc,
	0x39, 0x72, 0x3f, 0xec, 0xe6, 0x2f, 0x00, 0x00, 0xff, 0xff, 0xe5, 0xc9, 0x53, 0x3d, 0xb1, 0x01,
	0x00, 0x00,
}
