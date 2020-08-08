// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package serviceav1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// EchoAClient is the client API for EchoA service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type EchoAClient interface {
	Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingResponse, error)
}

type echoAClient struct {
	cc grpc.ClientConnInterface
}

func NewEchoAClient(cc grpc.ClientConnInterface) EchoAClient {
	return &echoAClient{cc}
}

func (c *echoAClient) Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingResponse, error) {
	out := new(PingResponse)
	err := c.cc.Invoke(ctx, "/parthpower.servicea.v1.EchoA/Ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EchoAServer is the server API for EchoA service.
// All implementations must embed UnimplementedEchoAServer
// for forward compatibility
type EchoAServer interface {
	Ping(context.Context, *PingRequest) (*PingResponse, error)
	mustEmbedUnimplementedEchoAServer()
}

// UnimplementedEchoAServer must be embedded to have forward compatible implementations.
type UnimplementedEchoAServer struct {
}

func (*UnimplementedEchoAServer) Ping(context.Context, *PingRequest) (*PingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (*UnimplementedEchoAServer) mustEmbedUnimplementedEchoAServer() {}

func RegisterEchoAServer(s *grpc.Server, srv EchoAServer) {
	s.RegisterService(&_EchoA_serviceDesc, srv)
}

func _EchoA_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EchoAServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/parthpower.servicea.v1.EchoA/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EchoAServer).Ping(ctx, req.(*PingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _EchoA_serviceDesc = grpc.ServiceDesc{
	ServiceName: "parthpower.servicea.v1.EchoA",
	HandlerType: (*EchoAServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _EchoA_Ping_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "servicea.proto",
}
