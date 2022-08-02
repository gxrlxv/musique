package usecase

import (
	"context"
	v1 "github.com/gxrlxv/musique/music_service/api/track/v1"
)

type trackUseCase struct {
}

func NewTrackUseCase() *trackUseCase {
	return &trackUseCase{}
}

func (tu *trackUseCase) AddTrack(ctx context.Context, trackId string) (*v1.AddTrackReply, error) {
	//TODO implement me
	panic("implement me")
}

func (tu *trackUseCase) RemoveTrack(ctx context.Context, trackId string) (*v1.RemoveTrackReply, error) {
	//TODO implement me

	panic("implement me")
}

func (tu *trackUseCase) GetTracks(ctx context.Context) (*v1.GetTracksReply, error) {
	//TODO implement me
	panic("implement me")
}
