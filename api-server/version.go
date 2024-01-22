package api_server

import "github.com/gin-gonic/gin"

type Version struct {
	Versions string `json:"versions"`
}

func GetVersion(c *gin.Context) {
	c.JSON(200, Version{
		Versions: "karma v0.1",
	})
}
