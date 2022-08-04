// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.21.1
// source: artist/v1/artist.proto

package v1

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type CreateAlbumRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AlbumTitle   string `protobuf:"bytes,1,opt,name=album_title,json=albumTitle,proto3" json:"album_title,omitempty"`
	ReleaseYear  string `protobuf:"bytes,2,opt,name=release_year,json=releaseYear,proto3" json:"release_year,omitempty"`
	ArtistName   string `protobuf:"bytes,3,opt,name=artist_name,json=artistName,proto3" json:"artist_name,omitempty"`
	GenreTitle   string `protobuf:"bytes,4,opt,name=genre_title,json=genreTitle,proto3" json:"genre_title,omitempty"`
	Milliseconds int64  `protobuf:"varint,5,opt,name=milliseconds,proto3" json:"milliseconds,omitempty"`
	Bytes        int64  `protobuf:"varint,6,opt,name=bytes,proto3" json:"bytes,omitempty"`
}

func (x *CreateAlbumRequest) Reset() {
	*x = CreateAlbumRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_artist_v1_artist_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateAlbumRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateAlbumRequest) ProtoMessage() {}

func (x *CreateAlbumRequest) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use CreateAlbumRequest.ProtoReflect.Descriptor instead.
func (*CreateAlbumRequest) Descriptor() ([]byte, []int) {
	return file_artist_v1_artist_proto_rawDescGZIP(), []int{0}
}

func (x *CreateAlbumRequest) GetAlbumTitle() string {
	if x != nil {
		return x.AlbumTitle
	}
	return ""
}

func (x *CreateAlbumRequest) GetReleaseYear() string {
	if x != nil {
		return x.ReleaseYear
	}
	return ""
}

func (x *CreateAlbumRequest) GetArtistName() string {
	if x != nil {
		return x.ArtistName
	}
	return ""
}

func (x *CreateAlbumRequest) GetGenreTitle() string {
	if x != nil {
		return x.GenreTitle
	}
	return ""
}

func (x *CreateAlbumRequest) GetMilliseconds() int64 {
	if x != nil {
		return x.Milliseconds
	}
	return 0
}

func (x *CreateAlbumRequest) GetBytes() int64 {
	if x != nil {
		return x.Bytes
	}
	return 0
}

type CreateAlbumReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TrackId string `protobuf:"bytes,1,opt,name=track_id,json=trackId,proto3" json:"track_id,omitempty"`
}

func (x *CreateAlbumReply) Reset() {
	*x = CreateAlbumReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_artist_v1_artist_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateAlbumReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateAlbumReply) ProtoMessage() {}

func (x *CreateAlbumReply) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use CreateAlbumReply.ProtoReflect.Descriptor instead.
func (*CreateAlbumReply) Descriptor() ([]byte, []int) {
	return file_artist_v1_artist_proto_rawDescGZIP(), []int{1}
}

func (x *CreateAlbumReply) GetTrackId() string {
	if x != nil {
		return x.TrackId
	}
	return ""
}

type DeleteAlbumRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AlbumId string `protobuf:"bytes,1,opt,name=album_id,json=albumId,proto3" json:"album_id,omitempty"`
}

func (x *DeleteAlbumRequest) Reset() {
	*x = DeleteAlbumRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_artist_v1_artist_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteAlbumRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteAlbumRequest) ProtoMessage() {}

func (x *DeleteAlbumRequest) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use DeleteAlbumRequest.ProtoReflect.Descriptor instead.
func (*DeleteAlbumRequest) Descriptor() ([]byte, []int) {
	return file_artist_v1_artist_proto_rawDescGZIP(), []int{2}
}

func (x *DeleteAlbumRequest) GetAlbumId() string {
	if x != nil {
		return x.AlbumId
	}
	return ""
}

type DeleteAlbumReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AlbumId string `protobuf:"bytes,1,opt,name=album_id,json=albumId,proto3" json:"album_id,omitempty"`
}

func (x *DeleteAlbumReply) Reset() {
	*x = DeleteAlbumReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_artist_v1_artist_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteAlbumReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteAlbumReply) ProtoMessage() {}

func (x *DeleteAlbumReply) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use DeleteAlbumReply.ProtoReflect.Descriptor instead.
func (*DeleteAlbumReply) Descriptor() ([]byte, []int) {
	return file_artist_v1_artist_proto_rawDescGZIP(), []int{3}
}

func (x *DeleteAlbumReply) GetAlbumId() string {
	if x != nil {
		return x.AlbumId
	}
	return ""
}

var File_artist_v1_artist_proto protoreflect.FileDescriptor

var file_artist_v1_artist_proto_rawDesc = []byte{
	0x0a, 0x16, 0x61, 0x72, 0x74, 0x69, 0x73, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x72, 0x74, 0x69,
	0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x72,
	0x74, 0x69, 0x73, 0x74, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd4, 0x01, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x41, 0x6c, 0x62, 0x75, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b,
	0x61, 0x6c, 0x62, 0x75, 0x6d, 0x5f, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x61, 0x6c, 0x62, 0x75, 0x6d, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x21, 0x0a,
	0x0c, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x5f, 0x79, 0x65, 0x61, 0x72, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x59, 0x65, 0x61, 0x72,
	0x12, 0x1f, 0x0a, 0x0b, 0x61, 0x72, 0x74, 0x69, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x61, 0x72, 0x74, 0x69, 0x73, 0x74, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x67, 0x65, 0x6e, 0x72, 0x65, 0x5f, 0x74, 0x69, 0x74, 0x6c, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x67, 0x65, 0x6e, 0x72, 0x65, 0x54, 0x69, 0x74,
	0x6c, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x6d, 0x69, 0x6c, 0x6c, 0x69, 0x73, 0x65, 0x63, 0x6f, 0x6e,
	0x64, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x6d, 0x69, 0x6c, 0x6c, 0x69, 0x73,
	0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x62, 0x79, 0x74, 0x65, 0x73, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x62, 0x79, 0x74, 0x65, 0x73, 0x22, 0x2d, 0x0a, 0x10,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x6c, 0x62, 0x75, 0x6d, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x12, 0x19, 0x0a, 0x08, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x49, 0x64, 0x22, 0x2f, 0x0a, 0x12, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x6c, 0x62, 0x75, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x19, 0x0a, 0x08, 0x61, 0x6c, 0x62, 0x75, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x6c, 0x62, 0x75, 0x6d, 0x49, 0x64, 0x22, 0x2d, 0x0a, 0x10,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x6c, 0x62, 0x75, 0x6d, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x12, 0x19, 0x0a, 0x08, 0x61, 0x6c, 0x62, 0x75, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x61, 0x6c, 0x62, 0x75, 0x6d, 0x49, 0x64, 0x32, 0xec, 0x01, 0x0a, 0x06,
	0x41, 0x72, 0x74, 0x69, 0x73, 0x74, 0x12, 0x6c, 0x0a, 0x0b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x41, 0x6c, 0x62, 0x75, 0x6d, 0x12, 0x21, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x72, 0x74, 0x69,
	0x73, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x6c, 0x62, 0x75,
	0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61,
	0x72, 0x74, 0x69, 0x73, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41,
	0x6c, 0x62, 0x75, 0x6d, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x19, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x13, 0x22, 0x0e, 0x2f, 0x61, 0x72, 0x74, 0x69, 0x73, 0x74, 0x2f, 0x61, 0x6c, 0x62, 0x75, 0x6d,
	0x73, 0x3a, 0x01, 0x2a, 0x12, 0x74, 0x0a, 0x0b, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x6c,
	0x62, 0x75, 0x6d, 0x12, 0x21, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x72, 0x74, 0x69, 0x73, 0x74,
	0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x6c, 0x62, 0x75, 0x6d, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x72, 0x74,
	0x69, 0x73, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x6c, 0x62,
	0x75, 0x6d, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x21, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1b, 0x2a,
	0x19, 0x2f, 0x61, 0x72, 0x74, 0x69, 0x73, 0x74, 0x2f, 0x61, 0x6c, 0x62, 0x75, 0x6d, 0x73, 0x2f,
	0x7b, 0x61, 0x6c, 0x62, 0x75, 0x6d, 0x5f, 0x69, 0x64, 0x7d, 0x42, 0x38, 0x5a, 0x36, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x78, 0x72, 0x6c, 0x78, 0x76, 0x2f,
	0x6d, 0x75, 0x73, 0x69, 0x71, 0x75, 0x65, 0x2f, 0x61, 0x72, 0x74, 0x69, 0x73, 0x74, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x72, 0x74, 0x69, 0x73,
	0x74, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
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

var file_artist_v1_artist_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_artist_v1_artist_proto_goTypes = []interface{}{
	(*CreateAlbumRequest)(nil), // 0: api.artist.v1.CreateAlbumRequest
	(*CreateAlbumReply)(nil),   // 1: api.artist.v1.CreateAlbumReply
	(*DeleteAlbumRequest)(nil), // 2: api.artist.v1.DeleteAlbumRequest
	(*DeleteAlbumReply)(nil),   // 3: api.artist.v1.DeleteAlbumReply
}
var file_artist_v1_artist_proto_depIdxs = []int32{
	0, // 0: api.artist.v1.Artist.CreateAlbum:input_type -> api.artist.v1.CreateAlbumRequest
	2, // 1: api.artist.v1.Artist.DeleteAlbum:input_type -> api.artist.v1.DeleteAlbumRequest
	1, // 2: api.artist.v1.Artist.CreateAlbum:output_type -> api.artist.v1.CreateAlbumReply
	3, // 3: api.artist.v1.Artist.DeleteAlbum:output_type -> api.artist.v1.DeleteAlbumReply
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
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
		file_artist_v1_artist_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
		file_artist_v1_artist_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
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
		file_artist_v1_artist_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
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
			NumMessages:   4,
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
