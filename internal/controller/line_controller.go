package controller

import (
	"net/http"

	"bus-api/internal/service"
	"github.com/gin-gonic/gin"
)

func GetLines(c *gin.Context) {
	lines := service.GetAllLines()
	c.JSON(http.StatusOK, lines)
}
