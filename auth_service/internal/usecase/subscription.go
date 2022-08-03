package usecase

import (
	"context"
	"github.com/gxrlxv/musique/auth_service/internal/domain"
	"time"
)

type SubscriptionRepository interface {
	CreateSubscription(ctx context.Context, subscription *domain.Subscription) error
	GetBySubscriptionId(ctx context.Context, id int) (error, time.Duration)
}

func (a *authUseCase) NewSubscription(ctx context.Context, userId string) error {
	a.log.Info("New Subscription use case")

	err, subDuration :=

		domain.NewSubscription(userid)

	err := a.subscriptionRepo.CreateSubscription(ctx, subscription)
	if err != nil {
		return err
	}

	return nil
}
