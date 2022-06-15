// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.6.1
// source: teller.proto

package tellerpb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// TellerServiceClient is the client API for TellerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TellerServiceClient interface {
	GetTeller(ctx context.Context, in *GetTellerRequest, opts ...grpc.CallOption) (*GetTellerResponse, error)
	GetTellers(ctx context.Context, in *GetTellersRequest, opts ...grpc.CallOption) (*GetTellersResponse, error)
}

type tellerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTellerServiceClient(cc grpc.ClientConnInterface) TellerServiceClient {
	return &tellerServiceClient{cc}
}

func (c *tellerServiceClient) GetTeller(ctx context.Context, in *GetTellerRequest, opts ...grpc.CallOption) (*GetTellerResponse, error) {
	out := new(GetTellerResponse)
	err := c.cc.Invoke(ctx, "/main.TellerService/GetTeller", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tellerServiceClient) GetTellers(ctx context.Context, in *GetTellersRequest, opts ...grpc.CallOption) (*GetTellersResponse, error) {
	out := new(GetTellersResponse)
	err := c.cc.Invoke(ctx, "/main.TellerService/GetTellers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TellerServiceServer is the server API for TellerService service.
// All implementations must embed UnimplementedTellerServiceServer
// for forward compatibility
type TellerServiceServer interface {
	GetTeller(context.Context, *GetTellerRequest) (*GetTellerResponse, error)
	GetTellers(context.Context, *GetTellersRequest) (*GetTellersResponse, error)
	mustEmbedUnimplementedTellerServiceServer()
}

// UnimplementedTellerServiceServer must be embedded to have forward compatible implementations.
type UnimplementedTellerServiceServer struct {
}

func (UnimplementedTellerServiceServer) GetTeller(context.Context, *GetTellerRequest) (*GetTellerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTeller not implemented")
}
func (UnimplementedTellerServiceServer) GetTellers(context.Context, *GetTellersRequest) (*GetTellersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTellers not implemented")
}
func (UnimplementedTellerServiceServer) mustEmbedUnimplementedTellerServiceServer() {}

// UnsafeTellerServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TellerServiceServer will
// result in compilation errors.
type UnsafeTellerServiceServer interface {
	mustEmbedUnimplementedTellerServiceServer()
}

func RegisterTellerServiceServer(s grpc.ServiceRegistrar, srv TellerServiceServer) {
	s.RegisterService(&TellerService_ServiceDesc, srv)
}

func _TellerService_GetTeller_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTellerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TellerServiceServer).GetTeller(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/main.TellerService/GetTeller",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TellerServiceServer).GetTeller(ctx, req.(*GetTellerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TellerService_GetTellers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTellersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TellerServiceServer).GetTellers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/main.TellerService/GetTellers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TellerServiceServer).GetTellers(ctx, req.(*GetTellersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TellerService_ServiceDesc is the grpc.ServiceDesc for TellerService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TellerService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "main.TellerService",
	HandlerType: (*TellerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetTeller",
			Handler:    _TellerService_GetTeller_Handler,
		},
		{
			MethodName: "GetTellers",
			Handler:    _TellerService_GetTellers_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "teller.proto",
}