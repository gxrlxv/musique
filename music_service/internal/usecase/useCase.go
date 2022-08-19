package usecase

import (
	"context"
	"github.com/gxrlxv/musique/music_service/internal/domain"
	"github.com/gxrlxv/musique/music_service/internal/repository"
	"github.com/sirupsen/logrus"
)

type MusicUseCase interface {
	AddTrack(ctx context.Context, playlistId, trackId string) (bool, error)
	AddAlbum(ctx context.Context, playlistId, albumId string) (bool, error)
	DeleteTrack(ctx context.Context, playlistId, trackId string) (bool, error)
	GetTrack(ctx context.Context, playlistId, trackId string) (domain.Track, error)
	GetAllTracks(ctx context.Context, playlistId string) ([]domain.Track, error)
}

type UseCase struct {
	MusicUseCase
}

func NewUseCase(repo *repository.Repository, log *logrus.Logger) *UseCase {
	return &UseCase{
		MusicUseCase: NewMusicUseCase(repo.PlaylistRepository, repo.TrackRepository, repo.AlbumRepository, log),
	}
}
