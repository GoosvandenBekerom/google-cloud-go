// Copyright 2022 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.12
// source: google/cloud/channel/v1/common.proto

package channelpb

import (
	reflect "reflect"
	sync "sync"

	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Enum to specify the institute type.
type EduData_InstituteType int32

const (
	// Not used.
	EduData_INSTITUTE_TYPE_UNSPECIFIED EduData_InstituteType = 0
	// Elementary/Secondary Schools & Districts
	EduData_K12 EduData_InstituteType = 1
	// Higher Education Universities & Colleges
	EduData_UNIVERSITY EduData_InstituteType = 2
)

// Enum value maps for EduData_InstituteType.
var (
	EduData_InstituteType_name = map[int32]string{
		0: "INSTITUTE_TYPE_UNSPECIFIED",
		1: "K12",
		2: "UNIVERSITY",
	}
	EduData_InstituteType_value = map[string]int32{
		"INSTITUTE_TYPE_UNSPECIFIED": 0,
		"K12":                        1,
		"UNIVERSITY":                 2,
	}
)

func (x EduData_InstituteType) Enum() *EduData_InstituteType {
	p := new(EduData_InstituteType)
	*p = x
	return p
}

func (x EduData_InstituteType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (EduData_InstituteType) Descriptor() protoreflect.EnumDescriptor {
	return file_google_cloud_channel_v1_common_proto_enumTypes[0].Descriptor()
}

func (EduData_InstituteType) Type() protoreflect.EnumType {
	return &file_google_cloud_channel_v1_common_proto_enumTypes[0]
}

func (x EduData_InstituteType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use EduData_InstituteType.Descriptor instead.
func (EduData_InstituteType) EnumDescriptor() ([]byte, []int) {
	return file_google_cloud_channel_v1_common_proto_rawDescGZIP(), []int{0, 0}
}

// Number of students and staff the institute has.
type EduData_InstituteSize int32

const (
	// Not used.
	EduData_INSTITUTE_SIZE_UNSPECIFIED EduData_InstituteSize = 0
	// 1 - 100
	EduData_SIZE_1_100 EduData_InstituteSize = 1
	// 101 - 500
	EduData_SIZE_101_500 EduData_InstituteSize = 2
	// 501 - 1,000
	EduData_SIZE_501_1000 EduData_InstituteSize = 3
	// 1,001 - 2,000
	EduData_SIZE_1001_2000 EduData_InstituteSize = 4
	// 2,001 - 5,000
	EduData_SIZE_2001_5000 EduData_InstituteSize = 5
	// 5,001 - 10,000
	EduData_SIZE_5001_10000 EduData_InstituteSize = 6
	// 10,001 +
	EduData_SIZE_10001_OR_MORE EduData_InstituteSize = 7
)

// Enum value maps for EduData_InstituteSize.
var (
	EduData_InstituteSize_name = map[int32]string{
		0: "INSTITUTE_SIZE_UNSPECIFIED",
		1: "SIZE_1_100",
		2: "SIZE_101_500",
		3: "SIZE_501_1000",
		4: "SIZE_1001_2000",
		5: "SIZE_2001_5000",
		6: "SIZE_5001_10000",
		7: "SIZE_10001_OR_MORE",
	}
	EduData_InstituteSize_value = map[string]int32{
		"INSTITUTE_SIZE_UNSPECIFIED": 0,
		"SIZE_1_100":                 1,
		"SIZE_101_500":               2,
		"SIZE_501_1000":              3,
		"SIZE_1001_2000":             4,
		"SIZE_2001_5000":             5,
		"SIZE_5001_10000":            6,
		"SIZE_10001_OR_MORE":         7,
	}
)

func (x EduData_InstituteSize) Enum() *EduData_InstituteSize {
	p := new(EduData_InstituteSize)
	*p = x
	return p
}

func (x EduData_InstituteSize) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (EduData_InstituteSize) Descriptor() protoreflect.EnumDescriptor {
	return file_google_cloud_channel_v1_common_proto_enumTypes[1].Descriptor()
}

func (EduData_InstituteSize) Type() protoreflect.EnumType {
	return &file_google_cloud_channel_v1_common_proto_enumTypes[1]
}

func (x EduData_InstituteSize) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use EduData_InstituteSize.Descriptor instead.
func (EduData_InstituteSize) EnumDescriptor() ([]byte, []int) {
	return file_google_cloud_channel_v1_common_proto_rawDescGZIP(), []int{0, 1}
}

// CustomerType of the customer
type CloudIdentityInfo_CustomerType int32

const (
	// Not used.
	CloudIdentityInfo_CUSTOMER_TYPE_UNSPECIFIED CloudIdentityInfo_CustomerType = 0
	// Domain-owning customer which needs domain verification to use services.
	CloudIdentityInfo_DOMAIN CloudIdentityInfo_CustomerType = 1
	// Team customer which needs email verification to use services.
	CloudIdentityInfo_TEAM CloudIdentityInfo_CustomerType = 2
)

// Enum value maps for CloudIdentityInfo_CustomerType.
var (
	CloudIdentityInfo_CustomerType_name = map[int32]string{
		0: "CUSTOMER_TYPE_UNSPECIFIED",
		1: "DOMAIN",
		2: "TEAM",
	}
	CloudIdentityInfo_CustomerType_value = map[string]int32{
		"CUSTOMER_TYPE_UNSPECIFIED": 0,
		"DOMAIN":                    1,
		"TEAM":                      2,
	}
)

func (x CloudIdentityInfo_CustomerType) Enum() *CloudIdentityInfo_CustomerType {
	p := new(CloudIdentityInfo_CustomerType)
	*p = x
	return p
}

func (x CloudIdentityInfo_CustomerType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CloudIdentityInfo_CustomerType) Descriptor() protoreflect.EnumDescriptor {
	return file_google_cloud_channel_v1_common_proto_enumTypes[2].Descriptor()
}

func (CloudIdentityInfo_CustomerType) Type() protoreflect.EnumType {
	return &file_google_cloud_channel_v1_common_proto_enumTypes[2]
}

func (x CloudIdentityInfo_CustomerType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CloudIdentityInfo_CustomerType.Descriptor instead.
func (CloudIdentityInfo_CustomerType) EnumDescriptor() ([]byte, []int) {
	return file_google_cloud_channel_v1_common_proto_rawDescGZIP(), []int{1, 0}
}

// Required Edu Attributes
type EduData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Designated institute type of customer.
	InstituteType EduData_InstituteType `protobuf:"varint,1,opt,name=institute_type,json=instituteType,proto3,enum=google.cloud.channel.v1.EduData_InstituteType" json:"institute_type,omitempty"`
	// Size of the institute.
	InstituteSize EduData_InstituteSize `protobuf:"varint,2,opt,name=institute_size,json=instituteSize,proto3,enum=google.cloud.channel.v1.EduData_InstituteSize" json:"institute_size,omitempty"`
	// Web address for the edu customer's institution.
	Website string `protobuf:"bytes,3,opt,name=website,proto3" json:"website,omitempty"`
}

func (x *EduData) Reset() {
	*x = EduData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_channel_v1_common_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EduData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EduData) ProtoMessage() {}

func (x *EduData) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_channel_v1_common_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EduData.ProtoReflect.Descriptor instead.
func (*EduData) Descriptor() ([]byte, []int) {
	return file_google_cloud_channel_v1_common_proto_rawDescGZIP(), []int{0}
}

func (x *EduData) GetInstituteType() EduData_InstituteType {
	if x != nil {
		return x.InstituteType
	}
	return EduData_INSTITUTE_TYPE_UNSPECIFIED
}

func (x *EduData) GetInstituteSize() EduData_InstituteSize {
	if x != nil {
		return x.InstituteSize
	}
	return EduData_INSTITUTE_SIZE_UNSPECIFIED
}

func (x *EduData) GetWebsite() string {
	if x != nil {
		return x.Website
	}
	return ""
}

// Cloud Identity information for the Cloud Channel Customer.
type CloudIdentityInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// CustomerType indicates verification type needed for using services.
	CustomerType CloudIdentityInfo_CustomerType `protobuf:"varint,1,opt,name=customer_type,json=customerType,proto3,enum=google.cloud.channel.v1.CloudIdentityInfo_CustomerType" json:"customer_type,omitempty"`
	// Output only. The primary domain name.
	PrimaryDomain string `protobuf:"bytes,9,opt,name=primary_domain,json=primaryDomain,proto3" json:"primary_domain,omitempty"`
	// Output only. Whether the domain is verified.
	// This field is not returned for a Customer's cloud_identity_info resource.
	// Partners can use the domains.get() method of the Workspace SDK's
	// Directory API, or listen to the PRIMARY_DOMAIN_VERIFIED Pub/Sub event in
	// to track domain verification of their resolve Workspace customers.
	IsDomainVerified bool `protobuf:"varint,4,opt,name=is_domain_verified,json=isDomainVerified,proto3" json:"is_domain_verified,omitempty"`
	// The alternate email.
	AlternateEmail string `protobuf:"bytes,6,opt,name=alternate_email,json=alternateEmail,proto3" json:"alternate_email,omitempty"`
	// Phone number associated with the Cloud Identity.
	PhoneNumber string `protobuf:"bytes,7,opt,name=phone_number,json=phoneNumber,proto3" json:"phone_number,omitempty"`
	// Language code.
	LanguageCode string `protobuf:"bytes,8,opt,name=language_code,json=languageCode,proto3" json:"language_code,omitempty"`
	// Output only. URI of Customer's Admin console dashboard.
	AdminConsoleUri string `protobuf:"bytes,10,opt,name=admin_console_uri,json=adminConsoleUri,proto3" json:"admin_console_uri,omitempty"`
	// Edu information about the customer.
	EduData *EduData `protobuf:"bytes,22,opt,name=edu_data,json=eduData,proto3" json:"edu_data,omitempty"`
}

func (x *CloudIdentityInfo) Reset() {
	*x = CloudIdentityInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_channel_v1_common_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CloudIdentityInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CloudIdentityInfo) ProtoMessage() {}

func (x *CloudIdentityInfo) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_channel_v1_common_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CloudIdentityInfo.ProtoReflect.Descriptor instead.
func (*CloudIdentityInfo) Descriptor() ([]byte, []int) {
	return file_google_cloud_channel_v1_common_proto_rawDescGZIP(), []int{1}
}

func (x *CloudIdentityInfo) GetCustomerType() CloudIdentityInfo_CustomerType {
	if x != nil {
		return x.CustomerType
	}
	return CloudIdentityInfo_CUSTOMER_TYPE_UNSPECIFIED
}

func (x *CloudIdentityInfo) GetPrimaryDomain() string {
	if x != nil {
		return x.PrimaryDomain
	}
	return ""
}

func (x *CloudIdentityInfo) GetIsDomainVerified() bool {
	if x != nil {
		return x.IsDomainVerified
	}
	return false
}

func (x *CloudIdentityInfo) GetAlternateEmail() string {
	if x != nil {
		return x.AlternateEmail
	}
	return ""
}

func (x *CloudIdentityInfo) GetPhoneNumber() string {
	if x != nil {
		return x.PhoneNumber
	}
	return ""
}

func (x *CloudIdentityInfo) GetLanguageCode() string {
	if x != nil {
		return x.LanguageCode
	}
	return ""
}

func (x *CloudIdentityInfo) GetAdminConsoleUri() string {
	if x != nil {
		return x.AdminConsoleUri
	}
	return ""
}

func (x *CloudIdentityInfo) GetEduData() *EduData {
	if x != nil {
		return x.EduData
	}
	return nil
}

// Data type and value of a parameter.
type Value struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The kind of value.
	//
	// Types that are assignable to Kind:
	//	*Value_Int64Value
	//	*Value_StringValue
	//	*Value_DoubleValue
	//	*Value_ProtoValue
	//	*Value_BoolValue
	Kind isValue_Kind `protobuf_oneof:"kind"`
}

func (x *Value) Reset() {
	*x = Value{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_channel_v1_common_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Value) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Value) ProtoMessage() {}

func (x *Value) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_channel_v1_common_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Value.ProtoReflect.Descriptor instead.
func (*Value) Descriptor() ([]byte, []int) {
	return file_google_cloud_channel_v1_common_proto_rawDescGZIP(), []int{2}
}

func (m *Value) GetKind() isValue_Kind {
	if m != nil {
		return m.Kind
	}
	return nil
}

func (x *Value) GetInt64Value() int64 {
	if x, ok := x.GetKind().(*Value_Int64Value); ok {
		return x.Int64Value
	}
	return 0
}

func (x *Value) GetStringValue() string {
	if x, ok := x.GetKind().(*Value_StringValue); ok {
		return x.StringValue
	}
	return ""
}

func (x *Value) GetDoubleValue() float64 {
	if x, ok := x.GetKind().(*Value_DoubleValue); ok {
		return x.DoubleValue
	}
	return 0
}

func (x *Value) GetProtoValue() *anypb.Any {
	if x, ok := x.GetKind().(*Value_ProtoValue); ok {
		return x.ProtoValue
	}
	return nil
}

func (x *Value) GetBoolValue() bool {
	if x, ok := x.GetKind().(*Value_BoolValue); ok {
		return x.BoolValue
	}
	return false
}

type isValue_Kind interface {
	isValue_Kind()
}

type Value_Int64Value struct {
	// Represents an int64 value.
	Int64Value int64 `protobuf:"varint,1,opt,name=int64_value,json=int64Value,proto3,oneof"`
}

type Value_StringValue struct {
	// Represents a string value.
	StringValue string `protobuf:"bytes,2,opt,name=string_value,json=stringValue,proto3,oneof"`
}

type Value_DoubleValue struct {
	// Represents a double value.
	DoubleValue float64 `protobuf:"fixed64,3,opt,name=double_value,json=doubleValue,proto3,oneof"`
}

type Value_ProtoValue struct {
	// Represents an 'Any' proto value.
	ProtoValue *anypb.Any `protobuf:"bytes,4,opt,name=proto_value,json=protoValue,proto3,oneof"`
}

type Value_BoolValue struct {
	// Represents a boolean value.
	BoolValue bool `protobuf:"varint,5,opt,name=bool_value,json=boolValue,proto3,oneof"`
}

func (*Value_Int64Value) isValue_Kind() {}

func (*Value_StringValue) isValue_Kind() {}

func (*Value_DoubleValue) isValue_Kind() {}

func (*Value_ProtoValue) isValue_Kind() {}

func (*Value_BoolValue) isValue_Kind() {}

// Information needed to create an Admin User for Google Workspace.
type AdminUser struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Primary email of the admin user.
	Email string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	// Given name of the admin user.
	GivenName string `protobuf:"bytes,2,opt,name=given_name,json=givenName,proto3" json:"given_name,omitempty"`
	// Family name of the admin user.
	FamilyName string `protobuf:"bytes,3,opt,name=family_name,json=familyName,proto3" json:"family_name,omitempty"`
}

func (x *AdminUser) Reset() {
	*x = AdminUser{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_channel_v1_common_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AdminUser) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdminUser) ProtoMessage() {}

func (x *AdminUser) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_channel_v1_common_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdminUser.ProtoReflect.Descriptor instead.
func (*AdminUser) Descriptor() ([]byte, []int) {
	return file_google_cloud_channel_v1_common_proto_rawDescGZIP(), []int{3}
}

func (x *AdminUser) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *AdminUser) GetGivenName() string {
	if x != nil {
		return x.GivenName
	}
	return ""
}

func (x *AdminUser) GetFamilyName() string {
	if x != nil {
		return x.FamilyName
	}
	return ""
}

var File_google_cloud_channel_v1_common_proto protoreflect.FileDescriptor

var file_google_cloud_channel_v1_common_proto_rawDesc = []byte{
	0x0a, 0x24, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x63,
	0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x17, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x2e, 0x76, 0x31, 0x1a,
	0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c,
	0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd7, 0x03, 0x0a, 0x07,
	0x45, 0x64, 0x75, 0x44, 0x61, 0x74, 0x61, 0x12, 0x55, 0x0a, 0x0e, 0x69, 0x6e, 0x73, 0x74, 0x69,
	0x74, 0x75, 0x74, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x2e, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x63,
	0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x64, 0x75, 0x44, 0x61, 0x74,
	0x61, 0x2e, 0x49, 0x6e, 0x73, 0x74, 0x69, 0x74, 0x75, 0x74, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52,
	0x0d, 0x69, 0x6e, 0x73, 0x74, 0x69, 0x74, 0x75, 0x74, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x55,
	0x0a, 0x0e, 0x69, 0x6e, 0x73, 0x74, 0x69, 0x74, 0x75, 0x74, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2e, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x2e, 0x76, 0x31,
	0x2e, 0x45, 0x64, 0x75, 0x44, 0x61, 0x74, 0x61, 0x2e, 0x49, 0x6e, 0x73, 0x74, 0x69, 0x74, 0x75,
	0x74, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x52, 0x0d, 0x69, 0x6e, 0x73, 0x74, 0x69, 0x74, 0x75, 0x74,
	0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x77, 0x65, 0x62, 0x73, 0x69, 0x74, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x77, 0x65, 0x62, 0x73, 0x69, 0x74, 0x65, 0x22,
	0x48, 0x0a, 0x0d, 0x49, 0x6e, 0x73, 0x74, 0x69, 0x74, 0x75, 0x74, 0x65, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x1e, 0x0a, 0x1a, 0x49, 0x4e, 0x53, 0x54, 0x49, 0x54, 0x55, 0x54, 0x45, 0x5f, 0x54, 0x59,
	0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00,
	0x12, 0x07, 0x0a, 0x03, 0x4b, 0x31, 0x32, 0x10, 0x01, 0x12, 0x0e, 0x0a, 0x0a, 0x55, 0x4e, 0x49,
	0x56, 0x45, 0x52, 0x53, 0x49, 0x54, 0x59, 0x10, 0x02, 0x22, 0xb9, 0x01, 0x0a, 0x0d, 0x49, 0x6e,
	0x73, 0x74, 0x69, 0x74, 0x75, 0x74, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x1e, 0x0a, 0x1a, 0x49,
	0x4e, 0x53, 0x54, 0x49, 0x54, 0x55, 0x54, 0x45, 0x5f, 0x53, 0x49, 0x5a, 0x45, 0x5f, 0x55, 0x4e,
	0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0e, 0x0a, 0x0a, 0x53,
	0x49, 0x5a, 0x45, 0x5f, 0x31, 0x5f, 0x31, 0x30, 0x30, 0x10, 0x01, 0x12, 0x10, 0x0a, 0x0c, 0x53,
	0x49, 0x5a, 0x45, 0x5f, 0x31, 0x30, 0x31, 0x5f, 0x35, 0x30, 0x30, 0x10, 0x02, 0x12, 0x11, 0x0a,
	0x0d, 0x53, 0x49, 0x5a, 0x45, 0x5f, 0x35, 0x30, 0x31, 0x5f, 0x31, 0x30, 0x30, 0x30, 0x10, 0x03,
	0x12, 0x12, 0x0a, 0x0e, 0x53, 0x49, 0x5a, 0x45, 0x5f, 0x31, 0x30, 0x30, 0x31, 0x5f, 0x32, 0x30,
	0x30, 0x30, 0x10, 0x04, 0x12, 0x12, 0x0a, 0x0e, 0x53, 0x49, 0x5a, 0x45, 0x5f, 0x32, 0x30, 0x30,
	0x31, 0x5f, 0x35, 0x30, 0x30, 0x30, 0x10, 0x05, 0x12, 0x13, 0x0a, 0x0f, 0x53, 0x49, 0x5a, 0x45,
	0x5f, 0x35, 0x30, 0x30, 0x31, 0x5f, 0x31, 0x30, 0x30, 0x30, 0x30, 0x10, 0x06, 0x12, 0x16, 0x0a,
	0x12, 0x53, 0x49, 0x5a, 0x45, 0x5f, 0x31, 0x30, 0x30, 0x30, 0x31, 0x5f, 0x4f, 0x52, 0x5f, 0x4d,
	0x4f, 0x52, 0x45, 0x10, 0x07, 0x22, 0xf4, 0x03, 0x0a, 0x11, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x49,
	0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x5c, 0x0a, 0x0d, 0x63,
	0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x37, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6c, 0x6f,
	0x75, 0x64, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x43,
	0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0c, 0x63, 0x75, 0x73,
	0x74, 0x6f, 0x6d, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x12, 0x2a, 0x0a, 0x0e, 0x70, 0x72, 0x69,
	0x6d, 0x61, 0x72, 0x79, 0x5f, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x18, 0x09, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x0d, 0x70, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x44,
	0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x12, 0x31, 0x0a, 0x12, 0x69, 0x73, 0x5f, 0x64, 0x6f, 0x6d, 0x61,
	0x69, 0x6e, 0x5f, 0x76, 0x65, 0x72, 0x69, 0x66, 0x69, 0x65, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x08, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x10, 0x69, 0x73, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e,
	0x56, 0x65, 0x72, 0x69, 0x66, 0x69, 0x65, 0x64, 0x12, 0x27, 0x0a, 0x0f, 0x61, 0x6c, 0x74, 0x65,
	0x72, 0x6e, 0x61, 0x74, 0x65, 0x5f, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0e, 0x61, 0x6c, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x74, 0x65, 0x45, 0x6d, 0x61, 0x69,
	0x6c, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x12, 0x23, 0x0a, 0x0d, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65,
	0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x6c, 0x61, 0x6e,
	0x67, 0x75, 0x61, 0x67, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x2f, 0x0a, 0x11, 0x61, 0x64, 0x6d,
	0x69, 0x6e, 0x5f, 0x63, 0x6f, 0x6e, 0x73, 0x6f, 0x6c, 0x65, 0x5f, 0x75, 0x72, 0x69, 0x18, 0x0a,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x0f, 0x61, 0x64, 0x6d, 0x69, 0x6e,
	0x43, 0x6f, 0x6e, 0x73, 0x6f, 0x6c, 0x65, 0x55, 0x72, 0x69, 0x12, 0x3b, 0x0a, 0x08, 0x65, 0x64,
	0x75, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x16, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x63, 0x68, 0x61, 0x6e,
	0x6e, 0x65, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x64, 0x75, 0x44, 0x61, 0x74, 0x61, 0x52, 0x07,
	0x65, 0x64, 0x75, 0x44, 0x61, 0x74, 0x61, 0x22, 0x43, 0x0a, 0x0c, 0x43, 0x75, 0x73, 0x74, 0x6f,
	0x6d, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1d, 0x0a, 0x19, 0x43, 0x55, 0x53, 0x54, 0x4f,
	0x4d, 0x45, 0x52, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49,
	0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x44, 0x4f, 0x4d, 0x41, 0x49, 0x4e,
	0x10, 0x01, 0x12, 0x08, 0x0a, 0x04, 0x54, 0x45, 0x41, 0x4d, 0x10, 0x02, 0x22, 0xd6, 0x01, 0x0a,
	0x05, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x21, 0x0a, 0x0b, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x5f,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x48, 0x00, 0x52, 0x0a, 0x69,
	0x6e, 0x74, 0x36, 0x34, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x23, 0x0a, 0x0c, 0x73, 0x74, 0x72,
	0x69, 0x6e, 0x67, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48,
	0x00, 0x52, 0x0b, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x23,
	0x0a, 0x0c, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x01, 0x48, 0x00, 0x52, 0x0b, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x12, 0x37, 0x0a, 0x0b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x5f, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x48, 0x00,
	0x52, 0x0a, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x1f, 0x0a, 0x0a,
	0x62, 0x6f, 0x6f, 0x6c, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08,
	0x48, 0x00, 0x52, 0x09, 0x62, 0x6f, 0x6f, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x06, 0x0a,
	0x04, 0x6b, 0x69, 0x6e, 0x64, 0x22, 0x61, 0x0a, 0x09, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x55, 0x73,
	0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1d, 0x0a, 0x0a, 0x67, 0x69, 0x76, 0x65,
	0x6e, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x67, 0x69,
	0x76, 0x65, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x66, 0x61, 0x6d, 0x69, 0x6c,
	0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x66, 0x61,
	0x6d, 0x69, 0x6c, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x42, 0x63, 0x0a, 0x1b, 0x63, 0x6f, 0x6d, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x63, 0x68, 0x61,
	0x6e, 0x6e, 0x65, 0x6c, 0x2e, 0x76, 0x31, 0x42, 0x0b, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x35, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x2f, 0x63, 0x68, 0x61, 0x6e,
	0x6e, 0x65, 0x6c, 0x2f, 0x61, 0x70, 0x69, 0x76, 0x31, 0x2f, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65,
	0x6c, 0x70, 0x62, 0x3b, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x70, 0x62, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_cloud_channel_v1_common_proto_rawDescOnce sync.Once
	file_google_cloud_channel_v1_common_proto_rawDescData = file_google_cloud_channel_v1_common_proto_rawDesc
)

func file_google_cloud_channel_v1_common_proto_rawDescGZIP() []byte {
	file_google_cloud_channel_v1_common_proto_rawDescOnce.Do(func() {
		file_google_cloud_channel_v1_common_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_channel_v1_common_proto_rawDescData)
	})
	return file_google_cloud_channel_v1_common_proto_rawDescData
}

var file_google_cloud_channel_v1_common_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_google_cloud_channel_v1_common_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_google_cloud_channel_v1_common_proto_goTypes = []interface{}{
	(EduData_InstituteType)(0),          // 0: google.cloud.channel.v1.EduData.InstituteType
	(EduData_InstituteSize)(0),          // 1: google.cloud.channel.v1.EduData.InstituteSize
	(CloudIdentityInfo_CustomerType)(0), // 2: google.cloud.channel.v1.CloudIdentityInfo.CustomerType
	(*EduData)(nil),                     // 3: google.cloud.channel.v1.EduData
	(*CloudIdentityInfo)(nil),           // 4: google.cloud.channel.v1.CloudIdentityInfo
	(*Value)(nil),                       // 5: google.cloud.channel.v1.Value
	(*AdminUser)(nil),                   // 6: google.cloud.channel.v1.AdminUser
	(*anypb.Any)(nil),                   // 7: google.protobuf.Any
}
var file_google_cloud_channel_v1_common_proto_depIdxs = []int32{
	0, // 0: google.cloud.channel.v1.EduData.institute_type:type_name -> google.cloud.channel.v1.EduData.InstituteType
	1, // 1: google.cloud.channel.v1.EduData.institute_size:type_name -> google.cloud.channel.v1.EduData.InstituteSize
	2, // 2: google.cloud.channel.v1.CloudIdentityInfo.customer_type:type_name -> google.cloud.channel.v1.CloudIdentityInfo.CustomerType
	3, // 3: google.cloud.channel.v1.CloudIdentityInfo.edu_data:type_name -> google.cloud.channel.v1.EduData
	7, // 4: google.cloud.channel.v1.Value.proto_value:type_name -> google.protobuf.Any
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_google_cloud_channel_v1_common_proto_init() }
func file_google_cloud_channel_v1_common_proto_init() {
	if File_google_cloud_channel_v1_common_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_cloud_channel_v1_common_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EduData); i {
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
		file_google_cloud_channel_v1_common_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CloudIdentityInfo); i {
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
		file_google_cloud_channel_v1_common_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Value); i {
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
		file_google_cloud_channel_v1_common_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AdminUser); i {
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
	file_google_cloud_channel_v1_common_proto_msgTypes[2].OneofWrappers = []interface{}{
		(*Value_Int64Value)(nil),
		(*Value_StringValue)(nil),
		(*Value_DoubleValue)(nil),
		(*Value_ProtoValue)(nil),
		(*Value_BoolValue)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_cloud_channel_v1_common_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_cloud_channel_v1_common_proto_goTypes,
		DependencyIndexes: file_google_cloud_channel_v1_common_proto_depIdxs,
		EnumInfos:         file_google_cloud_channel_v1_common_proto_enumTypes,
		MessageInfos:      file_google_cloud_channel_v1_common_proto_msgTypes,
	}.Build()
	File_google_cloud_channel_v1_common_proto = out.File
	file_google_cloud_channel_v1_common_proto_rawDesc = nil
	file_google_cloud_channel_v1_common_proto_goTypes = nil
	file_google_cloud_channel_v1_common_proto_depIdxs = nil
}
