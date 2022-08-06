package repository

import (
	"context"
	"github.com/gxrlxv/musique/artist_service/internal/domain"
	"github.com/gxrlxv/musique/artist_service/pkg/client/postgresql"
	"github.com/sirupsen/logrus"
)

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

func (ar *albumRepository) CreateAlbum(ctx context.Context, album *domain.Album) (string, error) {
	q := `
			INSERT INTO public.album
    			(title, artist_id, release_year)
			VALUES
    			($1,$2,$3)
			RETURNING id`

	var albumId string

	err := ar.client.QueryRow(ctx, q, album.Title, album.ArtistId, album.ReleaseYear).Scan(&albumId)
	if err != nil {
		ar.log.Error(err)
		return "", err
	}

	return albumId, nil
}

func (ar *albumRepository) DeleteAlbum(ctx context.Context, albumId string) error {
	q := `
			DELETE FROM public.album
    		WHERE id = $1`

	_, err := ar.client.Exec(ctx, q, albumId)
	if err != nil {
		ar.log.Error(err)
		return err
	}

	return nil
}
