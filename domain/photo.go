package domain

type Photo struct {
	ID        string    `gorm:"primaryKey;type:VARCHAR(50)" json:"id"`
	Title     string    `gorm:"type:VARCHAR(50);not null" valid:"required" json:"title"`
	Caption   string    `json:"caption"`
	PhotoUrl  string    `gorm:"not null" valid:"required" json:"photoUrl"`
	UserID    string    `gorm:"type:VARCHAR(50);not null" json:"userId"`
	User      User      `gorm:"foreignKey:UserID;constraint:onUpdate:CASCADE,onDelete:CASCADE"`
	CreatedAt string    `gorm:"type:timestamp;not null" json:"createdAt"`
	UpdatedAt string    `gorm:"type:timestamp;not null" json:"updatedAt"`
	Comments  []Comment `json:"comments"`
}

type PhotoRepository interface {
	AddPhoto()
	GetPhotos()
	UpdatePhoto()
	DeletePhoto()
}
