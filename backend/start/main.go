package main

import "github.com/gin-gonic/gin"

func main()  {
	router := gin.Default()

	router.GET("/ping", func(context *gin.Context) {
		context.Writer.Write([]byte("pong"))
	})

	router.Run(":8888")

}
