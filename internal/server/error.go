package server

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type ErrorResponse struct {
	Code  int    `json:"code"`
	Error string `json:"error"`
}

const (
	InternalServerError = "InternalServerError"
	NoPermission        = "no permission"
)

func ErrorJSONResponse(c *gin.Context, status int, message string) {
	errResp := ErrorResponse{
		Code:  status,
		Error: message,
	}
	c.JSON(status, errResp)
}

func ErrorInternalServerError(c *gin.Context) {
	ErrorJSONResponse(c, http.StatusInternalServerError, InternalServerError)
	return
}

func ErrorConflict(c *gin.Context, message string) {
	ErrorJSONResponse(c, http.StatusConflict, message)
}

func ErrorNotFound(c *gin.Context, message string) {
	ErrorJSONResponse(c, http.StatusNotFound, message)
}

func ErrorUnauthorized(c *gin.Context, message string) {
	ErrorJSONResponse(c, http.StatusUnauthorized, message)
}

func ErrorBadRequest(c *gin.Context, message string) {
	ErrorJSONResponse(c, http.StatusBadRequest, message)
}

func ErrorNoPermission(c *gin.Context, message string) {
	ErrorJSONResponse(c, http.StatusForbidden, message)
}
