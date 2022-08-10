package usecase

import (
	"context"
	"github.com/gxrlxv/musique/music_service/internal/domain"
	"github.com/sirupsen/logrus"
)

type PlaylistRepository interface {
	SaveTrack(ctx context.Context, playlistId, trackId string) error
	DeleteTrack(ctx context.Context, playlistId, trackId string) error
	GetTrack(ctx context.Context, playlistId, trackId string) error
	GetAllTracks(ctx context.Context, playlistId string) ([]string, error)
}

type TrackRepository interface {
	GetTrack(ctx context.Context, trackId string) (domain.Track, error)
	GetTracksByAlbumId(ctx context.Context, albumId string) ([]domain.Track, error)
}

type musicUseCase struct {
	playlistRepo PlaylistRepository
	trackRepo    TrackRepository
	log          *logrus.Logger
}

func NewMusicUseCase(musicRepo PlaylistRepository, trackRepo TrackRepository, log *logrus.Logger) *musicUseCase {
	return &musicUseCase{
		playlistRepo: musicRepo,
		trackRepo:    trackRepo,
		log:          log,
	}
}

func (m *musicUseCase) AddTrack(ctx context.Context, playlistId, trackId string) (bool, error) {
	_, err := m.trackRepo.GetTrack(ctx, trackId)
	if err != nil {
		return false, err
	}

	err = m.playlistRepo.SaveTrack(ctx, playlistId, trackId)
	if err != nil {
		return false, err
	}

	return true, err
}

func (m *musicUseCase) AddAlbum(ctx context.Context, playlistId, albumId string) (bool, error) {
	tracks, err := m.trackRepo.GetTracksByAlbumId(ctx, albumId)
	if err != nil {
		return false, err
	}

	for i := range tracks {
		err := m.playlistRepo.SaveTrack(ctx, playlistId, tracks[i].Id)
		if err != nil {
			return false, err
		}
	}

	return true, err
}

func (m *musicUseCase) DeleteTrack(ctx context.Context, playlistId, trackId string) (bool, error) {
	err := m.playlistRepo.DeleteTrack(ctx, playlistId, trackId)
	if err != nil {
		return false, err
	}

	return true, err
}

func (m *musicUseCase) GetTrack(ctx context.Context, playlistId, trackId string) (domain.Track, error) {
	err := m.playlistRepo.GetTrack(ctx, playlistId, trackId)
	if err != nil {
		return domain.Track{}, err
	}

	track, err := m.trackRepo.GetTrack(ctx, trackId)
	if err != nil {
		return domain.Track{}, err
	}

	return track, err
}

func (m *musicUseCase) GetAllTracks(ctx context.Context, playlistId string) ([]domain.Track, error) {
	tracksId, err := m.playlistRepo.GetAllTracks(ctx, playlistId)
	if err != nil {
		return nil, err
	}

	tracks := make([]domain.Track, len(tracksId))

	for i := range tracksId {
		track, err := m.trackRepo.GetTrack(ctx, tracksId[i])
		if err != nil {
			return nil, err
		}

		tracks = append(tracks, track)
	}

	return tracks, err
}
