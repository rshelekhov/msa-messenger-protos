// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        (unknown)
// source: api/subscriber/v1/subscriber_service.proto

package subscriber

import (
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

var File_api_subscriber_v1_subscriber_service_proto protoreflect.FileDescriptor

var file_api_subscriber_v1_subscriber_service_proto_rawDesc = []byte{
	0x0a, 0x2a, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x72,
	0x2f, 0x76, 0x31, 0x2f, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x72, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11, 0x61, 0x70,
	0x69, 0x2e, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x1a,
	0x2b, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x72, 0x2f,
	0x76, 0x31, 0x2f, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x72, 0x5f, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0x82, 0x07, 0x0a,
	0x11, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x6d, 0x0a, 0x10, 0x53, 0x65, 0x6e, 0x64, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64,
	0x49, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x12, 0x2a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x75, 0x62,
	0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x46,
	0x72, 0x69, 0x65, 0x6e, 0x64, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x2b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69,
	0x62, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x46, 0x72, 0x69, 0x65, 0x6e,
	0x64, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x73, 0x0a, 0x12, 0x41, 0x63, 0x63, 0x65, 0x70, 0x74, 0x46, 0x72, 0x69, 0x65, 0x6e,
	0x64, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x12, 0x2c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x75,
	0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x63, 0x63, 0x65,
	0x70, 0x74, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x75, 0x62, 0x73,
	0x63, 0x72, 0x69, 0x62, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x63, 0x63, 0x65, 0x70, 0x74,
	0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x76, 0x0a, 0x13, 0x44, 0x65, 0x63, 0x6c, 0x69, 0x6e,
	0x65, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x12, 0x2d, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x72, 0x2e, 0x76,
	0x31, 0x2e, 0x44, 0x65, 0x63, 0x6c, 0x69, 0x6e, 0x65, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x49,
	0x6e, 0x76, 0x69, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2e, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x72, 0x2e, 0x76, 0x31,
	0x2e, 0x44, 0x65, 0x63, 0x6c, 0x69, 0x6e, 0x65, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x49, 0x6e,
	0x76, 0x69, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x6a,
	0x0a, 0x0f, 0x47, 0x65, 0x74, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x49, 0x6e, 0x76, 0x69, 0x74,
	0x65, 0x12, 0x29, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62,
	0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x49,
	0x6e, 0x76, 0x69, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2a, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x72, 0x2e, 0x76, 0x31,
	0x2e, 0x47, 0x65, 0x74, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x76, 0x0a, 0x13, 0x47, 0x65,
	0x74, 0x41, 0x6c, 0x6c, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x65,
	0x73, 0x12, 0x2d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62,
	0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x46, 0x72, 0x69, 0x65,
	0x6e, 0x64, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x2e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65,
	0x72, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x46, 0x72, 0x69, 0x65, 0x6e,
	0x64, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x6d, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x50,
	0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x2a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x75, 0x62,
	0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x46, 0x72,
	0x69, 0x65, 0x6e, 0x64, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x2b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69,
	0x62, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64,
	0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x5b, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x73, 0x12,
	0x24, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x72,
	0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x75, 0x62, 0x73,
	0x63, 0x72, 0x69, 0x62, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x46, 0x72, 0x69,
	0x65, 0x6e, 0x64, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x61,
	0x0a, 0x0c, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x12, 0x26,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x72, 0x2e,
	0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x75, 0x62,
	0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x42, 0xde, 0x01, 0x0a, 0x15, 0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x75,
	0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x42, 0x16, 0x53, 0x75, 0x62,
	0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x47, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x72, 0x73, 0x68, 0x65, 0x6c, 0x65, 0x6b, 0x68, 0x6f, 0x76, 0x2f, 0x6d, 0x73, 0x61,
	0x2d, 0x6d, 0x65, 0x73, 0x73, 0x65, 0x6e, 0x67, 0x65, 0x72, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x73, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x72,
	0x2f, 0x76, 0x31, 0x3b, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x72, 0xa2, 0x02,
	0x03, 0x41, 0x53, 0x58, 0xaa, 0x02, 0x11, 0x41, 0x70, 0x69, 0x2e, 0x53, 0x75, 0x62, 0x73, 0x63,
	0x72, 0x69, 0x62, 0x65, 0x72, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x11, 0x41, 0x70, 0x69, 0x5c, 0x53,
	0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x72, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x1d, 0x41,
	0x70, 0x69, 0x5c, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x72, 0x5c, 0x56, 0x31,
	0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x13, 0x41,
	0x70, 0x69, 0x3a, 0x3a, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x72, 0x3a, 0x3a,
	0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_api_subscriber_v1_subscriber_service_proto_goTypes = []any{
	(*SendFriendInviteRequest)(nil),     // 0: api.subscriber.v1.SendFriendInviteRequest
	(*AcceptFriendInviteRequest)(nil),   // 1: api.subscriber.v1.AcceptFriendInviteRequest
	(*DeclineFriendInviteRequest)(nil),  // 2: api.subscriber.v1.DeclineFriendInviteRequest
	(*GetFriendInviteRequest)(nil),      // 3: api.subscriber.v1.GetFriendInviteRequest
	(*GetAllFriendInvitesRequest)(nil),  // 4: api.subscriber.v1.GetAllFriendInvitesRequest
	(*GetFriendProfileRequest)(nil),     // 5: api.subscriber.v1.GetFriendProfileRequest
	(*GetFriendsRequest)(nil),           // 6: api.subscriber.v1.GetFriendsRequest
	(*DeleteFriendRequest)(nil),         // 7: api.subscriber.v1.DeleteFriendRequest
	(*SendFriendInviteResponse)(nil),    // 8: api.subscriber.v1.SendFriendInviteResponse
	(*AcceptFriendInviteResponse)(nil),  // 9: api.subscriber.v1.AcceptFriendInviteResponse
	(*DeclineFriendInviteResponse)(nil), // 10: api.subscriber.v1.DeclineFriendInviteResponse
	(*GetFriendInviteResponse)(nil),     // 11: api.subscriber.v1.GetFriendInviteResponse
	(*GetAllFriendInvitesResponse)(nil), // 12: api.subscriber.v1.GetAllFriendInvitesResponse
	(*GetFriendProfileResponse)(nil),    // 13: api.subscriber.v1.GetFriendProfileResponse
	(*GetFriendsResponse)(nil),          // 14: api.subscriber.v1.GetFriendsResponse
	(*DeleteFriendResponse)(nil),        // 15: api.subscriber.v1.DeleteFriendResponse
}
var file_api_subscriber_v1_subscriber_service_proto_depIdxs = []int32{
	0,  // 0: api.subscriber.v1.SubscriberService.SendFriendInvite:input_type -> api.subscriber.v1.SendFriendInviteRequest
	1,  // 1: api.subscriber.v1.SubscriberService.AcceptFriendInvite:input_type -> api.subscriber.v1.AcceptFriendInviteRequest
	2,  // 2: api.subscriber.v1.SubscriberService.DeclineFriendInvite:input_type -> api.subscriber.v1.DeclineFriendInviteRequest
	3,  // 3: api.subscriber.v1.SubscriberService.GetFriendInvite:input_type -> api.subscriber.v1.GetFriendInviteRequest
	4,  // 4: api.subscriber.v1.SubscriberService.GetAllFriendInvites:input_type -> api.subscriber.v1.GetAllFriendInvitesRequest
	5,  // 5: api.subscriber.v1.SubscriberService.GetFriendProfile:input_type -> api.subscriber.v1.GetFriendProfileRequest
	6,  // 6: api.subscriber.v1.SubscriberService.GetFriends:input_type -> api.subscriber.v1.GetFriendsRequest
	7,  // 7: api.subscriber.v1.SubscriberService.DeleteFriend:input_type -> api.subscriber.v1.DeleteFriendRequest
	8,  // 8: api.subscriber.v1.SubscriberService.SendFriendInvite:output_type -> api.subscriber.v1.SendFriendInviteResponse
	9,  // 9: api.subscriber.v1.SubscriberService.AcceptFriendInvite:output_type -> api.subscriber.v1.AcceptFriendInviteResponse
	10, // 10: api.subscriber.v1.SubscriberService.DeclineFriendInvite:output_type -> api.subscriber.v1.DeclineFriendInviteResponse
	11, // 11: api.subscriber.v1.SubscriberService.GetFriendInvite:output_type -> api.subscriber.v1.GetFriendInviteResponse
	12, // 12: api.subscriber.v1.SubscriberService.GetAllFriendInvites:output_type -> api.subscriber.v1.GetAllFriendInvitesResponse
	13, // 13: api.subscriber.v1.SubscriberService.GetFriendProfile:output_type -> api.subscriber.v1.GetFriendProfileResponse
	14, // 14: api.subscriber.v1.SubscriberService.GetFriends:output_type -> api.subscriber.v1.GetFriendsResponse
	15, // 15: api.subscriber.v1.SubscriberService.DeleteFriend:output_type -> api.subscriber.v1.DeleteFriendResponse
	8,  // [8:16] is the sub-list for method output_type
	0,  // [0:8] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_api_subscriber_v1_subscriber_service_proto_init() }
func file_api_subscriber_v1_subscriber_service_proto_init() {
	if File_api_subscriber_v1_subscriber_service_proto != nil {
		return
	}
	file_api_subscriber_v1_subscriber_messages_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_subscriber_v1_subscriber_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_subscriber_v1_subscriber_service_proto_goTypes,
		DependencyIndexes: file_api_subscriber_v1_subscriber_service_proto_depIdxs,
	}.Build()
	File_api_subscriber_v1_subscriber_service_proto = out.File
	file_api_subscriber_v1_subscriber_service_proto_rawDesc = nil
	file_api_subscriber_v1_subscriber_service_proto_goTypes = nil
	file_api_subscriber_v1_subscriber_service_proto_depIdxs = nil
}
