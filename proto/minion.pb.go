// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.12.4
// source: proto/minion.proto

package proto

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

type SendCommandToMinion struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Time             string `protobuf:"bytes,1,opt,name=Time,proto3" json:"Time,omitempty"`                  // Time time as a string
	Scheduletask     bool   `protobuf:"varint,2,opt,name=Scheduletask,proto3" json:"Scheduletask,omitempty"` // Don't wait to execute it now
	Timeout          int64  `protobuf:"varint,3,opt,name=Timeout,proto3" json:"Timeout,omitempty"`           // to be used later
	CommandType      string `protobuf:"bytes,4,opt,name=CommandType,proto3" json:"CommandType,omitempty"`    // the command type (based on structs)
	MessageToClient  string `protobuf:"bytes,5,opt,name=MessageToClient,proto3" json:"MessageToClient,omitempty"`
	MarshaledCommand string `protobuf:"bytes,6,opt,name=MarshaledCommand,proto3" json:"MarshaledCommand,omitempty"` // The marshaled struct
}

func (x *SendCommandToMinion) Reset() {
	*x = SendCommandToMinion{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_minion_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendCommandToMinion) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendCommandToMinion) ProtoMessage() {}

func (x *SendCommandToMinion) ProtoReflect() protoreflect.Message {
	mi := &file_proto_minion_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendCommandToMinion.ProtoReflect.Descriptor instead.
func (*SendCommandToMinion) Descriptor() ([]byte, []int) {
	return file_proto_minion_proto_rawDescGZIP(), []int{0}
}

func (x *SendCommandToMinion) GetTime() string {
	if x != nil {
		return x.Time
	}
	return ""
}

func (x *SendCommandToMinion) GetScheduletask() bool {
	if x != nil {
		return x.Scheduletask
	}
	return false
}

func (x *SendCommandToMinion) GetTimeout() int64 {
	if x != nil {
		return x.Timeout
	}
	return 0
}

func (x *SendCommandToMinion) GetCommandType() string {
	if x != nil {
		return x.CommandType
	}
	return ""
}

func (x *SendCommandToMinion) GetMessageToClient() string {
	if x != nil {
		return x.MessageToClient
	}
	return ""
}

func (x *SendCommandToMinion) GetMarshaledCommand() string {
	if x != nil {
		return x.MarshaledCommand
	}
	return ""
}

type CommandFromMinion struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name              string `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
	MessageFromClient string `protobuf:"bytes,2,opt,name=MessageFromClient,proto3" json:"MessageFromClient,omitempty"`
	MessageToClient   string `protobuf:"bytes,3,opt,name=MessageToClient,proto3" json:"MessageToClient,omitempty"`
	ReadyToReceive    bool   `protobuf:"varint,4,opt,name=ReadyToReceive,proto3" json:"ReadyToReceive,omitempty"` // True == client will accept a new command
	Result            string `protobuf:"bytes,5,opt,name=Result,proto3" json:"Result,omitempty"`                  // The stdout/stderr output from the client
	Success           bool   `protobuf:"varint,6,opt,name=Success,proto3" json:"Success,omitempty"`               // Was the command successful
}

func (x *CommandFromMinion) Reset() {
	*x = CommandFromMinion{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_minion_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommandFromMinion) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommandFromMinion) ProtoMessage() {}

func (x *CommandFromMinion) ProtoReflect() protoreflect.Message {
	mi := &file_proto_minion_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommandFromMinion.ProtoReflect.Descriptor instead.
func (*CommandFromMinion) Descriptor() ([]byte, []int) {
	return file_proto_minion_proto_rawDescGZIP(), []int{1}
}

func (x *CommandFromMinion) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CommandFromMinion) GetMessageFromClient() string {
	if x != nil {
		return x.MessageFromClient
	}
	return ""
}

func (x *CommandFromMinion) GetMessageToClient() string {
	if x != nil {
		return x.MessageToClient
	}
	return ""
}

func (x *CommandFromMinion) GetReadyToReceive() bool {
	if x != nil {
		return x.ReadyToReceive
	}
	return false
}

func (x *CommandFromMinion) GetResult() string {
	if x != nil {
		return x.Result
	}
	return ""
}

func (x *CommandFromMinion) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type CreateMinion struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *CreateMinion) Reset() {
	*x = CreateMinion{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_minion_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateMinion) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateMinion) ProtoMessage() {}

func (x *CreateMinion) ProtoReflect() protoreflect.Message {
	mi := &file_proto_minion_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateMinion.ProtoReflect.Descriptor instead.
func (*CreateMinion) Descriptor() ([]byte, []int) {
	return file_proto_minion_proto_rawDescGZIP(), []int{2}
}

func (x *CreateMinion) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type MinionCreateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Authenticated  bool   `protobuf:"varint,1,opt,name=Authenticated,proto3" json:"Authenticated,omitempty"`
	MinionId       int64  `protobuf:"varint,2,opt,name=MinionId,proto3" json:"MinionId,omitempty"`
	RegisteredName string `protobuf:"bytes,3,opt,name=RegisteredName,proto3" json:"RegisteredName,omitempty"`
}

func (x *MinionCreateResponse) Reset() {
	*x = MinionCreateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_minion_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MinionCreateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MinionCreateResponse) ProtoMessage() {}

func (x *MinionCreateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_minion_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MinionCreateResponse.ProtoReflect.Descriptor instead.
func (*MinionCreateResponse) Descriptor() ([]byte, []int) {
	return file_proto_minion_proto_rawDescGZIP(), []int{3}
}

func (x *MinionCreateResponse) GetAuthenticated() bool {
	if x != nil {
		return x.Authenticated
	}
	return false
}

func (x *MinionCreateResponse) GetMinionId() int64 {
	if x != nil {
		return x.MinionId
	}
	return 0
}

func (x *MinionCreateResponse) GetRegisteredName() string {
	if x != nil {
		return x.RegisteredName
	}
	return ""
}

type CmdRunFromClient struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name              string `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
	MessageFromClient string `protobuf:"bytes,2,opt,name=MessageFromClient,proto3" json:"MessageFromClient,omitempty"`
	MessageToClient   string `protobuf:"bytes,3,opt,name=MessageToClient,proto3" json:"MessageToClient,omitempty"`
	TargetMinions     string `protobuf:"bytes,4,opt,name=TargetMinions,proto3" json:"TargetMinions,omitempty"`
	Command           string `protobuf:"bytes,5,opt,name=Command,proto3" json:"Command,omitempty"`
}

func (x *CmdRunFromClient) Reset() {
	*x = CmdRunFromClient{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_minion_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CmdRunFromClient) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CmdRunFromClient) ProtoMessage() {}

func (x *CmdRunFromClient) ProtoReflect() protoreflect.Message {
	mi := &file_proto_minion_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CmdRunFromClient.ProtoReflect.Descriptor instead.
func (*CmdRunFromClient) Descriptor() ([]byte, []int) {
	return file_proto_minion_proto_rawDescGZIP(), []int{4}
}

func (x *CmdRunFromClient) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CmdRunFromClient) GetMessageFromClient() string {
	if x != nil {
		return x.MessageFromClient
	}
	return ""
}

func (x *CmdRunFromClient) GetMessageToClient() string {
	if x != nil {
		return x.MessageToClient
	}
	return ""
}

func (x *CmdRunFromClient) GetTargetMinions() string {
	if x != nil {
		return x.TargetMinions
	}
	return ""
}

func (x *CmdRunFromClient) GetCommand() string {
	if x != nil {
		return x.Command
	}
	return ""
}

type CmdRunReturnResultToMinion struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MinionCmdResult []*MinionCmdRunResult `protobuf:"bytes,1,rep,name=minionCmdResult,proto3" json:"minionCmdResult,omitempty"` //
}

func (x *CmdRunReturnResultToMinion) Reset() {
	*x = CmdRunReturnResultToMinion{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_minion_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CmdRunReturnResultToMinion) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CmdRunReturnResultToMinion) ProtoMessage() {}

func (x *CmdRunReturnResultToMinion) ProtoReflect() protoreflect.Message {
	mi := &file_proto_minion_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CmdRunReturnResultToMinion.ProtoReflect.Descriptor instead.
func (*CmdRunReturnResultToMinion) Descriptor() ([]byte, []int) {
	return file_proto_minion_proto_rawDescGZIP(), []int{5}
}

func (x *CmdRunReturnResultToMinion) GetMinionCmdResult() []*MinionCmdRunResult {
	if x != nil {
		return x.MinionCmdResult
	}
	return nil
}

type MinionCmdRunResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool   `protobuf:"varint,1,opt,name=Success,proto3" json:"Success,omitempty"` // True == the command was successful, False == the command failed
	Result  string `protobuf:"bytes,2,opt,name=Result,proto3" json:"Result,omitempty"`    // The stdout/stderr output from the minion that ran the command
}

func (x *MinionCmdRunResult) Reset() {
	*x = MinionCmdRunResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_minion_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MinionCmdRunResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MinionCmdRunResult) ProtoMessage() {}

func (x *MinionCmdRunResult) ProtoReflect() protoreflect.Message {
	mi := &file_proto_minion_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MinionCmdRunResult.ProtoReflect.Descriptor instead.
func (*MinionCmdRunResult) Descriptor() ([]byte, []int) {
	return file_proto_minion_proto_rawDescGZIP(), []int{6}
}

func (x *MinionCmdRunResult) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *MinionCmdRunResult) GetResult() string {
	if x != nil {
		return x.Result
	}
	return ""
}

var File_proto_minion_proto protoreflect.FileDescriptor

var file_proto_minion_proto_rawDesc = []byte{
	0x0a, 0x12, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6d, 0x69, 0x6e, 0x69, 0x6f, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x6d, 0x69, 0x6e, 0x69, 0x6f, 0x6e, 0x22, 0xdf, 0x01, 0x0a,
	0x13, 0x53, 0x65, 0x6e, 0x64, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x54, 0x6f, 0x4d, 0x69,
	0x6e, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x53, 0x63, 0x68, 0x65,
	0x64, 0x75, 0x6c, 0x65, 0x74, 0x61, 0x73, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c,
	0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x74, 0x61, 0x73, 0x6b, 0x12, 0x18, 0x0a, 0x07,
	0x54, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x54,
	0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e,
	0x64, 0x54, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x43, 0x6f, 0x6d,
	0x6d, 0x61, 0x6e, 0x64, 0x54, 0x79, 0x70, 0x65, 0x12, 0x28, 0x0a, 0x0f, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x54, 0x6f, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0f, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x43, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x12, 0x2a, 0x0a, 0x10, 0x4d, 0x61, 0x72, 0x73, 0x68, 0x61, 0x6c, 0x65, 0x64, 0x43,
	0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x4d, 0x61,
	0x72, 0x73, 0x68, 0x61, 0x6c, 0x65, 0x64, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x22, 0xd9,
	0x01, 0x0a, 0x11, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x46, 0x72, 0x6f, 0x6d, 0x4d, 0x69,
	0x6e, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x2c, 0x0a, 0x11, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x46, 0x72, 0x6f, 0x6d, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x11, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x46, 0x72, 0x6f, 0x6d,
	0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x12, 0x28, 0x0a, 0x0f, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x54, 0x6f, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0f, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x12, 0x26, 0x0a, 0x0e, 0x52, 0x65, 0x61, 0x64, 0x79, 0x54, 0x6f, 0x52, 0x65, 0x63, 0x65, 0x69,
	0x76, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0e, 0x52, 0x65, 0x61, 0x64, 0x79, 0x54,
	0x6f, 0x52, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x52, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x12, 0x18, 0x0a, 0x07, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x07, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x22, 0x22, 0x0a, 0x0c, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x4d, 0x69, 0x6e, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x80,
	0x01, 0x0a, 0x14, 0x4d, 0x69, 0x6e, 0x69, 0x6f, 0x6e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x24, 0x0a, 0x0d, 0x41, 0x75, 0x74, 0x68, 0x65,
	0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0d,
	0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x65, 0x64, 0x12, 0x1a, 0x0a,
	0x08, 0x4d, 0x69, 0x6e, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x08, 0x4d, 0x69, 0x6e, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x26, 0x0a, 0x0e, 0x52, 0x65, 0x67,
	0x69, 0x73, 0x74, 0x65, 0x72, 0x65, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x65, 0x64, 0x4e, 0x61, 0x6d,
	0x65, 0x22, 0xbe, 0x01, 0x0a, 0x10, 0x43, 0x6d, 0x64, 0x52, 0x75, 0x6e, 0x46, 0x72, 0x6f, 0x6d,
	0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x2c, 0x0a, 0x11, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x46, 0x72, 0x6f, 0x6d, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x46, 0x72,
	0x6f, 0x6d, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x12, 0x28, 0x0a, 0x0f, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x54, 0x6f, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0f, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x43, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x12, 0x24, 0x0a, 0x0d, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x4d, 0x69, 0x6e, 0x69,
	0x6f, 0x6e, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x54, 0x61, 0x72, 0x67, 0x65,
	0x74, 0x4d, 0x69, 0x6e, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x43, 0x6f, 0x6d, 0x6d,
	0x61, 0x6e, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x43, 0x6f, 0x6d, 0x6d, 0x61,
	0x6e, 0x64, 0x22, 0x62, 0x0a, 0x1a, 0x43, 0x6d, 0x64, 0x52, 0x75, 0x6e, 0x52, 0x65, 0x74, 0x75,
	0x72, 0x6e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x54, 0x6f, 0x4d, 0x69, 0x6e, 0x69, 0x6f, 0x6e,
	0x12, 0x44, 0x0a, 0x0f, 0x6d, 0x69, 0x6e, 0x69, 0x6f, 0x6e, 0x43, 0x6d, 0x64, 0x52, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x6d, 0x69, 0x6e, 0x69,
	0x6f, 0x6e, 0x2e, 0x4d, 0x69, 0x6e, 0x69, 0x6f, 0x6e, 0x43, 0x6d, 0x64, 0x52, 0x75, 0x6e, 0x52,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x0f, 0x6d, 0x69, 0x6e, 0x69, 0x6f, 0x6e, 0x43, 0x6d, 0x64,
	0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x46, 0x0a, 0x12, 0x4d, 0x69, 0x6e, 0x69, 0x6f, 0x6e,
	0x43, 0x6d, 0x64, 0x52, 0x75, 0x6e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x18, 0x0a, 0x07,
	0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x53,
	0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x32, 0xf1,
	0x01, 0x0a, 0x0d, 0x4d, 0x69, 0x6e, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x49, 0x0a, 0x11, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x4e, 0x65, 0x77, 0x4d,
	0x69, 0x6e, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x2e, 0x6d, 0x69, 0x6e, 0x69, 0x6f, 0x6e, 0x2e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x4d, 0x69, 0x6e, 0x69, 0x6f, 0x6e, 0x1a, 0x1c, 0x2e, 0x6d, 0x69,
	0x6e, 0x69, 0x6f, 0x6e, 0x2e, 0x4d, 0x69, 0x6e, 0x69, 0x6f, 0x6e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x49, 0x0a, 0x0b, 0x47,
	0x65, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x73, 0x12, 0x19, 0x2e, 0x6d, 0x69, 0x6e,
	0x69, 0x6f, 0x6e, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x46, 0x72, 0x6f, 0x6d, 0x4d,
	0x69, 0x6e, 0x69, 0x6f, 0x6e, 0x1a, 0x1b, 0x2e, 0x6d, 0x69, 0x6e, 0x69, 0x6f, 0x6e, 0x2e, 0x53,
	0x65, 0x6e, 0x64, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x54, 0x6f, 0x4d, 0x69, 0x6e, 0x69,
	0x6f, 0x6e, 0x22, 0x00, 0x30, 0x01, 0x12, 0x4a, 0x0a, 0x06, 0x43, 0x6d, 0x64, 0x52, 0x75, 0x6e,
	0x12, 0x18, 0x2e, 0x6d, 0x69, 0x6e, 0x69, 0x6f, 0x6e, 0x2e, 0x43, 0x6d, 0x64, 0x52, 0x75, 0x6e,
	0x46, 0x72, 0x6f, 0x6d, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x1a, 0x22, 0x2e, 0x6d, 0x69, 0x6e,
	0x69, 0x6f, 0x6e, 0x2e, 0x43, 0x6d, 0x64, 0x52, 0x75, 0x6e, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e,
	0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x54, 0x6f, 0x4d, 0x69, 0x6e, 0x69, 0x6f, 0x6e, 0x22, 0x00,
	0x30, 0x01, 0x42, 0x25, 0x5a, 0x23, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x62, 0x62, 0x72, 0x75, 0x75, 0x6e, 0x2f, 0x62, 0x62, 0x72, 0x75, 0x75, 0x6e, 0x2f, 0x67,
	0x61, 0x6c, 0x74, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_proto_minion_proto_rawDescOnce sync.Once
	file_proto_minion_proto_rawDescData = file_proto_minion_proto_rawDesc
)

func file_proto_minion_proto_rawDescGZIP() []byte {
	file_proto_minion_proto_rawDescOnce.Do(func() {
		file_proto_minion_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_minion_proto_rawDescData)
	})
	return file_proto_minion_proto_rawDescData
}

var file_proto_minion_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_proto_minion_proto_goTypes = []interface{}{
	(*SendCommandToMinion)(nil),        // 0: minion.SendCommandToMinion
	(*CommandFromMinion)(nil),          // 1: minion.CommandFromMinion
	(*CreateMinion)(nil),               // 2: minion.CreateMinion
	(*MinionCreateResponse)(nil),       // 3: minion.MinionCreateResponse
	(*CmdRunFromClient)(nil),           // 4: minion.CmdRunFromClient
	(*CmdRunReturnResultToMinion)(nil), // 5: minion.CmdRunReturnResultToMinion
	(*MinionCmdRunResult)(nil),         // 6: minion.MinionCmdRunResult
}
var file_proto_minion_proto_depIdxs = []int32{
	6, // 0: minion.CmdRunReturnResultToMinion.minionCmdResult:type_name -> minion.MinionCmdRunResult
	2, // 1: minion.MinionService.RegisterNewMinion:input_type -> minion.CreateMinion
	1, // 2: minion.MinionService.GetCommands:input_type -> minion.CommandFromMinion
	4, // 3: minion.MinionService.CmdRun:input_type -> minion.CmdRunFromClient
	3, // 4: minion.MinionService.RegisterNewMinion:output_type -> minion.MinionCreateResponse
	0, // 5: minion.MinionService.GetCommands:output_type -> minion.SendCommandToMinion
	5, // 6: minion.MinionService.CmdRun:output_type -> minion.CmdRunReturnResultToMinion
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_minion_proto_init() }
func file_proto_minion_proto_init() {
	if File_proto_minion_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_minion_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendCommandToMinion); i {
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
		file_proto_minion_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommandFromMinion); i {
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
		file_proto_minion_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateMinion); i {
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
		file_proto_minion_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MinionCreateResponse); i {
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
		file_proto_minion_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CmdRunFromClient); i {
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
		file_proto_minion_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CmdRunReturnResultToMinion); i {
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
		file_proto_minion_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MinionCmdRunResult); i {
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
			RawDescriptor: file_proto_minion_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_minion_proto_goTypes,
		DependencyIndexes: file_proto_minion_proto_depIdxs,
		MessageInfos:      file_proto_minion_proto_msgTypes,
	}.Build()
	File_proto_minion_proto = out.File
	file_proto_minion_proto_rawDesc = nil
	file_proto_minion_proto_goTypes = nil
	file_proto_minion_proto_depIdxs = nil
}
