package domain

import (
	"context"
	"time"
)

type Comment struct {
	ID        string     `gorm:"primaryKey;type:VARCHAR(50)" json:"id" binding:"required"`
	UserID    string     `gorm:"type:VARCHAR(50);not null" json:"userId"`
	User      User       `gorm:"foreignKey:UserID;constraint:opUpdate:CASCADE,onDelete:CASCADE"`
	PhotoID   string     `gorm:"type:VARCHAR(50);not null" form:"photoId" json:"photoId"`
	Photo     Photo      `gorm:"foreignKey:PhotoID;constraint:opUpdate:CASCADE,onDelete:CASCADE"`
	Message   string     `gorm:"not null" valid:"required" form:"message" json:"message" binding:"required"`
	CreatedAt *time.Time `gorm:"not null;autoCreateTime" json:"createdAt"`
	UpdatedAt *time.Time `gorm:"not null;autoCreateTime" json:"updatedAt"`
}

type CommentUseCase interface {
	AddComment(context.Context, *Comment) error
	GetComments(ctx context.Context) ([]Comment, error)
	UpdateComment(ctx context.Context, comment *Comment) error
	DeleteComment(ctx context.Context, id string) error
}

type CommentRepository interface {
	AddComment(ctx context.Context, comment *Comment) error
	GetComments(ctx context.Context) (comments []Comment, err error)
	UpdateComment(ctx context.Context, comment *Comment) error
	DeleteComment(ctx context.Context, id string) error
}
