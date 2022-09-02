package uploadController

import (
	"github.com/gin-gonic/gin"
)

func GetUploads(c *gin.Context) {

}

func CreateUploads(c *gin.Context) {
	// bucketName := os.Getenv("MINIO_BUCKET")
	// var inputs *ent.Upload

	// if err := c.ShouldBindJSON(&inputs); err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
	// 	return
	// }

	// upload, err := MinioClient.GetObject(context.Background(), bucketName, "haj", minio.GetObjectOptions{})

	// if err != nil {
	// 	fmt.Println(err.Error())
	// 	c.AbortWithStatus(http.StatusBadRequest)
	// 	return
	// }
	// c.JSON(http.StatusOK, upload)
}

func EditUpload(c *gin.Context) {

}

func GetUpload(c *gin.Context) {

}

func DeleteUpload(c *gin.Context) {

}
