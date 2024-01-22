package api_server

import (
	"github.com/disism/karma/internal/database"
	"github.com/disism/karma/internal/server"
	"github.com/gin-gonic/gin"
	"golang.org/x/exp/slog"
)

func GetDevices(c *gin.Context) {
	client, err := database.New(c)
	if err != nil {
		slog.Error("get devices new database error: ", err.Error())
		server.ErrorInternalServerError(c)
	}

	if err := server.NewServer(c, client.Client).GetDevices(); err != nil {
		slog.Error("get devices error: ", err.Error())
		server.ErrorInternalServerError(c)
	}
}

func DeleteDevice(c *gin.Context) {
	client, err := database.New(c)
	if err != nil {
		slog.Error("delete device new database error: ", err.Error())
		server.ErrorInternalServerError(c)
	}

	if err := server.NewServer(c, client.Client).DeleteDevice(); err != nil {
		slog.Error("delete device error: ", err.Error())
		server.ErrorInternalServerError(c)
	}
}
