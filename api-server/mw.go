package api_server

import (
	"github.com/disism/karma/internal/database"
	"github.com/disism/karma/internal/server"
	"github.com/gin-gonic/gin"
	"golang.org/x/exp/slog"
)

func AuthMW() gin.HandlerFunc {
	return func(c *gin.Context) {
		client, err := database.New(c)
		if err != nil {
			slog.Error("auth middleware new database error: ", err.Error())
			server.ErrorInternalServerError(c)
			c.Abort()
		}

		if err := server.NewServer(c, client.Client).ValidateJWT(); err != nil {
			server.ErrorUnauthorized(c, err.Error())
			c.Abort()
		}

		c.Next()
	}
}

func CORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, PATCH, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Authorization")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
			return
		}

		c.Next()
	}
}
