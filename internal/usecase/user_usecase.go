package usecase

import (
	"context"
	"github.com/andriwhyu/simple-go-user-management/internal/domain"
	"strings"
)

type userUsecase struct {
	userRepo domain.UserRepository
}

func NewUserUsecase(userRepo domain.UserRepository) domain.UserUsecase {
	return &userUsecase{userRepo: userRepo}
}

func (u *userUsecase) GetByID(ctx context.Context, id int) (*domain.User, error) {
	return nil, nil
}

func (u *userUsecase) Create(ctx context.Context, name, email string) (*domain.User, error) {
	return nil, nil
}

func (u *userUsecase) GetAll(ctx context.Context) ([]*domain.User, error) {
	return nil, nil
}

func (u *userUsecase) Update(ctx context.Context, id int, name, email string) (*domain.User, error) {
	return nil, nil
}

func (u *userUsecase) Delete(ctx context.Context, id int) error {
	return nil
}

// isValidEmail performs basic email validation
func isValidEmail(email string) bool {
	// Basic email validation
	if !strings.Contains(email, "@") {
		return false
	}
	parts := strings.Split(email, "@")
	if len(parts) != 2 {
		return false
	}
	if len(parts[0]) == 0 || len(parts[1]) == 0 {
		return false
	}
	if !strings.Contains(parts[1], ".") {
		return false
	}
	return true
}
