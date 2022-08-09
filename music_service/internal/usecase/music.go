package usecase

import (
	"context"
	v1 "github.com/gxrlxv/musique/music_service/api/music/v1"
	"github.com/sirupsen/logrus"
)

type MusicRepository interface {
}

type musicUseCase struct {
	musicRepo MusicRepository
	log       *logrus.Logger
}

func NewMusicUseCase(musicRepo MusicRepository, log *logrus.Logger) *musicUseCase {
	return &musicUseCase{
		musicRepo: musicRepo,
		log:       log,
	}
}

func (tu *musicUseCase) AddTrack(ctx context.Context, trackId string) (*v1.AddTrackReply, error) {
	//TODO implement me
	panic("implement me")
}

func (tu *musicUseCase) RemoveTrack(ctx context.Context, trackId string) (*v1.RemoveTrackReply, error) {
	//TODO implement me

	panic("implement me")
}

func (tu *musicUseCase) GetTracks(ctx context.Context) (*v1.GetTracksReply, error) {
	//TODO implement me
	panic("implement me")
}
