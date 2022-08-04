package usecase

import (
	"context"
	"github.com/gxrlxv/musique/artist_service/internal/domain"
	"github.com/sirupsen/logrus"
)

type AlbumRepository interface {
	CreateAlbum(ctx context.Context, album domain.Album) error
	DeleteAlbum(ctx context.Context, albumId string) error
}

type albumUseCase struct {
	albumRepo AlbumRepository
	log       *logrus.Logger
}

func NewAlbumUseCase(albumRepo AlbumRepository, log *logrus.Logger) *albumUseCase {
	return &albumUseCase{
		albumRepo: albumRepo,
		log:       log,
	}
}

func (albumUseCase) CreateAlbum(ctx context.Context, album domain.CreateAlbumDTO) error {
	//TODO implement me\
	panic("implement me")
}

func (albumUseCase) DeleteAlbum(ctx context.Context, albumId string) error {
	//TODO implement me
	panic("implement me")
}
