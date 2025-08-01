// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        (unknown)
// source: api/notification/v1/notification_types.proto

package notification

import (
	_ "buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
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

// Platform options for user devices
type Platform int32

const (
	// Default unspecified value
	Platform_PLATFORM_UNSPECIFIED Platform = 0
	// Web platform
	Platform_PLATFORM_WEB Platform = 1
	// iOS platform
	Platform_PLATFORM_IOS Platform = 2
	// Android platform
	Platform_PLATFORM_ANDROID Platform = 3
)

// Enum value maps for Platform.
var (
	Platform_name = map[int32]string{
		0: "PLATFORM_UNSPECIFIED",
		1: "PLATFORM_WEB",
		2: "PLATFORM_IOS",
		3: "PLATFORM_ANDROID",
	}
	Platform_value = map[string]int32{
		"PLATFORM_UNSPECIFIED": 0,
		"PLATFORM_WEB":         1,
		"PLATFORM_IOS":         2,
		"PLATFORM_ANDROID":     3,
	}
)

func (x Platform) Enum() *Platform {
	p := new(Platform)
	*p = x
	return p
}

func (x Platform) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Platform) Descriptor() protoreflect.EnumDescriptor {
	return file_api_notification_v1_notification_types_proto_enumTypes[0].Descriptor()
}

func (Platform) Type() protoreflect.EnumType {
	return &file_api_notification_v1_notification_types_proto_enumTypes[0]
}

func (x Platform) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Platform.Descriptor instead.
func (Platform) EnumDescriptor() ([]byte, []int) {
	return file_api_notification_v1_notification_types_proto_rawDescGZIP(), []int{0}
}

// DeviceRegistration represents a registered device for notifications
type DeviceRegistration struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Unique identifier for the registration
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// User ID associated with this device
	UserId string `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	// Device information
	DeviceData *UserDeviceData `protobuf:"bytes,3,opt,name=device_data,json=deviceData,proto3" json:"device_data,omitempty"`
	// When the device was registered
	RegisteredAt *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=registered_at,json=registeredAt,proto3" json:"registered_at,omitempty"`
}

func (x *DeviceRegistration) Reset() {
	*x = DeviceRegistration{}
	mi := &file_api_notification_v1_notification_types_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeviceRegistration) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeviceRegistration) ProtoMessage() {}

func (x *DeviceRegistration) ProtoReflect() protoreflect.Message {
	mi := &file_api_notification_v1_notification_types_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeviceRegistration.ProtoReflect.Descriptor instead.
func (*DeviceRegistration) Descriptor() ([]byte, []int) {
	return file_api_notification_v1_notification_types_proto_rawDescGZIP(), []int{0}
}

func (x *DeviceRegistration) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *DeviceRegistration) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *DeviceRegistration) GetDeviceData() *UserDeviceData {
	if x != nil {
		return x.DeviceData
	}
	return nil
}

func (x *DeviceRegistration) GetRegisteredAt() *timestamppb.Timestamp {
	if x != nil {
		return x.RegisteredAt
	}
	return nil
}

// UserDeviceData contains information about the user's device for authentication and notifications
type UserDeviceData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Unique identifier for the user's device
	DeviceId string `protobuf:"bytes,1,opt,name=device_id,json=deviceId,proto3" json:"device_id,omitempty"`
	// Platform type (Web, iOS, Android)
	Platform Platform `protobuf:"varint,2,opt,name=platform,proto3,enum=api.notification.v1.Platform" json:"platform,omitempty"`
	// Only one of app_version or browser_version should be set based on the platform
	//
	// Types that are assignable to Version:
	//
	//	*UserDeviceData_AppVersion
	//	*UserDeviceData_BrowserVersion
	Version isUserDeviceData_Version `protobuf_oneof:"version"`
	// User's IP address
	IpAddress string `protobuf:"bytes,5,opt,name=ip_address,json=ipAddress,proto3" json:"ip_address,omitempty"`
	// User agent string from the device
	UserAgent string `protobuf:"bytes,6,opt,name=user_agent,json=userAgent,proto3" json:"user_agent,omitempty"`
}

func (x *UserDeviceData) Reset() {
	*x = UserDeviceData{}
	mi := &file_api_notification_v1_notification_types_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UserDeviceData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserDeviceData) ProtoMessage() {}

func (x *UserDeviceData) ProtoReflect() protoreflect.Message {
	mi := &file_api_notification_v1_notification_types_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserDeviceData.ProtoReflect.Descriptor instead.
func (*UserDeviceData) Descriptor() ([]byte, []int) {
	return file_api_notification_v1_notification_types_proto_rawDescGZIP(), []int{1}
}

func (x *UserDeviceData) GetDeviceId() string {
	if x != nil {
		return x.DeviceId
	}
	return ""
}

func (x *UserDeviceData) GetPlatform() Platform {
	if x != nil {
		return x.Platform
	}
	return Platform_PLATFORM_UNSPECIFIED
}

func (m *UserDeviceData) GetVersion() isUserDeviceData_Version {
	if m != nil {
		return m.Version
	}
	return nil
}

func (x *UserDeviceData) GetAppVersion() string {
	if x, ok := x.GetVersion().(*UserDeviceData_AppVersion); ok {
		return x.AppVersion
	}
	return ""
}

func (x *UserDeviceData) GetBrowserVersion() string {
	if x, ok := x.GetVersion().(*UserDeviceData_BrowserVersion); ok {
		return x.BrowserVersion
	}
	return ""
}

func (x *UserDeviceData) GetIpAddress() string {
	if x != nil {
		return x.IpAddress
	}
	return ""
}

func (x *UserDeviceData) GetUserAgent() string {
	if x != nil {
		return x.UserAgent
	}
	return ""
}

type isUserDeviceData_Version interface {
	isUserDeviceData_Version()
}

type UserDeviceData_AppVersion struct {
	// Mobile app version (for iOS/Android)
	AppVersion string `protobuf:"bytes,3,opt,name=app_version,json=appVersion,proto3,oneof"`
}

type UserDeviceData_BrowserVersion struct {
	// Browser version (for web platform)
	BrowserVersion string `protobuf:"bytes,4,opt,name=browser_version,json=browserVersion,proto3,oneof"`
}

func (*UserDeviceData_AppVersion) isUserDeviceData_Version() {}

func (*UserDeviceData_BrowserVersion) isUserDeviceData_Version() {}

var File_api_notification_v1_notification_types_proto protoreflect.FileDescriptor

var file_api_notification_v1_notification_types_proto_rawDesc = []byte{
	0x0a, 0x2c, 0x61, 0x70, 0x69, 0x2f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13,
	0x61, 0x70, 0x69, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x62, 0x75, 0x66, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61,
	0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xc4, 0x01, 0x0a, 0x12, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x67, 0x69,
	0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49,
	0x64, 0x12, 0x44, 0x0a, 0x0b, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x64, 0x61, 0x74, 0x61,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6e, 0x6f, 0x74,
	0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65,
	0x72, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x44, 0x61, 0x74, 0x61, 0x52, 0x0a, 0x64, 0x65, 0x76,
	0x69, 0x63, 0x65, 0x44, 0x61, 0x74, 0x61, 0x12, 0x3f, 0x0a, 0x0d, 0x72, 0x65, 0x67, 0x69, 0x73,
	0x74, 0x65, 0x72, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0c, 0x72, 0x65, 0x67, 0x69,
	0x73, 0x74, 0x65, 0x72, 0x65, 0x64, 0x41, 0x74, 0x22, 0xb4, 0x02, 0x0a, 0x0e, 0x55, 0x73, 0x65,
	0x72, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x44, 0x61, 0x74, 0x61, 0x12, 0x28, 0x0a, 0x09, 0x64,
	0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0b,
	0xba, 0x48, 0x08, 0xc8, 0x01, 0x01, 0x72, 0x03, 0xb0, 0x01, 0x01, 0x52, 0x08, 0x64, 0x65, 0x76,
	0x69, 0x63, 0x65, 0x49, 0x64, 0x12, 0x46, 0x0a, 0x08, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72,
	0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6e, 0x6f,
	0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x6c,
	0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x42, 0x0b, 0xba, 0x48, 0x08, 0xc8, 0x01, 0x01, 0x82, 0x01,
	0x02, 0x10, 0x01, 0x52, 0x08, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x12, 0x21, 0x0a,
	0x0b, 0x61, 0x70, 0x70, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x48, 0x00, 0x52, 0x0a, 0x61, 0x70, 0x70, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x12, 0x29, 0x0a, 0x0f, 0x62, 0x72, 0x6f, 0x77, 0x73, 0x65, 0x72, 0x5f, 0x76, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0e, 0x62, 0x72, 0x6f,
	0x77, 0x73, 0x65, 0x72, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x29, 0x0a, 0x0a, 0x69,
	0x70, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x0a, 0xba, 0x48, 0x07, 0xc8, 0x01, 0x01, 0x72, 0x02, 0x70, 0x01, 0x52, 0x09, 0x69, 0x70, 0x41,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x25, 0x0a, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x61,
	0x67, 0x65, 0x6e, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x42, 0x06, 0xba, 0x48, 0x03, 0xc8,
	0x01, 0x01, 0x52, 0x09, 0x75, 0x73, 0x65, 0x72, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x42, 0x10, 0x0a,
	0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x05, 0xba, 0x48, 0x02, 0x08, 0x01, 0x2a,
	0x5e, 0x0a, 0x08, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x12, 0x18, 0x0a, 0x14, 0x50,
	0x4c, 0x41, 0x54, 0x46, 0x4f, 0x52, 0x4d, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46,
	0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x10, 0x0a, 0x0c, 0x50, 0x4c, 0x41, 0x54, 0x46, 0x4f, 0x52,
	0x4d, 0x5f, 0x57, 0x45, 0x42, 0x10, 0x01, 0x12, 0x10, 0x0a, 0x0c, 0x50, 0x4c, 0x41, 0x54, 0x46,
	0x4f, 0x52, 0x4d, 0x5f, 0x49, 0x4f, 0x53, 0x10, 0x02, 0x12, 0x14, 0x0a, 0x10, 0x50, 0x4c, 0x41,
	0x54, 0x46, 0x4f, 0x52, 0x4d, 0x5f, 0x41, 0x4e, 0x44, 0x52, 0x4f, 0x49, 0x44, 0x10, 0x03, 0x42,
	0xec, 0x01, 0x0a, 0x17, 0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6e, 0x6f, 0x74, 0x69,
	0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x42, 0x16, 0x4e, 0x6f, 0x74,
	0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x73, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x4b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x72, 0x73, 0x68, 0x65, 0x6c, 0x65, 0x6b, 0x68, 0x6f, 0x76, 0x2f, 0x6d, 0x73, 0x61,
	0x2d, 0x6d, 0x65, 0x73, 0x73, 0x65, 0x6e, 0x67, 0x65, 0x72, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x73, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x3b, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0xa2, 0x02, 0x03, 0x41, 0x4e, 0x58, 0xaa, 0x02, 0x13, 0x41, 0x70, 0x69, 0x2e, 0x4e,
	0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x56, 0x31, 0xca, 0x02,
	0x13, 0x41, 0x70, 0x69, 0x5c, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x1f, 0x41, 0x70, 0x69, 0x5c, 0x4e, 0x6f, 0x74, 0x69, 0x66,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x15, 0x41, 0x70, 0x69, 0x3a, 0x3a, 0x4e, 0x6f,
	0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_notification_v1_notification_types_proto_rawDescOnce sync.Once
	file_api_notification_v1_notification_types_proto_rawDescData = file_api_notification_v1_notification_types_proto_rawDesc
)

func file_api_notification_v1_notification_types_proto_rawDescGZIP() []byte {
	file_api_notification_v1_notification_types_proto_rawDescOnce.Do(func() {
		file_api_notification_v1_notification_types_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_notification_v1_notification_types_proto_rawDescData)
	})
	return file_api_notification_v1_notification_types_proto_rawDescData
}

var file_api_notification_v1_notification_types_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_api_notification_v1_notification_types_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_api_notification_v1_notification_types_proto_goTypes = []any{
	(Platform)(0),                 // 0: api.notification.v1.Platform
	(*DeviceRegistration)(nil),    // 1: api.notification.v1.DeviceRegistration
	(*UserDeviceData)(nil),        // 2: api.notification.v1.UserDeviceData
	(*timestamppb.Timestamp)(nil), // 3: google.protobuf.Timestamp
}
var file_api_notification_v1_notification_types_proto_depIdxs = []int32{
	2, // 0: api.notification.v1.DeviceRegistration.device_data:type_name -> api.notification.v1.UserDeviceData
	3, // 1: api.notification.v1.DeviceRegistration.registered_at:type_name -> google.protobuf.Timestamp
	0, // 2: api.notification.v1.UserDeviceData.platform:type_name -> api.notification.v1.Platform
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_api_notification_v1_notification_types_proto_init() }
func file_api_notification_v1_notification_types_proto_init() {
	if File_api_notification_v1_notification_types_proto != nil {
		return
	}
	file_api_notification_v1_notification_types_proto_msgTypes[1].OneofWrappers = []any{
		(*UserDeviceData_AppVersion)(nil),
		(*UserDeviceData_BrowserVersion)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_notification_v1_notification_types_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_notification_v1_notification_types_proto_goTypes,
		DependencyIndexes: file_api_notification_v1_notification_types_proto_depIdxs,
		EnumInfos:         file_api_notification_v1_notification_types_proto_enumTypes,
		MessageInfos:      file_api_notification_v1_notification_types_proto_msgTypes,
	}.Build()
	File_api_notification_v1_notification_types_proto = out.File
	file_api_notification_v1_notification_types_proto_rawDesc = nil
	file_api_notification_v1_notification_types_proto_goTypes = nil
	file_api_notification_v1_notification_types_proto_depIdxs = nil
}
