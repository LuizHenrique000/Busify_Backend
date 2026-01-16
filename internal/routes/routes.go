package routes

import (
	"bus-api/internal/controller"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	api := r.Group("/api")
	{
		api.GET("/lines", controller.GetLines)
	}
}
