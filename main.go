package main

import (
	"context"
	"education/api"
	"education/seed"

	"github.com/gin-gonic/gin"
)

func main() {

	// seed.SeedUser(context.Background())
	seed.SeedPlaylist(context.Background())
	// seed.SeedUpload(context.Background())
	router := gin.Default()

	api.RoutesPool(router)

	router.Run(":3030")
}
