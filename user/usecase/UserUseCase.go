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

func (userUseCase *userUseCase) AddUser(ctx context.Context, user *domain.User) (err error) {
	ctx, cancel := context.WithTimeout(ctx, userUseCase.contextTimeOut)

	defer cancel()

	err = userUseCase.userRepository.AddUser(ctx, user)

	return
}

func (userUseCase *userUseCase) GetUsername(ctx context.Context, username string) (user string, err error) {
	return user, err
}

func (userUseCase *userUseCase) GetEmail(ctx context.Context, email string) (mail string, err error) {
	return mail, err
}

func (userUseCase *userUseCase) GetUsers(ctx context.Context) (user []domain.User, err error) {
	return user, err
}
func (userUseCase *userUseCase) DeleteUser(ctx context.Context, user string) (err error) {
	return err
}
