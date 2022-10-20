package domain

import "time"

type session struct {
	RefreshToken string
	ExpireAt     time.Time
}
