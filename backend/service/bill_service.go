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
	GetBillByDate(userId int, startDate, endDate string) map[string][]*datamodals.BillCategoryResult
}

func NewBillService() IBillService {
	return &BillService{billRepository:repository.NewBillRepository(), categoryRepository:repository.NewCategoryRepository()}
}

type BillService struct {
	billRepository repository.IBillRepository
	categoryRepository repository.ICategoryRepository
}

func (b BillService) GetBillByDate(userId int, startDate, endDate string) map[string][]*datamodals.BillCategoryResult {
	var returnMap = make(map[string][]*datamodals.BillCategoryResult)
	results := b.billRepository.SelectBillsByDate(userId, startDate, endDate)

	for _, result := range results {
		if returnMap[result.SelectedTime] == nil {
			var sliceA []*datamodals.BillCategoryResult
			returnMap[result.SelectedTime] = append(sliceA, result)
		} else {
			returnMap[result.SelectedTime] = append(returnMap[result.SelectedTime], result)
		}
	}

	return returnMap
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

