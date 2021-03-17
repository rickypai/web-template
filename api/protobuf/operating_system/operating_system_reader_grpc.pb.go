// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package operating_system

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

// OperatingSystemReaderClient is the client API for OperatingSystemReader service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OperatingSystemReaderClient interface {
	GetOneByID(ctx context.Context, in *GetOneByIDRequest, opts ...grpc.CallOption) (*GetOneByIDResponse, error)
	GetManyByIDs(ctx context.Context, in *GetManyByIDsRequest, opts ...grpc.CallOption) (*GetManyByIDsResponse, error)
	ListByCursor(ctx context.Context, in *ListByCursorRequest, opts ...grpc.CallOption) (*ListByCursorResponse, error)
	ListByPage(ctx context.Context, in *ListByPageRequest, opts ...grpc.CallOption) (*ListByPageResponse, error)
	ListByPrefix(ctx context.Context, in *ListByPrefixRequest, opts ...grpc.CallOption) (*ListByPrefixResponse, error)
}

type operatingSystemReaderClient struct {
	cc grpc.ClientConnInterface
}

func NewOperatingSystemReaderClient(cc grpc.ClientConnInterface) OperatingSystemReaderClient {
	return &operatingSystemReaderClient{cc}
}

func (c *operatingSystemReaderClient) GetOneByID(ctx context.Context, in *GetOneByIDRequest, opts ...grpc.CallOption) (*GetOneByIDResponse, error) {
	out := new(GetOneByIDResponse)
	err := c.cc.Invoke(ctx, "/operating_system.OperatingSystemReader/GetOneByID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *operatingSystemReaderClient) GetManyByIDs(ctx context.Context, in *GetManyByIDsRequest, opts ...grpc.CallOption) (*GetManyByIDsResponse, error) {
	out := new(GetManyByIDsResponse)
	err := c.cc.Invoke(ctx, "/operating_system.OperatingSystemReader/GetManyByIDs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *operatingSystemReaderClient) ListByCursor(ctx context.Context, in *ListByCursorRequest, opts ...grpc.CallOption) (*ListByCursorResponse, error) {
	out := new(ListByCursorResponse)
	err := c.cc.Invoke(ctx, "/operating_system.OperatingSystemReader/ListByCursor", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *operatingSystemReaderClient) ListByPage(ctx context.Context, in *ListByPageRequest, opts ...grpc.CallOption) (*ListByPageResponse, error) {
	out := new(ListByPageResponse)
	err := c.cc.Invoke(ctx, "/operating_system.OperatingSystemReader/ListByPage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *operatingSystemReaderClient) ListByPrefix(ctx context.Context, in *ListByPrefixRequest, opts ...grpc.CallOption) (*ListByPrefixResponse, error) {
	out := new(ListByPrefixResponse)
	err := c.cc.Invoke(ctx, "/operating_system.OperatingSystemReader/ListByPrefix", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OperatingSystemReaderServer is the server API for OperatingSystemReader service.
// All implementations must embed UnimplementedOperatingSystemReaderServer
// for forward compatibility
type OperatingSystemReaderServer interface {
	GetOneByID(context.Context, *GetOneByIDRequest) (*GetOneByIDResponse, error)
	GetManyByIDs(context.Context, *GetManyByIDsRequest) (*GetManyByIDsResponse, error)
	ListByCursor(context.Context, *ListByCursorRequest) (*ListByCursorResponse, error)
	ListByPage(context.Context, *ListByPageRequest) (*ListByPageResponse, error)
	ListByPrefix(context.Context, *ListByPrefixRequest) (*ListByPrefixResponse, error)
	mustEmbedUnimplementedOperatingSystemReaderServer()
}

// UnimplementedOperatingSystemReaderServer must be embedded to have forward compatible implementations.
type UnimplementedOperatingSystemReaderServer struct {
}

func (UnimplementedOperatingSystemReaderServer) GetOneByID(context.Context, *GetOneByIDRequest) (*GetOneByIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOneByID not implemented")
}
func (UnimplementedOperatingSystemReaderServer) GetManyByIDs(context.Context, *GetManyByIDsRequest) (*GetManyByIDsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetManyByIDs not implemented")
}
func (UnimplementedOperatingSystemReaderServer) ListByCursor(context.Context, *ListByCursorRequest) (*ListByCursorResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListByCursor not implemented")
}
func (UnimplementedOperatingSystemReaderServer) ListByPage(context.Context, *ListByPageRequest) (*ListByPageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListByPage not implemented")
}
func (UnimplementedOperatingSystemReaderServer) ListByPrefix(context.Context, *ListByPrefixRequest) (*ListByPrefixResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListByPrefix not implemented")
}
func (UnimplementedOperatingSystemReaderServer) mustEmbedUnimplementedOperatingSystemReaderServer() {}

// UnsafeOperatingSystemReaderServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OperatingSystemReaderServer will
// result in compilation errors.
type UnsafeOperatingSystemReaderServer interface {
	mustEmbedUnimplementedOperatingSystemReaderServer()
}

func RegisterOperatingSystemReaderServer(s grpc.ServiceRegistrar, srv OperatingSystemReaderServer) {
	s.RegisterService(&OperatingSystemReader_ServiceDesc, srv)
}

func _OperatingSystemReader_GetOneByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOneByIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OperatingSystemReaderServer).GetOneByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/operating_system.OperatingSystemReader/GetOneByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OperatingSystemReaderServer).GetOneByID(ctx, req.(*GetOneByIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OperatingSystemReader_GetManyByIDs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetManyByIDsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OperatingSystemReaderServer).GetManyByIDs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/operating_system.OperatingSystemReader/GetManyByIDs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OperatingSystemReaderServer).GetManyByIDs(ctx, req.(*GetManyByIDsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OperatingSystemReader_ListByCursor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListByCursorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OperatingSystemReaderServer).ListByCursor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/operating_system.OperatingSystemReader/ListByCursor",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OperatingSystemReaderServer).ListByCursor(ctx, req.(*ListByCursorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OperatingSystemReader_ListByPage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListByPageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OperatingSystemReaderServer).ListByPage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/operating_system.OperatingSystemReader/ListByPage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OperatingSystemReaderServer).ListByPage(ctx, req.(*ListByPageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OperatingSystemReader_ListByPrefix_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListByPrefixRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OperatingSystemReaderServer).ListByPrefix(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/operating_system.OperatingSystemReader/ListByPrefix",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OperatingSystemReaderServer).ListByPrefix(ctx, req.(*ListByPrefixRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// OperatingSystemReader_ServiceDesc is the grpc.ServiceDesc for OperatingSystemReader service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var OperatingSystemReader_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "operating_system.OperatingSystemReader",
	HandlerType: (*OperatingSystemReaderServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetOneByID",
			Handler:    _OperatingSystemReader_GetOneByID_Handler,
		},
		{
			MethodName: "GetManyByIDs",
			Handler:    _OperatingSystemReader_GetManyByIDs_Handler,
		},
		{
			MethodName: "ListByCursor",
			Handler:    _OperatingSystemReader_ListByCursor_Handler,
		},
		{
			MethodName: "ListByPage",
			Handler:    _OperatingSystemReader_ListByPage_Handler,
		},
		{
			MethodName: "ListByPrefix",
			Handler:    _OperatingSystemReader_ListByPrefix_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protobuf/operating_system/operating_system_reader.proto",
}