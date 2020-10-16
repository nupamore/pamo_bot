package services

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/translate"
	"github.com/diamondburned/arikawa/api"
	_ "github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-mysql/driver" // mysql driver
)

// DiscordAPI : discord api
var DiscordAPI *api.Client

// DB : db
var DB *sql.DB

// DBsetup : db init
func DBsetup() {
	// db setup
	dsn := fmt.Sprintf("%s:%s@(%s)/%s?%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_NAME"),
		"charset=utf8&parseTime=True&loc=Local",
	)
	db, err := sql.Open("mysql", dsn)
	DB = db

	if err != nil {
		log.Println("AWS init fail")
		panic(err)
	}
}

// AWStranslate : aws translate client
var AWStranslate *translate.Client

// AWSsetup : aws init
func AWSsetup() {
	cfg, err := config.LoadDefaultConfig()
	if err != nil {
		log.Println("AWS init fail")
		panic(err)
	}
	AWStranslate = translate.NewFromConfig(cfg)
}
