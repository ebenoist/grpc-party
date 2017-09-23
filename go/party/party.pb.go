// Code generated by protoc-gen-go. DO NOT EDIT.
// source: party.proto

/*
Package party is a generated protocol buffer package.

It is generated from these files:
	party.proto

It has these top-level messages:
	User
	PartyResp
*/
package party

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type User struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *User) Reset()                    { *m = User{} }
func (m *User) String() string            { return proto.CompactTextString(m) }
func (*User) ProtoMessage()               {}
func (*User) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *User) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type PartyResp struct {
	Msg string `protobuf:"bytes,1,opt,name=msg" json:"msg,omitempty"`
}

func (m *PartyResp) Reset()                    { *m = PartyResp{} }
func (m *PartyResp) String() string            { return proto.CompactTextString(m) }
func (*PartyResp) ProtoMessage()               {}
func (*PartyResp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *PartyResp) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func init() {
	proto.RegisterType((*User)(nil), "User")
	proto.RegisterType((*PartyResp)(nil), "PartyResp")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Party service

type PartyClient interface {
	PartyHard(ctx context.Context, in *User, opts ...grpc.CallOption) (*PartyResp, error)
}

type partyClient struct {
	cc *grpc.ClientConn
}

func NewPartyClient(cc *grpc.ClientConn) PartyClient {
	return &partyClient{cc}
}

func (c *partyClient) PartyHard(ctx context.Context, in *User, opts ...grpc.CallOption) (*PartyResp, error) {
	out := new(PartyResp)
	err := grpc.Invoke(ctx, "/Party/PartyHard", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Party service

type PartyServer interface {
	PartyHard(context.Context, *User) (*PartyResp, error)
}

func RegisterPartyServer(s *grpc.Server, srv PartyServer) {
	s.RegisterService(&_Party_serviceDesc, srv)
}

func _Party_PartyHard_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PartyServer).PartyHard(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Party/PartyHard",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PartyServer).PartyHard(ctx, req.(*User))
	}
	return interceptor(ctx, in, info, handler)
}

var _Party_serviceDesc = grpc.ServiceDesc{
	ServiceName: "Party",
	HandlerType: (*PartyServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PartyHard",
			Handler:    _Party_PartyHard_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "party.proto",
}

func init() { proto.RegisterFile("party.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 115 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2e, 0x48, 0x2c, 0x2a,
	0xa9, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x57, 0x92, 0xe2, 0x62, 0x09, 0x2d, 0x4e, 0x2d, 0x12,
	0x12, 0xe2, 0x62, 0xc9, 0x4b, 0xcc, 0x4d, 0x95, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x02, 0xb3,
	0x95, 0x64, 0xb9, 0x38, 0x03, 0x40, 0x4a, 0x83, 0x52, 0x8b, 0x0b, 0x84, 0x04, 0xb8, 0x98, 0x73,
	0x8b, 0xd3, 0xa1, 0xf2, 0x20, 0xa6, 0x91, 0x26, 0x17, 0x2b, 0x58, 0x5a, 0x48, 0x01, 0xaa, 0xce,
	0x23, 0xb1, 0x28, 0x45, 0x88, 0x55, 0x0f, 0x64, 0x9e, 0x14, 0x97, 0x1e, 0x5c, 0xab, 0x12, 0x43,
	0x12, 0x1b, 0xd8, 0x32, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0x09, 0xcf, 0x2c, 0x77, 0x7b,
	0x00, 0x00, 0x00,
}
