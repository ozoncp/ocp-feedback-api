// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.1
// source: feedback-service.proto

package ocp_feedback_api

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_feedback_service_proto protoreflect.FileDescriptor

var file_feedback_service_proto_rawDesc = []byte{
	0x0a, 0x16, 0x66, 0x65, 0x65, 0x64, 0x62, 0x61, 0x63, 0x6b, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x6f, 0x63, 0x70, 0x2e, 0x66, 0x65,
	0x65, 0x64, 0x62, 0x61, 0x63, 0x6b, 0x2e, 0x61, 0x70, 0x69, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x66, 0x65, 0x65, 0x64, 0x62, 0x61,
	0x63, 0x6b, 0x2d, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x17, 0x70, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x61, 0x6c, 0x2d, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0x88, 0x0b, 0x0a, 0x0e, 0x4f,
	0x63, 0x70, 0x46, 0x65, 0x65, 0x64, 0x62, 0x61, 0x63, 0x6b, 0x41, 0x70, 0x69, 0x12, 0x80, 0x01,
	0x0a, 0x10, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x46, 0x65, 0x65, 0x64, 0x62, 0x61, 0x63, 0x6b,
	0x56, 0x31, 0x12, 0x29, 0x2e, 0x6f, 0x63, 0x70, 0x2e, 0x66, 0x65, 0x65, 0x64, 0x62, 0x61, 0x63,
	0x6b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x46, 0x65, 0x65, 0x64,
	0x62, 0x61, 0x63, 0x6b, 0x56, 0x31, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2a, 0x2e,
	0x6f, 0x63, 0x70, 0x2e, 0x66, 0x65, 0x65, 0x64, 0x62, 0x61, 0x63, 0x6b, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x46, 0x65, 0x65, 0x64, 0x62, 0x61, 0x63, 0x6b, 0x56,
	0x31, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x15, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x0f, 0x22, 0x0d, 0x2f, 0x76, 0x31, 0x2f, 0x66, 0x65, 0x65, 0x64, 0x62, 0x61, 0x63, 0x6b, 0x73,
	0x12, 0x8f, 0x01, 0x0a, 0x15, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4d, 0x75, 0x6c, 0x74, 0x69,
	0x46, 0x65, 0x65, 0x64, 0x62, 0x61, 0x63, 0x6b, 0x56, 0x31, 0x12, 0x2e, 0x2e, 0x6f, 0x63, 0x70,
	0x2e, 0x66, 0x65, 0x65, 0x64, 0x62, 0x61, 0x63, 0x6b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x46, 0x65, 0x65, 0x64, 0x62, 0x61, 0x63,
	0x6b, 0x56, 0x31, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2f, 0x2e, 0x6f, 0x63, 0x70,
	0x2e, 0x66, 0x65, 0x65, 0x64, 0x62, 0x61, 0x63, 0x6b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x46, 0x65, 0x65, 0x64, 0x62, 0x61, 0x63,
	0x6b, 0x56, 0x31, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x15, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x0f, 0x22, 0x0d, 0x2f, 0x76, 0x31, 0x2f, 0x66, 0x65, 0x65, 0x64, 0x62, 0x61, 0x63,
	0x6b, 0x73, 0x12, 0x8e, 0x01, 0x0a, 0x10, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x46, 0x65, 0x65,
	0x64, 0x62, 0x61, 0x63, 0x6b, 0x56, 0x31, 0x12, 0x29, 0x2e, 0x6f, 0x63, 0x70, 0x2e, 0x66, 0x65,
	0x65, 0x64, 0x62, 0x61, 0x63, 0x6b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76,
	0x65, 0x46, 0x65, 0x65, 0x64, 0x62, 0x61, 0x63, 0x6b, 0x56, 0x31, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x2a, 0x2e, 0x6f, 0x63, 0x70, 0x2e, 0x66, 0x65, 0x65, 0x64, 0x62, 0x61, 0x63,
	0x6b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x46, 0x65, 0x65, 0x64,
	0x62, 0x61, 0x63, 0x6b, 0x56, 0x31, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x23,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1d, 0x2a, 0x1b, 0x2f, 0x76, 0x31, 0x2f, 0x66, 0x65, 0x65, 0x64,
	0x62, 0x61, 0x63, 0x6b, 0x73, 0x2f, 0x7b, 0x66, 0x65, 0x65, 0x64, 0x62, 0x61, 0x63, 0x6b, 0x5f,
	0x69, 0x64, 0x7d, 0x12, 0x94, 0x01, 0x0a, 0x12, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65,
	0x46, 0x65, 0x65, 0x64, 0x62, 0x61, 0x63, 0x6b, 0x56, 0x31, 0x12, 0x2b, 0x2e, 0x6f, 0x63, 0x70,
	0x2e, 0x66, 0x65, 0x65, 0x64, 0x62, 0x61, 0x63, 0x6b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x44, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x46, 0x65, 0x65, 0x64, 0x62, 0x61, 0x63, 0x6b, 0x56, 0x31,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2c, 0x2e, 0x6f, 0x63, 0x70, 0x2e, 0x66, 0x65,
	0x65, 0x64, 0x62, 0x61, 0x63, 0x6b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x44, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x62, 0x65, 0x46, 0x65, 0x65, 0x64, 0x62, 0x61, 0x63, 0x6b, 0x56, 0x31, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x23, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1d, 0x12, 0x1b, 0x2f,
	0x76, 0x31, 0x2f, 0x66, 0x65, 0x65, 0x64, 0x62, 0x61, 0x63, 0x6b, 0x73, 0x2f, 0x7b, 0x66, 0x65,
	0x65, 0x64, 0x62, 0x61, 0x63, 0x6b, 0x5f, 0x69, 0x64, 0x7d, 0x12, 0x7d, 0x0a, 0x0f, 0x4c, 0x69,
	0x73, 0x74, 0x46, 0x65, 0x65, 0x64, 0x62, 0x61, 0x63, 0x6b, 0x73, 0x56, 0x31, 0x12, 0x28, 0x2e,
	0x6f, 0x63, 0x70, 0x2e, 0x66, 0x65, 0x65, 0x64, 0x62, 0x61, 0x63, 0x6b, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x4c, 0x69, 0x73, 0x74, 0x46, 0x65, 0x65, 0x64, 0x62, 0x61, 0x63, 0x6b, 0x73, 0x56, 0x31,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x29, 0x2e, 0x6f, 0x63, 0x70, 0x2e, 0x66, 0x65,
	0x65, 0x64, 0x62, 0x61, 0x63, 0x6b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x46,
	0x65, 0x65, 0x64, 0x62, 0x61, 0x63, 0x6b, 0x73, 0x56, 0x31, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x15, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0f, 0x12, 0x0d, 0x2f, 0x76, 0x31, 0x2f,
	0x66, 0x65, 0x65, 0x64, 0x62, 0x61, 0x63, 0x6b, 0x73, 0x12, 0x80, 0x01, 0x0a, 0x10, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x61, 0x6c, 0x56, 0x31, 0x12, 0x29,
	0x2e, 0x6f, 0x63, 0x70, 0x2e, 0x66, 0x65, 0x65, 0x64, 0x62, 0x61, 0x63, 0x6b, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x61, 0x6c,
	0x56, 0x31, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2a, 0x2e, 0x6f, 0x63, 0x70, 0x2e,
	0x66, 0x65, 0x65, 0x64, 0x62, 0x61, 0x63, 0x6b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x61, 0x6c, 0x56, 0x31, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x15, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0f, 0x22, 0x0d, 0x2f,
	0x76, 0x31, 0x2f, 0x70, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x61, 0x6c, 0x73, 0x12, 0x8f, 0x01, 0x0a,
	0x15, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x50, 0x72, 0x6f, 0x70,
	0x6f, 0x73, 0x61, 0x6c, 0x56, 0x31, 0x12, 0x2e, 0x2e, 0x6f, 0x63, 0x70, 0x2e, 0x66, 0x65, 0x65,
	0x64, 0x62, 0x61, 0x63, 0x6b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x4d, 0x75, 0x6c, 0x74, 0x69, 0x50, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x61, 0x6c, 0x56, 0x31, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2f, 0x2e, 0x6f, 0x63, 0x70, 0x2e, 0x66, 0x65, 0x65,
	0x64, 0x62, 0x61, 0x63, 0x6b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x4d, 0x75, 0x6c, 0x74, 0x69, 0x50, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x61, 0x6c, 0x56, 0x31, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x15, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0f, 0x22,
	0x0d, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x61, 0x6c, 0x73, 0x12, 0x8e,
	0x01, 0x0a, 0x10, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x50, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x61,
	0x6c, 0x56, 0x31, 0x12, 0x29, 0x2e, 0x6f, 0x63, 0x70, 0x2e, 0x66, 0x65, 0x65, 0x64, 0x62, 0x61,
	0x63, 0x6b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x50, 0x72, 0x6f,
	0x70, 0x6f, 0x73, 0x61, 0x6c, 0x56, 0x31, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2a,
	0x2e, 0x6f, 0x63, 0x70, 0x2e, 0x66, 0x65, 0x65, 0x64, 0x62, 0x61, 0x63, 0x6b, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x50, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x61, 0x6c,
	0x56, 0x31, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x23, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x1d, 0x2a, 0x1b, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x61, 0x6c,
	0x73, 0x2f, 0x7b, 0x70, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x61, 0x6c, 0x5f, 0x69, 0x64, 0x7d, 0x12,
	0x94, 0x01, 0x0a, 0x12, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x50, 0x72, 0x6f, 0x70,
	0x6f, 0x73, 0x61, 0x6c, 0x56, 0x31, 0x12, 0x2b, 0x2e, 0x6f, 0x63, 0x70, 0x2e, 0x66, 0x65, 0x65,
	0x64, 0x62, 0x61, 0x63, 0x6b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x62, 0x65, 0x50, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x61, 0x6c, 0x56, 0x31, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x2c, 0x2e, 0x6f, 0x63, 0x70, 0x2e, 0x66, 0x65, 0x65, 0x64, 0x62, 0x61,
	0x63, 0x6b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x50,
	0x72, 0x6f, 0x70, 0x6f, 0x73, 0x61, 0x6c, 0x56, 0x31, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x23, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1d, 0x12, 0x1b, 0x2f, 0x76, 0x31, 0x2f, 0x70,
	0x72, 0x6f, 0x70, 0x6f, 0x73, 0x61, 0x6c, 0x73, 0x2f, 0x7b, 0x70, 0x72, 0x6f, 0x70, 0x6f, 0x73,
	0x61, 0x6c, 0x5f, 0x69, 0x64, 0x7d, 0x12, 0x7d, 0x0a, 0x0f, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x72,
	0x6f, 0x70, 0x6f, 0x73, 0x61, 0x6c, 0x73, 0x56, 0x31, 0x12, 0x28, 0x2e, 0x6f, 0x63, 0x70, 0x2e,
	0x66, 0x65, 0x65, 0x64, 0x62, 0x61, 0x63, 0x6b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4c, 0x69, 0x73,
	0x74, 0x50, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x61, 0x6c, 0x73, 0x56, 0x31, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x29, 0x2e, 0x6f, 0x63, 0x70, 0x2e, 0x66, 0x65, 0x65, 0x64, 0x62, 0x61,
	0x63, 0x6b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x70, 0x6f,
	0x73, 0x61, 0x6c, 0x73, 0x56, 0x31, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x15,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0f, 0x12, 0x0d, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x72, 0x6f, 0x70,
	0x6f, 0x73, 0x61, 0x6c, 0x73, 0x42, 0x4a, 0x5a, 0x48, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x6f, 0x7a, 0x6f, 0x6e, 0x63, 0x70, 0x2f, 0x6f, 0x63, 0x70, 0x2d, 0x66,
	0x65, 0x65, 0x64, 0x62, 0x61, 0x63, 0x6b, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x6b, 0x67, 0x2f,
	0x6f, 0x63, 0x70, 0x2d, 0x66, 0x65, 0x65, 0x64, 0x62, 0x61, 0x63, 0x6b, 0x2d, 0x61, 0x70, 0x69,
	0x3b, 0x6f, 0x63, 0x70, 0x5f, 0x66, 0x65, 0x65, 0x64, 0x62, 0x61, 0x63, 0x6b, 0x5f, 0x61, 0x70,
	0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_feedback_service_proto_goTypes = []interface{}{
	(*CreateFeedbackV1Request)(nil),       // 0: ocp.feedback.api.CreateFeedbackV1Request
	(*CreateMultiFeedbackV1Request)(nil),  // 1: ocp.feedback.api.CreateMultiFeedbackV1Request
	(*RemoveFeedbackV1Request)(nil),       // 2: ocp.feedback.api.RemoveFeedbackV1Request
	(*DescribeFeedbackV1Request)(nil),     // 3: ocp.feedback.api.DescribeFeedbackV1Request
	(*ListFeedbacksV1Request)(nil),        // 4: ocp.feedback.api.ListFeedbacksV1Request
	(*CreateProposalV1Request)(nil),       // 5: ocp.feedback.api.CreateProposalV1Request
	(*CreateMultiProposalV1Request)(nil),  // 6: ocp.feedback.api.CreateMultiProposalV1Request
	(*RemoveProposalV1Request)(nil),       // 7: ocp.feedback.api.RemoveProposalV1Request
	(*DescribeProposalV1Request)(nil),     // 8: ocp.feedback.api.DescribeProposalV1Request
	(*ListProposalsV1Request)(nil),        // 9: ocp.feedback.api.ListProposalsV1Request
	(*CreateFeedbackV1Response)(nil),      // 10: ocp.feedback.api.CreateFeedbackV1Response
	(*CreateMultiFeedbackV1Response)(nil), // 11: ocp.feedback.api.CreateMultiFeedbackV1Response
	(*RemoveFeedbackV1Response)(nil),      // 12: ocp.feedback.api.RemoveFeedbackV1Response
	(*DescribeFeedbackV1Response)(nil),    // 13: ocp.feedback.api.DescribeFeedbackV1Response
	(*ListFeedbacksV1Response)(nil),       // 14: ocp.feedback.api.ListFeedbacksV1Response
	(*CreateProposalV1Response)(nil),      // 15: ocp.feedback.api.CreateProposalV1Response
	(*CreateMultiProposalV1Response)(nil), // 16: ocp.feedback.api.CreateMultiProposalV1Response
	(*RemoveProposalV1Response)(nil),      // 17: ocp.feedback.api.RemoveProposalV1Response
	(*DescribeProposalV1Response)(nil),    // 18: ocp.feedback.api.DescribeProposalV1Response
	(*ListProposalsV1Response)(nil),       // 19: ocp.feedback.api.ListProposalsV1Response
}
var file_feedback_service_proto_depIdxs = []int32{
	0,  // 0: ocp.feedback.api.OcpFeedbackApi.CreateFeedbackV1:input_type -> ocp.feedback.api.CreateFeedbackV1Request
	1,  // 1: ocp.feedback.api.OcpFeedbackApi.CreateMultiFeedbackV1:input_type -> ocp.feedback.api.CreateMultiFeedbackV1Request
	2,  // 2: ocp.feedback.api.OcpFeedbackApi.RemoveFeedbackV1:input_type -> ocp.feedback.api.RemoveFeedbackV1Request
	3,  // 3: ocp.feedback.api.OcpFeedbackApi.DescribeFeedbackV1:input_type -> ocp.feedback.api.DescribeFeedbackV1Request
	4,  // 4: ocp.feedback.api.OcpFeedbackApi.ListFeedbacksV1:input_type -> ocp.feedback.api.ListFeedbacksV1Request
	5,  // 5: ocp.feedback.api.OcpFeedbackApi.CreateProposalV1:input_type -> ocp.feedback.api.CreateProposalV1Request
	6,  // 6: ocp.feedback.api.OcpFeedbackApi.CreateMultiProposalV1:input_type -> ocp.feedback.api.CreateMultiProposalV1Request
	7,  // 7: ocp.feedback.api.OcpFeedbackApi.RemoveProposalV1:input_type -> ocp.feedback.api.RemoveProposalV1Request
	8,  // 8: ocp.feedback.api.OcpFeedbackApi.DescribeProposalV1:input_type -> ocp.feedback.api.DescribeProposalV1Request
	9,  // 9: ocp.feedback.api.OcpFeedbackApi.ListProposalsV1:input_type -> ocp.feedback.api.ListProposalsV1Request
	10, // 10: ocp.feedback.api.OcpFeedbackApi.CreateFeedbackV1:output_type -> ocp.feedback.api.CreateFeedbackV1Response
	11, // 11: ocp.feedback.api.OcpFeedbackApi.CreateMultiFeedbackV1:output_type -> ocp.feedback.api.CreateMultiFeedbackV1Response
	12, // 12: ocp.feedback.api.OcpFeedbackApi.RemoveFeedbackV1:output_type -> ocp.feedback.api.RemoveFeedbackV1Response
	13, // 13: ocp.feedback.api.OcpFeedbackApi.DescribeFeedbackV1:output_type -> ocp.feedback.api.DescribeFeedbackV1Response
	14, // 14: ocp.feedback.api.OcpFeedbackApi.ListFeedbacksV1:output_type -> ocp.feedback.api.ListFeedbacksV1Response
	15, // 15: ocp.feedback.api.OcpFeedbackApi.CreateProposalV1:output_type -> ocp.feedback.api.CreateProposalV1Response
	16, // 16: ocp.feedback.api.OcpFeedbackApi.CreateMultiProposalV1:output_type -> ocp.feedback.api.CreateMultiProposalV1Response
	17, // 17: ocp.feedback.api.OcpFeedbackApi.RemoveProposalV1:output_type -> ocp.feedback.api.RemoveProposalV1Response
	18, // 18: ocp.feedback.api.OcpFeedbackApi.DescribeProposalV1:output_type -> ocp.feedback.api.DescribeProposalV1Response
	19, // 19: ocp.feedback.api.OcpFeedbackApi.ListProposalsV1:output_type -> ocp.feedback.api.ListProposalsV1Response
	10, // [10:20] is the sub-list for method output_type
	0,  // [0:10] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_feedback_service_proto_init() }
func file_feedback_service_proto_init() {
	if File_feedback_service_proto != nil {
		return
	}
	file_feedback_messages_proto_init()
	file_proposal_messages_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_feedback_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_feedback_service_proto_goTypes,
		DependencyIndexes: file_feedback_service_proto_depIdxs,
	}.Build()
	File_feedback_service_proto = out.File
	file_feedback_service_proto_rawDesc = nil
	file_feedback_service_proto_goTypes = nil
	file_feedback_service_proto_depIdxs = nil
}
