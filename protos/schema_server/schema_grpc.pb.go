// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.6.1
// source: schema.proto

package schema_server

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

// SchemaServerClient is the client API for SchemaServer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SchemaServerClient interface {
	// returns schema name, vendor, version, and files path(s)
	GetSchemaDetails(ctx context.Context, in *GetSchemaDetailsRequest, opts ...grpc.CallOption) (*GetSchemaDetailsResponse, error)
	// lists known schemas with name, vendor, version and status
	ListSchema(ctx context.Context, in *ListSchemaRequest, opts ...grpc.CallOption) (*ListSchemaResponse, error)
	// returns the schema of an item identified by a gNMI-like path
	GetSchema(ctx context.Context, in *GetSchemaRequest, opts ...grpc.CallOption) (*GetSchemaResponse, error)
	// creates a schema
	CreateSchema(ctx context.Context, in *CreateSchemaRequest, opts ...grpc.CallOption) (*CreateSchemaResponse, error)
	// trigger schema reload
	ReloadSchema(ctx context.Context, in *ReloadSchemaRequest, opts ...grpc.CallOption) (*ReloadSchemaResponse, error)
	// delete a schema
	DeleteSchema(ctx context.Context, in *DeleteSchemaRequest, opts ...grpc.CallOption) (*DeleteSchemaResponse, error)
	// client stream RPC to upload yang files to the server:
	// - uses CreateSchema as a first message
	// - then N intermediate UploadSchemaFile, initial, bytes, hash for each file
	// - and ends with an UploadSchemaFinalize{}
	UploadSchema(ctx context.Context, opts ...grpc.CallOption) (SchemaServer_UploadSchemaClient, error)
	// ToPath converts a list of items into a schema.proto.Path
	ToPath(ctx context.Context, in *ToPathRequest, opts ...grpc.CallOption) (*ToPathResponse, error)
	// ExpandPath returns a list of sub paths given a single path
	ExpandPath(ctx context.Context, in *ExpandPathRequest, opts ...grpc.CallOption) (*ExpandPathResponse, error)
}

type schemaServerClient struct {
	cc grpc.ClientConnInterface
}

func NewSchemaServerClient(cc grpc.ClientConnInterface) SchemaServerClient {
	return &schemaServerClient{cc}
}

func (c *schemaServerClient) GetSchemaDetails(ctx context.Context, in *GetSchemaDetailsRequest, opts ...grpc.CallOption) (*GetSchemaDetailsResponse, error) {
	out := new(GetSchemaDetailsResponse)
	err := c.cc.Invoke(ctx, "/schema.proto.SchemaServer/GetSchemaDetails", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *schemaServerClient) ListSchema(ctx context.Context, in *ListSchemaRequest, opts ...grpc.CallOption) (*ListSchemaResponse, error) {
	out := new(ListSchemaResponse)
	err := c.cc.Invoke(ctx, "/schema.proto.SchemaServer/ListSchema", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *schemaServerClient) GetSchema(ctx context.Context, in *GetSchemaRequest, opts ...grpc.CallOption) (*GetSchemaResponse, error) {
	out := new(GetSchemaResponse)
	err := c.cc.Invoke(ctx, "/schema.proto.SchemaServer/GetSchema", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *schemaServerClient) CreateSchema(ctx context.Context, in *CreateSchemaRequest, opts ...grpc.CallOption) (*CreateSchemaResponse, error) {
	out := new(CreateSchemaResponse)
	err := c.cc.Invoke(ctx, "/schema.proto.SchemaServer/CreateSchema", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *schemaServerClient) ReloadSchema(ctx context.Context, in *ReloadSchemaRequest, opts ...grpc.CallOption) (*ReloadSchemaResponse, error) {
	out := new(ReloadSchemaResponse)
	err := c.cc.Invoke(ctx, "/schema.proto.SchemaServer/ReloadSchema", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *schemaServerClient) DeleteSchema(ctx context.Context, in *DeleteSchemaRequest, opts ...grpc.CallOption) (*DeleteSchemaResponse, error) {
	out := new(DeleteSchemaResponse)
	err := c.cc.Invoke(ctx, "/schema.proto.SchemaServer/DeleteSchema", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *schemaServerClient) UploadSchema(ctx context.Context, opts ...grpc.CallOption) (SchemaServer_UploadSchemaClient, error) {
	stream, err := c.cc.NewStream(ctx, &SchemaServer_ServiceDesc.Streams[0], "/schema.proto.SchemaServer/UploadSchema", opts...)
	if err != nil {
		return nil, err
	}
	x := &schemaServerUploadSchemaClient{stream}
	return x, nil
}

type SchemaServer_UploadSchemaClient interface {
	Send(*UploadSchemaRequest) error
	CloseAndRecv() (*UploadSchemaResponse, error)
	grpc.ClientStream
}

type schemaServerUploadSchemaClient struct {
	grpc.ClientStream
}

func (x *schemaServerUploadSchemaClient) Send(m *UploadSchemaRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *schemaServerUploadSchemaClient) CloseAndRecv() (*UploadSchemaResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(UploadSchemaResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *schemaServerClient) ToPath(ctx context.Context, in *ToPathRequest, opts ...grpc.CallOption) (*ToPathResponse, error) {
	out := new(ToPathResponse)
	err := c.cc.Invoke(ctx, "/schema.proto.SchemaServer/ToPath", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *schemaServerClient) ExpandPath(ctx context.Context, in *ExpandPathRequest, opts ...grpc.CallOption) (*ExpandPathResponse, error) {
	out := new(ExpandPathResponse)
	err := c.cc.Invoke(ctx, "/schema.proto.SchemaServer/ExpandPath", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SchemaServerServer is the server API for SchemaServer service.
// All implementations must embed UnimplementedSchemaServerServer
// for forward compatibility
type SchemaServerServer interface {
	// returns schema name, vendor, version, and files path(s)
	GetSchemaDetails(context.Context, *GetSchemaDetailsRequest) (*GetSchemaDetailsResponse, error)
	// lists known schemas with name, vendor, version and status
	ListSchema(context.Context, *ListSchemaRequest) (*ListSchemaResponse, error)
	// returns the schema of an item identified by a gNMI-like path
	GetSchema(context.Context, *GetSchemaRequest) (*GetSchemaResponse, error)
	// creates a schema
	CreateSchema(context.Context, *CreateSchemaRequest) (*CreateSchemaResponse, error)
	// trigger schema reload
	ReloadSchema(context.Context, *ReloadSchemaRequest) (*ReloadSchemaResponse, error)
	// delete a schema
	DeleteSchema(context.Context, *DeleteSchemaRequest) (*DeleteSchemaResponse, error)
	// client stream RPC to upload yang files to the server:
	// - uses CreateSchema as a first message
	// - then N intermediate UploadSchemaFile, initial, bytes, hash for each file
	// - and ends with an UploadSchemaFinalize{}
	UploadSchema(SchemaServer_UploadSchemaServer) error
	// ToPath converts a list of items into a schema.proto.Path
	ToPath(context.Context, *ToPathRequest) (*ToPathResponse, error)
	// ExpandPath returns a list of sub paths given a single path
	ExpandPath(context.Context, *ExpandPathRequest) (*ExpandPathResponse, error)
	mustEmbedUnimplementedSchemaServerServer()
}

// UnimplementedSchemaServerServer must be embedded to have forward compatible implementations.
type UnimplementedSchemaServerServer struct {
}

func (UnimplementedSchemaServerServer) GetSchemaDetails(context.Context, *GetSchemaDetailsRequest) (*GetSchemaDetailsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSchemaDetails not implemented")
}
func (UnimplementedSchemaServerServer) ListSchema(context.Context, *ListSchemaRequest) (*ListSchemaResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListSchema not implemented")
}
func (UnimplementedSchemaServerServer) GetSchema(context.Context, *GetSchemaRequest) (*GetSchemaResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSchema not implemented")
}
func (UnimplementedSchemaServerServer) CreateSchema(context.Context, *CreateSchemaRequest) (*CreateSchemaResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateSchema not implemented")
}
func (UnimplementedSchemaServerServer) ReloadSchema(context.Context, *ReloadSchemaRequest) (*ReloadSchemaResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReloadSchema not implemented")
}
func (UnimplementedSchemaServerServer) DeleteSchema(context.Context, *DeleteSchemaRequest) (*DeleteSchemaResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteSchema not implemented")
}
func (UnimplementedSchemaServerServer) UploadSchema(SchemaServer_UploadSchemaServer) error {
	return status.Errorf(codes.Unimplemented, "method UploadSchema not implemented")
}
func (UnimplementedSchemaServerServer) ToPath(context.Context, *ToPathRequest) (*ToPathResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ToPath not implemented")
}
func (UnimplementedSchemaServerServer) ExpandPath(context.Context, *ExpandPathRequest) (*ExpandPathResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExpandPath not implemented")
}
func (UnimplementedSchemaServerServer) mustEmbedUnimplementedSchemaServerServer() {}

// UnsafeSchemaServerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SchemaServerServer will
// result in compilation errors.
type UnsafeSchemaServerServer interface {
	mustEmbedUnimplementedSchemaServerServer()
}

func RegisterSchemaServerServer(s grpc.ServiceRegistrar, srv SchemaServerServer) {
	s.RegisterService(&SchemaServer_ServiceDesc, srv)
}

func _SchemaServer_GetSchemaDetails_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSchemaDetailsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SchemaServerServer).GetSchemaDetails(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/schema.proto.SchemaServer/GetSchemaDetails",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SchemaServerServer).GetSchemaDetails(ctx, req.(*GetSchemaDetailsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SchemaServer_ListSchema_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListSchemaRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SchemaServerServer).ListSchema(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/schema.proto.SchemaServer/ListSchema",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SchemaServerServer).ListSchema(ctx, req.(*ListSchemaRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SchemaServer_GetSchema_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSchemaRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SchemaServerServer).GetSchema(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/schema.proto.SchemaServer/GetSchema",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SchemaServerServer).GetSchema(ctx, req.(*GetSchemaRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SchemaServer_CreateSchema_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateSchemaRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SchemaServerServer).CreateSchema(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/schema.proto.SchemaServer/CreateSchema",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SchemaServerServer).CreateSchema(ctx, req.(*CreateSchemaRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SchemaServer_ReloadSchema_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReloadSchemaRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SchemaServerServer).ReloadSchema(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/schema.proto.SchemaServer/ReloadSchema",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SchemaServerServer).ReloadSchema(ctx, req.(*ReloadSchemaRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SchemaServer_DeleteSchema_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteSchemaRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SchemaServerServer).DeleteSchema(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/schema.proto.SchemaServer/DeleteSchema",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SchemaServerServer).DeleteSchema(ctx, req.(*DeleteSchemaRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SchemaServer_UploadSchema_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(SchemaServerServer).UploadSchema(&schemaServerUploadSchemaServer{stream})
}

type SchemaServer_UploadSchemaServer interface {
	SendAndClose(*UploadSchemaResponse) error
	Recv() (*UploadSchemaRequest, error)
	grpc.ServerStream
}

type schemaServerUploadSchemaServer struct {
	grpc.ServerStream
}

func (x *schemaServerUploadSchemaServer) SendAndClose(m *UploadSchemaResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *schemaServerUploadSchemaServer) Recv() (*UploadSchemaRequest, error) {
	m := new(UploadSchemaRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _SchemaServer_ToPath_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ToPathRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SchemaServerServer).ToPath(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/schema.proto.SchemaServer/ToPath",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SchemaServerServer).ToPath(ctx, req.(*ToPathRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SchemaServer_ExpandPath_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExpandPathRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SchemaServerServer).ExpandPath(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/schema.proto.SchemaServer/ExpandPath",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SchemaServerServer).ExpandPath(ctx, req.(*ExpandPathRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SchemaServer_ServiceDesc is the grpc.ServiceDesc for SchemaServer service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SchemaServer_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "schema.proto.SchemaServer",
	HandlerType: (*SchemaServerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetSchemaDetails",
			Handler:    _SchemaServer_GetSchemaDetails_Handler,
		},
		{
			MethodName: "ListSchema",
			Handler:    _SchemaServer_ListSchema_Handler,
		},
		{
			MethodName: "GetSchema",
			Handler:    _SchemaServer_GetSchema_Handler,
		},
		{
			MethodName: "CreateSchema",
			Handler:    _SchemaServer_CreateSchema_Handler,
		},
		{
			MethodName: "ReloadSchema",
			Handler:    _SchemaServer_ReloadSchema_Handler,
		},
		{
			MethodName: "DeleteSchema",
			Handler:    _SchemaServer_DeleteSchema_Handler,
		},
		{
			MethodName: "ToPath",
			Handler:    _SchemaServer_ToPath_Handler,
		},
		{
			MethodName: "ExpandPath",
			Handler:    _SchemaServer_ExpandPath_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "UploadSchema",
			Handler:       _SchemaServer_UploadSchema_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "schema.proto",
}
