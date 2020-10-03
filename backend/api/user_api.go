package api

import (
	"accounting/service"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func login(context *gin.Context)  {
	code := context.Query("code")
	nickName := context.Query("nickName")
	logrus.Info(code)

	userService := service.NewUserService()
	body, err := userService.Login(code, nickName)
	if err != nil {
		logrus.Error(err)
	}

	context.Writer.Write([]byte(body))
}
