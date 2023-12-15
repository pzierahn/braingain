// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.25.1
// source: documents.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ChunkIDs struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items []string `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *ChunkIDs) Reset() {
	*x = ChunkIDs{}
	if protoimpl.UnsafeEnabled {
		mi := &file_documents_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChunkIDs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChunkIDs) ProtoMessage() {}

func (x *ChunkIDs) ProtoReflect() protoreflect.Message {
	mi := &file_documents_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChunkIDs.ProtoReflect.Descriptor instead.
func (*ChunkIDs) Descriptor() ([]byte, []int) {
	return file_documents_proto_rawDescGZIP(), []int{0}
}

func (x *ChunkIDs) GetItems() []string {
	if x != nil {
		return x.Items
	}
	return nil
}

type Chunk struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	DocumentId string `protobuf:"bytes,2,opt,name=document_id,json=documentId,proto3" json:"document_id,omitempty"`
	Filename   string `protobuf:"bytes,3,opt,name=filename,proto3" json:"filename,omitempty"`
	Page       uint32 `protobuf:"varint,4,opt,name=page,proto3" json:"page,omitempty"`
}

func (x *Chunk) Reset() {
	*x = Chunk{}
	if protoimpl.UnsafeEnabled {
		mi := &file_documents_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Chunk) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Chunk) ProtoMessage() {}

func (x *Chunk) ProtoReflect() protoreflect.Message {
	mi := &file_documents_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Chunk.ProtoReflect.Descriptor instead.
func (*Chunk) Descriptor() ([]byte, []int) {
	return file_documents_proto_rawDescGZIP(), []int{1}
}

func (x *Chunk) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Chunk) GetDocumentId() string {
	if x != nil {
		return x.DocumentId
	}
	return ""
}

func (x *Chunk) GetFilename() string {
	if x != nil {
		return x.Filename
	}
	return ""
}

func (x *Chunk) GetPage() uint32 {
	if x != nil {
		return x.Page
	}
	return 0
}

type Chunks struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items []*Chunk `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *Chunks) Reset() {
	*x = Chunks{}
	if protoimpl.UnsafeEnabled {
		mi := &file_documents_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Chunks) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Chunks) ProtoMessage() {}

func (x *Chunks) ProtoReflect() protoreflect.Message {
	mi := &file_documents_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Chunks.ProtoReflect.Descriptor instead.
func (*Chunks) Descriptor() ([]byte, []int) {
	return file_documents_proto_rawDescGZIP(), []int{2}
}

func (x *Chunks) GetItems() []*Chunk {
	if x != nil {
		return x.Items
	}
	return nil
}

type SearchQuery struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Query        string  `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`
	CollectionId string  `protobuf:"bytes,2,opt,name=collection_id,json=collectionId,proto3" json:"collection_id,omitempty"`
	Threshold    float32 `protobuf:"fixed32,3,opt,name=threshold,proto3" json:"threshold,omitempty"`
	Limit        uint32  `protobuf:"varint,4,opt,name=limit,proto3" json:"limit,omitempty"`
}

func (x *SearchQuery) Reset() {
	*x = SearchQuery{}
	if protoimpl.UnsafeEnabled {
		mi := &file_documents_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchQuery) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchQuery) ProtoMessage() {}

func (x *SearchQuery) ProtoReflect() protoreflect.Message {
	mi := &file_documents_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchQuery.ProtoReflect.Descriptor instead.
func (*SearchQuery) Descriptor() ([]byte, []int) {
	return file_documents_proto_rawDescGZIP(), []int{3}
}

func (x *SearchQuery) GetQuery() string {
	if x != nil {
		return x.Query
	}
	return ""
}

func (x *SearchQuery) GetCollectionId() string {
	if x != nil {
		return x.CollectionId
	}
	return ""
}

func (x *SearchQuery) GetThreshold() float32 {
	if x != nil {
		return x.Threshold
	}
	return 0
}

func (x *SearchQuery) GetLimit() uint32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type SearchResults struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items []*SearchResults_Document `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *SearchResults) Reset() {
	*x = SearchResults{}
	if protoimpl.UnsafeEnabled {
		mi := &file_documents_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchResults) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchResults) ProtoMessage() {}

func (x *SearchResults) ProtoReflect() protoreflect.Message {
	mi := &file_documents_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchResults.ProtoReflect.Descriptor instead.
func (*SearchResults) Descriptor() ([]byte, []int) {
	return file_documents_proto_rawDescGZIP(), []int{4}
}

func (x *SearchResults) GetItems() []*SearchResults_Document {
	if x != nil {
		return x.Items
	}
	return nil
}

type IndexProgress struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TotalPages     uint32 `protobuf:"varint,1,opt,name=totalPages,proto3" json:"totalPages,omitempty"`
	ProcessedPages uint32 `protobuf:"varint,2,opt,name=processedPages,proto3" json:"processedPages,omitempty"`
}

func (x *IndexProgress) Reset() {
	*x = IndexProgress{}
	if protoimpl.UnsafeEnabled {
		mi := &file_documents_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IndexProgress) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IndexProgress) ProtoMessage() {}

func (x *IndexProgress) ProtoReflect() protoreflect.Message {
	mi := &file_documents_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IndexProgress.ProtoReflect.Descriptor instead.
func (*IndexProgress) Descriptor() ([]byte, []int) {
	return file_documents_proto_rawDescGZIP(), []int{5}
}

func (x *IndexProgress) GetTotalPages() uint32 {
	if x != nil {
		return x.TotalPages
	}
	return 0
}

func (x *IndexProgress) GetProcessedPages() uint32 {
	if x != nil {
		return x.ProcessedPages
	}
	return 0
}

type Document struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	CollectionId string `protobuf:"bytes,2,opt,name=collection_id,json=collectionId,proto3" json:"collection_id,omitempty"`
	Filename     string `protobuf:"bytes,3,opt,name=filename,proto3" json:"filename,omitempty"`
	Path         string `protobuf:"bytes,4,opt,name=path,proto3" json:"path,omitempty"`
}

func (x *Document) Reset() {
	*x = Document{}
	if protoimpl.UnsafeEnabled {
		mi := &file_documents_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Document) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Document) ProtoMessage() {}

func (x *Document) ProtoReflect() protoreflect.Message {
	mi := &file_documents_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Document.ProtoReflect.Descriptor instead.
func (*Document) Descriptor() ([]byte, []int) {
	return file_documents_proto_rawDescGZIP(), []int{6}
}

func (x *Document) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Document) GetCollectionId() string {
	if x != nil {
		return x.CollectionId
	}
	return ""
}

func (x *Document) GetFilename() string {
	if x != nil {
		return x.Filename
	}
	return ""
}

func (x *Document) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

type DocumentFilter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Query        string `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`
	CollectionId string `protobuf:"bytes,2,opt,name=collection_id,json=collectionId,proto3" json:"collection_id,omitempty"`
}

func (x *DocumentFilter) Reset() {
	*x = DocumentFilter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_documents_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DocumentFilter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DocumentFilter) ProtoMessage() {}

func (x *DocumentFilter) ProtoReflect() protoreflect.Message {
	mi := &file_documents_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DocumentFilter.ProtoReflect.Descriptor instead.
func (*DocumentFilter) Descriptor() ([]byte, []int) {
	return file_documents_proto_rawDescGZIP(), []int{7}
}

func (x *DocumentFilter) GetQuery() string {
	if x != nil {
		return x.Query
	}
	return ""
}

func (x *DocumentFilter) GetCollectionId() string {
	if x != nil {
		return x.CollectionId
	}
	return ""
}

type Documents struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items []*Documents_Document `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *Documents) Reset() {
	*x = Documents{}
	if protoimpl.UnsafeEnabled {
		mi := &file_documents_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Documents) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Documents) ProtoMessage() {}

func (x *Documents) ProtoReflect() protoreflect.Message {
	mi := &file_documents_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Documents.ProtoReflect.Descriptor instead.
func (*Documents) Descriptor() ([]byte, []int) {
	return file_documents_proto_rawDescGZIP(), []int{8}
}

func (x *Documents) GetItems() []*Documents_Document {
	if x != nil {
		return x.Items
	}
	return nil
}

type SearchResults_Document struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         string  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	DocumentId string  `protobuf:"bytes,2,opt,name=document_id,json=documentId,proto3" json:"document_id,omitempty"`
	Filename   string  `protobuf:"bytes,3,opt,name=filename,proto3" json:"filename,omitempty"`
	Content    string  `protobuf:"bytes,4,opt,name=content,proto3" json:"content,omitempty"`
	Page       uint32  `protobuf:"varint,5,opt,name=page,proto3" json:"page,omitempty"`
	Score      float32 `protobuf:"fixed32,6,opt,name=score,proto3" json:"score,omitempty"`
}

func (x *SearchResults_Document) Reset() {
	*x = SearchResults_Document{}
	if protoimpl.UnsafeEnabled {
		mi := &file_documents_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchResults_Document) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchResults_Document) ProtoMessage() {}

func (x *SearchResults_Document) ProtoReflect() protoreflect.Message {
	mi := &file_documents_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchResults_Document.ProtoReflect.Descriptor instead.
func (*SearchResults_Document) Descriptor() ([]byte, []int) {
	return file_documents_proto_rawDescGZIP(), []int{4, 0}
}

func (x *SearchResults_Document) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *SearchResults_Document) GetDocumentId() string {
	if x != nil {
		return x.DocumentId
	}
	return ""
}

func (x *SearchResults_Document) GetFilename() string {
	if x != nil {
		return x.Filename
	}
	return ""
}

func (x *SearchResults_Document) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *SearchResults_Document) GetPage() uint32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *SearchResults_Document) GetScore() float32 {
	if x != nil {
		return x.Score
	}
	return 0
}

type Documents_Document struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Filename string `protobuf:"bytes,2,opt,name=filename,proto3" json:"filename,omitempty"`
	Pages    uint32 `protobuf:"varint,3,opt,name=pages,proto3" json:"pages,omitempty"`
}

func (x *Documents_Document) Reset() {
	*x = Documents_Document{}
	if protoimpl.UnsafeEnabled {
		mi := &file_documents_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Documents_Document) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Documents_Document) ProtoMessage() {}

func (x *Documents_Document) ProtoReflect() protoreflect.Message {
	mi := &file_documents_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Documents_Document.ProtoReflect.Descriptor instead.
func (*Documents_Document) Descriptor() ([]byte, []int) {
	return file_documents_proto_rawDescGZIP(), []int{8, 0}
}

func (x *Documents_Document) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Documents_Document) GetFilename() string {
	if x != nil {
		return x.Filename
	}
	return ""
}

func (x *Documents_Document) GetPages() uint32 {
	if x != nil {
		return x.Pages
	}
	return 0
}

var File_documents_proto protoreflect.FileDescriptor

var file_documents_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x20, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x62, 0x72, 0x61, 0x69,
	0x6e, 0x62, 0x6f, 0x6f, 0x73, 0x74, 0x2e, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x73,
	0x2e, 0x76, 0x31, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x20, 0x0a, 0x08, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x49, 0x44, 0x73, 0x12, 0x14, 0x0a, 0x05,
	0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x69, 0x74, 0x65,
	0x6d, 0x73, 0x22, 0x68, 0x0a, 0x05, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x64,
	0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08,
	0x66, 0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x66, 0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x22, 0x47, 0x0a, 0x06,
	0x43, 0x68, 0x75, 0x6e, 0x6b, 0x73, 0x12, 0x3d, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74,
	0x2e, 0x62, 0x72, 0x61, 0x69, 0x6e, 0x62, 0x6f, 0x6f, 0x73, 0x74, 0x2e, 0x64, 0x6f, 0x63, 0x75,
	0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x52, 0x05,
	0x69, 0x74, 0x65, 0x6d, 0x73, 0x22, 0x7c, 0x0a, 0x0b, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x51,
	0x75, 0x65, 0x72, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x12, 0x23, 0x0a, 0x0d, 0x63, 0x6f,
	0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0c, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12,
	0x1c, 0x0a, 0x09, 0x74, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x02, 0x52, 0x09, 0x74, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x12, 0x14, 0x0a,
	0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x6c, 0x69,
	0x6d, 0x69, 0x74, 0x22, 0xfd, 0x01, 0x0a, 0x0d, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x73, 0x12, 0x4e, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x38, 0x2e, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2e,
	0x62, 0x72, 0x61, 0x69, 0x6e, 0x62, 0x6f, 0x6f, 0x73, 0x74, 0x2e, 0x64, 0x6f, 0x63, 0x75, 0x6d,
	0x65, 0x6e, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x73, 0x2e, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x05,
	0x69, 0x74, 0x65, 0x6d, 0x73, 0x1a, 0x9b, 0x01, 0x0a, 0x08, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65,
	0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e,
	0x74, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67,
	0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x73, 0x63,
	0x6f, 0x72, 0x65, 0x22, 0x57, 0x0a, 0x0d, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x50, 0x72, 0x6f, 0x67,
	0x72, 0x65, 0x73, 0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x50, 0x61, 0x67,
	0x65, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x50,
	0x61, 0x67, 0x65, 0x73, 0x12, 0x26, 0x0a, 0x0e, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x65,
	0x64, 0x50, 0x61, 0x67, 0x65, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0e, 0x70, 0x72,
	0x6f, 0x63, 0x65, 0x73, 0x73, 0x65, 0x64, 0x50, 0x61, 0x67, 0x65, 0x73, 0x22, 0x6f, 0x0a, 0x08,
	0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x63, 0x6f, 0x6c, 0x6c,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0c, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x1a, 0x0a,
	0x08, 0x66, 0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x66, 0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74,
	0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x22, 0x4b, 0x0a,
	0x0e, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12,
	0x14, 0x0a, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x71, 0x75, 0x65, 0x72, 0x79, 0x12, 0x23, 0x0a, 0x0d, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x6f,
	0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x22, 0xa5, 0x01, 0x0a, 0x09, 0x44,
	0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x4a, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x34, 0x2e, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69,
	0x6e, 0x74, 0x2e, 0x62, 0x72, 0x61, 0x69, 0x6e, 0x62, 0x6f, 0x6f, 0x73, 0x74, 0x2e, 0x64, 0x6f,
	0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x6f, 0x63, 0x75, 0x6d,
	0x65, 0x6e, 0x74, 0x73, 0x2e, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x05, 0x69,
	0x74, 0x65, 0x6d, 0x73, 0x1a, 0x4c, 0x0a, 0x08, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x1a, 0x0a, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x70, 0x61, 0x67, 0x65, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x70, 0x61, 0x67,
	0x65, 0x73, 0x32, 0xc9, 0x04, 0x0a, 0x0f, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x65, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x30,
	0x2e, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x62, 0x72, 0x61, 0x69, 0x6e, 0x62,
	0x6f, 0x6f, 0x73, 0x74, 0x2e, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x76,
	0x31, 0x2e, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72,
	0x1a, 0x2b, 0x2e, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x62, 0x72, 0x61, 0x69,
	0x6e, 0x62, 0x6f, 0x6f, 0x73, 0x74, 0x2e, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x73,
	0x2e, 0x76, 0x31, 0x2e, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x66, 0x0a,
	0x05, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x2a, 0x2e, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e,
	0x74, 0x2e, 0x62, 0x72, 0x61, 0x69, 0x6e, 0x62, 0x6f, 0x6f, 0x73, 0x74, 0x2e, 0x64, 0x6f, 0x63,
	0x75, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65,
	0x6e, 0x74, 0x1a, 0x2f, 0x2e, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x62, 0x72,
	0x61, 0x69, 0x6e, 0x62, 0x6f, 0x6f, 0x73, 0x74, 0x2e, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e,
	0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x50, 0x72, 0x6f, 0x67, 0x72,
	0x65, 0x73, 0x73, 0x30, 0x01, 0x12, 0x4c, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12,
	0x2a, 0x2e, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x62, 0x72, 0x61, 0x69, 0x6e,
	0x62, 0x6f, 0x6f, 0x73, 0x74, 0x2e, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2e,
	0x76, 0x31, 0x2e, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x12, 0x4c, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x2a, 0x2e,
	0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x62, 0x72, 0x61, 0x69, 0x6e, 0x62, 0x6f,
	0x6f, 0x73, 0x74, 0x2e, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x76, 0x31,
	0x2e, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x12, 0x68, 0x0a, 0x06, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x12, 0x2d, 0x2e, 0x65, 0x6e,
	0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x62, 0x72, 0x61, 0x69, 0x6e, 0x62, 0x6f, 0x6f, 0x73,
	0x74, 0x2e, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53,
	0x65, 0x61, 0x72, 0x63, 0x68, 0x51, 0x75, 0x65, 0x72, 0x79, 0x1a, 0x2f, 0x2e, 0x65, 0x6e, 0x64,
	0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x62, 0x72, 0x61, 0x69, 0x6e, 0x62, 0x6f, 0x6f, 0x73, 0x74,
	0x2e, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65,
	0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x12, 0x61, 0x0a, 0x09, 0x47,
	0x65, 0x74, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x73, 0x12, 0x2a, 0x2e, 0x65, 0x6e, 0x64, 0x70, 0x6f,
	0x69, 0x6e, 0x74, 0x2e, 0x62, 0x72, 0x61, 0x69, 0x6e, 0x62, 0x6f, 0x6f, 0x73, 0x74, 0x2e, 0x64,
	0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x68, 0x75, 0x6e,
	0x6b, 0x49, 0x44, 0x73, 0x1a, 0x28, 0x2e, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2e,
	0x62, 0x72, 0x61, 0x69, 0x6e, 0x62, 0x6f, 0x6f, 0x73, 0x74, 0x2e, 0x64, 0x6f, 0x63, 0x75, 0x6d,
	0x65, 0x6e, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x73, 0x42, 0x09,
	0x5a, 0x07, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_documents_proto_rawDescOnce sync.Once
	file_documents_proto_rawDescData = file_documents_proto_rawDesc
)

func file_documents_proto_rawDescGZIP() []byte {
	file_documents_proto_rawDescOnce.Do(func() {
		file_documents_proto_rawDescData = protoimpl.X.CompressGZIP(file_documents_proto_rawDescData)
	})
	return file_documents_proto_rawDescData
}

var file_documents_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_documents_proto_goTypes = []interface{}{
	(*ChunkIDs)(nil),               // 0: endpoint.brainboost.documents.v1.ChunkIDs
	(*Chunk)(nil),                  // 1: endpoint.brainboost.documents.v1.Chunk
	(*Chunks)(nil),                 // 2: endpoint.brainboost.documents.v1.Chunks
	(*SearchQuery)(nil),            // 3: endpoint.brainboost.documents.v1.SearchQuery
	(*SearchResults)(nil),          // 4: endpoint.brainboost.documents.v1.SearchResults
	(*IndexProgress)(nil),          // 5: endpoint.brainboost.documents.v1.IndexProgress
	(*Document)(nil),               // 6: endpoint.brainboost.documents.v1.Document
	(*DocumentFilter)(nil),         // 7: endpoint.brainboost.documents.v1.DocumentFilter
	(*Documents)(nil),              // 8: endpoint.brainboost.documents.v1.Documents
	(*SearchResults_Document)(nil), // 9: endpoint.brainboost.documents.v1.SearchResults.Document
	(*Documents_Document)(nil),     // 10: endpoint.brainboost.documents.v1.Documents.Document
	(*emptypb.Empty)(nil),          // 11: google.protobuf.Empty
}
var file_documents_proto_depIdxs = []int32{
	1,  // 0: endpoint.brainboost.documents.v1.Chunks.items:type_name -> endpoint.brainboost.documents.v1.Chunk
	9,  // 1: endpoint.brainboost.documents.v1.SearchResults.items:type_name -> endpoint.brainboost.documents.v1.SearchResults.Document
	10, // 2: endpoint.brainboost.documents.v1.Documents.items:type_name -> endpoint.brainboost.documents.v1.Documents.Document
	7,  // 3: endpoint.brainboost.documents.v1.DocumentService.List:input_type -> endpoint.brainboost.documents.v1.DocumentFilter
	6,  // 4: endpoint.brainboost.documents.v1.DocumentService.Index:input_type -> endpoint.brainboost.documents.v1.Document
	6,  // 5: endpoint.brainboost.documents.v1.DocumentService.Delete:input_type -> endpoint.brainboost.documents.v1.Document
	6,  // 6: endpoint.brainboost.documents.v1.DocumentService.Update:input_type -> endpoint.brainboost.documents.v1.Document
	3,  // 7: endpoint.brainboost.documents.v1.DocumentService.Search:input_type -> endpoint.brainboost.documents.v1.SearchQuery
	0,  // 8: endpoint.brainboost.documents.v1.DocumentService.GetChunks:input_type -> endpoint.brainboost.documents.v1.ChunkIDs
	8,  // 9: endpoint.brainboost.documents.v1.DocumentService.List:output_type -> endpoint.brainboost.documents.v1.Documents
	5,  // 10: endpoint.brainboost.documents.v1.DocumentService.Index:output_type -> endpoint.brainboost.documents.v1.IndexProgress
	11, // 11: endpoint.brainboost.documents.v1.DocumentService.Delete:output_type -> google.protobuf.Empty
	11, // 12: endpoint.brainboost.documents.v1.DocumentService.Update:output_type -> google.protobuf.Empty
	4,  // 13: endpoint.brainboost.documents.v1.DocumentService.Search:output_type -> endpoint.brainboost.documents.v1.SearchResults
	2,  // 14: endpoint.brainboost.documents.v1.DocumentService.GetChunks:output_type -> endpoint.brainboost.documents.v1.Chunks
	9,  // [9:15] is the sub-list for method output_type
	3,  // [3:9] is the sub-list for method input_type
	3,  // [3:3] is the sub-list for extension type_name
	3,  // [3:3] is the sub-list for extension extendee
	0,  // [0:3] is the sub-list for field type_name
}

func init() { file_documents_proto_init() }
func file_documents_proto_init() {
	if File_documents_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_documents_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChunkIDs); i {
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
		file_documents_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Chunk); i {
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
		file_documents_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Chunks); i {
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
		file_documents_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchQuery); i {
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
		file_documents_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchResults); i {
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
		file_documents_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IndexProgress); i {
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
		file_documents_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Document); i {
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
		file_documents_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DocumentFilter); i {
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
		file_documents_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Documents); i {
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
		file_documents_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchResults_Document); i {
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
		file_documents_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Documents_Document); i {
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
			RawDescriptor: file_documents_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_documents_proto_goTypes,
		DependencyIndexes: file_documents_proto_depIdxs,
		MessageInfos:      file_documents_proto_msgTypes,
	}.Build()
	File_documents_proto = out.File
	file_documents_proto_rawDesc = nil
	file_documents_proto_goTypes = nil
	file_documents_proto_depIdxs = nil
}
