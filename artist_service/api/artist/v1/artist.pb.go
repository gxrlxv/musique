// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.21.1
// source: artist/v1/artist.proto

package v1

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

type CreateSingleRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CreateSingleRequest) Reset() {
	*x = CreateSingleRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_artist_v1_artist_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateSingleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateSingleRequest) ProtoMessage() {}

func (x *CreateSingleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_artist_v1_artist_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateSingleRequest.ProtoReflect.Descriptor instead.
func (*CreateSingleRequest) Descriptor() ([]byte, []int) {
	return file_artist_v1_artist_proto_rawDescGZIP(), []int{0}
}

type CreateSingleReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CreateSingleReply) Reset() {
	*x = CreateSingleReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_artist_v1_artist_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateSingleReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateSingleReply) ProtoMessage() {}

func (x *CreateSingleReply) ProtoReflect() protoreflect.Message {
	mi := &file_artist_v1_artist_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateSingleReply.ProtoReflect.Descriptor instead.
func (*CreateSingleReply) Descriptor() ([]byte, []int) {
	return file_artist_v1_artist_proto_rawDescGZIP(), []int{1}
}

type CreateAlbumRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CreateAlbumRequest) Reset() {
	*x = CreateAlbumRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_artist_v1_artist_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateAlbumRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateAlbumRequest) ProtoMessage() {}

func (x *CreateAlbumRequest) ProtoReflect() protoreflect.Message {
	mi := &file_artist_v1_artist_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateAlbumRequest.ProtoReflect.Descriptor instead.
func (*CreateAlbumRequest) Descriptor() ([]byte, []int) {
	return file_artist_v1_artist_proto_rawDescGZIP(), []int{2}
}

type CreateAlbumReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CreateAlbumReply) Reset() {
	*x = CreateAlbumReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_artist_v1_artist_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateAlbumReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateAlbumReply) ProtoMessage() {}

func (x *CreateAlbumReply) ProtoReflect() protoreflect.Message {
	mi := &file_artist_v1_artist_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateAlbumReply.ProtoReflect.Descriptor instead.
func (*CreateAlbumReply) Descriptor() ([]byte, []int) {
	return file_artist_v1_artist_proto_rawDescGZIP(), []int{3}
}

type DeleteSingleRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteSingleRequest) Reset() {
	*x = DeleteSingleRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_artist_v1_artist_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteSingleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteSingleRequest) ProtoMessage() {}

func (x *DeleteSingleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_artist_v1_artist_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteSingleRequest.ProtoReflect.Descriptor instead.
func (*DeleteSingleRequest) Descriptor() ([]byte, []int) {
	return file_artist_v1_artist_proto_rawDescGZIP(), []int{4}
}

type DeleteSingleReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteSingleReply) Reset() {
	*x = DeleteSingleReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_artist_v1_artist_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteSingleReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteSingleReply) ProtoMessage() {}

func (x *DeleteSingleReply) ProtoReflect() protoreflect.Message {
	mi := &file_artist_v1_artist_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteSingleReply.ProtoReflect.Descriptor instead.
func (*DeleteSingleReply) Descriptor() ([]byte, []int) {
	return file_artist_v1_artist_proto_rawDescGZIP(), []int{5}
}

type DeleteAlbumRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteAlbumRequest) Reset() {
	*x = DeleteAlbumRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_artist_v1_artist_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteAlbumRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteAlbumRequest) ProtoMessage() {}

func (x *DeleteAlbumRequest) ProtoReflect() protoreflect.Message {
	mi := &file_artist_v1_artist_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteAlbumRequest.ProtoReflect.Descriptor instead.
func (*DeleteAlbumRequest) Descriptor() ([]byte, []int) {
	return file_artist_v1_artist_proto_rawDescGZIP(), []int{6}
}

type DeleteAlbumReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteAlbumReply) Reset() {
	*x = DeleteAlbumReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_artist_v1_artist_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteAlbumReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteAlbumReply) ProtoMessage() {}

func (x *DeleteAlbumReply) ProtoReflect() protoreflect.Message {
	mi := &file_artist_v1_artist_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteAlbumReply.ProtoReflect.Descriptor instead.
func (*DeleteAlbumReply) Descriptor() ([]byte, []int) {
	return file_artist_v1_artist_proto_rawDescGZIP(), []int{7}
}

var File_artist_v1_artist_proto protoreflect.FileDescriptor

var file_artist_v1_artist_proto_rawDesc = []byte{
	0x0a, 0x16, 0x61, 0x72, 0x74, 0x69, 0x73, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x72, 0x74, 0x69,
	0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x72,
	0x74, 0x69, 0x73, 0x74, 0x2e, 0x76, 0x31, 0x22, 0x15, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x53, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x13,
	0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x22, 0x14, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x6c, 0x62,
	0x75, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x12, 0x0a, 0x10, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x41, 0x6c, 0x62, 0x75, 0x6d, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x15, 0x0a,
	0x13, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x22, 0x13, 0x0a, 0x11, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x69,
	0x6e, 0x67, 0x6c, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x14, 0x0a, 0x12, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x41, 0x6c, 0x62, 0x75, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22,
	0x12, 0x0a, 0x10, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x6c, 0x62, 0x75, 0x6d, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x32, 0xe2, 0x02, 0x0a, 0x06, 0x41, 0x72, 0x74, 0x69, 0x73, 0x74, 0x12, 0x56,
	0x0a, 0x0c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x12, 0x22,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x72, 0x74, 0x69, 0x73, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x20, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x72, 0x74, 0x69, 0x73, 0x74, 0x2e,
	0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x53, 0x0a, 0x0b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x41, 0x6c, 0x62, 0x75, 0x6d, 0x12, 0x21, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x72, 0x74, 0x69,
	0x73, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x6c, 0x62, 0x75,
	0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61,
	0x72, 0x74, 0x69, 0x73, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41,
	0x6c, 0x62, 0x75, 0x6d, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x56, 0x0a, 0x0c, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x12, 0x22, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x61, 0x72, 0x74, 0x69, 0x73, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x53, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x20, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x72, 0x74, 0x69, 0x73, 0x74, 0x2e, 0x76, 0x31, 0x2e,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x22, 0x00, 0x12, 0x53, 0x0a, 0x0b, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x6c, 0x62,
	0x75, 0x6d, 0x12, 0x21, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x72, 0x74, 0x69, 0x73, 0x74, 0x2e,
	0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x6c, 0x62, 0x75, 0x6d, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x72, 0x74, 0x69,
	0x73, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x6c, 0x62, 0x75,
	0x6d, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x42, 0x38, 0x5a, 0x36, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x78, 0x72, 0x6c, 0x78, 0x76, 0x2f, 0x6d, 0x75,
	0x73, 0x69, 0x71, 0x75, 0x65, 0x2f, 0x61, 0x72, 0x74, 0x69, 0x73, 0x74, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x72, 0x74, 0x69, 0x73, 0x74, 0x2f,
	0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_artist_v1_artist_proto_rawDescOnce sync.Once
	file_artist_v1_artist_proto_rawDescData = file_artist_v1_artist_proto_rawDesc
)

func file_artist_v1_artist_proto_rawDescGZIP() []byte {
	file_artist_v1_artist_proto_rawDescOnce.Do(func() {
		file_artist_v1_artist_proto_rawDescData = protoimpl.X.CompressGZIP(file_artist_v1_artist_proto_rawDescData)
	})
	return file_artist_v1_artist_proto_rawDescData
}

var file_artist_v1_artist_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_artist_v1_artist_proto_goTypes = []interface{}{
	(*CreateSingleRequest)(nil), // 0: api.artist.v1.CreateSingleRequest
	(*CreateSingleReply)(nil),   // 1: api.artist.v1.CreateSingleReply
	(*CreateAlbumRequest)(nil),  // 2: api.artist.v1.CreateAlbumRequest
	(*CreateAlbumReply)(nil),    // 3: api.artist.v1.CreateAlbumReply
	(*DeleteSingleRequest)(nil), // 4: api.artist.v1.DeleteSingleRequest
	(*DeleteSingleReply)(nil),   // 5: api.artist.v1.DeleteSingleReply
	(*DeleteAlbumRequest)(nil),  // 6: api.artist.v1.DeleteAlbumRequest
	(*DeleteAlbumReply)(nil),    // 7: api.artist.v1.DeleteAlbumReply
}
var file_artist_v1_artist_proto_depIdxs = []int32{
	0, // 0: api.artist.v1.Artist.CreateSingle:input_type -> api.artist.v1.CreateSingleRequest
	2, // 1: api.artist.v1.Artist.CreateAlbum:input_type -> api.artist.v1.CreateAlbumRequest
	4, // 2: api.artist.v1.Artist.DeleteSingle:input_type -> api.artist.v1.DeleteSingleRequest
	6, // 3: api.artist.v1.Artist.DeleteAlbum:input_type -> api.artist.v1.DeleteAlbumRequest
	1, // 4: api.artist.v1.Artist.CreateSingle:output_type -> api.artist.v1.CreateSingleReply
	3, // 5: api.artist.v1.Artist.CreateAlbum:output_type -> api.artist.v1.CreateAlbumReply
	5, // 6: api.artist.v1.Artist.DeleteSingle:output_type -> api.artist.v1.DeleteSingleReply
	7, // 7: api.artist.v1.Artist.DeleteAlbum:output_type -> api.artist.v1.DeleteAlbumReply
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_artist_v1_artist_proto_init() }
func file_artist_v1_artist_proto_init() {
	if File_artist_v1_artist_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_artist_v1_artist_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateSingleRequest); i {
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
		file_artist_v1_artist_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateSingleReply); i {
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
		file_artist_v1_artist_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateAlbumRequest); i {
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
		file_artist_v1_artist_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateAlbumReply); i {
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
		file_artist_v1_artist_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteSingleRequest); i {
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
		file_artist_v1_artist_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteSingleReply); i {
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
		file_artist_v1_artist_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteAlbumRequest); i {
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
		file_artist_v1_artist_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteAlbumReply); i {
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
			RawDescriptor: file_artist_v1_artist_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_artist_v1_artist_proto_goTypes,
		DependencyIndexes: file_artist_v1_artist_proto_depIdxs,
		MessageInfos:      file_artist_v1_artist_proto_msgTypes,
	}.Build()
	File_artist_v1_artist_proto = out.File
	file_artist_v1_artist_proto_rawDesc = nil
	file_artist_v1_artist_proto_goTypes = nil
	file_artist_v1_artist_proto_depIdxs = nil
}
