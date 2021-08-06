// Code generated by protoc-gen-go. DO NOT EDIT.
// source: event.proto

package event

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type IngestRequest_Entity int32

const (
	IngestRequest_UNKNOWN_ENTITY IngestRequest_Entity = 0
	IngestRequest_PRODUCT        IngestRequest_Entity = 1
	IngestRequest_CATEGORY       IngestRequest_Entity = 2
	IngestRequest_BASKET         IngestRequest_Entity = 3
)

var IngestRequest_Entity_name = map[int32]string{
	0: "UNKNOWN_ENTITY",
	1: "PRODUCT",
	2: "CATEGORY",
	3: "BASKET",
}

var IngestRequest_Entity_value = map[string]int32{
	"UNKNOWN_ENTITY": 0,
	"PRODUCT":        1,
	"CATEGORY":       2,
	"BASKET":         3,
}

func (x IngestRequest_Entity) String() string {
	return proto.EnumName(IngestRequest_Entity_name, int32(x))
}

func (IngestRequest_Entity) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_2d17a9d3f0ddf27e, []int{0, 0}
}

type IngestRequest_Operation int32

const (
	IngestRequest_UNKNOWN_OPERATION IngestRequest_Operation = 0
	IngestRequest_CREATED           IngestRequest_Operation = 1
	IngestRequest_UPDATED           IngestRequest_Operation = 2
	IngestRequest_DELETED           IngestRequest_Operation = 3
)

var IngestRequest_Operation_name = map[int32]string{
	0: "UNKNOWN_OPERATION",
	1: "CREATED",
	2: "UPDATED",
	3: "DELETED",
}

var IngestRequest_Operation_value = map[string]int32{
	"UNKNOWN_OPERATION": 0,
	"CREATED":           1,
	"UPDATED":           2,
	"DELETED":           3,
}

func (x IngestRequest_Operation) String() string {
	return proto.EnumName(IngestRequest_Operation_name, int32(x))
}

func (IngestRequest_Operation) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_2d17a9d3f0ddf27e, []int{0, 1}
}

type IngestRequest struct {
	RequestId            string                  `protobuf:"bytes,1,opt,name=requestId,proto3" json:"requestId,omitempty"`
	Entity               IngestRequest_Entity    `protobuf:"varint,2,opt,name=entity,proto3,enum=event.IngestRequest_Entity" json:"entity,omitempty"`
	Operation            IngestRequest_Operation `protobuf:"varint,3,opt,name=operation,proto3,enum=event.IngestRequest_Operation" json:"operation,omitempty"`
	Payload              string                  `protobuf:"bytes,4,opt,name=payload,proto3" json:"payload,omitempty"`
	Timestamp            *timestamp.Timestamp    `protobuf:"bytes,5,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *IngestRequest) Reset()         { *m = IngestRequest{} }
func (m *IngestRequest) String() string { return proto.CompactTextString(m) }
func (*IngestRequest) ProtoMessage()    {}
func (*IngestRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2d17a9d3f0ddf27e, []int{0}
}

func (m *IngestRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IngestRequest.Unmarshal(m, b)
}
func (m *IngestRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IngestRequest.Marshal(b, m, deterministic)
}
func (m *IngestRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IngestRequest.Merge(m, src)
}
func (m *IngestRequest) XXX_Size() int {
	return xxx_messageInfo_IngestRequest.Size(m)
}
func (m *IngestRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_IngestRequest.DiscardUnknown(m)
}

var xxx_messageInfo_IngestRequest proto.InternalMessageInfo

func (m *IngestRequest) GetRequestId() string {
	if m != nil {
		return m.RequestId
	}
	return ""
}

func (m *IngestRequest) GetEntity() IngestRequest_Entity {
	if m != nil {
		return m.Entity
	}
	return IngestRequest_UNKNOWN_ENTITY
}

func (m *IngestRequest) GetOperation() IngestRequest_Operation {
	if m != nil {
		return m.Operation
	}
	return IngestRequest_UNKNOWN_OPERATION
}

func (m *IngestRequest) GetPayload() string {
	if m != nil {
		return m.Payload
	}
	return ""
}

func (m *IngestRequest) GetTimestamp() *timestamp.Timestamp {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

type IngestResponse struct {
	RequestId            string               `protobuf:"bytes,1,opt,name=requestId,proto3" json:"requestId,omitempty"`
	IsSuccessful         *wrappers.BoolValue  `protobuf:"bytes,2,opt,name=isSuccessful,proto3" json:"isSuccessful,omitempty"`
	Message              string               `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
	Timestamp            *timestamp.Timestamp `protobuf:"bytes,4,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *IngestResponse) Reset()         { *m = IngestResponse{} }
func (m *IngestResponse) String() string { return proto.CompactTextString(m) }
func (*IngestResponse) ProtoMessage()    {}
func (*IngestResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_2d17a9d3f0ddf27e, []int{1}
}

func (m *IngestResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IngestResponse.Unmarshal(m, b)
}
func (m *IngestResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IngestResponse.Marshal(b, m, deterministic)
}
func (m *IngestResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IngestResponse.Merge(m, src)
}
func (m *IngestResponse) XXX_Size() int {
	return xxx_messageInfo_IngestResponse.Size(m)
}
func (m *IngestResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_IngestResponse.DiscardUnknown(m)
}

var xxx_messageInfo_IngestResponse proto.InternalMessageInfo

func (m *IngestResponse) GetRequestId() string {
	if m != nil {
		return m.RequestId
	}
	return ""
}

func (m *IngestResponse) GetIsSuccessful() *wrappers.BoolValue {
	if m != nil {
		return m.IsSuccessful
	}
	return nil
}

func (m *IngestResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *IngestResponse) GetTimestamp() *timestamp.Timestamp {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

func init() {
	proto.RegisterEnum("event.IngestRequest_Entity", IngestRequest_Entity_name, IngestRequest_Entity_value)
	proto.RegisterEnum("event.IngestRequest_Operation", IngestRequest_Operation_name, IngestRequest_Operation_value)
	proto.RegisterType((*IngestRequest)(nil), "event.IngestRequest")
	proto.RegisterType((*IngestResponse)(nil), "event.IngestResponse")
}

func init() { proto.RegisterFile("event.proto", fileDescriptor_2d17a9d3f0ddf27e) }

var fileDescriptor_2d17a9d3f0ddf27e = []byte{
	// 424 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0xcf, 0x6a, 0xdb, 0x40,
	0x10, 0xc6, 0x23, 0x3b, 0x91, 0xab, 0x71, 0x6a, 0xd4, 0xa5, 0x01, 0xe1, 0x96, 0xd4, 0xe8, 0xe4,
	0x93, 0x0c, 0xce, 0xa1, 0x3d, 0x94, 0x82, 0xff, 0x2c, 0xad, 0x48, 0x91, 0xcc, 0x5a, 0x6e, 0x49,
	0x2f, 0x45, 0x71, 0x26, 0x46, 0x20, 0x6b, 0x55, 0xed, 0x2a, 0x25, 0xf7, 0x3e, 0x59, 0x9f, 0xac,
	0x68, 0xb7, 0x1b, 0x63, 0x63, 0x28, 0xbd, 0xe9, 0x9b, 0xf9, 0xcd, 0x68, 0xbe, 0x99, 0x85, 0x2e,
	0x3e, 0x60, 0x21, 0x83, 0xb2, 0xe2, 0x92, 0x93, 0x33, 0x25, 0xfa, 0x97, 0x1b, 0xce, 0x37, 0x39,
	0x8e, 0x54, 0xf0, 0xb6, 0xbe, 0x1f, 0xfd, 0xac, 0xd2, 0xb2, 0xc4, 0x4a, 0x68, 0xac, 0xff, 0xe6,
	0x30, 0x2f, 0xb3, 0x2d, 0x0a, 0x99, 0x6e, 0x4b, 0x0d, 0xf8, 0xbf, 0xda, 0xf0, 0x3c, 0x2c, 0x36,
	0x28, 0x24, 0xc3, 0x1f, 0x35, 0x0a, 0x49, 0x5e, 0x83, 0x53, 0xe9, 0xcf, 0xf0, 0xce, 0xb3, 0x06,
	0xd6, 0xd0, 0x61, 0xbb, 0x00, 0xb9, 0x02, 0x1b, 0x0b, 0x99, 0xc9, 0x47, 0xaf, 0x35, 0xb0, 0x86,
	0xbd, 0xf1, 0xab, 0x40, 0x4f, 0xb5, 0xd7, 0x23, 0xa0, 0x0a, 0x61, 0x7f, 0x51, 0xf2, 0x1e, 0x1c,
	0x5e, 0x62, 0x95, 0xca, 0x8c, 0x17, 0x5e, 0x5b, 0xd5, 0x5d, 0x1e, 0xad, 0x8b, 0x0d, 0xc5, 0x76,
	0x05, 0xc4, 0x83, 0x4e, 0x99, 0x3e, 0xe6, 0x3c, 0xbd, 0xf3, 0x4e, 0xd5, 0x38, 0x46, 0x92, 0x77,
	0xe0, 0x3c, 0xf9, 0xf1, 0xce, 0x06, 0xd6, 0xb0, 0x3b, 0xee, 0x07, 0xda, 0x71, 0x60, 0x1c, 0x07,
	0x89, 0x21, 0xd8, 0x0e, 0xf6, 0x67, 0x60, 0xeb, 0x19, 0x09, 0x81, 0xde, 0x2a, 0xba, 0x8e, 0xe2,
	0xaf, 0xd1, 0x77, 0x1a, 0x25, 0x61, 0x72, 0xe3, 0x9e, 0x90, 0x2e, 0x74, 0x16, 0x2c, 0x9e, 0xaf,
	0x66, 0x89, 0x6b, 0x91, 0x73, 0x78, 0x36, 0x9b, 0x24, 0xf4, 0x63, 0xcc, 0x6e, 0xdc, 0x16, 0x01,
	0xb0, 0xa7, 0x93, 0xe5, 0x35, 0x4d, 0xdc, 0xb6, 0x1f, 0x82, 0xf3, 0x34, 0x30, 0xb9, 0x80, 0x17,
	0xa6, 0x4f, 0xbc, 0xa0, 0x6c, 0x92, 0x84, 0x71, 0xa4, 0x5b, 0xcd, 0x18, 0x9d, 0x24, 0x74, 0xee,
	0x5a, 0x8d, 0x58, 0x2d, 0xe6, 0x4a, 0xb4, 0x1a, 0x31, 0xa7, 0x9f, 0x69, 0x23, 0xda, 0xfe, 0x6f,
	0x0b, 0x7a, 0x66, 0x15, 0xa2, 0xe4, 0x85, 0xc0, 0x7f, 0xdc, 0xe1, 0x03, 0x9c, 0x67, 0x62, 0x59,
	0xaf, 0xd7, 0x28, 0xc4, 0x7d, 0x9d, 0xab, 0x6b, 0x1c, 0x73, 0x3f, 0xe5, 0x3c, 0xff, 0x92, 0xe6,
	0x35, 0xb2, 0x3d, 0xbe, 0x59, 0xea, 0x16, 0x85, 0x48, 0x37, 0xa8, 0x0e, 0xe2, 0x30, 0x23, 0xf7,
	0x97, 0x7a, 0xfa, 0x1f, 0x4b, 0x1d, 0x7f, 0x32, 0x4f, 0x69, 0x89, 0xd5, 0x43, 0xb6, 0x46, 0xf2,
	0x16, 0x6c, 0x1d, 0x20, 0x2f, 0x8f, 0x9d, 0xbb, 0x7f, 0x71, 0x10, 0xd5, 0xce, 0xfd, 0x93, 0xa9,
	0xf3, 0xad, 0x13, 0x8c, 0x54, 0xee, 0xd6, 0x56, 0xff, 0xbc, 0xfa, 0x13, 0x00, 0x00, 0xff, 0xff,
	0xb2, 0xd7, 0xec, 0xfb, 0xfe, 0x02, 0x00, 0x00,
}