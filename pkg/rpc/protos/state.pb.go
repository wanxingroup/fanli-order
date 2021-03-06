// Code generated by protoc-gen-go. DO NOT EDIT.
// source: state.proto

package protos

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

type PingRequest struct {
	Message string `protobuf:"bytes,1,opt,name=message" json:"message,omitempty"`
}

func (m *PingRequest) Reset()                    { *m = PingRequest{} }
func (m *PingRequest) String() string            { return proto.CompactTextString(m) }
func (*PingRequest) ProtoMessage()               {}
func (*PingRequest) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{0} }

func (m *PingRequest) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type PingReply struct {
	Message string `protobuf:"bytes,1,opt,name=message" json:"message,omitempty"`
}

func (m *PingReply) Reset()                    { *m = PingReply{} }
func (m *PingReply) String() string            { return proto.CompactTextString(m) }
func (*PingReply) ProtoMessage()               {}
func (*PingReply) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{1} }

func (m *PingReply) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*PingRequest)(nil), "protos.PingRequest")
	proto.RegisterType((*PingReply)(nil), "protos.PingReply")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for StatusController service

type StatusControllerClient interface {
	Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingReply, error)
}

type statusControllerClient struct {
	cc *grpc.ClientConn
}

func NewStatusControllerClient(cc *grpc.ClientConn) StatusControllerClient {
	return &statusControllerClient{cc}
}

func (c *statusControllerClient) Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingReply, error) {
	out := new(PingReply)
	err := grpc.Invoke(ctx, "/protos.StatusController/Ping", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for StatusController service

type StatusControllerServer interface {
	Ping(context.Context, *PingRequest) (*PingReply, error)
}

func RegisterStatusControllerServer(s *grpc.Server, srv StatusControllerServer) {
	s.RegisterService(&_StatusController_serviceDesc, srv)
}

func _StatusController_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StatusControllerServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.StatusController/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StatusControllerServer).Ping(ctx, req.(*PingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _StatusController_serviceDesc = grpc.ServiceDesc{
	ServiceName: "protos.StatusController",
	HandlerType: (*StatusControllerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _StatusController_Ping_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "state.proto",
}

func init() { proto.RegisterFile("state.proto", fileDescriptor3) }

var fileDescriptor3 = []byte{
	// 135 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2e, 0x2e, 0x49, 0x2c,
	0x49, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x03, 0x53, 0xc5, 0x4a, 0xea, 0x5c, 0xdc,
	0x01, 0x99, 0x79, 0xe9, 0x41, 0xa9, 0x85, 0xa5, 0xa9, 0xc5, 0x25, 0x42, 0x12, 0x5c, 0xec, 0xb9,
	0xa9, 0xc5, 0xc5, 0x89, 0xe9, 0xa9, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0x30, 0xae, 0x92,
	0x2a, 0x17, 0x27, 0x44, 0x61, 0x41, 0x4e, 0x25, 0x6e, 0x65, 0x46, 0x2e, 0x5c, 0x02, 0xc1, 0x25,
	0x89, 0x25, 0xa5, 0xc5, 0xce, 0xf9, 0x79, 0x25, 0x45, 0xf9, 0x39, 0x39, 0xa9, 0x45, 0x42, 0x06,
	0x5c, 0x2c, 0x20, 0xad, 0x42, 0xc2, 0x10, 0xbb, 0x8b, 0xf5, 0x90, 0x6c, 0x94, 0x12, 0x44, 0x15,
	0x2c, 0xc8, 0xa9, 0x54, 0x62, 0x48, 0x82, 0xb8, 0xce, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0xd9,
	0x5c, 0xf0, 0xef, 0xb3, 0x00, 0x00, 0x00,
}
