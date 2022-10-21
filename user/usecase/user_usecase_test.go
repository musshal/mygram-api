package usecase_test

import (
	"context"
	"mygram-api/domain"
	"mygram-api/domain/mocks"
	"testing"

	userUseCase "mygram-api/user/usecase"

	"github.com/asaskevich/govalidator"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestRegister(t *testing.T) {
	mockUserRepository := new(mocks.UserRepository)

	t.Run("success", func(t *testing.T) {
		tempMockRegisterUser := domain.User{
			Age:      8,
			Email:    "johndoe@example.com",
			Password: "secret",
			Username: "johndoe",
		}

		tempMockRegisterUser.ID = "user-123"

		mockUserRepository.On("Register", mock.Anything, mock.AnythingOfType("*domain.User")).Return(nil).Once()

		userUseCase := userUseCase.NewUserUseCase(mockUserRepository)
		err := userUseCase.Register(context.Background(), &tempMockRegisterUser)

		assert.NoError(t, err)

		_, err = govalidator.ValidateStruct(tempMockRegisterUser)

		assert.NoError(t, err)

		mockUserRepository.AssertExpectations(t)
	})

	t.Run("not contain needed property", func(t *testing.T) {
		tempMockRegisterUser := domain.User{
			Age:   8,
			Email: "johndoe@example.com",
		}

		tempMockRegisterUser.ID = "user-123"

		mockUserRepository.On("Register", mock.Anything, mock.AnythingOfType("*domain.User")).Return(nil).Once()

		userUseCase := userUseCase.NewUserUseCase(mockUserRepository)
		err := userUseCase.Register(context.Background(), &tempMockRegisterUser)

		assert.NoError(t, err)

		_, err = govalidator.ValidateStruct(tempMockRegisterUser)

		assert.Error(t, err)

		mockUserRepository.AssertExpectations(t)
	})

	t.Run("age under limit number", func(t *testing.T) {
		tempMockRegisterUser := domain.User{
			Age:      7,
			Email:    "johndoe@example.com",
			Password: "secret",
			Username: "johndoe",
		}

		tempMockRegisterUser.ID = "user-123"

		mockUserRepository.On("Register", mock.Anything, mock.AnythingOfType("*domain.User")).Return(nil).Once()

		userUseCase := userUseCase.NewUserUseCase(mockUserRepository)
		err := userUseCase.Register(context.Background(), &tempMockRegisterUser)

		assert.NoError(t, err)

		_, err = govalidator.ValidateStruct(tempMockRegisterUser)

		assert.Error(t, err)

		mockUserRepository.AssertExpectations(t)
	})

	t.Run("password under limit character", func(t *testing.T) {
		tempMockRegisterUser := domain.User{
			Age:      8,
			Email:    "johndoe@example.com",
			Password: "scrt",
			Username: "johndoe",
		}

		tempMockRegisterUser.ID = "user-123"

		mockUserRepository.On("Register", mock.Anything, mock.AnythingOfType("*domain.User")).Return(nil).Once()

		userUseCase := userUseCase.NewUserUseCase(mockUserRepository)
		err := userUseCase.Register(context.Background(), &tempMockRegisterUser)

		assert.NoError(t, err)

		_, err = govalidator.ValidateStruct(tempMockRegisterUser)

		assert.Error(t, err)

		mockUserRepository.AssertExpectations(t)
	})
}
