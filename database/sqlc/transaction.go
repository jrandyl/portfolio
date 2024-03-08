package sqlc

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/google/uuid"
)

type Store interface {
	Querier
	RegisterTx(c context.Context, arg RegisterTxParams) (RegisterTxResult, error)
}

type Transaction struct {
	*Queries
	db *sql.DB
}

func NewTransaction(db *sql.DB) Store {
	return &Transaction{
		Queries: New(db),
		db:      db,
	}
}

func (t *Transaction) execTransaction(ctx context.Context, fn func(*Queries) error) error {
	tx, err := t.db.BeginTx(ctx, &sql.TxOptions{
		Isolation: sql.LevelSerializable,
	})

	if err != nil {
		return err
	}

	q := New(tx)
	err = fn(q)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("transaction error: %v, rollback error: %v", err, rbErr)
		}
		return err
	}

	return tx.Commit()
}

type RegisterTxParams struct {
	Username      string `json:"username"`
	Password      []byte `json:"password"`
	FirstName     string `json:"first_name"`
	MiddleName    string `json:"middle_name"`
	LastName      string `json:"last_name"`
	FullName      string `json:"full_name"`
	Email         string `json:"email"`
	ContactNumber string `json:"contact_number"`
}

type UserResponse struct {
	ID                string    `json:"id"`
	Username          string    `json:"username"`
	PasswordUpdatedAt string    `json:"password_updated_at"`
	LastLogin         string    `json:"last_login"`
	Status            string    `json:"status"`
	CreatedAt         time.Time `json:"created_at"`
}

type RegisterTxResult struct {
	User    UserResponse `json:"user"`
	Profile Profile      `json:"profile"`
}

func (t *Transaction) RegisterTx(c context.Context, arg RegisterTxParams) (RegisterTxResult, error) {
	var result RegisterTxResult

	err := t.execTransaction(c, func(q *Queries) error {
		var err error

		userID := uuid.New().String()
		user, err := q.CreateUser(c, CreateUserParams{
			ID:       userID,
			Username: arg.Username,
			Password: arg.Password,
		})

		if err != nil {
			return err
		}

		result.User = UserResponse{
			ID:                user.ID,
			Username:          user.Username,
			PasswordUpdatedAt: user.PasswordUpdatedAt,
			LastLogin:         user.LastLogin,
			Status:            user.Status,
			CreatedAt:         user.CreatedAt,
		}

		profileID := uuid.New().String()

		result.Profile, err = q.CreateProfile(c, CreateProfileParams{
			ID:            profileID,
			UserID:        result.User.ID,
			FirstName:     arg.FirstName,
			MiddleName:    arg.MiddleName,
			LastName:      arg.LastName,
			FullName:      arg.FullName,
			Email:         arg.Email,
			ContactNumber: arg.ContactNumber,
		})

		if err != nil {
			return err
		}

		return nil

	})

	return result, err
}
