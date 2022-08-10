package usecase

import (
	"context"
	"github.com/gxrlxv/musique/music_service/internal/domain"
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

func (m *musicUseCase) AddTrack(ctx context.Context, playlistId, trackId string) (bool, error) {
	//TODO implement me
	panic("implement me")
}

func (m *musicUseCase) AddAlbum(ctx context.Context, playlistId, albumId string) (bool, error) {
	//TODO implement me
	panic("implement me")
}

func (m *musicUseCase) DeleteTrack(ctx context.Context, playlistId, trackId string) (bool, error) {
	//TODO implement me
	panic("implement me")
}

func (m *musicUseCase) GetTrack(ctx context.Context, playlistId, trackId string) (domain.Track, error) {
	//TODO implement me
	panic("implement me")
}

func (m *musicUseCase) GetAllTracks(ctx context.Context, playlistId string) ([]domain.Track, error) {
	//TODO implement me
	panic("implement me")
}
