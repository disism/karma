package server

import (
	"github.com/disism/karma/ent"
	"github.com/gin-gonic/gin"
)

type Server struct {
	ctx    *gin.Context
	client *ent.Client
}

func NewServer(ctx *gin.Context, client *ent.Client) *Server {
	return &Server{ctx: ctx, client: client}
}

type UsersServer interface {
	// Login is a function that handles user authentication.
	//
	// It does not take any parameters.
	// It returns an error if there is an issue with the authentication process.
	Login() error

	// CreateUser creates a user.
	//
	// It returns an error if there was a problem creating the user.
	CreateUser() error
}

type DevicesServer interface {
	GetDevices() error

	DeleteDevice() error
}

type IPFSServer interface {
	IPFSAddFiles() error
}

type SavedServer interface {
	AddSaves() error

	GetSaves() error

	GetSaved() error

	EditSaved() error

	DelSaved() error

	LinkDir() error

	UnlinkDir() error
}

type DirServer interface {
	ListDir() error

	ListDirs() error

	RenameDir() error

	MVDir() error

	RMDir() error
}
