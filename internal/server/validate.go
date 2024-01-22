package server

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/disism/godis/jwt"
	"github.com/gin-gonic/gin"
	"golang.org/x/exp/slog"
)

const (
	NoAuthorizationHeader      = "no authorization header"
	TokenAuthenticationFailure = "token authentication failure"
	TokenInvalid               = "invalid authorization"
	Unauthorized               = "unauthorized"
)

const (
	ContextUserIDKEY   = "UserID"
	ContextDeviceIDKEY = "DeviceID"
)

func GetUserID(ctx *gin.Context) uint64 {
	const z = 0
	g, exists := ctx.Get(ContextUserIDKEY)
	if !exists {
		ErrorUnauthorized(ctx, Unauthorized)
		return z
	}

	id, err := strconv.ParseUint(g.(string), 10, 64)
	if err != nil {
		ErrorInternalServerError(ctx)
		return z
	}
	return id
}

func GetDeviceID(ctx *gin.Context) string {
	g, exists := ctx.Get(ContextDeviceIDKEY)
	if !exists {
		ErrorUnauthorized(ctx, Unauthorized)
		return ""
	}
	if id, ok := g.(uint64); ok {
		return strconv.FormatUint(id, 10)
	}
	ErrorInternalServerError(ctx)
	return ""
}

func (s *Server) ValidateJWT() error {
	a := s.ctx.GetHeader("Authorization")
	if a == "" {
		return fmt.Errorf(NoAuthorizationHeader)
	}

	bearer := strings.Split(a, "Bearer ")
	if len(bearer) != 2 {
		return fmt.Errorf(TokenInvalid)
	}

	validate, err := jwt.Validate(bearer[1], JWTSecret())
	if err != nil {
		return fmt.Errorf(TokenAuthenticationFailure)
	}

	id, err := strconv.ParseUint(validate.ID, 10, 64)
	if err != nil {
		slog.Error("validate jwt parse id error: ", err)
		return fmt.Errorf(Unauthorized)
	}

	device, err := s.client.Devices.Get(s.ctx, id)
	if err != nil {
		slog.Error("validate jwt get device error: ", err)
		return fmt.Errorf(Unauthorized)
	}

	s.ctx.Set(ContextUserIDKEY, validate.Subject)
	s.ctx.Set(ContextDeviceIDKEY, device.ID)
	return nil
}
