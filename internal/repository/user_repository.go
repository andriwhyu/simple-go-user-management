package repository

import (
	"context"
	"database/sql"
	"github.com/andriwhyu/simple-go-user-management/internal/domain"
)

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) domain.UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Create(ctx context.Context, user *domain.User) error {
	return nil
}

func (r *userRepository) GetByID(ctx context.Context, id int) (*domain.User, error) {
	return nil, nil
}

func (r *userRepository) GetByEmail(ctx context.Context, email string) (*domain.User, error) {
	return nil, nil
}

func (r *userRepository) GetAll(ctx context.Context) ([]*domain.User, error) {
	return nil, nil
}
func (r *userRepository) Update(ctx context.Context, user *domain.User) error {
	return nil
}

func (r *userRepository) Delete(ctx context.Context, id int) error {
	return nil
}
