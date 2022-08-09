package usecase

import (
	"context"
	"github.com/gxrlxv/musique/auth_service/internal/domain"
	"time"
)

var freeSubId = 1

type SubscriptionRepository interface {
	CreateSubscription(ctx context.Context, subscription domain.Subscription) error
	GetBySubscriptionId(ctx context.Context, id int) (time.Duration, error)
}

func (a *authUseCase) NewSubscription(ctx context.Context, userId string) error {
	a.log.Info("New Subscription use case")

	subDuration, err := a.subscriptionRepo.GetBySubscriptionId(ctx, freeSubId)
	if err != nil {
		return err
	}

	subscription := domain.NewSubscription(userId, freeSubId, subDuration)

	err = a.subscriptionRepo.CreateSubscription(ctx, subscription)
	if err != nil {
		return err
	}

	return err
}
