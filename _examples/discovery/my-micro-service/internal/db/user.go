package db

import (
	"github.com/god-jay/gools/_examples/discovery/my-micro-service/internal/model"
	"gorm.io/gorm"
)

type User struct {
	*gorm.DB
}

func (u *User) GetById(id int32) (*model.User, error) {
	var user model.User
	err := u.Where("id", id).Find(&user).Error
	return &user, err
}
