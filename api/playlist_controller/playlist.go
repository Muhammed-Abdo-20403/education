package playlistController

import (
	"education/config"
	"education/ent"
	"education/ent/playlist"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetPlaylists(c *gin.Context) {

	playlist, err := config.Client.Playlist.Query().All(c)

	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusBadRequest)
	}

	c.JSON(http.StatusOK, gin.H{"playlist": playlist})

}

func CreatePlaylist(c *gin.Context) {
	var inputs *ent.Playlist

	if err := c.ShouldBindJSON(&inputs); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	playlists, err := config.Client.Playlist.Create().
		SetUserID(inputs.UserID).
		SetTitle(inputs.Title).
		Save(c)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	c.JSON(http.StatusOK, playlists)
}

func EditPlaylist(c *gin.Context) {
	var inputs *ent.Playlist

	if err := c.ShouldBindJSON(&inputs); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	playlistID, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		fmt.Println("--------------------------------------------")
		fmt.Println(err.Error())
		fmt.Println("--------------------------------------------")
		return
	}

	playlist, err := config.Client.Playlist.UpdateOneID(playlistID).SetTitle(inputs.Title).Save(c)

	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	c.JSON(http.StatusOK, playlist)
}

func GetPlaylist(c *gin.Context) {

	playlistID, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		fmt.Println("---------------------------------------")
		fmt.Println(err.Error())
		fmt.Println("---------------------------------------")
		return
	}

	playlist, err := config.Client.Playlist.Query().Where(playlist.ID(playlistID)).First(c)

	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	c.JSON(http.StatusOK, playlist)
}

func DeletePlaylist(c *gin.Context) {
	playlistID, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		fmt.Println("---------------------------------------")
		fmt.Println(err.Error())
		fmt.Println("---------------------------------------")
	}

	playlist := config.Client.Playlist.DeleteOneID(playlistID).Exec(c)

	c.JSON(http.StatusCreated, playlist)
}
