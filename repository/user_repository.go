package repository

import (
	"accounting/datamodals"
	"accounting/util"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type IUserRepository interface {
	InsertUser(user *datamodals.User) error
}

func NewUserRepository() IUserRepository {
	db, err := util.Connection()
	if err != nil {
		logrus.Error(err)
	}
	return &UserRepository{db:db}
}

type UserRepository struct {
	db *gorm.DB
}

func (u UserRepository) InsertUser(user *datamodals.User) error {
	if err := u.db.Create(user).Error; err != nil {
		return err
	}
	return nil
}

