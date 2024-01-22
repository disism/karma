package api_server

import (
	"github.com/disism/karma/internal/database"
	"github.com/disism/karma/internal/server"
	"github.com/gin-gonic/gin"
	"golang.org/x/exp/slog"
)

func Login(c *gin.Context) {
	client, err := database.New(c)
	if err != nil {
		slog.Error("login new database error: ", err.Error())
		server.ErrorInternalServerError(c)
		return
	}

	if err := server.NewServer(c, client.Client).Login(); err != nil {
		slog.Error("login error: ", err.Error())
		return
	}

}
