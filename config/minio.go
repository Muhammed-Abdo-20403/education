package config

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

var minioClient *minio.Client

func init() {

	var err error
	err = godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	endpoint := os.Getenv("MINIO_ENDPOINT")
	bucketName := os.Getenv("MINIO_BUCKET")
	accessKeyID := os.Getenv("MINIO_ACCESSKEY")
	secretAccessKey := os.Getenv("MINIO_SECRETKEY")
	useSSL := false

	minioClient, err = minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: useSSL,
	})

	if err != nil {
		fmt.Println("-----------------111111---------------------------")
		fmt.Println(err.Error())
		fmt.Println("--------------------------------------------")
		log.Fatalln(err)
	}

	// mimeType := "application/vnd.oasis.opendocument.text"
	file, err := os.Open("/home/mohammed/Desktop/education/public/haj.odt")

	if err != nil {
		fmt.Println("-------------222222-------------------------------")
		fmt.Println(err.Error())
		fmt.Println("--------------------------------------------")
		return
	}
	defer file.Close()

	fileStat, err := file.Stat()
	if err != nil {
		fmt.Println("-------------33333333333-------------------------------")
		fmt.Println(err.Error())
		fmt.Println("--------------------------------------------")
		return
	}

	upload, err := minioClient.PutObject(context.Background(), bucketName, "Hag", file, fileStat.Size(), minio.PutObjectOptions{ContentType: "application/octet-stream"})

	if err != nil {
		fmt.Println("-----------------44444444444444444444---------------------------")
		fmt.Println(err.Error(), upload.Location)
		fmt.Println("--------------------------------------------")
		return
	}
	fmt.Println("Successfully uploaded bytes: ", upload)

}
