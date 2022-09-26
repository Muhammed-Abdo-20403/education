package seed

import (
	"context"
	"education/config"
	"education/ent"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	jsoniter "github.com/json-iterator/go"
	"github.com/mitchellh/mapstructure"
)

func SeedPlaylist(ctx context.Context) error {
	var fileJSON []gin.H
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	file, err := os.ReadFile("/home/mohammed/Desktop/education/public/playlist.json")
	if err != nil {
		fmt.Println(err.Error())
	}

	err = json.Unmarshal(file, &fileJSON)
	if err != nil {
		fmt.Println(err.Error())
	}
	bulk := make([]*ent.PlaylistCreate, len(fileJSON))
	for i, p := range fileJSON {
		var playlist ent.Playlist
		err := mapstructure.Decode(p, &playlist)
		if err != nil {
			fmt.Println(err.Error())
		}
		bulk[i] = config.Client.Playlist.Create().
			SetUserID(playlist.UserID).
			SetTitle(playlist.Title)
	}
	_, err = config.Client.Playlist.CreateBulk(bulk...).Save(ctx)
	if err != nil {
		return fmt.Errorf("creating Playlist %w", err)
	}

	return nil
}
