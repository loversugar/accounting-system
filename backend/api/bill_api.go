package api

import (
	"accounting/datamodals"
	"accounting/service"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

type Bill struct {
	UserId string `json:"userId" form:"userId"`
	Consumption float32 `json:"consumption" form:"consumption"`
	Note string `json:"note" form:"note"`
	SelectedTime time.Time `json:"selectedTime" form:"selectedTime"`
	CategoryId int `json:"categoryId" form:"categoryId`
}

func addBill(ctx *gin.Context) {
	var bill Bill
	ctx.BindJSON(&bill)

	innerBill := &datamodals.Bill{
		Consumption:bill.Consumption, Note:bill.Note,
		UserId:bill.UserId, SelectedTime:bill.SelectedTime, CreateTime:time.Now()}

	err := billService.AddBill(innerBill, bill.CategoryId)
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

func getBillByMonth(ctx *gin.Context) {
	userId, err := strconv.Atoi(ctx.Query("userId"))
	if err != nil {

	}
	month, err := strconv.Atoi(ctx.Query("month"))
	if err != nil {

	}

	billService := service.NewBillService()
	billService.GetBillByMonth(userId, month)
}
