package token

import "time"

type Maker interface {
	CreateToken(userID int64, email string, role string, duration time.Duration) (string, error)
	VerifyToken(token string) (*Payload, error)
}
