package playlistController

import (
	"education/config"
	"education/ent"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetPlaylists(c *gin.Context) {
	var outputs *ent.Playlist

	if err := c.ShouldBindJSON(outputs); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	playlist, err := config.Client.Playlist.Query().All(c)

	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusBadRequest)
	}

	c.JSON(http.StatusOK, playlist)

}

func CreatePlaylist(c *gin.Context) {

}

func EditPlaylist(c *gin.Context) {

}

func GetPlaylist(c *gin.Context) {

}

func DeletePlaylist(c *gin.Context) {

}
