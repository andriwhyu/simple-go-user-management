package usecase

import (
	"context"
	"errors"
	"github.com/andriwhyu/simple-go-user-management/internal/domain"
	"github.com/andriwhyu/simple-go-user-management/internal/utils"
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
	// Validate input
	if strings.TrimSpace(name) == "" {
		return nil, errors.New("name is required")
	}
	if strings.TrimSpace(email) == "" {
		return nil, errors.New("email is required")
	}
	if !utils.IsValidEmail(email) {
		return nil, errors.New("invalid email format")
	}

	// Check if email already exists
	existingUser, err := u.userRepo.GetByEmail(ctx, email)
	if err != nil {
		return nil, err
	}
	if existingUser != nil {
		return nil, errors.New("email already exists")
	}

	user := &domain.User{
		Name:  name,
		Email: email,
	}

	err = u.userRepo.Create(ctx, user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *userUsecase) GetAll(ctx context.Context) ([]*domain.User, error) {
	users, err := u.userRepo.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (u *userUsecase) Update(ctx context.Context, id int, name, email string) (*domain.User, error) {
	return nil, nil
}

func (u *userUsecase) Delete(ctx context.Context, id int) error {
	return nil
}
