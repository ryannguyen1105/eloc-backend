package service

import (
	"github.com/ryannguyen1105/eloc-backend/common/token"
	"github.com/ryannguyen1105/eloc-backend/services/eloc_auth/config"
	db "github.com/ryannguyen1105/eloc-backend/services/eloc_auth/db/sqlc"
)

type AuthService struct {
	store      db.Store
	tokenMaker token.Maker
	config     config.Config
}

func NewAuthService(store db.Store, tokenMaker token.Maker, config config.Config) *AuthService {
	return &AuthService{
		store:      store,
		tokenMaker: tokenMaker,
		config:     config,
	}
}
