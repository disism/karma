package api_server

import (
	"github.com/gin-gonic/gin"
)

func GetNodeInfo(c *gin.Context) {

	c.JSON(200, gin.H{
		"karma": "https://karma.disism.com",
	})
}
