package repository

import (
	"context"
	"github.com/gxrlxv/musique/artist_service/pkg/client/postgresql"
	"github.com/sirupsen/logrus"
)

type artistRepository struct {
	client postgresql.Client
	log    *logrus.Logger
}

func NewArtistRepository(client postgresql.Client, log *logrus.Logger) *artistRepository {
	return &artistRepository{
		client: client,
		log:    log,
	}
}

func (ar *artistRepository) GetByName(ctx context.Context, name string) (string, error) {
	q := `
		SELECT id
		FROM public.artist
		WHERE title = $1
	`

	var artistId string
	err := ar.client.QueryRow(ctx, q, name).Scan(&artistId)
	if err != nil {
		ar.log.Error(err)
		return "", err
	}

	return artistId, nil
}
