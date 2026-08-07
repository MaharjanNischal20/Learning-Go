package main

import (
	"github.com/gin-gonic/gin"
)

var albums = []album{
	{ID: 1, Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: 2, Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: 3, Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func main() {
	r := gin.Default()
	r.GET("/albums", getAlbums)
	r.POST("/albums", addAlbums)
	r.GET("/albums/:id", getAlbumByID)

	err := r.Run("localhost:8080")
	if err != nil {
		return
	}
}
