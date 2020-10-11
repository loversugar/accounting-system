package repository

import (
	"accounting/datamodals"
	"accounting/util"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type ICategoryRepository interface {
	InsertCategory(category *datamodals.Category) error

	SelectCategoryById(id int) *datamodals.Category

	GetCategoriesByUserId(userId int) []*datamodals.Category
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

func (c CategoryRepository) GetCategoriesByUserId(userId int) []*datamodals.Category {
	var categories []*datamodals.Category

	if userId == 0 {
		c.db.Debug().Where("user_id is null and deleted = 0").Find(&categories)
	} else {
		c.db.Where("(user_id = ? and deleted = 0) or (user_id = null and deleted = 0)", userId).Find(&categories)
	}

	return categories
}

func (c CategoryRepository) SelectCategoryById(id int) *datamodals.Category {
	var category datamodals.Category
	c.db.Where("id = ?", id).First(&category)
	return &category
}

func (c CategoryRepository) InsertCategory(category *datamodals.Category) error {
	return c.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(category).Error; err != nil {
			return err
		}
		return nil
	})
}
