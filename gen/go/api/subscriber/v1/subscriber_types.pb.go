// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        (unknown)
// source: api/subscriber/v1/subscriber_types.proto

package subscriber

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// RequestDirection indicates whether invites are incoming to or outgoing from a user
type RequestDirection int32

const (
	// Default unspecified value
	RequestDirection_REQUEST_DIRECTION_UNSPECIFIED RequestDirection = 0
	// Invites received by the user
	RequestDirection_REQUEST_DIRECTION_INCOMING RequestDirection = 1
	// Invites sent by the user
	RequestDirection_REQUEST_DIRECTION_OUTGOING RequestDirection = 2
)

// Enum value maps for RequestDirection.
var (
	RequestDirection_name = map[int32]string{
		0: "REQUEST_DIRECTION_UNSPECIFIED",
		1: "REQUEST_DIRECTION_INCOMING",
		2: "REQUEST_DIRECTION_OUTGOING",
	}
	RequestDirection_value = map[string]int32{
		"REQUEST_DIRECTION_UNSPECIFIED": 0,
		"REQUEST_DIRECTION_INCOMING":    1,
		"REQUEST_DIRECTION_OUTGOING":    2,
	}
)

func (x RequestDirection) Enum() *RequestDirection {
	p := new(RequestDirection)
	*p = x
	return p
}

func (x RequestDirection) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RequestDirection) Descriptor() protoreflect.EnumDescriptor {
	return file_api_subscriber_v1_subscriber_types_proto_enumTypes[0].Descriptor()
}

func (RequestDirection) Type() protoreflect.EnumType {
	return &file_api_subscriber_v1_subscriber_types_proto_enumTypes[0]
}

func (x RequestDirection) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use RequestDirection.Descriptor instead.
func (RequestDirection) EnumDescriptor() ([]byte, []int) {
	return file_api_subscriber_v1_subscriber_types_proto_rawDescGZIP(), []int{0}
}

// RequestStatus represents the current state of a friend invite
type RequestStatus int32

const (
	// Default unspecified value
	RequestStatus_REQUEST_STATUS_UNSPECIFIED RequestStatus = 0
	// Invite is waiting for response
	RequestStatus_REQUEST_STATUS_PENDING RequestStatus = 1
	// Invite has been accepted
	RequestStatus_REQUEST_STATUS_ACCEPTED RequestStatus = 2
	// Invite has been declined
	RequestStatus_REQUEST_STATUS_DECLINED RequestStatus = 3
)

// Enum value maps for RequestStatus.
var (
	RequestStatus_name = map[int32]string{
		0: "REQUEST_STATUS_UNSPECIFIED",
		1: "REQUEST_STATUS_PENDING",
		2: "REQUEST_STATUS_ACCEPTED",
		3: "REQUEST_STATUS_DECLINED",
	}
	RequestStatus_value = map[string]int32{
		"REQUEST_STATUS_UNSPECIFIED": 0,
		"REQUEST_STATUS_PENDING":     1,
		"REQUEST_STATUS_ACCEPTED":    2,
		"REQUEST_STATUS_DECLINED":    3,
	}
)

func (x RequestStatus) Enum() *RequestStatus {
	p := new(RequestStatus)
	*p = x
	return p
}

func (x RequestStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RequestStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_api_subscriber_v1_subscriber_types_proto_enumTypes[1].Descriptor()
}

func (RequestStatus) Type() protoreflect.EnumType {
	return &file_api_subscriber_v1_subscriber_types_proto_enumTypes[1]
}

func (x RequestStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use RequestStatus.Descriptor instead.
func (RequestStatus) EnumDescriptor() ([]byte, []int) {
	return file_api_subscriber_v1_subscriber_types_proto_rawDescGZIP(), []int{1}
}

// FriendInvite contains details about a friend request between users
type FriendInvite struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Unique identifier for the invite
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// ID of the user who sent the invite
	FromUserId string `protobuf:"bytes,2,opt,name=from_user_id,json=fromUserId,proto3" json:"from_user_id,omitempty"`
	// ID of the user who received the invite
	ToUserId string `protobuf:"bytes,3,opt,name=to_user_id,json=toUserId,proto3" json:"to_user_id,omitempty"`
	// Optional message included with the invite
	Message string `protobuf:"bytes,4,opt,name=message,proto3" json:"message,omitempty"`
	// Current status of the invite
	Status RequestStatus `protobuf:"varint,5,opt,name=status,proto3,enum=api.subscriber.v1.RequestStatus" json:"status,omitempty"`
	// When the invite was created
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	// When the invite was last updated
	UpdatedAt *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *FriendInvite) Reset() {
	*x = FriendInvite{}
	mi := &file_api_subscriber_v1_subscriber_types_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FriendInvite) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FriendInvite) ProtoMessage() {}

func (x *FriendInvite) ProtoReflect() protoreflect.Message {
	mi := &file_api_subscriber_v1_subscriber_types_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FriendInvite.ProtoReflect.Descriptor instead.
func (*FriendInvite) Descriptor() ([]byte, []int) {
	return file_api_subscriber_v1_subscriber_types_proto_rawDescGZIP(), []int{0}
}

func (x *FriendInvite) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *FriendInvite) GetFromUserId() string {
	if x != nil {
		return x.FromUserId
	}
	return ""
}

func (x *FriendInvite) GetToUserId() string {
	if x != nil {
		return x.ToUserId
	}
	return ""
}

func (x *FriendInvite) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *FriendInvite) GetStatus() RequestStatus {
	if x != nil {
		return x.Status
	}
	return RequestStatus_REQUEST_STATUS_UNSPECIFIED
}

func (x *FriendInvite) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *FriendInvite) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

// Friend contains information about a user's friend
type Friend struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID of the friend
	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	// Display name of the friend
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// When the friendship was established
	FriendsSince *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=friends_since,json=friendsSince,proto3" json:"friends_since,omitempty"`
}

func (x *Friend) Reset() {
	*x = Friend{}
	mi := &file_api_subscriber_v1_subscriber_types_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Friend) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Friend) ProtoMessage() {}

func (x *Friend) ProtoReflect() protoreflect.Message {
	mi := &file_api_subscriber_v1_subscriber_types_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Friend.ProtoReflect.Descriptor instead.
func (*Friend) Descriptor() ([]byte, []int) {
	return file_api_subscriber_v1_subscriber_types_proto_rawDescGZIP(), []int{1}
}

func (x *Friend) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *Friend) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Friend) GetFriendsSince() *timestamppb.Timestamp {
	if x != nil {
		return x.FriendsSince
	}
	return nil
}

var File_api_subscriber_v1_subscriber_types_proto protoreflect.FileDescriptor

var file_api_subscriber_v1_subscriber_types_proto_rawDesc = []byte{
	0x0a, 0x28, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x72,
	0x2f, 0x76, 0x31, 0x2f, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x72, 0x5f, 0x74,
	0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11, 0x61, 0x70, 0x69, 0x2e,
	0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa8,
	0x02, 0x0a, 0x0c, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x20, 0x0a, 0x0c, 0x66, 0x72, 0x6f, 0x6d, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x66, 0x72, 0x6f, 0x6d, 0x55, 0x73, 0x65, 0x72, 0x49,
	0x64, 0x12, 0x1c, 0x0a, 0x0a, 0x74, 0x6f, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x6f, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12,
	0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x38, 0x0a, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x20, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61,
	0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x39,
	0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x76, 0x0a, 0x06, 0x46, 0x72, 0x69,
	0x65, 0x6e, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x3f, 0x0a, 0x0d, 0x66, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x73, 0x5f, 0x73, 0x69, 0x6e, 0x63,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x0c, 0x66, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x73, 0x53, 0x69, 0x6e, 0x63,
	0x65, 0x2a, 0x75, 0x0a, 0x10, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x44, 0x69, 0x72, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x21, 0x0a, 0x1d, 0x52, 0x45, 0x51, 0x55, 0x45, 0x53, 0x54,
	0x5f, 0x44, 0x49, 0x52, 0x45, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45,
	0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x1e, 0x0a, 0x1a, 0x52, 0x45, 0x51, 0x55,
	0x45, 0x53, 0x54, 0x5f, 0x44, 0x49, 0x52, 0x45, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x49, 0x4e,
	0x43, 0x4f, 0x4d, 0x49, 0x4e, 0x47, 0x10, 0x01, 0x12, 0x1e, 0x0a, 0x1a, 0x52, 0x45, 0x51, 0x55,
	0x45, 0x53, 0x54, 0x5f, 0x44, 0x49, 0x52, 0x45, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x4f, 0x55,
	0x54, 0x47, 0x4f, 0x49, 0x4e, 0x47, 0x10, 0x02, 0x2a, 0x85, 0x01, 0x0a, 0x0d, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1e, 0x0a, 0x1a, 0x52, 0x45,
	0x51, 0x55, 0x45, 0x53, 0x54, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x55, 0x4e, 0x53,
	0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x1a, 0x0a, 0x16, 0x52, 0x45,
	0x51, 0x55, 0x45, 0x53, 0x54, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x50, 0x45, 0x4e,
	0x44, 0x49, 0x4e, 0x47, 0x10, 0x01, 0x12, 0x1b, 0x0a, 0x17, 0x52, 0x45, 0x51, 0x55, 0x45, 0x53,
	0x54, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x41, 0x43, 0x43, 0x45, 0x50, 0x54, 0x45,
	0x44, 0x10, 0x02, 0x12, 0x1b, 0x0a, 0x17, 0x52, 0x45, 0x51, 0x55, 0x45, 0x53, 0x54, 0x5f, 0x53,
	0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x44, 0x45, 0x43, 0x4c, 0x49, 0x4e, 0x45, 0x44, 0x10, 0x03,
	0x42, 0xdc, 0x01, 0x0a, 0x15, 0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x75, 0x62,
	0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x42, 0x14, 0x53, 0x75, 0x62, 0x73,
	0x63, 0x72, 0x69, 0x62, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x50, 0x01, 0x5a, 0x47, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x72,
	0x73, 0x68, 0x65, 0x6c, 0x65, 0x6b, 0x68, 0x6f, 0x76, 0x2f, 0x6d, 0x73, 0x61, 0x2d, 0x6d, 0x65,
	0x73, 0x73, 0x65, 0x6e, 0x67, 0x65, 0x72, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x72, 0x2f, 0x76, 0x31,
	0x3b, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x72, 0xa2, 0x02, 0x03, 0x41, 0x53,
	0x58, 0xaa, 0x02, 0x11, 0x41, 0x70, 0x69, 0x2e, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62,
	0x65, 0x72, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x11, 0x41, 0x70, 0x69, 0x5c, 0x53, 0x75, 0x62, 0x73,
	0x63, 0x72, 0x69, 0x62, 0x65, 0x72, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x1d, 0x41, 0x70, 0x69, 0x5c,
	0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x72, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50,
	0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x13, 0x41, 0x70, 0x69, 0x3a,
	0x3a, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x72, 0x3a, 0x3a, 0x56, 0x31, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_subscriber_v1_subscriber_types_proto_rawDescOnce sync.Once
	file_api_subscriber_v1_subscriber_types_proto_rawDescData = file_api_subscriber_v1_subscriber_types_proto_rawDesc
)

func file_api_subscriber_v1_subscriber_types_proto_rawDescGZIP() []byte {
	file_api_subscriber_v1_subscriber_types_proto_rawDescOnce.Do(func() {
		file_api_subscriber_v1_subscriber_types_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_subscriber_v1_subscriber_types_proto_rawDescData)
	})
	return file_api_subscriber_v1_subscriber_types_proto_rawDescData
}

var file_api_subscriber_v1_subscriber_types_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_api_subscriber_v1_subscriber_types_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_api_subscriber_v1_subscriber_types_proto_goTypes = []any{
	(RequestDirection)(0),         // 0: api.subscriber.v1.RequestDirection
	(RequestStatus)(0),            // 1: api.subscriber.v1.RequestStatus
	(*FriendInvite)(nil),          // 2: api.subscriber.v1.FriendInvite
	(*Friend)(nil),                // 3: api.subscriber.v1.Friend
	(*timestamppb.Timestamp)(nil), // 4: google.protobuf.Timestamp
}
var file_api_subscriber_v1_subscriber_types_proto_depIdxs = []int32{
	1, // 0: api.subscriber.v1.FriendInvite.status:type_name -> api.subscriber.v1.RequestStatus
	4, // 1: api.subscriber.v1.FriendInvite.created_at:type_name -> google.protobuf.Timestamp
	4, // 2: api.subscriber.v1.FriendInvite.updated_at:type_name -> google.protobuf.Timestamp
	4, // 3: api.subscriber.v1.Friend.friends_since:type_name -> google.protobuf.Timestamp
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_api_subscriber_v1_subscriber_types_proto_init() }
func file_api_subscriber_v1_subscriber_types_proto_init() {
	if File_api_subscriber_v1_subscriber_types_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_subscriber_v1_subscriber_types_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_subscriber_v1_subscriber_types_proto_goTypes,
		DependencyIndexes: file_api_subscriber_v1_subscriber_types_proto_depIdxs,
		EnumInfos:         file_api_subscriber_v1_subscriber_types_proto_enumTypes,
		MessageInfos:      file_api_subscriber_v1_subscriber_types_proto_msgTypes,
	}.Build()
	File_api_subscriber_v1_subscriber_types_proto = out.File
	file_api_subscriber_v1_subscriber_types_proto_rawDesc = nil
	file_api_subscriber_v1_subscriber_types_proto_goTypes = nil
	file_api_subscriber_v1_subscriber_types_proto_depIdxs = nil
}
