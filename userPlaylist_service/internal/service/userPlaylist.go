package service

import (
	"context"
	v1 "github.com/gxrlxv/musique/music_service/api/userPlaylist/v1"
	"github.com/sirupsen/logrus"
)

type TrackUseCase interface {
	AddTrack(ctx context.Context, trackId string) (*v1.AddTrackReply, error)
	RemoveTrack(ctx context.Context, trackId string) (*v1.RemoveTrackReply, error)
	GetTracks(ctx context.Context) (*v1.GetTracksReply, error)
}

type TrackService struct {
	v1.UnimplementedTrackServer
	uc  TrackUseCase
	log *logrus.Logger
}

func NewTrackService(useCase TrackUseCase) *TrackService {
	return &TrackService{
		UnimplementedTrackServer: v1.UnimplementedTrackServer{},
		uc:                       useCase,
	}
}

func (ts *TrackService) AddTrack(ctx context.Context, in *v1.AddTrackRequest) (*v1.AddTrackReply, error) {
	return &v1.AddTrackReply{}, nil
}
func (ts *TrackService) RemoveTrack(ctx context.Context, in *v1.RemoveTrackRequest) (*v1.RemoveTrackReply, error) {
	return &v1.RemoveTrackReply{}, nil
}
func (ts *TrackService) GetTracks(ctx context.Context, in *v1.GetTracksRequest) (*v1.GetTracksReply, error) {
	return &v1.GetTracksReply{}, nil
}
