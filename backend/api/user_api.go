package api

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func login(context *gin.Context)  {
	code := context.Query("code")
	nickName := context.Query("nickName")
	logrus.Info(code)

	body, err := userService.Login(code, nickName)
	if err != nil {
		logrus.Error(err)
	}

	context.Writer.Write([]byte(body))
}
