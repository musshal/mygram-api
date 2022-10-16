package domain

import "context"

type Photo struct {
	ID        string    `gorm:"primaryKey;type:VARCHAR(50)" json:"id"`
	Title     string    `gorm:"type:VARCHAR(50);not null" valid:"required" form:"title" json:"title"`
	Caption   string    `form:"caption" json:"caption"`
	PhotoUrl  string    `gorm:"not null" valid:"required" form:"photoUrl" json:"photoUrl"`
	UserID    string    `gorm:"type:VARCHAR(50);not null" json:"userId"`
	User      User      `gorm:"foreignKey:UserID;constraint:onUpdate:CASCADE,onDelete:CASCADE"`
	CreatedAt string    `gorm:"type:timestamp;not null" json:"createdAt"`
	UpdatedAt string    `gorm:"type:timestamp;not null" json:"updatedAt"`
	Comments  []Comment `json:"comments"`
}

type PhotoUseCase interface {
	AddPhotoUseCase(context.Context, *User) error
	GetPhotosUseCase(ctx context.Context) ([]Photo, error)
	UpdatePhotoUseCase(ctx context.Context, photo *Photo) error
	DeletePhotoUseCase(ctx context.Context, id string) error
}

type PhotoRepository interface {
	AddPhoto(ctx context.Context, photo *Photo) error
	GetPhotos(ctx context.Context) (photos []Photo, err error)
	UpdatePhoto(ctx context.Context, photo *Photo) error
	DeletePhoto(ctx context.Context, id string) error
}
