package main

import (
	"accounting/api"
	"github.com/gin-gonic/gin"
)

func main()  {
	router := gin.Default()

	router.GET("/ping", func(context *gin.Context) {
		context.Writer.Write([]byte("pong"))
	})

	api.Routes(router)

	router.Run(":8888")

}
