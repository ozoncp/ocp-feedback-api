// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.1
// source: proposal-messages.proto

package ocp_feedback_api

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
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

// Proposal messages
type Proposal struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProposalId uint64 `protobuf:"varint,1,opt,name=proposal_id,json=proposalId,proto3" json:"proposal_id,omitempty"`
	UserId     uint64 `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	LessonId   uint64 `protobuf:"varint,3,opt,name=lesson_id,json=lessonId,proto3" json:"lesson_id,omitempty"`
	DocumentId uint64 `protobuf:"varint,4,opt,name=document_id,json=documentId,proto3" json:"document_id,omitempty"`
}

func (x *Proposal) Reset() {
	*x = Proposal{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proposal_messages_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Proposal) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Proposal) ProtoMessage() {}

func (x *Proposal) ProtoReflect() protoreflect.Message {
	mi := &file_proposal_messages_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Proposal.ProtoReflect.Descriptor instead.
func (*Proposal) Descriptor() ([]byte, []int) {
	return file_proposal_messages_proto_rawDescGZIP(), []int{0}
}

func (x *Proposal) GetProposalId() uint64 {
	if x != nil {
		return x.ProposalId
	}
	return 0
}

func (x *Proposal) GetUserId() uint64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *Proposal) GetLessonId() uint64 {
	if x != nil {
		return x.LessonId
	}
	return 0
}

func (x *Proposal) GetDocumentId() uint64 {
	if x != nil {
		return x.DocumentId
	}
	return 0
}

type CreateProposalV1Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId     uint64 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	LessonId   uint64 `protobuf:"varint,2,opt,name=lesson_id,json=lessonId,proto3" json:"lesson_id,omitempty"`
	DocumentId uint64 `protobuf:"varint,3,opt,name=document_id,json=documentId,proto3" json:"document_id,omitempty"`
}

func (x *CreateProposalV1Request) Reset() {
	*x = CreateProposalV1Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proposal_messages_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateProposalV1Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateProposalV1Request) ProtoMessage() {}

func (x *CreateProposalV1Request) ProtoReflect() protoreflect.Message {
	mi := &file_proposal_messages_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateProposalV1Request.ProtoReflect.Descriptor instead.
func (*CreateProposalV1Request) Descriptor() ([]byte, []int) {
	return file_proposal_messages_proto_rawDescGZIP(), []int{1}
}

func (x *CreateProposalV1Request) GetUserId() uint64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *CreateProposalV1Request) GetLessonId() uint64 {
	if x != nil {
		return x.LessonId
	}
	return 0
}

func (x *CreateProposalV1Request) GetDocumentId() uint64 {
	if x != nil {
		return x.DocumentId
	}
	return 0
}

type CreateProposalV1Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProposalId uint64 `protobuf:"varint,1,opt,name=proposal_id,json=proposalId,proto3" json:"proposal_id,omitempty"`
}

func (x *CreateProposalV1Response) Reset() {
	*x = CreateProposalV1Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proposal_messages_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateProposalV1Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateProposalV1Response) ProtoMessage() {}

func (x *CreateProposalV1Response) ProtoReflect() protoreflect.Message {
	mi := &file_proposal_messages_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateProposalV1Response.ProtoReflect.Descriptor instead.
func (*CreateProposalV1Response) Descriptor() ([]byte, []int) {
	return file_proposal_messages_proto_rawDescGZIP(), []int{2}
}

func (x *CreateProposalV1Response) GetProposalId() uint64 {
	if x != nil {
		return x.ProposalId
	}
	return 0
}

type RemoveProposalV1Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProposalId uint64 `protobuf:"varint,1,opt,name=proposal_id,json=proposalId,proto3" json:"proposal_id,omitempty"`
}

func (x *RemoveProposalV1Request) Reset() {
	*x = RemoveProposalV1Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proposal_messages_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveProposalV1Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveProposalV1Request) ProtoMessage() {}

func (x *RemoveProposalV1Request) ProtoReflect() protoreflect.Message {
	mi := &file_proposal_messages_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveProposalV1Request.ProtoReflect.Descriptor instead.
func (*RemoveProposalV1Request) Descriptor() ([]byte, []int) {
	return file_proposal_messages_proto_rawDescGZIP(), []int{3}
}

func (x *RemoveProposalV1Request) GetProposalId() uint64 {
	if x != nil {
		return x.ProposalId
	}
	return 0
}

type RemoveProposalV1Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *RemoveProposalV1Response) Reset() {
	*x = RemoveProposalV1Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proposal_messages_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveProposalV1Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveProposalV1Response) ProtoMessage() {}

func (x *RemoveProposalV1Response) ProtoReflect() protoreflect.Message {
	mi := &file_proposal_messages_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveProposalV1Response.ProtoReflect.Descriptor instead.
func (*RemoveProposalV1Response) Descriptor() ([]byte, []int) {
	return file_proposal_messages_proto_rawDescGZIP(), []int{4}
}

type DescribeProposalV1Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProposalId uint64 `protobuf:"varint,1,opt,name=proposal_id,json=proposalId,proto3" json:"proposal_id,omitempty"`
}

func (x *DescribeProposalV1Request) Reset() {
	*x = DescribeProposalV1Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proposal_messages_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DescribeProposalV1Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DescribeProposalV1Request) ProtoMessage() {}

func (x *DescribeProposalV1Request) ProtoReflect() protoreflect.Message {
	mi := &file_proposal_messages_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DescribeProposalV1Request.ProtoReflect.Descriptor instead.
func (*DescribeProposalV1Request) Descriptor() ([]byte, []int) {
	return file_proposal_messages_proto_rawDescGZIP(), []int{5}
}

func (x *DescribeProposalV1Request) GetProposalId() uint64 {
	if x != nil {
		return x.ProposalId
	}
	return 0
}

type DescribeProposalV1Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Proposal *Proposal `protobuf:"bytes,1,opt,name=proposal,proto3" json:"proposal,omitempty"`
}

func (x *DescribeProposalV1Response) Reset() {
	*x = DescribeProposalV1Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proposal_messages_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DescribeProposalV1Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DescribeProposalV1Response) ProtoMessage() {}

func (x *DescribeProposalV1Response) ProtoReflect() protoreflect.Message {
	mi := &file_proposal_messages_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DescribeProposalV1Response.ProtoReflect.Descriptor instead.
func (*DescribeProposalV1Response) Descriptor() ([]byte, []int) {
	return file_proposal_messages_proto_rawDescGZIP(), []int{6}
}

func (x *DescribeProposalV1Response) GetProposal() *Proposal {
	if x != nil {
		return x.Proposal
	}
	return nil
}

type ListProposalsV1Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Limit  uint64 `protobuf:"varint,1,opt,name=limit,proto3" json:"limit,omitempty"`
	Offset uint64 `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"`
}

func (x *ListProposalsV1Request) Reset() {
	*x = ListProposalsV1Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proposal_messages_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListProposalsV1Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListProposalsV1Request) ProtoMessage() {}

func (x *ListProposalsV1Request) ProtoReflect() protoreflect.Message {
	mi := &file_proposal_messages_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListProposalsV1Request.ProtoReflect.Descriptor instead.
func (*ListProposalsV1Request) Descriptor() ([]byte, []int) {
	return file_proposal_messages_proto_rawDescGZIP(), []int{7}
}

func (x *ListProposalsV1Request) GetLimit() uint64 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *ListProposalsV1Request) GetOffset() uint64 {
	if x != nil {
		return x.Offset
	}
	return 0
}

type ListProposalsV1Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Proposals []*Proposal `protobuf:"bytes,1,rep,name=proposals,proto3" json:"proposals,omitempty"`
}

func (x *ListProposalsV1Response) Reset() {
	*x = ListProposalsV1Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proposal_messages_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListProposalsV1Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListProposalsV1Response) ProtoMessage() {}

func (x *ListProposalsV1Response) ProtoReflect() protoreflect.Message {
	mi := &file_proposal_messages_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListProposalsV1Response.ProtoReflect.Descriptor instead.
func (*ListProposalsV1Response) Descriptor() ([]byte, []int) {
	return file_proposal_messages_proto_rawDescGZIP(), []int{8}
}

func (x *ListProposalsV1Response) GetProposals() []*Proposal {
	if x != nil {
		return x.Proposals
	}
	return nil
}

var File_proposal_messages_proto protoreflect.FileDescriptor

var file_proposal_messages_proto_rawDesc = []byte{
	0x0a, 0x17, 0x70, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x61, 0x6c, 0x2d, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x6f, 0x63, 0x70, 0x2e, 0x66,
	0x65, 0x65, 0x64, 0x62, 0x61, 0x63, 0x6b, 0x2e, 0x61, 0x70, 0x69, 0x1a, 0x41, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f,
	0x78, 0x79, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x76, 0x61,
	0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f,
	0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x82,
	0x01, 0x0a, 0x08, 0x50, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x61, 0x6c, 0x12, 0x1f, 0x0a, 0x0b, 0x70,
	0x72, 0x6f, 0x70, 0x6f, 0x73, 0x61, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x0a, 0x70, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x61, 0x6c, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07,
	0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x6c, 0x65, 0x73, 0x73, 0x6f, 0x6e, 0x5f,
	0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x6c, 0x65, 0x73, 0x73, 0x6f, 0x6e,
	0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x69,
	0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e,
	0x74, 0x49, 0x64, 0x22, 0x8b, 0x01, 0x0a, 0x17, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x72,
	0x6f, 0x70, 0x6f, 0x73, 0x61, 0x6c, 0x56, 0x31, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x20, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04,
	0x42, 0x07, 0xfa, 0x42, 0x04, 0x32, 0x02, 0x20, 0x00, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49,
	0x64, 0x12, 0x24, 0x0a, 0x09, 0x6c, 0x65, 0x73, 0x73, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x04, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x32, 0x02, 0x20, 0x00, 0x52, 0x08, 0x6c,
	0x65, 0x73, 0x73, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x28, 0x0a, 0x0b, 0x64, 0x6f, 0x63, 0x75, 0x6d,
	0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x42, 0x07, 0xfa, 0x42,
	0x04, 0x32, 0x02, 0x20, 0x00, 0x52, 0x0a, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x49,
	0x64, 0x22, 0x3b, 0x0a, 0x18, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x70, 0x6f,
	0x73, 0x61, 0x6c, 0x56, 0x31, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a,
	0x0b, 0x70, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x61, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x0a, 0x70, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x61, 0x6c, 0x49, 0x64, 0x22, 0x43,
	0x0a, 0x17, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x50, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x61, 0x6c,
	0x56, 0x31, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x28, 0x0a, 0x0b, 0x70, 0x72, 0x6f,
	0x70, 0x6f, 0x73, 0x61, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x42, 0x07,
	0xfa, 0x42, 0x04, 0x32, 0x02, 0x20, 0x00, 0x52, 0x0a, 0x70, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x61,
	0x6c, 0x49, 0x64, 0x22, 0x1a, 0x0a, 0x18, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x50, 0x72, 0x6f,
	0x70, 0x6f, 0x73, 0x61, 0x6c, 0x56, 0x31, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x45, 0x0a, 0x19, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x50, 0x72, 0x6f, 0x70, 0x6f,
	0x73, 0x61, 0x6c, 0x56, 0x31, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x28, 0x0a, 0x0b,
	0x70, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x61, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x04, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x32, 0x02, 0x20, 0x00, 0x52, 0x0a, 0x70, 0x72, 0x6f, 0x70,
	0x6f, 0x73, 0x61, 0x6c, 0x49, 0x64, 0x22, 0x54, 0x0a, 0x1a, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x62, 0x65, 0x50, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x61, 0x6c, 0x56, 0x31, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x36, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x61, 0x6c,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x6f, 0x63, 0x70, 0x2e, 0x66, 0x65, 0x65,
	0x64, 0x62, 0x61, 0x63, 0x6b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x50, 0x72, 0x6f, 0x70, 0x6f, 0x73,
	0x61, 0x6c, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x61, 0x6c, 0x22, 0x4f, 0x0a, 0x16,
	0x4c, 0x69, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x61, 0x6c, 0x73, 0x56, 0x31, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x04, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x32, 0x02, 0x20, 0x00, 0x52, 0x05,
	0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x22, 0x53, 0x0a,
	0x17, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x61, 0x6c, 0x73, 0x56, 0x31,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x38, 0x0a, 0x09, 0x70, 0x72, 0x6f, 0x70,
	0x6f, 0x73, 0x61, 0x6c, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x6f, 0x63,
	0x70, 0x2e, 0x66, 0x65, 0x65, 0x64, 0x62, 0x61, 0x63, 0x6b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x50,
	0x72, 0x6f, 0x70, 0x6f, 0x73, 0x61, 0x6c, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x61,
	0x6c, 0x73, 0x42, 0x4a, 0x5a, 0x48, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x6f, 0x7a, 0x6f, 0x6e, 0x63, 0x70, 0x2f, 0x6f, 0x63, 0x70, 0x2d, 0x66, 0x65, 0x65, 0x64,
	0x62, 0x61, 0x63, 0x6b, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x6f, 0x63, 0x70,
	0x2d, 0x66, 0x65, 0x65, 0x64, 0x62, 0x61, 0x63, 0x6b, 0x2d, 0x61, 0x70, 0x69, 0x3b, 0x6f, 0x63,
	0x70, 0x5f, 0x66, 0x65, 0x65, 0x64, 0x62, 0x61, 0x63, 0x6b, 0x5f, 0x61, 0x70, 0x69, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proposal_messages_proto_rawDescOnce sync.Once
	file_proposal_messages_proto_rawDescData = file_proposal_messages_proto_rawDesc
)

func file_proposal_messages_proto_rawDescGZIP() []byte {
	file_proposal_messages_proto_rawDescOnce.Do(func() {
		file_proposal_messages_proto_rawDescData = protoimpl.X.CompressGZIP(file_proposal_messages_proto_rawDescData)
	})
	return file_proposal_messages_proto_rawDescData
}

var file_proposal_messages_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_proposal_messages_proto_goTypes = []interface{}{
	(*Proposal)(nil),                   // 0: ocp.feedback.api.Proposal
	(*CreateProposalV1Request)(nil),    // 1: ocp.feedback.api.CreateProposalV1Request
	(*CreateProposalV1Response)(nil),   // 2: ocp.feedback.api.CreateProposalV1Response
	(*RemoveProposalV1Request)(nil),    // 3: ocp.feedback.api.RemoveProposalV1Request
	(*RemoveProposalV1Response)(nil),   // 4: ocp.feedback.api.RemoveProposalV1Response
	(*DescribeProposalV1Request)(nil),  // 5: ocp.feedback.api.DescribeProposalV1Request
	(*DescribeProposalV1Response)(nil), // 6: ocp.feedback.api.DescribeProposalV1Response
	(*ListProposalsV1Request)(nil),     // 7: ocp.feedback.api.ListProposalsV1Request
	(*ListProposalsV1Response)(nil),    // 8: ocp.feedback.api.ListProposalsV1Response
}
var file_proposal_messages_proto_depIdxs = []int32{
	0, // 0: ocp.feedback.api.DescribeProposalV1Response.proposal:type_name -> ocp.feedback.api.Proposal
	0, // 1: ocp.feedback.api.ListProposalsV1Response.proposals:type_name -> ocp.feedback.api.Proposal
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_proposal_messages_proto_init() }
func file_proposal_messages_proto_init() {
	if File_proposal_messages_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proposal_messages_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Proposal); i {
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
		file_proposal_messages_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateProposalV1Request); i {
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
		file_proposal_messages_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateProposalV1Response); i {
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
		file_proposal_messages_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoveProposalV1Request); i {
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
		file_proposal_messages_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoveProposalV1Response); i {
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
		file_proposal_messages_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DescribeProposalV1Request); i {
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
		file_proposal_messages_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DescribeProposalV1Response); i {
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
		file_proposal_messages_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListProposalsV1Request); i {
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
		file_proposal_messages_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListProposalsV1Response); i {
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
			RawDescriptor: file_proposal_messages_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proposal_messages_proto_goTypes,
		DependencyIndexes: file_proposal_messages_proto_depIdxs,
		MessageInfos:      file_proposal_messages_proto_msgTypes,
	}.Build()
	File_proposal_messages_proto = out.File
	file_proposal_messages_proto_rawDesc = nil
	file_proposal_messages_proto_goTypes = nil
	file_proposal_messages_proto_depIdxs = nil
}