// Code generated by protoc-gen-go. DO NOT EDIT.
// source: user.proto

package pb

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

type UserInfoReq struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserInfoReq) Reset()         { *m = UserInfoReq{} }
func (m *UserInfoReq) String() string { return proto.CompactTextString(m) }
func (*UserInfoReq) ProtoMessage()    {}
func (*UserInfoReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{0}
}

func (m *UserInfoReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserInfoReq.Unmarshal(m, b)
}
func (m *UserInfoReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserInfoReq.Marshal(b, m, deterministic)
}
func (m *UserInfoReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserInfoReq.Merge(m, src)
}
func (m *UserInfoReq) XXX_Size() int {
	return xxx_messageInfo_UserInfoReq.Size(m)
}
func (m *UserInfoReq) XXX_DiscardUnknown() {
	xxx_messageInfo_UserInfoReq.DiscardUnknown(m)
}

var xxx_messageInfo_UserInfoReq proto.InternalMessageInfo

func (m *UserInfoReq) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type UserInfoRsp struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Hobby                []string `protobuf:"bytes,3,rep,name=hobby,proto3" json:"hobby,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserInfoRsp) Reset()         { *m = UserInfoRsp{} }
func (m *UserInfoRsp) String() string { return proto.CompactTextString(m) }
func (*UserInfoRsp) ProtoMessage()    {}
func (*UserInfoRsp) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{1}
}

func (m *UserInfoRsp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserInfoRsp.Unmarshal(m, b)
}
func (m *UserInfoRsp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserInfoRsp.Marshal(b, m, deterministic)
}
func (m *UserInfoRsp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserInfoRsp.Merge(m, src)
}
func (m *UserInfoRsp) XXX_Size() int {
	return xxx_messageInfo_UserInfoRsp.Size(m)
}
func (m *UserInfoRsp) XXX_DiscardUnknown() {
	xxx_messageInfo_UserInfoRsp.DiscardUnknown(m)
}

var xxx_messageInfo_UserInfoRsp proto.InternalMessageInfo

func (m *UserInfoRsp) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *UserInfoRsp) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *UserInfoRsp) GetHobby() []string {
	if m != nil {
		return m.Hobby
	}
	return nil
}

func init() {
	proto.RegisterType((*UserInfoReq)(nil), "pb.UserInfoReq")
	proto.RegisterType((*UserInfoRsp)(nil), "pb.UserInfoRsp")
}

func init() { proto.RegisterFile("user.proto", fileDescriptor_116e343673f7ffaf) }

var fileDescriptor_116e343673f7ffaf = []byte{
	// 150 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2a, 0x2d, 0x4e, 0x2d,
	0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2a, 0x48, 0x52, 0x52, 0xe4, 0xe2, 0x0e, 0x2d,
	0x4e, 0x2d, 0xf2, 0xcc, 0x4b, 0xcb, 0x0f, 0x4a, 0x2d, 0x14, 0x12, 0xe2, 0x62, 0xc9, 0x4b, 0xcc,
	0x4d, 0x95, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x02, 0xb3, 0x95, 0xdc, 0x91, 0x94, 0x14, 0x17,
	0x08, 0xf1, 0x71, 0x31, 0x65, 0xa6, 0x80, 0x15, 0x30, 0x07, 0x31, 0x65, 0xa6, 0xc0, 0xb5, 0x30,
	0x21, 0xb4, 0x08, 0x89, 0x70, 0xb1, 0x66, 0xe4, 0x27, 0x25, 0x55, 0x4a, 0x30, 0x2b, 0x30, 0x6b,
	0x70, 0x06, 0x41, 0x38, 0x46, 0x96, 0x5c, 0x2c, 0x20, 0x83, 0x84, 0x0c, 0xb9, 0xb8, 0xd3, 0x53,
	0x4b, 0x60, 0x66, 0x0a, 0xf1, 0xeb, 0x15, 0x24, 0xe9, 0x21, 0x39, 0x42, 0x0a, 0x55, 0xa0, 0xb8,
	0x40, 0x89, 0x21, 0x89, 0x0d, 0xec, 0x62, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0xdb, 0x22,
	0x17, 0x6e, 0xbf, 0x00, 0x00, 0x00,
}
