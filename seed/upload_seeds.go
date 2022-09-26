package seed

// func SeedUpload(c context.Context) error {
// 	var fileJSON []gin.H
// 	var json = jsoniter.ConfigCompatibleWithStandardLibrary
// 	file, err := os.ReadFile("/home/mohammed/Desktop/education/public/upload.json")
// 	if err != nil {
// 		fmt.Println(err.Error())
// 	}

// 	err = json.Unmarshal(file, &fileJSON)
// 	if err != nil {
// 		fmt.Println(err.Error())
// 	}
// 	bulk := make([]*ent.UploadCreate, len(fileJSON))

// 	for i, up := range fileJSON {
// 		var upload ent.Uplaod
// 		err := mapstructure.Decode(up, &upload)
// 		if err != nil {
// 			fmt.Println(err.Error())
// 		}
// 		bucketName := os.Getenv("MINIO_BUCKET")
// 		mimeType := "application/octet-stream"
// 		file_uuid := (uuid.New()).String()

// 		bulk[i] = config.Client.Upload.Create().
// 			SetPlaylistID(upload.PlaylistID).
// 			SetUserID(upload.UserID).
// 			SetUUID(file_uuid).
// 			SetName(file.file).
// 			SetMimeType(mimeType).
// 			SetSize(int(file.Size))

// 		config.MinioClient.PutObject(context.Background(), bucketName, file_uuid, openFile, file.Size, minio.PutObjectOptions{ContentType: mimeType})
// 	}
// 	return nil
// }
