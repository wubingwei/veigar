package api

import "github.com/gin-gonic/gin"

func InitRouter(apiRouter *gin.RouterGroup) {
	apiRouter.GET("/admin/:id", Admin)
}
