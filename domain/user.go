package domain

type User struct {
	ID           string         `gorm:"primaryKey;type:VARCHAR(50)" json:"id"`
	Username     string         `gorm:"type:VARCHAR(50);unique;not null" valid:"required" json:"username"`
	Email        string         `gorm:"type:VARCHAR(50);unique;not null" valid:"required" json:"email"`
	Password     string         `gorm:"not null" valid:"required" json:"password"`
	Age          uint           `gorm:"not null" valid:"required" json:"age"`
	CreatedAt    string         `gorm:"type:timestamp;not null" json:"createdAt"`
	UpdatedAt    string         `gorm:"type:timestamp;not null" json:"updatedAt"`
	Photos       []Photo        `json:"photos"`
	SocialMedias []SocialMedias `json:"socialMedias"`
}

type UserUseCase interface {
	UserRegister()
	UserLogin()
}

type UserRepository interface {
	AddUser()
	GetUsers()
	DeleteUser()
}
