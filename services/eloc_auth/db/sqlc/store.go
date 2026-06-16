package db

import (
	"context"
	"database/sql"
	"fmt"
	"time"
)

type Store struct {
	*Queries
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{
		db:      db,
		Queries: New(db),
	}
}

func (store *Store) execTx(ctx context.Context, fn func(queries *Queries) error) error {
	tx, err := store.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	q := New(tx)
	err = fn(q)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx err: %v, rollback: %v", err, rbErr)
		}
		return err
	}
	return tx.Commit()
}

type RegisterTxParams struct {
	Email        string    `json:"email"`
	PasswordHash string    `json:"password_hash"`
	FullName     string    `json:"full_name"`
	RoleID       int64     `json:"role_id"`
	RefreshToken string    `json:"refresh_token"`
	ExpiresAt    time.Time `json:"expires_at"`
}

type RegisterTxResult struct {
	User  User      `json:"user"`
	Token UserToken `json:"user_token"`
}

func (store *Store) RegisterTx(ctx context.Context, arg RegisterTxParams) (RegisterTxResult, error) {
	var result RegisterTxResult

	err := store.execTx(ctx, func(q *Queries) error {
		var err error

		result.User, err = q.CreateUser(ctx, CreateUserParams{
			Email:        arg.Email,
			PasswordHash: arg.PasswordHash,
			Fullname:     arg.FullName,
			RoleID:       arg.RoleID,
		})
		if err != nil {
			return err
		}
		result.Token, err = q.CreateUserToken(ctx, CreateUserTokenParams{
			UserID:       result.User.ID,
			RefreshToken: arg.RefreshToken,
			ExpiresAt:    arg.ExpiresAt,
		})
		if err != nil {
			return err
		}

		return nil
	})
	return result, err
}
