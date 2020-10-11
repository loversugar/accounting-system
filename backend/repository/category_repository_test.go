package repository

import (
	"fmt"
	"testing"
)

func TestCategoryRepository_InsertCategory(t *testing.T) {
	categoryRepo := NewCategoryRepository()
	categories := categoryRepo.GetCategoriesByUserId(0);
	for _, category := range categories {
		fmt.Println(category.CategoryName)
	}
}
