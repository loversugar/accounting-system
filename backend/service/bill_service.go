package service

import "accounting/repository"

type IBillService interface {
	AddBill()
}

func NewBillService() IBillService {
	return &BillService{billRepository:repository.NewBillRepository()}
}

type BillService struct {
	billRepository repository.IBillRepository
}

func (b BillService) AddBill() {

}

