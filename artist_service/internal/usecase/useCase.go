package usecase

import (
	"context"
	"github.com/gxrlxv/musique/artist_service/internal/domain"
	"github.com/gxrlxv/musique/artist_service/internal/repository"
	"github.com/sirupsen/logrus"
)

type ArtistUseCase interface {
	CreateAlbum(ctx context.Context, album domain.CreateAlbumDTO) (string, error)
	DeleteAlbum(ctx context.Context, albumId string) (bool, error)
	DeleteTrack(ctx context.Context, albumId, trackId string) (bool, error)
	AddTrack(ctx context.Context, albumId string, track domain.CreateTrackDTO) (bool, error)
	CheckPermission(ctx context.Context) error
}

type UseCase struct {
	ArtistUseCase
}

func NewUseCase(repo *repository.Repository, log *logrus.Logger) *UseCase {
	return &UseCase{
		ArtistUseCase: NewArtistUseCase(repo.AlbumRepository, repo.TrackRepository, repo.GenreRepository,
			repo.ArtistRepository, log),
	}
}
