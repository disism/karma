package api_server

import (
	"github.com/disism/karma/internal/database"
	"github.com/disism/karma/internal/server"
	"github.com/gin-gonic/gin"
	"golang.org/x/exp/slog"
)

func CreateUser(c *gin.Context) {
	client, err := database.New(c)
	if err != nil {
		slog.Error("create user new database error: ", err.Error())
		server.ErrorInternalServerError(c)
	}

	if err := server.NewServer(c, client.Client).CreateUser(); err != nil {
		slog.Error("create user error: ", err.Error())
		server.ErrorInternalServerError(c)
	}
}
