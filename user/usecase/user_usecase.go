package usecase

import (
	"context"
	"mygram-api/domain"
)

type userUseCase struct {
	userRepository domain.UserRepository
}

func NewUserUseCase(userRepository domain.UserRepository) *userUseCase {
	return &userUseCase{userRepository}
}

func (userUseCase *userUseCase) UserRegister(ctx context.Context, user *domain.User) (err error) {
	if err = userUseCase.userRepository.UserRegister(ctx, user); err != nil {
		return err
	}

	return
}

func (userUseCase *userUseCase) UserLogin(ctx context.Context, user *domain.User) (err error) {
	if err = userUseCase.userRepository.UserLogin(ctx, user); err != nil {
		return nil
	}

	return
}

func (userUseCase *userUseCase) UpdateUser(ctx context.Context, user domain.User, id string) (u domain.User, err error) {
	u, err = userUseCase.userRepository.UpdateUser(ctx, user, id)

	if err != nil {
		return u, err
	}

	return u, nil
}

func (userUseCase *userUseCase) DeleteUser(ctx context.Context, id string) (err error) {
	if err = userUseCase.userRepository.DeleteUser(ctx, id); err != nil {
		return err
	}

	return
}
