package api

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func login(context *gin.Context)  {
	code := context.Query("code")
	logrus.Info(code)


	context.Writer.Write(body)
}

func register(context *gin.Context)  {

}
