// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.29.3
// source: proto/service.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Pokemon struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name          string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Type          []string               `protobuf:"bytes,3,rep,name=type,proto3" json:"type,omitempty"`
	Abilities     []string               `protobuf:"bytes,4,rep,name=abilities,proto3" json:"abilities,omitempty"`
	Height        float64                `protobuf:"fixed64,5,opt,name=height,proto3" json:"height,omitempty"`
	Weight        float64                `protobuf:"fixed64,6,opt,name=weight,proto3" json:"weight,omitempty"`
	PokedexId     int32                  `protobuf:"varint,7,opt,name=pokedex_id,json=pokedexId,proto3" json:"pokedex_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Pokemon) Reset() {
	*x = Pokemon{}
	mi := &file_proto_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Pokemon) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Pokemon) ProtoMessage() {}

func (x *Pokemon) ProtoReflect() protoreflect.Message {
	mi := &file_proto_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Pokemon.ProtoReflect.Descriptor instead.
func (*Pokemon) Descriptor() ([]byte, []int) {
	return file_proto_service_proto_rawDescGZIP(), []int{0}
}

func (x *Pokemon) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Pokemon) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Pokemon) GetType() []string {
	if x != nil {
		return x.Type
	}
	return nil
}

func (x *Pokemon) GetAbilities() []string {
	if x != nil {
		return x.Abilities
	}
	return nil
}

func (x *Pokemon) GetHeight() float64 {
	if x != nil {
		return x.Height
	}
	return 0
}

func (x *Pokemon) GetWeight() float64 {
	if x != nil {
		return x.Weight
	}
	return 0
}

func (x *Pokemon) GetPokedexId() int32 {
	if x != nil {
		return x.PokedexId
	}
	return 0
}

type CreatePokemonRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Name          string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Type          []string               `protobuf:"bytes,2,rep,name=type,proto3" json:"type,omitempty"`
	Abilities     []string               `protobuf:"bytes,3,rep,name=abilities,proto3" json:"abilities,omitempty"`
	Height        float64                `protobuf:"fixed64,4,opt,name=height,proto3" json:"height,omitempty"`
	Weight        float64                `protobuf:"fixed64,5,opt,name=weight,proto3" json:"weight,omitempty"`
	PokedexId     int32                  `protobuf:"varint,6,opt,name=pokedex_id,json=pokedexId,proto3" json:"pokedex_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreatePokemonRequest) Reset() {
	*x = CreatePokemonRequest{}
	mi := &file_proto_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreatePokemonRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePokemonRequest) ProtoMessage() {}

func (x *CreatePokemonRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePokemonRequest.ProtoReflect.Descriptor instead.
func (*CreatePokemonRequest) Descriptor() ([]byte, []int) {
	return file_proto_service_proto_rawDescGZIP(), []int{1}
}

func (x *CreatePokemonRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreatePokemonRequest) GetType() []string {
	if x != nil {
		return x.Type
	}
	return nil
}

func (x *CreatePokemonRequest) GetAbilities() []string {
	if x != nil {
		return x.Abilities
	}
	return nil
}

func (x *CreatePokemonRequest) GetHeight() float64 {
	if x != nil {
		return x.Height
	}
	return 0
}

func (x *CreatePokemonRequest) GetWeight() float64 {
	if x != nil {
		return x.Weight
	}
	return 0
}

func (x *CreatePokemonRequest) GetPokedexId() int32 {
	if x != nil {
		return x.PokedexId
	}
	return 0
}

type GetPokemonRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetPokemonRequest) Reset() {
	*x = GetPokemonRequest{}
	mi := &file_proto_service_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetPokemonRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPokemonRequest) ProtoMessage() {}

func (x *GetPokemonRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_service_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPokemonRequest.ProtoReflect.Descriptor instead.
func (*GetPokemonRequest) Descriptor() ([]byte, []int) {
	return file_proto_service_proto_rawDescGZIP(), []int{2}
}

func (x *GetPokemonRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type ListPokemonsRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// For future pagination support
	PageSize      int32 `protobuf:"varint,1,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	PageNumber    int32 `protobuf:"varint,2,opt,name=page_number,json=pageNumber,proto3" json:"page_number,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListPokemonsRequest) Reset() {
	*x = ListPokemonsRequest{}
	mi := &file_proto_service_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListPokemonsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListPokemonsRequest) ProtoMessage() {}

func (x *ListPokemonsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_service_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListPokemonsRequest.ProtoReflect.Descriptor instead.
func (*ListPokemonsRequest) Descriptor() ([]byte, []int) {
	return file_proto_service_proto_rawDescGZIP(), []int{3}
}

func (x *ListPokemonsRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *ListPokemonsRequest) GetPageNumber() int32 {
	if x != nil {
		return x.PageNumber
	}
	return 0
}

type ListPokemonsResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Pokemons      []*Pokemon             `protobuf:"bytes,1,rep,name=pokemons,proto3" json:"pokemons,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListPokemonsResponse) Reset() {
	*x = ListPokemonsResponse{}
	mi := &file_proto_service_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListPokemonsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListPokemonsResponse) ProtoMessage() {}

func (x *ListPokemonsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_service_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListPokemonsResponse.ProtoReflect.Descriptor instead.
func (*ListPokemonsResponse) Descriptor() ([]byte, []int) {
	return file_proto_service_proto_rawDescGZIP(), []int{4}
}

func (x *ListPokemonsResponse) GetPokemons() []*Pokemon {
	if x != nil {
		return x.Pokemons
	}
	return nil
}

type UpdatePokemonRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Pokemon       *Pokemon               `protobuf:"bytes,2,opt,name=pokemon,proto3" json:"pokemon,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdatePokemonRequest) Reset() {
	*x = UpdatePokemonRequest{}
	mi := &file_proto_service_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdatePokemonRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdatePokemonRequest) ProtoMessage() {}

func (x *UpdatePokemonRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_service_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdatePokemonRequest.ProtoReflect.Descriptor instead.
func (*UpdatePokemonRequest) Descriptor() ([]byte, []int) {
	return file_proto_service_proto_rawDescGZIP(), []int{5}
}

func (x *UpdatePokemonRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdatePokemonRequest) GetPokemon() *Pokemon {
	if x != nil {
		return x.Pokemon
	}
	return nil
}

type DeletePokemonRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeletePokemonRequest) Reset() {
	*x = DeletePokemonRequest{}
	mi := &file_proto_service_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeletePokemonRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeletePokemonRequest) ProtoMessage() {}

func (x *DeletePokemonRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_service_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeletePokemonRequest.ProtoReflect.Descriptor instead.
func (*DeletePokemonRequest) Descriptor() ([]byte, []int) {
	return file_proto_service_proto_rawDescGZIP(), []int{6}
}

func (x *DeletePokemonRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type DeletePokemonResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Success       bool                   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Message       string                 `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeletePokemonResponse) Reset() {
	*x = DeletePokemonResponse{}
	mi := &file_proto_service_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeletePokemonResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeletePokemonResponse) ProtoMessage() {}

func (x *DeletePokemonResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_service_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeletePokemonResponse.ProtoReflect.Descriptor instead.
func (*DeletePokemonResponse) Descriptor() ([]byte, []int) {
	return file_proto_service_proto_rawDescGZIP(), []int{7}
}

func (x *DeletePokemonResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *DeletePokemonResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type PokemonResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Pokemon       *Pokemon               `protobuf:"bytes,1,opt,name=pokemon,proto3" json:"pokemon,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PokemonResponse) Reset() {
	*x = PokemonResponse{}
	mi := &file_proto_service_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PokemonResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PokemonResponse) ProtoMessage() {}

func (x *PokemonResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_service_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PokemonResponse.ProtoReflect.Descriptor instead.
func (*PokemonResponse) Descriptor() ([]byte, []int) {
	return file_proto_service_proto_rawDescGZIP(), []int{8}
}

func (x *PokemonResponse) GetPokemon() *Pokemon {
	if x != nil {
		return x.Pokemon
	}
	return nil
}

type User struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Username      string                 `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	Password      string                 `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"` // Should be hashed when stored
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *User) Reset() {
	*x = User{}
	mi := &file_proto_service_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_proto_service_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_proto_service_proto_rawDescGZIP(), []int{9}
}

func (x *User) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *User) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *User) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type CreateUserRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Username      string                 `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password      string                 `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateUserRequest) Reset() {
	*x = CreateUserRequest{}
	mi := &file_proto_service_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateUserRequest) ProtoMessage() {}

func (x *CreateUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_service_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateUserRequest.ProtoReflect.Descriptor instead.
func (*CreateUserRequest) Descriptor() ([]byte, []int) {
	return file_proto_service_proto_rawDescGZIP(), []int{10}
}

func (x *CreateUserRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *CreateUserRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type GetUserRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetUserRequest) Reset() {
	*x = GetUserRequest{}
	mi := &file_proto_service_proto_msgTypes[11]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserRequest) ProtoMessage() {}

func (x *GetUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_service_proto_msgTypes[11]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserRequest.ProtoReflect.Descriptor instead.
func (*GetUserRequest) Descriptor() ([]byte, []int) {
	return file_proto_service_proto_rawDescGZIP(), []int{11}
}

func (x *GetUserRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetUserByUsernameRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Username      string                 `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetUserByUsernameRequest) Reset() {
	*x = GetUserByUsernameRequest{}
	mi := &file_proto_service_proto_msgTypes[12]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetUserByUsernameRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserByUsernameRequest) ProtoMessage() {}

func (x *GetUserByUsernameRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_service_proto_msgTypes[12]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserByUsernameRequest.ProtoReflect.Descriptor instead.
func (*GetUserByUsernameRequest) Descriptor() ([]byte, []int) {
	return file_proto_service_proto_rawDescGZIP(), []int{12}
}

func (x *GetUserByUsernameRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

type UserResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	User          *User                  `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UserResponse) Reset() {
	*x = UserResponse{}
	mi := &file_proto_service_proto_msgTypes[13]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserResponse) ProtoMessage() {}

func (x *UserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_service_proto_msgTypes[13]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserResponse.ProtoReflect.Descriptor instead.
func (*UserResponse) Descriptor() ([]byte, []int) {
	return file_proto_service_proto_rawDescGZIP(), []int{13}
}

func (x *UserResponse) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

type LoginRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Username      string                 `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password      string                 `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *LoginRequest) Reset() {
	*x = LoginRequest{}
	mi := &file_proto_service_proto_msgTypes[14]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LoginRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginRequest) ProtoMessage() {}

func (x *LoginRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_service_proto_msgTypes[14]
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
	return file_proto_service_proto_rawDescGZIP(), []int{14}
}

func (x *LoginRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *LoginRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type LoginResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Token         string                 `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *LoginResponse) Reset() {
	*x = LoginResponse{}
	mi := &file_proto_service_proto_msgTypes[15]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LoginResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginResponse) ProtoMessage() {}

func (x *LoginResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_service_proto_msgTypes[15]
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
	return file_proto_service_proto_rawDescGZIP(), []int{15}
}

func (x *LoginResponse) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

var File_proto_service_proto protoreflect.FileDescriptor

const file_proto_service_proto_rawDesc = "" +
	"\n" +
	"\x13proto/service.proto\x12\x04grpc\"\xae\x01\n" +
	"\aPokemon\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\x12\x12\n" +
	"\x04name\x18\x02 \x01(\tR\x04name\x12\x12\n" +
	"\x04type\x18\x03 \x03(\tR\x04type\x12\x1c\n" +
	"\tabilities\x18\x04 \x03(\tR\tabilities\x12\x16\n" +
	"\x06height\x18\x05 \x01(\x01R\x06height\x12\x16\n" +
	"\x06weight\x18\x06 \x01(\x01R\x06weight\x12\x1d\n" +
	"\n" +
	"pokedex_id\x18\a \x01(\x05R\tpokedexId\"\xab\x01\n" +
	"\x14CreatePokemonRequest\x12\x12\n" +
	"\x04name\x18\x01 \x01(\tR\x04name\x12\x12\n" +
	"\x04type\x18\x02 \x03(\tR\x04type\x12\x1c\n" +
	"\tabilities\x18\x03 \x03(\tR\tabilities\x12\x16\n" +
	"\x06height\x18\x04 \x01(\x01R\x06height\x12\x16\n" +
	"\x06weight\x18\x05 \x01(\x01R\x06weight\x12\x1d\n" +
	"\n" +
	"pokedex_id\x18\x06 \x01(\x05R\tpokedexId\"#\n" +
	"\x11GetPokemonRequest\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\"S\n" +
	"\x13ListPokemonsRequest\x12\x1b\n" +
	"\tpage_size\x18\x01 \x01(\x05R\bpageSize\x12\x1f\n" +
	"\vpage_number\x18\x02 \x01(\x05R\n" +
	"pageNumber\"A\n" +
	"\x14ListPokemonsResponse\x12)\n" +
	"\bpokemons\x18\x01 \x03(\v2\r.grpc.PokemonR\bpokemons\"O\n" +
	"\x14UpdatePokemonRequest\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\x12'\n" +
	"\apokemon\x18\x02 \x01(\v2\r.grpc.PokemonR\apokemon\"&\n" +
	"\x14DeletePokemonRequest\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\"K\n" +
	"\x15DeletePokemonResponse\x12\x18\n" +
	"\asuccess\x18\x01 \x01(\bR\asuccess\x12\x18\n" +
	"\amessage\x18\x02 \x01(\tR\amessage\":\n" +
	"\x0fPokemonResponse\x12'\n" +
	"\apokemon\x18\x01 \x01(\v2\r.grpc.PokemonR\apokemon\"N\n" +
	"\x04User\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\x12\x1a\n" +
	"\busername\x18\x02 \x01(\tR\busername\x12\x1a\n" +
	"\bpassword\x18\x03 \x01(\tR\bpassword\"K\n" +
	"\x11CreateUserRequest\x12\x1a\n" +
	"\busername\x18\x01 \x01(\tR\busername\x12\x1a\n" +
	"\bpassword\x18\x02 \x01(\tR\bpassword\" \n" +
	"\x0eGetUserRequest\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\"6\n" +
	"\x18GetUserByUsernameRequest\x12\x1a\n" +
	"\busername\x18\x01 \x01(\tR\busername\".\n" +
	"\fUserResponse\x12\x1e\n" +
	"\x04user\x18\x01 \x01(\v2\n" +
	".grpc.UserR\x04user\"F\n" +
	"\fLoginRequest\x12\x1a\n" +
	"\busername\x18\x01 \x01(\tR\busername\x12\x1a\n" +
	"\bpassword\x18\x02 \x01(\tR\bpassword\"%\n" +
	"\rLoginResponse\x12\x14\n" +
	"\x05token\x18\x01 \x01(\tR\x05token2\xf1\x02\n" +
	"\x0ePokemonService\x12D\n" +
	"\rCreatePokemon\x12\x1a.grpc.CreatePokemonRequest\x1a\x15.grpc.PokemonResponse\"\x00\x12>\n" +
	"\n" +
	"GetPokemon\x12\x17.grpc.GetPokemonRequest\x1a\x15.grpc.PokemonResponse\"\x00\x12G\n" +
	"\fListPokemons\x12\x19.grpc.ListPokemonsRequest\x1a\x1a.grpc.ListPokemonsResponse\"\x00\x12D\n" +
	"\rUpdatePokemon\x12\x1a.grpc.UpdatePokemonRequest\x1a\x15.grpc.PokemonResponse\"\x00\x12J\n" +
	"\rDeletePokemon\x12\x1a.grpc.DeletePokemonRequest\x1a\x1b.grpc.DeletePokemonResponse\"\x002\x80\x02\n" +
	"\vUserService\x12;\n" +
	"\n" +
	"CreateUser\x12\x17.grpc.CreateUserRequest\x1a\x12.grpc.UserResponse\"\x00\x125\n" +
	"\aGetUser\x12\x14.grpc.GetUserRequest\x1a\x12.grpc.UserResponse\"\x00\x12I\n" +
	"\x11GetUserByUsername\x12\x1e.grpc.GetUserByUsernameRequest\x1a\x12.grpc.UserResponse\"\x00\x122\n" +
	"\x05Login\x12\x12.grpc.LoginRequest\x1a\x13.grpc.LoginResponse\"\x00B Z\x1eproject-is/services/grpc/protob\x06proto3"

var (
	file_proto_service_proto_rawDescOnce sync.Once
	file_proto_service_proto_rawDescData []byte
)

func file_proto_service_proto_rawDescGZIP() []byte {
	file_proto_service_proto_rawDescOnce.Do(func() {
		file_proto_service_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_proto_service_proto_rawDesc), len(file_proto_service_proto_rawDesc)))
	})
	return file_proto_service_proto_rawDescData
}

var file_proto_service_proto_msgTypes = make([]protoimpl.MessageInfo, 16)
var file_proto_service_proto_goTypes = []any{
	(*Pokemon)(nil),                  // 0: grpc.Pokemon
	(*CreatePokemonRequest)(nil),     // 1: grpc.CreatePokemonRequest
	(*GetPokemonRequest)(nil),        // 2: grpc.GetPokemonRequest
	(*ListPokemonsRequest)(nil),      // 3: grpc.ListPokemonsRequest
	(*ListPokemonsResponse)(nil),     // 4: grpc.ListPokemonsResponse
	(*UpdatePokemonRequest)(nil),     // 5: grpc.UpdatePokemonRequest
	(*DeletePokemonRequest)(nil),     // 6: grpc.DeletePokemonRequest
	(*DeletePokemonResponse)(nil),    // 7: grpc.DeletePokemonResponse
	(*PokemonResponse)(nil),          // 8: grpc.PokemonResponse
	(*User)(nil),                     // 9: grpc.User
	(*CreateUserRequest)(nil),        // 10: grpc.CreateUserRequest
	(*GetUserRequest)(nil),           // 11: grpc.GetUserRequest
	(*GetUserByUsernameRequest)(nil), // 12: grpc.GetUserByUsernameRequest
	(*UserResponse)(nil),             // 13: grpc.UserResponse
	(*LoginRequest)(nil),             // 14: grpc.LoginRequest
	(*LoginResponse)(nil),            // 15: grpc.LoginResponse
}
var file_proto_service_proto_depIdxs = []int32{
	0,  // 0: grpc.ListPokemonsResponse.pokemons:type_name -> grpc.Pokemon
	0,  // 1: grpc.UpdatePokemonRequest.pokemon:type_name -> grpc.Pokemon
	0,  // 2: grpc.PokemonResponse.pokemon:type_name -> grpc.Pokemon
	9,  // 3: grpc.UserResponse.user:type_name -> grpc.User
	1,  // 4: grpc.PokemonService.CreatePokemon:input_type -> grpc.CreatePokemonRequest
	2,  // 5: grpc.PokemonService.GetPokemon:input_type -> grpc.GetPokemonRequest
	3,  // 6: grpc.PokemonService.ListPokemons:input_type -> grpc.ListPokemonsRequest
	5,  // 7: grpc.PokemonService.UpdatePokemon:input_type -> grpc.UpdatePokemonRequest
	6,  // 8: grpc.PokemonService.DeletePokemon:input_type -> grpc.DeletePokemonRequest
	10, // 9: grpc.UserService.CreateUser:input_type -> grpc.CreateUserRequest
	11, // 10: grpc.UserService.GetUser:input_type -> grpc.GetUserRequest
	12, // 11: grpc.UserService.GetUserByUsername:input_type -> grpc.GetUserByUsernameRequest
	14, // 12: grpc.UserService.Login:input_type -> grpc.LoginRequest
	8,  // 13: grpc.PokemonService.CreatePokemon:output_type -> grpc.PokemonResponse
	8,  // 14: grpc.PokemonService.GetPokemon:output_type -> grpc.PokemonResponse
	4,  // 15: grpc.PokemonService.ListPokemons:output_type -> grpc.ListPokemonsResponse
	8,  // 16: grpc.PokemonService.UpdatePokemon:output_type -> grpc.PokemonResponse
	7,  // 17: grpc.PokemonService.DeletePokemon:output_type -> grpc.DeletePokemonResponse
	13, // 18: grpc.UserService.CreateUser:output_type -> grpc.UserResponse
	13, // 19: grpc.UserService.GetUser:output_type -> grpc.UserResponse
	13, // 20: grpc.UserService.GetUserByUsername:output_type -> grpc.UserResponse
	15, // 21: grpc.UserService.Login:output_type -> grpc.LoginResponse
	13, // [13:22] is the sub-list for method output_type
	4,  // [4:13] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_proto_service_proto_init() }
func file_proto_service_proto_init() {
	if File_proto_service_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_proto_service_proto_rawDesc), len(file_proto_service_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   16,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_proto_service_proto_goTypes,
		DependencyIndexes: file_proto_service_proto_depIdxs,
		MessageInfos:      file_proto_service_proto_msgTypes,
	}.Build()
	File_proto_service_proto = out.File
	file_proto_service_proto_goTypes = nil
	file_proto_service_proto_depIdxs = nil
}
