package repository

import (
	"context"
	"github.com/gxrlxv/musique/artist_service/internal/domain"
	"github.com/gxrlxv/musique/artist_service/pkg/client/postgresql"
	"github.com/sirupsen/logrus"
)

type AlbumRepository interface {
	CreateAlbum(ctx context.Context, album domain.Album) (string, error)
	DeleteAlbum(ctx context.Context, albumId string) error
	GetAlbumByID(ctx context.Context, albumId string) (domain.Album, error)
	UpdateAlbum(ctx context.Context, albumDTO domain.UpdateAlbumDTO) error
}

type TrackRepository interface {
	SaveTrack(ctx context.Context, track domain.Track) error
	DeleteTrack(ctx context.Context, albumId, trackId string) error
	DeleteTracks(ctx context.Context, albumId string) error
}

type GenreRepository interface {
	GetByTitle(ctx context.Context, title string) (int, error)
}

type ArtistRepository interface {
	GetByName(ctx context.Context, name string) (string, error)
}

type Repository struct {
	AlbumRepository
	TrackRepository
	GenreRepository
	ArtistRepository
}

func NewRepositiry(client postgresql.Client, log *logrus.Logger) *Repository {
	return &Repository{
		AlbumRepository:  NewAlbumRepository(client, log),
		TrackRepository:  NewTrackRepository(client, log),
		GenreRepository:  NewGenreRepository(client, log),
		ArtistRepository: NewArtistRepository(client, log),
	}
}
