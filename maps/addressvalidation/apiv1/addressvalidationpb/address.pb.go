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
// source: google/maps/addressvalidation/v1/address.proto

package addressvalidationpb

import (
	reflect "reflect"
	sync "sync"

	_ "google.golang.org/genproto/googleapis/api/annotations"
	postaladdress "google.golang.org/genproto/googleapis/type/postaladdress"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// The different possible values for confirmation levels.
type AddressComponent_ConfirmationLevel int32

const (
	// Default value. This value is unused.
	AddressComponent_CONFIRMATION_LEVEL_UNSPECIFIED AddressComponent_ConfirmationLevel = 0
	// We were able to verify that this component exists and makes sense in the
	// context of the rest of the address.
	AddressComponent_CONFIRMED AddressComponent_ConfirmationLevel = 1
	// This component could not be confirmed, but it is plausible that it
	// exists. For example, a street number within a known valid range of
	// numbers on a street where specific house numbers are not known.
	AddressComponent_UNCONFIRMED_BUT_PLAUSIBLE AddressComponent_ConfirmationLevel = 2
	// This component was not confirmed and is likely to be wrong. For
	// example, a neighborhood that does not fit the rest of the address.
	AddressComponent_UNCONFIRMED_AND_SUSPICIOUS AddressComponent_ConfirmationLevel = 3
)

// Enum value maps for AddressComponent_ConfirmationLevel.
var (
	AddressComponent_ConfirmationLevel_name = map[int32]string{
		0: "CONFIRMATION_LEVEL_UNSPECIFIED",
		1: "CONFIRMED",
		2: "UNCONFIRMED_BUT_PLAUSIBLE",
		3: "UNCONFIRMED_AND_SUSPICIOUS",
	}
	AddressComponent_ConfirmationLevel_value = map[string]int32{
		"CONFIRMATION_LEVEL_UNSPECIFIED": 0,
		"CONFIRMED":                      1,
		"UNCONFIRMED_BUT_PLAUSIBLE":      2,
		"UNCONFIRMED_AND_SUSPICIOUS":     3,
	}
)

func (x AddressComponent_ConfirmationLevel) Enum() *AddressComponent_ConfirmationLevel {
	p := new(AddressComponent_ConfirmationLevel)
	*p = x
	return p
}

func (x AddressComponent_ConfirmationLevel) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AddressComponent_ConfirmationLevel) Descriptor() protoreflect.EnumDescriptor {
	return file_google_maps_addressvalidation_v1_address_proto_enumTypes[0].Descriptor()
}

func (AddressComponent_ConfirmationLevel) Type() protoreflect.EnumType {
	return &file_google_maps_addressvalidation_v1_address_proto_enumTypes[0]
}

func (x AddressComponent_ConfirmationLevel) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AddressComponent_ConfirmationLevel.Descriptor instead.
func (AddressComponent_ConfirmationLevel) EnumDescriptor() ([]byte, []int) {
	return file_google_maps_addressvalidation_v1_address_proto_rawDescGZIP(), []int{1, 0}
}

// Details of the post-processed address. Post-processing includes
// correcting misspelled parts of the address, replacing incorrect parts, and
// inferring missing parts.
type Address struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The post-processed address, formatted as a single-line address following
	// the address formatting rules of the region where the address is located.
	FormattedAddress string `protobuf:"bytes,2,opt,name=formatted_address,json=formattedAddress,proto3" json:"formatted_address,omitempty"`
	// The validated address represented as a postal address.
	PostalAddress *postaladdress.PostalAddress `protobuf:"bytes,3,opt,name=postal_address,json=postalAddress,proto3" json:"postal_address,omitempty"`
	// Unordered list. The individual address components of the formatted and
	// corrected address, along with validation information. This provides
	// information on the validation status of the individual components.
	//
	// Address components are not ordered in a particular way. Do not make any
	// assumptions on the ordering of the address components in the list.
	AddressComponents []*AddressComponent `protobuf:"bytes,4,rep,name=address_components,json=addressComponents,proto3" json:"address_components,omitempty"`
	// The types of components that were expected to be present in a correctly
	// formatted mailing address but were not found in the input AND could
	// not be inferred. Components of this type are not present in
	// `formatted_address`, `postal_address`, or `address_components`. An
	// example might be `['street_number', 'route']` for an input like
	// "Boulder, Colorado, 80301, USA". The list of possible types can be found
	// [here](https://developers.google.com/maps/documentation/geocoding/requests-geocoding#Types).
	MissingComponentTypes []string `protobuf:"bytes,5,rep,name=missing_component_types,json=missingComponentTypes,proto3" json:"missing_component_types,omitempty"`
	// The types of the components that are present in the `address_components`
	// but could not be confirmed to be correct. This field is provided for the
	// sake of convenience: its contents are equivalent to iterating through the
	// `address_components` to find the types of all the components where the
	// [confirmation_level][google.maps.addressvalidation.v1.AddressComponent.confirmation_level]
	// is not
	// [CONFIRMED][google.maps.addressvalidation.v1.AddressComponent.ConfirmationLevel.CONFIRMED]
	// or the
	// [inferred][google.maps.addressvalidation.v1.AddressComponent.inferred]
	// flag is not set to `true`. The list of possible types can be found
	// [here](https://developers.google.com/maps/documentation/geocoding/requests-geocoding#Types).
	UnconfirmedComponentTypes []string `protobuf:"bytes,6,rep,name=unconfirmed_component_types,json=unconfirmedComponentTypes,proto3" json:"unconfirmed_component_types,omitempty"`
	// Any tokens in the input that could not be resolved. This might be an
	// input that was not recognized as a valid part of an address (for example
	// in an input like "123235253253 Main St, San Francisco, CA, 94105", the
	// unresolved tokens may look like `["123235253253"]` since that does not
	// look like a valid street number.
	UnresolvedTokens []string `protobuf:"bytes,7,rep,name=unresolved_tokens,json=unresolvedTokens,proto3" json:"unresolved_tokens,omitempty"`
}

func (x *Address) Reset() {
	*x = Address{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_maps_addressvalidation_v1_address_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Address) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Address) ProtoMessage() {}

func (x *Address) ProtoReflect() protoreflect.Message {
	mi := &file_google_maps_addressvalidation_v1_address_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Address.ProtoReflect.Descriptor instead.
func (*Address) Descriptor() ([]byte, []int) {
	return file_google_maps_addressvalidation_v1_address_proto_rawDescGZIP(), []int{0}
}

func (x *Address) GetFormattedAddress() string {
	if x != nil {
		return x.FormattedAddress
	}
	return ""
}

func (x *Address) GetPostalAddress() *postaladdress.PostalAddress {
	if x != nil {
		return x.PostalAddress
	}
	return nil
}

func (x *Address) GetAddressComponents() []*AddressComponent {
	if x != nil {
		return x.AddressComponents
	}
	return nil
}

func (x *Address) GetMissingComponentTypes() []string {
	if x != nil {
		return x.MissingComponentTypes
	}
	return nil
}

func (x *Address) GetUnconfirmedComponentTypes() []string {
	if x != nil {
		return x.UnconfirmedComponentTypes
	}
	return nil
}

func (x *Address) GetUnresolvedTokens() []string {
	if x != nil {
		return x.UnresolvedTokens
	}
	return nil
}

// Represents an address component, such as a street, city, or state.
type AddressComponent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The name for this component.
	ComponentName *ComponentName `protobuf:"bytes,1,opt,name=component_name,json=componentName,proto3" json:"component_name,omitempty"`
	// The type of the address component. See
	// [Table 2: Additional types returned by the Places
	// service](https://developers.google.com/places/web-service/supported_types#table2)
	// for a list of possible types.
	ComponentType string `protobuf:"bytes,2,opt,name=component_type,json=componentType,proto3" json:"component_type,omitempty"`
	// Indicates the level of certainty that we have that the component
	// is correct.
	ConfirmationLevel AddressComponent_ConfirmationLevel `protobuf:"varint,3,opt,name=confirmation_level,json=confirmationLevel,proto3,enum=google.maps.addressvalidation.v1.AddressComponent_ConfirmationLevel" json:"confirmation_level,omitempty"`
	// Indicates that the component was not part of the input, but we
	// inferred it for the address location and believe it should be provided
	// for a complete address.
	Inferred bool `protobuf:"varint,4,opt,name=inferred,proto3" json:"inferred,omitempty"`
	// Indicates the spelling of the component name was corrected in a minor way,
	// for example by switching two characters that appeared in the wrong order.
	// This indicates a cosmetic change.
	SpellCorrected bool `protobuf:"varint,5,opt,name=spell_corrected,json=spellCorrected,proto3" json:"spell_corrected,omitempty"`
	// Indicates the name of the component was replaced with a completely
	// different one, for example a wrong postal code being replaced with one that
	// is correct for the address. This is not a cosmetic change, the input
	// component has been changed to a different one.
	Replaced bool `protobuf:"varint,6,opt,name=replaced,proto3" json:"replaced,omitempty"`
	// Indicates an address component that is not expected to be present in a
	// postal address for the given region. We have retained it only because it
	// was part of the input.
	Unexpected bool `protobuf:"varint,7,opt,name=unexpected,proto3" json:"unexpected,omitempty"`
}

func (x *AddressComponent) Reset() {
	*x = AddressComponent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_maps_addressvalidation_v1_address_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddressComponent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddressComponent) ProtoMessage() {}

func (x *AddressComponent) ProtoReflect() protoreflect.Message {
	mi := &file_google_maps_addressvalidation_v1_address_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddressComponent.ProtoReflect.Descriptor instead.
func (*AddressComponent) Descriptor() ([]byte, []int) {
	return file_google_maps_addressvalidation_v1_address_proto_rawDescGZIP(), []int{1}
}

func (x *AddressComponent) GetComponentName() *ComponentName {
	if x != nil {
		return x.ComponentName
	}
	return nil
}

func (x *AddressComponent) GetComponentType() string {
	if x != nil {
		return x.ComponentType
	}
	return ""
}

func (x *AddressComponent) GetConfirmationLevel() AddressComponent_ConfirmationLevel {
	if x != nil {
		return x.ConfirmationLevel
	}
	return AddressComponent_CONFIRMATION_LEVEL_UNSPECIFIED
}

func (x *AddressComponent) GetInferred() bool {
	if x != nil {
		return x.Inferred
	}
	return false
}

func (x *AddressComponent) GetSpellCorrected() bool {
	if x != nil {
		return x.SpellCorrected
	}
	return false
}

func (x *AddressComponent) GetReplaced() bool {
	if x != nil {
		return x.Replaced
	}
	return false
}

func (x *AddressComponent) GetUnexpected() bool {
	if x != nil {
		return x.Unexpected
	}
	return false
}

// A wrapper for the name of the component.
type ComponentName struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The name text. For example, "5th Avenue" for a street name or "1253" for a
	// street number.
	Text string `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
	// The BCP-47 language code. This will not be present if the component name is
	// not associated with a language, such as a street number.
	LanguageCode string `protobuf:"bytes,2,opt,name=language_code,json=languageCode,proto3" json:"language_code,omitempty"`
}

func (x *ComponentName) Reset() {
	*x = ComponentName{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_maps_addressvalidation_v1_address_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ComponentName) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ComponentName) ProtoMessage() {}

func (x *ComponentName) ProtoReflect() protoreflect.Message {
	mi := &file_google_maps_addressvalidation_v1_address_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ComponentName.ProtoReflect.Descriptor instead.
func (*ComponentName) Descriptor() ([]byte, []int) {
	return file_google_maps_addressvalidation_v1_address_proto_rawDescGZIP(), []int{2}
}

func (x *ComponentName) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *ComponentName) GetLanguageCode() string {
	if x != nil {
		return x.LanguageCode
	}
	return ""
}

var File_google_maps_addressvalidation_v1_address_proto protoreflect.FileDescriptor

var file_google_maps_addressvalidation_v1_address_proto_rawDesc = []byte{
	0x0a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x6d, 0x61, 0x70, 0x73, 0x2f, 0x61, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f,
	0x76, 0x31, 0x2f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x6d, 0x61, 0x70, 0x73, 0x2e, 0x61, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66,
	0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x74, 0x79, 0x70, 0x65,
	0x2f, 0x70, 0x6f, 0x73, 0x74, 0x61, 0x6c, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x86, 0x03, 0x0a, 0x07, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x12, 0x2b, 0x0a, 0x11, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x74, 0x65, 0x64, 0x5f, 0x61,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x66, 0x6f,
	0x72, 0x6d, 0x61, 0x74, 0x74, 0x65, 0x64, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x41,
	0x0a, 0x0e, 0x70, 0x6f, 0x73, 0x74, 0x61, 0x6c, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x74, 0x79, 0x70, 0x65, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x61, 0x6c, 0x41, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x52, 0x0d, 0x70, 0x6f, 0x73, 0x74, 0x61, 0x6c, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x12, 0x66, 0x0a, 0x12, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x5f, 0x63, 0x6f, 0x6d,
	0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x32, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x6d, 0x61, 0x70, 0x73, 0x2e, 0x61, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31,
	0x2e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e,
	0x74, 0x42, 0x03, 0xe0, 0x41, 0x06, 0x52, 0x11, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x43,
	0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x36, 0x0a, 0x17, 0x6d, 0x69, 0x73,
	0x73, 0x69, 0x6e, 0x67, 0x5f, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x5f, 0x74,
	0x79, 0x70, 0x65, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x52, 0x15, 0x6d, 0x69, 0x73, 0x73,
	0x69, 0x6e, 0x67, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65,
	0x73, 0x12, 0x3e, 0x0a, 0x1b, 0x75, 0x6e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x65, 0x64,
	0x5f, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x73,
	0x18, 0x06, 0x20, 0x03, 0x28, 0x09, 0x52, 0x19, 0x75, 0x6e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x72,
	0x6d, 0x65, 0x64, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65,
	0x73, 0x12, 0x2b, 0x0a, 0x11, 0x75, 0x6e, 0x72, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x64, 0x5f,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x09, 0x52, 0x10, 0x75, 0x6e,
	0x72, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x64, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x22, 0x8f,
	0x04, 0x0a, 0x10, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x6e,
	0x65, 0x6e, 0x74, 0x12, 0x56, 0x0a, 0x0e, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2f, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x6d, 0x61, 0x70, 0x73, 0x2e, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x43,
	0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x0d, 0x63, 0x6f,
	0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x63,
	0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0d, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x73, 0x0a, 0x12, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x44,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x6d, 0x61, 0x70, 0x73, 0x2e, 0x61, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76,
	0x31, 0x2e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65,
	0x6e, 0x74, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4c,
	0x65, 0x76, 0x65, 0x6c, 0x52, 0x11, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x69, 0x6e, 0x66, 0x65, 0x72,
	0x72, 0x65, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x69, 0x6e, 0x66, 0x65, 0x72,
	0x72, 0x65, 0x64, 0x12, 0x27, 0x0a, 0x0f, 0x73, 0x70, 0x65, 0x6c, 0x6c, 0x5f, 0x63, 0x6f, 0x72,
	0x72, 0x65, 0x63, 0x74, 0x65, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0e, 0x73, 0x70,
	0x65, 0x6c, 0x6c, 0x43, 0x6f, 0x72, 0x72, 0x65, 0x63, 0x74, 0x65, 0x64, 0x12, 0x1a, 0x0a, 0x08,
	0x72, 0x65, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08,
	0x72, 0x65, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x75, 0x6e, 0x65, 0x78,
	0x70, 0x65, 0x63, 0x74, 0x65, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x75, 0x6e,
	0x65, 0x78, 0x70, 0x65, 0x63, 0x74, 0x65, 0x64, 0x22, 0x85, 0x01, 0x0a, 0x11, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x22,
	0x0a, 0x1e, 0x43, 0x4f, 0x4e, 0x46, 0x49, 0x52, 0x4d, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x4c,
	0x45, 0x56, 0x45, 0x4c, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44,
	0x10, 0x00, 0x12, 0x0d, 0x0a, 0x09, 0x43, 0x4f, 0x4e, 0x46, 0x49, 0x52, 0x4d, 0x45, 0x44, 0x10,
	0x01, 0x12, 0x1d, 0x0a, 0x19, 0x55, 0x4e, 0x43, 0x4f, 0x4e, 0x46, 0x49, 0x52, 0x4d, 0x45, 0x44,
	0x5f, 0x42, 0x55, 0x54, 0x5f, 0x50, 0x4c, 0x41, 0x55, 0x53, 0x49, 0x42, 0x4c, 0x45, 0x10, 0x02,
	0x12, 0x1e, 0x0a, 0x1a, 0x55, 0x4e, 0x43, 0x4f, 0x4e, 0x46, 0x49, 0x52, 0x4d, 0x45, 0x44, 0x5f,
	0x41, 0x4e, 0x44, 0x5f, 0x53, 0x55, 0x53, 0x50, 0x49, 0x43, 0x49, 0x4f, 0x55, 0x53, 0x10, 0x03,
	0x22, 0x48, 0x0a, 0x0d, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x74, 0x65, 0x78, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67,
	0x65, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x6c, 0x61,
	0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x42, 0x89, 0x02, 0x0a, 0x24, 0x63,
	0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x6d, 0x61, 0x70, 0x73, 0x2e, 0x61,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x76, 0x31, 0x42, 0x0c, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x50, 0x01, 0x5a, 0x58, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x2f, 0x6d, 0x61, 0x70, 0x73, 0x2f, 0x61, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f,
	0x61, 0x70, 0x69, 0x76, 0x31, 0x2f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x76, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x70, 0x62, 0x3b, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x70, 0x62, 0xf8, 0x01, 0x01,
	0xa2, 0x02, 0x07, 0x47, 0x4d, 0x50, 0x41, 0x56, 0x56, 0x31, 0xaa, 0x02, 0x20, 0x47, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x4d, 0x61, 0x70, 0x73, 0x2e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x20,
	0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x4d, 0x61, 0x70, 0x73, 0x5c, 0x41, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5c, 0x56, 0x31,
	0xea, 0x02, 0x23, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x4d, 0x61, 0x70, 0x73, 0x3a,
	0x3a, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_maps_addressvalidation_v1_address_proto_rawDescOnce sync.Once
	file_google_maps_addressvalidation_v1_address_proto_rawDescData = file_google_maps_addressvalidation_v1_address_proto_rawDesc
)

func file_google_maps_addressvalidation_v1_address_proto_rawDescGZIP() []byte {
	file_google_maps_addressvalidation_v1_address_proto_rawDescOnce.Do(func() {
		file_google_maps_addressvalidation_v1_address_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_maps_addressvalidation_v1_address_proto_rawDescData)
	})
	return file_google_maps_addressvalidation_v1_address_proto_rawDescData
}

var file_google_maps_addressvalidation_v1_address_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_maps_addressvalidation_v1_address_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_google_maps_addressvalidation_v1_address_proto_goTypes = []interface{}{
	(AddressComponent_ConfirmationLevel)(0), // 0: google.maps.addressvalidation.v1.AddressComponent.ConfirmationLevel
	(*Address)(nil),                         // 1: google.maps.addressvalidation.v1.Address
	(*AddressComponent)(nil),                // 2: google.maps.addressvalidation.v1.AddressComponent
	(*ComponentName)(nil),                   // 3: google.maps.addressvalidation.v1.ComponentName
	(*postaladdress.PostalAddress)(nil),     // 4: google.type.PostalAddress
}
var file_google_maps_addressvalidation_v1_address_proto_depIdxs = []int32{
	4, // 0: google.maps.addressvalidation.v1.Address.postal_address:type_name -> google.type.PostalAddress
	2, // 1: google.maps.addressvalidation.v1.Address.address_components:type_name -> google.maps.addressvalidation.v1.AddressComponent
	3, // 2: google.maps.addressvalidation.v1.AddressComponent.component_name:type_name -> google.maps.addressvalidation.v1.ComponentName
	0, // 3: google.maps.addressvalidation.v1.AddressComponent.confirmation_level:type_name -> google.maps.addressvalidation.v1.AddressComponent.ConfirmationLevel
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_google_maps_addressvalidation_v1_address_proto_init() }
func file_google_maps_addressvalidation_v1_address_proto_init() {
	if File_google_maps_addressvalidation_v1_address_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_maps_addressvalidation_v1_address_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Address); i {
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
		file_google_maps_addressvalidation_v1_address_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddressComponent); i {
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
		file_google_maps_addressvalidation_v1_address_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ComponentName); i {
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
			RawDescriptor: file_google_maps_addressvalidation_v1_address_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_maps_addressvalidation_v1_address_proto_goTypes,
		DependencyIndexes: file_google_maps_addressvalidation_v1_address_proto_depIdxs,
		EnumInfos:         file_google_maps_addressvalidation_v1_address_proto_enumTypes,
		MessageInfos:      file_google_maps_addressvalidation_v1_address_proto_msgTypes,
	}.Build()
	File_google_maps_addressvalidation_v1_address_proto = out.File
	file_google_maps_addressvalidation_v1_address_proto_rawDesc = nil
	file_google_maps_addressvalidation_v1_address_proto_goTypes = nil
	file_google_maps_addressvalidation_v1_address_proto_depIdxs = nil
}
