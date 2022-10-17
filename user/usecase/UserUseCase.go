package usecase

import (
	"context"
	"mygram-api/domain"
	"time"
)

type userUseCase struct {
	userRepository domain.UserRepository
	contextTimeOut time.Duration
}

func NewUser(userRepository domain.UserRepository, timeout time.Duration) domain.UserUseCase {
	return &userUseCase{
		userRepository: userRepository,
		contextTimeOut: timeout,
	}
}

func (userUseCase *userUseCase) CheckUsername(ctx context.Context, username string) (err error) {
	ctx, cancel := context.WithTimeout(ctx, userUseCase.contextTimeOut)

	defer cancel()

	return userUseCase.userRepository.CheckUsername(ctx, username)
}

func (userUseCase *userUseCase) CheckEmail(ctx context.Context, email string) (err error) {
	ctx, cancel := context.WithTimeout(ctx, userUseCase.contextTimeOut)

	defer cancel()

	return userUseCase.userRepository.CheckEmail(ctx, email)
}

func (userUseCase *userUseCase) AddUser(ctx context.Context, user *domain.User) (err error) {
	ctx, cancel := context.WithTimeout(ctx, userUseCase.contextTimeOut)

	defer cancel()

	if err := userUseCase.userRepository.CheckUsername(ctx, user.Username); err != nil {
		return domain.ErrConflict
	}

	if err := userUseCase.userRepository.CheckEmail(ctx, user.Email); err != nil {
		return domain.ErrConflict
	}

	err = userUseCase.userRepository.AddUser(ctx, user)

	return
}

func (userUseCase *userUseCase) GetUsers(ctx context.Context) (users []domain.User, err error) {
	ctx, cancel := context.WithTimeout(ctx, userUseCase.contextTimeOut)

	defer cancel()

	return userUseCase.userRepository.GetUsers(ctx)
}

func (userUseCase *userUseCase) CheckUser(ctx context.Context, id string) (err error) {
	ctx, cancel := context.WithTimeout(ctx, userUseCase.contextTimeOut)

	defer cancel()

	return userUseCase.userRepository.CheckUser(ctx, id)
}

func (userUseCase *userUseCase) DeleteUser(ctx context.Context, id string) (err error) {
	ctx, cancel := context.WithTimeout(ctx, userUseCase.contextTimeOut)

	defer cancel()

	if err := userUseCase.userRepository.CheckUser(ctx, id); err != nil {
		return err
	}

	return userUseCase.userRepository.DeleteUser(ctx, id)
}
