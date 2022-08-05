package repository

import (
	"context"
	"github.com/gxrlxv/musique/pkg/client/postgresql"
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

func (artistRepository) GetByName(ctx context.Context, name string) (string, error) {
	//TODO implement me
	panic("implement me")
}
