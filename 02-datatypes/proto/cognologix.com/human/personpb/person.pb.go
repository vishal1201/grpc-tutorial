// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.6.1
// source: person.proto

package personpb

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	any "github.com/golang/protobuf/ptypes/any"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type PhoneNumber_PhoneType int32

const (
	PhoneNumber_MOBILE PhoneNumber_PhoneType = 0
	PhoneNumber_HOME   PhoneNumber_PhoneType = 1
	PhoneNumber_WORK   PhoneNumber_PhoneType = 2
)

// Enum value maps for PhoneNumber_PhoneType.
var (
	PhoneNumber_PhoneType_name = map[int32]string{
		0: "MOBILE",
		1: "HOME",
		2: "WORK",
	}
	PhoneNumber_PhoneType_value = map[string]int32{
		"MOBILE": 0,
		"HOME":   1,
		"WORK":   2,
	}
)

func (x PhoneNumber_PhoneType) Enum() *PhoneNumber_PhoneType {
	p := new(PhoneNumber_PhoneType)
	*p = x
	return p
}

func (x PhoneNumber_PhoneType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PhoneNumber_PhoneType) Descriptor() protoreflect.EnumDescriptor {
	return file_person_proto_enumTypes[0].Descriptor()
}

func (PhoneNumber_PhoneType) Type() protoreflect.EnumType {
	return &file_person_proto_enumTypes[0]
}

func (x PhoneNumber_PhoneType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PhoneNumber_PhoneType.Descriptor instead.
func (PhoneNumber_PhoneType) EnumDescriptor() ([]byte, []int) {
	return file_person_proto_rawDescGZIP(), []int{0, 0}
}

type PhoneNumber struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Number string                `protobuf:"bytes,1,opt,name=number,proto3" json:"number,omitempty"`
	PhType PhoneNumber_PhoneType `protobuf:"varint,2,opt,name=ph_type,json=phType,proto3,enum=human.PhoneNumber_PhoneType" json:"ph_type,omitempty"`
}

func (x *PhoneNumber) Reset() {
	*x = PhoneNumber{}
	if protoimpl.UnsafeEnabled {
		mi := &file_person_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PhoneNumber) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PhoneNumber) ProtoMessage() {}

func (x *PhoneNumber) ProtoReflect() protoreflect.Message {
	mi := &file_person_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PhoneNumber.ProtoReflect.Descriptor instead.
func (*PhoneNumber) Descriptor() ([]byte, []int) {
	return file_person_proto_rawDescGZIP(), []int{0}
}

func (x *PhoneNumber) GetNumber() string {
	if x != nil {
		return x.Number
	}
	return ""
}

func (x *PhoneNumber) GetPhType() PhoneNumber_PhoneType {
	if x != nil {
		return x.PhType
	}
	return PhoneNumber_MOBILE
}

type Permanent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Line string `protobuf:"bytes,1,opt,name=line,proto3" json:"line,omitempty"`
}

func (x *Permanent) Reset() {
	*x = Permanent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_person_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Permanent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Permanent) ProtoMessage() {}

func (x *Permanent) ProtoReflect() protoreflect.Message {
	mi := &file_person_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Permanent.ProtoReflect.Descriptor instead.
func (*Permanent) Descriptor() ([]byte, []int) {
	return file_person_proto_rawDescGZIP(), []int{1}
}

func (x *Permanent) GetLine() string {
	if x != nil {
		return x.Line
	}
	return ""
}

type Communication struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Line string `protobuf:"bytes,1,opt,name=line,proto3" json:"line,omitempty"`
}

func (x *Communication) Reset() {
	*x = Communication{}
	if protoimpl.UnsafeEnabled {
		mi := &file_person_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Communication) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Communication) ProtoMessage() {}

func (x *Communication) ProtoReflect() protoreflect.Message {
	mi := &file_person_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Communication.ProtoReflect.Descriptor instead.
func (*Communication) Descriptor() ([]byte, []int) {
	return file_person_proto_rawDescGZIP(), []int{2}
}

func (x *Communication) GetLine() string {
	if x != nil {
		return x.Line
	}
	return ""
}

type Person struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int64                `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	FirstName   string               `protobuf:"bytes,2,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	LastName    string               `protobuf:"bytes,3,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	Age         int32                `protobuf:"varint,4,opt,name=age,proto3" json:"age,omitempty"`
	Phones      []*PhoneNumber       `protobuf:"bytes,5,rep,name=phones,proto3" json:"phones,omitempty"`
	LastUpdated *timestamp.Timestamp `protobuf:"bytes,6,opt,name=last_updated,json=lastUpdated,proto3" json:"last_updated,omitempty"`
	Extra       []*any.Any           `protobuf:"bytes,7,rep,name=extra,proto3" json:"extra,omitempty"`
	// Types that are assignable to AddressType:
	//	*Person_Permanent
	//	*Person_Communication
	AddressType isPerson_AddressType `protobuf_oneof:"address_type"`
	Family      map[string]*Person   `protobuf:"bytes,10,rep,name=family,proto3" json:"family,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *Person) Reset() {
	*x = Person{}
	if protoimpl.UnsafeEnabled {
		mi := &file_person_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Person) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Person) ProtoMessage() {}

func (x *Person) ProtoReflect() protoreflect.Message {
	mi := &file_person_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Person.ProtoReflect.Descriptor instead.
func (*Person) Descriptor() ([]byte, []int) {
	return file_person_proto_rawDescGZIP(), []int{3}
}

func (x *Person) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Person) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *Person) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *Person) GetAge() int32 {
	if x != nil {
		return x.Age
	}
	return 0
}

func (x *Person) GetPhones() []*PhoneNumber {
	if x != nil {
		return x.Phones
	}
	return nil
}

func (x *Person) GetLastUpdated() *timestamp.Timestamp {
	if x != nil {
		return x.LastUpdated
	}
	return nil
}

func (x *Person) GetExtra() []*any.Any {
	if x != nil {
		return x.Extra
	}
	return nil
}

func (m *Person) GetAddressType() isPerson_AddressType {
	if m != nil {
		return m.AddressType
	}
	return nil
}

func (x *Person) GetPermanent() *Permanent {
	if x, ok := x.GetAddressType().(*Person_Permanent); ok {
		return x.Permanent
	}
	return nil
}

func (x *Person) GetCommunication() *Communication {
	if x, ok := x.GetAddressType().(*Person_Communication); ok {
		return x.Communication
	}
	return nil
}

func (x *Person) GetFamily() map[string]*Person {
	if x != nil {
		return x.Family
	}
	return nil
}

type isPerson_AddressType interface {
	isPerson_AddressType()
}

type Person_Permanent struct {
	Permanent *Permanent `protobuf:"bytes,8,opt,name=permanent,proto3,oneof"`
}

type Person_Communication struct {
	Communication *Communication `protobuf:"bytes,9,opt,name=communication,proto3,oneof"`
}

func (*Person_Permanent) isPerson_AddressType() {}

func (*Person_Communication) isPerson_AddressType() {}

type GetPersonRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetPersonRequest) Reset() {
	*x = GetPersonRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_person_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPersonRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPersonRequest) ProtoMessage() {}

func (x *GetPersonRequest) ProtoReflect() protoreflect.Message {
	mi := &file_person_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPersonRequest.ProtoReflect.Descriptor instead.
func (*GetPersonRequest) Descriptor() ([]byte, []int) {
	return file_person_proto_rawDescGZIP(), []int{4}
}

func (x *GetPersonRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GetPersonResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Person *Person `protobuf:"bytes,1,opt,name=person,proto3" json:"person,omitempty"`
}

func (x *GetPersonResponse) Reset() {
	*x = GetPersonResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_person_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPersonResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPersonResponse) ProtoMessage() {}

func (x *GetPersonResponse) ProtoReflect() protoreflect.Message {
	mi := &file_person_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPersonResponse.ProtoReflect.Descriptor instead.
func (*GetPersonResponse) Descriptor() ([]byte, []int) {
	return file_person_proto_rawDescGZIP(), []int{5}
}

func (x *GetPersonResponse) GetPerson() *Person {
	if x != nil {
		return x.Person
	}
	return nil
}

type PutPersonRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Person []*Person `protobuf:"bytes,1,rep,name=person,proto3" json:"person,omitempty"`
}

func (x *PutPersonRequest) Reset() {
	*x = PutPersonRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_person_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PutPersonRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PutPersonRequest) ProtoMessage() {}

func (x *PutPersonRequest) ProtoReflect() protoreflect.Message {
	mi := &file_person_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PutPersonRequest.ProtoReflect.Descriptor instead.
func (*PutPersonRequest) Descriptor() ([]byte, []int) {
	return file_person_proto_rawDescGZIP(), []int{6}
}

func (x *PutPersonRequest) GetPerson() []*Person {
	if x != nil {
		return x.Person
	}
	return nil
}

type PutPersonResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result bool    `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
	Id     []int64 `protobuf:"varint,2,rep,packed,name=id,proto3" json:"id,omitempty"`
}

func (x *PutPersonResponse) Reset() {
	*x = PutPersonResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_person_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PutPersonResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PutPersonResponse) ProtoMessage() {}

func (x *PutPersonResponse) ProtoReflect() protoreflect.Message {
	mi := &file_person_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PutPersonResponse.ProtoReflect.Descriptor instead.
func (*PutPersonResponse) Descriptor() ([]byte, []int) {
	return file_person_proto_rawDescGZIP(), []int{7}
}

func (x *PutPersonResponse) GetResult() bool {
	if x != nil {
		return x.Result
	}
	return false
}

func (x *PutPersonResponse) GetId() []int64 {
	if x != nil {
		return x.Id
	}
	return nil
}

var File_person_proto protoreflect.FileDescriptor

var file_person_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05,
	0x68, 0x75, 0x6d, 0x61, 0x6e, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x89, 0x01, 0x0a, 0x0b, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x35, 0x0a, 0x07, 0x70, 0x68, 0x5f,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1c, 0x2e, 0x68, 0x75, 0x6d,
	0x61, 0x6e, 0x2e, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x2e, 0x50,
	0x68, 0x6f, 0x6e, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x06, 0x70, 0x68, 0x54, 0x79, 0x70, 0x65,
	0x22, 0x2b, 0x0a, 0x09, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0a, 0x0a,
	0x06, 0x4d, 0x4f, 0x42, 0x49, 0x4c, 0x45, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x48, 0x4f, 0x4d,
	0x45, 0x10, 0x01, 0x12, 0x08, 0x0a, 0x04, 0x57, 0x4f, 0x52, 0x4b, 0x10, 0x02, 0x22, 0x1f, 0x0a,
	0x09, 0x50, 0x65, 0x72, 0x6d, 0x61, 0x6e, 0x65, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6c, 0x69,
	0x6e, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6c, 0x69, 0x6e, 0x65, 0x22, 0x23,
	0x0a, 0x0d, 0x43, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x12, 0x0a, 0x04, 0x6c, 0x69, 0x6e, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6c,
	0x69, 0x6e, 0x65, 0x22, 0xfa, 0x03, 0x0a, 0x06, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1d,
	0x0a, 0x0a, 0x66, 0x69, 0x72, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a,
	0x09, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x67,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x61, 0x67, 0x65, 0x12, 0x2a, 0x0a, 0x06,
	0x70, 0x68, 0x6f, 0x6e, 0x65, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x68,
	0x75, 0x6d, 0x61, 0x6e, 0x2e, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72,
	0x52, 0x06, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x73, 0x12, 0x3d, 0x0a, 0x0c, 0x6c, 0x61, 0x73, 0x74,
	0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0b, 0x6c, 0x61, 0x73, 0x74,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x12, 0x2a, 0x0a, 0x05, 0x65, 0x78, 0x74, 0x72, 0x61,
	0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x05, 0x65, 0x78,
	0x74, 0x72, 0x61, 0x12, 0x30, 0x0a, 0x09, 0x70, 0x65, 0x72, 0x6d, 0x61, 0x6e, 0x65, 0x6e, 0x74,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x68, 0x75, 0x6d, 0x61, 0x6e, 0x2e, 0x50,
	0x65, 0x72, 0x6d, 0x61, 0x6e, 0x65, 0x6e, 0x74, 0x48, 0x00, 0x52, 0x09, 0x70, 0x65, 0x72, 0x6d,
	0x61, 0x6e, 0x65, 0x6e, 0x74, 0x12, 0x3c, 0x0a, 0x0d, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x68,
	0x75, 0x6d, 0x61, 0x6e, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x48, 0x00, 0x52, 0x0d, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x31, 0x0a, 0x06, 0x66, 0x61, 0x6d, 0x69, 0x6c, 0x79, 0x18, 0x0a, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x68, 0x75, 0x6d, 0x61, 0x6e, 0x2e, 0x50, 0x65, 0x72, 0x73,
	0x6f, 0x6e, 0x2e, 0x46, 0x61, 0x6d, 0x69, 0x6c, 0x79, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06,
	0x66, 0x61, 0x6d, 0x69, 0x6c, 0x79, 0x1a, 0x48, 0x0a, 0x0b, 0x46, 0x61, 0x6d, 0x69, 0x6c, 0x79,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x23, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x68, 0x75, 0x6d, 0x61, 0x6e, 0x2e, 0x50,
	0x65, 0x72, 0x73, 0x6f, 0x6e, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01,
	0x42, 0x0e, 0x0a, 0x0c, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x5f, 0x74, 0x79, 0x70, 0x65,
	0x22, 0x22, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x02, 0x69, 0x64, 0x22, 0x3a, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x50, 0x65, 0x72, 0x73, 0x6f,
	0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x25, 0x0a, 0x06, 0x70, 0x65, 0x72,
	0x73, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x68, 0x75, 0x6d, 0x61,
	0x6e, 0x2e, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x52, 0x06, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e,
	0x22, 0x39, 0x0a, 0x10, 0x50, 0x75, 0x74, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x25, 0x0a, 0x06, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x68, 0x75, 0x6d, 0x61, 0x6e, 0x2e, 0x50, 0x65, 0x72,
	0x73, 0x6f, 0x6e, 0x52, 0x06, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x22, 0x3b, 0x0a, 0x11, 0x50,
	0x75, 0x74, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x32, 0xd5, 0x01, 0x0a, 0x0d, 0x50, 0x65, 0x72,
	0x73, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x40, 0x0a, 0x09, 0x47, 0x65,
	0x74, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x12, 0x17, 0x2e, 0x68, 0x75, 0x6d, 0x61, 0x6e, 0x2e,
	0x47, 0x65, 0x74, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x18, 0x2e, 0x68, 0x75, 0x6d, 0x61, 0x6e, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x65, 0x72, 0x73,
	0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x40, 0x0a, 0x09,
	0x50, 0x75, 0x74, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x12, 0x17, 0x2e, 0x68, 0x75, 0x6d, 0x61,
	0x6e, 0x2e, 0x50, 0x75, 0x74, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x18, 0x2e, 0x68, 0x75, 0x6d, 0x61, 0x6e, 0x2e, 0x50, 0x75, 0x74, 0x50, 0x65,
	0x72, 0x73, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x40,
	0x0a, 0x09, 0x50, 0x75, 0x74, 0x50, 0x65, 0x6f, 0x70, 0x6c, 0x65, 0x12, 0x17, 0x2e, 0x68, 0x75,
	0x6d, 0x61, 0x6e, 0x2e, 0x50, 0x75, 0x74, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x68, 0x75, 0x6d, 0x61, 0x6e, 0x2e, 0x50, 0x75, 0x74,
	0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x42, 0x21, 0x48, 0x01, 0x5a, 0x1d, 0x63, 0x6f, 0x67, 0x6e, 0x6f, 0x6c, 0x6f, 0x67, 0x69, 0x78,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x68, 0x75, 0x6d, 0x61, 0x6e, 0x2f, 0x70, 0x65, 0x72, 0x73, 0x6f,
	0x6e, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_person_proto_rawDescOnce sync.Once
	file_person_proto_rawDescData = file_person_proto_rawDesc
)

func file_person_proto_rawDescGZIP() []byte {
	file_person_proto_rawDescOnce.Do(func() {
		file_person_proto_rawDescData = protoimpl.X.CompressGZIP(file_person_proto_rawDescData)
	})
	return file_person_proto_rawDescData
}

var file_person_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_person_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_person_proto_goTypes = []interface{}{
	(PhoneNumber_PhoneType)(0),  // 0: human.PhoneNumber.PhoneType
	(*PhoneNumber)(nil),         // 1: human.PhoneNumber
	(*Permanent)(nil),           // 2: human.Permanent
	(*Communication)(nil),       // 3: human.Communication
	(*Person)(nil),              // 4: human.Person
	(*GetPersonRequest)(nil),    // 5: human.GetPersonRequest
	(*GetPersonResponse)(nil),   // 6: human.GetPersonResponse
	(*PutPersonRequest)(nil),    // 7: human.PutPersonRequest
	(*PutPersonResponse)(nil),   // 8: human.PutPersonResponse
	nil,                         // 9: human.Person.FamilyEntry
	(*timestamp.Timestamp)(nil), // 10: google.protobuf.Timestamp
	(*any.Any)(nil),             // 11: google.protobuf.Any
}
var file_person_proto_depIdxs = []int32{
	0,  // 0: human.PhoneNumber.ph_type:type_name -> human.PhoneNumber.PhoneType
	1,  // 1: human.Person.phones:type_name -> human.PhoneNumber
	10, // 2: human.Person.last_updated:type_name -> google.protobuf.Timestamp
	11, // 3: human.Person.extra:type_name -> google.protobuf.Any
	2,  // 4: human.Person.permanent:type_name -> human.Permanent
	3,  // 5: human.Person.communication:type_name -> human.Communication
	9,  // 6: human.Person.family:type_name -> human.Person.FamilyEntry
	4,  // 7: human.GetPersonResponse.person:type_name -> human.Person
	4,  // 8: human.PutPersonRequest.person:type_name -> human.Person
	4,  // 9: human.Person.FamilyEntry.value:type_name -> human.Person
	5,  // 10: human.PersonService.GetPerson:input_type -> human.GetPersonRequest
	7,  // 11: human.PersonService.PutPerson:input_type -> human.PutPersonRequest
	7,  // 12: human.PersonService.PutPeople:input_type -> human.PutPersonRequest
	6,  // 13: human.PersonService.GetPerson:output_type -> human.GetPersonResponse
	8,  // 14: human.PersonService.PutPerson:output_type -> human.PutPersonResponse
	8,  // 15: human.PersonService.PutPeople:output_type -> human.PutPersonResponse
	13, // [13:16] is the sub-list for method output_type
	10, // [10:13] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_person_proto_init() }
func file_person_proto_init() {
	if File_person_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_person_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PhoneNumber); i {
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
		file_person_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Permanent); i {
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
		file_person_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Communication); i {
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
		file_person_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Person); i {
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
		file_person_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPersonRequest); i {
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
		file_person_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPersonResponse); i {
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
		file_person_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PutPersonRequest); i {
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
		file_person_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PutPersonResponse); i {
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
	file_person_proto_msgTypes[3].OneofWrappers = []interface{}{
		(*Person_Permanent)(nil),
		(*Person_Communication)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_person_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_person_proto_goTypes,
		DependencyIndexes: file_person_proto_depIdxs,
		EnumInfos:         file_person_proto_enumTypes,
		MessageInfos:      file_person_proto_msgTypes,
	}.Build()
	File_person_proto = out.File
	file_person_proto_rawDesc = nil
	file_person_proto_goTypes = nil
	file_person_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// PersonServiceClient is the client API for PersonService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PersonServiceClient interface {
	GetPerson(ctx context.Context, in *GetPersonRequest, opts ...grpc.CallOption) (*GetPersonResponse, error)
	PutPerson(ctx context.Context, in *PutPersonRequest, opts ...grpc.CallOption) (*PutPersonResponse, error)
	PutPeople(ctx context.Context, in *PutPersonRequest, opts ...grpc.CallOption) (*PutPersonResponse, error)
}

type personServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPersonServiceClient(cc grpc.ClientConnInterface) PersonServiceClient {
	return &personServiceClient{cc}
}

func (c *personServiceClient) GetPerson(ctx context.Context, in *GetPersonRequest, opts ...grpc.CallOption) (*GetPersonResponse, error) {
	out := new(GetPersonResponse)
	err := c.cc.Invoke(ctx, "/human.PersonService/GetPerson", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *personServiceClient) PutPerson(ctx context.Context, in *PutPersonRequest, opts ...grpc.CallOption) (*PutPersonResponse, error) {
	out := new(PutPersonResponse)
	err := c.cc.Invoke(ctx, "/human.PersonService/PutPerson", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *personServiceClient) PutPeople(ctx context.Context, in *PutPersonRequest, opts ...grpc.CallOption) (*PutPersonResponse, error) {
	out := new(PutPersonResponse)
	err := c.cc.Invoke(ctx, "/human.PersonService/PutPeople", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PersonServiceServer is the server API for PersonService service.
type PersonServiceServer interface {
	GetPerson(context.Context, *GetPersonRequest) (*GetPersonResponse, error)
	PutPerson(context.Context, *PutPersonRequest) (*PutPersonResponse, error)
	PutPeople(context.Context, *PutPersonRequest) (*PutPersonResponse, error)
}

// UnimplementedPersonServiceServer can be embedded to have forward compatible implementations.
type UnimplementedPersonServiceServer struct {
}

func (*UnimplementedPersonServiceServer) GetPerson(context.Context, *GetPersonRequest) (*GetPersonResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPerson not implemented")
}
func (*UnimplementedPersonServiceServer) PutPerson(context.Context, *PutPersonRequest) (*PutPersonResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PutPerson not implemented")
}
func (*UnimplementedPersonServiceServer) PutPeople(context.Context, *PutPersonRequest) (*PutPersonResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PutPeople not implemented")
}

func RegisterPersonServiceServer(s *grpc.Server, srv PersonServiceServer) {
	s.RegisterService(&_PersonService_serviceDesc, srv)
}

func _PersonService_GetPerson_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPersonRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PersonServiceServer).GetPerson(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/human.PersonService/GetPerson",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PersonServiceServer).GetPerson(ctx, req.(*GetPersonRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PersonService_PutPerson_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PutPersonRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PersonServiceServer).PutPerson(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/human.PersonService/PutPerson",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PersonServiceServer).PutPerson(ctx, req.(*PutPersonRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PersonService_PutPeople_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PutPersonRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PersonServiceServer).PutPeople(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/human.PersonService/PutPeople",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PersonServiceServer).PutPeople(ctx, req.(*PutPersonRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _PersonService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "human.PersonService",
	HandlerType: (*PersonServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetPerson",
			Handler:    _PersonService_GetPerson_Handler,
		},
		{
			MethodName: "PutPerson",
			Handler:    _PersonService_PutPerson_Handler,
		},
		{
			MethodName: "PutPeople",
			Handler:    _PersonService_PutPeople_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "person.proto",
}
