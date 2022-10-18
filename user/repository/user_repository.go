package repository

import (
	"context"
	"errors"
	"mygram-api/domain"
	"mygram-api/helpers"
	"time"

	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db}
}

func (userRepository *userRepository) Register(ctx context.Context, user *domain.User) (err error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)

	defer cancel()

	err = userRepository.db.Debug().WithContext(ctx).Create(&user).Error

	if err != nil {
		return err
	}

	return
}

func (userRepository *userRepository) Login(ctx context.Context, user *domain.User) (err error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)

	defer cancel()

	password := user.Password

	if err = userRepository.db.Debug().WithContext(ctx).Where("email = ?", user.Email).Take(&user).Error; err != nil {
		return err
	}

	if isValid := helpers.Compare([]byte(user.Password), []byte(password)); !isValid {
		return errors.New("the credential you entered are wrong")
	}

	return
}

func (userRepository *userRepository) Update(ctx context.Context, user domain.User, id uint) (u domain.User, err error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)

	defer cancel()

	u = domain.User{}

	if err = userRepository.db.Debug().First(&u, id).Error; err != nil {
		return u, err
	}

	if err = userRepository.db.Debug().WithContext(ctx).Model(&u).Where("id = ?", id).Updates(user).Error; err != nil {
		return u, err
	}

	return u, nil
}

func (userRepository *userRepository) Delete(ctx context.Context, id uint) (err error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)

	defer cancel()

	if err = userRepository.db.Debug().WithContext(ctx).Delete(&domain.User{}, id).Error; err != nil {
		return err
	}

	return
}
