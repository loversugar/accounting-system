package service

import (
	"accounting/datamodals"
	"accounting/exception"
	"accounting/repository"
	"net/http"
)

type IBillService interface {
	AddBill(bill *datamodals.Bill, categoryId int) *exception.ApiException
	GetBillByMonth(userId int, month int)
}

func NewBillService() IBillService {
	return &BillService{billRepository:repository.NewBillRepository(), categoryRepository:repository.NewCategoryRepository()}
}

type BillService struct {
	billRepository repository.IBillRepository
	categoryRepository repository.ICategoryRepository
}

func (b BillService) GetBillByMonth(userId int, month int) {
	panic("implement me")
}

func (b BillService) AddBill(bill *datamodals.Bill, categoryId int) *exception.ApiException {
	category := b.categoryRepository.SelectCategoryById(categoryId)

	if category == nil {
		return &exception.ApiException{Code:http.StatusBadRequest, Message:"category is not existed", Data:nil}
	}

	billCategory := &datamodals.BillCategory{
		CategoryID:categoryId, CategoryName:category.CategoryName, CategoryUrl:category.CategoryUrl}

	b.billRepository.InsertBill(bill, billCategory)

	return nil
}

