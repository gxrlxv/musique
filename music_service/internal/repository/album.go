package repository

import (
	"context"
	"github.com/gxrlxv/musique/music_service/internal/domain"
	"github.com/gxrlxv/musique/music_service/pkg/client/postgresql"
	"github.com/jackc/pgx/v4"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var ErrAlbumNotFound = status.Error(codes.NotFound, "album with given id not found")

type albumRepository struct {
	client postgresql.Client
	log    *logrus.Logger
}

func NewAlbumRepository(client postgresql.Client, log *logrus.Logger) *albumRepository {
	return &albumRepository{
		client: client,
		log:    log,
	}
}

func (ar *albumRepository) GetAlbum(ctx context.Context, albumId string) (domain.Album, error) {
	q := `
			SELECT id, title, number_tracks, release_year
			FROM public.album
			WHERE id = $1`

	var album domain.Album

	err := ar.client.QueryRow(ctx, q, albumId).Scan(&album.Id, &album.Title, &album.NumberTracks, &album.ReleaseYear)
	if err != nil {
		ar.log.Error(err)
		if err == pgx.ErrNoRows {
			return domain.Album{}, ErrAlbumNotFound
		}
		return domain.Album{}, err
	}

	return album, err
}
