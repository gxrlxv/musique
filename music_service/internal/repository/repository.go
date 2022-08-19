package repository

import (
	"context"
	"github.com/gxrlxv/musique/music_service/internal/domain"
	"github.com/gxrlxv/musique/music_service/pkg/client/postgresql"
	"github.com/sirupsen/logrus"
)

type PlaylistRepository interface {
	SaveTrack(ctx context.Context, playlistId, trackId string) error
	DeleteTrack(ctx context.Context, playlistId, trackId string) error
	GetTrack(ctx context.Context, playlistId, trackId string) error
	GetAllTracks(ctx context.Context, playlistId string) ([]string, error)
	GetPlaylist(ctx context.Context, playlistId string) (domain.Playlist, error)
	UpdatePlaylist(ctx context.Context, playlist domain.Playlist) error
}

type TrackRepository interface {
	GetTrack(ctx context.Context, trackId string) (domain.Track, error)
	GetTracksByAlbumId(ctx context.Context, albumId string) ([]domain.Track, error)
}

type AlbumRepository interface {
	GetAlbum(ctx context.Context, albumId string) (domain.Album, error)
}

type Repository struct {
	PlaylistRepository
	TrackRepository
	AlbumRepository
}

func NewRepository(client postgresql.Client, log *logrus.Logger) *Repository {
	return &Repository{
		PlaylistRepository: NewPlaylistRepository(client, log),
		TrackRepository:    NewTrackRepository(client, log),
		AlbumRepository:    NewAlbumRepository(client, log),
	}
}
