package domain

type Comment struct {
	ID        string `gorm:"primaryKey;type:VARCHAR(50)" json:"id"`
	UserID    string `gorm:"type:VARCHAR(50);not null" json:"userId"`
	User      User   `gorm:"foreignKey:UserID;constraint:opUpdate:CASCADE,onDelete:CASCADE"`
	PhotoID   string `gorm:"type:VARCHAR(50);not null" json:"photoId"`
	Photo     Photo  `gorm:"foreignKey:PhotoID;constraint:opUpdate:CASCADE,onDelete:CASCADE"`
	Message   string `gorm:"not null" valid:"required" json:"message"`
	CreatedAt string `gorm:"type:timestamp;not null" json:"createdAt"`
	UpdatedAt string `gorm:"type:timestamp;not null" json:"updatedAt"`
}

type CommentRepository interface {
	AddComment()
	GetComments()
	UpdateComment()
	DeleteComment()
}
