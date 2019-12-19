package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/wubingwei/veigar/handler/api"
	"net/http"
)

func HealthCheck(c *gin.Context) {
	c.JSON(
		http.StatusOK, gin.H{
			"code":    0,
			"message": "healthy",
			"error":   "",
		})
}

func InitRouter(engine *gin.Engine) {
	engine.GET("/healthCheck", HealthCheck)

	apiRouter := engine.Group("/api")
	api.InitRouter(apiRouter)
}
