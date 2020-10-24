package api

import (
	"accounting/datamodals"
	"github.com/gin-gonic/gin"
	"strconv"
	"strings"
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

func getBillByDate(ctx *gin.Context) {
	userId, err := strconv.Atoi(ctx.Query("userId"))
	if err != nil {

	}

	startDate := strings.Trim(ctx.Query("startDate"), " ")
	endDate := strings.Trim(ctx.Query("endDate"), " ")

	if len(startDate) == 0 {

	}

	if len(endDate) == 0 {

	}

	billService.GetBillByDate(userId, startDate, endDate);
}
