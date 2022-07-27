package service

import (
	"context"
	v1 "github.com/gxrlxv/musique/artist_service/api/artist/v1"
)

type ArtistService struct {
	v1.UnimplementedArtistServer
}

func NewArtistService() *ArtistService {
	return &ArtistService{v1.UnimplementedArtistServer{}}
}

func (a *ArtistService) CreateSingle(ctx context.Context, in *v1.CreateSingleRequest) (*v1.CreateSingleReply, error) {
	return &v1.CreateSingleReply{}, nil
}
func (a *ArtistService) CreateAlbum(ctx context.Context, in *v1.CreateAlbumRequest) (*v1.CreateAlbumReply, error) {
	return &v1.CreateAlbumReply{}, nil
}
func (a *ArtistService) DeleteSingle(ctx context.Context, in *v1.DeleteSingleRequest) (*v1.DeleteSingleReply, error) {
	return &v1.DeleteSingleReply{}, nil
}
func (a *ArtistService) DeleteAlbum(ctx context.Context, in *v1.DeleteAlbumRequest) (*v1.DeleteAlbumReply, error) {
	return &v1.DeleteAlbumReply{}, nil
}
