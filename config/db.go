package config

import (
	"context"
	"fmt"
	"log"
	"os"

	// "education/ent"
	// "education/ent/migrate"

	"entgo.io/ent/examples/fs/ent"
	"entgo.io/ent/examples/fs/ent/migrate"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var Client *ent.Client

func init() {
	var err error
	err = godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	Dbdriver := os.Getenv("DB_DRIVER")
	DbHost := os.Getenv("DB_HOST")
	DbUser := os.Getenv("DB_USER")
	DbPassword := os.Getenv("DB_PASSWORD")
	DbName := os.Getenv("DB_NAME")
	DbPort := os.Getenv("DB_PORT")

	DBURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", DbUser, DbPassword, DbHost, DbPort, DbName)

	client, err := ent.Open(Dbdriver, DBURL)

	if err != nil {
		log.Fatalf("failed connecting to mysql: %v", err)
	}

	ctx := context.Background()

	err = client.Schema.Create(
		ctx,
		migrate.WithDropIndex(true),
		migrate.WithDropColumn(true),
	)

	if err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

}
