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
	if id < 1 {
		return nil, errors.New("invalid user ID")
	}

	user, err := u.userRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return user, nil
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
	if id < 1 {
		return nil, errors.New("invalid user ID")
	}

	// Check if user exists
	user, err := u.userRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if name == "" {
		name = user.Name
	}

	if email == "" {
		email = user.Email
	}

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

	// Check if email is being changed and if new email already exists
	if user.Email != email {
		existingUser, err := u.userRepo.GetByEmail(ctx, email)
		if err != nil {
			return nil, err
		}

		if existingUser != nil {
			return nil, errors.New("email already exists")
		}
	}

	user.Name = name
	user.Email = email

	err = u.userRepo.Update(ctx, user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *userUsecase) Delete(ctx context.Context, id int) error {
	if id < 1 {
		return errors.New("invalid user ID")
	}

	err := u.userRepo.Delete(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
