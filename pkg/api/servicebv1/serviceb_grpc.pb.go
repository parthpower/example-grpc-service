// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package servicebv1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// EchoBClient is the client API for EchoB service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type EchoBClient interface {
	Flip(ctx context.Context, in *FlipRequest, opts ...grpc.CallOption) (*FlipResponse, error)
}

type echoBClient struct {
	cc grpc.ClientConnInterface
}

func NewEchoBClient(cc grpc.ClientConnInterface) EchoBClient {
	return &echoBClient{cc}
}

func (c *echoBClient) Flip(ctx context.Context, in *FlipRequest, opts ...grpc.CallOption) (*FlipResponse, error) {
	out := new(FlipResponse)
	err := c.cc.Invoke(ctx, "/parthpower.serviceb.v1.EchoB/Flip", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EchoBServer is the server API for EchoB service.
// All implementations must embed UnimplementedEchoBServer
// for forward compatibility
type EchoBServer interface {
	Flip(context.Context, *FlipRequest) (*FlipResponse, error)
	mustEmbedUnimplementedEchoBServer()
}

// UnimplementedEchoBServer must be embedded to have forward compatible implementations.
type UnimplementedEchoBServer struct {
}

func (*UnimplementedEchoBServer) Flip(context.Context, *FlipRequest) (*FlipResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Flip not implemented")
}
func (*UnimplementedEchoBServer) mustEmbedUnimplementedEchoBServer() {}

func RegisterEchoBServer(s *grpc.Server, srv EchoBServer) {
	s.RegisterService(&_EchoB_serviceDesc, srv)
}

func _EchoB_Flip_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FlipRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EchoBServer).Flip(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/parthpower.serviceb.v1.EchoB/Flip",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EchoBServer).Flip(ctx, req.(*FlipRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _EchoB_serviceDesc = grpc.ServiceDesc{
	ServiceName: "parthpower.serviceb.v1.EchoB",
	HandlerType: (*EchoBServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Flip",
			Handler:    _EchoB_Flip_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "serviceb.proto",
}
