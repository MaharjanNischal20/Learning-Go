package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func addAlbums(c *gin.Context) {
	var newAlbums []album

	if err := c.BindJSON(&newAlbums); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	for _, newAlbum := range newAlbums {
		albums = append(albums, newAlbum)
	}
	c.IndentedJSON(http.StatusCreated, newAlbums)
}

func getAlbumByID(c *gin.Context) {
	id := c.Param("id")

	intId, err := strconv.Atoi(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	for _, album := range albums {
		if album.ID == intId {
			c.IndentedJSON(http.StatusOK, album)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{
		"error": "Album not found",
	})
}
