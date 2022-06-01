// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.14.0
// source: fleet/pb/fleet.proto

package pb

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

type AgentByIDReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AgentID string `protobuf:"bytes,1,opt,name=agentID,proto3" json:"agentID,omitempty"`
	OwnerID string `protobuf:"bytes,2,opt,name=ownerID,proto3" json:"ownerID,omitempty"`
}

func (x *AgentByIDReq) Reset() {
	*x = AgentByIDReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fleet_pb_fleet_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AgentByIDReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AgentByIDReq) ProtoMessage() {}

func (x *AgentByIDReq) ProtoReflect() protoreflect.Message {
	mi := &file_fleet_pb_fleet_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AgentByIDReq.ProtoReflect.Descriptor instead.
func (*AgentByIDReq) Descriptor() ([]byte, []int) {
	return file_fleet_pb_fleet_proto_rawDescGZIP(), []int{0}
}

func (x *AgentByIDReq) GetAgentID() string {
	if x != nil {
		return x.AgentID
	}
	return ""
}

func (x *AgentByIDReq) GetOwnerID() string {
	if x != nil {
		return x.OwnerID
	}
	return ""
}

type AgentRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name    string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Channel string `protobuf:"bytes,3,opt,name=channel,proto3" json:"channel,omitempty"`
}

func (x *AgentRes) Reset() {
	*x = AgentRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fleet_pb_fleet_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AgentRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AgentRes) ProtoMessage() {}

func (x *AgentRes) ProtoReflect() protoreflect.Message {
	mi := &file_fleet_pb_fleet_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AgentRes.ProtoReflect.Descriptor instead.
func (*AgentRes) Descriptor() ([]byte, []int) {
	return file_fleet_pb_fleet_proto_rawDescGZIP(), []int{1}
}

func (x *AgentRes) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *AgentRes) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *AgentRes) GetChannel() string {
	if x != nil {
		return x.Channel
	}
	return ""
}

type AgentGroupByIDReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AgentGroupID string `protobuf:"bytes,1,opt,name=agentGroupID,proto3" json:"agentGroupID,omitempty"`
	OwnerID      string `protobuf:"bytes,2,opt,name=ownerID,proto3" json:"ownerID,omitempty"`
}

func (x *AgentGroupByIDReq) Reset() {
	*x = AgentGroupByIDReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fleet_pb_fleet_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AgentGroupByIDReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AgentGroupByIDReq) ProtoMessage() {}

func (x *AgentGroupByIDReq) ProtoReflect() protoreflect.Message {
	mi := &file_fleet_pb_fleet_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AgentGroupByIDReq.ProtoReflect.Descriptor instead.
func (*AgentGroupByIDReq) Descriptor() ([]byte, []int) {
	return file_fleet_pb_fleet_proto_rawDescGZIP(), []int{2}
}

func (x *AgentGroupByIDReq) GetAgentGroupID() string {
	if x != nil {
		return x.AgentGroupID
	}
	return ""
}

func (x *AgentGroupByIDReq) GetOwnerID() string {
	if x != nil {
		return x.OwnerID
	}
	return ""
}

type AgentGroupRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name    string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Channel string `protobuf:"bytes,3,opt,name=channel,proto3" json:"channel,omitempty"`
}

func (x *AgentGroupRes) Reset() {
	*x = AgentGroupRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fleet_pb_fleet_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AgentGroupRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AgentGroupRes) ProtoMessage() {}

func (x *AgentGroupRes) ProtoReflect() protoreflect.Message {
	mi := &file_fleet_pb_fleet_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AgentGroupRes.ProtoReflect.Descriptor instead.
func (*AgentGroupRes) Descriptor() ([]byte, []int) {
	return file_fleet_pb_fleet_proto_rawDescGZIP(), []int{3}
}

func (x *AgentGroupRes) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *AgentGroupRes) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *AgentGroupRes) GetChannel() string {
	if x != nil {
		return x.Channel
	}
	return ""
}

type AgentInfoByChannelIDReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Channel string `protobuf:"bytes,1,opt,name=channel,proto3" json:"channel,omitempty"`
}

func (x *AgentInfoByChannelIDReq) Reset() {
	*x = AgentInfoByChannelIDReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fleet_pb_fleet_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AgentInfoByChannelIDReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AgentInfoByChannelIDReq) ProtoMessage() {}

func (x *AgentInfoByChannelIDReq) ProtoReflect() protoreflect.Message {
	mi := &file_fleet_pb_fleet_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AgentInfoByChannelIDReq.ProtoReflect.Descriptor instead.
func (*AgentInfoByChannelIDReq) Descriptor() ([]byte, []int) {
	return file_fleet_pb_fleet_proto_rawDescGZIP(), []int{4}
}

func (x *AgentInfoByChannelIDReq) GetChannel() string {
	if x != nil {
		return x.Channel
	}
	return ""
}

type AgentInfoRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OwnerID   string            `protobuf:"bytes,1,opt,name=ownerID,proto3" json:"ownerID,omitempty"`
	AgentName string            `protobuf:"bytes,2,opt,name=agentName,proto3" json:"agentName,omitempty"`
	AgentTags map[string]string `protobuf:"bytes,3,rep,name=agentTags,proto3" json:"agentTags,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	OrbTags   map[string]string `protobuf:"bytes,4,rep,name=orbTags,proto3" json:"orbTags,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *AgentInfoRes) Reset() {
	*x = AgentInfoRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fleet_pb_fleet_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AgentInfoRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AgentInfoRes) ProtoMessage() {}

func (x *AgentInfoRes) ProtoReflect() protoreflect.Message {
	mi := &file_fleet_pb_fleet_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AgentInfoRes.ProtoReflect.Descriptor instead.
func (*AgentInfoRes) Descriptor() ([]byte, []int) {
	return file_fleet_pb_fleet_proto_rawDescGZIP(), []int{5}
}

func (x *AgentInfoRes) GetOwnerID() string {
	if x != nil {
		return x.OwnerID
	}
	return ""
}

func (x *AgentInfoRes) GetAgentName() string {
	if x != nil {
		return x.AgentName
	}
	return ""
}

func (x *AgentInfoRes) GetAgentTags() map[string]string {
	if x != nil {
		return x.AgentTags
	}
	return nil
}

func (x *AgentInfoRes) GetOrbTags() map[string]string {
	if x != nil {
		return x.OrbTags
	}
	return nil
}

var File_fleet_pb_fleet_proto protoreflect.FileDescriptor

var file_fleet_pb_fleet_proto_rawDesc = []byte{
	0x0a, 0x14, 0x66, 0x6c, 0x65, 0x65, 0x74, 0x2f, 0x70, 0x62, 0x2f, 0x66, 0x6c, 0x65, 0x65, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x66, 0x6c, 0x65, 0x65, 0x74, 0x22, 0x42, 0x0a,
	0x0c, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x71, 0x12, 0x18, 0x0a,
	0x07, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x61, 0x67, 0x65, 0x6e, 0x74, 0x49, 0x44, 0x12, 0x18, 0x0a, 0x07, 0x6f, 0x77, 0x6e, 0x65, 0x72,
	0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x49,
	0x44, 0x22, 0x48, 0x0a, 0x08, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x22, 0x51, 0x0a, 0x11, 0x41,
	0x67, 0x65, 0x6e, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x71,
	0x12, 0x22, 0x0a, 0x0c, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x44,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x49, 0x44, 0x12, 0x18, 0x0a, 0x07, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x49, 0x44, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x49, 0x44, 0x22, 0x4d,
	0x0a, 0x0d, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x73, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x22, 0x33, 0x0a,
	0x17, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x42, 0x79, 0x43, 0x68, 0x61, 0x6e,
	0x6e, 0x65, 0x6c, 0x49, 0x44, 0x52, 0x65, 0x71, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x68, 0x61, 0x6e,
	0x6e, 0x65, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x68, 0x61, 0x6e, 0x6e,
	0x65, 0x6c, 0x22, 0xbe, 0x02, 0x0a, 0x0c, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f,
	0x52, 0x65, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x49, 0x44, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x49, 0x44, 0x12, 0x1c, 0x0a,
	0x09, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x40, 0x0a, 0x09, 0x61,
	0x67, 0x65, 0x6e, 0x74, 0x54, 0x61, 0x67, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x22,
	0x2e, 0x66, 0x6c, 0x65, 0x65, 0x74, 0x2e, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f,
	0x52, 0x65, 0x73, 0x2e, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x54, 0x61, 0x67, 0x73, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x52, 0x09, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x54, 0x61, 0x67, 0x73, 0x12, 0x3a, 0x0a,
	0x07, 0x6f, 0x72, 0x62, 0x54, 0x61, 0x67, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20,
	0x2e, 0x66, 0x6c, 0x65, 0x65, 0x74, 0x2e, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f,
	0x52, 0x65, 0x73, 0x2e, 0x4f, 0x72, 0x62, 0x54, 0x61, 0x67, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x52, 0x07, 0x6f, 0x72, 0x62, 0x54, 0x61, 0x67, 0x73, 0x1a, 0x3c, 0x0a, 0x0e, 0x41, 0x67, 0x65,
	0x6e, 0x74, 0x54, 0x61, 0x67, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b,
	0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x3a, 0x0a, 0x0c, 0x4f, 0x72, 0x62, 0x54, 0x61,
	0x67, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a,
	0x02, 0x38, 0x01, 0x32, 0xe6, 0x01, 0x0a, 0x0c, 0x46, 0x6c, 0x65, 0x65, 0x74, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x37, 0x0a, 0x0d, 0x52, 0x65, 0x74, 0x72, 0x69, 0x65, 0x76, 0x65,
	0x41, 0x67, 0x65, 0x6e, 0x74, 0x12, 0x13, 0x2e, 0x66, 0x6c, 0x65, 0x65, 0x74, 0x2e, 0x41, 0x67,
	0x65, 0x6e, 0x74, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x71, 0x1a, 0x0f, 0x2e, 0x66, 0x6c, 0x65,
	0x65, 0x74, 0x2e, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x22, 0x00, 0x12, 0x46, 0x0a,
	0x12, 0x52, 0x65, 0x74, 0x72, 0x69, 0x65, 0x76, 0x65, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x47, 0x72,
	0x6f, 0x75, 0x70, 0x12, 0x18, 0x2e, 0x66, 0x6c, 0x65, 0x65, 0x74, 0x2e, 0x41, 0x67, 0x65, 0x6e,
	0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x71, 0x1a, 0x14, 0x2e,
	0x66, 0x6c, 0x65, 0x65, 0x74, 0x2e, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0x52, 0x65, 0x73, 0x22, 0x00, 0x12, 0x55, 0x0a, 0x1c, 0x52, 0x65, 0x74, 0x72, 0x69, 0x65, 0x76,
	0x65, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x42, 0x79, 0x43, 0x68, 0x61, 0x6e,
	0x6e, 0x65, 0x6c, 0x49, 0x44, 0x12, 0x1e, 0x2e, 0x66, 0x6c, 0x65, 0x65, 0x74, 0x2e, 0x41, 0x67,
	0x65, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x42, 0x79, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c,
	0x49, 0x44, 0x52, 0x65, 0x71, 0x1a, 0x13, 0x2e, 0x66, 0x6c, 0x65, 0x65, 0x74, 0x2e, 0x41, 0x67,
	0x65, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x22, 0x00, 0x42, 0x0a, 0x5a, 0x08,
	0x66, 0x6c, 0x65, 0x65, 0x74, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_fleet_pb_fleet_proto_rawDescOnce sync.Once
	file_fleet_pb_fleet_proto_rawDescData = file_fleet_pb_fleet_proto_rawDesc
)

func file_fleet_pb_fleet_proto_rawDescGZIP() []byte {
	file_fleet_pb_fleet_proto_rawDescOnce.Do(func() {
		file_fleet_pb_fleet_proto_rawDescData = protoimpl.X.CompressGZIP(file_fleet_pb_fleet_proto_rawDescData)
	})
	return file_fleet_pb_fleet_proto_rawDescData
}

var file_fleet_pb_fleet_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_fleet_pb_fleet_proto_goTypes = []interface{}{
	(*AgentByIDReq)(nil),            // 0: fleet.AgentByIDReq
	(*AgentRes)(nil),                // 1: fleet.AgentRes
	(*AgentGroupByIDReq)(nil),       // 2: fleet.AgentGroupByIDReq
	(*AgentGroupRes)(nil),           // 3: fleet.AgentGroupRes
	(*AgentInfoByChannelIDReq)(nil), // 4: fleet.AgentInfoByChannelIDReq
	(*AgentInfoRes)(nil),            // 5: fleet.AgentInfoRes
	nil,                             // 6: fleet.AgentInfoRes.AgentTagsEntry
	nil,                             // 7: fleet.AgentInfoRes.OrbTagsEntry
}
var file_fleet_pb_fleet_proto_depIdxs = []int32{
	6, // 0: fleet.AgentInfoRes.agentTags:type_name -> fleet.AgentInfoRes.AgentTagsEntry
	7, // 1: fleet.AgentInfoRes.orbTags:type_name -> fleet.AgentInfoRes.OrbTagsEntry
	0, // 2: fleet.FleetService.RetrieveAgent:input_type -> fleet.AgentByIDReq
	2, // 3: fleet.FleetService.RetrieveAgentGroup:input_type -> fleet.AgentGroupByIDReq
	4, // 4: fleet.FleetService.RetrieveAgentInfoByChannelID:input_type -> fleet.AgentInfoByChannelIDReq
	1, // 5: fleet.FleetService.RetrieveAgent:output_type -> fleet.AgentRes
	3, // 6: fleet.FleetService.RetrieveAgentGroup:output_type -> fleet.AgentGroupRes
	5, // 7: fleet.FleetService.RetrieveAgentInfoByChannelID:output_type -> fleet.AgentInfoRes
	5, // [5:8] is the sub-list for method output_type
	2, // [2:5] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_fleet_pb_fleet_proto_init() }
func file_fleet_pb_fleet_proto_init() {
	if File_fleet_pb_fleet_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_fleet_pb_fleet_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AgentByIDReq); i {
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
		file_fleet_pb_fleet_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AgentRes); i {
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
		file_fleet_pb_fleet_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AgentGroupByIDReq); i {
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
		file_fleet_pb_fleet_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AgentGroupRes); i {
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
		file_fleet_pb_fleet_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AgentInfoByChannelIDReq); i {
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
		file_fleet_pb_fleet_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AgentInfoRes); i {
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
			RawDescriptor: file_fleet_pb_fleet_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_fleet_pb_fleet_proto_goTypes,
		DependencyIndexes: file_fleet_pb_fleet_proto_depIdxs,
		MessageInfos:      file_fleet_pb_fleet_proto_msgTypes,
	}.Build()
	File_fleet_pb_fleet_proto = out.File
	file_fleet_pb_fleet_proto_rawDesc = nil
	file_fleet_pb_fleet_proto_goTypes = nil
	file_fleet_pb_fleet_proto_depIdxs = nil
}
