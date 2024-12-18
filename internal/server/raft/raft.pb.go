// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v5.28.3
// source: raft.proto

package raft

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

type Entry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Entry) Reset() {
	*x = Entry{}
	mi := &file_raft_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Entry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Entry) ProtoMessage() {}

func (x *Entry) ProtoReflect() protoreflect.Message {
	mi := &file_raft_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Entry.ProtoReflect.Descriptor instead.
func (*Entry) Descriptor() ([]byte, []int) {
	return file_raft_proto_rawDescGZIP(), []int{0}
}

func (x *Entry) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type EntryInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Term         int32    `protobuf:"varint,1,opt,name=term,proto3" json:"term,omitempty"`
	LeaderId     int32    `protobuf:"varint,2,opt,name=leaderId,proto3" json:"leaderId,omitempty"`
	PrevLogIndex int32    `protobuf:"varint,3,opt,name=prevLogIndex,proto3" json:"prevLogIndex,omitempty"`
	PrevLogTerm  int32    `protobuf:"varint,4,opt,name=prevLogTerm,proto3" json:"prevLogTerm,omitempty"`
	LeaderCommit int32    `protobuf:"varint,5,opt,name=leaderCommit,proto3" json:"leaderCommit,omitempty"`
	Entries      []*Entry `protobuf:"bytes,6,rep,name=entries,proto3" json:"entries,omitempty"`
}

func (x *EntryInput) Reset() {
	*x = EntryInput{}
	mi := &file_raft_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *EntryInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EntryInput) ProtoMessage() {}

func (x *EntryInput) ProtoReflect() protoreflect.Message {
	mi := &file_raft_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EntryInput.ProtoReflect.Descriptor instead.
func (*EntryInput) Descriptor() ([]byte, []int) {
	return file_raft_proto_rawDescGZIP(), []int{1}
}

func (x *EntryInput) GetTerm() int32 {
	if x != nil {
		return x.Term
	}
	return 0
}

func (x *EntryInput) GetLeaderId() int32 {
	if x != nil {
		return x.LeaderId
	}
	return 0
}

func (x *EntryInput) GetPrevLogIndex() int32 {
	if x != nil {
		return x.PrevLogIndex
	}
	return 0
}

func (x *EntryInput) GetPrevLogTerm() int32 {
	if x != nil {
		return x.PrevLogTerm
	}
	return 0
}

func (x *EntryInput) GetLeaderCommit() int32 {
	if x != nil {
		return x.LeaderCommit
	}
	return 0
}

func (x *EntryInput) GetEntries() []*Entry {
	if x != nil {
		return x.Entries
	}
	return nil
}

type EntryResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Term    int32 `protobuf:"varint,1,opt,name=term,proto3" json:"term,omitempty"`
	Success bool  `protobuf:"varint,2,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *EntryResult) Reset() {
	*x = EntryResult{}
	mi := &file_raft_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *EntryResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EntryResult) ProtoMessage() {}

func (x *EntryResult) ProtoReflect() protoreflect.Message {
	mi := &file_raft_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EntryResult.ProtoReflect.Descriptor instead.
func (*EntryResult) Descriptor() ([]byte, []int) {
	return file_raft_proto_rawDescGZIP(), []int{2}
}

func (x *EntryResult) GetTerm() int32 {
	if x != nil {
		return x.Term
	}
	return 0
}

func (x *EntryResult) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type VoteInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Term         int32 `protobuf:"varint,1,opt,name=term,proto3" json:"term,omitempty"`
	CandidateId  int32 `protobuf:"varint,2,opt,name=candidateId,proto3" json:"candidateId,omitempty"`
	LastLogIndex int32 `protobuf:"varint,3,opt,name=lastLogIndex,proto3" json:"lastLogIndex,omitempty"`
	LastLogTerm  int32 `protobuf:"varint,4,opt,name=lastLogTerm,proto3" json:"lastLogTerm,omitempty"`
}

func (x *VoteInput) Reset() {
	*x = VoteInput{}
	mi := &file_raft_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *VoteInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VoteInput) ProtoMessage() {}

func (x *VoteInput) ProtoReflect() protoreflect.Message {
	mi := &file_raft_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VoteInput.ProtoReflect.Descriptor instead.
func (*VoteInput) Descriptor() ([]byte, []int) {
	return file_raft_proto_rawDescGZIP(), []int{3}
}

func (x *VoteInput) GetTerm() int32 {
	if x != nil {
		return x.Term
	}
	return 0
}

func (x *VoteInput) GetCandidateId() int32 {
	if x != nil {
		return x.CandidateId
	}
	return 0
}

func (x *VoteInput) GetLastLogIndex() int32 {
	if x != nil {
		return x.LastLogIndex
	}
	return 0
}

func (x *VoteInput) GetLastLogTerm() int32 {
	if x != nil {
		return x.LastLogTerm
	}
	return 0
}

type VoteResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Term        int32 `protobuf:"varint,1,opt,name=term,proto3" json:"term,omitempty"`
	VoteGranted bool  `protobuf:"varint,2,opt,name=voteGranted,proto3" json:"voteGranted,omitempty"`
}

func (x *VoteResult) Reset() {
	*x = VoteResult{}
	mi := &file_raft_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *VoteResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VoteResult) ProtoMessage() {}

func (x *VoteResult) ProtoReflect() protoreflect.Message {
	mi := &file_raft_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VoteResult.ProtoReflect.Descriptor instead.
func (*VoteResult) Descriptor() ([]byte, []int) {
	return file_raft_proto_rawDescGZIP(), []int{4}
}

func (x *VoteResult) GetTerm() int32 {
	if x != nil {
		return x.Term
	}
	return 0
}

func (x *VoteResult) GetVoteGranted() bool {
	if x != nil {
		return x.VoteGranted
	}
	return false
}

type File struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name    string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Content string `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
}

func (x *File) Reset() {
	*x = File{}
	mi := &file_raft_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *File) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*File) ProtoMessage() {}

func (x *File) ProtoReflect() protoreflect.Message {
	mi := &file_raft_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use File.ProtoReflect.Descriptor instead.
func (*File) Descriptor() ([]byte, []int) {
	return file_raft_proto_rawDescGZIP(), []int{5}
}

func (x *File) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *File) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

type Name struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Name) Reset() {
	*x = Name{}
	mi := &file_raft_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Name) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Name) ProtoMessage() {}

func (x *Name) ProtoReflect() protoreflect.Message {
	mi := &file_raft_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Name.ProtoReflect.Descriptor instead.
func (*Name) Descriptor() ([]byte, []int) {
	return file_raft_proto_rawDescGZIP(), []int{6}
}

func (x *Name) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type WriteResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ok bool `protobuf:"varint,1,opt,name=ok,proto3" json:"ok,omitempty"`
}

func (x *WriteResult) Reset() {
	*x = WriteResult{}
	mi := &file_raft_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *WriteResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WriteResult) ProtoMessage() {}

func (x *WriteResult) ProtoReflect() protoreflect.Message {
	mi := &file_raft_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WriteResult.ProtoReflect.Descriptor instead.
func (*WriteResult) Descriptor() ([]byte, []int) {
	return file_raft_proto_rawDescGZIP(), []int{7}
}

func (x *WriteResult) GetOk() bool {
	if x != nil {
		return x.Ok
	}
	return false
}

type ReadResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Conten string `protobuf:"bytes,1,opt,name=conten,proto3" json:"conten,omitempty"`
}

func (x *ReadResult) Reset() {
	*x = ReadResult{}
	mi := &file_raft_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ReadResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadResult) ProtoMessage() {}

func (x *ReadResult) ProtoReflect() protoreflect.Message {
	mi := &file_raft_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadResult.ProtoReflect.Descriptor instead.
func (*ReadResult) Descriptor() ([]byte, []int) {
	return file_raft_proto_rawDescGZIP(), []int{8}
}

func (x *ReadResult) GetConten() string {
	if x != nil {
		return x.Conten
	}
	return ""
}

var File_raft_proto protoreflect.FileDescriptor

var file_raft_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x72, 0x61, 0x66, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x72, 0x61,
	0x66, 0x74, 0x22, 0x1b, 0x0a, 0x05, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22,
	0xcd, 0x01, 0x0a, 0x0a, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x12,
	0x0a, 0x04, 0x74, 0x65, 0x72, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x74, 0x65,
	0x72, 0x6d, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x49, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x6c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x22,
	0x0a, 0x0c, 0x70, 0x72, 0x65, 0x76, 0x4c, 0x6f, 0x67, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x70, 0x72, 0x65, 0x76, 0x4c, 0x6f, 0x67, 0x49, 0x6e, 0x64,
	0x65, 0x78, 0x12, 0x20, 0x0a, 0x0b, 0x70, 0x72, 0x65, 0x76, 0x4c, 0x6f, 0x67, 0x54, 0x65, 0x72,
	0x6d, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x70, 0x72, 0x65, 0x76, 0x4c, 0x6f, 0x67,
	0x54, 0x65, 0x72, 0x6d, 0x12, 0x22, 0x0a, 0x0c, 0x6c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x43, 0x6f,
	0x6d, 0x6d, 0x69, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x6c, 0x65, 0x61, 0x64,
	0x65, 0x72, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x12, 0x25, 0x0a, 0x07, 0x65, 0x6e, 0x74, 0x72,
	0x69, 0x65, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x72, 0x61, 0x66, 0x74,
	0x2e, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x65, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x22,
	0x3b, 0x0a, 0x0b, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x12,
	0x0a, 0x04, 0x74, 0x65, 0x72, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x74, 0x65,
	0x72, 0x6d, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x22, 0x87, 0x01, 0x0a,
	0x09, 0x56, 0x6f, 0x74, 0x65, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65,
	0x72, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x74, 0x65, 0x72, 0x6d, 0x12, 0x20,
	0x0a, 0x0b, 0x63, 0x61, 0x6e, 0x64, 0x69, 0x64, 0x61, 0x74, 0x65, 0x49, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x0b, 0x63, 0x61, 0x6e, 0x64, 0x69, 0x64, 0x61, 0x74, 0x65, 0x49, 0x64,
	0x12, 0x22, 0x0a, 0x0c, 0x6c, 0x61, 0x73, 0x74, 0x4c, 0x6f, 0x67, 0x49, 0x6e, 0x64, 0x65, 0x78,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x6c, 0x61, 0x73, 0x74, 0x4c, 0x6f, 0x67, 0x49,
	0x6e, 0x64, 0x65, 0x78, 0x12, 0x20, 0x0a, 0x0b, 0x6c, 0x61, 0x73, 0x74, 0x4c, 0x6f, 0x67, 0x54,
	0x65, 0x72, 0x6d, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x6c, 0x61, 0x73, 0x74, 0x4c,
	0x6f, 0x67, 0x54, 0x65, 0x72, 0x6d, 0x22, 0x42, 0x0a, 0x0a, 0x56, 0x6f, 0x74, 0x65, 0x52, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x72, 0x6d, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x04, 0x74, 0x65, 0x72, 0x6d, 0x12, 0x20, 0x0a, 0x0b, 0x76, 0x6f, 0x74, 0x65,
	0x47, 0x72, 0x61, 0x6e, 0x74, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x76,
	0x6f, 0x74, 0x65, 0x47, 0x72, 0x61, 0x6e, 0x74, 0x65, 0x64, 0x22, 0x34, 0x0a, 0x04, 0x46, 0x69,
	0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x22, 0x1a, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x1d, 0x0a, 0x0b,
	0x57, 0x72, 0x69, 0x74, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x6f,
	0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x02, 0x6f, 0x6b, 0x22, 0x24, 0x0a, 0x0a, 0x52,
	0x65, 0x61, 0x64, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x63, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x32, 0x75, 0x0a, 0x0b, 0x52, 0x61, 0x66, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x34, 0x0a, 0x0d, 0x41, 0x70, 0x70, 0x65, 0x6e, 0x64, 0x45, 0x6e, 0x74, 0x72, 0x69, 0x65,
	0x73, 0x12, 0x10, 0x2e, 0x72, 0x61, 0x66, 0x74, 0x2e, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x49, 0x6e,
	0x70, 0x75, 0x74, 0x1a, 0x11, 0x2e, 0x72, 0x61, 0x66, 0x74, 0x2e, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x30, 0x0a, 0x0b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x56, 0x6f, 0x74, 0x65, 0x12, 0x0f, 0x2e, 0x72, 0x61, 0x66, 0x74, 0x2e, 0x56, 0x6f, 0x74,
	0x65, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a, 0x10, 0x2e, 0x72, 0x61, 0x66, 0x74, 0x2e, 0x56, 0x6f,
	0x74, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x32, 0x5b, 0x0a, 0x0b, 0x46, 0x69, 0x6c, 0x65,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x26, 0x0a, 0x05, 0x57, 0x72, 0x69, 0x74, 0x65,
	0x12, 0x0a, 0x2e, 0x72, 0x61, 0x66, 0x74, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x1a, 0x11, 0x2e, 0x72,
	0x61, 0x66, 0x74, 0x2e, 0x57, 0x72, 0x69, 0x74, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12,
	0x24, 0x0a, 0x04, 0x52, 0x65, 0x61, 0x64, 0x12, 0x0a, 0x2e, 0x72, 0x61, 0x66, 0x74, 0x2e, 0x4e,
	0x61, 0x6d, 0x65, 0x1a, 0x10, 0x2e, 0x72, 0x61, 0x66, 0x74, 0x2e, 0x52, 0x65, 0x61, 0x64, 0x52,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x42, 0x2f, 0x5a, 0x2d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x73, 0x68, 0x75, 0x70, 0x65, 0x64, 0x6e, 0x65, 0x6b, 0x61, 0x72,
	0x2f, 0x72, 0x61, 0x66, 0x74, 0x2d, 0x67, 0x6f, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61,
	0x6c, 0x2f, 0x72, 0x61, 0x66, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_raft_proto_rawDescOnce sync.Once
	file_raft_proto_rawDescData = file_raft_proto_rawDesc
)

func file_raft_proto_rawDescGZIP() []byte {
	file_raft_proto_rawDescOnce.Do(func() {
		file_raft_proto_rawDescData = protoimpl.X.CompressGZIP(file_raft_proto_rawDescData)
	})
	return file_raft_proto_rawDescData
}

var file_raft_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_raft_proto_goTypes = []any{
	(*Entry)(nil),       // 0: raft.Entry
	(*EntryInput)(nil),  // 1: raft.EntryInput
	(*EntryResult)(nil), // 2: raft.EntryResult
	(*VoteInput)(nil),   // 3: raft.VoteInput
	(*VoteResult)(nil),  // 4: raft.VoteResult
	(*File)(nil),        // 5: raft.File
	(*Name)(nil),        // 6: raft.Name
	(*WriteResult)(nil), // 7: raft.WriteResult
	(*ReadResult)(nil),  // 8: raft.ReadResult
}
var file_raft_proto_depIdxs = []int32{
	0, // 0: raft.EntryInput.entries:type_name -> raft.Entry
	1, // 1: raft.RaftService.AppendEntries:input_type -> raft.EntryInput
	3, // 2: raft.RaftService.RequestVote:input_type -> raft.VoteInput
	5, // 3: raft.FileService.Write:input_type -> raft.File
	6, // 4: raft.FileService.Read:input_type -> raft.Name
	2, // 5: raft.RaftService.AppendEntries:output_type -> raft.EntryResult
	4, // 6: raft.RaftService.RequestVote:output_type -> raft.VoteResult
	7, // 7: raft.FileService.Write:output_type -> raft.WriteResult
	8, // 8: raft.FileService.Read:output_type -> raft.ReadResult
	5, // [5:9] is the sub-list for method output_type
	1, // [1:5] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_raft_proto_init() }
func file_raft_proto_init() {
	if File_raft_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_raft_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_raft_proto_goTypes,
		DependencyIndexes: file_raft_proto_depIdxs,
		MessageInfos:      file_raft_proto_msgTypes,
	}.Build()
	File_raft_proto = out.File
	file_raft_proto_rawDesc = nil
	file_raft_proto_goTypes = nil
	file_raft_proto_depIdxs = nil
}
