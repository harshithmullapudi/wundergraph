// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: starwars.proto

package starwars

import (
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

type Type int32

const (
	Type_HUMAN Type = 0
	Type_DROID Type = 1
)

// Enum value maps for Type.
var (
	Type_name = map[int32]string{
		0: "HUMAN",
		1: "DROID",
	}
	Type_value = map[string]int32{
		"HUMAN": 0,
		"DROID": 1,
	}
)

func (x Type) Enum() *Type {
	p := new(Type)
	*p = x
	return p
}

func (x Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Type) Descriptor() protoreflect.EnumDescriptor {
	return file_starwars_proto_enumTypes[0].Descriptor()
}

func (Type) Type() protoreflect.EnumType {
	return &file_starwars_proto_enumTypes[0]
}

func (x Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Type.Descriptor instead.
func (Type) EnumDescriptor() ([]byte, []int) {
	return file_starwars_proto_rawDescGZIP(), []int{0}
}

type Episode int32

const (
	Episode__       Episode = 0
	Episode_NEWHOPE Episode = 1
	Episode_EMPIRE  Episode = 2
	Episode_JEDI    Episode = 3
)

// Enum value maps for Episode.
var (
	Episode_name = map[int32]string{
		0: "_",
		1: "NEWHOPE",
		2: "EMPIRE",
		3: "JEDI",
	}
	Episode_value = map[string]int32{
		"_":       0,
		"NEWHOPE": 1,
		"EMPIRE":  2,
		"JEDI":    3,
	}
)

func (x Episode) Enum() *Episode {
	p := new(Episode)
	*p = x
	return p
}

func (x Episode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Episode) Descriptor() protoreflect.EnumDescriptor {
	return file_starwars_proto_enumTypes[1].Descriptor()
}

func (Episode) Type() protoreflect.EnumType {
	return &file_starwars_proto_enumTypes[1]
}

func (x Episode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Episode.Descriptor instead.
func (Episode) EnumDescriptor() ([]byte, []int) {
	return file_starwars_proto_rawDescGZIP(), []int{1}
}

type Character struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id              string       `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name            string       `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Friends         []*Character `protobuf:"bytes,3,rep,name=friends,proto3" json:"friends,omitempty"`
	AppearsIn       []Episode    `protobuf:"varint,4,rep,packed,name=appears_in,json=appearsIn,proto3,enum=starwars.Episode" json:"appears_in,omitempty"`
	HomePlanet      string       `protobuf:"bytes,5,opt,name=home_planet,json=homePlanet,proto3" json:"home_planet,omitempty"`
	PrimaryFunction string       `protobuf:"bytes,6,opt,name=primary_function,json=primaryFunction,proto3" json:"primary_function,omitempty"`
	Type            Type         `protobuf:"varint,7,opt,name=type,proto3,enum=starwars.Type" json:"type,omitempty"`
}

func (x *Character) Reset() {
	*x = Character{}
	if protoimpl.UnsafeEnabled {
		mi := &file_starwars_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Character) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Character) ProtoMessage() {}

func (x *Character) ProtoReflect() protoreflect.Message {
	mi := &file_starwars_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Character.ProtoReflect.Descriptor instead.
func (*Character) Descriptor() ([]byte, []int) {
	return file_starwars_proto_rawDescGZIP(), []int{0}
}

func (x *Character) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Character) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Character) GetFriends() []*Character {
	if x != nil {
		return x.Friends
	}
	return nil
}

func (x *Character) GetAppearsIn() []Episode {
	if x != nil {
		return x.AppearsIn
	}
	return nil
}

func (x *Character) GetHomePlanet() string {
	if x != nil {
		return x.HomePlanet
	}
	return ""
}

func (x *Character) GetPrimaryFunction() string {
	if x != nil {
		return x.PrimaryFunction
	}
	return ""
}

func (x *Character) GetType() Type {
	if x != nil {
		return x.Type
	}
	return Type_HUMAN
}

type GetHeroRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Episode Episode `protobuf:"varint,1,opt,name=episode,proto3,enum=starwars.Episode" json:"episode,omitempty"`
}

func (x *GetHeroRequest) Reset() {
	*x = GetHeroRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_starwars_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetHeroRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetHeroRequest) ProtoMessage() {}

func (x *GetHeroRequest) ProtoReflect() protoreflect.Message {
	mi := &file_starwars_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetHeroRequest.ProtoReflect.Descriptor instead.
func (*GetHeroRequest) Descriptor() ([]byte, []int) {
	return file_starwars_proto_rawDescGZIP(), []int{1}
}

func (x *GetHeroRequest) GetEpisode() Episode {
	if x != nil {
		return x.Episode
	}
	return Episode__
}

type GetHumanRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetHumanRequest) Reset() {
	*x = GetHumanRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_starwars_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetHumanRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetHumanRequest) ProtoMessage() {}

func (x *GetHumanRequest) ProtoReflect() protoreflect.Message {
	mi := &file_starwars_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetHumanRequest.ProtoReflect.Descriptor instead.
func (*GetHumanRequest) Descriptor() ([]byte, []int) {
	return file_starwars_proto_rawDescGZIP(), []int{2}
}

func (x *GetHumanRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetDroidRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetDroidRequest) Reset() {
	*x = GetDroidRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_starwars_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDroidRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDroidRequest) ProtoMessage() {}

func (x *GetDroidRequest) ProtoReflect() protoreflect.Message {
	mi := &file_starwars_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDroidRequest.ProtoReflect.Descriptor instead.
func (*GetDroidRequest) Descriptor() ([]byte, []int) {
	return file_starwars_proto_rawDescGZIP(), []int{3}
}

func (x *GetDroidRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type ListEmptyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListEmptyRequest) Reset() {
	*x = ListEmptyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_starwars_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListEmptyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListEmptyRequest) ProtoMessage() {}

func (x *ListEmptyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_starwars_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListEmptyRequest.ProtoReflect.Descriptor instead.
func (*ListEmptyRequest) Descriptor() ([]byte, []int) {
	return file_starwars_proto_rawDescGZIP(), []int{4}
}

type ListHumansResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Humans []*Character `protobuf:"bytes,1,rep,name=humans,proto3" json:"humans,omitempty"`
}

func (x *ListHumansResponse) Reset() {
	*x = ListHumansResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_starwars_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListHumansResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListHumansResponse) ProtoMessage() {}

func (x *ListHumansResponse) ProtoReflect() protoreflect.Message {
	mi := &file_starwars_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListHumansResponse.ProtoReflect.Descriptor instead.
func (*ListHumansResponse) Descriptor() ([]byte, []int) {
	return file_starwars_proto_rawDescGZIP(), []int{5}
}

func (x *ListHumansResponse) GetHumans() []*Character {
	if x != nil {
		return x.Humans
	}
	return nil
}

type ListDroidsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Droids []*Character `protobuf:"bytes,1,rep,name=droids,proto3" json:"droids,omitempty"`
}

func (x *ListDroidsResponse) Reset() {
	*x = ListDroidsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_starwars_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListDroidsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListDroidsResponse) ProtoMessage() {}

func (x *ListDroidsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_starwars_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListDroidsResponse.ProtoReflect.Descriptor instead.
func (*ListDroidsResponse) Descriptor() ([]byte, []int) {
	return file_starwars_proto_rawDescGZIP(), []int{6}
}

func (x *ListDroidsResponse) GetDroids() []*Character {
	if x != nil {
		return x.Droids
	}
	return nil
}

var File_starwars_proto protoreflect.FileDescriptor

var file_starwars_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x73, 0x74, 0x61, 0x72, 0x77, 0x61, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x08, 0x73, 0x74, 0x61, 0x72, 0x77, 0x61, 0x72, 0x73, 0x22, 0x80, 0x02, 0x0a, 0x09, 0x43,
	0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x2d, 0x0a, 0x07,
	0x66, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e,
	0x73, 0x74, 0x61, 0x72, 0x77, 0x61, 0x72, 0x73, 0x2e, 0x43, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74,
	0x65, 0x72, 0x52, 0x07, 0x66, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x73, 0x12, 0x30, 0x0a, 0x0a, 0x61,
	0x70, 0x70, 0x65, 0x61, 0x72, 0x73, 0x5f, 0x69, 0x6e, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0e, 0x32,
	0x11, 0x2e, 0x73, 0x74, 0x61, 0x72, 0x77, 0x61, 0x72, 0x73, 0x2e, 0x45, 0x70, 0x69, 0x73, 0x6f,
	0x64, 0x65, 0x52, 0x09, 0x61, 0x70, 0x70, 0x65, 0x61, 0x72, 0x73, 0x49, 0x6e, 0x12, 0x1f, 0x0a,
	0x0b, 0x68, 0x6f, 0x6d, 0x65, 0x5f, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x74, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x68, 0x6f, 0x6d, 0x65, 0x50, 0x6c, 0x61, 0x6e, 0x65, 0x74, 0x12, 0x29,
	0x0a, 0x10, 0x70, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x5f, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x70, 0x72, 0x69, 0x6d, 0x61, 0x72,
	0x79, 0x46, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x22, 0x0a, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0e, 0x2e, 0x73, 0x74, 0x61, 0x72, 0x77, 0x61,
	0x72, 0x73, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x22, 0x3d, 0x0a,
	0x0e, 0x47, 0x65, 0x74, 0x48, 0x65, 0x72, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x2b, 0x0a, 0x07, 0x65, 0x70, 0x69, 0x73, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x11, 0x2e, 0x73, 0x74, 0x61, 0x72, 0x77, 0x61, 0x72, 0x73, 0x2e, 0x45, 0x70, 0x69, 0x73,
	0x6f, 0x64, 0x65, 0x52, 0x07, 0x65, 0x70, 0x69, 0x73, 0x6f, 0x64, 0x65, 0x22, 0x21, 0x0a, 0x0f,
	0x47, 0x65, 0x74, 0x48, 0x75, 0x6d, 0x61, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22,
	0x21, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x44, 0x72, 0x6f, 0x69, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x22, 0x12, 0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x41, 0x0a, 0x12, 0x4c, 0x69, 0x73, 0x74, 0x48, 0x75,
	0x6d, 0x61, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2b, 0x0a, 0x06,
	0x68, 0x75, 0x6d, 0x61, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x73,
	0x74, 0x61, 0x72, 0x77, 0x61, 0x72, 0x73, 0x2e, 0x43, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65,
	0x72, 0x52, 0x06, 0x68, 0x75, 0x6d, 0x61, 0x6e, 0x73, 0x22, 0x41, 0x0a, 0x12, 0x4c, 0x69, 0x73,
	0x74, 0x44, 0x72, 0x6f, 0x69, 0x64, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x2b, 0x0a, 0x06, 0x64, 0x72, 0x6f, 0x69, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x13, 0x2e, 0x73, 0x74, 0x61, 0x72, 0x77, 0x61, 0x72, 0x73, 0x2e, 0x43, 0x68, 0x61, 0x72, 0x61,
	0x63, 0x74, 0x65, 0x72, 0x52, 0x06, 0x64, 0x72, 0x6f, 0x69, 0x64, 0x73, 0x2a, 0x1c, 0x0a, 0x04,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x09, 0x0a, 0x05, 0x48, 0x55, 0x4d, 0x41, 0x4e, 0x10, 0x00, 0x12,
	0x09, 0x0a, 0x05, 0x44, 0x52, 0x4f, 0x49, 0x44, 0x10, 0x01, 0x2a, 0x33, 0x0a, 0x07, 0x45, 0x70,
	0x69, 0x73, 0x6f, 0x64, 0x65, 0x12, 0x05, 0x0a, 0x01, 0x5f, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07,
	0x4e, 0x45, 0x57, 0x48, 0x4f, 0x50, 0x45, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x45, 0x4d, 0x50,
	0x49, 0x52, 0x45, 0x10, 0x02, 0x12, 0x08, 0x0a, 0x04, 0x4a, 0x45, 0x44, 0x49, 0x10, 0x03, 0x32,
	0xd3, 0x02, 0x0a, 0x0f, 0x53, 0x74, 0x61, 0x72, 0x77, 0x61, 0x72, 0x73, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x38, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x48, 0x65, 0x72, 0x6f, 0x12, 0x18,
	0x2e, 0x73, 0x74, 0x61, 0x72, 0x77, 0x61, 0x72, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x48, 0x65, 0x72,
	0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x73, 0x74, 0x61, 0x72, 0x77,
	0x61, 0x72, 0x73, 0x2e, 0x43, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x12, 0x3a, 0x0a,
	0x08, 0x47, 0x65, 0x74, 0x48, 0x75, 0x6d, 0x61, 0x6e, 0x12, 0x19, 0x2e, 0x73, 0x74, 0x61, 0x72,
	0x77, 0x61, 0x72, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x48, 0x75, 0x6d, 0x61, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x73, 0x74, 0x61, 0x72, 0x77, 0x61, 0x72, 0x73, 0x2e,
	0x43, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x12, 0x3a, 0x0a, 0x08, 0x47, 0x65, 0x74,
	0x44, 0x72, 0x6f, 0x69, 0x64, 0x12, 0x19, 0x2e, 0x73, 0x74, 0x61, 0x72, 0x77, 0x61, 0x72, 0x73,
	0x2e, 0x47, 0x65, 0x74, 0x44, 0x72, 0x6f, 0x69, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x13, 0x2e, 0x73, 0x74, 0x61, 0x72, 0x77, 0x61, 0x72, 0x73, 0x2e, 0x43, 0x68, 0x61, 0x72,
	0x61, 0x63, 0x74, 0x65, 0x72, 0x12, 0x46, 0x0a, 0x0a, 0x4c, 0x69, 0x73, 0x74, 0x48, 0x75, 0x6d,
	0x61, 0x6e, 0x73, 0x12, 0x1a, 0x2e, 0x73, 0x74, 0x61, 0x72, 0x77, 0x61, 0x72, 0x73, 0x2e, 0x4c,
	0x69, 0x73, 0x74, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1c, 0x2e, 0x73, 0x74, 0x61, 0x72, 0x77, 0x61, 0x72, 0x73, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x48,
	0x75, 0x6d, 0x61, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x46, 0x0a,
	0x0a, 0x4c, 0x69, 0x73, 0x74, 0x44, 0x72, 0x6f, 0x69, 0x64, 0x73, 0x12, 0x1a, 0x2e, 0x73, 0x74,
	0x61, 0x72, 0x77, 0x61, 0x72, 0x73, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x73, 0x74, 0x61, 0x72, 0x77, 0x61,
	0x72, 0x73, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x44, 0x72, 0x6f, 0x69, 0x64, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x4b, 0x5a, 0x49, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x77, 0x75, 0x6e, 0x64, 0x65, 0x72, 0x67, 0x72, 0x61, 0x70, 0x68, 0x2f,
	0x77, 0x75, 0x6e, 0x64, 0x65, 0x72, 0x67, 0x72, 0x61, 0x70, 0x68, 0x2f, 0x70, 0x6b, 0x67, 0x2f,
	0x64, 0x61, 0x74, 0x61, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x67, 0x72, 0x70, 0x63,
	0x2f, 0x74, 0x65, 0x73, 0x74, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x73, 0x74, 0x61, 0x72, 0x77, 0x61,
	0x72, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_starwars_proto_rawDescOnce sync.Once
	file_starwars_proto_rawDescData = file_starwars_proto_rawDesc
)

func file_starwars_proto_rawDescGZIP() []byte {
	file_starwars_proto_rawDescOnce.Do(func() {
		file_starwars_proto_rawDescData = protoimpl.X.CompressGZIP(file_starwars_proto_rawDescData)
	})
	return file_starwars_proto_rawDescData
}

var file_starwars_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_starwars_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_starwars_proto_goTypes = []interface{}{
	(Type)(0),                  // 0: starwars.Type
	(Episode)(0),               // 1: starwars.Episode
	(*Character)(nil),          // 2: starwars.Character
	(*GetHeroRequest)(nil),     // 3: starwars.GetHeroRequest
	(*GetHumanRequest)(nil),    // 4: starwars.GetHumanRequest
	(*GetDroidRequest)(nil),    // 5: starwars.GetDroidRequest
	(*ListEmptyRequest)(nil),   // 6: starwars.ListEmptyRequest
	(*ListHumansResponse)(nil), // 7: starwars.ListHumansResponse
	(*ListDroidsResponse)(nil), // 8: starwars.ListDroidsResponse
}
var file_starwars_proto_depIdxs = []int32{
	2,  // 0: starwars.Character.friends:type_name -> starwars.Character
	1,  // 1: starwars.Character.appears_in:type_name -> starwars.Episode
	0,  // 2: starwars.Character.type:type_name -> starwars.Type
	1,  // 3: starwars.GetHeroRequest.episode:type_name -> starwars.Episode
	2,  // 4: starwars.ListHumansResponse.humans:type_name -> starwars.Character
	2,  // 5: starwars.ListDroidsResponse.droids:type_name -> starwars.Character
	3,  // 6: starwars.StarwarsService.GetHero:input_type -> starwars.GetHeroRequest
	4,  // 7: starwars.StarwarsService.GetHuman:input_type -> starwars.GetHumanRequest
	5,  // 8: starwars.StarwarsService.GetDroid:input_type -> starwars.GetDroidRequest
	6,  // 9: starwars.StarwarsService.ListHumans:input_type -> starwars.ListEmptyRequest
	6,  // 10: starwars.StarwarsService.ListDroids:input_type -> starwars.ListEmptyRequest
	2,  // 11: starwars.StarwarsService.GetHero:output_type -> starwars.Character
	2,  // 12: starwars.StarwarsService.GetHuman:output_type -> starwars.Character
	2,  // 13: starwars.StarwarsService.GetDroid:output_type -> starwars.Character
	7,  // 14: starwars.StarwarsService.ListHumans:output_type -> starwars.ListHumansResponse
	8,  // 15: starwars.StarwarsService.ListDroids:output_type -> starwars.ListDroidsResponse
	11, // [11:16] is the sub-list for method output_type
	6,  // [6:11] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_starwars_proto_init() }
func file_starwars_proto_init() {
	if File_starwars_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_starwars_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Character); i {
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
		file_starwars_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetHeroRequest); i {
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
		file_starwars_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetHumanRequest); i {
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
		file_starwars_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetDroidRequest); i {
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
		file_starwars_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListEmptyRequest); i {
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
		file_starwars_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListHumansResponse); i {
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
		file_starwars_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListDroidsResponse); i {
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
			RawDescriptor: file_starwars_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_starwars_proto_goTypes,
		DependencyIndexes: file_starwars_proto_depIdxs,
		EnumInfos:         file_starwars_proto_enumTypes,
		MessageInfos:      file_starwars_proto_msgTypes,
	}.Build()
	File_starwars_proto = out.File
	file_starwars_proto_rawDesc = nil
	file_starwars_proto_goTypes = nil
	file_starwars_proto_depIdxs = nil
}