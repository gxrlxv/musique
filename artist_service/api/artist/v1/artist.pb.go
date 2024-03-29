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

type Track struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title        string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Genre        string `protobuf:"bytes,2,opt,name=genre,proto3" json:"genre,omitempty"`
	Milliseconds int64  `protobuf:"varint,3,opt,name=milliseconds,proto3" json:"milliseconds,omitempty"`
	Bytes        int64  `protobuf:"varint,4,opt,name=bytes,proto3" json:"bytes,omitempty"`
}

func (x *Track) Reset() {
	*x = Track{}
	if protoimpl.UnsafeEnabled {
		mi := &file_artist_v1_artist_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Track) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Track) ProtoMessage() {}

func (x *Track) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use Track.ProtoReflect.Descriptor instead.
func (*Track) Descriptor() ([]byte, []int) {
	return file_artist_v1_artist_proto_rawDescGZIP(), []int{0}
}

func (x *Track) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Track) GetGenre() string {
	if x != nil {
		return x.Genre
	}
	return ""
}

func (x *Track) GetMilliseconds() int64 {
	if x != nil {
		return x.Milliseconds
	}
	return 0
}

func (x *Track) GetBytes() int64 {
	if x != nil {
		return x.Bytes
	}
	return 0
}

type CreateAlbumRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AlbumTitle  string   `protobuf:"bytes,1,opt,name=album_title,json=albumTitle,proto3" json:"album_title,omitempty"`
	ReleaseYear int32    `protobuf:"varint,2,opt,name=release_year,json=releaseYear,proto3" json:"release_year,omitempty"`
	ArtistName  string   `protobuf:"bytes,3,opt,name=artist_name,json=artistName,proto3" json:"artist_name,omitempty"`
	Tracks      []*Track `protobuf:"bytes,4,rep,name=tracks,proto3" json:"tracks,omitempty"`
}

func (x *CreateAlbumRequest) Reset() {
	*x = CreateAlbumRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_artist_v1_artist_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateAlbumRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateAlbumRequest) ProtoMessage() {}

func (x *CreateAlbumRequest) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use CreateAlbumRequest.ProtoReflect.Descriptor instead.
func (*CreateAlbumRequest) Descriptor() ([]byte, []int) {
	return file_artist_v1_artist_proto_rawDescGZIP(), []int{1}
}

func (x *CreateAlbumRequest) GetAlbumTitle() string {
	if x != nil {
		return x.AlbumTitle
	}
	return ""
}

func (x *CreateAlbumRequest) GetReleaseYear() int32 {
	if x != nil {
		return x.ReleaseYear
	}
	return 0
}

func (x *CreateAlbumRequest) GetArtistName() string {
	if x != nil {
		return x.ArtistName
	}
	return ""
}

func (x *CreateAlbumRequest) GetTracks() []*Track {
	if x != nil {
		return x.Tracks
	}
	return nil
}

type CreateAlbumReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AlbumId string `protobuf:"bytes,1,opt,name=album_id,json=albumId,proto3" json:"album_id,omitempty"`
}

func (x *CreateAlbumReply) Reset() {
	*x = CreateAlbumReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_artist_v1_artist_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateAlbumReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateAlbumReply) ProtoMessage() {}

func (x *CreateAlbumReply) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use CreateAlbumReply.ProtoReflect.Descriptor instead.
func (*CreateAlbumReply) Descriptor() ([]byte, []int) {
	return file_artist_v1_artist_proto_rawDescGZIP(), []int{2}
}

func (x *CreateAlbumReply) GetAlbumId() string {
	if x != nil {
		return x.AlbumId
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
		mi := &file_artist_v1_artist_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteAlbumRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteAlbumRequest) ProtoMessage() {}

func (x *DeleteAlbumRequest) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use DeleteAlbumRequest.ProtoReflect.Descriptor instead.
func (*DeleteAlbumRequest) Descriptor() ([]byte, []int) {
	return file_artist_v1_artist_proto_rawDescGZIP(), []int{3}
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

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *DeleteAlbumReply) Reset() {
	*x = DeleteAlbumReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_artist_v1_artist_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteAlbumReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteAlbumReply) ProtoMessage() {}

func (x *DeleteAlbumReply) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use DeleteAlbumReply.ProtoReflect.Descriptor instead.
func (*DeleteAlbumReply) Descriptor() ([]byte, []int) {
	return file_artist_v1_artist_proto_rawDescGZIP(), []int{4}
}

func (x *DeleteAlbumReply) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type DeleteTrackRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AlbumId string `protobuf:"bytes,1,opt,name=album_id,json=albumId,proto3" json:"album_id,omitempty"`
	TrackId string `protobuf:"bytes,2,opt,name=track_id,json=trackId,proto3" json:"track_id,omitempty"`
}

func (x *DeleteTrackRequest) Reset() {
	*x = DeleteTrackRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_artist_v1_artist_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteTrackRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteTrackRequest) ProtoMessage() {}

func (x *DeleteTrackRequest) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use DeleteTrackRequest.ProtoReflect.Descriptor instead.
func (*DeleteTrackRequest) Descriptor() ([]byte, []int) {
	return file_artist_v1_artist_proto_rawDescGZIP(), []int{5}
}

func (x *DeleteTrackRequest) GetAlbumId() string {
	if x != nil {
		return x.AlbumId
	}
	return ""
}

func (x *DeleteTrackRequest) GetTrackId() string {
	if x != nil {
		return x.TrackId
	}
	return ""
}

type DeleteTrackReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *DeleteTrackReply) Reset() {
	*x = DeleteTrackReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_artist_v1_artist_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteTrackReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteTrackReply) ProtoMessage() {}

func (x *DeleteTrackReply) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use DeleteTrackReply.ProtoReflect.Descriptor instead.
func (*DeleteTrackReply) Descriptor() ([]byte, []int) {
	return file_artist_v1_artist_proto_rawDescGZIP(), []int{6}
}

func (x *DeleteTrackReply) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type AddTrackRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AlbumId string `protobuf:"bytes,1,opt,name=album_id,json=albumId,proto3" json:"album_id,omitempty"`
	Track   *Track `protobuf:"bytes,2,opt,name=track,proto3" json:"track,omitempty"`
}

func (x *AddTrackRequest) Reset() {
	*x = AddTrackRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_artist_v1_artist_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddTrackRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddTrackRequest) ProtoMessage() {}

func (x *AddTrackRequest) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use AddTrackRequest.ProtoReflect.Descriptor instead.
func (*AddTrackRequest) Descriptor() ([]byte, []int) {
	return file_artist_v1_artist_proto_rawDescGZIP(), []int{7}
}

func (x *AddTrackRequest) GetAlbumId() string {
	if x != nil {
		return x.AlbumId
	}
	return ""
}

func (x *AddTrackRequest) GetTrack() *Track {
	if x != nil {
		return x.Track
	}
	return nil
}

type AddTrackReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *AddTrackReply) Reset() {
	*x = AddTrackReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_artist_v1_artist_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddTrackReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddTrackReply) ProtoMessage() {}

func (x *AddTrackReply) ProtoReflect() protoreflect.Message {
	mi := &file_artist_v1_artist_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddTrackReply.ProtoReflect.Descriptor instead.
func (*AddTrackReply) Descriptor() ([]byte, []int) {
	return file_artist_v1_artist_proto_rawDescGZIP(), []int{8}
}

func (x *AddTrackReply) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

var File_artist_v1_artist_proto protoreflect.FileDescriptor

var file_artist_v1_artist_proto_rawDesc = []byte{
	0x0a, 0x16, 0x61, 0x72, 0x74, 0x69, 0x73, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x72, 0x74, 0x69,
	0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x72,
	0x74, 0x69, 0x73, 0x74, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x6d, 0x0a, 0x05, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x12, 0x14,
	0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74,
	0x69, 0x74, 0x6c, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x67, 0x65, 0x6e, 0x72, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x67, 0x65, 0x6e, 0x72, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x6d, 0x69,
	0x6c, 0x6c, 0x69, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x0c, 0x6d, 0x69, 0x6c, 0x6c, 0x69, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x12, 0x14,
	0x0a, 0x05, 0x62, 0x79, 0x74, 0x65, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x62,
	0x79, 0x74, 0x65, 0x73, 0x22, 0xa7, 0x01, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41,
	0x6c, 0x62, 0x75, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x61,
	0x6c, 0x62, 0x75, 0x6d, 0x5f, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x61, 0x6c, 0x62, 0x75, 0x6d, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x21, 0x0a, 0x0c,
	0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x5f, 0x79, 0x65, 0x61, 0x72, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x0b, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x59, 0x65, 0x61, 0x72, 0x12,
	0x1f, 0x0a, 0x0b, 0x61, 0x72, 0x74, 0x69, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x61, 0x72, 0x74, 0x69, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x2c, 0x0a, 0x06, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x14, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x72, 0x74, 0x69, 0x73, 0x74, 0x2e, 0x76, 0x31,
	0x2e, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x52, 0x06, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x73, 0x22, 0x2d,
	0x0a, 0x10, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x6c, 0x62, 0x75, 0x6d, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x12, 0x19, 0x0a, 0x08, 0x61, 0x6c, 0x62, 0x75, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x6c, 0x62, 0x75, 0x6d, 0x49, 0x64, 0x22, 0x2f, 0x0a,
	0x12, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x6c, 0x62, 0x75, 0x6d, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x61, 0x6c, 0x62, 0x75, 0x6d, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x6c, 0x62, 0x75, 0x6d, 0x49, 0x64, 0x22, 0x2c,
	0x0a, 0x10, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x6c, 0x62, 0x75, 0x6d, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x22, 0x4a, 0x0a, 0x12,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x61, 0x6c, 0x62, 0x75, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x6c, 0x62, 0x75, 0x6d, 0x49, 0x64, 0x12, 0x19, 0x0a,
	0x08, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x49, 0x64, 0x22, 0x2c, 0x0a, 0x10, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x18, 0x0a, 0x07,
	0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73,
	0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x22, 0x58, 0x0a, 0x0f, 0x41, 0x64, 0x64, 0x54, 0x72, 0x61,
	0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x61, 0x6c, 0x62,
	0x75, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x6c, 0x62,
	0x75, 0x6d, 0x49, 0x64, 0x12, 0x2a, 0x0a, 0x05, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x72, 0x74, 0x69, 0x73, 0x74,
	0x2e, 0x76, 0x31, 0x2e, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x52, 0x05, 0x74, 0x72, 0x61, 0x63, 0x6b,
	0x22, 0x29, 0x0a, 0x0d, 0x41, 0x64, 0x64, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x32, 0xdd, 0x03, 0x0a, 0x06,
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
	0x7b, 0x61, 0x6c, 0x62, 0x75, 0x6d, 0x5f, 0x69, 0x64, 0x7d, 0x12, 0x7f, 0x0a, 0x0b, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x12, 0x21, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x61, 0x72, 0x74, 0x69, 0x73, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x54, 0x72, 0x61, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x61, 0x72, 0x74, 0x69, 0x73, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x2c, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x26, 0x2a, 0x24, 0x2f, 0x61, 0x72, 0x74, 0x69, 0x73, 0x74, 0x2f, 0x61,
	0x6c, 0x62, 0x75, 0x6d, 0x73, 0x2f, 0x7b, 0x61, 0x6c, 0x62, 0x75, 0x6d, 0x5f, 0x69, 0x64, 0x7d,
	0x2f, 0x7b, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x5f, 0x69, 0x64, 0x7d, 0x12, 0x6e, 0x0a, 0x08, 0x41,
	0x64, 0x64, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x12, 0x1e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x72,
	0x74, 0x69, 0x73, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x64, 0x64, 0x54, 0x72, 0x61, 0x63, 0x6b,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x72,
	0x74, 0x69, 0x73, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x64, 0x64, 0x54, 0x72, 0x61, 0x63, 0x6b,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x24, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1e, 0x22, 0x19, 0x2f,
	0x61, 0x72, 0x74, 0x69, 0x73, 0x74, 0x2f, 0x61, 0x6c, 0x62, 0x75, 0x6d, 0x73, 0x2f, 0x7b, 0x61,
	0x6c, 0x62, 0x75, 0x6d, 0x5f, 0x69, 0x64, 0x7d, 0x3a, 0x01, 0x2a, 0x42, 0x38, 0x5a, 0x36, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x78, 0x72, 0x6c, 0x78, 0x76,
	0x2f, 0x6d, 0x75, 0x73, 0x69, 0x71, 0x75, 0x65, 0x2f, 0x61, 0x72, 0x74, 0x69, 0x73, 0x74, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x72, 0x74, 0x69,
	0x73, 0x74, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
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

var file_artist_v1_artist_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_artist_v1_artist_proto_goTypes = []interface{}{
	(*Track)(nil),              // 0: api.artist.v1.Track
	(*CreateAlbumRequest)(nil), // 1: api.artist.v1.CreateAlbumRequest
	(*CreateAlbumReply)(nil),   // 2: api.artist.v1.CreateAlbumReply
	(*DeleteAlbumRequest)(nil), // 3: api.artist.v1.DeleteAlbumRequest
	(*DeleteAlbumReply)(nil),   // 4: api.artist.v1.DeleteAlbumReply
	(*DeleteTrackRequest)(nil), // 5: api.artist.v1.DeleteTrackRequest
	(*DeleteTrackReply)(nil),   // 6: api.artist.v1.DeleteTrackReply
	(*AddTrackRequest)(nil),    // 7: api.artist.v1.AddTrackRequest
	(*AddTrackReply)(nil),      // 8: api.artist.v1.AddTrackReply
}
var file_artist_v1_artist_proto_depIdxs = []int32{
	0, // 0: api.artist.v1.CreateAlbumRequest.tracks:type_name -> api.artist.v1.Track
	0, // 1: api.artist.v1.AddTrackRequest.track:type_name -> api.artist.v1.Track
	1, // 2: api.artist.v1.Artist.CreateAlbum:input_type -> api.artist.v1.CreateAlbumRequest
	3, // 3: api.artist.v1.Artist.DeleteAlbum:input_type -> api.artist.v1.DeleteAlbumRequest
	5, // 4: api.artist.v1.Artist.DeleteTrack:input_type -> api.artist.v1.DeleteTrackRequest
	7, // 5: api.artist.v1.Artist.AddTrack:input_type -> api.artist.v1.AddTrackRequest
	2, // 6: api.artist.v1.Artist.CreateAlbum:output_type -> api.artist.v1.CreateAlbumReply
	4, // 7: api.artist.v1.Artist.DeleteAlbum:output_type -> api.artist.v1.DeleteAlbumReply
	6, // 8: api.artist.v1.Artist.DeleteTrack:output_type -> api.artist.v1.DeleteTrackReply
	8, // 9: api.artist.v1.Artist.AddTrack:output_type -> api.artist.v1.AddTrackReply
	6, // [6:10] is the sub-list for method output_type
	2, // [2:6] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_artist_v1_artist_proto_init() }
func file_artist_v1_artist_proto_init() {
	if File_artist_v1_artist_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_artist_v1_artist_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Track); i {
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
		file_artist_v1_artist_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
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
		file_artist_v1_artist_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
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
		file_artist_v1_artist_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
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
		file_artist_v1_artist_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteTrackRequest); i {
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
			switch v := v.(*DeleteTrackReply); i {
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
			switch v := v.(*AddTrackRequest); i {
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
		file_artist_v1_artist_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddTrackReply); i {
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
			NumMessages:   9,
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
