// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.21.1
// source: music/v1/music.proto

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
		mi := &file_music_v1_music_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Track) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Track) ProtoMessage() {}

func (x *Track) ProtoReflect() protoreflect.Message {
	mi := &file_music_v1_music_proto_msgTypes[0]
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
	return file_music_v1_music_proto_rawDescGZIP(), []int{0}
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

type AddTrackRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PlaylistId string `protobuf:"bytes,1,opt,name=playlist_id,json=playlistId,proto3" json:"playlist_id,omitempty"`
	TrackId    string `protobuf:"bytes,2,opt,name=track_id,json=trackId,proto3" json:"track_id,omitempty"`
}

func (x *AddTrackRequest) Reset() {
	*x = AddTrackRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_music_v1_music_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddTrackRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddTrackRequest) ProtoMessage() {}

func (x *AddTrackRequest) ProtoReflect() protoreflect.Message {
	mi := &file_music_v1_music_proto_msgTypes[1]
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
	return file_music_v1_music_proto_rawDescGZIP(), []int{1}
}

func (x *AddTrackRequest) GetPlaylistId() string {
	if x != nil {
		return x.PlaylistId
	}
	return ""
}

func (x *AddTrackRequest) GetTrackId() string {
	if x != nil {
		return x.TrackId
	}
	return ""
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
		mi := &file_music_v1_music_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddTrackReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddTrackReply) ProtoMessage() {}

func (x *AddTrackReply) ProtoReflect() protoreflect.Message {
	mi := &file_music_v1_music_proto_msgTypes[2]
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
	return file_music_v1_music_proto_rawDescGZIP(), []int{2}
}

func (x *AddTrackReply) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type AddAlbumRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PlaylistId string `protobuf:"bytes,1,opt,name=playlist_id,json=playlistId,proto3" json:"playlist_id,omitempty"`
	AlbumId    string `protobuf:"bytes,2,opt,name=album_id,json=albumId,proto3" json:"album_id,omitempty"`
}

func (x *AddAlbumRequest) Reset() {
	*x = AddAlbumRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_music_v1_music_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddAlbumRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddAlbumRequest) ProtoMessage() {}

func (x *AddAlbumRequest) ProtoReflect() protoreflect.Message {
	mi := &file_music_v1_music_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddAlbumRequest.ProtoReflect.Descriptor instead.
func (*AddAlbumRequest) Descriptor() ([]byte, []int) {
	return file_music_v1_music_proto_rawDescGZIP(), []int{3}
}

func (x *AddAlbumRequest) GetPlaylistId() string {
	if x != nil {
		return x.PlaylistId
	}
	return ""
}

func (x *AddAlbumRequest) GetAlbumId() string {
	if x != nil {
		return x.AlbumId
	}
	return ""
}

type AddAlbumReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *AddAlbumReply) Reset() {
	*x = AddAlbumReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_music_v1_music_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddAlbumReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddAlbumReply) ProtoMessage() {}

func (x *AddAlbumReply) ProtoReflect() protoreflect.Message {
	mi := &file_music_v1_music_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddAlbumReply.ProtoReflect.Descriptor instead.
func (*AddAlbumReply) Descriptor() ([]byte, []int) {
	return file_music_v1_music_proto_rawDescGZIP(), []int{4}
}

func (x *AddAlbumReply) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type DeleteTrackRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PlaylistId string `protobuf:"bytes,1,opt,name=playlist_id,json=playlistId,proto3" json:"playlist_id,omitempty"`
	TrackId    string `protobuf:"bytes,2,opt,name=track_id,json=trackId,proto3" json:"track_id,omitempty"`
}

func (x *DeleteTrackRequest) Reset() {
	*x = DeleteTrackRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_music_v1_music_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteTrackRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteTrackRequest) ProtoMessage() {}

func (x *DeleteTrackRequest) ProtoReflect() protoreflect.Message {
	mi := &file_music_v1_music_proto_msgTypes[5]
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
	return file_music_v1_music_proto_rawDescGZIP(), []int{5}
}

func (x *DeleteTrackRequest) GetPlaylistId() string {
	if x != nil {
		return x.PlaylistId
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
		mi := &file_music_v1_music_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteTrackReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteTrackReply) ProtoMessage() {}

func (x *DeleteTrackReply) ProtoReflect() protoreflect.Message {
	mi := &file_music_v1_music_proto_msgTypes[6]
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
	return file_music_v1_music_proto_rawDescGZIP(), []int{6}
}

func (x *DeleteTrackReply) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type GetTrackRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PlaylistId string `protobuf:"bytes,1,opt,name=playlist_id,json=playlistId,proto3" json:"playlist_id,omitempty"`
	TrackId    string `protobuf:"bytes,2,opt,name=track_id,json=trackId,proto3" json:"track_id,omitempty"`
}

func (x *GetTrackRequest) Reset() {
	*x = GetTrackRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_music_v1_music_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTrackRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTrackRequest) ProtoMessage() {}

func (x *GetTrackRequest) ProtoReflect() protoreflect.Message {
	mi := &file_music_v1_music_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTrackRequest.ProtoReflect.Descriptor instead.
func (*GetTrackRequest) Descriptor() ([]byte, []int) {
	return file_music_v1_music_proto_rawDescGZIP(), []int{7}
}

func (x *GetTrackRequest) GetPlaylistId() string {
	if x != nil {
		return x.PlaylistId
	}
	return ""
}

func (x *GetTrackRequest) GetTrackId() string {
	if x != nil {
		return x.TrackId
	}
	return ""
}

type GetTrackReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Track *Track `protobuf:"bytes,1,opt,name=track,proto3" json:"track,omitempty"`
}

func (x *GetTrackReply) Reset() {
	*x = GetTrackReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_music_v1_music_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTrackReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTrackReply) ProtoMessage() {}

func (x *GetTrackReply) ProtoReflect() protoreflect.Message {
	mi := &file_music_v1_music_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTrackReply.ProtoReflect.Descriptor instead.
func (*GetTrackReply) Descriptor() ([]byte, []int) {
	return file_music_v1_music_proto_rawDescGZIP(), []int{8}
}

func (x *GetTrackReply) GetTrack() *Track {
	if x != nil {
		return x.Track
	}
	return nil
}

type GetTracksRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PlaylistId string `protobuf:"bytes,1,opt,name=playlist_id,json=playlistId,proto3" json:"playlist_id,omitempty"`
}

func (x *GetTracksRequest) Reset() {
	*x = GetTracksRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_music_v1_music_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTracksRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTracksRequest) ProtoMessage() {}

func (x *GetTracksRequest) ProtoReflect() protoreflect.Message {
	mi := &file_music_v1_music_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTracksRequest.ProtoReflect.Descriptor instead.
func (*GetTracksRequest) Descriptor() ([]byte, []int) {
	return file_music_v1_music_proto_rawDescGZIP(), []int{9}
}

func (x *GetTracksRequest) GetPlaylistId() string {
	if x != nil {
		return x.PlaylistId
	}
	return ""
}

type GetTracksReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tracks []*Track `protobuf:"bytes,1,rep,name=tracks,proto3" json:"tracks,omitempty"`
}

func (x *GetTracksReply) Reset() {
	*x = GetTracksReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_music_v1_music_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTracksReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTracksReply) ProtoMessage() {}

func (x *GetTracksReply) ProtoReflect() protoreflect.Message {
	mi := &file_music_v1_music_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTracksReply.ProtoReflect.Descriptor instead.
func (*GetTracksReply) Descriptor() ([]byte, []int) {
	return file_music_v1_music_proto_rawDescGZIP(), []int{10}
}

func (x *GetTracksReply) GetTracks() []*Track {
	if x != nil {
		return x.Tracks
	}
	return nil
}

var File_music_v1_music_proto protoreflect.FileDescriptor

var file_music_v1_music_proto_rawDesc = []byte{
	0x0a, 0x14, 0x6d, 0x75, 0x73, 0x69, 0x63, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x75, 0x73, 0x69, 0x63,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x61, 0x70, 0x69, 0x2e, 0x74, 0x72, 0x61, 0x63,
	0x6b, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x6d, 0x0a, 0x05, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x12, 0x14, 0x0a, 0x05, 0x74,
	0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x67, 0x65, 0x6e, 0x72, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x67, 0x65, 0x6e, 0x72, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x6d, 0x69, 0x6c, 0x6c, 0x69,
	0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x6d,
	0x69, 0x6c, 0x6c, 0x69, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x62,
	0x79, 0x74, 0x65, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x62, 0x79, 0x74, 0x65,
	0x73, 0x22, 0x4d, 0x0a, 0x0f, 0x41, 0x64, 0x64, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x6c, 0x61, 0x79, 0x6c, 0x69, 0x73, 0x74,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x6c, 0x61, 0x79, 0x6c,
	0x69, 0x73, 0x74, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x5f, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x49, 0x64,
	0x22, 0x29, 0x0a, 0x0d, 0x41, 0x64, 0x64, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x22, 0x4d, 0x0a, 0x0f, 0x41,
	0x64, 0x64, 0x41, 0x6c, 0x62, 0x75, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f,
	0x0a, 0x0b, 0x70, 0x6c, 0x61, 0x79, 0x6c, 0x69, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x6c, 0x61, 0x79, 0x6c, 0x69, 0x73, 0x74, 0x49, 0x64, 0x12,
	0x19, 0x0a, 0x08, 0x61, 0x6c, 0x62, 0x75, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x61, 0x6c, 0x62, 0x75, 0x6d, 0x49, 0x64, 0x22, 0x29, 0x0a, 0x0d, 0x41, 0x64,
	0x64, 0x41, 0x6c, 0x62, 0x75, 0x6d, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x73,
	0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x22, 0x50, 0x0a, 0x12, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54,
	0x72, 0x61, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x70,
	0x6c, 0x61, 0x79, 0x6c, 0x69, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x70, 0x6c, 0x61, 0x79, 0x6c, 0x69, 0x73, 0x74, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08,
	0x74, 0x72, 0x61, 0x63, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x74, 0x72, 0x61, 0x63, 0x6b, 0x49, 0x64, 0x22, 0x2c, 0x0a, 0x10, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x73,
	0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x22, 0x4d, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x54, 0x72, 0x61, 0x63,
	0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x6c, 0x61, 0x79,
	0x6c, 0x69, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70,
	0x6c, 0x61, 0x79, 0x6c, 0x69, 0x73, 0x74, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x74, 0x72, 0x61,
	0x63, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x74, 0x72, 0x61,
	0x63, 0x6b, 0x49, 0x64, 0x22, 0x3a, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x54, 0x72, 0x61, 0x63, 0x6b,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x29, 0x0a, 0x05, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x74, 0x72, 0x61, 0x63, 0x6b,
	0x2e, 0x76, 0x31, 0x2e, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x52, 0x05, 0x74, 0x72, 0x61, 0x63, 0x6b,
	0x22, 0x33, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x6c, 0x61, 0x79, 0x6c, 0x69, 0x73, 0x74,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x6c, 0x61, 0x79, 0x6c,
	0x69, 0x73, 0x74, 0x49, 0x64, 0x22, 0x3d, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x54, 0x72, 0x61, 0x63,
	0x6b, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x2b, 0x0a, 0x06, 0x74, 0x72, 0x61, 0x63, 0x6b,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x74, 0x72,
	0x61, 0x63, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x52, 0x06, 0x74, 0x72,
	0x61, 0x63, 0x6b, 0x73, 0x32, 0xd2, 0x04, 0x0a, 0x05, 0x4d, 0x75, 0x73, 0x69, 0x63, 0x12, 0x76,
	0x0a, 0x08, 0x41, 0x64, 0x64, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x12, 0x1d, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x64, 0x64, 0x54, 0x72, 0x61,
	0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x74, 0x72, 0x61, 0x63, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x64, 0x64, 0x54, 0x72, 0x61, 0x63,
	0x6b, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x2e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x28, 0x22, 0x23,
	0x2f, 0x70, 0x6c, 0x61, 0x79, 0x6c, 0x69, 0x73, 0x74, 0x73, 0x2f, 0x7b, 0x70, 0x6c, 0x61, 0x79,
	0x6c, 0x69, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x7d, 0x2f, 0x7b, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x5f,
	0x69, 0x64, 0x7d, 0x3a, 0x01, 0x2a, 0x12, 0x76, 0x0a, 0x08, 0x41, 0x64, 0x64, 0x41, 0x6c, 0x62,
	0x75, 0x6d, 0x12, 0x1d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x2e, 0x76,
	0x31, 0x2e, 0x41, 0x64, 0x64, 0x41, 0x6c, 0x62, 0x75, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x2e, 0x76, 0x31,
	0x2e, 0x41, 0x64, 0x64, 0x41, 0x6c, 0x62, 0x75, 0x6d, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x2e,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x28, 0x22, 0x23, 0x2f, 0x70, 0x6c, 0x61, 0x79, 0x6c, 0x69, 0x73,
	0x74, 0x73, 0x2f, 0x7b, 0x70, 0x6c, 0x61, 0x79, 0x6c, 0x69, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x7d,
	0x2f, 0x7b, 0x61, 0x6c, 0x62, 0x75, 0x6d, 0x5f, 0x69, 0x64, 0x7d, 0x3a, 0x01, 0x2a, 0x12, 0x77,
	0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x20, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x74,
	0x72, 0x61, 0x63, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x72,
	0x61, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x54, 0x72, 0x61, 0x63, 0x6b, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x2b, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x25, 0x2a, 0x23, 0x2f, 0x70, 0x6c, 0x61, 0x79, 0x6c, 0x69, 0x73, 0x74, 0x73, 0x2f, 0x7b,
	0x70, 0x6c, 0x61, 0x79, 0x6c, 0x69, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x7d, 0x2f, 0x7b, 0x74, 0x72,
	0x61, 0x63, 0x6b, 0x5f, 0x69, 0x64, 0x7d, 0x12, 0x73, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x54, 0x72,
	0x61, 0x63, 0x6b, 0x12, 0x1d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x2e,
	0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x2e, 0x76,
	0x31, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22,
	0x2b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x25, 0x12, 0x23, 0x2f, 0x70, 0x6c, 0x61, 0x79, 0x6c, 0x69,
	0x73, 0x74, 0x73, 0x2f, 0x7b, 0x70, 0x6c, 0x61, 0x79, 0x6c, 0x69, 0x73, 0x74, 0x5f, 0x69, 0x64,
	0x7d, 0x2f, 0x7b, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x5f, 0x69, 0x64, 0x7d, 0x12, 0x6b, 0x0a, 0x09,
	0x47, 0x65, 0x74, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x73, 0x12, 0x1e, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x74, 0x72, 0x61, 0x63, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x72, 0x61, 0x63,
	0x6b, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x74, 0x72, 0x61, 0x63, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x72, 0x61, 0x63,
	0x6b, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x20, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1a, 0x12,
	0x18, 0x2f, 0x70, 0x6c, 0x61, 0x79, 0x6c, 0x69, 0x73, 0x74, 0x73, 0x2f, 0x7b, 0x70, 0x6c, 0x61,
	0x79, 0x6c, 0x69, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x7d, 0x42, 0x36, 0x5a, 0x34, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x78, 0x72, 0x6c, 0x78, 0x76, 0x2f, 0x6d,
	0x75, 0x73, 0x69, 0x71, 0x75, 0x65, 0x2f, 0x6d, 0x75, 0x73, 0x69, 0x63, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6d, 0x75, 0x73, 0x69, 0x63, 0x2f, 0x76,
	0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_music_v1_music_proto_rawDescOnce sync.Once
	file_music_v1_music_proto_rawDescData = file_music_v1_music_proto_rawDesc
)

func file_music_v1_music_proto_rawDescGZIP() []byte {
	file_music_v1_music_proto_rawDescOnce.Do(func() {
		file_music_v1_music_proto_rawDescData = protoimpl.X.CompressGZIP(file_music_v1_music_proto_rawDescData)
	})
	return file_music_v1_music_proto_rawDescData
}

var file_music_v1_music_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_music_v1_music_proto_goTypes = []interface{}{
	(*Track)(nil),              // 0: api.track.v1.Track
	(*AddTrackRequest)(nil),    // 1: api.track.v1.AddTrackRequest
	(*AddTrackReply)(nil),      // 2: api.track.v1.AddTrackReply
	(*AddAlbumRequest)(nil),    // 3: api.track.v1.AddAlbumRequest
	(*AddAlbumReply)(nil),      // 4: api.track.v1.AddAlbumReply
	(*DeleteTrackRequest)(nil), // 5: api.track.v1.DeleteTrackRequest
	(*DeleteTrackReply)(nil),   // 6: api.track.v1.DeleteTrackReply
	(*GetTrackRequest)(nil),    // 7: api.track.v1.GetTrackRequest
	(*GetTrackReply)(nil),      // 8: api.track.v1.GetTrackReply
	(*GetTracksRequest)(nil),   // 9: api.track.v1.GetTracksRequest
	(*GetTracksReply)(nil),     // 10: api.track.v1.GetTracksReply
}
var file_music_v1_music_proto_depIdxs = []int32{
	0,  // 0: api.track.v1.GetTrackReply.track:type_name -> api.track.v1.Track
	0,  // 1: api.track.v1.GetTracksReply.tracks:type_name -> api.track.v1.Track
	1,  // 2: api.track.v1.Music.AddTrack:input_type -> api.track.v1.AddTrackRequest
	3,  // 3: api.track.v1.Music.AddAlbum:input_type -> api.track.v1.AddAlbumRequest
	5,  // 4: api.track.v1.Music.Delete:input_type -> api.track.v1.DeleteTrackRequest
	7,  // 5: api.track.v1.Music.GetTrack:input_type -> api.track.v1.GetTrackRequest
	9,  // 6: api.track.v1.Music.GetTracks:input_type -> api.track.v1.GetTracksRequest
	2,  // 7: api.track.v1.Music.AddTrack:output_type -> api.track.v1.AddTrackReply
	4,  // 8: api.track.v1.Music.AddAlbum:output_type -> api.track.v1.AddAlbumReply
	6,  // 9: api.track.v1.Music.Delete:output_type -> api.track.v1.DeleteTrackReply
	8,  // 10: api.track.v1.Music.GetTrack:output_type -> api.track.v1.GetTrackReply
	10, // 11: api.track.v1.Music.GetTracks:output_type -> api.track.v1.GetTracksReply
	7,  // [7:12] is the sub-list for method output_type
	2,  // [2:7] is the sub-list for method input_type
	2,  // [2:2] is the sub-list for extension type_name
	2,  // [2:2] is the sub-list for extension extendee
	0,  // [0:2] is the sub-list for field type_name
}

func init() { file_music_v1_music_proto_init() }
func file_music_v1_music_proto_init() {
	if File_music_v1_music_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_music_v1_music_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_music_v1_music_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
		file_music_v1_music_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
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
		file_music_v1_music_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddAlbumRequest); i {
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
		file_music_v1_music_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddAlbumReply); i {
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
		file_music_v1_music_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
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
		file_music_v1_music_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
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
		file_music_v1_music_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTrackRequest); i {
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
		file_music_v1_music_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTrackReply); i {
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
		file_music_v1_music_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTracksRequest); i {
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
		file_music_v1_music_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTracksReply); i {
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
			RawDescriptor: file_music_v1_music_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_music_v1_music_proto_goTypes,
		DependencyIndexes: file_music_v1_music_proto_depIdxs,
		MessageInfos:      file_music_v1_music_proto_msgTypes,
	}.Build()
	File_music_v1_music_proto = out.File
	file_music_v1_music_proto_rawDesc = nil
	file_music_v1_music_proto_goTypes = nil
	file_music_v1_music_proto_depIdxs = nil
}
