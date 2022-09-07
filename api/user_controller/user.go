package userController

import (
	"education/config"
	"education/ent"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {

	users, err := config.Client.User.Query().All(c)

	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusBadRequest)
	}

	c.JSON(http.StatusOK, users)
}

func CreateUsers(c *gin.Context) {
	var inputs *ent.User

	if err := c.ShouldBindJSON(&inputs); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	user, err := config.Client.User.
		Create().
		SetName(inputs.Name).
		SetChannelName(inputs.ChannelName).
		SetEmail(inputs.Email).
		SetSpecialist(inputs.Specialist).
		SetAge(inputs.Age).
		SetPhone(inputs.Phone).
		SetLanguage(inputs.Language).
		SetCountry(inputs.Country).
		SetShorBio(inputs.ShorBio).
		Save(c)

	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusBadRequest)
	}

	c.JSON(http.StatusCreated, user)
}

func EditUser(c *gin.Context) {

}

func GetUser(c *gin.Context) {

}

func DeleteUser(c *gin.Context) {

}
