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
	//TODO implement me
	panic("implement me")
}

func (ar *albumRepository) DeleteAlbum(ctx context.Context, albumId string) error {
	//TODO implement me
	panic("implement me")
}
