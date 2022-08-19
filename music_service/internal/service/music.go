package service

import (
	"context"
	v1 "github.com/gxrlxv/musique/music_service/api/music/v1"
	"github.com/gxrlxv/musique/music_service/internal/usecase"
	"github.com/sirupsen/logrus"
)

type MusicService struct {
	v1.UnimplementedMusicServer
	uc  *usecase.UseCase
	log *logrus.Logger
}

func NewMusicService(useCase *usecase.UseCase, log *logrus.Logger) *MusicService {
	return &MusicService{
		UnimplementedMusicServer: v1.UnimplementedMusicServer{},
		uc:                       useCase,
		log:                      log,
	}
}

func (ms *MusicService) AddTrack(ctx context.Context, in *v1.AddTrackRequest) (*v1.AddTrackReply, error) {
	ms.log.Info("Add track service")

	success, err := ms.uc.AddTrack(ctx, in.PlaylistId, in.TrackId)
	if err != nil {
		return &v1.AddTrackReply{
			Success: success,
		}, err
	}

	return &v1.AddTrackReply{
		Success: success,
	}, err
}

func (ms *MusicService) AddAlbum(ctx context.Context, in *v1.AddAlbumRequest) (*v1.AddAlbumReply, error) {
	ms.log.Info("Add album service")

	success, err := ms.uc.AddAlbum(ctx, in.PlaylistId, in.AlbumId)
	if err != nil {
		return &v1.AddAlbumReply{
			Success: success,
		}, err
	}

	return &v1.AddAlbumReply{
		Success: success,
	}, err
}

func (ms *MusicService) DeleteTrack(ctx context.Context, in *v1.DeleteTrackRequest) (*v1.DeleteTrackReply, error) {
	ms.log.Info("Delete track service")

	success, err := ms.uc.DeleteTrack(ctx, in.PlaylistId, in.TrackId)
	if err != nil {
		return &v1.DeleteTrackReply{
			Success: success,
		}, err
	}

	return &v1.DeleteTrackReply{
		Success: success,
	}, err
}

func (ms *MusicService) GetTrack(ctx context.Context, in *v1.GetTrackRequest) (*v1.GetTrackReply, error) {
	ms.log.Info("Get track service")

	track, err := ms.uc.GetTrack(ctx, in.PlaylistId, in.TrackId)
	if err != nil {
		return &v1.GetTrackReply{
			Track: &v1.Track{},
		}, err
	}

	return &v1.GetTrackReply{
		Track: &v1.Track{
			Title:        track.Title,
			Milliseconds: track.Milliseconds,
			Bytes:        track.Bytes,
		},
	}, err
}

func (ms *MusicService) GetAllTracks(ctx context.Context, in *v1.GetTracksRequest) (*v1.GetTracksReply, error) {
	ms.log.Info("Get all tracks service")

	tracksDTO, err := ms.uc.GetAllTracks(ctx, in.PlaylistId)
	if err != nil {
		return &v1.GetTracksReply{
			Tracks: []*v1.Track{},
		}, err
	}

	tracks := make([]*v1.Track, 0, len(tracksDTO))

	for i := range tracksDTO {
		track := &v1.Track{
			Title:        tracksDTO[i].Title,
			Milliseconds: tracksDTO[i].Milliseconds,
			Bytes:        tracksDTO[i].Bytes,
		}

		tracks = append(tracks, track)
	}

	return &v1.GetTracksReply{
		Tracks: tracks,
	}, err
}
