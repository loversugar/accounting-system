package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func Authorize() gin.HandlerFunc {
	return func(c *gin.Context) {
		logrus.Info("invoked")
		c.Next()
	}
}
