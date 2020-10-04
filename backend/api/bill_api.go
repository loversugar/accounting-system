package api

import (
	"accounting/service"
	"github.com/gin-gonic/gin"
)

type Bill struct {
	UserId string `json:"userId" form:"user_id"`
	Consumption float32 `json:"consumption" form:"consumption"`
	CategoryId int `json:"categoryId" form:"categoryId`
}

func addBill(ctx *gin.Context) {
	var bill Bill
	ctx.BindJSON(&bill)
	billService := service.NewBillService()
	billService.AddBill(&bill)
}
