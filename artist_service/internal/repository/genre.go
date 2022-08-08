package repository

import (
	"context"
	"github.com/gxrlxv/musique/artist_service/pkg/client/postgresql"
	"github.com/sirupsen/logrus"
)

type genreRepository struct {
	client postgresql.Client
	log    *logrus.Logger
}

func NewGenreRepository(client postgresql.Client, log *logrus.Logger) *genreRepository {
	return &genreRepository{
		client: client,
		log:    log,
	}
}

func (gr *genreRepository) GetByTitle(ctx context.Context, title string) (int, error) {
	q := `
		SELECT id
		FROM public.genre
		WHERE title = $1
	`

	var genreId int
	err := gr.client.QueryRow(ctx, q, title).Scan(&genreId)
	if err != nil {
		gr.log.Error(err)
		return 0, err
	}

	return genreId, err
}
