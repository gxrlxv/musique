package domain

import "time"

type Subscription struct {
	UserId         string    `json:"user_id"`
	SubscriptionId int       `json:"subscription_id"`
	StartDate      time.Time `json:"start_date"`
	EndDate        time.Time `json:"end_date"`
}

func NewSubscription(userId string, subId int, duration time.Duration) *Subscription {
	return &Subscription{
		UserId:         userId,
		SubscriptionId: subId,
		StartDate:      time.Now(),
		EndDate:        time.Now().Add(duration),
	}
}
