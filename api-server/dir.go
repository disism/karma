package api_server

import (
	"github.com/disism/karma/internal/database"
	"github.com/disism/karma/internal/server"
	"github.com/gin-gonic/gin"
	"golang.org/x/exp/slog"
)

func MKDir(c *gin.Context) {
	client, err := database.New(c)
	if err != nil {
		slog.Error("make dir new database error: ", err.Error())
		server.ErrorInternalServerError(c)
	}

	if err := server.NewServer(c, client.Client).MKDir(); err != nil {
		slog.Error("make dir error: ", err.Error())
		server.ErrorInternalServerError(c)
	}
}

func ListDir(c *gin.Context) {
	client, err := database.New(c)
	if err != nil {
		slog.Error("list dir new database error: ", err.Error())
		server.ErrorInternalServerError(c)
	}

	if err := server.NewServer(c, client.Client).ListDir(); err != nil {
		slog.Error("list dir error: ", err.Error())
		server.ErrorInternalServerError(c)
	}
}

func ListDirs(c *gin.Context) {
	client, err := database.New(c)
	if err != nil {
		slog.Error("list dirs new database error: ", err.Error())
		server.ErrorInternalServerError(c)
	}

	if err := server.NewServer(c, client.Client).ListDirs(); err != nil {
		slog.Error("list dirs error: ", err.Error())
		server.ErrorInternalServerError(c)
	}
}

func RenameDir(c *gin.Context) {
	client, err := database.New(c)
	if err != nil {
		slog.Error("rename dir new database error: ", err.Error())
		server.ErrorInternalServerError(c)
	}

	if err := server.NewServer(c, client.Client).RenameDir(); err != nil {
		slog.Error("rename dir error: ", err.Error())
		server.ErrorInternalServerError(c)
	}
}

func MVDir(c *gin.Context) {
	client, err := database.New(c)
	if err != nil {
		slog.Error("mv dir new database error: ", err.Error())
		server.ErrorInternalServerError(c)
	}

	if err := server.NewServer(c, client.Client).MVDir(); err != nil {
		slog.Error("mv dir error: ", err.Error())
		server.ErrorInternalServerError(c)
	}
}

func RMDir(c *gin.Context) {
	client, err := database.New(c)
	if err != nil {
		slog.Error("rm dir new database error: ", err.Error())
		server.ErrorInternalServerError(c)
	}

	if err := server.NewServer(c, client.Client).RMDir(); err != nil {
		slog.Error("rm dir error: ", err.Error())
		server.ErrorInternalServerError(c)
	}
}
