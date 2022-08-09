package service

import (
	"context"
	v1 "github.com/gxrlxv/musique/music_service/api/music/v1"
	"github.com/sirupsen/logrus"
)

type MusicUseCase interface {
	AddTrack(ctx context.Context, trackId string) (*v1.AddTrackReply, error)
	RemoveTrack(ctx context.Context, trackId string) (*v1.RemoveTrackReply, error)
	GetTracks(ctx context.Context) (*v1.GetTracksReply, error)
}

type MusicService struct {
	v1.UnimplementedMusicServer
	uc  MusicUseCase
	log *logrus.Logger
}

func NewMusicService(useCase MusicUseCase, log *logrus.Logger) *MusicService {
	return &MusicService{
		UnimplementedMusicServer: v1.UnimplementedMusicServer{},
		uc:                       useCase,
		log:                      log,
	}
}

func (ts *MusicService) AddTrack(ctx context.Context, in *v1.AddTrackRequest) (*v1.AddTrackReply, error) {
	return &v1.AddTrackReply{}, nil
}
func (ts *MusicService) RemoveTrack(ctx context.Context, in *v1.RemoveTrackRequest) (*v1.RemoveTrackReply, error) {
	return &v1.RemoveTrackReply{}, nil
}
func (ts *MusicService) GetTracks(ctx context.Context, in *v1.GetTracksRequest) (*v1.GetTracksReply, error) {
	return &v1.GetTracksReply{}, nil
}
