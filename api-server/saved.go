package api_server

import (
	"github.com/disism/karma/internal/database"
	"github.com/disism/karma/internal/server"
	"github.com/gin-gonic/gin"
	"golang.org/x/exp/slog"
)

func CreateSaves(c *gin.Context) {
	client, err := database.New(c)
	if err != nil {
		slog.Error("create saved new database error: ", err.Error())
		server.ErrorInternalServerError(c)
	}

	if err := server.NewServer(c, client.Client).CreateSaves(); err != nil {
		slog.Error("create saved error: ", err.Error())
		server.ErrorInternalServerError(c)
	}

}

func GetSaves(c *gin.Context) {
	client, err := database.New(c)
	if err != nil {
		slog.Error("get saves new database error: ", err.Error())
		server.ErrorInternalServerError(c)
	}

	if err := server.NewServer(c, client.Client).GetSaves(); err != nil {
		slog.Error("get saves error: ", err.Error())
		server.ErrorInternalServerError(c)
	}
}

func GetSaved(c *gin.Context) {
	client, err := database.New(c)
	if err != nil {
		slog.Error("get saved new database error: ", err.Error())
		server.ErrorInternalServerError(c)
	}

	if err := server.NewServer(c, client.Client).GetSaved(); err != nil {
		slog.Error("get saved error: ", err.Error())
		server.ErrorInternalServerError(c)
	}
}

func EditSaved(c *gin.Context) {
	client, err := database.New(c)
	if err != nil {
		slog.Error("edit saved new database error: ", err.Error())
		server.ErrorInternalServerError(c)
	}

	if err := server.NewServer(c, client.Client).EditSaved(); err != nil {
		slog.Error("edit saved error: ", err.Error())
		server.ErrorInternalServerError(c)
	}
}

func DelSaved(c *gin.Context) {
	client, err := database.New(c)
	if err != nil {
		slog.Error("del saved new database error: ", err.Error())
		server.ErrorInternalServerError(c)
	}

	if err := server.NewServer(c, client.Client).DelSaved(); err != nil {
		slog.Error("del saved error: ", err.Error())
		server.ErrorInternalServerError(c)
	}
}

func LinkDir(c *gin.Context) {
	client, err := database.New(c)
	if err != nil {
		slog.Error("link dir new database error: ", err.Error())
		server.ErrorInternalServerError(c)
	}

	if err := server.NewServer(c, client.Client).LinkDir(); err != nil {
		slog.Error("link dir error: ", err.Error())
		server.ErrorInternalServerError(c)
	}
}

func UnlinkDir(c *gin.Context) {
	client, err := database.New(c)
	if err != nil {
		slog.Error("unlink dir new database error: ", err.Error())
		server.ErrorInternalServerError(c)
	}

	if err := server.NewServer(c, client.Client).UnlinkDir(); err != nil {
		slog.Error("unlink dir error: ", err.Error())
		server.ErrorInternalServerError(c)
	}
}
