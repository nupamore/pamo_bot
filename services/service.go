package services

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/translate"
	"github.com/diamondburned/arikawa/api"
	"github.com/nupamore/pamo_bot/configs"
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
		configs.Env["DB_USER"],
		configs.Env["DB_PASS"],
		configs.Env["DB_HOST"],
		configs.Env["DB_NAME"],
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
