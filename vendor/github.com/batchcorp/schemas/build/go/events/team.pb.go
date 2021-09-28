// Code generated by protoc-gen-go. DO NOT EDIT.
// source: team.proto

package events

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

type Team struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Account              *Account `protobuf:"bytes,3,opt,name=account,proto3" json:"account,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Team) Reset()         { *m = Team{} }
func (m *Team) String() string { return proto.CompactTextString(m) }
func (*Team) ProtoMessage()    {}
func (*Team) Descriptor() ([]byte, []int) {
	return fileDescriptor_8b4e9e93d7b2c6bb, []int{0}
}

func (m *Team) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Team.Unmarshal(m, b)
}
func (m *Team) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Team.Marshal(b, m, deterministic)
}
func (m *Team) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Team.Merge(m, src)
}
func (m *Team) XXX_Size() int {
	return xxx_messageInfo_Team.Size(m)
}
func (m *Team) XXX_DiscardUnknown() {
	xxx_messageInfo_Team.DiscardUnknown(m)
}

var xxx_messageInfo_Team proto.InternalMessageInfo

func (m *Team) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Team) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Team) GetAccount() *Account {
	if m != nil {
		return m.Account
	}
	return nil
}

func init() {
	proto.RegisterType((*Team)(nil), "events.Team")
}

func init() { proto.RegisterFile("team.proto", fileDescriptor_8b4e9e93d7b2c6bb) }

var fileDescriptor_8b4e9e93d7b2c6bb = []byte{
	// 166 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2a, 0x49, 0x4d, 0xcc,
	0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x4b, 0x2d, 0x4b, 0xcd, 0x2b, 0x29, 0x96, 0xe2,
	0x4d, 0x4c, 0x4e, 0xce, 0x2f, 0xcd, 0x2b, 0x81, 0x08, 0x2b, 0x85, 0x72, 0xb1, 0x84, 0xa4, 0x26,
	0xe6, 0x0a, 0xf1, 0x71, 0x31, 0x65, 0xa6, 0x48, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x06, 0x31, 0x65,
	0xa6, 0x08, 0x09, 0x71, 0xb1, 0xe4, 0x25, 0xe6, 0xa6, 0x4a, 0x30, 0x81, 0x45, 0xc0, 0x6c, 0x21,
	0x4d, 0x2e, 0x76, 0xa8, 0x66, 0x09, 0x66, 0x05, 0x46, 0x0d, 0x6e, 0x23, 0x7e, 0x3d, 0x88, 0xa1,
	0x7a, 0x8e, 0x10, 0xe1, 0x20, 0x98, 0xbc, 0x93, 0x5e, 0x94, 0x4e, 0x7a, 0x66, 0x49, 0x46, 0x69,
	0x92, 0x5e, 0x72, 0x7e, 0xae, 0x7e, 0x52, 0x62, 0x49, 0x72, 0x46, 0x72, 0x7e, 0x51, 0x81, 0x7e,
	0x71, 0x72, 0x46, 0x6a, 0x6e, 0x62, 0xb1, 0x7e, 0x52, 0x69, 0x66, 0x4e, 0x8a, 0x7e, 0x7a, 0xbe,
	0x3e, 0xc4, 0x80, 0x24, 0x36, 0xb0, 0x6b, 0x8c, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0x8f, 0xc5,
	0x25, 0xd4, 0xb2, 0x00, 0x00, 0x00,
}