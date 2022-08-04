package service

import (
	"context"
	v1 "github.com/gxrlxv/musique/artist_service/api/artist/v1"
	"github.com/gxrlxv/musique/artist_service/internal/domain"
	"github.com/sirupsen/logrus"
)

type ArtistUseCase interface {
	CreateAlbum(ctx context.Context, album domain.CreateAlbumDTO) error
	DeleteAlbum(ctx context.Context, albumId string) error
}

type ArtistService struct {
	v1.UnimplementedArtistServer
	uc  ArtistUseCase
	log *logrus.Logger
}

func NewArtistService(useCase ArtistUseCase, log *logrus.Logger) *ArtistService {
	return &ArtistService{
		UnimplementedArtistServer: v1.UnimplementedArtistServer{},
		uc:                        useCase,
		log:                       log,
	}
}

func (as *ArtistService) CreateAlbum(ctx context.Context, in *v1.CreateAlbumRequest) (*v1.CreateAlbumReply, error) {

	albumDTO := domain.CreateAlbumDTO{
		Title:       in.AlbumTitle,
		ReleaseYear: in.ReleaseYear,
		ArtistName:  in.ArtistName,
	}

	err := as.uc.CreateAlbum(ctx, albumDTO)
	if err != nil {
		return nil, err
	}
	return &v1.CreateAlbumReply{
		TrackId: "",
	}, nil
}

func (as *ArtistService) DeleteAlbum(ctx context.Context, in *v1.DeleteAlbumRequest) (*v1.DeleteAlbumReply, error) {

	err := as.uc.DeleteAlbum(ctx, in.AlbumId)
	if err != nil {
		return nil, err
	}

	return &v1.DeleteAlbumReply{
		AlbumId: "",
	}, nil
}