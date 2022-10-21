package usecase_test

import (
	"context"
	"mygram-api/domain"
	"mygram-api/domain/mocks"
	"mygram-api/helpers"
	"testing"
	"time"

	userUseCase "mygram-api/user/usecase"

	"github.com/asaskevich/govalidator"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestRegister(t *testing.T) {
	mockRegisteredUser := domain.User{
		ID:       "user-123",
		Age:      8,
		Email:    "johndoe@example.com",
		Password: "secret",
		Username: "johndoe",
	}

	mockUserRepository := new(mocks.UserRepository)
	userUseCase := userUseCase.NewUserUseCase(mockUserRepository)

	t.Run("register user correctly", func(t *testing.T) {
		tempMockRegisterUser := domain.User{
			Age:      8,
			Email:    "johndoe@example.com",
			Password: "secret",
			Username: "johndoe",
		}

		tempMockRegisterUser.ID = "user-123"

		mockUserRepository.On("Register", mock.Anything, mock.AnythingOfType("*domain.User")).Return(nil).Once()

		err := userUseCase.Register(context.Background(), &tempMockRegisterUser)

		assert.NoError(t, err)

		_, err = govalidator.ValidateStruct(tempMockRegisterUser)

		assert.NoError(t, err)
		assert.Equal(t, tempMockRegisterUser.ID, mockRegisteredUser.ID)
		assert.Equal(t, tempMockRegisterUser.Age, mockRegisteredUser.Age)
		assert.Equal(t, tempMockRegisterUser.Email, mockRegisteredUser.Email)
		assert.Equal(t, tempMockRegisterUser.Password, mockRegisteredUser.Password)
		assert.Equal(t, tempMockRegisterUser.Username, mockRegisteredUser.Username)
		mockUserRepository.AssertExpectations(t)
	})

	t.Run("register user with not contain needed property", func(t *testing.T) {
		tempMockRegisterUser := domain.User{
			Age:   8,
			Email: "johndoe@example.com",
		}

		tempMockRegisterUser.ID = "user-123"

		mockUserRepository.On("Register", mock.Anything, mock.AnythingOfType("*domain.User")).Return(nil).Once()

		err := userUseCase.Register(context.Background(), &tempMockRegisterUser)

		assert.NoError(t, err)

		_, err = govalidator.ValidateStruct(tempMockRegisterUser)

		assert.Error(t, err)

		mockUserRepository.AssertExpectations(t)
	})

	t.Run("register user with age under limit number", func(t *testing.T) {
		tempMockRegisterUser := domain.User{
			Age:      7,
			Email:    "johndoe@example.com",
			Password: "secret",
			Username: "johndoe",
		}

		tempMockRegisterUser.ID = "user-123"

		mockUserRepository.On("Register", mock.Anything, mock.AnythingOfType("*domain.User")).Return(nil).Once()

		err := userUseCase.Register(context.Background(), &tempMockRegisterUser)

		assert.NoError(t, err)

		_, err = govalidator.ValidateStruct(tempMockRegisterUser)

		assert.Error(t, err)

		mockUserRepository.AssertExpectations(t)
	})

	t.Run("register user with invalid email format", func(t *testing.T) {
		tempMockRegisterUser := domain.User{
			Age:      8,
			Email:    "johndoe",
			Password: "secret",
			Username: "johndoe",
		}

		tempMockRegisterUser.ID = "user-123"

		mockUserRepository.On("Register", mock.Anything, mock.AnythingOfType("*domain.User")).Return(nil).Once()

		err := userUseCase.Register(context.Background(), &tempMockRegisterUser)

		assert.NoError(t, err)

		_, err = govalidator.ValidateStruct(tempMockRegisterUser)

		assert.Error(t, err)

		mockUserRepository.AssertExpectations(t)
	})

	t.Run("register user with password under limit character", func(t *testing.T) {
		tempMockRegisterUser := domain.User{
			Age:      8,
			Email:    "johndoe@example.com",
			Password: "scrt",
			Username: "johndoe",
		}

		tempMockRegisterUser.ID = "user-123"

		mockUserRepository.On("Register", mock.Anything, mock.AnythingOfType("*domain.User")).Return(nil).Once()

		err := userUseCase.Register(context.Background(), &tempMockRegisterUser)

		assert.NoError(t, err)

		_, err = govalidator.ValidateStruct(tempMockRegisterUser)

		assert.Error(t, err)

		mockUserRepository.AssertExpectations(t)
	})
}

func TestLogin(t *testing.T) {
	mockRegisteredUser := domain.User{
		ID:       "user-123",
		Age:      8,
		Email:    "johndoe@example.com",
		Password: helpers.Hash("secret"),
		Username: "johndoe",
	}

	mockUserRepository := new(mocks.UserRepository)
	userUseCase := userUseCase.NewUserUseCase(mockUserRepository)

	t.Run("login user correctly", func(t *testing.T) {
		tempMockLoginUser := domain.User{
			Email:    "johndoe@example.com",
			Password: "secret",
		}

		mockUserRepository.On("Login", mock.Anything, mock.AnythingOfType("*domain.User")).Return(nil).Once()

		err := userUseCase.Login(context.Background(), &tempMockLoginUser)

		assert.NoError(t, err)

		assert.Equal(t, tempMockLoginUser.Email, mockRegisteredUser.Email)

		isValid := helpers.Compare([]byte(mockRegisteredUser.Password), []byte(tempMockLoginUser.Password))

		assert.True(t, isValid)
		mockUserRepository.AssertExpectations(t)
	})

	t.Run("login user with not registered email", func(t *testing.T) {
		tempMockLoginUser := domain.User{
			Email:    "lorem@example.com",
			Password: "secret",
		}

		mockUserRepository.On("Login", mock.Anything, mock.AnythingOfType("*domain.User")).Return(nil).Once()

		err := userUseCase.Login(context.Background(), &tempMockLoginUser)

		assert.NoError(t, err)

		assert.NotEqual(t, tempMockLoginUser.Email, mockRegisteredUser.Email)

		isValid := helpers.Compare([]byte(mockRegisteredUser.Password), []byte(tempMockLoginUser.Password))

		assert.True(t, isValid)
		mockUserRepository.AssertExpectations(t)
	})

	t.Run("login user with invalid password", func(t *testing.T) {
		tempMockLoginUser := domain.User{
			Email:    "johndoe@example.com",
			Password: "scrt",
		}

		mockUserRepository.On("Login", mock.Anything, mock.AnythingOfType("*domain.User")).Return(nil).Once()

		err := userUseCase.Login(context.Background(), &tempMockLoginUser)

		assert.NoError(t, err)

		assert.Equal(t, tempMockLoginUser.Email, mockRegisteredUser.Email)

		isValid := helpers.Compare([]byte(mockRegisteredUser.Password), []byte(tempMockLoginUser.Password))

		assert.False(t, isValid)
		mockUserRepository.AssertExpectations(t)
	})
}

func TestUpdate(t *testing.T) {
	date := time.Now()
	mockUpdatedUser := domain.User{
		ID:        "user-123",
		Email:     "newjohndoe@example.com",
		Username:  "newjohndoe",
		Age:       8,
		UpdatedAt: &date,
	}

	mockUserRepository := new(mocks.UserRepository)
	userUseCase := userUseCase.NewUserUseCase(mockUserRepository)

	t.Run("updated user correctly", func(t *testing.T) {
		tempMockUpdateUser := domain.User{
			Email:    "newjohndoe@example.com",
			Username: "newjohndoe",
		}

		mockUserRepository.On("Update", mock.Anything, mock.AnythingOfType("domain.User")).Return(mockUpdatedUser, nil).Once()

		user, err := userUseCase.Update(context.Background(), tempMockUpdateUser)

		tempMockUpdatedUser := domain.User{
			ID:        "user-123",
			Email:     tempMockUpdateUser.Email,
			Username:  tempMockUpdateUser.Username,
			Age:       8,
			UpdatedAt: &date,
		}

		assert.NoError(t, err)
		assert.Equal(t, tempMockUpdatedUser, user)
		assert.NoError(t, err)
		assert.Equal(t, tempMockUpdateUser.Email, mockUpdatedUser.Email)
		assert.Equal(t, tempMockUpdatedUser.Username, mockUpdatedUser.Username)
		mockUserRepository.AssertExpectations(t)
	})
}
