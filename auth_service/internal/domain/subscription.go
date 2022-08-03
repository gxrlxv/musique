package domain

import "time"

type Subscription struct {
	UserId         string    `json:"user_id"`
	SubscriptionId int       `json:"subscription_id"`
	StartDate      time.Time `json:"start_date"`
	EndDate        time.Time `json:"end_date"`
}

func NewSubscription(userId string) *subscription {
	return &subscription{
		userId:         userId,
		subscriptionId: 0,
		startDate:      time.Now(),
		endDate:        time.Now(),
	}
}
