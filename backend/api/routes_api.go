package api

import "github.com/gin-gonic/gin"

func Routes(router *gin.Engine)  {
	v1 := router.Group("accounting-system")
	userApi := v1.Group("user")
	{
		userApi.GET("login", login)
		userApi.POST("register", register)
	}
}
