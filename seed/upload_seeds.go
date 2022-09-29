package seed

import (
	"context"
	"education/config"
	"education/ent"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	jsoniter "github.com/json-iterator/go"
	"github.com/minio/minio-go/v7"
	"github.com/mitchellh/mapstructure"
)

func SeedUpload(c context.Context) error {
	var fileJSON []gin.H
	var json = jsoniter.ConfigCompatibleWithStandardLibrary

	file, err := os.ReadFile("/home/mohammed/Desktop/education/public/upload.json")

	if err != nil {
		fmt.Println(err.Error())
	}

	err = json.Unmarshal(file, &fileJSON)
	if err != nil {
		fmt.Println(err.Error())
	}
	bulk := make([]*ent.UploadCreate, len(fileJSON))

	for i, up := range fileJSON {
		var inputs ent.Upload
		var fileSize int64
		fileSize = 5

		err := mapstructure.Decode(up, &inputs)

		if err != nil {
			fmt.Println(err.Error())
		}

		bucketName := os.Getenv("MINIO_BUCKET")
		mimeType := "application/octet-stream"
		file_uuid := (uuid.New()).String()
		filePath, err := os.Open("/home/mohammed/Desktop/education/public/necklace.odt")
		if err != nil {
			fmt.Println(err)
		}

		bulk[i] = config.Client.Upload.Create().
			SetPlaylistID(inputs.PlaylistID).
			SetUserID(inputs.UserID).
			SetUUID(file_uuid).
			SetName(inputs.Name).
			SetMimeType(mimeType).
			SetSize(int(fileSize)).
			SetTitle(inputs.Title)

		config.MinioClient.PutObject(context.Background(), bucketName, file_uuid, filePath, fileSize, minio.PutObjectOptions{ContentType: mimeType})
	}

	_, err = config.Client.Upload.CreateBulk(bulk...).Save(c)

	if err != nil {
		fmt.Println(err.Error())
	}
	return nil
}
