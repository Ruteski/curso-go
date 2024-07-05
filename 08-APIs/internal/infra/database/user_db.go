package database

import (
	"apis/internal/entity"
	"gorm.io/gorm"
)

type User struct {
	//DBSql *sql.DB
	DB *gorm.DB
}

func NewUser(db *gorm.DB) *User {
	return &User{db}
}

func (u *User) CreateUser(user *entity.User) error {
	return u.DB.Create(&user).Error
}

func (u *User) FindByEmail(email string) (*entity.User, error) {
	var user entity.User
	if err := u.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}
