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
// source: google/cloud/websecurityscanner/v1/finding_addon.proto

package websecurityscannerpb

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Types of XSS attack vector.
type Xss_AttackVector int32

const (
	// Unknown attack vector.
	Xss_ATTACK_VECTOR_UNSPECIFIED Xss_AttackVector = 0
	// The attack comes from fuzzing the browser's localStorage.
	Xss_LOCAL_STORAGE Xss_AttackVector = 1
	// The attack comes from fuzzing the browser's sessionStorage.
	Xss_SESSION_STORAGE Xss_AttackVector = 2
	// The attack comes from fuzzing the window's name property.
	Xss_WINDOW_NAME Xss_AttackVector = 3
	// The attack comes from fuzzing the referrer property.
	Xss_REFERRER Xss_AttackVector = 4
	// The attack comes from fuzzing an input element.
	Xss_FORM_INPUT Xss_AttackVector = 5
	// The attack comes from fuzzing the browser's cookies.
	Xss_COOKIE Xss_AttackVector = 6
	// The attack comes from hijacking the post messaging mechanism.
	Xss_POST_MESSAGE Xss_AttackVector = 7
	// The attack comes from fuzzing parameters in the url.
	Xss_GET_PARAMETERS Xss_AttackVector = 8
	// The attack comes from fuzzing the fragment in the url.
	Xss_URL_FRAGMENT Xss_AttackVector = 9
	// The attack comes from fuzzing the HTML comments.
	Xss_HTML_COMMENT Xss_AttackVector = 10
	// The attack comes from fuzzing the POST parameters.
	Xss_POST_PARAMETERS Xss_AttackVector = 11
	// The attack comes from fuzzing the protocol.
	Xss_PROTOCOL Xss_AttackVector = 12
	// The attack comes from the server side and is stored.
	Xss_STORED_XSS Xss_AttackVector = 13
	// The attack is a Same-Origin Method Execution attack via a GET parameter.
	Xss_SAME_ORIGIN Xss_AttackVector = 14
	// The attack payload is received from a third-party host via a URL that is
	// user-controllable
	Xss_USER_CONTROLLABLE_URL Xss_AttackVector = 15
)

// Enum value maps for Xss_AttackVector.
var (
	Xss_AttackVector_name = map[int32]string{
		0:  "ATTACK_VECTOR_UNSPECIFIED",
		1:  "LOCAL_STORAGE",
		2:  "SESSION_STORAGE",
		3:  "WINDOW_NAME",
		4:  "REFERRER",
		5:  "FORM_INPUT",
		6:  "COOKIE",
		7:  "POST_MESSAGE",
		8:  "GET_PARAMETERS",
		9:  "URL_FRAGMENT",
		10: "HTML_COMMENT",
		11: "POST_PARAMETERS",
		12: "PROTOCOL",
		13: "STORED_XSS",
		14: "SAME_ORIGIN",
		15: "USER_CONTROLLABLE_URL",
	}
	Xss_AttackVector_value = map[string]int32{
		"ATTACK_VECTOR_UNSPECIFIED": 0,
		"LOCAL_STORAGE":             1,
		"SESSION_STORAGE":           2,
		"WINDOW_NAME":               3,
		"REFERRER":                  4,
		"FORM_INPUT":                5,
		"COOKIE":                    6,
		"POST_MESSAGE":              7,
		"GET_PARAMETERS":            8,
		"URL_FRAGMENT":              9,
		"HTML_COMMENT":              10,
		"POST_PARAMETERS":           11,
		"PROTOCOL":                  12,
		"STORED_XSS":                13,
		"SAME_ORIGIN":               14,
		"USER_CONTROLLABLE_URL":     15,
	}
)

func (x Xss_AttackVector) Enum() *Xss_AttackVector {
	p := new(Xss_AttackVector)
	*p = x
	return p
}

func (x Xss_AttackVector) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Xss_AttackVector) Descriptor() protoreflect.EnumDescriptor {
	return file_google_cloud_websecurityscanner_v1_finding_addon_proto_enumTypes[0].Descriptor()
}

func (Xss_AttackVector) Type() protoreflect.EnumType {
	return &file_google_cloud_websecurityscanner_v1_finding_addon_proto_enumTypes[0]
}

func (x Xss_AttackVector) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Xss_AttackVector.Descriptor instead.
func (Xss_AttackVector) EnumDescriptor() ([]byte, []int) {
	return file_google_cloud_websecurityscanner_v1_finding_addon_proto_rawDescGZIP(), []int{5, 0}
}

// Locations within a request where XML was substituted.
type Xxe_Location int32

const (
	// Unknown Location.
	Xxe_LOCATION_UNSPECIFIED Xxe_Location = 0
	// The XML payload replaced the complete request body.
	Xxe_COMPLETE_REQUEST_BODY Xxe_Location = 1
)

// Enum value maps for Xxe_Location.
var (
	Xxe_Location_name = map[int32]string{
		0: "LOCATION_UNSPECIFIED",
		1: "COMPLETE_REQUEST_BODY",
	}
	Xxe_Location_value = map[string]int32{
		"LOCATION_UNSPECIFIED":  0,
		"COMPLETE_REQUEST_BODY": 1,
	}
)

func (x Xxe_Location) Enum() *Xxe_Location {
	p := new(Xxe_Location)
	*p = x
	return p
}

func (x Xxe_Location) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Xxe_Location) Descriptor() protoreflect.EnumDescriptor {
	return file_google_cloud_websecurityscanner_v1_finding_addon_proto_enumTypes[1].Descriptor()
}

func (Xxe_Location) Type() protoreflect.EnumType {
	return &file_google_cloud_websecurityscanner_v1_finding_addon_proto_enumTypes[1]
}

func (x Xxe_Location) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Xxe_Location.Descriptor instead.
func (Xxe_Location) EnumDescriptor() ([]byte, []int) {
	return file_google_cloud_websecurityscanner_v1_finding_addon_proto_rawDescGZIP(), []int{6, 0}
}

// ! Information about a vulnerability with an HTML.
type Form struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ! The URI where to send the form when it's submitted.
	ActionUri string `protobuf:"bytes,1,opt,name=action_uri,json=actionUri,proto3" json:"action_uri,omitempty"`
	// ! The names of form fields related to the vulnerability.
	Fields []string `protobuf:"bytes,2,rep,name=fields,proto3" json:"fields,omitempty"`
}

func (x *Form) Reset() {
	*x = Form{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_websecurityscanner_v1_finding_addon_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Form) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Form) ProtoMessage() {}

func (x *Form) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_websecurityscanner_v1_finding_addon_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Form.ProtoReflect.Descriptor instead.
func (*Form) Descriptor() ([]byte, []int) {
	return file_google_cloud_websecurityscanner_v1_finding_addon_proto_rawDescGZIP(), []int{0}
}

func (x *Form) GetActionUri() string {
	if x != nil {
		return x.ActionUri
	}
	return ""
}

func (x *Form) GetFields() []string {
	if x != nil {
		return x.Fields
	}
	return nil
}

// Information reported for an outdated library.
type OutdatedLibrary struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The name of the outdated library.
	LibraryName string `protobuf:"bytes,1,opt,name=library_name,json=libraryName,proto3" json:"library_name,omitempty"`
	// The version number.
	Version string `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
	// URLs to learn more information about the vulnerabilities in the library.
	LearnMoreUrls []string `protobuf:"bytes,3,rep,name=learn_more_urls,json=learnMoreUrls,proto3" json:"learn_more_urls,omitempty"`
}

func (x *OutdatedLibrary) Reset() {
	*x = OutdatedLibrary{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_websecurityscanner_v1_finding_addon_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OutdatedLibrary) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OutdatedLibrary) ProtoMessage() {}

func (x *OutdatedLibrary) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_websecurityscanner_v1_finding_addon_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OutdatedLibrary.ProtoReflect.Descriptor instead.
func (*OutdatedLibrary) Descriptor() ([]byte, []int) {
	return file_google_cloud_websecurityscanner_v1_finding_addon_proto_rawDescGZIP(), []int{1}
}

func (x *OutdatedLibrary) GetLibraryName() string {
	if x != nil {
		return x.LibraryName
	}
	return ""
}

func (x *OutdatedLibrary) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *OutdatedLibrary) GetLearnMoreUrls() []string {
	if x != nil {
		return x.LearnMoreUrls
	}
	return nil
}

// Information regarding any resource causing the vulnerability such
// as JavaScript sources, image, audio files, etc.
type ViolatingResource struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The MIME type of this resource.
	ContentType string `protobuf:"bytes,1,opt,name=content_type,json=contentType,proto3" json:"content_type,omitempty"`
	// URL of this violating resource.
	ResourceUrl string `protobuf:"bytes,2,opt,name=resource_url,json=resourceUrl,proto3" json:"resource_url,omitempty"`
}

func (x *ViolatingResource) Reset() {
	*x = ViolatingResource{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_websecurityscanner_v1_finding_addon_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ViolatingResource) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ViolatingResource) ProtoMessage() {}

func (x *ViolatingResource) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_websecurityscanner_v1_finding_addon_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ViolatingResource.ProtoReflect.Descriptor instead.
func (*ViolatingResource) Descriptor() ([]byte, []int) {
	return file_google_cloud_websecurityscanner_v1_finding_addon_proto_rawDescGZIP(), []int{2}
}

func (x *ViolatingResource) GetContentType() string {
	if x != nil {
		return x.ContentType
	}
	return ""
}

func (x *ViolatingResource) GetResourceUrl() string {
	if x != nil {
		return x.ResourceUrl
	}
	return ""
}

// Information about vulnerable request parameters.
type VulnerableParameters struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The vulnerable parameter names.
	ParameterNames []string `protobuf:"bytes,1,rep,name=parameter_names,json=parameterNames,proto3" json:"parameter_names,omitempty"`
}

func (x *VulnerableParameters) Reset() {
	*x = VulnerableParameters{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_websecurityscanner_v1_finding_addon_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VulnerableParameters) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VulnerableParameters) ProtoMessage() {}

func (x *VulnerableParameters) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_websecurityscanner_v1_finding_addon_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VulnerableParameters.ProtoReflect.Descriptor instead.
func (*VulnerableParameters) Descriptor() ([]byte, []int) {
	return file_google_cloud_websecurityscanner_v1_finding_addon_proto_rawDescGZIP(), []int{3}
}

func (x *VulnerableParameters) GetParameterNames() []string {
	if x != nil {
		return x.ParameterNames
	}
	return nil
}

// Information about vulnerable or missing HTTP Headers.
type VulnerableHeaders struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// List of vulnerable headers.
	Headers []*VulnerableHeaders_Header `protobuf:"bytes,1,rep,name=headers,proto3" json:"headers,omitempty"`
	// List of missing headers.
	MissingHeaders []*VulnerableHeaders_Header `protobuf:"bytes,2,rep,name=missing_headers,json=missingHeaders,proto3" json:"missing_headers,omitempty"`
}

func (x *VulnerableHeaders) Reset() {
	*x = VulnerableHeaders{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_websecurityscanner_v1_finding_addon_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VulnerableHeaders) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VulnerableHeaders) ProtoMessage() {}

func (x *VulnerableHeaders) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_websecurityscanner_v1_finding_addon_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VulnerableHeaders.ProtoReflect.Descriptor instead.
func (*VulnerableHeaders) Descriptor() ([]byte, []int) {
	return file_google_cloud_websecurityscanner_v1_finding_addon_proto_rawDescGZIP(), []int{4}
}

func (x *VulnerableHeaders) GetHeaders() []*VulnerableHeaders_Header {
	if x != nil {
		return x.Headers
	}
	return nil
}

func (x *VulnerableHeaders) GetMissingHeaders() []*VulnerableHeaders_Header {
	if x != nil {
		return x.MissingHeaders
	}
	return nil
}

// Information reported for an XSS.
type Xss struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Stack traces leading to the point where the XSS occurred.
	StackTraces []string `protobuf:"bytes,1,rep,name=stack_traces,json=stackTraces,proto3" json:"stack_traces,omitempty"`
	// An error message generated by a javascript breakage.
	ErrorMessage string `protobuf:"bytes,2,opt,name=error_message,json=errorMessage,proto3" json:"error_message,omitempty"`
	// The attack vector of the payload triggering this XSS.
	AttackVector Xss_AttackVector `protobuf:"varint,3,opt,name=attack_vector,json=attackVector,proto3,enum=google.cloud.websecurityscanner.v1.Xss_AttackVector" json:"attack_vector,omitempty"`
	// The reproduction url for the seeding POST request of a Stored XSS.
	StoredXssSeedingUrl string `protobuf:"bytes,4,opt,name=stored_xss_seeding_url,json=storedXssSeedingUrl,proto3" json:"stored_xss_seeding_url,omitempty"`
}

func (x *Xss) Reset() {
	*x = Xss{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_websecurityscanner_v1_finding_addon_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Xss) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Xss) ProtoMessage() {}

func (x *Xss) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_websecurityscanner_v1_finding_addon_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Xss.ProtoReflect.Descriptor instead.
func (*Xss) Descriptor() ([]byte, []int) {
	return file_google_cloud_websecurityscanner_v1_finding_addon_proto_rawDescGZIP(), []int{5}
}

func (x *Xss) GetStackTraces() []string {
	if x != nil {
		return x.StackTraces
	}
	return nil
}

func (x *Xss) GetErrorMessage() string {
	if x != nil {
		return x.ErrorMessage
	}
	return ""
}

func (x *Xss) GetAttackVector() Xss_AttackVector {
	if x != nil {
		return x.AttackVector
	}
	return Xss_ATTACK_VECTOR_UNSPECIFIED
}

func (x *Xss) GetStoredXssSeedingUrl() string {
	if x != nil {
		return x.StoredXssSeedingUrl
	}
	return ""
}

// Information reported for an XXE.
type Xxe struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The XML string that triggered the XXE vulnerability. Non-payload values
	// might be redacted.
	PayloadValue string `protobuf:"bytes,1,opt,name=payload_value,json=payloadValue,proto3" json:"payload_value,omitempty"`
	// Location within the request where the payload was placed.
	PayloadLocation Xxe_Location `protobuf:"varint,2,opt,name=payload_location,json=payloadLocation,proto3,enum=google.cloud.websecurityscanner.v1.Xxe_Location" json:"payload_location,omitempty"`
}

func (x *Xxe) Reset() {
	*x = Xxe{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_websecurityscanner_v1_finding_addon_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Xxe) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Xxe) ProtoMessage() {}

func (x *Xxe) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_websecurityscanner_v1_finding_addon_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Xxe.ProtoReflect.Descriptor instead.
func (*Xxe) Descriptor() ([]byte, []int) {
	return file_google_cloud_websecurityscanner_v1_finding_addon_proto_rawDescGZIP(), []int{6}
}

func (x *Xxe) GetPayloadValue() string {
	if x != nil {
		return x.PayloadValue
	}
	return ""
}

func (x *Xxe) GetPayloadLocation() Xxe_Location {
	if x != nil {
		return x.PayloadLocation
	}
	return Xxe_LOCATION_UNSPECIFIED
}

// Describes a HTTP Header.
type VulnerableHeaders_Header struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Header name.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Header value.
	Value string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *VulnerableHeaders_Header) Reset() {
	*x = VulnerableHeaders_Header{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_websecurityscanner_v1_finding_addon_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VulnerableHeaders_Header) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VulnerableHeaders_Header) ProtoMessage() {}

func (x *VulnerableHeaders_Header) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_websecurityscanner_v1_finding_addon_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VulnerableHeaders_Header.ProtoReflect.Descriptor instead.
func (*VulnerableHeaders_Header) Descriptor() ([]byte, []int) {
	return file_google_cloud_websecurityscanner_v1_finding_addon_proto_rawDescGZIP(), []int{4, 0}
}

func (x *VulnerableHeaders_Header) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *VulnerableHeaders_Header) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

var File_google_cloud_websecurityscanner_v1_finding_addon_proto protoreflect.FileDescriptor

var file_google_cloud_websecurityscanner_v1_finding_addon_proto_rawDesc = []byte{
	0x0a, 0x36, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x77,
	0x65, 0x62, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x73, 0x63, 0x61, 0x6e, 0x6e, 0x65,
	0x72, 0x2f, 0x76, 0x31, 0x2f, 0x66, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x5f, 0x61, 0x64, 0x64,
	0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x22, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x77, 0x65, 0x62, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69,
	0x74, 0x79, 0x73, 0x63, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x22, 0x3d, 0x0a, 0x04,
	0x46, 0x6f, 0x72, 0x6d, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x75,
	0x72, 0x69, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x55, 0x72, 0x69, 0x12, 0x16, 0x0a, 0x06, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x06, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x22, 0x76, 0x0a, 0x0f, 0x4f,
	0x75, 0x74, 0x64, 0x61, 0x74, 0x65, 0x64, 0x4c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x12, 0x21,
	0x0a, 0x0c, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x26, 0x0a, 0x0f, 0x6c,
	0x65, 0x61, 0x72, 0x6e, 0x5f, 0x6d, 0x6f, 0x72, 0x65, 0x5f, 0x75, 0x72, 0x6c, 0x73, 0x18, 0x03,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x0d, 0x6c, 0x65, 0x61, 0x72, 0x6e, 0x4d, 0x6f, 0x72, 0x65, 0x55,
	0x72, 0x6c, 0x73, 0x22, 0x59, 0x0a, 0x11, 0x56, 0x69, 0x6f, 0x6c, 0x61, 0x74, 0x69, 0x6e, 0x67,
	0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x55, 0x72, 0x6c, 0x22, 0x3f,
	0x0a, 0x14, 0x56, 0x75, 0x6c, 0x6e, 0x65, 0x72, 0x61, 0x62, 0x6c, 0x65, 0x50, 0x61, 0x72, 0x61,
	0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x12, 0x27, 0x0a, 0x0f, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x65,
	0x74, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x0e, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x22,
	0x86, 0x02, 0x0a, 0x11, 0x56, 0x75, 0x6c, 0x6e, 0x65, 0x72, 0x61, 0x62, 0x6c, 0x65, 0x48, 0x65,
	0x61, 0x64, 0x65, 0x72, 0x73, 0x12, 0x56, 0x0a, 0x07, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x3c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x77, 0x65, 0x62, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74,
	0x79, 0x73, 0x63, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x56, 0x75, 0x6c, 0x6e,
	0x65, 0x72, 0x61, 0x62, 0x6c, 0x65, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x2e, 0x48, 0x65,
	0x61, 0x64, 0x65, 0x72, 0x52, 0x07, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x12, 0x65, 0x0a,
	0x0f, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x5f, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x3c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x77, 0x65, 0x62, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74,
	0x79, 0x73, 0x63, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x56, 0x75, 0x6c, 0x6e,
	0x65, 0x72, 0x61, 0x62, 0x6c, 0x65, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x2e, 0x48, 0x65,
	0x61, 0x64, 0x65, 0x72, 0x52, 0x0e, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x48, 0x65, 0x61,
	0x64, 0x65, 0x72, 0x73, 0x1a, 0x32, 0x0a, 0x06, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x99, 0x04, 0x0a, 0x03, 0x58, 0x73, 0x73,
	0x12, 0x21, 0x0a, 0x0c, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x5f, 0x74, 0x72, 0x61, 0x63, 0x65, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0b, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x54, 0x72, 0x61,
	0x63, 0x65, 0x73, 0x12, 0x23, 0x0a, 0x0d, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x59, 0x0a, 0x0d, 0x61, 0x74, 0x74, 0x61,
	0x63, 0x6b, 0x5f, 0x76, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x34, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x77,
	0x65, 0x62, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x73, 0x63, 0x61, 0x6e, 0x6e, 0x65,
	0x72, 0x2e, 0x76, 0x31, 0x2e, 0x58, 0x73, 0x73, 0x2e, 0x41, 0x74, 0x74, 0x61, 0x63, 0x6b, 0x56,
	0x65, 0x63, 0x74, 0x6f, 0x72, 0x52, 0x0c, 0x61, 0x74, 0x74, 0x61, 0x63, 0x6b, 0x56, 0x65, 0x63,
	0x74, 0x6f, 0x72, 0x12, 0x33, 0x0a, 0x16, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x64, 0x5f, 0x78, 0x73,
	0x73, 0x5f, 0x73, 0x65, 0x65, 0x64, 0x69, 0x6e, 0x67, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x13, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x64, 0x58, 0x73, 0x73, 0x53, 0x65,
	0x65, 0x64, 0x69, 0x6e, 0x67, 0x55, 0x72, 0x6c, 0x22, 0xb9, 0x02, 0x0a, 0x0c, 0x41, 0x74, 0x74,
	0x61, 0x63, 0x6b, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x12, 0x1d, 0x0a, 0x19, 0x41, 0x54, 0x54,
	0x41, 0x43, 0x4b, 0x5f, 0x56, 0x45, 0x43, 0x54, 0x4f, 0x52, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45,
	0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x11, 0x0a, 0x0d, 0x4c, 0x4f, 0x43, 0x41,
	0x4c, 0x5f, 0x53, 0x54, 0x4f, 0x52, 0x41, 0x47, 0x45, 0x10, 0x01, 0x12, 0x13, 0x0a, 0x0f, 0x53,
	0x45, 0x53, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x53, 0x54, 0x4f, 0x52, 0x41, 0x47, 0x45, 0x10, 0x02,
	0x12, 0x0f, 0x0a, 0x0b, 0x57, 0x49, 0x4e, 0x44, 0x4f, 0x57, 0x5f, 0x4e, 0x41, 0x4d, 0x45, 0x10,
	0x03, 0x12, 0x0c, 0x0a, 0x08, 0x52, 0x45, 0x46, 0x45, 0x52, 0x52, 0x45, 0x52, 0x10, 0x04, 0x12,
	0x0e, 0x0a, 0x0a, 0x46, 0x4f, 0x52, 0x4d, 0x5f, 0x49, 0x4e, 0x50, 0x55, 0x54, 0x10, 0x05, 0x12,
	0x0a, 0x0a, 0x06, 0x43, 0x4f, 0x4f, 0x4b, 0x49, 0x45, 0x10, 0x06, 0x12, 0x10, 0x0a, 0x0c, 0x50,
	0x4f, 0x53, 0x54, 0x5f, 0x4d, 0x45, 0x53, 0x53, 0x41, 0x47, 0x45, 0x10, 0x07, 0x12, 0x12, 0x0a,
	0x0e, 0x47, 0x45, 0x54, 0x5f, 0x50, 0x41, 0x52, 0x41, 0x4d, 0x45, 0x54, 0x45, 0x52, 0x53, 0x10,
	0x08, 0x12, 0x10, 0x0a, 0x0c, 0x55, 0x52, 0x4c, 0x5f, 0x46, 0x52, 0x41, 0x47, 0x4d, 0x45, 0x4e,
	0x54, 0x10, 0x09, 0x12, 0x10, 0x0a, 0x0c, 0x48, 0x54, 0x4d, 0x4c, 0x5f, 0x43, 0x4f, 0x4d, 0x4d,
	0x45, 0x4e, 0x54, 0x10, 0x0a, 0x12, 0x13, 0x0a, 0x0f, 0x50, 0x4f, 0x53, 0x54, 0x5f, 0x50, 0x41,
	0x52, 0x41, 0x4d, 0x45, 0x54, 0x45, 0x52, 0x53, 0x10, 0x0b, 0x12, 0x0c, 0x0a, 0x08, 0x50, 0x52,
	0x4f, 0x54, 0x4f, 0x43, 0x4f, 0x4c, 0x10, 0x0c, 0x12, 0x0e, 0x0a, 0x0a, 0x53, 0x54, 0x4f, 0x52,
	0x45, 0x44, 0x5f, 0x58, 0x53, 0x53, 0x10, 0x0d, 0x12, 0x0f, 0x0a, 0x0b, 0x53, 0x41, 0x4d, 0x45,
	0x5f, 0x4f, 0x52, 0x49, 0x47, 0x49, 0x4e, 0x10, 0x0e, 0x12, 0x19, 0x0a, 0x15, 0x55, 0x53, 0x45,
	0x52, 0x5f, 0x43, 0x4f, 0x4e, 0x54, 0x52, 0x4f, 0x4c, 0x4c, 0x41, 0x42, 0x4c, 0x45, 0x5f, 0x55,
	0x52, 0x4c, 0x10, 0x0f, 0x22, 0xc8, 0x01, 0x0a, 0x03, 0x58, 0x78, 0x65, 0x12, 0x23, 0x0a, 0x0d,
	0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0c, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x12, 0x5b, 0x0a, 0x10, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x6c, 0x6f, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x30, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x77, 0x65, 0x62, 0x73, 0x65,
	0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x73, 0x63, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x2e, 0x76, 0x31,
	0x2e, 0x58, 0x78, 0x65, 0x2e, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0f, 0x70,
	0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x3f,
	0x0a, 0x08, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x14, 0x4c, 0x4f,
	0x43, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49,
	0x45, 0x44, 0x10, 0x00, 0x12, 0x19, 0x0a, 0x15, 0x43, 0x4f, 0x4d, 0x50, 0x4c, 0x45, 0x54, 0x45,
	0x5f, 0x52, 0x45, 0x51, 0x55, 0x45, 0x53, 0x54, 0x5f, 0x42, 0x4f, 0x44, 0x59, 0x10, 0x01, 0x42,
	0x87, 0x02, 0x0a, 0x26, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x77, 0x65, 0x62, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79,
	0x73, 0x63, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x42, 0x11, 0x46, 0x69, 0x6e, 0x64,
	0x69, 0x6e, 0x67, 0x41, 0x64, 0x64, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a,
	0x56, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x67, 0x6f, 0x2f, 0x77, 0x65, 0x62, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79,
	0x73, 0x63, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x69, 0x76, 0x31, 0x2f, 0x77, 0x65,
	0x62, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x73, 0x63, 0x61, 0x6e, 0x6e, 0x65, 0x72,
	0x70, 0x62, 0x3b, 0x77, 0x65, 0x62, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x73, 0x63,
	0x61, 0x6e, 0x6e, 0x65, 0x72, 0x70, 0x62, 0xaa, 0x02, 0x22, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x57, 0x65, 0x62, 0x53, 0x65, 0x63, 0x75, 0x72, 0x69,
	0x74, 0x79, 0x53, 0x63, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x22, 0x47,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x57, 0x65, 0x62, 0x53,
	0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x53, 0x63, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x5c, 0x56,
	0x31, 0xea, 0x02, 0x25, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x43, 0x6c, 0x6f, 0x75,
	0x64, 0x3a, 0x3a, 0x57, 0x65, 0x62, 0x53, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x53, 0x63,
	0x61, 0x6e, 0x6e, 0x65, 0x72, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_google_cloud_websecurityscanner_v1_finding_addon_proto_rawDescOnce sync.Once
	file_google_cloud_websecurityscanner_v1_finding_addon_proto_rawDescData = file_google_cloud_websecurityscanner_v1_finding_addon_proto_rawDesc
)

func file_google_cloud_websecurityscanner_v1_finding_addon_proto_rawDescGZIP() []byte {
	file_google_cloud_websecurityscanner_v1_finding_addon_proto_rawDescOnce.Do(func() {
		file_google_cloud_websecurityscanner_v1_finding_addon_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_websecurityscanner_v1_finding_addon_proto_rawDescData)
	})
	return file_google_cloud_websecurityscanner_v1_finding_addon_proto_rawDescData
}

var file_google_cloud_websecurityscanner_v1_finding_addon_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_google_cloud_websecurityscanner_v1_finding_addon_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_google_cloud_websecurityscanner_v1_finding_addon_proto_goTypes = []interface{}{
	(Xss_AttackVector)(0),            // 0: google.cloud.websecurityscanner.v1.Xss.AttackVector
	(Xxe_Location)(0),                // 1: google.cloud.websecurityscanner.v1.Xxe.Location
	(*Form)(nil),                     // 2: google.cloud.websecurityscanner.v1.Form
	(*OutdatedLibrary)(nil),          // 3: google.cloud.websecurityscanner.v1.OutdatedLibrary
	(*ViolatingResource)(nil),        // 4: google.cloud.websecurityscanner.v1.ViolatingResource
	(*VulnerableParameters)(nil),     // 5: google.cloud.websecurityscanner.v1.VulnerableParameters
	(*VulnerableHeaders)(nil),        // 6: google.cloud.websecurityscanner.v1.VulnerableHeaders
	(*Xss)(nil),                      // 7: google.cloud.websecurityscanner.v1.Xss
	(*Xxe)(nil),                      // 8: google.cloud.websecurityscanner.v1.Xxe
	(*VulnerableHeaders_Header)(nil), // 9: google.cloud.websecurityscanner.v1.VulnerableHeaders.Header
}
var file_google_cloud_websecurityscanner_v1_finding_addon_proto_depIdxs = []int32{
	9, // 0: google.cloud.websecurityscanner.v1.VulnerableHeaders.headers:type_name -> google.cloud.websecurityscanner.v1.VulnerableHeaders.Header
	9, // 1: google.cloud.websecurityscanner.v1.VulnerableHeaders.missing_headers:type_name -> google.cloud.websecurityscanner.v1.VulnerableHeaders.Header
	0, // 2: google.cloud.websecurityscanner.v1.Xss.attack_vector:type_name -> google.cloud.websecurityscanner.v1.Xss.AttackVector
	1, // 3: google.cloud.websecurityscanner.v1.Xxe.payload_location:type_name -> google.cloud.websecurityscanner.v1.Xxe.Location
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_google_cloud_websecurityscanner_v1_finding_addon_proto_init() }
func file_google_cloud_websecurityscanner_v1_finding_addon_proto_init() {
	if File_google_cloud_websecurityscanner_v1_finding_addon_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_cloud_websecurityscanner_v1_finding_addon_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Form); i {
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
		file_google_cloud_websecurityscanner_v1_finding_addon_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OutdatedLibrary); i {
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
		file_google_cloud_websecurityscanner_v1_finding_addon_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ViolatingResource); i {
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
		file_google_cloud_websecurityscanner_v1_finding_addon_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VulnerableParameters); i {
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
		file_google_cloud_websecurityscanner_v1_finding_addon_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VulnerableHeaders); i {
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
		file_google_cloud_websecurityscanner_v1_finding_addon_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Xss); i {
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
		file_google_cloud_websecurityscanner_v1_finding_addon_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Xxe); i {
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
		file_google_cloud_websecurityscanner_v1_finding_addon_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VulnerableHeaders_Header); i {
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
			RawDescriptor: file_google_cloud_websecurityscanner_v1_finding_addon_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_cloud_websecurityscanner_v1_finding_addon_proto_goTypes,
		DependencyIndexes: file_google_cloud_websecurityscanner_v1_finding_addon_proto_depIdxs,
		EnumInfos:         file_google_cloud_websecurityscanner_v1_finding_addon_proto_enumTypes,
		MessageInfos:      file_google_cloud_websecurityscanner_v1_finding_addon_proto_msgTypes,
	}.Build()
	File_google_cloud_websecurityscanner_v1_finding_addon_proto = out.File
	file_google_cloud_websecurityscanner_v1_finding_addon_proto_rawDesc = nil
	file_google_cloud_websecurityscanner_v1_finding_addon_proto_goTypes = nil
	file_google_cloud_websecurityscanner_v1_finding_addon_proto_depIdxs = nil
}
