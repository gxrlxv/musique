package usecase

import (
	"context"
	"github.com/gxrlxv/musique/music_service/internal/domain"
	"github.com/gxrlxv/musique/music_service/internal/repository"
	"github.com/sirupsen/logrus"
)

type musicUseCase struct {
	playlistRepo repository.PlaylistRepository
	trackRepo    repository.TrackRepository
	albumRepo    repository.AlbumRepository
	log          *logrus.Logger
}

func NewMusicUseCase(playlistRepo repository.PlaylistRepository, trackRepo repository.TrackRepository,
	albumRepo repository.AlbumRepository, log *logrus.Logger) *musicUseCase {
	return &musicUseCase{
		playlistRepo: playlistRepo,
		trackRepo:    trackRepo,
		albumRepo:    albumRepo,
		log:          log,
	}
}

func (m *musicUseCase) AddTrack(ctx context.Context, playlistId, trackId string) (bool, error) {
	// Проверяем существует ли плейлист
	playlist, err := m.playlistRepo.GetPlaylist(ctx, playlistId)
	if err != nil {
		return false, err
	}

	// Проверяем существует ли трека в бд
	_, err = m.trackRepo.GetTrack(ctx, trackId)
	if err != nil {
		return false, err
	}

	// Сохраняем трек в плейлист
	err = m.playlistRepo.SaveTrack(ctx, playlistId, trackId)
	if err != nil {
		return false, err
	}

	// Обновляем количество треков в плейлисте
	playlist.NumberTracks += 1

	err = m.playlistRepo.UpdatePlaylist(ctx, playlist)
	if err != nil {
		return false, err
	}

	return true, err
}

func (m *musicUseCase) AddAlbum(ctx context.Context, playlistId, albumId string) (bool, error) {
	// Проверяем существует ли плейлист
	playlist, err := m.playlistRepo.GetPlaylist(ctx, playlistId)
	if err != nil {
		return false, err
	}

	// Проверяем существует ли альбом
	_, err = m.albumRepo.GetAlbum(ctx, albumId)
	if err != nil {
		return false, err
	}

	// Получаем слайс айдишников всех треков в альбоме
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

	// Обновляем количество треков в плейлисте
	playlist.NumberTracks += len(tracks)

	err = m.playlistRepo.UpdatePlaylist(ctx, playlist)
	if err != nil {
		return false, err
	}

	return true, err
}

func (m *musicUseCase) DeleteTrack(ctx context.Context, playlistId, trackId string) (bool, error) {
	// Проверка существует ли трек в плейлисте
	err := m.playlistRepo.GetTrack(ctx, playlistId, trackId)
	if err != nil {
		return false, err
	}

	err = m.playlistRepo.DeleteTrack(ctx, playlistId, trackId)
	if err != nil {
		return false, err
	}

	return true, err
}

func (m *musicUseCase) GetTrack(ctx context.Context, playlistId, trackId string) (domain.Track, error) {
	// Проверяем существует ли плейлист
	_, err := m.playlistRepo.GetPlaylist(ctx, playlistId)
	if err != nil {
		return domain.Track{}, err
	}

	err = m.playlistRepo.GetTrack(ctx, playlistId, trackId)
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
	// Проверяем существует ли плейлист
	_, err := m.playlistRepo.GetPlaylist(ctx, playlistId)
	if err != nil {
		return nil, err
	}

	// Получаем слайс айдишников всех треков в плейлисте
	tracksId, err := m.playlistRepo.GetAllTracks(ctx, playlistId)
	if err != nil {
		return nil, err
	}

	tracks := make([]domain.Track, 0, len(tracksId))

	for i := range tracksId {
		track, err := m.trackRepo.GetTrack(ctx, tracksId[i])
		if err != nil {
			return nil, err
		}

		tracks = append(tracks, track)
	}

	return tracks, err
}
