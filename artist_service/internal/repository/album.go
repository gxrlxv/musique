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
    			(title, artist_id, release_year, number_tracks)
			VALUES
    			($1,$2,$3,$4)
			RETURNING id`

	var albumId string

	err := ar.client.QueryRow(ctx, q, album.Title, album.ArtistId, album.ReleaseYear, album.NumberTracks).Scan(&albumId)
	if err != nil {
		ar.log.Error(err)
		return "", err
	}

	return albumId, err
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

	return err
}
func (ar *albumRepository) GetAlbumByID(ctx context.Context, albumId string) (domain.Album, error) {
	q := `
			SELECT id, title, artist_id, number_tracks, release_year
			FROM public.album 
			WHERE id = $1`

	var album domain.Album

	err := ar.client.QueryRow(ctx, q, albumId).Scan(&album.Id, &album.Title, &album.ArtistId, &album.NumberTracks, &album.ReleaseYear)
	if err != nil {
		ar.log.Error(err)
		return domain.Album{}, err
	}

	return album, err
}

func (ar *albumRepository) UpdateAlbum(ctx context.Context, albumDTO domain.UpdateAlbumDTO) error {
	q := `
			UPDATE public.album 
			SET title = $1, number_tracks = $2 
			WHERE id = $3`

	_, err := ar.client.Exec(ctx, q, albumDTO.Title, albumDTO.NumberTracks, albumDTO.Id)
	if err != nil {
		ar.log.Error(err)
		return err
	}

	return err
}
