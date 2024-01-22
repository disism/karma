package server

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Success(c *gin.Context, resp any) {
	c.JSON(http.StatusOK, resp)
}
