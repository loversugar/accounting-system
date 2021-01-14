package main

import (
	"accounting/api"
	"accounting/middleware"
	"github.com/gin-gonic/gin"
)


func main()  {
	router := gin.Default()

	router.Use(middleware.LoggerToFile())
	router.GET("/ping", func(context *gin.Context) {
		context.Writer.Write([]byte("pong"))
	})

	api.Routes(router)


	router.Run(":8888")

}
