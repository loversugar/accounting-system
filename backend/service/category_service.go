package service

import (
	"accounting/datamodals"
	"accounting/repository"
)

type ICategoryService interface {
	GetCategories(userId int) []*datamodals.Category
}

func NewCategoryService() ICategoryService {
	return &CategoryService{repository:repository.NewCategoryRepository()}
}

type CategoryService struct {
	repository repository.ICategoryRepository
}

func (c CategoryService) GetCategories(userId int) []*datamodals.Category {
	return c.repository.GetCategoriesByUserId(userId)
}
