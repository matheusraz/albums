package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
	{ID: "4", Title: "First Internation Job", Artist: "Raz", Price: 6250},
}

func GetAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func PostAlbum(c *gin.Context) {

	var newAlbum album

	err := c.BindJSON(&newAlbum)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "Could not flush album into database :("})
		return
	}

	albums = append(albums, newAlbum)

	c.JSON(http.StatusCreated, gin.H{"msg": "Album added succesfully"})

}

func Health(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"msg": "Server Up and Running!"})
}

func GetAlbumByID(c *gin.Context) {
	// albumId := c.GetHeader("albumId")
	albumId := c.Param("albumId")
	// albumId := c.Query("albumId")

	for _, album := range albums {
		if album.ID == albumId {
			c.JSON(http.StatusFound, album)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"msg": "Album could not be found :("})
}
