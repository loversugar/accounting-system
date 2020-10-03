package api

import (
	"accounting/service"
	"github.com/gin-gonic/gin"
)

func addBill(ctx *gin.Context) {
	billService := service.NewBillService()
	billService.AddBill()
}
