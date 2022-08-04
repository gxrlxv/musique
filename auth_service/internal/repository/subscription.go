package repository

import (
	"context"
	"github.com/gxrlxv/musique/auth_service/internal/domain"
	"github.com/gxrlxv/musique/auth_service/pkg/client/postgresql"
	"github.com/sirupsen/logrus"
	"time"
)

type subscriptionRepository struct {
	client postgresql.Client
	log    *logrus.Logger
}

func NewSubscriptionRepository(client postgresql.Client, log *logrus.Logger) *subscriptionRepository {
	return &subscriptionRepository{
		client: client,
		log:    log,
	}
}

func (sr *subscriptionRepository) CreateSubscription(ctx context.Context, subscription *domain.Subscription) error {
	q := `
			INSERT INTO public.subscription_validity
    			(subscription_id, start_date, end_date, user_id)
			VALUES
    			($1,$2,$3,$4)`

	_, err := sr.client.Exec(ctx, q, subscription.SubscriptionId, subscription.StartDate, subscription.EndDate, subscription.UserId)
	if err != nil {
		sr.log.Error(err)
		return err
	}

	return nil
}

func (sr *subscriptionRepository) GetBySubscriptionId(ctx context.Context, id int) (time.Duration, error) {
	q := `
		SELECT duration 
		FROM public.subscription 
		WHERE id = $1
	`

	var duration time.Duration
	err := sr.client.QueryRow(ctx, q, id).Scan(&duration)
	if err != nil {
		sr.log.Error(err)
		return 0, err
	}

	return duration, nil
}
