package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

var MinioClient *minio.Client

func init() {

	var err error
	err = godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	endpoint := os.Getenv("MINIO_ENDPOINT")
	accessKeyID := os.Getenv("MINIO_ACCESSKEY")
	secretAccessKey := os.Getenv("MINIO_SECRETKEY")
	useSSL := false

	MinioClient, err = minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: useSSL,
	})

	if err != nil {
		fmt.Println("-----------------111111---------------------------")
		fmt.Println(err.Error())
		fmt.Println("--------------------------------------------")
		log.Fatalln(err)
	}
}
