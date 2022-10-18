package domain

import (
	"context"
	"time"
)

type Photo struct {
	ID        uint       `gorm:"primaryKey" json:"id"`
	Title     string     `gorm:"type:VARCHAR(50);not null" valid:"required" form:"title" json:"title" example:"Tittle"`
	Caption   string     `form:"caption" json:"caption"`
	PhotoUrl  string     `gorm:"not null" valid:"required" form:"photoUrl" json:"photoUrl" example:"https://www.example.com/image.jpg"`
	UserID    uint       `gorm:"not null" json:"userId"`
	User      User       `gorm:"foreignKey:UserID;constraint:onUpdate:CASCADE,onDelete:CASCADE"`
	CreatedAt *time.Time `gorm:"not null;autoCreateTime" json:"createdAt,omitempty"`
	UpdatedAt *time.Time `gorm:"not null;autoCreateTime" json:"updatedAt,omitempty"`
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
