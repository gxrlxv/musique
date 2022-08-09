package domain

import "time"

type Session struct {
	UserId       string    `json:"user_id"`
	RefreshToken string    `json:"refresh_token"`
	ExpiresAt    time.Time `json:"expires_at"`
}

func NewSession(userId, refreshToken string, ttl time.Duration) Session {
	return Session{
		UserId:       userId,
		RefreshToken: refreshToken,
		ExpiresAt:    time.Now().Add(ttl),
	}
}
