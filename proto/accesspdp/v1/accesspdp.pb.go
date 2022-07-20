// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        (unknown)
// source: accesspdp/v1/accesspdp.proto

package v1

import (
	v1 "github.com/virtru/access-pdp/proto/attributes/v1"
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

type DetermineAccessRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DataAttributes       []*v1.AttributeInstance              `protobuf:"bytes,1,rep,name=data_attributes,json=dataAttributes,proto3" json:"data_attributes,omitempty"`
	EntityAttributeSets  map[string]*ListOfAttributeInstances `protobuf:"bytes,2,rep,name=entity_attribute_sets,json=entityAttributeSets,proto3" json:"entity_attribute_sets,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	AttributeDefinitions []*v1.AttributeDefinition            `protobuf:"bytes,3,rep,name=attribute_definitions,json=attributeDefinitions,proto3" json:"attribute_definitions,omitempty"`
}

func (x *DetermineAccessRequest) Reset() {
	*x = DetermineAccessRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_accesspdp_v1_accesspdp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DetermineAccessRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DetermineAccessRequest) ProtoMessage() {}

func (x *DetermineAccessRequest) ProtoReflect() protoreflect.Message {
	mi := &file_accesspdp_v1_accesspdp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DetermineAccessRequest.ProtoReflect.Descriptor instead.
func (*DetermineAccessRequest) Descriptor() ([]byte, []int) {
	return file_accesspdp_v1_accesspdp_proto_rawDescGZIP(), []int{0}
}

func (x *DetermineAccessRequest) GetDataAttributes() []*v1.AttributeInstance {
	if x != nil {
		return x.DataAttributes
	}
	return nil
}

func (x *DetermineAccessRequest) GetEntityAttributeSets() map[string]*ListOfAttributeInstances {
	if x != nil {
		return x.EntityAttributeSets
	}
	return nil
}

func (x *DetermineAccessRequest) GetAttributeDefinitions() []*v1.AttributeDefinition {
	if x != nil {
		return x.AttributeDefinitions
	}
	return nil
}

type ListOfAttributeInstances struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AttributeInstances []*v1.AttributeInstance `protobuf:"bytes,1,rep,name=attribute_instances,json=attributeInstances,proto3" json:"attribute_instances,omitempty"`
}

func (x *ListOfAttributeInstances) Reset() {
	*x = ListOfAttributeInstances{}
	if protoimpl.UnsafeEnabled {
		mi := &file_accesspdp_v1_accesspdp_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListOfAttributeInstances) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListOfAttributeInstances) ProtoMessage() {}

func (x *ListOfAttributeInstances) ProtoReflect() protoreflect.Message {
	mi := &file_accesspdp_v1_accesspdp_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListOfAttributeInstances.ProtoReflect.Descriptor instead.
func (*ListOfAttributeInstances) Descriptor() ([]byte, []int) {
	return file_accesspdp_v1_accesspdp_proto_rawDescGZIP(), []int{1}
}

func (x *ListOfAttributeInstances) GetAttributeInstances() []*v1.AttributeInstance {
	if x != nil {
		return x.AttributeInstances
	}
	return nil
}

type DetermineAccessResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Access  *bool             `protobuf:"varint,1,opt,name=access,proto3,oneof" json:"access,omitempty"`
	Results []*DataRuleResult `protobuf:"bytes,2,rep,name=results,proto3" json:"results,omitempty"`
}

func (x *DetermineAccessResponse) Reset() {
	*x = DetermineAccessResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_accesspdp_v1_accesspdp_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DetermineAccessResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DetermineAccessResponse) ProtoMessage() {}

func (x *DetermineAccessResponse) ProtoReflect() protoreflect.Message {
	mi := &file_accesspdp_v1_accesspdp_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DetermineAccessResponse.ProtoReflect.Descriptor instead.
func (*DetermineAccessResponse) Descriptor() ([]byte, []int) {
	return file_accesspdp_v1_accesspdp_proto_rawDescGZIP(), []int{2}
}

func (x *DetermineAccessResponse) GetAccess() bool {
	if x != nil && x.Access != nil {
		return *x.Access
	}
	return false
}

func (x *DetermineAccessResponse) GetResults() []*DataRuleResult {
	if x != nil {
		return x.Results
	}
	return nil
}

type DataRuleResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Passed         *bool                   `protobuf:"varint,1,opt,name=passed,proto3,oneof" json:"passed,omitempty"`
	RuleDefinition *v1.AttributeDefinition `protobuf:"bytes,2,opt,name=rule_definition,json=ruleDefinition,proto3,oneof" json:"rule_definition,omitempty"`
	ValueFailures  []*ValueFailure         `protobuf:"bytes,3,rep,name=value_failures,json=valueFailures,proto3" json:"value_failures,omitempty"`
}

func (x *DataRuleResult) Reset() {
	*x = DataRuleResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_accesspdp_v1_accesspdp_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DataRuleResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DataRuleResult) ProtoMessage() {}

func (x *DataRuleResult) ProtoReflect() protoreflect.Message {
	mi := &file_accesspdp_v1_accesspdp_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DataRuleResult.ProtoReflect.Descriptor instead.
func (*DataRuleResult) Descriptor() ([]byte, []int) {
	return file_accesspdp_v1_accesspdp_proto_rawDescGZIP(), []int{3}
}

func (x *DataRuleResult) GetPassed() bool {
	if x != nil && x.Passed != nil {
		return *x.Passed
	}
	return false
}

func (x *DataRuleResult) GetRuleDefinition() *v1.AttributeDefinition {
	if x != nil {
		return x.RuleDefinition
	}
	return nil
}

func (x *DataRuleResult) GetValueFailures() []*ValueFailure {
	if x != nil {
		return x.ValueFailures
	}
	return nil
}

type ValueFailure struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DataAttribute *v1.AttributeInstance `protobuf:"bytes,1,opt,name=data_attribute,json=dataAttribute,proto3,oneof" json:"data_attribute,omitempty"`
	Message       *string               `protobuf:"bytes,2,opt,name=message,proto3,oneof" json:"message,omitempty"`
}

func (x *ValueFailure) Reset() {
	*x = ValueFailure{}
	if protoimpl.UnsafeEnabled {
		mi := &file_accesspdp_v1_accesspdp_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ValueFailure) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValueFailure) ProtoMessage() {}

func (x *ValueFailure) ProtoReflect() protoreflect.Message {
	mi := &file_accesspdp_v1_accesspdp_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ValueFailure.ProtoReflect.Descriptor instead.
func (*ValueFailure) Descriptor() ([]byte, []int) {
	return file_accesspdp_v1_accesspdp_proto_rawDescGZIP(), []int{4}
}

func (x *ValueFailure) GetDataAttribute() *v1.AttributeInstance {
	if x != nil {
		return x.DataAttribute
	}
	return nil
}

func (x *ValueFailure) GetMessage() string {
	if x != nil && x.Message != nil {
		return *x.Message
	}
	return ""
}

var File_accesspdp_v1_accesspdp_proto protoreflect.FileDescriptor

var file_accesspdp_v1_accesspdp_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x70, 0x64, 0x70, 0x2f, 0x76, 0x31, 0x2f, 0x61,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x70, 0x64, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c,
	0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x70, 0x64, 0x70, 0x2e, 0x76, 0x31, 0x1a, 0x1e, 0x61, 0x74,
	0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x74, 0x74, 0x72,
	0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9f, 0x03, 0x0a,
	0x16, 0x44, 0x65, 0x74, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x65, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x49, 0x0a, 0x0f, 0x64, 0x61, 0x74, 0x61, 0x5f,
	0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x20, 0x2e, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x2e, 0x76, 0x31,
	0x2e, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e,
	0x63, 0x65, 0x52, 0x0e, 0x64, 0x61, 0x74, 0x61, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74,
	0x65, 0x73, 0x12, 0x71, 0x0a, 0x15, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x5f, 0x61, 0x74, 0x74,
	0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x5f, 0x73, 0x65, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x3d, 0x2e, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x70, 0x64, 0x70, 0x2e, 0x76, 0x31,
	0x2e, 0x44, 0x65, 0x74, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x65, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x41, 0x74,
	0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x53, 0x65, 0x74, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x52, 0x13, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74,
	0x65, 0x53, 0x65, 0x74, 0x73, 0x12, 0x57, 0x0a, 0x15, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75,
	0x74, 0x65, 0x5f, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x03,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65,
	0x73, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x44, 0x65,
	0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x14, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62,
	0x75, 0x74, 0x65, 0x44, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x1a, 0x6e,
	0x0a, 0x18, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74,
	0x65, 0x53, 0x65, 0x74, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x3c, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x61, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x70, 0x64, 0x70, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4f,
	0x66, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e,
	0x63, 0x65, 0x73, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x6d,
	0x0a, 0x18, 0x4c, 0x69, 0x73, 0x74, 0x4f, 0x66, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74,
	0x65, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x12, 0x51, 0x0a, 0x13, 0x61, 0x74,
	0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x5f, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62,
	0x75, 0x74, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74,
	0x65, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x12, 0x61, 0x74, 0x74, 0x72, 0x69,
	0x62, 0x75, 0x74, 0x65, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x22, 0x79, 0x0a,
	0x17, 0x44, 0x65, 0x74, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x65, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1b, 0x0a, 0x06, 0x61, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x48, 0x00, 0x52, 0x06, 0x61, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x88, 0x01, 0x01, 0x12, 0x36, 0x0a, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x70,
	0x64, 0x70, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x52, 0x75, 0x6c, 0x65, 0x52, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x52, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x42, 0x09, 0x0a,
	0x07, 0x5f, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x22, 0xe1, 0x01, 0x0a, 0x0e, 0x44, 0x61, 0x74,
	0x61, 0x52, 0x75, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x1b, 0x0a, 0x06, 0x70,
	0x61, 0x73, 0x73, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x48, 0x00, 0x52, 0x06, 0x70,
	0x61, 0x73, 0x73, 0x65, 0x64, 0x88, 0x01, 0x01, 0x12, 0x50, 0x0a, 0x0f, 0x72, 0x75, 0x6c, 0x65,
	0x5f, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x22, 0x2e, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x2e, 0x76,
	0x31, 0x2e, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x44, 0x65, 0x66, 0x69, 0x6e,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x01, 0x52, 0x0e, 0x72, 0x75, 0x6c, 0x65, 0x44, 0x65, 0x66,
	0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x88, 0x01, 0x01, 0x12, 0x41, 0x0a, 0x0e, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x5f, 0x66, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x70, 0x64, 0x70, 0x2e, 0x76,
	0x31, 0x2e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x46, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x52, 0x0d,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x46, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x73, 0x42, 0x09, 0x0a,
	0x07, 0x5f, 0x70, 0x61, 0x73, 0x73, 0x65, 0x64, 0x42, 0x12, 0x0a, 0x10, 0x5f, 0x72, 0x75, 0x6c,
	0x65, 0x5f, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x9a, 0x01, 0x0a,
	0x0c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x46, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x12, 0x4c, 0x0a,
	0x0e, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74,
	0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x49,
	0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x48, 0x00, 0x52, 0x0d, 0x64, 0x61, 0x74, 0x61, 0x41,
	0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x88, 0x01, 0x01, 0x12, 0x1d, 0x0a, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x88, 0x01, 0x01, 0x42, 0x11, 0x0a, 0x0f, 0x5f, 0x64,
	0x61, 0x74, 0x61, 0x5f, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x42, 0x0a, 0x0a,
	0x08, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0x75, 0x0a, 0x11, 0x41, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x50, 0x44, 0x50, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x60,
	0x0a, 0x0f, 0x44, 0x65, 0x74, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x65, 0x41, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x12, 0x24, 0x2e, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x70, 0x64, 0x70, 0x2e, 0x76, 0x31,
	0x2e, 0x44, 0x65, 0x74, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x65, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x70, 0x64, 0x70, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x74, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x65,
	0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x30, 0x01,
	0x42, 0x31, 0x5a, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x76,
	0x69, 0x72, 0x74, 0x72, 0x75, 0x2f, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x2d, 0x70, 0x64, 0x70,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x70, 0x64, 0x70,
	0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_accesspdp_v1_accesspdp_proto_rawDescOnce sync.Once
	file_accesspdp_v1_accesspdp_proto_rawDescData = file_accesspdp_v1_accesspdp_proto_rawDesc
)

func file_accesspdp_v1_accesspdp_proto_rawDescGZIP() []byte {
	file_accesspdp_v1_accesspdp_proto_rawDescOnce.Do(func() {
		file_accesspdp_v1_accesspdp_proto_rawDescData = protoimpl.X.CompressGZIP(file_accesspdp_v1_accesspdp_proto_rawDescData)
	})
	return file_accesspdp_v1_accesspdp_proto_rawDescData
}

var file_accesspdp_v1_accesspdp_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_accesspdp_v1_accesspdp_proto_goTypes = []interface{}{
	(*DetermineAccessRequest)(nil),   // 0: accesspdp.v1.DetermineAccessRequest
	(*ListOfAttributeInstances)(nil), // 1: accesspdp.v1.ListOfAttributeInstances
	(*DetermineAccessResponse)(nil),  // 2: accesspdp.v1.DetermineAccessResponse
	(*DataRuleResult)(nil),           // 3: accesspdp.v1.DataRuleResult
	(*ValueFailure)(nil),             // 4: accesspdp.v1.ValueFailure
	nil,                              // 5: accesspdp.v1.DetermineAccessRequest.EntityAttributeSetsEntry
	(*v1.AttributeInstance)(nil),     // 6: attributes.v1.AttributeInstance
	(*v1.AttributeDefinition)(nil),   // 7: attributes.v1.AttributeDefinition
}
var file_accesspdp_v1_accesspdp_proto_depIdxs = []int32{
	6,  // 0: accesspdp.v1.DetermineAccessRequest.data_attributes:type_name -> attributes.v1.AttributeInstance
	5,  // 1: accesspdp.v1.DetermineAccessRequest.entity_attribute_sets:type_name -> accesspdp.v1.DetermineAccessRequest.EntityAttributeSetsEntry
	7,  // 2: accesspdp.v1.DetermineAccessRequest.attribute_definitions:type_name -> attributes.v1.AttributeDefinition
	6,  // 3: accesspdp.v1.ListOfAttributeInstances.attribute_instances:type_name -> attributes.v1.AttributeInstance
	3,  // 4: accesspdp.v1.DetermineAccessResponse.results:type_name -> accesspdp.v1.DataRuleResult
	7,  // 5: accesspdp.v1.DataRuleResult.rule_definition:type_name -> attributes.v1.AttributeDefinition
	4,  // 6: accesspdp.v1.DataRuleResult.value_failures:type_name -> accesspdp.v1.ValueFailure
	6,  // 7: accesspdp.v1.ValueFailure.data_attribute:type_name -> attributes.v1.AttributeInstance
	1,  // 8: accesspdp.v1.DetermineAccessRequest.EntityAttributeSetsEntry.value:type_name -> accesspdp.v1.ListOfAttributeInstances
	0,  // 9: accesspdp.v1.AccessPDPEndpoint.DetermineAccess:input_type -> accesspdp.v1.DetermineAccessRequest
	2,  // 10: accesspdp.v1.AccessPDPEndpoint.DetermineAccess:output_type -> accesspdp.v1.DetermineAccessResponse
	10, // [10:11] is the sub-list for method output_type
	9,  // [9:10] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_accesspdp_v1_accesspdp_proto_init() }
func file_accesspdp_v1_accesspdp_proto_init() {
	if File_accesspdp_v1_accesspdp_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_accesspdp_v1_accesspdp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DetermineAccessRequest); i {
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
		file_accesspdp_v1_accesspdp_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListOfAttributeInstances); i {
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
		file_accesspdp_v1_accesspdp_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DetermineAccessResponse); i {
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
		file_accesspdp_v1_accesspdp_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DataRuleResult); i {
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
		file_accesspdp_v1_accesspdp_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ValueFailure); i {
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
	file_accesspdp_v1_accesspdp_proto_msgTypes[2].OneofWrappers = []interface{}{}
	file_accesspdp_v1_accesspdp_proto_msgTypes[3].OneofWrappers = []interface{}{}
	file_accesspdp_v1_accesspdp_proto_msgTypes[4].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_accesspdp_v1_accesspdp_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_accesspdp_v1_accesspdp_proto_goTypes,
		DependencyIndexes: file_accesspdp_v1_accesspdp_proto_depIdxs,
		MessageInfos:      file_accesspdp_v1_accesspdp_proto_msgTypes,
	}.Build()
	File_accesspdp_v1_accesspdp_proto = out.File
	file_accesspdp_v1_accesspdp_proto_rawDesc = nil
	file_accesspdp_v1_accesspdp_proto_goTypes = nil
	file_accesspdp_v1_accesspdp_proto_depIdxs = nil
}
