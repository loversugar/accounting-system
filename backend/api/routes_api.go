package api

import "github.com/gin-gonic/gin"

func Routes(router *gin.Engine)  {
	v1 := router.Group("accounting-system")
	userApi := v1.Group("user")
	{
		userApi.GET("login", login)
	}
	billApi := v1.Group("bill")
	{
		billApi.POST("addBill", addBill)
	}
}
