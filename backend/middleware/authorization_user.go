package middleware

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type AA struct {
	name string
	password string

}

func Authorize() gin.HandlerFunc {
	return func(c *gin.Context) {
		aa := &AA{name:"test", password:"test"}
		log.Info("cc", aa);
		c.Next()
	}
}
