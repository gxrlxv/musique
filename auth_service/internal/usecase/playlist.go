package usecase

import "context"

type PlaylistRepository interface {
	Create(ctx context.Context, userId, title string)
}

func (a *authUseCase) NewPlaylist(ctx context.Context, userId string) error {
	//TODO implement me

	panic("implement me")
}
