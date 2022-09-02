package main

import (
	"education/api"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	api.RoutesPool(router)

	router.Run(":3030")
}
