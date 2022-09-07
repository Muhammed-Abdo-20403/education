package userController

import (
	"education/config"
	"education/ent"
	"education/ent/user"
	"fmt"
	"net/http"
	"strconv"

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
	var inputs *ent.User

	if err := c.ShouldBindJSON(&inputs); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	userID, _ := strconv.Atoi(c.Param("id"))

	user, err := config.Client.User.
		UpdateOneID(userID).
		SetName(inputs.Name).
		SetEmail(inputs.Email).
		Save(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
	}

	c.JSON(http.StatusOK, user)
}

func GetUser(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
	}

	user, err := config.Client.User.
		Query().
		Where(user.ID(userID)).
		First(c)

	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusBadRequest)
	}

	c.JSON(http.StatusOK, user)
}

func DeleteUser(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	user := config.Client.User.DeleteOneID(userID).Exec(c)

	c.JSON(http.StatusCreated, user)
}
