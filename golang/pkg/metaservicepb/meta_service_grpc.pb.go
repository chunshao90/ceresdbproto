// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.20.1
// source: meta_service.proto

package metaservicepb

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

// HoraeMetaRpcServiceClient is the client API for HoraeMetaRpcService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type HoraeMetaRpcServiceClient interface {
	AllocSchemaID(ctx context.Context, in *AllocSchemaIdRequest, opts ...grpc.CallOption) (*AllocSchemaIdResponse, error)
	GetTablesOfShards(ctx context.Context, in *GetTablesOfShardsRequest, opts ...grpc.CallOption) (*GetTablesOfShardsResponse, error)
	CreateTable(ctx context.Context, in *CreateTableRequest, opts ...grpc.CallOption) (*CreateTableResponse, error)
	DropTable(ctx context.Context, in *DropTableRequest, opts ...grpc.CallOption) (*DropTableResponse, error)
	RouteTables(ctx context.Context, in *RouteTablesRequest, opts ...grpc.CallOption) (*RouteTablesResponse, error)
	GetNodes(ctx context.Context, in *GetNodesRequest, opts ...grpc.CallOption) (*GetNodesResponse, error)
	NodeHeartbeat(ctx context.Context, in *NodeHeartbeatRequest, opts ...grpc.CallOption) (*NodeHeartbeatResponse, error)
}

type horaeMetaRpcServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewHoraeMetaRpcServiceClient(cc grpc.ClientConnInterface) HoraeMetaRpcServiceClient {
	return &horaeMetaRpcServiceClient{cc}
}

func (c *horaeMetaRpcServiceClient) AllocSchemaID(ctx context.Context, in *AllocSchemaIdRequest, opts ...grpc.CallOption) (*AllocSchemaIdResponse, error) {
	out := new(AllocSchemaIdResponse)
	err := c.cc.Invoke(ctx, "/meta_service.HoraeMetaRpcService/AllocSchemaID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *horaeMetaRpcServiceClient) GetTablesOfShards(ctx context.Context, in *GetTablesOfShardsRequest, opts ...grpc.CallOption) (*GetTablesOfShardsResponse, error) {
	out := new(GetTablesOfShardsResponse)
	err := c.cc.Invoke(ctx, "/meta_service.HoraeMetaRpcService/GetTablesOfShards", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *horaeMetaRpcServiceClient) CreateTable(ctx context.Context, in *CreateTableRequest, opts ...grpc.CallOption) (*CreateTableResponse, error) {
	out := new(CreateTableResponse)
	err := c.cc.Invoke(ctx, "/meta_service.HoraeMetaRpcService/CreateTable", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *horaeMetaRpcServiceClient) DropTable(ctx context.Context, in *DropTableRequest, opts ...grpc.CallOption) (*DropTableResponse, error) {
	out := new(DropTableResponse)
	err := c.cc.Invoke(ctx, "/meta_service.HoraeMetaRpcService/DropTable", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *horaeMetaRpcServiceClient) RouteTables(ctx context.Context, in *RouteTablesRequest, opts ...grpc.CallOption) (*RouteTablesResponse, error) {
	out := new(RouteTablesResponse)
	err := c.cc.Invoke(ctx, "/meta_service.HoraeMetaRpcService/RouteTables", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *horaeMetaRpcServiceClient) GetNodes(ctx context.Context, in *GetNodesRequest, opts ...grpc.CallOption) (*GetNodesResponse, error) {
	out := new(GetNodesResponse)
	err := c.cc.Invoke(ctx, "/meta_service.HoraeMetaRpcService/GetNodes", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *horaeMetaRpcServiceClient) NodeHeartbeat(ctx context.Context, in *NodeHeartbeatRequest, opts ...grpc.CallOption) (*NodeHeartbeatResponse, error) {
	out := new(NodeHeartbeatResponse)
	err := c.cc.Invoke(ctx, "/meta_service.HoraeMetaRpcService/NodeHeartbeat", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HoraeMetaRpcServiceServer is the server API for HoraeMetaRpcService service.
// All implementations must embed UnimplementedHoraeMetaRpcServiceServer
// for forward compatibility
type HoraeMetaRpcServiceServer interface {
	AllocSchemaID(context.Context, *AllocSchemaIdRequest) (*AllocSchemaIdResponse, error)
	GetTablesOfShards(context.Context, *GetTablesOfShardsRequest) (*GetTablesOfShardsResponse, error)
	CreateTable(context.Context, *CreateTableRequest) (*CreateTableResponse, error)
	DropTable(context.Context, *DropTableRequest) (*DropTableResponse, error)
	RouteTables(context.Context, *RouteTablesRequest) (*RouteTablesResponse, error)
	GetNodes(context.Context, *GetNodesRequest) (*GetNodesResponse, error)
	NodeHeartbeat(context.Context, *NodeHeartbeatRequest) (*NodeHeartbeatResponse, error)
	mustEmbedUnimplementedHoraeMetaRpcServiceServer()
}

// UnimplementedHoraeMetaRpcServiceServer must be embedded to have forward compatible implementations.
type UnimplementedHoraeMetaRpcServiceServer struct {
}

func (UnimplementedHoraeMetaRpcServiceServer) AllocSchemaID(context.Context, *AllocSchemaIdRequest) (*AllocSchemaIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AllocSchemaID not implemented")
}
func (UnimplementedHoraeMetaRpcServiceServer) GetTablesOfShards(context.Context, *GetTablesOfShardsRequest) (*GetTablesOfShardsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTablesOfShards not implemented")
}
func (UnimplementedHoraeMetaRpcServiceServer) CreateTable(context.Context, *CreateTableRequest) (*CreateTableResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTable not implemented")
}
func (UnimplementedHoraeMetaRpcServiceServer) DropTable(context.Context, *DropTableRequest) (*DropTableResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DropTable not implemented")
}
func (UnimplementedHoraeMetaRpcServiceServer) RouteTables(context.Context, *RouteTablesRequest) (*RouteTablesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RouteTables not implemented")
}
func (UnimplementedHoraeMetaRpcServiceServer) GetNodes(context.Context, *GetNodesRequest) (*GetNodesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNodes not implemented")
}
func (UnimplementedHoraeMetaRpcServiceServer) NodeHeartbeat(context.Context, *NodeHeartbeatRequest) (*NodeHeartbeatResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NodeHeartbeat not implemented")
}
func (UnimplementedHoraeMetaRpcServiceServer) mustEmbedUnimplementedHoraeMetaRpcServiceServer() {}

// UnsafeHoraeMetaRpcServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to HoraeMetaRpcServiceServer will
// result in compilation errors.
type UnsafeHoraeMetaRpcServiceServer interface {
	mustEmbedUnimplementedHoraeMetaRpcServiceServer()
}

func RegisterHoraeMetaRpcServiceServer(s grpc.ServiceRegistrar, srv HoraeMetaRpcServiceServer) {
	s.RegisterService(&HoraeMetaRpcService_ServiceDesc, srv)
}

func _HoraeMetaRpcService_AllocSchemaID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AllocSchemaIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HoraeMetaRpcServiceServer).AllocSchemaID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/meta_service.HoraeMetaRpcService/AllocSchemaID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HoraeMetaRpcServiceServer).AllocSchemaID(ctx, req.(*AllocSchemaIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HoraeMetaRpcService_GetTablesOfShards_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTablesOfShardsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HoraeMetaRpcServiceServer).GetTablesOfShards(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/meta_service.HoraeMetaRpcService/GetTablesOfShards",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HoraeMetaRpcServiceServer).GetTablesOfShards(ctx, req.(*GetTablesOfShardsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HoraeMetaRpcService_CreateTable_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTableRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HoraeMetaRpcServiceServer).CreateTable(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/meta_service.HoraeMetaRpcService/CreateTable",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HoraeMetaRpcServiceServer).CreateTable(ctx, req.(*CreateTableRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HoraeMetaRpcService_DropTable_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DropTableRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HoraeMetaRpcServiceServer).DropTable(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/meta_service.HoraeMetaRpcService/DropTable",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HoraeMetaRpcServiceServer).DropTable(ctx, req.(*DropTableRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HoraeMetaRpcService_RouteTables_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RouteTablesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HoraeMetaRpcServiceServer).RouteTables(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/meta_service.HoraeMetaRpcService/RouteTables",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HoraeMetaRpcServiceServer).RouteTables(ctx, req.(*RouteTablesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HoraeMetaRpcService_GetNodes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetNodesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HoraeMetaRpcServiceServer).GetNodes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/meta_service.HoraeMetaRpcService/GetNodes",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HoraeMetaRpcServiceServer).GetNodes(ctx, req.(*GetNodesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HoraeMetaRpcService_NodeHeartbeat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NodeHeartbeatRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HoraeMetaRpcServiceServer).NodeHeartbeat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/meta_service.HoraeMetaRpcService/NodeHeartbeat",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HoraeMetaRpcServiceServer).NodeHeartbeat(ctx, req.(*NodeHeartbeatRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// HoraeMetaRpcService_ServiceDesc is the grpc.ServiceDesc for HoraeMetaRpcService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var HoraeMetaRpcService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "meta_service.HoraeMetaRpcService",
	HandlerType: (*HoraeMetaRpcServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AllocSchemaID",
			Handler:    _HoraeMetaRpcService_AllocSchemaID_Handler,
		},
		{
			MethodName: "GetTablesOfShards",
			Handler:    _HoraeMetaRpcService_GetTablesOfShards_Handler,
		},
		{
			MethodName: "CreateTable",
			Handler:    _HoraeMetaRpcService_CreateTable_Handler,
		},
		{
			MethodName: "DropTable",
			Handler:    _HoraeMetaRpcService_DropTable_Handler,
		},
		{
			MethodName: "RouteTables",
			Handler:    _HoraeMetaRpcService_RouteTables_Handler,
		},
		{
			MethodName: "GetNodes",
			Handler:    _HoraeMetaRpcService_GetNodes_Handler,
		},
		{
			MethodName: "NodeHeartbeat",
			Handler:    _HoraeMetaRpcService_NodeHeartbeat_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "meta_service.proto",
}
