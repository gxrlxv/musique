package usecase

import (
	"context"
)

type PlaylistRepository interface {
	Create(ctx context.Context, userId string) error
}

func (a *authUseCase) NewPlaylist(ctx context.Context, userId string) error {
	a.log.Info("NewPlaylist use case")

	err := a.playlistRepo.Create(ctx, userId)
	if err != nil {
		a.log.Error(err)
		return internalErr(err)
	}

	return nil
}
