// Copyright 2023 Google LLC
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
// source: google/cloud/aiplatform/v1/user_action_reference.proto

package aiplatformpb

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

// References an API call. It contains more information about long running
// operation and Jobs that are triggered by the API call.
type UserActionReference struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Reference:
	//	*UserActionReference_Operation
	//	*UserActionReference_DataLabelingJob
	Reference isUserActionReference_Reference `protobuf_oneof:"reference"`
	// The method name of the API RPC call. For example,
	// "/google.cloud.aiplatform.{apiVersion}.DatasetService.CreateDataset"
	Method string `protobuf:"bytes,3,opt,name=method,proto3" json:"method,omitempty"`
}

func (x *UserActionReference) Reset() {
	*x = UserActionReference{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_aiplatform_v1_user_action_reference_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserActionReference) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserActionReference) ProtoMessage() {}

func (x *UserActionReference) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_aiplatform_v1_user_action_reference_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserActionReference.ProtoReflect.Descriptor instead.
func (*UserActionReference) Descriptor() ([]byte, []int) {
	return file_google_cloud_aiplatform_v1_user_action_reference_proto_rawDescGZIP(), []int{0}
}

func (m *UserActionReference) GetReference() isUserActionReference_Reference {
	if m != nil {
		return m.Reference
	}
	return nil
}

func (x *UserActionReference) GetOperation() string {
	if x, ok := x.GetReference().(*UserActionReference_Operation); ok {
		return x.Operation
	}
	return ""
}

func (x *UserActionReference) GetDataLabelingJob() string {
	if x, ok := x.GetReference().(*UserActionReference_DataLabelingJob); ok {
		return x.DataLabelingJob
	}
	return ""
}

func (x *UserActionReference) GetMethod() string {
	if x != nil {
		return x.Method
	}
	return ""
}

type isUserActionReference_Reference interface {
	isUserActionReference_Reference()
}

type UserActionReference_Operation struct {
	// For API calls that return a long running operation.
	// Resource name of the long running operation.
	// Format:
	// `projects/{project}/locations/{location}/operations/{operation}`
	Operation string `protobuf:"bytes,1,opt,name=operation,proto3,oneof"`
}

type UserActionReference_DataLabelingJob struct {
	// For API calls that start a LabelingJob.
	// Resource name of the LabelingJob.
	// Format:
	// `projects/{project}/locations/{location}/dataLabelingJobs/{data_labeling_job}`
	DataLabelingJob string `protobuf:"bytes,2,opt,name=data_labeling_job,json=dataLabelingJob,proto3,oneof"`
}

func (*UserActionReference_Operation) isUserActionReference_Reference() {}

func (*UserActionReference_DataLabelingJob) isUserActionReference_Reference() {}

var File_google_cloud_aiplatform_v1_user_action_reference_proto protoreflect.FileDescriptor

var file_google_cloud_aiplatform_v1_user_action_reference_proto_rawDesc = []byte{
	0x0a, 0x36, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x61,
	0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e,
	0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1a, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72,
	0x6d, 0x2e, 0x76, 0x31, 0x22, 0x88, 0x01, 0x0a, 0x13, 0x55, 0x73, 0x65, 0x72, 0x41, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x1e, 0x0a, 0x09,
	0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48,
	0x00, 0x52, 0x09, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2c, 0x0a, 0x11,
	0x64, 0x61, 0x74, 0x61, 0x5f, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x69, 0x6e, 0x67, 0x5f, 0x6a, 0x6f,
	0x62, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0f, 0x64, 0x61, 0x74, 0x61, 0x4c,
	0x61, 0x62, 0x65, 0x6c, 0x69, 0x6e, 0x67, 0x4a, 0x6f, 0x62, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x65,
	0x74, 0x68, 0x6f, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x65, 0x74, 0x68,
	0x6f, 0x64, 0x42, 0x0b, 0x0a, 0x09, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x42,
	0xd6, 0x01, 0x0a, 0x1e, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e,
	0x76, 0x31, 0x42, 0x18, 0x55, 0x73, 0x65, 0x72, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x67, 0x6f, 0x2f, 0x61, 0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x61,
	0x70, 0x69, 0x76, 0x31, 0x2f, 0x61, 0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x70,
	0x62, 0x3b, 0x61, 0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x70, 0x62, 0xaa, 0x02,
	0x1a, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x41, 0x49,
	0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x1a, 0x47, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x41, 0x49, 0x50, 0x6c, 0x61,
	0x74, 0x66, 0x6f, 0x72, 0x6d, 0x5c, 0x56, 0x31, 0xea, 0x02, 0x1d, 0x47, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x3a, 0x3a, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x3a, 0x3a, 0x41, 0x49, 0x50, 0x6c, 0x61, 0x74,
	0x66, 0x6f, 0x72, 0x6d, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_cloud_aiplatform_v1_user_action_reference_proto_rawDescOnce sync.Once
	file_google_cloud_aiplatform_v1_user_action_reference_proto_rawDescData = file_google_cloud_aiplatform_v1_user_action_reference_proto_rawDesc
)

func file_google_cloud_aiplatform_v1_user_action_reference_proto_rawDescGZIP() []byte {
	file_google_cloud_aiplatform_v1_user_action_reference_proto_rawDescOnce.Do(func() {
		file_google_cloud_aiplatform_v1_user_action_reference_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_aiplatform_v1_user_action_reference_proto_rawDescData)
	})
	return file_google_cloud_aiplatform_v1_user_action_reference_proto_rawDescData
}

var file_google_cloud_aiplatform_v1_user_action_reference_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_cloud_aiplatform_v1_user_action_reference_proto_goTypes = []interface{}{
	(*UserActionReference)(nil), // 0: google.cloud.aiplatform.v1.UserActionReference
}
var file_google_cloud_aiplatform_v1_user_action_reference_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_google_cloud_aiplatform_v1_user_action_reference_proto_init() }
func file_google_cloud_aiplatform_v1_user_action_reference_proto_init() {
	if File_google_cloud_aiplatform_v1_user_action_reference_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_cloud_aiplatform_v1_user_action_reference_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserActionReference); i {
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
	file_google_cloud_aiplatform_v1_user_action_reference_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*UserActionReference_Operation)(nil),
		(*UserActionReference_DataLabelingJob)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_cloud_aiplatform_v1_user_action_reference_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_cloud_aiplatform_v1_user_action_reference_proto_goTypes,
		DependencyIndexes: file_google_cloud_aiplatform_v1_user_action_reference_proto_depIdxs,
		MessageInfos:      file_google_cloud_aiplatform_v1_user_action_reference_proto_msgTypes,
	}.Build()
	File_google_cloud_aiplatform_v1_user_action_reference_proto = out.File
	file_google_cloud_aiplatform_v1_user_action_reference_proto_rawDesc = nil
	file_google_cloud_aiplatform_v1_user_action_reference_proto_goTypes = nil
	file_google_cloud_aiplatform_v1_user_action_reference_proto_depIdxs = nil
}
