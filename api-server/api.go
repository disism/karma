package api_server

import (
	"github.com/gin-gonic/gin"
)

func Run() error {
	r := gin.Default()

	r.Use(CORS())

	r.GET("/version", GetVersion)
	r.GET("/.well-known/node-info", GetNodeInfo)

	r.POST("/users/create", CreateUser)
	r.POST("/login", Login)

	r.Use(AuthMW())

	devices := r.Group("/_devices/v1")
	{
		devices.GET("", GetDevices)
		devices.DELETE("/:id", DeleteDevice)
	}

	ipfs := r.Group("/_ipfs/v1")
	{
		ipfs.POST("/add", IPFSAddFiles)
	}
	saved := r.Group("/_saved/v1")
	{

		saved.POST("", CreateSaves)
		saved.GET("", GetSaves)
		saved.GET(":id", GetSaved)
		saved.PUT(":id", EditSaved)
		saved.DELETE(":id", DelSaved)
		saved.PUT(":id/link", LinkDir)
		saved.PUT(":id/unlink", UnlinkDir)

		dirs := saved.Group("/dirs")
		{
			dirs.POST("", MKDir)
			dirs.GET("", ListDir)
			dirs.GET("/all", ListDirs)
			dirs.PATCH("/:id/name", RenameDir)
			dirs.PUT("/:id/mv/:new_id", MVDir)
			dirs.DELETE("/:id", RMDir)
		}
	}

	if err := r.Run(":7330"); err != nil {
		return err
	}
	return nil
}
