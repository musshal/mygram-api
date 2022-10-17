package test

import (
	"mygram-api/domain"
	"mygram-api/domain/mocks"
	"mygram-api/user/usecase"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"golang.org/x/net/context"
)

func TestAddUser(test *testing.T) {
	mockUserRepository := new(mocks.UserRepository)

	mockUser := domain.User{
		ID:       "user-1",
		Username: "musshal",
		Email:    "musshal@mail.com",
		Password: "secret",
		Age:      22,
	}

	test.Run("success", func(test *testing.T) {
		tempMockUser := mockUser
		tempMockUser.ID = "user-1"

		mockUserRepository.On("GetUsername", mock.Anything, mock.AnythingOfType("string")).Return(string, domain.ErrNotFound).Once()
		mockUserRepository.On("GetEmail", mock.Anything, mock.AnythingOfType("string")).Return(string, domain.ErrNotFound).Once()
		mockUserRepository.On("AddUser", mock.Anything, mock.AnythingOfType("*domain.User")).Return(nil).Once()

		mockUserRepository := new(mocks.UserRepository)

		ucase := usecase.NewUser(mockUserRepository, time.Second*2)

		err := ucase.AddUser(context.TODO(), &tempMockUser)

		assert.NotError(test, err)
		assert.Equal(test, mockUser.ID, tempMockUser.ID)
		assert.Equal(test, mockUser.Username, tempMockUser.Username)
		assert.Equal(test, mockUser.Email, tempMockUser.Email)
		assert.Equal(test, mockUser.Password, tempMockUser.Password)
		assert.Equal(test, mockUser.Age, tempMockUser.Age)

		mockUserRepository.AssertExpectations(test)
	})

	test.Run("fail", func(test *testing.T) {

	})
}
