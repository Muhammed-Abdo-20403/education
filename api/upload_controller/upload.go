package uploadController

import (
	"context"
	"education/config"
	"education/ent"
	"education/ent/upload"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/minio/minio-go/v7"
)

func GetUploads(c *gin.Context) {
	uploads, err := config.Client.Upload.Query().All(c)

	if err != nil {
		fmt.Println("-------------------------------")
		fmt.Println(err.Error())
		fmt.Println("-------------------------------")
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
	}
	c.JSON(http.StatusOK, uploads)

}

func CreateUploads(c *gin.Context) {
	bucketName := os.Getenv("MINIO_BUCKET")
	mimeType := "application/octet-stream"
	file, err := c.FormFile("mytestfile")
	openFile, _ := file.Open()
	if err != nil {
		fmt.Println("--------------------------------------------")
		fmt.Println(err.Error())
		fmt.Println("--------------------------------------------")
		return
	}

	file_uuid := (uuid.New()).String()

	if err != nil {
		fmt.Println("--------------------------------------------")
		fmt.Println(err.Error())
		fmt.Println("--------------------------------------------")
	}

	upload, err := config.MinioClient.PutObject(context.Background(), bucketName, file_uuid, openFile, file.Size, minio.PutObjectOptions{ContentType: mimeType})

	if err != nil {
		fmt.Println("--------------------------------------------")
		fmt.Println(err.Error(), upload.Location)
		fmt.Println("--------------------------------------------")
		return
	}
	fmt.Println("Successfully uploaded bytes: ", upload)

	uploads, err := config.Client.Upload.Create().SetPlaylistID(1).SetUserID(1).SetUUID(file_uuid).SetName(file.Filename).SetMimeType(mimeType).SetSize(int(file.Size)).Save(c)

	if err != nil {
		fmt.Println("----------------------------------")
		fmt.Println(err.Error())
		fmt.Println("----------------------------------")
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, uploads)
}

func EditUpload(c *gin.Context) {
	var inputs *ent.Upload

	if err := c.ShouldBindJSON(&inputs); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
	}

	uploadID, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		fmt.Println("------------------------------")
		fmt.Println(err.Error())
		fmt.Println("------------------------------")
	}

	uploads, err := config.Client.Upload.UpdateOneID(uploadID).SetTitle(inputs.Title).Save(c)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
	}

	c.JSON(http.StatusCreated, uploads)
}

func GetUpload(c *gin.Context) {
	uploadID, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		fmt.Println("-----------------------------------")
		fmt.Println(err.Error())
		fmt.Println("-----------------------------------")
	}

	upload, err := config.Client.Upload.Query().Where(upload.ID(uploadID)).First(c)

	if err != nil {
		fmt.Println("-----------------------------------")
		fmt.Println(err.Error())
		fmt.Println("-----------------------------------")
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
	}

	c.JSON(http.StatusOK, upload)
}

func DeleteUpload(c *gin.Context) {
	bucketName := os.Getenv("MINIO_BUCKET")

	uploadID, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	u, err := config.Client.Upload.Query().Where(upload.ID(uploadID)).First(c)

	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	err = config.Client.Upload.DeleteOneID(u.ID).Exec(c)

	if err != nil {

		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	opts := minio.RemoveObjectOptions{
		ForceDelete:      true,
		GovernanceBypass: true,
	}

	err = config.MinioClient.RemoveObject(context.Background(), bucketName, u.UUID, opts)

	if err != nil {
		fmt.Println(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.Status(http.StatusOK)

}
