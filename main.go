package main

import (
	"github.com/matheusraz/albums/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Server endpoints
	router.GET("/healthz", routes.Health)
	router.GET("/albums", routes.GetAlbums)
	router.POST("/album", routes.PostAlbum)
	router.GET("/album/:albumId", routes.GetAlbumByID)

	// Run Server
	router.Run()
}
