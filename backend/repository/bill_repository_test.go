package repository

import (
	"accounting/datamodals"
	"testing"
	"time"
)

func TestBillRepository_InsertBill(t *testing.T) {
	billRepo := NewBillRepository()
	bill := &datamodals.Bill{OpenId:"test", Username:"zs", Consumption:-9, CreateTime:time.Now()}
	err := billRepo.InsertBill(bill, 1)
	if err != nil {
		t.Error("failed")
		return
	}
	t.Log("success")
}