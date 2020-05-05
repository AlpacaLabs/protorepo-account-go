// Code generated by protoc-gen-go. DO NOT EDIT.
// source: alpacalabs/account/v1/phone_number.proto

package v1

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type PhoneNumber struct {
	Id                   string               `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	CreatedAt            *timestamp.Timestamp `protobuf:"bytes,2,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	AccountId            string               `protobuf:"bytes,3,opt,name=account_id,json=accountId,proto3" json:"account_id,omitempty"`
	PhoneNumber          string               `protobuf:"bytes,4,opt,name=phone_number,json=phoneNumber,proto3" json:"phone_number,omitempty"`
	Confirmed            bool                 `protobuf:"varint,5,opt,name=confirmed,proto3" json:"confirmed,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *PhoneNumber) Reset()         { *m = PhoneNumber{} }
func (m *PhoneNumber) String() string { return proto.CompactTextString(m) }
func (*PhoneNumber) ProtoMessage()    {}
func (*PhoneNumber) Descriptor() ([]byte, []int) {
	return fileDescriptor_132c973f88ab5901, []int{0}
}

func (m *PhoneNumber) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PhoneNumber.Unmarshal(m, b)
}
func (m *PhoneNumber) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PhoneNumber.Marshal(b, m, deterministic)
}
func (m *PhoneNumber) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PhoneNumber.Merge(m, src)
}
func (m *PhoneNumber) XXX_Size() int {
	return xxx_messageInfo_PhoneNumber.Size(m)
}
func (m *PhoneNumber) XXX_DiscardUnknown() {
	xxx_messageInfo_PhoneNumber.DiscardUnknown(m)
}

var xxx_messageInfo_PhoneNumber proto.InternalMessageInfo

func (m *PhoneNumber) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *PhoneNumber) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *PhoneNumber) GetAccountId() string {
	if m != nil {
		return m.AccountId
	}
	return ""
}

func (m *PhoneNumber) GetPhoneNumber() string {
	if m != nil {
		return m.PhoneNumber
	}
	return ""
}

func (m *PhoneNumber) GetConfirmed() bool {
	if m != nil {
		return m.Confirmed
	}
	return false
}

func init() {
	proto.RegisterType((*PhoneNumber)(nil), "alpacalabs.account.v1.PhoneNumber")
}

func init() {
	proto.RegisterFile("alpacalabs/account/v1/phone_number.proto", fileDescriptor_132c973f88ab5901)
}

var fileDescriptor_132c973f88ab5901 = []byte{
	// 303 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0xb1, 0x6a, 0xf3, 0x30,
	0x1c, 0xc4, 0x91, 0xf3, 0x7d, 0xa5, 0x56, 0x42, 0x07, 0x43, 0xc0, 0x0d, 0x2d, 0x4d, 0x3b, 0x79,
	0x89, 0x84, 0xdb, 0xa9, 0x5b, 0x9d, 0xad, 0x10, 0x4a, 0x30, 0x25, 0x43, 0x09, 0x18, 0x59, 0x52,
	0x1c, 0x81, 0x65, 0x09, 0x47, 0xce, 0x03, 0xb5, 0x5b, 0x9f, 0xa0, 0xcf, 0xd0, 0xa7, 0x2a, 0x96,
	0x15, 0xec, 0x21, 0x9b, 0x38, 0xee, 0x7e, 0xba, 0xfb, 0xc3, 0x88, 0x94, 0x9a, 0x50, 0x52, 0x92,
	0xfc, 0x80, 0x09, 0xa5, 0xaa, 0xa9, 0x0c, 0x3e, 0xc6, 0x58, 0xef, 0x55, 0xc5, 0xb3, 0xaa, 0x91,
	0x39, 0xaf, 0x91, 0xae, 0x95, 0x51, 0xc1, 0xb4, 0x77, 0x22, 0xe7, 0x44, 0xc7, 0x78, 0x76, 0x57,
	0x28, 0x55, 0x94, 0x1c, 0x5b, 0x53, 0xde, 0xec, 0xb0, 0x11, 0x92, 0x1f, 0x0c, 0x91, 0xba, 0xcb,
	0x3d, 0xfc, 0x00, 0x38, 0x5e, 0xb7, 0xb8, 0x37, 0x4b, 0x0b, 0xae, 0xa0, 0x27, 0x58, 0x08, 0xe6,
	0x20, 0xf2, 0x53, 0x4f, 0xb0, 0xe0, 0x19, 0x42, 0x5a, 0x73, 0x62, 0x38, 0xcb, 0x88, 0x09, 0xbd,
	0x39, 0x88, 0xc6, 0x8f, 0x33, 0xd4, 0x51, 0xd1, 0x89, 0x8a, 0xde, 0x4f, 0xd4, 0xd4, 0x77, 0xee,
	0xc4, 0x04, 0xb7, 0x10, 0xba, 0x26, 0x99, 0x60, 0xe1, 0xc8, 0x22, 0x7d, 0xa7, 0xbc, 0xb2, 0xe0,
	0x1e, 0x4e, 0x86, 0x3b, 0xc2, 0x7f, 0xd6, 0x30, 0xd6, 0x83, 0x32, 0x37, 0xd0, 0xa7, 0xaa, 0xda,
	0x89, 0x5a, 0x72, 0x16, 0xfe, 0x9f, 0x83, 0xe8, 0x32, 0xed, 0x85, 0xe5, 0x17, 0x80, 0xd7, 0x54,
	0x49, 0x74, 0x76, 0xf9, 0x72, 0x92, 0x74, 0xef, 0x75, 0xdb, 0x71, 0x0d, 0x3e, 0x5e, 0x0a, 0x61,
	0xf6, 0x4d, 0x8e, 0xa8, 0x92, 0x38, 0xb1, 0x89, 0x55, 0x7b, 0x55, 0x3b, 0xa1, 0xe6, 0x5a, 0x2d,
	0x5c, 0x76, 0x51, 0x28, 0x7c, 0xf6, 0xea, 0x9f, 0xde, 0x28, 0x59, 0x25, 0xdf, 0xde, 0xb4, 0x4f,
	0x23, 0xf7, 0x07, 0xda, 0xc4, 0xbf, 0x43, 0x7d, 0xeb, 0xf4, 0xed, 0x26, 0xce, 0x2f, 0xec, 0x0f,
	0x4f, 0x7f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x28, 0x53, 0x32, 0x2e, 0xcb, 0x01, 0x00, 0x00,
}
