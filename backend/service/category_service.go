package service

import (
	"accounting/datamodals"
	"accounting/repository"
)

type ICategoryService interface {
	getCategories(userId int) []*datamodals.Category
}

func NewCategoryService() ICategoryService {
	return &CategoryService{repository:repository.NewCategoryRepository()}
}

type CategoryService struct {
	repository repository.ICategoryRepository
}

func (c CategoryService) getCategories(userId int) []*datamodals.Category {

}

