package api

import "github.com/gin-gonic/gin"

func getCategories(ctx *gin.Context) {
	ctx.JSON(200, categoryService.GetCategories(0))
}
