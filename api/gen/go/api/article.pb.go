// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0-devel
// 	protoc        v3.15.5
// source: api/article.proto

package api

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_article_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_api_article_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_api_article_proto_rawDescGZIP(), []int{0}
}

type Article struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	AuthorID string   `protobuf:"bytes,2,opt,name=authorID,proto3" json:"authorID,omitempty"`
	Title    string   `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
	Tags     []string `protobuf:"bytes,4,rep,name=tags,proto3" json:"tags,omitempty"`
	Brief    string   `protobuf:"bytes,5,opt,name=brief,proto3" json:"brief,omitempty"`
	Content  string   `protobuf:"bytes,6,opt,name=content,proto3" json:"content,omitempty"`
	CreateAt int32    `protobuf:"varint,7,opt,name=createAt,proto3" json:"createAt,omitempty"`
	UpdateAt int32    `protobuf:"varint,8,opt,name=updateAt,proto3" json:"updateAt,omitempty"`
}

func (x *Article) Reset() {
	*x = Article{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_article_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Article) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Article) ProtoMessage() {}

func (x *Article) ProtoReflect() protoreflect.Message {
	mi := &file_api_article_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Article.ProtoReflect.Descriptor instead.
func (*Article) Descriptor() ([]byte, []int) {
	return file_api_article_proto_rawDescGZIP(), []int{1}
}

func (x *Article) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Article) GetAuthorID() string {
	if x != nil {
		return x.AuthorID
	}
	return ""
}

func (x *Article) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Article) GetTags() []string {
	if x != nil {
		return x.Tags
	}
	return nil
}

func (x *Article) GetBrief() string {
	if x != nil {
		return x.Brief
	}
	return ""
}

func (x *Article) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *Article) GetCreateAt() int32 {
	if x != nil {
		return x.CreateAt
	}
	return 0
}

func (x *Article) GetUpdateAt() int32 {
	if x != nil {
		return x.UpdateAt
	}
	return 0
}

type ArticleID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *ArticleID) Reset() {
	*x = ArticleID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_article_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ArticleID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ArticleID) ProtoMessage() {}

func (x *ArticleID) ProtoReflect() protoreflect.Message {
	mi := &file_api_article_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ArticleID.ProtoReflect.Descriptor instead.
func (*ArticleID) Descriptor() ([]byte, []int) {
	return file_api_article_proto_rawDescGZIP(), []int{2}
}

func (x *ArticleID) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type ListArticleReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Offset   int32  `protobuf:"varint,1,opt,name=offset,proto3" json:"offset,omitempty"`
	Limit    int32  `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	AuthorID string `protobuf:"bytes,3,opt,name=authorID,proto3" json:"authorID,omitempty"`
}

func (x *ListArticleReq) Reset() {
	*x = ListArticleReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_article_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListArticleReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListArticleReq) ProtoMessage() {}

func (x *ListArticleReq) ProtoReflect() protoreflect.Message {
	mi := &file_api_article_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListArticleReq.ProtoReflect.Descriptor instead.
func (*ListArticleReq) Descriptor() ([]byte, []int) {
	return file_api_article_proto_rawDescGZIP(), []int{3}
}

func (x *ListArticleReq) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *ListArticleReq) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *ListArticleReq) GetAuthorID() string {
	if x != nil {
		return x.AuthorID
	}
	return ""
}

type ListArticleRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Articles []*Article `protobuf:"bytes,1,rep,name=articles,proto3" json:"articles,omitempty"`
}

func (x *ListArticleRes) Reset() {
	*x = ListArticleRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_article_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListArticleRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListArticleRes) ProtoMessage() {}

func (x *ListArticleRes) ProtoReflect() protoreflect.Message {
	mi := &file_api_article_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListArticleRes.ProtoReflect.Descriptor instead.
func (*ListArticleRes) Descriptor() ([]byte, []int) {
	return file_api_article_proto_rawDescGZIP(), []int{4}
}

func (x *ListArticleRes) GetArticles() []*Article {
	if x != nil {
		return x.Articles
	}
	return nil
}

type Author struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Key  string `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Author) Reset() {
	*x = Author{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_article_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Author) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Author) ProtoMessage() {}

func (x *Author) ProtoReflect() protoreflect.Message {
	mi := &file_api_article_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Author.ProtoReflect.Descriptor instead.
func (*Author) Descriptor() ([]byte, []int) {
	return file_api_article_proto_rawDescGZIP(), []int{5}
}

func (x *Author) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Author) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *Author) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type AuthorID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *AuthorID) Reset() {
	*x = AuthorID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_article_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthorID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthorID) ProtoMessage() {}

func (x *AuthorID) ProtoReflect() protoreflect.Message {
	mi := &file_api_article_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthorID.ProtoReflect.Descriptor instead.
func (*AuthorID) Descriptor() ([]byte, []int) {
	return file_api_article_proto_rawDescGZIP(), []int{6}
}

func (x *AuthorID) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type ListArticleMetaReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Offset   int32  `protobuf:"varint,1,opt,name=offset,proto3" json:"offset,omitempty"`
	Limit    int32  `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	AuthorID string `protobuf:"bytes,3,opt,name=authorID,proto3" json:"authorID,omitempty"`
}

func (x *ListArticleMetaReq) Reset() {
	*x = ListArticleMetaReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_article_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListArticleMetaReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListArticleMetaReq) ProtoMessage() {}

func (x *ListArticleMetaReq) ProtoReflect() protoreflect.Message {
	mi := &file_api_article_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListArticleMetaReq.ProtoReflect.Descriptor instead.
func (*ListArticleMetaReq) Descriptor() ([]byte, []int) {
	return file_api_article_proto_rawDescGZIP(), []int{7}
}

func (x *ListArticleMetaReq) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *ListArticleMetaReq) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *ListArticleMetaReq) GetAuthorID() string {
	if x != nil {
		return x.AuthorID
	}
	return ""
}

type ArticleMeta struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	AuthorID   string   `protobuf:"bytes,2,opt,name=authorID,proto3" json:"authorID,omitempty"`
	AuthorName string   `protobuf:"bytes,3,opt,name=authorName,proto3" json:"authorName,omitempty"`
	Title      string   `protobuf:"bytes,4,opt,name=title,proto3" json:"title,omitempty"`
	Tags       []string `protobuf:"bytes,5,rep,name=tags,proto3" json:"tags,omitempty"`
	Brief      string   `protobuf:"bytes,6,opt,name=brief,proto3" json:"brief,omitempty"`
	CreateAt   int32    `protobuf:"varint,7,opt,name=createAt,proto3" json:"createAt,omitempty"`
	UpdateAt   int32    `protobuf:"varint,8,opt,name=updateAt,proto3" json:"updateAt,omitempty"`
}

func (x *ArticleMeta) Reset() {
	*x = ArticleMeta{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_article_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ArticleMeta) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ArticleMeta) ProtoMessage() {}

func (x *ArticleMeta) ProtoReflect() protoreflect.Message {
	mi := &file_api_article_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ArticleMeta.ProtoReflect.Descriptor instead.
func (*ArticleMeta) Descriptor() ([]byte, []int) {
	return file_api_article_proto_rawDescGZIP(), []int{8}
}

func (x *ArticleMeta) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ArticleMeta) GetAuthorID() string {
	if x != nil {
		return x.AuthorID
	}
	return ""
}

func (x *ArticleMeta) GetAuthorName() string {
	if x != nil {
		return x.AuthorName
	}
	return ""
}

func (x *ArticleMeta) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *ArticleMeta) GetTags() []string {
	if x != nil {
		return x.Tags
	}
	return nil
}

func (x *ArticleMeta) GetBrief() string {
	if x != nil {
		return x.Brief
	}
	return ""
}

func (x *ArticleMeta) GetCreateAt() int32 {
	if x != nil {
		return x.CreateAt
	}
	return 0
}

func (x *ArticleMeta) GetUpdateAt() int32 {
	if x != nil {
		return x.UpdateAt
	}
	return 0
}

type ListArticleMetaRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ArticleMetas []*ArticleMeta `protobuf:"bytes,1,rep,name=articleMetas,proto3" json:"articleMetas,omitempty"`
}

func (x *ListArticleMetaRes) Reset() {
	*x = ListArticleMetaRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_article_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListArticleMetaRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListArticleMetaRes) ProtoMessage() {}

func (x *ListArticleMetaRes) ProtoReflect() protoreflect.Message {
	mi := &file_api_article_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListArticleMetaRes.ProtoReflect.Descriptor instead.
func (*ListArticleMetaRes) Descriptor() ([]byte, []int) {
	return file_api_article_proto_rawDescGZIP(), []int{9}
}

func (x *ListArticleMetaRes) GetArticleMetas() []*ArticleMeta {
	if x != nil {
		return x.ArticleMetas
	}
	return nil
}

var File_api_article_proto protoreflect.FileDescriptor

var file_api_article_proto_rawDesc = []byte{
	0x0a, 0x11, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x03, 0x61, 0x70, 0x69, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22,
	0xc7, 0x01, 0x0a, 0x07, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x61,
	0x75, 0x74, 0x68, 0x6f, 0x72, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x61,
	0x75, 0x74, 0x68, 0x6f, 0x72, 0x49, 0x44, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x74, 0x61, 0x67,
	0x73, 0x12, 0x14, 0x0a, 0x05, 0x62, 0x72, 0x69, 0x65, 0x66, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x62, 0x72, 0x69, 0x65, 0x66, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x74, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x08, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x74, 0x12, 0x1a, 0x0a,
	0x08, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x08, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x74, 0x22, 0x1b, 0x0a, 0x09, 0x41, 0x72, 0x74,
	0x69, 0x63, 0x6c, 0x65, 0x49, 0x44, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x5a, 0x0a, 0x0e, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x72,
	0x74, 0x69, 0x63, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73,
	0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74,
	0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72,
	0x49, 0x44, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72,
	0x49, 0x44, 0x22, 0x3a, 0x0a, 0x0e, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c,
	0x65, 0x52, 0x65, 0x73, 0x12, 0x28, 0x0a, 0x08, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x41, 0x72, 0x74,
	0x69, 0x63, 0x6c, 0x65, 0x52, 0x08, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x73, 0x22, 0x3e,
	0x0a, 0x06, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x1a,
	0x0a, 0x08, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x49, 0x44, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x5e, 0x0a, 0x12, 0x4c, 0x69,
	0x73, 0x74, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x52, 0x65, 0x71,
	0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x1a,
	0x0a, 0x08, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x49, 0x44, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x49, 0x44, 0x22, 0xd1, 0x01, 0x0a, 0x0b, 0x41,
	0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x61, 0x75,
	0x74, 0x68, 0x6f, 0x72, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x61, 0x75,
	0x74, 0x68, 0x6f, 0x72, 0x49, 0x44, 0x12, 0x1e, 0x0a, 0x0a, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72,
	0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x61, 0x75, 0x74, 0x68,
	0x6f, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x74, 0x61, 0x67, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x74, 0x61, 0x67, 0x73,
	0x12, 0x14, 0x0a, 0x05, 0x62, 0x72, 0x69, 0x65, 0x66, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x62, 0x72, 0x69, 0x65, 0x66, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x41, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x41, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x74, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x74, 0x22, 0x4a,
	0x0a, 0x12, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x4d, 0x65, 0x74,
	0x61, 0x52, 0x65, 0x73, 0x12, 0x34, 0x0a, 0x0c, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x4d,
	0x65, 0x74, 0x61, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x52, 0x0c, 0x61, 0x72,
	0x74, 0x69, 0x63, 0x6c, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x73, 0x32, 0x9f, 0x06, 0x0a, 0x0e, 0x41,
	0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x2d, 0x0a,
	0x04, 0x50, 0x69, 0x6e, 0x67, 0x12, 0x0a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x1a, 0x0a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x0d, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x07, 0x12, 0x05, 0x2f, 0x70, 0x69, 0x6e, 0x67, 0x12, 0x51, 0x0a, 0x11,
	0x41, 0x64, 0x64, 0x4f, 0x72, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x75, 0x74, 0x68, 0x6f,
	0x72, 0x12, 0x0b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x1a, 0x0d,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x49, 0x44, 0x22, 0x20, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x1a, 0x22, 0x15, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x64, 0x64, 0x4f, 0x72,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x3a, 0x01, 0x2a, 0x12,
	0x48, 0x0a, 0x0e, 0x44, 0x65, 0x6c, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x42, 0x79, 0x4b, 0x65,
	0x79, 0x12, 0x0b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x1a, 0x0a,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x1d, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x17, 0x22, 0x12, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x65, 0x6c, 0x41, 0x75, 0x74, 0x68, 0x6f,
	0x72, 0x42, 0x79, 0x4b, 0x65, 0x79, 0x3a, 0x01, 0x2a, 0x12, 0x60, 0x0a, 0x0f, 0x4c, 0x69, 0x73,
	0x74, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x12, 0x17, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x4d, 0x65,
	0x74, 0x61, 0x52, 0x65, 0x71, 0x1a, 0x17, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4c, 0x69, 0x73, 0x74,
	0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x52, 0x65, 0x73, 0x22, 0x1b,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x15, 0x12, 0x13, 0x2f, 0x76, 0x31, 0x2f, 0x6c, 0x69, 0x73, 0x74,
	0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x12, 0x3e, 0x0a, 0x09, 0x50,
	0x75, 0x74, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x12, 0x0b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x41,
	0x75, 0x74, 0x68, 0x6f, 0x72, 0x1a, 0x0d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x41, 0x75, 0x74, 0x68,
	0x6f, 0x72, 0x49, 0x44, 0x22, 0x15, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0f, 0x22, 0x0a, 0x2f, 0x76,
	0x31, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x3a, 0x01, 0x2a, 0x12, 0x3a, 0x0a, 0x09, 0x44,
	0x65, 0x6c, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x12, 0x0d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x41,
	0x75, 0x74, 0x68, 0x6f, 0x72, 0x49, 0x44, 0x1a, 0x0a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x22, 0x12, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0c, 0x2a, 0x0a, 0x2f, 0x76, 0x31,
	0x2f, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x12, 0x42, 0x0a, 0x0a, 0x50, 0x75, 0x74, 0x41, 0x72,
	0x74, 0x69, 0x63, 0x6c, 0x65, 0x12, 0x0c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x41, 0x72, 0x74, 0x69,
	0x63, 0x6c, 0x65, 0x1a, 0x0e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c,
	0x65, 0x49, 0x44, 0x22, 0x16, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x10, 0x22, 0x0b, 0x2f, 0x76, 0x31,
	0x2f, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x3a, 0x01, 0x2a, 0x12, 0x44, 0x0a, 0x0a, 0x47,
	0x65, 0x74, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x12, 0x0e, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x49, 0x44, 0x1a, 0x0c, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x22, 0x18, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x12, 0x12,
	0x10, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x2f, 0x7b, 0x69, 0x64,
	0x7d, 0x12, 0x42, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x12,
	0x0e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x49, 0x44, 0x1a,
	0x0a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x18, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x12, 0x2a, 0x10, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65,
	0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0x46, 0x0a, 0x0d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41,
	0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x12, 0x0c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x41, 0x72, 0x74,
	0x69, 0x63, 0x6c, 0x65, 0x1a, 0x0a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x22, 0x1b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x15, 0x1a, 0x10, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x72,
	0x74, 0x69, 0x63, 0x6c, 0x65, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x3a, 0x01, 0x2a, 0x12, 0x4d, 0x0a,
	0x0b, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x12, 0x13, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x52, 0x65,
	0x71, 0x1a, 0x13, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x72, 0x74, 0x69,
	0x63, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x22, 0x14, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0e, 0x12, 0x0c,
	0x2f, 0x76, 0x31, 0x2f, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x73, 0x42, 0x2d, 0x5a, 0x2b,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x68, 0x61, 0x74, 0x6c, 0x6f,
	0x6e, 0x65, 0x6c, 0x79, 0x2f, 0x72, 0x70, 0x63, 0x2d, 0x6f, 0x70, 0x73, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_api_article_proto_rawDescOnce sync.Once
	file_api_article_proto_rawDescData = file_api_article_proto_rawDesc
)

func file_api_article_proto_rawDescGZIP() []byte {
	file_api_article_proto_rawDescOnce.Do(func() {
		file_api_article_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_article_proto_rawDescData)
	})
	return file_api_article_proto_rawDescData
}

var file_api_article_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_api_article_proto_goTypes = []interface{}{
	(*Empty)(nil),              // 0: api.Empty
	(*Article)(nil),            // 1: api.Article
	(*ArticleID)(nil),          // 2: api.ArticleID
	(*ListArticleReq)(nil),     // 3: api.ListArticleReq
	(*ListArticleRes)(nil),     // 4: api.ListArticleRes
	(*Author)(nil),             // 5: api.Author
	(*AuthorID)(nil),           // 6: api.AuthorID
	(*ListArticleMetaReq)(nil), // 7: api.ListArticleMetaReq
	(*ArticleMeta)(nil),        // 8: api.ArticleMeta
	(*ListArticleMetaRes)(nil), // 9: api.ListArticleMetaRes
}
var file_api_article_proto_depIdxs = []int32{
	1,  // 0: api.ListArticleRes.articles:type_name -> api.Article
	8,  // 1: api.ListArticleMetaRes.articleMetas:type_name -> api.ArticleMeta
	0,  // 2: api.ArticleService.Ping:input_type -> api.Empty
	5,  // 3: api.ArticleService.AddOrUpdateAuthor:input_type -> api.Author
	5,  // 4: api.ArticleService.DelAuthorByKey:input_type -> api.Author
	7,  // 5: api.ArticleService.ListArticleMeta:input_type -> api.ListArticleMetaReq
	5,  // 6: api.ArticleService.PutAuthor:input_type -> api.Author
	6,  // 7: api.ArticleService.DelAuthor:input_type -> api.AuthorID
	1,  // 8: api.ArticleService.PutArticle:input_type -> api.Article
	2,  // 9: api.ArticleService.GetArticle:input_type -> api.ArticleID
	2,  // 10: api.ArticleService.DelArticle:input_type -> api.ArticleID
	1,  // 11: api.ArticleService.UpdateArticle:input_type -> api.Article
	3,  // 12: api.ArticleService.ListArticle:input_type -> api.ListArticleReq
	0,  // 13: api.ArticleService.Ping:output_type -> api.Empty
	6,  // 14: api.ArticleService.AddOrUpdateAuthor:output_type -> api.AuthorID
	0,  // 15: api.ArticleService.DelAuthorByKey:output_type -> api.Empty
	9,  // 16: api.ArticleService.ListArticleMeta:output_type -> api.ListArticleMetaRes
	6,  // 17: api.ArticleService.PutAuthor:output_type -> api.AuthorID
	0,  // 18: api.ArticleService.DelAuthor:output_type -> api.Empty
	2,  // 19: api.ArticleService.PutArticle:output_type -> api.ArticleID
	1,  // 20: api.ArticleService.GetArticle:output_type -> api.Article
	0,  // 21: api.ArticleService.DelArticle:output_type -> api.Empty
	0,  // 22: api.ArticleService.UpdateArticle:output_type -> api.Empty
	4,  // 23: api.ArticleService.ListArticle:output_type -> api.ListArticleRes
	13, // [13:24] is the sub-list for method output_type
	2,  // [2:13] is the sub-list for method input_type
	2,  // [2:2] is the sub-list for extension type_name
	2,  // [2:2] is the sub-list for extension extendee
	0,  // [0:2] is the sub-list for field type_name
}

func init() { file_api_article_proto_init() }
func file_api_article_proto_init() {
	if File_api_article_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_article_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_article_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Article); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_article_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ArticleID); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_article_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListArticleReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_article_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListArticleRes); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_article_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Author); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_article_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthorID); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_article_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListArticleMetaReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_article_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ArticleMeta); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_article_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListArticleMetaRes); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_article_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_article_proto_goTypes,
		DependencyIndexes: file_api_article_proto_depIdxs,
		MessageInfos:      file_api_article_proto_msgTypes,
	}.Build()
	File_api_article_proto = out.File
	file_api_article_proto_rawDesc = nil
	file_api_article_proto_goTypes = nil
	file_api_article_proto_depIdxs = nil
}
