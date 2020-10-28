package repository

import (
	"testing"
)

func TestBillRepository_InsertBill(t *testing.T) {

}

func TestBillRepository_SelectBillsByDate(t *testing.T) {
	repository := NewBillRepository()
	bills := repository.SelectBillsByDate(1, "2020-11-01", "2020-11-29")
	for _, bill := range bills {
		t.Log(bill.SelectedTime)
	}
}