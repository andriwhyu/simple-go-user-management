package domain

import (
	"context"
	"time"
)

type User struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// UserRepository represents user repository contract
type UserRepository interface {
	GetByID(ctx context.Context, id int) (*User, error)
	Create(ctx context.Context, user *User) error
	GetByEmail(ctx context.Context, email string) (*User, error)
	GetAll(ctx context.Context) ([]*User, error)
	Update(ctx context.Context, user *User) error
	Delete(ctx context.Context, id int) error
}

type UserUsecase interface {
	GetByID(ctx context.Context, id int) (*User, error)
	Create(ctx context.Context, name, email string) (*User, error)
	GetAll(ctx context.Context) ([]*User, error)
	Update(ctx context.Context, id int, name, email string) (*User, error)
	Delete(ctx context.Context, id int) error
}
