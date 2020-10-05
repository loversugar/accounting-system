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
	err := billService.AddBill(&bill)
	if err != nil {
		ctx.JSON(err.Code, gin.H{
			"code": err.Code,
			"message": err.Message,
			"data": err.Data,
		})
	} else {
		ctx.JSON(200, gin.H{
			"code": 200,
			"message": nil,
			"data": nil,
		})
	}
}
