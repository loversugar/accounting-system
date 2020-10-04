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

func (c CategoryRepository) getCategoriesByUserId(userId int) *[]datamodals.Category   {
	var categories []datamodals.Category
	c.db.Where("user_id = ? or is_private = 0", userId).Find(&categories)
	return &categories
}

