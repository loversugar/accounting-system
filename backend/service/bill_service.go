package service

import (
	"accounting/api"
	"accounting/datamodals"
	"accounting/repository"
	"time"
)

type IBillService interface {
	AddBill(bill *api.Bill)
}

func NewBillService() IBillService {
	return &BillService{billRepository:repository.NewBillRepository(), categoryRepository:repository.NewCategoryRepository()}
}

type BillService struct {
	billRepository repository.IBillRepository
	categoryRepository repository.ICategoryRepository
}

func (b BillService) AddBill(bill *api.Bill) {
	innerBill := &datamodals.Bill{Consumption:bill.Consumption, UserId:bill.UserId, CreateTime:time.Now()}

	category := b.categoryRepository.SelectCategoryById(bill.CategoryId)

	billCategory := &datamodals.BillCategory{
		CategoryID:bill.CategoryId, CategoryName:category.CategoryName, CategoryUrl:category.CategoryUrl}

	b.billRepository.InsertBill(innerBill, billCategory)
}

