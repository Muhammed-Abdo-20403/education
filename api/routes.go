package api

import (
	uploadController "education/api/upload_controller"
	userController "education/api/user_controller"

	"github.com/gin-gonic/gin"
)

func RoutesPool(c *gin.Engine) {
	router := gin.Default()

	users := router.Group("/users")
	{
		users.GET("", userController.GetUsers)
		users.POST("", userController.CreateUsers)
		users.PUT("/:id", userController.EditUser)
		users.GET("/:id", userController.GetUser)
		users.DELETE("/:id", userController.DeleteUser)
	}

	uploads := router.Group("/uploads")
	{
		uploads.GET("", uploadController.GetUploads)
		uploads.POST("", uploadController.CreateUploads)
		uploads.PUT("/:id", uploadController.EditUpload)
		uploads.GET("/:id", uploadController.GetUpload)
		uploads.DELETE("/:id", uploadController.DeleteUpload)
	}
}
