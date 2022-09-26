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

func SeedUser(ctx context.Context) error {
	var fileJSON []gin.H
	var json = jsoniter.ConfigCompatibleWithStandardLibrary

	file, err := os.ReadFile("/home/mohammed/Desktop/education/public/user.json")

	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("Successfully Opened user.json")

	err = json.Unmarshal(file, &fileJSON)
	if err != nil {
		fmt.Println(err.Error())
	}

	bulk := make([]*ent.UserCreate, len(fileJSON))
	for i, u := range fileJSON {
		var user ent.User
		err := mapstructure.Decode(u, &user)
		if err != nil {
			fmt.Println(err.Error())
		}
		bulk[i] = config.Client.User.Create().
			SetName(user.Name).
			SetChannelName(user.ChannelName).
			SetEmail(user.Email).
			SetSpecialist(user.Specialist).
			SetAge(user.Age).
			SetPhone(user.Phone).
			SetLanguage(user.Language).
			SetCountry(user.Country).
			SetShorBio(user.ShorBio)
	}
	_, err = config.Client.User.CreateBulk(bulk...).Save(ctx)

	if err != nil {
		return fmt.Errorf("creating Users: %w", err)
	}

	return nil
}
