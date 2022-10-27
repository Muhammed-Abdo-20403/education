package main

import (
	"education/api"
	"education/config"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
)

func main() {
	// seed.SeedUser(context.Background())
	// seed.SeedPlaylist(context.Background())
	// seed.SeedUpload(context.Background())
	router := gin.Default()

	api.RoutesPool(router)

	RestClient := resty.New()

	resp, err := RestClient.R().
		EnableTrace().
		Get("http://127.0.0.1:3031/users/23")

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("  Status Code:", resp.StatusCode())
	fmt.Println("  Status     :", resp.Status())
	fmt.Println("  Proto      :", resp.Proto())
	fmt.Println("  Time       :", resp.Time())
	fmt.Println("  Received At:", resp.ReceivedAt())
	fmt.Println("  Body       :\n", resp)

	// config.SendGrid()
	config.Sms()

	router.Run(":3030")
}
