// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.1.0
// - protoc             v3.15.5
// source: api/article.proto

package api

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

// ArticleServiceClient is the client API for ArticleService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ArticleServiceClient interface {
	Ping(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error)
	DelAuthorByKey(ctx context.Context, in *Author, opts ...grpc.CallOption) (*Empty, error)
	ListArticleMeta(ctx context.Context, in *ListArticleMetaReq, opts ...grpc.CallOption) (*ListArticleMetaRes, error)
	PutAuthor(ctx context.Context, in *Author, opts ...grpc.CallOption) (*AuthorID, error)
	DelAuthor(ctx context.Context, in *AuthorID, opts ...grpc.CallOption) (*Empty, error)
	PutArticle(ctx context.Context, in *Article, opts ...grpc.CallOption) (*ArticleID, error)
	GetArticle(ctx context.Context, in *ArticleID, opts ...grpc.CallOption) (*Article, error)
	DelArticle(ctx context.Context, in *ArticleID, opts ...grpc.CallOption) (*Empty, error)
	UpdateArticle(ctx context.Context, in *Article, opts ...grpc.CallOption) (*Empty, error)
	ListArticle(ctx context.Context, in *ListArticleReq, opts ...grpc.CallOption) (*ListArticleRes, error)
}

type articleServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewArticleServiceClient(cc grpc.ClientConnInterface) ArticleServiceClient {
	return &articleServiceClient{cc}
}

func (c *articleServiceClient) Ping(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/api.ArticleService/Ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *articleServiceClient) DelAuthorByKey(ctx context.Context, in *Author, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/api.ArticleService/DelAuthorByKey", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *articleServiceClient) ListArticleMeta(ctx context.Context, in *ListArticleMetaReq, opts ...grpc.CallOption) (*ListArticleMetaRes, error) {
	out := new(ListArticleMetaRes)
	err := c.cc.Invoke(ctx, "/api.ArticleService/ListArticleMeta", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *articleServiceClient) PutAuthor(ctx context.Context, in *Author, opts ...grpc.CallOption) (*AuthorID, error) {
	out := new(AuthorID)
	err := c.cc.Invoke(ctx, "/api.ArticleService/PutAuthor", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *articleServiceClient) DelAuthor(ctx context.Context, in *AuthorID, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/api.ArticleService/DelAuthor", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *articleServiceClient) PutArticle(ctx context.Context, in *Article, opts ...grpc.CallOption) (*ArticleID, error) {
	out := new(ArticleID)
	err := c.cc.Invoke(ctx, "/api.ArticleService/PutArticle", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *articleServiceClient) GetArticle(ctx context.Context, in *ArticleID, opts ...grpc.CallOption) (*Article, error) {
	out := new(Article)
	err := c.cc.Invoke(ctx, "/api.ArticleService/GetArticle", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *articleServiceClient) DelArticle(ctx context.Context, in *ArticleID, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/api.ArticleService/DelArticle", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *articleServiceClient) UpdateArticle(ctx context.Context, in *Article, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/api.ArticleService/UpdateArticle", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *articleServiceClient) ListArticle(ctx context.Context, in *ListArticleReq, opts ...grpc.CallOption) (*ListArticleRes, error) {
	out := new(ListArticleRes)
	err := c.cc.Invoke(ctx, "/api.ArticleService/ListArticle", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ArticleServiceServer is the server API for ArticleService service.
// All implementations must embed UnimplementedArticleServiceServer
// for forward compatibility
type ArticleServiceServer interface {
	Ping(context.Context, *Empty) (*Empty, error)
	DelAuthorByKey(context.Context, *Author) (*Empty, error)
	ListArticleMeta(context.Context, *ListArticleMetaReq) (*ListArticleMetaRes, error)
	PutAuthor(context.Context, *Author) (*AuthorID, error)
	DelAuthor(context.Context, *AuthorID) (*Empty, error)
	PutArticle(context.Context, *Article) (*ArticleID, error)
	GetArticle(context.Context, *ArticleID) (*Article, error)
	DelArticle(context.Context, *ArticleID) (*Empty, error)
	UpdateArticle(context.Context, *Article) (*Empty, error)
	ListArticle(context.Context, *ListArticleReq) (*ListArticleRes, error)
	mustEmbedUnimplementedArticleServiceServer()
}

// UnimplementedArticleServiceServer must be embedded to have forward compatible implementations.
type UnimplementedArticleServiceServer struct {
}

func (UnimplementedArticleServiceServer) Ping(context.Context, *Empty) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (UnimplementedArticleServiceServer) DelAuthorByKey(context.Context, *Author) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DelAuthorByKey not implemented")
}
func (UnimplementedArticleServiceServer) ListArticleMeta(context.Context, *ListArticleMetaReq) (*ListArticleMetaRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListArticleMeta not implemented")
}
func (UnimplementedArticleServiceServer) PutAuthor(context.Context, *Author) (*AuthorID, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PutAuthor not implemented")
}
func (UnimplementedArticleServiceServer) DelAuthor(context.Context, *AuthorID) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DelAuthor not implemented")
}
func (UnimplementedArticleServiceServer) PutArticle(context.Context, *Article) (*ArticleID, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PutArticle not implemented")
}
func (UnimplementedArticleServiceServer) GetArticle(context.Context, *ArticleID) (*Article, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetArticle not implemented")
}
func (UnimplementedArticleServiceServer) DelArticle(context.Context, *ArticleID) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DelArticle not implemented")
}
func (UnimplementedArticleServiceServer) UpdateArticle(context.Context, *Article) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateArticle not implemented")
}
func (UnimplementedArticleServiceServer) ListArticle(context.Context, *ListArticleReq) (*ListArticleRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListArticle not implemented")
}
func (UnimplementedArticleServiceServer) mustEmbedUnimplementedArticleServiceServer() {}

// UnsafeArticleServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ArticleServiceServer will
// result in compilation errors.
type UnsafeArticleServiceServer interface {
	mustEmbedUnimplementedArticleServiceServer()
}

func RegisterArticleServiceServer(s grpc.ServiceRegistrar, srv ArticleServiceServer) {
	s.RegisterService(&ArticleService_ServiceDesc, srv)
}

func _ArticleService_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArticleServiceServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.ArticleService/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArticleServiceServer).Ping(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _ArticleService_DelAuthorByKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Author)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArticleServiceServer).DelAuthorByKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.ArticleService/DelAuthorByKey",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArticleServiceServer).DelAuthorByKey(ctx, req.(*Author))
	}
	return interceptor(ctx, in, info, handler)
}

func _ArticleService_ListArticleMeta_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListArticleMetaReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArticleServiceServer).ListArticleMeta(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.ArticleService/ListArticleMeta",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArticleServiceServer).ListArticleMeta(ctx, req.(*ListArticleMetaReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ArticleService_PutAuthor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Author)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArticleServiceServer).PutAuthor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.ArticleService/PutAuthor",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArticleServiceServer).PutAuthor(ctx, req.(*Author))
	}
	return interceptor(ctx, in, info, handler)
}

func _ArticleService_DelAuthor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthorID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArticleServiceServer).DelAuthor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.ArticleService/DelAuthor",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArticleServiceServer).DelAuthor(ctx, req.(*AuthorID))
	}
	return interceptor(ctx, in, info, handler)
}

func _ArticleService_PutArticle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Article)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArticleServiceServer).PutArticle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.ArticleService/PutArticle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArticleServiceServer).PutArticle(ctx, req.(*Article))
	}
	return interceptor(ctx, in, info, handler)
}

func _ArticleService_GetArticle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ArticleID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArticleServiceServer).GetArticle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.ArticleService/GetArticle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArticleServiceServer).GetArticle(ctx, req.(*ArticleID))
	}
	return interceptor(ctx, in, info, handler)
}

func _ArticleService_DelArticle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ArticleID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArticleServiceServer).DelArticle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.ArticleService/DelArticle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArticleServiceServer).DelArticle(ctx, req.(*ArticleID))
	}
	return interceptor(ctx, in, info, handler)
}

func _ArticleService_UpdateArticle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Article)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArticleServiceServer).UpdateArticle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.ArticleService/UpdateArticle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArticleServiceServer).UpdateArticle(ctx, req.(*Article))
	}
	return interceptor(ctx, in, info, handler)
}

func _ArticleService_ListArticle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListArticleReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArticleServiceServer).ListArticle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.ArticleService/ListArticle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArticleServiceServer).ListArticle(ctx, req.(*ListArticleReq))
	}
	return interceptor(ctx, in, info, handler)
}

// ArticleService_ServiceDesc is the grpc.ServiceDesc for ArticleService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ArticleService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.ArticleService",
	HandlerType: (*ArticleServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _ArticleService_Ping_Handler,
		},
		{
			MethodName: "DelAuthorByKey",
			Handler:    _ArticleService_DelAuthorByKey_Handler,
		},
		{
			MethodName: "ListArticleMeta",
			Handler:    _ArticleService_ListArticleMeta_Handler,
		},
		{
			MethodName: "PutAuthor",
			Handler:    _ArticleService_PutAuthor_Handler,
		},
		{
			MethodName: "DelAuthor",
			Handler:    _ArticleService_DelAuthor_Handler,
		},
		{
			MethodName: "PutArticle",
			Handler:    _ArticleService_PutArticle_Handler,
		},
		{
			MethodName: "GetArticle",
			Handler:    _ArticleService_GetArticle_Handler,
		},
		{
			MethodName: "DelArticle",
			Handler:    _ArticleService_DelArticle_Handler,
		},
		{
			MethodName: "UpdateArticle",
			Handler:    _ArticleService_UpdateArticle_Handler,
		},
		{
			MethodName: "ListArticle",
			Handler:    _ArticleService_ListArticle_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/article.proto",
}
