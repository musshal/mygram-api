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

func (userUseCase *userUseCase) Register(ctx context.Context, user *domain.User) (err error) {
	if err = userUseCase.userRepository.Register(ctx, user); err != nil {
		return err
	}

	return
}

func (userUseCase *userUseCase) Login(ctx context.Context, user *domain.User) (err error) {
	if err = userUseCase.userRepository.Login(ctx, user); err != nil {
		return nil
	}

	return
}

func (userUseCase *userUseCase) Update(ctx context.Context, user domain.User, id uint) (u domain.User, err error) {
	u, err = userUseCase.userRepository.Update(ctx, user, id)

	if err != nil {
		return u, err
	}

	return u, nil
}

func (userUseCase *userUseCase) Delete(ctx context.Context, id uint) (err error) {
	if err = userUseCase.userRepository.Delete(ctx, id); err != nil {
		return err
	}

	return
}
