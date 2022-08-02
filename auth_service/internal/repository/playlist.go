package repository

import (
	"context"
	"github.com/gxrlxv/musique/auth_service/pkg/client/postgresql"
	"github.com/sirupsen/logrus"
)

type playlistRepository struct {
	client postgresql.Client
	log    *logrus.Logger
}

func NewPlaylistRepository(client postgresql.Client, log *logrus.Logger) *playlistRepository {
	return &playlistRepository{
		client: client,
		log:    log,
	}
}

func (pr *playlistRepository) Create(ctx context.Context, userId, title string) {
	//TODO implement me
	panic("implement me")
}
