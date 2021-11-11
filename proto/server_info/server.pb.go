// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: proto/server_info/server.proto

package server_info

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

type ServerCommand int32

const (
	ServerCommand_HELLO_PING                ServerCommand = 0
	ServerCommand_CMD_ID_AUTH               ServerCommand = 1
	ServerCommand_CMD_ID_UPDATE_SERVER_INFO ServerCommand = 2
	ServerCommand_CMD_ID_CLIENT_ACTION      ServerCommand = 3
)

// Enum value maps for ServerCommand.
var (
	ServerCommand_name = map[int32]string{
		0: "HELLO_PING",
		1: "CMD_ID_AUTH",
		2: "CMD_ID_UPDATE_SERVER_INFO",
		3: "CMD_ID_CLIENT_ACTION",
	}
	ServerCommand_value = map[string]int32{
		"HELLO_PING":                0,
		"CMD_ID_AUTH":               1,
		"CMD_ID_UPDATE_SERVER_INFO": 2,
		"CMD_ID_CLIENT_ACTION":      3,
	}
)

func (x ServerCommand) Enum() *ServerCommand {
	p := new(ServerCommand)
	*p = x
	return p
}

func (x ServerCommand) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ServerCommand) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_server_info_server_proto_enumTypes[0].Descriptor()
}

func (ServerCommand) Type() protoreflect.EnumType {
	return &file_proto_server_info_server_proto_enumTypes[0]
}

func (x ServerCommand) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ServerCommand.Descriptor instead.
func (ServerCommand) EnumDescriptor() ([]byte, []int) {
	return file_proto_server_info_server_proto_rawDescGZIP(), []int{0}
}

type ServerReport struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ServerId    string        `protobuf:"bytes,1,opt,name=server_id,json=serverId,proto3" json:"server_id,omitempty"`
	ServerToken []byte        `protobuf:"bytes,2,opt,name=server_token,json=serverToken,proto3" json:"server_token,omitempty"`
	Command     ServerCommand `protobuf:"varint,3,opt,name=command,proto3,enum=server_info.ServerCommand" json:"command,omitempty"`
	CommandBuf  []byte        `protobuf:"bytes,4,opt,name=command_buf,json=commandBuf,proto3" json:"command_buf,omitempty"`
	SeqID       string        `protobuf:"bytes,5,opt,name=SeqID,proto3" json:"SeqID,omitempty"`
}

func (x *ServerReport) Reset() {
	*x = ServerReport{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_server_info_server_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServerReport) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServerReport) ProtoMessage() {}

func (x *ServerReport) ProtoReflect() protoreflect.Message {
	mi := &file_proto_server_info_server_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServerReport.ProtoReflect.Descriptor instead.
func (*ServerReport) Descriptor() ([]byte, []int) {
	return file_proto_server_info_server_proto_rawDescGZIP(), []int{0}
}

func (x *ServerReport) GetServerId() string {
	if x != nil {
		return x.ServerId
	}
	return ""
}

func (x *ServerReport) GetServerToken() []byte {
	if x != nil {
		return x.ServerToken
	}
	return nil
}

func (x *ServerReport) GetCommand() ServerCommand {
	if x != nil {
		return x.Command
	}
	return ServerCommand_HELLO_PING
}

func (x *ServerReport) GetCommandBuf() []byte {
	if x != nil {
		return x.CommandBuf
	}
	return nil
}

func (x *ServerReport) GetSeqID() string {
	if x != nil {
		return x.SeqID
	}
	return ""
}

type ServerInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Process string `protobuf:"bytes,1,opt,name=process,proto3" json:"process,omitempty"` // 进程信息
	Ipv4    string `protobuf:"bytes,2,opt,name=ipv4,proto3" json:"ipv4,omitempty"`
	Ipv6    string `protobuf:"bytes,3,opt,name=ipv6,proto3" json:"ipv6,omitempty"`
}

func (x *ServerInfo) Reset() {
	*x = ServerInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_server_info_server_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServerInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServerInfo) ProtoMessage() {}

func (x *ServerInfo) ProtoReflect() protoreflect.Message {
	mi := &file_proto_server_info_server_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServerInfo.ProtoReflect.Descriptor instead.
func (*ServerInfo) Descriptor() ([]byte, []int) {
	return file_proto_server_info_server_proto_rawDescGZIP(), []int{1}
}

func (x *ServerInfo) GetProcess() string {
	if x != nil {
		return x.Process
	}
	return ""
}

func (x *ServerInfo) GetIpv4() string {
	if x != nil {
		return x.Ipv4
	}
	return ""
}

func (x *ServerInfo) GetIpv6() string {
	if x != nil {
		return x.Ipv6
	}
	return ""
}

type SystemInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Cpu     []*CPU     `protobuf:"bytes,1,rep,name=cpu,proto3" json:"cpu,omitempty"`
	Load    *Load      `protobuf:"bytes,2,opt,name=load,proto3" json:"load,omitempty"`
	Memory  *Mem       `protobuf:"bytes,3,opt,name=memory,proto3" json:"memory,omitempty"`
	Network []*Network `protobuf:"bytes,4,rep,name=network,proto3" json:"network,omitempty"`
	Percent *Percent   `protobuf:"bytes,5,opt,name=Percent,proto3" json:"Percent,omitempty"`
}

func (x *SystemInfo) Reset() {
	*x = SystemInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_server_info_server_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SystemInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SystemInfo) ProtoMessage() {}

func (x *SystemInfo) ProtoReflect() protoreflect.Message {
	mi := &file_proto_server_info_server_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SystemInfo.ProtoReflect.Descriptor instead.
func (*SystemInfo) Descriptor() ([]byte, []int) {
	return file_proto_server_info_server_proto_rawDescGZIP(), []int{2}
}

func (x *SystemInfo) GetCpu() []*CPU {
	if x != nil {
		return x.Cpu
	}
	return nil
}

func (x *SystemInfo) GetLoad() *Load {
	if x != nil {
		return x.Load
	}
	return nil
}

func (x *SystemInfo) GetMemory() *Mem {
	if x != nil {
		return x.Memory
	}
	return nil
}

func (x *SystemInfo) GetNetwork() []*Network {
	if x != nil {
		return x.Network
	}
	return nil
}

func (x *SystemInfo) GetPercent() *Percent {
	if x != nil {
		return x.Percent
	}
	return nil
}

type CPU struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Cores     int32  `protobuf:"varint,1,opt,name=cores,proto3" json:"cores,omitempty"`
	ModelName string `protobuf:"bytes,2,opt,name=model_name,json=modelName,proto3" json:"model_name,omitempty"`
}

func (x *CPU) Reset() {
	*x = CPU{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_server_info_server_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CPU) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CPU) ProtoMessage() {}

func (x *CPU) ProtoReflect() protoreflect.Message {
	mi := &file_proto_server_info_server_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CPU.ProtoReflect.Descriptor instead.
func (*CPU) Descriptor() ([]byte, []int) {
	return file_proto_server_info_server_proto_rawDescGZIP(), []int{3}
}

func (x *CPU) GetCores() int32 {
	if x != nil {
		return x.Cores
	}
	return 0
}

func (x *CPU) GetModelName() string {
	if x != nil {
		return x.ModelName
	}
	return ""
}

type Load struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Load1  float32 `protobuf:"fixed32,1,opt,name=load1,proto3" json:"load1,omitempty"`
	Load5  float32 `protobuf:"fixed32,2,opt,name=load5,proto3" json:"load5,omitempty"`
	Load15 float32 `protobuf:"fixed32,3,opt,name=load15,proto3" json:"load15,omitempty"`
}

func (x *Load) Reset() {
	*x = Load{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_server_info_server_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Load) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Load) ProtoMessage() {}

func (x *Load) ProtoReflect() protoreflect.Message {
	mi := &file_proto_server_info_server_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Load.ProtoReflect.Descriptor instead.
func (*Load) Descriptor() ([]byte, []int) {
	return file_proto_server_info_server_proto_rawDescGZIP(), []int{4}
}

func (x *Load) GetLoad1() float32 {
	if x != nil {
		return x.Load1
	}
	return 0
}

func (x *Load) GetLoad5() float32 {
	if x != nil {
		return x.Load5
	}
	return 0
}

func (x *Load) GetLoad15() float32 {
	if x != nil {
		return x.Load15
	}
	return 0
}

type Mem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Available int64 `protobuf:"varint,1,opt,name=available,proto3" json:"available,omitempty"`
	Total     int64 `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`
	Used      int64 `protobuf:"varint,3,opt,name=used,proto3" json:"used,omitempty"`
}

func (x *Mem) Reset() {
	*x = Mem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_server_info_server_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Mem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Mem) ProtoMessage() {}

func (x *Mem) ProtoReflect() protoreflect.Message {
	mi := &file_proto_server_info_server_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Mem.ProtoReflect.Descriptor instead.
func (*Mem) Descriptor() ([]byte, []int) {
	return file_proto_server_info_server_proto_rawDescGZIP(), []int{5}
}

func (x *Mem) GetAvailable() int64 {
	if x != nil {
		return x.Available
	}
	return 0
}

func (x *Mem) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *Mem) GetUsed() int64 {
	if x != nil {
		return x.Used
	}
	return 0
}

type Network struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Interface string   `protobuf:"bytes,1,opt,name=interface,proto3" json:"interface,omitempty"`
	Address   []string `protobuf:"bytes,2,rep,name=address,proto3" json:"address,omitempty"`
	ByteRecv  int64    `protobuf:"varint,3,opt,name=byte_recv,json=byteRecv,proto3" json:"byte_recv,omitempty"`
	ByteSent  int64    `protobuf:"varint,4,opt,name=byte_sent,json=byteSent,proto3" json:"byte_sent,omitempty"`
}

func (x *Network) Reset() {
	*x = Network{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_server_info_server_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Network) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Network) ProtoMessage() {}

func (x *Network) ProtoReflect() protoreflect.Message {
	mi := &file_proto_server_info_server_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Network.ProtoReflect.Descriptor instead.
func (*Network) Descriptor() ([]byte, []int) {
	return file_proto_server_info_server_proto_rawDescGZIP(), []int{6}
}

func (x *Network) GetInterface() string {
	if x != nil {
		return x.Interface
	}
	return ""
}

func (x *Network) GetAddress() []string {
	if x != nil {
		return x.Address
	}
	return nil
}

func (x *Network) GetByteRecv() int64 {
	if x != nil {
		return x.ByteRecv
	}
	return 0
}

func (x *Network) GetByteSent() int64 {
	if x != nil {
		return x.ByteSent
	}
	return 0
}

type Percent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Cpu    float32 `protobuf:"fixed32,1,opt,name=cpu,proto3" json:"cpu,omitempty"`
	Disk   float32 `protobuf:"fixed32,2,opt,name=disk,proto3" json:"disk,omitempty"`
	Memory float32 `protobuf:"fixed32,3,opt,name=memory,proto3" json:"memory,omitempty"`
	Swap   float32 `protobuf:"fixed32,4,opt,name=swap,proto3" json:"swap,omitempty"`
}

func (x *Percent) Reset() {
	*x = Percent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_server_info_server_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Percent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Percent) ProtoMessage() {}

func (x *Percent) ProtoReflect() protoreflect.Message {
	mi := &file_proto_server_info_server_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Percent.ProtoReflect.Descriptor instead.
func (*Percent) Descriptor() ([]byte, []int) {
	return file_proto_server_info_server_proto_rawDescGZIP(), []int{7}
}

func (x *Percent) GetCpu() float32 {
	if x != nil {
		return x.Cpu
	}
	return 0
}

func (x *Percent) GetDisk() float32 {
	if x != nil {
		return x.Disk
	}
	return 0
}

func (x *Percent) GetMemory() float32 {
	if x != nil {
		return x.Memory
	}
	return 0
}

func (x *Percent) GetSwap() float32 {
	if x != nil {
		return x.Swap
	}
	return 0
}

var File_proto_server_info_server_proto protoreflect.FileDescriptor

var file_proto_server_info_server_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x69,
	0x6e, 0x66, 0x6f, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x0b, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x22, 0xbb, 0x01,
	0x0a, 0x0c, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x1b,
	0x0a, 0x09, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x0b, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x34,
	0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x1a, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x53, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x52, 0x07, 0x63, 0x6f, 0x6d,
	0x6d, 0x61, 0x6e, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x5f,
	0x62, 0x75, 0x66, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0a, 0x63, 0x6f, 0x6d, 0x6d, 0x61,
	0x6e, 0x64, 0x42, 0x75, 0x66, 0x12, 0x14, 0x0a, 0x05, 0x53, 0x65, 0x71, 0x49, 0x44, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x53, 0x65, 0x71, 0x49, 0x44, 0x22, 0x4e, 0x0a, 0x0a, 0x53,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x72, 0x6f,
	0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x72, 0x6f, 0x63,
	0x65, 0x73, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x69, 0x70, 0x76, 0x34, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x69, 0x70, 0x76, 0x34, 0x12, 0x12, 0x0a, 0x04, 0x69, 0x70, 0x76, 0x36, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x69, 0x70, 0x76, 0x36, 0x22, 0xe1, 0x01, 0x0a, 0x0a,
	0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x22, 0x0a, 0x03, 0x63, 0x70,
	0x75, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x43, 0x50, 0x55, 0x52, 0x03, 0x63, 0x70, 0x75, 0x12, 0x25,
	0x0a, 0x04, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x4c, 0x6f, 0x61, 0x64, 0x52,
	0x04, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x28, 0x0a, 0x06, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x69,
	0x6e, 0x66, 0x6f, 0x2e, 0x4d, 0x65, 0x6d, 0x52, 0x06, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x12,
	0x2e, 0x0a, 0x07, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x14, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x4e,
	0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x52, 0x07, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x12,
	0x2e, 0x0a, 0x07, 0x50, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x14, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x50,
	0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x52, 0x07, 0x50, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x22,
	0x3a, 0x0a, 0x03, 0x43, 0x50, 0x55, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x72, 0x65, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x63, 0x6f, 0x72, 0x65, 0x73, 0x12, 0x1d, 0x0a, 0x0a,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x4a, 0x0a, 0x04, 0x4c,
	0x6f, 0x61, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x6f, 0x61, 0x64, 0x31, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x02, 0x52, 0x05, 0x6c, 0x6f, 0x61, 0x64, 0x31, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x6f, 0x61,
	0x64, 0x35, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x6c, 0x6f, 0x61, 0x64, 0x35, 0x12,
	0x16, 0x0a, 0x06, 0x6c, 0x6f, 0x61, 0x64, 0x31, 0x35, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52,
	0x06, 0x6c, 0x6f, 0x61, 0x64, 0x31, 0x35, 0x22, 0x4d, 0x0a, 0x03, 0x4d, 0x65, 0x6d, 0x12, 0x1c,
	0x0a, 0x09, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x09, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x74, 0x6f, 0x74,
	0x61, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x04, 0x75, 0x73, 0x65, 0x64, 0x22, 0x7b, 0x0a, 0x07, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72,
	0x6b, 0x12, 0x1c, 0x0a, 0x09, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x1b, 0x0a, 0x09, 0x62, 0x79, 0x74,
	0x65, 0x5f, 0x72, 0x65, 0x63, 0x76, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x62, 0x79,
	0x74, 0x65, 0x52, 0x65, 0x63, 0x76, 0x12, 0x1b, 0x0a, 0x09, 0x62, 0x79, 0x74, 0x65, 0x5f, 0x73,
	0x65, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x62, 0x79, 0x74, 0x65, 0x53,
	0x65, 0x6e, 0x74, 0x22, 0x5b, 0x0a, 0x07, 0x50, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x12, 0x10,
	0x0a, 0x03, 0x63, 0x70, 0x75, 0x18, 0x01, 0x20, 0x01, 0x28, 0x02, 0x52, 0x03, 0x63, 0x70, 0x75,
	0x12, 0x12, 0x0a, 0x04, 0x64, 0x69, 0x73, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x04,
	0x64, 0x69, 0x73, 0x6b, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x02, 0x52, 0x06, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x12, 0x12, 0x0a, 0x04,
	0x73, 0x77, 0x61, 0x70, 0x18, 0x04, 0x20, 0x01, 0x28, 0x02, 0x52, 0x04, 0x73, 0x77, 0x61, 0x70,
	0x2a, 0x69, 0x0a, 0x0d, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e,
	0x64, 0x12, 0x0e, 0x0a, 0x0a, 0x48, 0x45, 0x4c, 0x4c, 0x4f, 0x5f, 0x50, 0x49, 0x4e, 0x47, 0x10,
	0x00, 0x12, 0x0f, 0x0a, 0x0b, 0x43, 0x4d, 0x44, 0x5f, 0x49, 0x44, 0x5f, 0x41, 0x55, 0x54, 0x48,
	0x10, 0x01, 0x12, 0x1d, 0x0a, 0x19, 0x43, 0x4d, 0x44, 0x5f, 0x49, 0x44, 0x5f, 0x55, 0x50, 0x44,
	0x41, 0x54, 0x45, 0x5f, 0x53, 0x45, 0x52, 0x56, 0x45, 0x52, 0x5f, 0x49, 0x4e, 0x46, 0x4f, 0x10,
	0x02, 0x12, 0x18, 0x0a, 0x14, 0x43, 0x4d, 0x44, 0x5f, 0x49, 0x44, 0x5f, 0x43, 0x4c, 0x49, 0x45,
	0x4e, 0x54, 0x5f, 0x41, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x03, 0x42, 0x28, 0x5a, 0x26, 0x43,
	0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x43, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2d, 0x53, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_server_info_server_proto_rawDescOnce sync.Once
	file_proto_server_info_server_proto_rawDescData = file_proto_server_info_server_proto_rawDesc
)

func file_proto_server_info_server_proto_rawDescGZIP() []byte {
	file_proto_server_info_server_proto_rawDescOnce.Do(func() {
		file_proto_server_info_server_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_server_info_server_proto_rawDescData)
	})
	return file_proto_server_info_server_proto_rawDescData
}

var file_proto_server_info_server_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_proto_server_info_server_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_proto_server_info_server_proto_goTypes = []interface{}{
	(ServerCommand)(0),   // 0: server_info.ServerCommand
	(*ServerReport)(nil), // 1: server_info.ServerReport
	(*ServerInfo)(nil),   // 2: server_info.ServerInfo
	(*SystemInfo)(nil),   // 3: server_info.SystemInfo
	(*CPU)(nil),          // 4: server_info.CPU
	(*Load)(nil),         // 5: server_info.Load
	(*Mem)(nil),          // 6: server_info.Mem
	(*Network)(nil),      // 7: server_info.Network
	(*Percent)(nil),      // 8: server_info.Percent
}
var file_proto_server_info_server_proto_depIdxs = []int32{
	0, // 0: server_info.ServerReport.command:type_name -> server_info.ServerCommand
	4, // 1: server_info.SystemInfo.cpu:type_name -> server_info.CPU
	5, // 2: server_info.SystemInfo.load:type_name -> server_info.Load
	6, // 3: server_info.SystemInfo.memory:type_name -> server_info.Mem
	7, // 4: server_info.SystemInfo.network:type_name -> server_info.Network
	8, // 5: server_info.SystemInfo.Percent:type_name -> server_info.Percent
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_proto_server_info_server_proto_init() }
func file_proto_server_info_server_proto_init() {
	if File_proto_server_info_server_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_server_info_server_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServerReport); i {
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
		file_proto_server_info_server_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServerInfo); i {
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
		file_proto_server_info_server_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SystemInfo); i {
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
		file_proto_server_info_server_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CPU); i {
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
		file_proto_server_info_server_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Load); i {
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
		file_proto_server_info_server_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Mem); i {
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
		file_proto_server_info_server_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Network); i {
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
		file_proto_server_info_server_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Percent); i {
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
			RawDescriptor: file_proto_server_info_server_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_server_info_server_proto_goTypes,
		DependencyIndexes: file_proto_server_info_server_proto_depIdxs,
		EnumInfos:         file_proto_server_info_server_proto_enumTypes,
		MessageInfos:      file_proto_server_info_server_proto_msgTypes,
	}.Build()
	File_proto_server_info_server_proto = out.File
	file_proto_server_info_server_proto_rawDesc = nil
	file_proto_server_info_server_proto_goTypes = nil
	file_proto_server_info_server_proto_depIdxs = nil
}
