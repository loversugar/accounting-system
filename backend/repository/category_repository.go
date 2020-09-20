package repository

import (
	"accounting/datamodals"
	"accounting/util"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type ICategoryRepository interface {
	InsertCategory(category *datamodals.Category) error
}

func NewCategoryRepository() ICategoryRepository {
	db, err := util.Connection()
	if err != nil {
		logrus.Error(err)
	}
	return &CategoryRepository{db:db}
}

type CategoryRepository struct {
	db *gorm.DB
}

func (c CategoryRepository) InsertCategory(category *datamodals.Category) error {
	return c.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(category).Error; err != nil {
			return err
		}
		return nil
	})
}
