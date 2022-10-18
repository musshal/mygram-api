package domain

import (
	"context"
	"time"
)

type Photo struct {
	ID        string     `gorm:"primaryKey;type:VARCHAR(50)" json:"id" binding:"required"`
	Title     string     `gorm:"type:VARCHAR(50);not null" valid:"required" form:"title" json:"title" binding:"required"`
	Caption   string     `form:"caption" json:"caption"`
	PhotoUrl  string     `gorm:"not null" valid:"required" form:"photoUrl" json:"photoUrl" binding:"required"`
	UserID    uint       `gorm:"not null" json:"userId"`
	User      User       `gorm:"foreignKey:UserID;constraint:onUpdate:CASCADE,onDelete:CASCADE"`
	CreatedAt *time.Time `gorm:"not null;autoCreateTime" json:"createdAt"`
	UpdatedAt *time.Time `gorm:"not null;autoCreateTime" json:"updatedAt"`
	Comments  []Comment  `json:"comments"`
}

type PhotoUseCase interface {
	AddPhoto(context.Context, *User) error
	GetPhotos(ctx context.Context) ([]Photo, error)
	UpdatePhoto(ctx context.Context, photo *Photo) error
	DeletePhoto(ctx context.Context, id string) error
}

type PhotoRepository interface {
	AddPhoto(ctx context.Context, photo *Photo) error
	GetPhotos(ctx context.Context) (photos []Photo, err error)
	UpdatePhoto(ctx context.Context, photo *Photo) error
	DeletePhoto(ctx context.Context, id string) error
}
