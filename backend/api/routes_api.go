package api

import (
	"accounting/service"
	"github.com/gin-gonic/gin"
)

var userService service.IUserService
var categoryService service.ICategoryService
var billService service.IBillService

func Routes(router *gin.Engine)  {
	userService = service.NewUserService()
	categoryService = service.NewCategoryService()
	billService = service.NewBillService()

	v1 := router.Group("accounting-system")
	userApi := v1.Group("user")
	{
		userApi.GET("login", login)
	}
	billApi := v1.Group("bill")
	{
		billApi.POST("addBill", addBill)
		billApi.GET("getBillByMonth", getBillByMonth)
	}
	categoryApi := v1.Group("category")
	{
		categoryApi.GET("getCategories", getCategories)
	}
}
