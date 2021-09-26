// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: searchterms.proto

package main

import (
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

type SearchTermRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Term string `protobuf:"bytes,1,opt,name=term,proto3" json:"term,omitempty"`
}

func (x *SearchTermRequest) Reset() {
	*x = SearchTermRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_searchterms_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchTermRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchTermRequest) ProtoMessage() {}

func (x *SearchTermRequest) ProtoReflect() protoreflect.Message {
	mi := &file_searchterms_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchTermRequest.ProtoReflect.Descriptor instead.
func (*SearchTermRequest) Descriptor() ([]byte, []int) {
	return file_searchterms_proto_rawDescGZIP(), []int{0}
}

func (x *SearchTermRequest) GetTerm() string {
	if x != nil {
		return x.Term
	}
	return ""
}

type SearchTermResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *SearchTermResponse) Reset() {
	*x = SearchTermResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_searchterms_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchTermResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchTermResponse) ProtoMessage() {}

func (x *SearchTermResponse) ProtoReflect() protoreflect.Message {
	mi := &file_searchterms_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchTermResponse.ProtoReflect.Descriptor instead.
func (*SearchTermResponse) Descriptor() ([]byte, []int) {
	return file_searchterms_proto_rawDescGZIP(), []int{1}
}

func (x *SearchTermResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type UpdateSearchTermsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Term string `protobuf:"bytes,1,opt,name=term,proto3" json:"term,omitempty"`
}

func (x *UpdateSearchTermsRequest) Reset() {
	*x = UpdateSearchTermsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_searchterms_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateSearchTermsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateSearchTermsRequest) ProtoMessage() {}

func (x *UpdateSearchTermsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_searchterms_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateSearchTermsRequest.ProtoReflect.Descriptor instead.
func (*UpdateSearchTermsRequest) Descriptor() ([]byte, []int) {
	return file_searchterms_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateSearchTermsRequest) GetTerm() string {
	if x != nil {
		return x.Term
	}
	return ""
}

type UpdateSearchTermsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateSearchTermsResponse) Reset() {
	*x = UpdateSearchTermsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_searchterms_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateSearchTermsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateSearchTermsResponse) ProtoMessage() {}

func (x *UpdateSearchTermsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_searchterms_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateSearchTermsResponse.ProtoReflect.Descriptor instead.
func (*UpdateSearchTermsResponse) Descriptor() ([]byte, []int) {
	return file_searchterms_proto_rawDescGZIP(), []int{3}
}

type SearchTermMetricsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SearchTermMetricsRequest) Reset() {
	*x = SearchTermMetricsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_searchterms_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchTermMetricsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchTermMetricsRequest) ProtoMessage() {}

func (x *SearchTermMetricsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_searchterms_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchTermMetricsRequest.ProtoReflect.Descriptor instead.
func (*SearchTermMetricsRequest) Descriptor() ([]byte, []int) {
	return file_searchterms_proto_rawDescGZIP(), []int{4}
}

type SearchTermMetrics struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Term        string `protobuf:"bytes,1,opt,name=term,proto3" json:"term,omitempty"`
	SearchCount int32  `protobuf:"varint,2,opt,name=searchCount,proto3" json:"searchCount,omitempty"`
}

func (x *SearchTermMetrics) Reset() {
	*x = SearchTermMetrics{}
	if protoimpl.UnsafeEnabled {
		mi := &file_searchterms_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchTermMetrics) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchTermMetrics) ProtoMessage() {}

func (x *SearchTermMetrics) ProtoReflect() protoreflect.Message {
	mi := &file_searchterms_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchTermMetrics.ProtoReflect.Descriptor instead.
func (*SearchTermMetrics) Descriptor() ([]byte, []int) {
	return file_searchterms_proto_rawDescGZIP(), []int{5}
}

func (x *SearchTermMetrics) GetTerm() string {
	if x != nil {
		return x.Term
	}
	return ""
}

func (x *SearchTermMetrics) GetSearchCount() int32 {
	if x != nil {
		return x.SearchCount
	}
	return 0
}

type SearchTermMetricsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Results []*SearchTermMetrics `protobuf:"bytes,1,rep,name=results,proto3" json:"results,omitempty"`
}

func (x *SearchTermMetricsResponse) Reset() {
	*x = SearchTermMetricsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_searchterms_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchTermMetricsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchTermMetricsResponse) ProtoMessage() {}

func (x *SearchTermMetricsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_searchterms_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchTermMetricsResponse.ProtoReflect.Descriptor instead.
func (*SearchTermMetricsResponse) Descriptor() ([]byte, []int) {
	return file_searchterms_proto_rawDescGZIP(), []int{6}
}

func (x *SearchTermMetricsResponse) GetResults() []*SearchTermMetrics {
	if x != nil {
		return x.Results
	}
	return nil
}

var File_searchterms_proto protoreflect.FileDescriptor

var file_searchterms_proto_rawDesc = []byte{
	0x0a, 0x11, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x74, 0x65, 0x72, 0x6d, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x13, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x74, 0x65, 0x72, 0x6d, 0x73,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x22, 0x27, 0x0a, 0x11, 0x53, 0x65, 0x61, 0x72,
	0x63, 0x68, 0x54, 0x65, 0x72, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a,
	0x04, 0x74, 0x65, 0x72, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x72,
	0x6d, 0x22, 0x2e, 0x0a, 0x12, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x54, 0x65, 0x72, 0x6d, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x22, 0x2e, 0x0a, 0x18, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x65, 0x61, 0x72, 0x63,
	0x68, 0x54, 0x65, 0x72, 0x6d, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a,
	0x04, 0x74, 0x65, 0x72, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x72,
	0x6d, 0x22, 0x1b, 0x0a, 0x19, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x65, 0x61, 0x72, 0x63,
	0x68, 0x54, 0x65, 0x72, 0x6d, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1a,
	0x0a, 0x18, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x54, 0x65, 0x72, 0x6d, 0x4d, 0x65, 0x74, 0x72,
	0x69, 0x63, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x49, 0x0a, 0x11, 0x53, 0x65,
	0x61, 0x72, 0x63, 0x68, 0x54, 0x65, 0x72, 0x6d, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x12,
	0x12, 0x0a, 0x04, 0x74, 0x65, 0x72, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74,
	0x65, 0x72, 0x6d, 0x12, 0x20, 0x0a, 0x0b, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x43, 0x6f, 0x75,
	0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68,
	0x43, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x5d, 0x0a, 0x19, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x54,
	0x65, 0x72, 0x6d, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x40, 0x0a, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x74, 0x65, 0x72, 0x6d,
	0x73, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68,
	0x54, 0x65, 0x72, 0x6d, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x52, 0x07, 0x72, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x73, 0x32, 0xcf, 0x02, 0x0a, 0x06, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x12,
	0x60, 0x0a, 0x0b, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x54, 0x65, 0x72, 0x6d, 0x73, 0x12, 0x26,
	0x2e, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x74, 0x65, 0x72, 0x6d, 0x73, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x54, 0x65, 0x72, 0x6d, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x74,
	0x65, 0x72, 0x6d, 0x73, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x65, 0x61,
	0x72, 0x63, 0x68, 0x54, 0x65, 0x72, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x6e, 0x0a, 0x0b, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x65, 0x72, 0x6d, 0x73,
	0x12, 0x2d, 0x2e, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x74, 0x65, 0x72, 0x6d, 0x73, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x65, 0x61,
	0x72, 0x63, 0x68, 0x54, 0x65, 0x72, 0x6d, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x2e, 0x2e, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x74, 0x65, 0x72, 0x6d, 0x73, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x65, 0x61, 0x72,
	0x63, 0x68, 0x54, 0x65, 0x72, 0x6d, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x73, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x4d, 0x65,
	0x74, 0x72, 0x69, 0x63, 0x73, 0x12, 0x2d, 0x2e, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x74, 0x65,
	0x72, 0x6d, 0x73, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x65, 0x61, 0x72,
	0x63, 0x68, 0x54, 0x65, 0x72, 0x6d, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x2e, 0x2e, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x74, 0x65, 0x72,
	0x6d, 0x73, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63,
	0x68, 0x54, 0x65, 0x72, 0x6d, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x16, 0x5a, 0x14, 0x2f, 0x73, 0x65, 0x61, 0x72, 0x63,
	0x68, 0x74, 0x65, 0x72, 0x6d, 0x73, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_searchterms_proto_rawDescOnce sync.Once
	file_searchterms_proto_rawDescData = file_searchterms_proto_rawDesc
)

func file_searchterms_proto_rawDescGZIP() []byte {
	file_searchterms_proto_rawDescOnce.Do(func() {
		file_searchterms_proto_rawDescData = protoimpl.X.CompressGZIP(file_searchterms_proto_rawDescData)
	})
	return file_searchterms_proto_rawDescData
}

var file_searchterms_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_searchterms_proto_goTypes = []interface{}{
	(*SearchTermRequest)(nil),         // 0: searchtermsprotobuf.SearchTermRequest
	(*SearchTermResponse)(nil),        // 1: searchtermsprotobuf.SearchTermResponse
	(*UpdateSearchTermsRequest)(nil),  // 2: searchtermsprotobuf.UpdateSearchTermsRequest
	(*UpdateSearchTermsResponse)(nil), // 3: searchtermsprotobuf.UpdateSearchTermsResponse
	(*SearchTermMetricsRequest)(nil),  // 4: searchtermsprotobuf.SearchTermMetricsRequest
	(*SearchTermMetrics)(nil),         // 5: searchtermsprotobuf.SearchTermMetrics
	(*SearchTermMetricsResponse)(nil), // 6: searchtermsprotobuf.SearchTermMetricsResponse
}
var file_searchterms_proto_depIdxs = []int32{
	5, // 0: searchtermsprotobuf.SearchTermMetricsResponse.results:type_name -> searchtermsprotobuf.SearchTermMetrics
	0, // 1: searchtermsprotobuf.Search.SearchTerms:input_type -> searchtermsprotobuf.SearchTermRequest
	2, // 2: searchtermsprotobuf.Search.UpdateTerms:input_type -> searchtermsprotobuf.UpdateSearchTermsRequest
	4, // 3: searchtermsprotobuf.Search.GetSearchMetrics:input_type -> searchtermsprotobuf.SearchTermMetricsRequest
	1, // 4: searchtermsprotobuf.Search.SearchTerms:output_type -> searchtermsprotobuf.SearchTermResponse
	3, // 5: searchtermsprotobuf.Search.UpdateTerms:output_type -> searchtermsprotobuf.UpdateSearchTermsResponse
	6, // 6: searchtermsprotobuf.Search.GetSearchMetrics:output_type -> searchtermsprotobuf.SearchTermMetricsResponse
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_searchterms_proto_init() }
func file_searchterms_proto_init() {
	if File_searchterms_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_searchterms_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchTermRequest); i {
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
		file_searchterms_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchTermResponse); i {
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
		file_searchterms_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateSearchTermsRequest); i {
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
		file_searchterms_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateSearchTermsResponse); i {
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
		file_searchterms_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchTermMetricsRequest); i {
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
		file_searchterms_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchTermMetrics); i {
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
		file_searchterms_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchTermMetricsResponse); i {
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
			RawDescriptor: file_searchterms_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_searchterms_proto_goTypes,
		DependencyIndexes: file_searchterms_proto_depIdxs,
		MessageInfos:      file_searchterms_proto_msgTypes,
	}.Build()
	File_searchterms_proto = out.File
	file_searchterms_proto_rawDesc = nil
	file_searchterms_proto_goTypes = nil
	file_searchterms_proto_depIdxs = nil
}