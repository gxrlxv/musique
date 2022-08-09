package service

import (
	"context"
	v1 "github.com/gxrlxv/musique/artist_service/api/artist/v1"
	"github.com/gxrlxv/musique/artist_service/internal/domain"
	"github.com/sirupsen/logrus"
)

type ArtistUseCase interface {
	CreateAlbum(ctx context.Context, album domain.CreateAlbumDTO) (string, error)
	DeleteAlbum(ctx context.Context, albumId string) (bool, error)
	DeleteTrack(ctx context.Context, albumId, trackId string) (bool, error)
	AddTrack(ctx context.Context, albumId string, track domain.CreateTrackDTO) (bool, error)
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
	as.log.Info("create album service")

	var tracksDTO []*domain.CreateTrackDTO

	for i := range in.Tracks {
		trackDTO := domain.CreateTrackDTO{
			Title:        in.Tracks[i].Title,
			Genre:        in.Tracks[i].Genre,
			Milliseconds: in.Tracks[i].Milliseconds,
			Bytes:        in.Tracks[i].Bytes,
		}

		tracksDTO = append(tracksDTO, &trackDTO)
	}

	albumDTO := domain.CreateAlbumDTO{
		Title:        in.AlbumTitle,
		ReleaseYear:  in.ReleaseYear,
		ArtistName:   in.ArtistName,
		NumberTracks: len(in.Tracks),
		Tracks:       tracksDTO,
	}

	albumId, err := as.uc.CreateAlbum(ctx, albumDTO)
	if err != nil {
		return &v1.CreateAlbumReply{}, err
	}

	return &v1.CreateAlbumReply{
		AlbumId: albumId,
	}, err
}

func (as *ArtistService) DeleteAlbum(ctx context.Context, in *v1.DeleteAlbumRequest) (*v1.DeleteAlbumReply, error) {
	as.log.Info("delete album service")

	success, err := as.uc.DeleteAlbum(ctx, in.AlbumId)
	if err != nil {
		return &v1.DeleteAlbumReply{
			Success: success,
		}, err
	}

	return &v1.DeleteAlbumReply{
		Success: success,
	}, err
}

func (as *ArtistService) DeleteTrack(ctx context.Context, in *v1.DeleteTrackRequest) (*v1.DeleteTrackReply, error) {
	as.log.Info("delete track service")

	success, err := as.uc.DeleteTrack(ctx, in.AlbumId, in.TrackId)
	if err != nil {
		return &v1.DeleteTrackReply{
			Success: success,
		}, err
	}

	return &v1.DeleteTrackReply{
		Success: success,
	}, err
}

func (as *ArtistService) AddTrack(ctx context.Context, in *v1.AddTrackRequest) (*v1.AddTrackReply, error) {
	as.log.Info("add track service")

	track := domain.CreateTrackDTO{
		Title:        in.Track.Title,
		Genre:        in.Track.Genre,
		Milliseconds: in.Track.Milliseconds,
		Bytes:        in.Track.Bytes,
	}

	success, err := as.uc.AddTrack(ctx, in.AlbumId, track)
	if err != nil {
		return &v1.AddTrackReply{}, err
	}

	return &v1.AddTrackReply{
		Success: success,
	}, nil
}