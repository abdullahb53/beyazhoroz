// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.20.1
// source: furkan_service.proto

package api

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

type ResponseState struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status int32      `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Mesaj  string     `protobuf:"bytes,2,opt,name=mesaj,proto3" json:"mesaj,omitempty"`
	Data   *UsersBook `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *ResponseState) Reset() {
	*x = ResponseState{}
	if protoimpl.UnsafeEnabled {
		mi := &file_furkan_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResponseState) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponseState) ProtoMessage() {}

func (x *ResponseState) ProtoReflect() protoreflect.Message {
	mi := &file_furkan_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponseState.ProtoReflect.Descriptor instead.
func (*ResponseState) Descriptor() ([]byte, []int) {
	return file_furkan_service_proto_rawDescGZIP(), []int{0}
}

func (x *ResponseState) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *ResponseState) GetMesaj() string {
	if x != nil {
		return x.Mesaj
	}
	return ""
}

func (x *ResponseState) GetData() *UsersBook {
	if x != nil {
		return x.Data
	}
	return nil
}

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid     string `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Name    string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Email   string `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	City    string `protobuf:"bytes,4,opt,name=city,proto3" json:"city,omitempty"`
	Country string `protobuf:"bytes,5,opt,name=country,proto3" json:"country,omitempty"`
	Explain string `protobuf:"bytes,6,opt,name=explain,proto3" json:"explain,omitempty"`
	Phone   string `protobuf:"bytes,7,opt,name=phone,proto3" json:"phone,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_furkan_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_furkan_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
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
	return file_furkan_service_proto_rawDescGZIP(), []int{1}
}

func (x *User) GetUid() string {
	if x != nil {
		return x.Uid
	}
	return ""
}

func (x *User) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *User) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *User) GetCity() string {
	if x != nil {
		return x.City
	}
	return ""
}

func (x *User) GetCountry() string {
	if x != nil {
		return x.Country
	}
	return ""
}

func (x *User) GetExplain() string {
	if x != nil {
		return x.Explain
	}
	return ""
}

func (x *User) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

// Our address book file is just one of these.
type UsersBook struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	People []*User `protobuf:"bytes,1,rep,name=people,proto3" json:"people,omitempty"`
}

func (x *UsersBook) Reset() {
	*x = UsersBook{}
	if protoimpl.UnsafeEnabled {
		mi := &file_furkan_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UsersBook) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UsersBook) ProtoMessage() {}

func (x *UsersBook) ProtoReflect() protoreflect.Message {
	mi := &file_furkan_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UsersBook.ProtoReflect.Descriptor instead.
func (*UsersBook) Descriptor() ([]byte, []int) {
	return file_furkan_service_proto_rawDescGZIP(), []int{2}
}

func (x *UsersBook) GetPeople() []*User {
	if x != nil {
		return x.People
	}
	return nil
}

type UserCreateReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PersonItem *User `protobuf:"bytes,1,opt,name=person_item,json=personItem,proto3" json:"person_item,omitempty"`
}

func (x *UserCreateReq) Reset() {
	*x = UserCreateReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_furkan_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserCreateReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserCreateReq) ProtoMessage() {}

func (x *UserCreateReq) ProtoReflect() protoreflect.Message {
	mi := &file_furkan_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserCreateReq.ProtoReflect.Descriptor instead.
func (*UserCreateReq) Descriptor() ([]byte, []int) {
	return file_furkan_service_proto_rawDescGZIP(), []int{3}
}

func (x *UserCreateReq) GetPersonItem() *User {
	if x != nil {
		return x.PersonItem
	}
	return nil
}

type UserCreateRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CreatedUserItems *ResponseState `protobuf:"bytes,1,opt,name=createdUserItems,proto3" json:"createdUserItems,omitempty"`
}

func (x *UserCreateRes) Reset() {
	*x = UserCreateRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_furkan_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserCreateRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserCreateRes) ProtoMessage() {}

func (x *UserCreateRes) ProtoReflect() protoreflect.Message {
	mi := &file_furkan_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserCreateRes.ProtoReflect.Descriptor instead.
func (*UserCreateRes) Descriptor() ([]byte, []int) {
	return file_furkan_service_proto_rawDescGZIP(), []int{4}
}

func (x *UserCreateRes) GetCreatedUserItems() *ResponseState {
	if x != nil {
		return x.CreatedUserItems
	}
	return nil
}

type UserListReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	City    string `protobuf:"bytes,1,opt,name=city,proto3" json:"city,omitempty"`
	Country string `protobuf:"bytes,2,opt,name=country,proto3" json:"country,omitempty"`
}

func (x *UserListReq) Reset() {
	*x = UserListReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_furkan_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserListReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserListReq) ProtoMessage() {}

func (x *UserListReq) ProtoReflect() protoreflect.Message {
	mi := &file_furkan_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserListReq.ProtoReflect.Descriptor instead.
func (*UserListReq) Descriptor() ([]byte, []int) {
	return file_furkan_service_proto_rawDescGZIP(), []int{5}
}

func (x *UserListReq) GetCity() string {
	if x != nil {
		return x.City
	}
	return ""
}

func (x *UserListReq) GetCountry() string {
	if x != nil {
		return x.Country
	}
	return ""
}

type UserListRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ListedUserItems *ResponseState `protobuf:"bytes,1,opt,name=ListedUserItems,proto3" json:"ListedUserItems,omitempty"`
}

func (x *UserListRes) Reset() {
	*x = UserListRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_furkan_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserListRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserListRes) ProtoMessage() {}

func (x *UserListRes) ProtoReflect() protoreflect.Message {
	mi := &file_furkan_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserListRes.ProtoReflect.Descriptor instead.
func (*UserListRes) Descriptor() ([]byte, []int) {
	return file_furkan_service_proto_rawDescGZIP(), []int{6}
}

func (x *UserListRes) GetListedUserItems() *ResponseState {
	if x != nil {
		return x.ListedUserItems
	}
	return nil
}

type UserFindReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid string `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
}

func (x *UserFindReq) Reset() {
	*x = UserFindReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_furkan_service_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserFindReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserFindReq) ProtoMessage() {}

func (x *UserFindReq) ProtoReflect() protoreflect.Message {
	mi := &file_furkan_service_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserFindReq.ProtoReflect.Descriptor instead.
func (*UserFindReq) Descriptor() ([]byte, []int) {
	return file_furkan_service_proto_rawDescGZIP(), []int{7}
}

func (x *UserFindReq) GetUid() string {
	if x != nil {
		return x.Uid
	}
	return ""
}

type UserFindRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FindedUserItem *ResponseState `protobuf:"bytes,1,opt,name=FindedUserItem,proto3" json:"FindedUserItem,omitempty"`
}

func (x *UserFindRes) Reset() {
	*x = UserFindRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_furkan_service_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserFindRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserFindRes) ProtoMessage() {}

func (x *UserFindRes) ProtoReflect() protoreflect.Message {
	mi := &file_furkan_service_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserFindRes.ProtoReflect.Descriptor instead.
func (*UserFindRes) Descriptor() ([]byte, []int) {
	return file_furkan_service_proto_rawDescGZIP(), []int{8}
}

func (x *UserFindRes) GetFindedUserItem() *ResponseState {
	if x != nil {
		return x.FindedUserItem
	}
	return nil
}

type UserDeleteReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid string `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
}

func (x *UserDeleteReq) Reset() {
	*x = UserDeleteReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_furkan_service_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserDeleteReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserDeleteReq) ProtoMessage() {}

func (x *UserDeleteReq) ProtoReflect() protoreflect.Message {
	mi := &file_furkan_service_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserDeleteReq.ProtoReflect.Descriptor instead.
func (*UserDeleteReq) Descriptor() ([]byte, []int) {
	return file_furkan_service_proto_rawDescGZIP(), []int{9}
}

func (x *UserDeleteReq) GetUid() string {
	if x != nil {
		return x.Uid
	}
	return ""
}

type UserDeleteRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserDeletedUserItem *ResponseState `protobuf:"bytes,1,opt,name=UserDeletedUserItem,proto3" json:"UserDeletedUserItem,omitempty"`
}

func (x *UserDeleteRes) Reset() {
	*x = UserDeleteRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_furkan_service_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserDeleteRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserDeleteRes) ProtoMessage() {}

func (x *UserDeleteRes) ProtoReflect() protoreflect.Message {
	mi := &file_furkan_service_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserDeleteRes.ProtoReflect.Descriptor instead.
func (*UserDeleteRes) Descriptor() ([]byte, []int) {
	return file_furkan_service_proto_rawDescGZIP(), []int{10}
}

func (x *UserDeleteRes) GetUserDeletedUserItem() *ResponseState {
	if x != nil {
		return x.UserDeletedUserItem
	}
	return nil
}

type UserUpdateReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PersonItem *User `protobuf:"bytes,1,opt,name=person_item,json=personItem,proto3" json:"person_item,omitempty"`
}

func (x *UserUpdateReq) Reset() {
	*x = UserUpdateReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_furkan_service_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserUpdateReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserUpdateReq) ProtoMessage() {}

func (x *UserUpdateReq) ProtoReflect() protoreflect.Message {
	mi := &file_furkan_service_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserUpdateReq.ProtoReflect.Descriptor instead.
func (*UserUpdateReq) Descriptor() ([]byte, []int) {
	return file_furkan_service_proto_rawDescGZIP(), []int{11}
}

func (x *UserUpdateReq) GetPersonItem() *User {
	if x != nil {
		return x.PersonItem
	}
	return nil
}

type UserUpdateRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserUpdatedItems *ResponseState `protobuf:"bytes,1,opt,name=UserUpdatedItems,proto3" json:"UserUpdatedItems,omitempty"`
}

func (x *UserUpdateRes) Reset() {
	*x = UserUpdateRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_furkan_service_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserUpdateRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserUpdateRes) ProtoMessage() {}

func (x *UserUpdateRes) ProtoReflect() protoreflect.Message {
	mi := &file_furkan_service_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserUpdateRes.ProtoReflect.Descriptor instead.
func (*UserUpdateRes) Descriptor() ([]byte, []int) {
	return file_furkan_service_proto_rawDescGZIP(), []int{12}
}

func (x *UserUpdateRes) GetUserUpdatedItems() *ResponseState {
	if x != nil {
		return x.UserUpdatedItems
	}
	return nil
}

var File_furkan_service_proto protoreflect.FileDescriptor

var file_furkan_service_proto_rawDesc = []byte{
	0x0a, 0x14, 0x66, 0x75, 0x72, 0x6b, 0x61, 0x6e, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x5d, 0x0a, 0x0d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x14, 0x0a, 0x05, 0x6d, 0x65, 0x73, 0x61, 0x6a, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x6d, 0x65, 0x73, 0x61, 0x6a, 0x12, 0x1e, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x73, 0x42, 0x6f, 0x6f, 0x6b, 0x52,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0xa0, 0x01, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x10,
	0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x69, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x69,
	0x74, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x69, 0x74, 0x79, 0x12, 0x18,
	0x0a, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x78, 0x70, 0x6c,
	0x61, 0x69, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x65, 0x78, 0x70, 0x6c, 0x61,
	0x69, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x22, 0x2a, 0x0a, 0x09, 0x55, 0x73, 0x65, 0x72,
	0x73, 0x42, 0x6f, 0x6f, 0x6b, 0x12, 0x1d, 0x0a, 0x06, 0x70, 0x65, 0x6f, 0x70, 0x6c, 0x65, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x05, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x06, 0x70, 0x65,
	0x6f, 0x70, 0x6c, 0x65, 0x22, 0x37, 0x0a, 0x0d, 0x55, 0x73, 0x65, 0x72, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x52, 0x65, 0x71, 0x12, 0x26, 0x0a, 0x0b, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x5f,
	0x69, 0x74, 0x65, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x05, 0x2e, 0x55, 0x73, 0x65,
	0x72, 0x52, 0x0a, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x49, 0x74, 0x65, 0x6d, 0x22, 0x4b, 0x0a,
	0x0d, 0x55, 0x73, 0x65, 0x72, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x12, 0x3a,
	0x0a, 0x10, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x55, 0x73, 0x65, 0x72, 0x49, 0x74, 0x65,
	0x6d, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x10, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x55, 0x73, 0x65, 0x72, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x22, 0x3b, 0x0a, 0x0b, 0x55, 0x73,
	0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x69, 0x74,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x69, 0x74, 0x79, 0x12, 0x18, 0x0a,
	0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x22, 0x47, 0x0a, 0x0b, 0x55, 0x73, 0x65, 0x72, 0x4c,
	0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x12, 0x38, 0x0a, 0x0f, 0x4c, 0x69, 0x73, 0x74, 0x65, 0x64,
	0x55, 0x73, 0x65, 0x72, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0e, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52,
	0x0f, 0x4c, 0x69, 0x73, 0x74, 0x65, 0x64, 0x55, 0x73, 0x65, 0x72, 0x49, 0x74, 0x65, 0x6d, 0x73,
	0x22, 0x1f, 0x0a, 0x0b, 0x55, 0x73, 0x65, 0x72, 0x46, 0x69, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x12,
	0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x69,
	0x64, 0x22, 0x45, 0x0a, 0x0b, 0x55, 0x73, 0x65, 0x72, 0x46, 0x69, 0x6e, 0x64, 0x52, 0x65, 0x73,
	0x12, 0x36, 0x0a, 0x0e, 0x46, 0x69, 0x6e, 0x64, 0x65, 0x64, 0x55, 0x73, 0x65, 0x72, 0x49, 0x74,
	0x65, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x0e, 0x46, 0x69, 0x6e, 0x64, 0x65, 0x64,
	0x55, 0x73, 0x65, 0x72, 0x49, 0x74, 0x65, 0x6d, 0x22, 0x21, 0x0a, 0x0d, 0x55, 0x73, 0x65, 0x72,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x69, 0x64, 0x22, 0x51, 0x0a, 0x0d, 0x55,
	0x73, 0x65, 0x72, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x12, 0x40, 0x0a, 0x13,
	0x55, 0x73, 0x65, 0x72, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x55, 0x73, 0x65, 0x72, 0x49,
	0x74, 0x65, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x13, 0x55, 0x73, 0x65, 0x72, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x55, 0x73, 0x65, 0x72, 0x49, 0x74, 0x65, 0x6d, 0x22, 0x37,
	0x0a, 0x0d, 0x55, 0x73, 0x65, 0x72, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x12,
	0x26, 0x0a, 0x0b, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x5f, 0x69, 0x74, 0x65, 0x6d, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x05, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x0a, 0x70, 0x65, 0x72,
	0x73, 0x6f, 0x6e, 0x49, 0x74, 0x65, 0x6d, 0x22, 0x4b, 0x0a, 0x0d, 0x55, 0x73, 0x65, 0x72, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x12, 0x3a, 0x0a, 0x10, 0x55, 0x73, 0x65, 0x72,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x53, 0x74, 0x61,
	0x74, 0x65, 0x52, 0x10, 0x55, 0x73, 0x65, 0x72, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x49,
	0x74, 0x65, 0x6d, 0x73, 0x32, 0xea, 0x01, 0x0a, 0x0d, 0x46, 0x75, 0x72, 0x6b, 0x61, 0x6e, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x2c, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x55, 0x73, 0x65, 0x72, 0x12, 0x0e, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x52, 0x65, 0x71, 0x1a, 0x0e, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x52, 0x65, 0x73, 0x12, 0x27, 0x0a, 0x09, 0x4c, 0x69, 0x73, 0x74, 0x55, 0x73, 0x65, 0x72,
	0x73, 0x12, 0x0c, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x1a,
	0x0c, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x12, 0x26, 0x0a,
	0x08, 0x46, 0x69, 0x6e, 0x64, 0x55, 0x73, 0x65, 0x72, 0x12, 0x0c, 0x2e, 0x55, 0x73, 0x65, 0x72,
	0x46, 0x69, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x1a, 0x0c, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x46, 0x69,
	0x6e, 0x64, 0x52, 0x65, 0x73, 0x12, 0x2c, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x55,
	0x73, 0x65, 0x72, 0x12, 0x0e, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x52, 0x65, 0x71, 0x1a, 0x0e, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x52, 0x65, 0x73, 0x12, 0x2c, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65,
	0x72, 0x12, 0x0e, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x71, 0x1a, 0x0e, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x73, 0x42, 0x06, 0x5a, 0x04, 0x61, 0x70, 0x69, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_furkan_service_proto_rawDescOnce sync.Once
	file_furkan_service_proto_rawDescData = file_furkan_service_proto_rawDesc
)

func file_furkan_service_proto_rawDescGZIP() []byte {
	file_furkan_service_proto_rawDescOnce.Do(func() {
		file_furkan_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_furkan_service_proto_rawDescData)
	})
	return file_furkan_service_proto_rawDescData
}

var file_furkan_service_proto_msgTypes = make([]protoimpl.MessageInfo, 13)
var file_furkan_service_proto_goTypes = []interface{}{
	(*ResponseState)(nil), // 0: ResponseState
	(*User)(nil),          // 1: User
	(*UsersBook)(nil),     // 2: UsersBook
	(*UserCreateReq)(nil), // 3: UserCreateReq
	(*UserCreateRes)(nil), // 4: UserCreateRes
	(*UserListReq)(nil),   // 5: UserListReq
	(*UserListRes)(nil),   // 6: UserListRes
	(*UserFindReq)(nil),   // 7: UserFindReq
	(*UserFindRes)(nil),   // 8: UserFindRes
	(*UserDeleteReq)(nil), // 9: UserDeleteReq
	(*UserDeleteRes)(nil), // 10: UserDeleteRes
	(*UserUpdateReq)(nil), // 11: UserUpdateReq
	(*UserUpdateRes)(nil), // 12: UserUpdateRes
}
var file_furkan_service_proto_depIdxs = []int32{
	2,  // 0: ResponseState.data:type_name -> UsersBook
	1,  // 1: UsersBook.people:type_name -> User
	1,  // 2: UserCreateReq.person_item:type_name -> User
	0,  // 3: UserCreateRes.createdUserItems:type_name -> ResponseState
	0,  // 4: UserListRes.ListedUserItems:type_name -> ResponseState
	0,  // 5: UserFindRes.FindedUserItem:type_name -> ResponseState
	0,  // 6: UserDeleteRes.UserDeletedUserItem:type_name -> ResponseState
	1,  // 7: UserUpdateReq.person_item:type_name -> User
	0,  // 8: UserUpdateRes.UserUpdatedItems:type_name -> ResponseState
	3,  // 9: FurkanService.CreateUser:input_type -> UserCreateReq
	5,  // 10: FurkanService.ListUsers:input_type -> UserListReq
	7,  // 11: FurkanService.FindUser:input_type -> UserFindReq
	9,  // 12: FurkanService.DeleteUser:input_type -> UserDeleteReq
	11, // 13: FurkanService.UpdateUser:input_type -> UserUpdateReq
	4,  // 14: FurkanService.CreateUser:output_type -> UserCreateRes
	6,  // 15: FurkanService.ListUsers:output_type -> UserListRes
	8,  // 16: FurkanService.FindUser:output_type -> UserFindRes
	10, // 17: FurkanService.DeleteUser:output_type -> UserDeleteRes
	12, // 18: FurkanService.UpdateUser:output_type -> UserUpdateRes
	14, // [14:19] is the sub-list for method output_type
	9,  // [9:14] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_furkan_service_proto_init() }
func file_furkan_service_proto_init() {
	if File_furkan_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_furkan_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResponseState); i {
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
		file_furkan_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*User); i {
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
		file_furkan_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UsersBook); i {
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
		file_furkan_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserCreateReq); i {
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
		file_furkan_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserCreateRes); i {
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
		file_furkan_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserListReq); i {
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
		file_furkan_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserListRes); i {
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
		file_furkan_service_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserFindReq); i {
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
		file_furkan_service_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserFindRes); i {
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
		file_furkan_service_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserDeleteReq); i {
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
		file_furkan_service_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserDeleteRes); i {
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
		file_furkan_service_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserUpdateReq); i {
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
		file_furkan_service_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserUpdateRes); i {
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
			RawDescriptor: file_furkan_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   13,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_furkan_service_proto_goTypes,
		DependencyIndexes: file_furkan_service_proto_depIdxs,
		MessageInfos:      file_furkan_service_proto_msgTypes,
	}.Build()
	File_furkan_service_proto = out.File
	file_furkan_service_proto_rawDesc = nil
	file_furkan_service_proto_goTypes = nil
	file_furkan_service_proto_depIdxs = nil
}
