package main

import (
	"context"
	"education/api"
	"education/seed"

	"github.com/gin-gonic/gin"
)

func main() {

	seed.Do(context.Background())
	seed.DO(context.Background())
	router := gin.Default()

	api.RoutesPool(router)

	router.Run(":3030")
}
