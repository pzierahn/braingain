// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v5.27.3
// source: document_service.proto

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	Document_List_FullMethodName   = "/chatbot.documents.v1.Document/List"
	Document_Rename_FullMethodName = "/chatbot.documents.v1.Document/Rename"
	Document_Delete_FullMethodName = "/chatbot.documents.v1.Document/Delete"
	Document_Index_FullMethodName  = "/chatbot.documents.v1.Document/Index"
	Document_Search_FullMethodName = "/chatbot.documents.v1.Document/Search"
)

// DocumentClient is the client API for Document service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DocumentClient interface {
	List(ctx context.Context, in *DocumentFilter, opts ...grpc.CallOption) (*DocumentList, error)
	Rename(ctx context.Context, in *RenameDocument, opts ...grpc.CallOption) (*emptypb.Empty, error)
	Delete(ctx context.Context, in *DocumentID, opts ...grpc.CallOption) (*emptypb.Empty, error)
	Index(ctx context.Context, in *IndexJob, opts ...grpc.CallOption) (Document_IndexClient, error)
	Search(ctx context.Context, in *SearchQuery, opts ...grpc.CallOption) (*SearchResults, error)
}

type documentClient struct {
	cc grpc.ClientConnInterface
}

func NewDocumentClient(cc grpc.ClientConnInterface) DocumentClient {
	return &documentClient{cc}
}

func (c *documentClient) List(ctx context.Context, in *DocumentFilter, opts ...grpc.CallOption) (*DocumentList, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DocumentList)
	err := c.cc.Invoke(ctx, Document_List_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *documentClient) Rename(ctx context.Context, in *RenameDocument, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, Document_Rename_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *documentClient) Delete(ctx context.Context, in *DocumentID, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, Document_Delete_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *documentClient) Index(ctx context.Context, in *IndexJob, opts ...grpc.CallOption) (Document_IndexClient, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &Document_ServiceDesc.Streams[0], Document_Index_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &documentIndexClient{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Document_IndexClient interface {
	Recv() (*IndexProgress, error)
	grpc.ClientStream
}

type documentIndexClient struct {
	grpc.ClientStream
}

func (x *documentIndexClient) Recv() (*IndexProgress, error) {
	m := new(IndexProgress)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *documentClient) Search(ctx context.Context, in *SearchQuery, opts ...grpc.CallOption) (*SearchResults, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SearchResults)
	err := c.cc.Invoke(ctx, Document_Search_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DocumentServer is the server API for Document service.
// All implementations must embed UnimplementedDocumentServer
// for forward compatibility
type DocumentServer interface {
	List(context.Context, *DocumentFilter) (*DocumentList, error)
	Rename(context.Context, *RenameDocument) (*emptypb.Empty, error)
	Delete(context.Context, *DocumentID) (*emptypb.Empty, error)
	Index(*IndexJob, Document_IndexServer) error
	Search(context.Context, *SearchQuery) (*SearchResults, error)
	mustEmbedUnimplementedDocumentServer()
}

// UnimplementedDocumentServer must be embedded to have forward compatible implementations.
type UnimplementedDocumentServer struct {
}

func (UnimplementedDocumentServer) List(context.Context, *DocumentFilter) (*DocumentList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedDocumentServer) Rename(context.Context, *RenameDocument) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Rename not implemented")
}
func (UnimplementedDocumentServer) Delete(context.Context, *DocumentID) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedDocumentServer) Index(*IndexJob, Document_IndexServer) error {
	return status.Errorf(codes.Unimplemented, "method Index not implemented")
}
func (UnimplementedDocumentServer) Search(context.Context, *SearchQuery) (*SearchResults, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Search not implemented")
}
func (UnimplementedDocumentServer) mustEmbedUnimplementedDocumentServer() {}

// UnsafeDocumentServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DocumentServer will
// result in compilation errors.
type UnsafeDocumentServer interface {
	mustEmbedUnimplementedDocumentServer()
}

func RegisterDocumentServer(s grpc.ServiceRegistrar, srv DocumentServer) {
	s.RegisterService(&Document_ServiceDesc, srv)
}

func _Document_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DocumentFilter)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DocumentServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Document_List_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DocumentServer).List(ctx, req.(*DocumentFilter))
	}
	return interceptor(ctx, in, info, handler)
}

func _Document_Rename_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RenameDocument)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DocumentServer).Rename(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Document_Rename_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DocumentServer).Rename(ctx, req.(*RenameDocument))
	}
	return interceptor(ctx, in, info, handler)
}

func _Document_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DocumentID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DocumentServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Document_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DocumentServer).Delete(ctx, req.(*DocumentID))
	}
	return interceptor(ctx, in, info, handler)
}

func _Document_Index_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(IndexJob)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(DocumentServer).Index(m, &documentIndexServer{ServerStream: stream})
}

type Document_IndexServer interface {
	Send(*IndexProgress) error
	grpc.ServerStream
}

type documentIndexServer struct {
	grpc.ServerStream
}

func (x *documentIndexServer) Send(m *IndexProgress) error {
	return x.ServerStream.SendMsg(m)
}

func _Document_Search_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DocumentServer).Search(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Document_Search_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DocumentServer).Search(ctx, req.(*SearchQuery))
	}
	return interceptor(ctx, in, info, handler)
}

// Document_ServiceDesc is the grpc.ServiceDesc for Document service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Document_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "chatbot.documents.v1.Document",
	HandlerType: (*DocumentServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "List",
			Handler:    _Document_List_Handler,
		},
		{
			MethodName: "Rename",
			Handler:    _Document_Rename_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _Document_Delete_Handler,
		},
		{
			MethodName: "Search",
			Handler:    _Document_Search_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Index",
			Handler:       _Document_Index_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "document_service.proto",
}
