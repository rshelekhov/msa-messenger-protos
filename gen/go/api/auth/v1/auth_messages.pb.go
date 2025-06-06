// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        (unknown)
// source: api/auth/v1/auth_messages.proto

package auth

import (
	_ "buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
	v1 "github.com/rshelekhov/msa-messenger-protos/gen/go/api/common/v1"
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

// LoginRequest contains credentials and device info for authentication
type LoginRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// User's email address
	Email string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	// User's password
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	// Device data for tracking user activity and sending notifications
	DeviceData *v1.UserDeviceData `protobuf:"bytes,3,opt,name=device_data,json=deviceData,proto3" json:"device_data,omitempty"`
}

func (x *LoginRequest) Reset() {
	*x = LoginRequest{}
	mi := &file_api_auth_v1_auth_messages_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LoginRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginRequest) ProtoMessage() {}

func (x *LoginRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_auth_v1_auth_messages_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginRequest.ProtoReflect.Descriptor instead.
func (*LoginRequest) Descriptor() ([]byte, []int) {
	return file_api_auth_v1_auth_messages_proto_rawDescGZIP(), []int{0}
}

func (x *LoginRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *LoginRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *LoginRequest) GetDeviceData() *v1.UserDeviceData {
	if x != nil {
		return x.DeviceData
	}
	return nil
}

// LoginResponse contains authentication tokens returned after successful login
type LoginResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Authentication tokens
	Tokens *Token `protobuf:"bytes,1,opt,name=tokens,proto3" json:"tokens,omitempty"`
}

func (x *LoginResponse) Reset() {
	*x = LoginResponse{}
	mi := &file_api_auth_v1_auth_messages_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LoginResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginResponse) ProtoMessage() {}

func (x *LoginResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_auth_v1_auth_messages_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginResponse.ProtoReflect.Descriptor instead.
func (*LoginResponse) Descriptor() ([]byte, []int) {
	return file_api_auth_v1_auth_messages_proto_rawDescGZIP(), []int{1}
}

func (x *LoginResponse) GetTokens() *Token {
	if x != nil {
		return x.Tokens
	}
	return nil
}

// RegisterUserRequest contains data needed to register a new user
type RegisterUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// User's email address
	Email string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	// User's password
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	// User's display name
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	// Device data for tracking user activity and sending notifications
	DeviceData *v1.UserDeviceData `protobuf:"bytes,4,opt,name=device_data,json=deviceData,proto3" json:"device_data,omitempty"`
}

func (x *RegisterUserRequest) Reset() {
	*x = RegisterUserRequest{}
	mi := &file_api_auth_v1_auth_messages_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RegisterUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterUserRequest) ProtoMessage() {}

func (x *RegisterUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_auth_v1_auth_messages_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterUserRequest.ProtoReflect.Descriptor instead.
func (*RegisterUserRequest) Descriptor() ([]byte, []int) {
	return file_api_auth_v1_auth_messages_proto_rawDescGZIP(), []int{2}
}

func (x *RegisterUserRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *RegisterUserRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *RegisterUserRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *RegisterUserRequest) GetDeviceData() *v1.UserDeviceData {
	if x != nil {
		return x.DeviceData
	}
	return nil
}

// RegisterUserResponse contains the newly registered user's ID
type RegisterUserResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Newly registered user's ID
	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	// Authentication tokens
	Tokens *Token `protobuf:"bytes,2,opt,name=tokens,proto3" json:"tokens,omitempty"`
}

func (x *RegisterUserResponse) Reset() {
	*x = RegisterUserResponse{}
	mi := &file_api_auth_v1_auth_messages_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RegisterUserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterUserResponse) ProtoMessage() {}

func (x *RegisterUserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_auth_v1_auth_messages_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterUserResponse.ProtoReflect.Descriptor instead.
func (*RegisterUserResponse) Descriptor() ([]byte, []int) {
	return file_api_auth_v1_auth_messages_proto_rawDescGZIP(), []int{3}
}

func (x *RegisterUserResponse) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *RegisterUserResponse) GetTokens() *Token {
	if x != nil {
		return x.Tokens
	}
	return nil
}

// RefreshTokenRequest contains data needed to refresh expired tokens
type RefreshTokenRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Valid refresh token
	RefreshToken string `protobuf:"bytes,1,opt,name=refresh_token,json=refreshToken,proto3" json:"refresh_token,omitempty"`
	// Information about the user's device
	DeviceData *v1.UserDeviceData `protobuf:"bytes,2,opt,name=device_data,json=deviceData,proto3" json:"device_data,omitempty"`
}

func (x *RefreshTokenRequest) Reset() {
	*x = RefreshTokenRequest{}
	mi := &file_api_auth_v1_auth_messages_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RefreshTokenRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RefreshTokenRequest) ProtoMessage() {}

func (x *RefreshTokenRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_auth_v1_auth_messages_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RefreshTokenRequest.ProtoReflect.Descriptor instead.
func (*RefreshTokenRequest) Descriptor() ([]byte, []int) {
	return file_api_auth_v1_auth_messages_proto_rawDescGZIP(), []int{4}
}

func (x *RefreshTokenRequest) GetRefreshToken() string {
	if x != nil {
		return x.RefreshToken
	}
	return ""
}

func (x *RefreshTokenRequest) GetDeviceData() *v1.UserDeviceData {
	if x != nil {
		return x.DeviceData
	}
	return nil
}

// RefreshTokenResponse contains newly issued authentication tokens
type RefreshTokenResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// New JWT token for API access
	AccessToken string `protobuf:"bytes,1,opt,name=access_token,json=accessToken,proto3" json:"access_token,omitempty"`
	// New refresh token
	RefreshToken string `protobuf:"bytes,2,opt,name=refresh_token,json=refreshToken,proto3" json:"refresh_token,omitempty"`
}

func (x *RefreshTokenResponse) Reset() {
	*x = RefreshTokenResponse{}
	mi := &file_api_auth_v1_auth_messages_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RefreshTokenResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RefreshTokenResponse) ProtoMessage() {}

func (x *RefreshTokenResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_auth_v1_auth_messages_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RefreshTokenResponse.ProtoReflect.Descriptor instead.
func (*RefreshTokenResponse) Descriptor() ([]byte, []int) {
	return file_api_auth_v1_auth_messages_proto_rawDescGZIP(), []int{5}
}

func (x *RefreshTokenResponse) GetAccessToken() string {
	if x != nil {
		return x.AccessToken
	}
	return ""
}

func (x *RefreshTokenResponse) GetRefreshToken() string {
	if x != nil {
		return x.RefreshToken
	}
	return ""
}

// UpdateEmailRequest contains data needed to update the user's email address
type UpdateEmailRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// User's ID
	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	// User's email address
	Email string `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
}

func (x *UpdateEmailRequest) Reset() {
	*x = UpdateEmailRequest{}
	mi := &file_api_auth_v1_auth_messages_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateEmailRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateEmailRequest) ProtoMessage() {}

func (x *UpdateEmailRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_auth_v1_auth_messages_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateEmailRequest.ProtoReflect.Descriptor instead.
func (*UpdateEmailRequest) Descriptor() ([]byte, []int) {
	return file_api_auth_v1_auth_messages_proto_rawDescGZIP(), []int{6}
}

func (x *UpdateEmailRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *UpdateEmailRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

// UpdateEmailResponse indicates if the email update operation was successful
type UpdateEmailResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Whether the email update operation succeeded
	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *UpdateEmailResponse) Reset() {
	*x = UpdateEmailResponse{}
	mi := &file_api_auth_v1_auth_messages_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateEmailResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateEmailResponse) ProtoMessage() {}

func (x *UpdateEmailResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_auth_v1_auth_messages_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateEmailResponse.ProtoReflect.Descriptor instead.
func (*UpdateEmailResponse) Descriptor() ([]byte, []int) {
	return file_api_auth_v1_auth_messages_proto_rawDescGZIP(), []int{7}
}

func (x *UpdateEmailResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

// LogoutRequest contains the device information for the session to terminate
type LogoutRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Information about the user's device
	DeviceData *v1.UserDeviceData `protobuf:"bytes,1,opt,name=device_data,json=deviceData,proto3" json:"device_data,omitempty"`
}

func (x *LogoutRequest) Reset() {
	*x = LogoutRequest{}
	mi := &file_api_auth_v1_auth_messages_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LogoutRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogoutRequest) ProtoMessage() {}

func (x *LogoutRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_auth_v1_auth_messages_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogoutRequest.ProtoReflect.Descriptor instead.
func (*LogoutRequest) Descriptor() ([]byte, []int) {
	return file_api_auth_v1_auth_messages_proto_rawDescGZIP(), []int{8}
}

func (x *LogoutRequest) GetDeviceData() *v1.UserDeviceData {
	if x != nil {
		return x.DeviceData
	}
	return nil
}

// LogoutResponse indicates if the logout operation was successful
type LogoutResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Whether the logout operation succeeded
	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *LogoutResponse) Reset() {
	*x = LogoutResponse{}
	mi := &file_api_auth_v1_auth_messages_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LogoutResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogoutResponse) ProtoMessage() {}

func (x *LogoutResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_auth_v1_auth_messages_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogoutResponse.ProtoReflect.Descriptor instead.
func (*LogoutResponse) Descriptor() ([]byte, []int) {
	return file_api_auth_v1_auth_messages_proto_rawDescGZIP(), []int{9}
}

func (x *LogoutResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

// RemoveUserRequest contains the device information for the session to terminate
type RemoveUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// User's ID
	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	// Information about the user's device
	DeviceData *v1.UserDeviceData `protobuf:"bytes,2,opt,name=device_data,json=deviceData,proto3" json:"device_data,omitempty"`
}

func (x *RemoveUserRequest) Reset() {
	*x = RemoveUserRequest{}
	mi := &file_api_auth_v1_auth_messages_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RemoveUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveUserRequest) ProtoMessage() {}

func (x *RemoveUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_auth_v1_auth_messages_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveUserRequest.ProtoReflect.Descriptor instead.
func (*RemoveUserRequest) Descriptor() ([]byte, []int) {
	return file_api_auth_v1_auth_messages_proto_rawDescGZIP(), []int{10}
}

func (x *RemoveUserRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *RemoveUserRequest) GetDeviceData() *v1.UserDeviceData {
	if x != nil {
		return x.DeviceData
	}
	return nil
}

// RemoveUserResponse indicates if the user removal operation was successful
type RemoveUserResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Whether the user removal operation succeeded
	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *RemoveUserResponse) Reset() {
	*x = RemoveUserResponse{}
	mi := &file_api_auth_v1_auth_messages_proto_msgTypes[11]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RemoveUserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveUserResponse) ProtoMessage() {}

func (x *RemoveUserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_auth_v1_auth_messages_proto_msgTypes[11]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveUserResponse.ProtoReflect.Descriptor instead.
func (*RemoveUserResponse) Descriptor() ([]byte, []int) {
	return file_api_auth_v1_auth_messages_proto_rawDescGZIP(), []int{11}
}

func (x *RemoveUserResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

// GetJWKSRequest is used to get the JWKS for the auth service
type GetJWKSRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetJWKSRequest) Reset() {
	*x = GetJWKSRequest{}
	mi := &file_api_auth_v1_auth_messages_proto_msgTypes[12]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetJWKSRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetJWKSRequest) ProtoMessage() {}

func (x *GetJWKSRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_auth_v1_auth_messages_proto_msgTypes[12]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetJWKSRequest.ProtoReflect.Descriptor instead.
func (*GetJWKSRequest) Descriptor() ([]byte, []int) {
	return file_api_auth_v1_auth_messages_proto_rawDescGZIP(), []int{12}
}

// GetJWKSResponse contains the JWKS for the auth service
type GetJWKSResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The JWKS for the auth service
	Jwks []*JWK `protobuf:"bytes,1,rep,name=jwks,proto3" json:"jwks,omitempty"`
}

func (x *GetJWKSResponse) Reset() {
	*x = GetJWKSResponse{}
	mi := &file_api_auth_v1_auth_messages_proto_msgTypes[13]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetJWKSResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetJWKSResponse) ProtoMessage() {}

func (x *GetJWKSResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_auth_v1_auth_messages_proto_msgTypes[13]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetJWKSResponse.ProtoReflect.Descriptor instead.
func (*GetJWKSResponse) Descriptor() ([]byte, []int) {
	return file_api_auth_v1_auth_messages_proto_rawDescGZIP(), []int{13}
}

func (x *GetJWKSResponse) GetJwks() []*JWK {
	if x != nil {
		return x.Jwks
	}
	return nil
}

// JWK represents a JSON Web Key structure containing the necessary fields
// for RSA public key construction
type JWK struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The specific cryptographic algorithm used with the key
	Alg string `protobuf:"bytes,1,opt,name=alg,proto3" json:"alg,omitempty"`
	// The family of cryptographic algorithms used with the key
	Kty string `protobuf:"bytes,2,opt,name=kty,proto3" json:"kty,omitempty"`
	// How the key was meant to be used; sig represents the signature
	Use string `protobuf:"bytes,3,opt,name=use,proto3" json:"use,omitempty"`
	// The unique identifier for the key
	Kid string `protobuf:"bytes,4,opt,name=kid,proto3" json:"kid,omitempty"`
	// The modulus for the RSA public key
	N string `protobuf:"bytes,5,opt,name=n,proto3" json:"n,omitempty"`
	// The exponent for the RSA public key
	E string `protobuf:"bytes,6,opt,name=e,proto3" json:"e,omitempty"`
}

func (x *JWK) Reset() {
	*x = JWK{}
	mi := &file_api_auth_v1_auth_messages_proto_msgTypes[14]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *JWK) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JWK) ProtoMessage() {}

func (x *JWK) ProtoReflect() protoreflect.Message {
	mi := &file_api_auth_v1_auth_messages_proto_msgTypes[14]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JWK.ProtoReflect.Descriptor instead.
func (*JWK) Descriptor() ([]byte, []int) {
	return file_api_auth_v1_auth_messages_proto_rawDescGZIP(), []int{14}
}

func (x *JWK) GetAlg() string {
	if x != nil {
		return x.Alg
	}
	return ""
}

func (x *JWK) GetKty() string {
	if x != nil {
		return x.Kty
	}
	return ""
}

func (x *JWK) GetUse() string {
	if x != nil {
		return x.Use
	}
	return ""
}

func (x *JWK) GetKid() string {
	if x != nil {
		return x.Kid
	}
	return ""
}

func (x *JWK) GetN() string {
	if x != nil {
		return x.N
	}
	return ""
}

func (x *JWK) GetE() string {
	if x != nil {
		return x.E
	}
	return ""
}

var File_api_auth_v1_auth_messages_proto protoreflect.FileDescriptor

var file_api_auth_v1_auth_messages_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x75,
	0x74, 0x68, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x0b, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x76, 0x31, 0x1a, 0x20,
	0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1c, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x75,
	0x74, 0x68, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b,
	0x62, 0x75, 0x66, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc1, 0x01, 0x0a, 0x0c,
	0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x25, 0x0a, 0x05,
	0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0f, 0xba, 0x48, 0x0c,
	0xc8, 0x01, 0x01, 0x72, 0x07, 0x10, 0x05, 0x18, 0xff, 0x01, 0x60, 0x01, 0x52, 0x05, 0x65, 0x6d,
	0x61, 0x69, 0x6c, 0x12, 0x42, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x26, 0xba, 0x48, 0x23, 0xc8, 0x01, 0x01, 0x72, 0x1e, 0x10,
	0x08, 0x18, 0xff, 0x01, 0x32, 0x17, 0x5e, 0x5b, 0x41, 0x2d, 0x5a, 0x61, 0x2d, 0x7a, 0x5c, 0x64,
	0x40, 0x24, 0x21, 0x25, 0x2a, 0x3f, 0x26, 0x5d, 0x7b, 0x38, 0x2c, 0x7d, 0x24, 0x52, 0x08, 0x70,
	0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x46, 0x0a, 0x0b, 0x64, 0x65, 0x76, 0x69, 0x63,
	0x65, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65,
	0x72, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x44, 0x61, 0x74, 0x61, 0x42, 0x06, 0xba, 0x48, 0x03,
	0xc8, 0x01, 0x01, 0x52, 0x0a, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x44, 0x61, 0x74, 0x61, 0x22,
	0x3b, 0x0a, 0x0d, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x2a, 0x0a, 0x06, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x12, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x06, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x22, 0xeb, 0x01, 0x0a,
	0x13, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x25, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x0f, 0xba, 0x48, 0x0c, 0xc8, 0x01, 0x01, 0x72, 0x07, 0x10, 0x05, 0x18,
	0xff, 0x01, 0x60, 0x01, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x42, 0x0a, 0x08, 0x70,
	0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x26, 0xba,
	0x48, 0x23, 0xc8, 0x01, 0x01, 0x72, 0x1e, 0x10, 0x08, 0x18, 0xff, 0x01, 0x32, 0x17, 0x5e, 0x5b,
	0x41, 0x2d, 0x5a, 0x61, 0x2d, 0x7a, 0x5c, 0x64, 0x40, 0x24, 0x21, 0x25, 0x2a, 0x3f, 0x26, 0x5d,
	0x7b, 0x38, 0x2c, 0x7d, 0x24, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12,
	0x21, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0d, 0xba,
	0x48, 0x0a, 0xc8, 0x01, 0x01, 0x72, 0x05, 0x10, 0x01, 0x18, 0xff, 0x01, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x46, 0x0a, 0x0b, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x64, 0x61, 0x74,
	0x61, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x44, 0x65, 0x76, 0x69,
	0x63, 0x65, 0x44, 0x61, 0x74, 0x61, 0x42, 0x06, 0xba, 0x48, 0x03, 0xc8, 0x01, 0x01, 0x52, 0x0a,
	0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x44, 0x61, 0x74, 0x61, 0x22, 0x5b, 0x0a, 0x14, 0x52, 0x65,
	0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x2a, 0x0a, 0x06, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52,
	0x06, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x22, 0x8a, 0x01, 0x0a, 0x13, 0x52, 0x65, 0x66, 0x72,
	0x65, 0x73, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x2b, 0x0a, 0x0d, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x06, 0xba, 0x48, 0x03, 0xc8, 0x01, 0x01, 0x52, 0x0c,
	0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x46, 0x0a, 0x0b,
	0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76,
	0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x44, 0x61, 0x74, 0x61,
	0x42, 0x06, 0xba, 0x48, 0x03, 0xc8, 0x01, 0x01, 0x52, 0x0a, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65,
	0x44, 0x61, 0x74, 0x61, 0x22, 0x5e, 0x0a, 0x14, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x21, 0x0a, 0x0c,
	0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12,
	0x23, 0x0a, 0x0d, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x61, 0x0a, 0x12, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x45, 0x6d,
	0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x24, 0x0a, 0x07, 0x75, 0x73,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0b, 0xba, 0x48, 0x08,
	0xc8, 0x01, 0x01, 0x72, 0x03, 0xb0, 0x01, 0x01, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64,
	0x12, 0x25, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x0f, 0xba, 0x48, 0x0c, 0xc8, 0x01, 0x01, 0x72, 0x07, 0x10, 0x05, 0x18, 0xff, 0x01, 0x60, 0x01,
	0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x22, 0x2f, 0x0a, 0x13, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x22, 0x57, 0x0a, 0x0d, 0x4c, 0x6f, 0x67, 0x6f,
	0x75, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x46, 0x0a, 0x0b, 0x64, 0x65, 0x76,
	0x69, 0x63, 0x65, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x55,
	0x73, 0x65, 0x72, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x44, 0x61, 0x74, 0x61, 0x42, 0x06, 0xba,
	0x48, 0x03, 0xc8, 0x01, 0x01, 0x52, 0x0a, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x44, 0x61, 0x74,
	0x61, 0x22, 0x2a, 0x0a, 0x0e, 0x4c, 0x6f, 0x67, 0x6f, 0x75, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x22, 0x81, 0x01,
	0x0a, 0x11, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x24, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x0b, 0xba, 0x48, 0x08, 0xc8, 0x01, 0x01, 0x72, 0x03, 0xb0, 0x01,
	0x01, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x46, 0x0a, 0x0b, 0x64, 0x65, 0x76,
	0x69, 0x63, 0x65, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x55,
	0x73, 0x65, 0x72, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x44, 0x61, 0x74, 0x61, 0x42, 0x06, 0xba,
	0x48, 0x03, 0xc8, 0x01, 0x01, 0x52, 0x0a, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x44, 0x61, 0x74,
	0x61, 0x22, 0x2e, 0x0a, 0x12, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x22, 0x10, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x4a, 0x57, 0x4b, 0x53, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x22, 0x37, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x4a, 0x57, 0x4b, 0x53, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x24, 0x0a, 0x04, 0x6a, 0x77, 0x6b, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e,
	0x76, 0x31, 0x2e, 0x4a, 0x57, 0x4b, 0x52, 0x04, 0x6a, 0x77, 0x6b, 0x73, 0x22, 0x69, 0x0a, 0x03,
	0x4a, 0x57, 0x4b, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x6c, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x61, 0x6c, 0x67, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6b, 0x74, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x73, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x69, 0x64,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x69, 0x64, 0x12, 0x0c, 0x0a, 0x01, 0x6e,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x01, 0x6e, 0x12, 0x0c, 0x0a, 0x01, 0x65, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x01, 0x65, 0x42, 0xaf, 0x01, 0x0a, 0x0f, 0x63, 0x6f, 0x6d, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x76, 0x31, 0x42, 0x11, 0x41, 0x75, 0x74,
	0x68, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01,
	0x5a, 0x3b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x72, 0x73, 0x68,
	0x65, 0x6c, 0x65, 0x6b, 0x68, 0x6f, 0x76, 0x2f, 0x6d, 0x73, 0x61, 0x2d, 0x6d, 0x65, 0x73, 0x73,
	0x65, 0x6e, 0x67, 0x65, 0x72, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x76, 0x31, 0x3b, 0x61, 0x75, 0x74, 0x68, 0xa2, 0x02, 0x03,
	0x41, 0x41, 0x58, 0xaa, 0x02, 0x0b, 0x41, 0x70, 0x69, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x2e, 0x56,
	0x31, 0xca, 0x02, 0x0b, 0x41, 0x70, 0x69, 0x5c, 0x41, 0x75, 0x74, 0x68, 0x5c, 0x56, 0x31, 0xe2,
	0x02, 0x17, 0x41, 0x70, 0x69, 0x5c, 0x41, 0x75, 0x74, 0x68, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50,
	0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0d, 0x41, 0x70, 0x69, 0x3a,
	0x3a, 0x41, 0x75, 0x74, 0x68, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_api_auth_v1_auth_messages_proto_rawDescOnce sync.Once
	file_api_auth_v1_auth_messages_proto_rawDescData = file_api_auth_v1_auth_messages_proto_rawDesc
)

func file_api_auth_v1_auth_messages_proto_rawDescGZIP() []byte {
	file_api_auth_v1_auth_messages_proto_rawDescOnce.Do(func() {
		file_api_auth_v1_auth_messages_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_auth_v1_auth_messages_proto_rawDescData)
	})
	return file_api_auth_v1_auth_messages_proto_rawDescData
}

var file_api_auth_v1_auth_messages_proto_msgTypes = make([]protoimpl.MessageInfo, 15)
var file_api_auth_v1_auth_messages_proto_goTypes = []any{
	(*LoginRequest)(nil),         // 0: api.auth.v1.LoginRequest
	(*LoginResponse)(nil),        // 1: api.auth.v1.LoginResponse
	(*RegisterUserRequest)(nil),  // 2: api.auth.v1.RegisterUserRequest
	(*RegisterUserResponse)(nil), // 3: api.auth.v1.RegisterUserResponse
	(*RefreshTokenRequest)(nil),  // 4: api.auth.v1.RefreshTokenRequest
	(*RefreshTokenResponse)(nil), // 5: api.auth.v1.RefreshTokenResponse
	(*UpdateEmailRequest)(nil),   // 6: api.auth.v1.UpdateEmailRequest
	(*UpdateEmailResponse)(nil),  // 7: api.auth.v1.UpdateEmailResponse
	(*LogoutRequest)(nil),        // 8: api.auth.v1.LogoutRequest
	(*LogoutResponse)(nil),       // 9: api.auth.v1.LogoutResponse
	(*RemoveUserRequest)(nil),    // 10: api.auth.v1.RemoveUserRequest
	(*RemoveUserResponse)(nil),   // 11: api.auth.v1.RemoveUserResponse
	(*GetJWKSRequest)(nil),       // 12: api.auth.v1.GetJWKSRequest
	(*GetJWKSResponse)(nil),      // 13: api.auth.v1.GetJWKSResponse
	(*JWK)(nil),                  // 14: api.auth.v1.JWK
	(*v1.UserDeviceData)(nil),    // 15: api.common.v1.UserDeviceData
	(*Token)(nil),                // 16: api.auth.v1.Token
}
var file_api_auth_v1_auth_messages_proto_depIdxs = []int32{
	15, // 0: api.auth.v1.LoginRequest.device_data:type_name -> api.common.v1.UserDeviceData
	16, // 1: api.auth.v1.LoginResponse.tokens:type_name -> api.auth.v1.Token
	15, // 2: api.auth.v1.RegisterUserRequest.device_data:type_name -> api.common.v1.UserDeviceData
	16, // 3: api.auth.v1.RegisterUserResponse.tokens:type_name -> api.auth.v1.Token
	15, // 4: api.auth.v1.RefreshTokenRequest.device_data:type_name -> api.common.v1.UserDeviceData
	15, // 5: api.auth.v1.LogoutRequest.device_data:type_name -> api.common.v1.UserDeviceData
	15, // 6: api.auth.v1.RemoveUserRequest.device_data:type_name -> api.common.v1.UserDeviceData
	14, // 7: api.auth.v1.GetJWKSResponse.jwks:type_name -> api.auth.v1.JWK
	8,  // [8:8] is the sub-list for method output_type
	8,  // [8:8] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_api_auth_v1_auth_messages_proto_init() }
func file_api_auth_v1_auth_messages_proto_init() {
	if File_api_auth_v1_auth_messages_proto != nil {
		return
	}
	file_api_auth_v1_auth_types_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_auth_v1_auth_messages_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   15,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_auth_v1_auth_messages_proto_goTypes,
		DependencyIndexes: file_api_auth_v1_auth_messages_proto_depIdxs,
		MessageInfos:      file_api_auth_v1_auth_messages_proto_msgTypes,
	}.Build()
	File_api_auth_v1_auth_messages_proto = out.File
	file_api_auth_v1_auth_messages_proto_rawDesc = nil
	file_api_auth_v1_auth_messages_proto_goTypes = nil
	file_api_auth_v1_auth_messages_proto_depIdxs = nil
}
