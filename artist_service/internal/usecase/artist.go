package usecase

import (
	"context"
	"github.com/gxrlxv/musique/artist_service/internal/domain"
	"github.com/sirupsen/logrus"
)

type AlbumRepository interface {
	CreateAlbum(ctx context.Context, album *domain.Album) (string, error)
	DeleteAlbum(ctx context.Context, albumId string) error
	GetAlbumByID(ctx context.Context, albumId string) (domain.Album, error)
}

type TrackRepository interface {
	SaveTrack(ctx context.Context, track *domain.Track) error
	DeleteTracks(ctx context.Context, albumId string) error
}

type GenreRepository interface {
	GetByTitle(ctx context.Context, title string) (int, error)
}

type ArtistRepository interface {
	GetByName(ctx context.Context, name string) (string, error)
}

type albumUseCase struct {
	albumRepo  AlbumRepository
	trackRepo  TrackRepository
	genreRepo  GenreRepository
	artistRepo ArtistRepository
	log        *logrus.Logger
}

func NewAlbumUseCase(albumRepo AlbumRepository, trackRepo TrackRepository, genreRepo GenreRepository, artistRepo ArtistRepository, log *logrus.Logger) *albumUseCase {
	return &albumUseCase{
		albumRepo:  albumRepo,
		trackRepo:  trackRepo,
		genreRepo:  genreRepo,
		artistRepo: artistRepo,
		log:        log,
	}
}

func (a *albumUseCase) CreateAlbum(ctx context.Context, albumDTO domain.CreateAlbumDTO) (string, error) {
	a.log.Info("create album use case")

	artistId, err := a.artistRepo.GetByName(ctx, albumDTO.ArtistName)
	if err != nil {
		a.log.Error(err)
		return "", err
	}

	album := domain.NewAlbum(albumDTO.Title, artistId, albumDTO.ReleaseYear, albumDTO.NumberTracks)

	albumId, err := a.albumRepo.CreateAlbum(ctx, album)
	if err != nil {
		a.log.Error(err)
		return "", err
	}

	for i := 0; i < albumDTO.NumberTracks; i++ {
		genreId, err := a.genreRepo.GetByTitle(ctx, albumDTO.Tracks[i].Genre)
		if err != nil {
			a.log.Error(err)
			return "", err
		}

		track := domain.NewTrack(albumDTO.Tracks[i].Title, artistId, albumId, albumDTO.ReleaseYear, genreId, albumDTO.Tracks[i].Milliseconds, albumDTO.Tracks[i].Bytes)

		err = a.trackRepo.SaveTrack(ctx, track)
		if err != nil {
			a.log.Error(err)
			return "", err
		}
	}

	return albumId, nil
}

func (a *albumUseCase) DeleteAlbum(ctx context.Context, albumId string) (bool, error) {
	a.log.Info("delete album use case")

	_, err := a.albumRepo.GetAlbumByID(ctx, albumId)
	if err != nil {
		return false, err
	}

	err = a.trackRepo.DeleteTracks(ctx, albumId)
	if err != nil {
		a.log.Error(err)
		return false, err
	}

	err = a.albumRepo.DeleteAlbum(ctx, albumId)
	if err != nil {
		a.log.Error(err)
		return false, err
	}

	return true, nil
}
