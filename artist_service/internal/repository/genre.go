package repository

import (
	"context"
	"github.com/gxrlxv/musique/pkg/client/postgresql"
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

func (genreRepository) GetByTitle(ctx context.Context, title string) (int, error) {
	//TODO implement me
	panic("implement me")
}
