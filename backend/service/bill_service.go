package service

import (
	"accounting/api"
	"accounting/datamodals"
	"accounting/exception"
	"accounting/repository"
	"net/http"
	"time"
)

type IBillService interface {
	AddBill(bill *api.Bill) *exception.ApiException
}

func NewBillService() IBillService {
	return &BillService{billRepository:repository.NewBillRepository(), categoryRepository:repository.NewCategoryRepository()}
}

type BillService struct {
	billRepository repository.IBillRepository
	categoryRepository repository.ICategoryRepository
}

func (b BillService) AddBill(bill *api.Bill) *exception.ApiException {
	innerBill := &datamodals.Bill{Consumption:bill.Consumption, UserId:bill.UserId, CreateTime:time.Now()}

	category := b.categoryRepository.SelectCategoryById(bill.CategoryId)

	if category == nil {
		return &exception.ApiException{Code:http.StatusBadRequest, Message:"category is not existed", Data:nil}
	}

	billCategory := &datamodals.BillCategory{
		CategoryID:bill.CategoryId, CategoryName:category.CategoryName, CategoryUrl:category.CategoryUrl}

	b.billRepository.InsertBill(innerBill, billCategory)
}

