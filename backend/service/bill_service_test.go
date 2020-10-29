package service

import (
	"testing"
)

func TestBillService_GetBillByDate(t *testing.T) {
	billService := NewBillService()
	billMap := billService.GetBillByDate(1, "2020-11-01", "2020-11-30")
	for key, value := range billMap {
		t.Log(key)
		t.Log(value)
	}
}