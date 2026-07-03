package service

import db "github.com/ryannguyen1105/eloc-backend/services/eloc_auth/db/sqlc"

type AuthService struct {
	store db.Store
}

func NewAuthService(store db.Store) *AuthService {
	return &AuthService{store: store}
}
