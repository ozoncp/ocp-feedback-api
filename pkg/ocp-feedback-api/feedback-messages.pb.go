// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.1
// source: feedback-messages.proto

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

type Feedback struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId      uint64 `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	ClassroomId uint64 `protobuf:"varint,3,opt,name=classroom_id,json=classroomId,proto3" json:"classroom_id,omitempty"`
	Comment     string `protobuf:"bytes,4,opt,name=comment,proto3" json:"comment,omitempty"`
}

func (x *Feedback) Reset() {
	*x = Feedback{}
	if protoimpl.UnsafeEnabled {
		mi := &file_feedback_messages_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Feedback) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Feedback) ProtoMessage() {}

func (x *Feedback) ProtoReflect() protoreflect.Message {
	mi := &file_feedback_messages_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Feedback.ProtoReflect.Descriptor instead.
func (*Feedback) Descriptor() ([]byte, []int) {
	return file_feedback_messages_proto_rawDescGZIP(), []int{0}
}

func (x *Feedback) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Feedback) GetUserId() uint64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *Feedback) GetClassroomId() uint64 {
	if x != nil {
		return x.ClassroomId
	}
	return 0
}

func (x *Feedback) GetComment() string {
	if x != nil {
		return x.Comment
	}
	return ""
}

type CreateFeedbackV1Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Feedback *Feedback `protobuf:"bytes,1,opt,name=feedback,proto3" json:"feedback,omitempty"`
}

func (x *CreateFeedbackV1Request) Reset() {
	*x = CreateFeedbackV1Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_feedback_messages_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateFeedbackV1Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateFeedbackV1Request) ProtoMessage() {}

func (x *CreateFeedbackV1Request) ProtoReflect() protoreflect.Message {
	mi := &file_feedback_messages_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateFeedbackV1Request.ProtoReflect.Descriptor instead.
func (*CreateFeedbackV1Request) Descriptor() ([]byte, []int) {
	return file_feedback_messages_proto_rawDescGZIP(), []int{1}
}

func (x *CreateFeedbackV1Request) GetFeedback() *Feedback {
	if x != nil {
		return x.Feedback
	}
	return nil
}

type CreateFeedbackV1Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Feedback uint64 `protobuf:"varint,1,opt,name=feedback,proto3" json:"feedback,omitempty"`
}

func (x *CreateFeedbackV1Response) Reset() {
	*x = CreateFeedbackV1Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_feedback_messages_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateFeedbackV1Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateFeedbackV1Response) ProtoMessage() {}

func (x *CreateFeedbackV1Response) ProtoReflect() protoreflect.Message {
	mi := &file_feedback_messages_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateFeedbackV1Response.ProtoReflect.Descriptor instead.
func (*CreateFeedbackV1Response) Descriptor() ([]byte, []int) {
	return file_feedback_messages_proto_rawDescGZIP(), []int{2}
}

func (x *CreateFeedbackV1Response) GetFeedback() uint64 {
	if x != nil {
		return x.Feedback
	}
	return 0
}

type CreateMultiFeedbackV1Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Feedbacks []*Feedback `protobuf:"bytes,1,rep,name=feedbacks,proto3" json:"feedbacks,omitempty"`
}

func (x *CreateMultiFeedbackV1Request) Reset() {
	*x = CreateMultiFeedbackV1Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_feedback_messages_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateMultiFeedbackV1Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateMultiFeedbackV1Request) ProtoMessage() {}

func (x *CreateMultiFeedbackV1Request) ProtoReflect() protoreflect.Message {
	mi := &file_feedback_messages_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateMultiFeedbackV1Request.ProtoReflect.Descriptor instead.
func (*CreateMultiFeedbackV1Request) Descriptor() ([]byte, []int) {
	return file_feedback_messages_proto_rawDescGZIP(), []int{3}
}

func (x *CreateMultiFeedbackV1Request) GetFeedbacks() []*Feedback {
	if x != nil {
		return x.Feedbacks
	}
	return nil
}

type CreateMultiFeedbackV1Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Feedbacks []uint64 `protobuf:"varint,1,rep,packed,name=feedbacks,proto3" json:"feedbacks,omitempty"`
}

func (x *CreateMultiFeedbackV1Response) Reset() {
	*x = CreateMultiFeedbackV1Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_feedback_messages_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateMultiFeedbackV1Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateMultiFeedbackV1Response) ProtoMessage() {}

func (x *CreateMultiFeedbackV1Response) ProtoReflect() protoreflect.Message {
	mi := &file_feedback_messages_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateMultiFeedbackV1Response.ProtoReflect.Descriptor instead.
func (*CreateMultiFeedbackV1Response) Descriptor() ([]byte, []int) {
	return file_feedback_messages_proto_rawDescGZIP(), []int{4}
}

func (x *CreateMultiFeedbackV1Response) GetFeedbacks() []uint64 {
	if x != nil {
		return x.Feedbacks
	}
	return nil
}

type RemoveFeedbackV1Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Feedback uint64 `protobuf:"varint,1,opt,name=feedback,proto3" json:"feedback,omitempty"`
}

func (x *RemoveFeedbackV1Request) Reset() {
	*x = RemoveFeedbackV1Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_feedback_messages_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveFeedbackV1Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveFeedbackV1Request) ProtoMessage() {}

func (x *RemoveFeedbackV1Request) ProtoReflect() protoreflect.Message {
	mi := &file_feedback_messages_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveFeedbackV1Request.ProtoReflect.Descriptor instead.
func (*RemoveFeedbackV1Request) Descriptor() ([]byte, []int) {
	return file_feedback_messages_proto_rawDescGZIP(), []int{5}
}

func (x *RemoveFeedbackV1Request) GetFeedback() uint64 {
	if x != nil {
		return x.Feedback
	}
	return 0
}

type RemoveFeedbackV1Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *RemoveFeedbackV1Response) Reset() {
	*x = RemoveFeedbackV1Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_feedback_messages_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveFeedbackV1Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveFeedbackV1Response) ProtoMessage() {}

func (x *RemoveFeedbackV1Response) ProtoReflect() protoreflect.Message {
	mi := &file_feedback_messages_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveFeedbackV1Response.ProtoReflect.Descriptor instead.
func (*RemoveFeedbackV1Response) Descriptor() ([]byte, []int) {
	return file_feedback_messages_proto_rawDescGZIP(), []int{6}
}

type DescribeFeedbackV1Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Feedback uint64 `protobuf:"varint,1,opt,name=feedback,proto3" json:"feedback,omitempty"`
}

func (x *DescribeFeedbackV1Request) Reset() {
	*x = DescribeFeedbackV1Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_feedback_messages_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DescribeFeedbackV1Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DescribeFeedbackV1Request) ProtoMessage() {}

func (x *DescribeFeedbackV1Request) ProtoReflect() protoreflect.Message {
	mi := &file_feedback_messages_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DescribeFeedbackV1Request.ProtoReflect.Descriptor instead.
func (*DescribeFeedbackV1Request) Descriptor() ([]byte, []int) {
	return file_feedback_messages_proto_rawDescGZIP(), []int{7}
}

func (x *DescribeFeedbackV1Request) GetFeedback() uint64 {
	if x != nil {
		return x.Feedback
	}
	return 0
}

type DescribeFeedbackV1Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Feedback *Feedback `protobuf:"bytes,1,opt,name=feedback,proto3" json:"feedback,omitempty"`
}

func (x *DescribeFeedbackV1Response) Reset() {
	*x = DescribeFeedbackV1Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_feedback_messages_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DescribeFeedbackV1Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DescribeFeedbackV1Response) ProtoMessage() {}

func (x *DescribeFeedbackV1Response) ProtoReflect() protoreflect.Message {
	mi := &file_feedback_messages_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DescribeFeedbackV1Response.ProtoReflect.Descriptor instead.
func (*DescribeFeedbackV1Response) Descriptor() ([]byte, []int) {
	return file_feedback_messages_proto_rawDescGZIP(), []int{8}
}

func (x *DescribeFeedbackV1Response) GetFeedback() *Feedback {
	if x != nil {
		return x.Feedback
	}
	return nil
}

type ListFeedbacksV1Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Limit  uint64 `protobuf:"varint,1,opt,name=limit,proto3" json:"limit,omitempty"`
	Offset uint64 `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"`
}

func (x *ListFeedbacksV1Request) Reset() {
	*x = ListFeedbacksV1Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_feedback_messages_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListFeedbacksV1Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListFeedbacksV1Request) ProtoMessage() {}

func (x *ListFeedbacksV1Request) ProtoReflect() protoreflect.Message {
	mi := &file_feedback_messages_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListFeedbacksV1Request.ProtoReflect.Descriptor instead.
func (*ListFeedbacksV1Request) Descriptor() ([]byte, []int) {
	return file_feedback_messages_proto_rawDescGZIP(), []int{9}
}

func (x *ListFeedbacksV1Request) GetLimit() uint64 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *ListFeedbacksV1Request) GetOffset() uint64 {
	if x != nil {
		return x.Offset
	}
	return 0
}

type ListFeedbacksV1Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Feedbacks []*Feedback `protobuf:"bytes,1,rep,name=feedbacks,proto3" json:"feedbacks,omitempty"`
}

func (x *ListFeedbacksV1Response) Reset() {
	*x = ListFeedbacksV1Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_feedback_messages_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListFeedbacksV1Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListFeedbacksV1Response) ProtoMessage() {}

func (x *ListFeedbacksV1Response) ProtoReflect() protoreflect.Message {
	mi := &file_feedback_messages_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListFeedbacksV1Response.ProtoReflect.Descriptor instead.
func (*ListFeedbacksV1Response) Descriptor() ([]byte, []int) {
	return file_feedback_messages_proto_rawDescGZIP(), []int{10}
}

func (x *ListFeedbacksV1Response) GetFeedbacks() []*Feedback {
	if x != nil {
		return x.Feedbacks
	}
	return nil
}

type UpdateFeedbackV1Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Feedback *Feedback `protobuf:"bytes,1,opt,name=feedback,proto3" json:"feedback,omitempty"`
}

func (x *UpdateFeedbackV1Request) Reset() {
	*x = UpdateFeedbackV1Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_feedback_messages_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateFeedbackV1Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateFeedbackV1Request) ProtoMessage() {}

func (x *UpdateFeedbackV1Request) ProtoReflect() protoreflect.Message {
	mi := &file_feedback_messages_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateFeedbackV1Request.ProtoReflect.Descriptor instead.
func (*UpdateFeedbackV1Request) Descriptor() ([]byte, []int) {
	return file_feedback_messages_proto_rawDescGZIP(), []int{11}
}

func (x *UpdateFeedbackV1Request) GetFeedback() *Feedback {
	if x != nil {
		return x.Feedback
	}
	return nil
}

type UpdateFeedbackV1Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateFeedbackV1Response) Reset() {
	*x = UpdateFeedbackV1Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_feedback_messages_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateFeedbackV1Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateFeedbackV1Response) ProtoMessage() {}

func (x *UpdateFeedbackV1Response) ProtoReflect() protoreflect.Message {
	mi := &file_feedback_messages_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateFeedbackV1Response.ProtoReflect.Descriptor instead.
func (*UpdateFeedbackV1Response) Descriptor() ([]byte, []int) {
	return file_feedback_messages_proto_rawDescGZIP(), []int{12}
}

var File_feedback_messages_proto protoreflect.FileDescriptor

var file_feedback_messages_proto_rawDesc = []byte{
	0x0a, 0x17, 0x66, 0x65, 0x65, 0x64, 0x62, 0x61, 0x63, 0x6b, 0x2d, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x6f, 0x63, 0x70, 0x2e, 0x66,
	0x65, 0x65, 0x64, 0x62, 0x61, 0x63, 0x6b, 0x2e, 0x61, 0x70, 0x69, 0x1a, 0x41, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f,
	0x78, 0x79, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x76, 0x61,
	0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f,
	0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8b,
	0x01, 0x0a, 0x08, 0x46, 0x65, 0x65, 0x64, 0x62, 0x61, 0x63, 0x6b, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x20, 0x0a, 0x07, 0x75,
	0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x42, 0x07, 0xfa, 0x42,
	0x04, 0x32, 0x02, 0x20, 0x00, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x2a, 0x0a,
	0x0c, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x72, 0x6f, 0x6f, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x04, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x32, 0x02, 0x20, 0x00, 0x52, 0x0b, 0x63, 0x6c,
	0x61, 0x73, 0x73, 0x72, 0x6f, 0x6f, 0x6d, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x07, 0x63, 0x6f, 0x6d,
	0x6d, 0x65, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72,
	0x02, 0x10, 0x01, 0x52, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x22, 0x5b, 0x0a, 0x17,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x46, 0x65, 0x65, 0x64, 0x62, 0x61, 0x63, 0x6b, 0x56, 0x31,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x40, 0x0a, 0x08, 0x66, 0x65, 0x65, 0x64, 0x62,
	0x61, 0x63, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x6f, 0x63, 0x70, 0x2e,
	0x66, 0x65, 0x65, 0x64, 0x62, 0x61, 0x63, 0x6b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x46, 0x65, 0x65,
	0x64, 0x62, 0x61, 0x63, 0x6b, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x8a, 0x01, 0x02, 0x10, 0x01, 0x52,
	0x08, 0x66, 0x65, 0x65, 0x64, 0x62, 0x61, 0x63, 0x6b, 0x22, 0x3f, 0x0a, 0x18, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x46, 0x65, 0x65, 0x64, 0x62, 0x61, 0x63, 0x6b, 0x56, 0x31, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x23, 0x0a, 0x08, 0x66, 0x65, 0x65, 0x64, 0x62, 0x61, 0x63,
	0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x32, 0x02, 0x20, 0x00,
	0x52, 0x08, 0x66, 0x65, 0x65, 0x64, 0x62, 0x61, 0x63, 0x6b, 0x22, 0x62, 0x0a, 0x1c, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x46, 0x65, 0x65, 0x64, 0x62, 0x61, 0x63,
	0x6b, 0x56, 0x31, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x42, 0x0a, 0x09, 0x66, 0x65,
	0x65, 0x64, 0x62, 0x61, 0x63, 0x6b, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x6f, 0x63, 0x70, 0x2e, 0x66, 0x65, 0x65, 0x64, 0x62, 0x61, 0x63, 0x6b, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x46, 0x65, 0x65, 0x64, 0x62, 0x61, 0x63, 0x6b, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x92, 0x01,
	0x02, 0x08, 0x01, 0x52, 0x09, 0x66, 0x65, 0x65, 0x64, 0x62, 0x61, 0x63, 0x6b, 0x73, 0x22, 0x47,
	0x0a, 0x1d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x46, 0x65, 0x65,
	0x64, 0x62, 0x61, 0x63, 0x6b, 0x56, 0x31, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x26, 0x0a, 0x09, 0x66, 0x65, 0x65, 0x64, 0x62, 0x61, 0x63, 0x6b, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x04, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x92, 0x01, 0x02, 0x08, 0x01, 0x52, 0x09, 0x66, 0x65,
	0x65, 0x64, 0x62, 0x61, 0x63, 0x6b, 0x73, 0x22, 0x3e, 0x0a, 0x17, 0x52, 0x65, 0x6d, 0x6f, 0x76,
	0x65, 0x46, 0x65, 0x65, 0x64, 0x62, 0x61, 0x63, 0x6b, 0x56, 0x31, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x23, 0x0a, 0x08, 0x66, 0x65, 0x65, 0x64, 0x62, 0x61, 0x63, 0x6b, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x04, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x32, 0x02, 0x20, 0x00, 0x52, 0x08, 0x66,
	0x65, 0x65, 0x64, 0x62, 0x61, 0x63, 0x6b, 0x22, 0x1a, 0x0a, 0x18, 0x52, 0x65, 0x6d, 0x6f, 0x76,
	0x65, 0x46, 0x65, 0x65, 0x64, 0x62, 0x61, 0x63, 0x6b, 0x56, 0x31, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x40, 0x0a, 0x19, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x46,
	0x65, 0x65, 0x64, 0x62, 0x61, 0x63, 0x6b, 0x56, 0x31, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x23, 0x0a, 0x08, 0x66, 0x65, 0x65, 0x64, 0x62, 0x61, 0x63, 0x6b, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x04, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x32, 0x02, 0x20, 0x00, 0x52, 0x08, 0x66, 0x65, 0x65,
	0x64, 0x62, 0x61, 0x63, 0x6b, 0x22, 0x5e, 0x0a, 0x1a, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62,
	0x65, 0x46, 0x65, 0x65, 0x64, 0x62, 0x61, 0x63, 0x6b, 0x56, 0x31, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x40, 0x0a, 0x08, 0x66, 0x65, 0x65, 0x64, 0x62, 0x61, 0x63, 0x6b, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x6f, 0x63, 0x70, 0x2e, 0x66, 0x65, 0x65, 0x64,
	0x62, 0x61, 0x63, 0x6b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x46, 0x65, 0x65, 0x64, 0x62, 0x61, 0x63,
	0x6b, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x8a, 0x01, 0x02, 0x10, 0x01, 0x52, 0x08, 0x66, 0x65, 0x65,
	0x64, 0x62, 0x61, 0x63, 0x6b, 0x22, 0x4f, 0x0a, 0x16, 0x4c, 0x69, 0x73, 0x74, 0x46, 0x65, 0x65,
	0x64, 0x62, 0x61, 0x63, 0x6b, 0x73, 0x56, 0x31, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x1d, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x42, 0x07,
	0xfa, 0x42, 0x04, 0x32, 0x02, 0x20, 0x00, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x16,
	0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06,
	0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x22, 0x53, 0x0a, 0x17, 0x4c, 0x69, 0x73, 0x74, 0x46, 0x65,
	0x65, 0x64, 0x62, 0x61, 0x63, 0x6b, 0x73, 0x56, 0x31, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x38, 0x0a, 0x09, 0x66, 0x65, 0x65, 0x64, 0x62, 0x61, 0x63, 0x6b, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x6f, 0x63, 0x70, 0x2e, 0x66, 0x65, 0x65, 0x64, 0x62,
	0x61, 0x63, 0x6b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x46, 0x65, 0x65, 0x64, 0x62, 0x61, 0x63, 0x6b,
	0x52, 0x09, 0x66, 0x65, 0x65, 0x64, 0x62, 0x61, 0x63, 0x6b, 0x73, 0x22, 0x5b, 0x0a, 0x17, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x46, 0x65, 0x65, 0x64, 0x62, 0x61, 0x63, 0x6b, 0x56, 0x31, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x40, 0x0a, 0x08, 0x66, 0x65, 0x65, 0x64, 0x62, 0x61,
	0x63, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x6f, 0x63, 0x70, 0x2e, 0x66,
	0x65, 0x65, 0x64, 0x62, 0x61, 0x63, 0x6b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x46, 0x65, 0x65, 0x64,
	0x62, 0x61, 0x63, 0x6b, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x8a, 0x01, 0x02, 0x10, 0x01, 0x52, 0x08,
	0x66, 0x65, 0x65, 0x64, 0x62, 0x61, 0x63, 0x6b, 0x22, 0x1a, 0x0a, 0x18, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x46, 0x65, 0x65, 0x64, 0x62, 0x61, 0x63, 0x6b, 0x56, 0x31, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x42, 0x4a, 0x5a, 0x48, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x6f, 0x7a, 0x6f, 0x6e, 0x63, 0x70, 0x2f, 0x6f, 0x63, 0x70, 0x2d, 0x66, 0x65,
	0x65, 0x64, 0x62, 0x61, 0x63, 0x6b, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x6f,
	0x63, 0x70, 0x2d, 0x66, 0x65, 0x65, 0x64, 0x62, 0x61, 0x63, 0x6b, 0x2d, 0x61, 0x70, 0x69, 0x3b,
	0x6f, 0x63, 0x70, 0x5f, 0x66, 0x65, 0x65, 0x64, 0x62, 0x61, 0x63, 0x6b, 0x5f, 0x61, 0x70, 0x69,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_feedback_messages_proto_rawDescOnce sync.Once
	file_feedback_messages_proto_rawDescData = file_feedback_messages_proto_rawDesc
)

func file_feedback_messages_proto_rawDescGZIP() []byte {
	file_feedback_messages_proto_rawDescOnce.Do(func() {
		file_feedback_messages_proto_rawDescData = protoimpl.X.CompressGZIP(file_feedback_messages_proto_rawDescData)
	})
	return file_feedback_messages_proto_rawDescData
}

var file_feedback_messages_proto_msgTypes = make([]protoimpl.MessageInfo, 13)
var file_feedback_messages_proto_goTypes = []interface{}{
	(*Feedback)(nil),                      // 0: ocp.feedback.api.Feedback
	(*CreateFeedbackV1Request)(nil),       // 1: ocp.feedback.api.CreateFeedbackV1Request
	(*CreateFeedbackV1Response)(nil),      // 2: ocp.feedback.api.CreateFeedbackV1Response
	(*CreateMultiFeedbackV1Request)(nil),  // 3: ocp.feedback.api.CreateMultiFeedbackV1Request
	(*CreateMultiFeedbackV1Response)(nil), // 4: ocp.feedback.api.CreateMultiFeedbackV1Response
	(*RemoveFeedbackV1Request)(nil),       // 5: ocp.feedback.api.RemoveFeedbackV1Request
	(*RemoveFeedbackV1Response)(nil),      // 6: ocp.feedback.api.RemoveFeedbackV1Response
	(*DescribeFeedbackV1Request)(nil),     // 7: ocp.feedback.api.DescribeFeedbackV1Request
	(*DescribeFeedbackV1Response)(nil),    // 8: ocp.feedback.api.DescribeFeedbackV1Response
	(*ListFeedbacksV1Request)(nil),        // 9: ocp.feedback.api.ListFeedbacksV1Request
	(*ListFeedbacksV1Response)(nil),       // 10: ocp.feedback.api.ListFeedbacksV1Response
	(*UpdateFeedbackV1Request)(nil),       // 11: ocp.feedback.api.UpdateFeedbackV1Request
	(*UpdateFeedbackV1Response)(nil),      // 12: ocp.feedback.api.UpdateFeedbackV1Response
}
var file_feedback_messages_proto_depIdxs = []int32{
	0, // 0: ocp.feedback.api.CreateFeedbackV1Request.feedback:type_name -> ocp.feedback.api.Feedback
	0, // 1: ocp.feedback.api.CreateMultiFeedbackV1Request.feedbacks:type_name -> ocp.feedback.api.Feedback
	0, // 2: ocp.feedback.api.DescribeFeedbackV1Response.feedback:type_name -> ocp.feedback.api.Feedback
	0, // 3: ocp.feedback.api.ListFeedbacksV1Response.feedbacks:type_name -> ocp.feedback.api.Feedback
	0, // 4: ocp.feedback.api.UpdateFeedbackV1Request.feedback:type_name -> ocp.feedback.api.Feedback
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_feedback_messages_proto_init() }
func file_feedback_messages_proto_init() {
	if File_feedback_messages_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_feedback_messages_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Feedback); i {
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
		file_feedback_messages_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateFeedbackV1Request); i {
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
		file_feedback_messages_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateFeedbackV1Response); i {
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
		file_feedback_messages_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateMultiFeedbackV1Request); i {
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
		file_feedback_messages_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateMultiFeedbackV1Response); i {
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
		file_feedback_messages_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoveFeedbackV1Request); i {
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
		file_feedback_messages_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoveFeedbackV1Response); i {
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
		file_feedback_messages_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DescribeFeedbackV1Request); i {
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
		file_feedback_messages_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DescribeFeedbackV1Response); i {
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
		file_feedback_messages_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListFeedbacksV1Request); i {
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
		file_feedback_messages_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListFeedbacksV1Response); i {
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
		file_feedback_messages_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateFeedbackV1Request); i {
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
		file_feedback_messages_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateFeedbackV1Response); i {
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
			RawDescriptor: file_feedback_messages_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   13,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_feedback_messages_proto_goTypes,
		DependencyIndexes: file_feedback_messages_proto_depIdxs,
		MessageInfos:      file_feedback_messages_proto_msgTypes,
	}.Build()
	File_feedback_messages_proto = out.File
	file_feedback_messages_proto_rawDesc = nil
	file_feedback_messages_proto_goTypes = nil
	file_feedback_messages_proto_depIdxs = nil
}
