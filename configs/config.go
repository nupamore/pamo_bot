package configs

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
)

const envFilePath = "configs/.env"

// Env : dotenv
var Env map[string]string

func init() {
	Env = map[string]string{
		// api
		"OAUTH_ENDPOINT": "https://discord.com/api/oauth2",
		"OAUTH_API":      "https://discord.com/api",
		"KAKAO_API":      "https://dapi.kakao.com",
		"NAVER_API":      "https://openapi.naver.com",

		// sercret
		"BOT_TOKEN":             "",
		"OAUTH_KEY":             "",
		"OAUTH_SECRET":          "",
		"DB_USER":               "",
		"DB_PASS":               "",
		"DB_HOST":               "",
		"DB_NAME":               "",
		"AWS_REGION":            "",
		"AWS_ACCESS_KEY_ID":     "",
		"AWS_SECRET_ACCESS_KEY": "",
		"KAKAO_KEY":             "",
		"NAVER_ID":              "",
		"NAVER_SECRET":          "",

		// config
		"BOT_PREFIX":         "",
		"BOT_STATUS":         "",
		"SERVER_PORT":        "",
		"OAUTH_CALLBACK":     "",
		"WEB_URL":            "",
		"LINK_RESERVE_PAGE":  "",
		"LINK_RESERVE_IMAGE": "",
		"LINK_RESERVE_VIDEO": "",
	}

	err := godotenv.Load(envFilePath)
	env, err := godotenv.Read(envFilePath)
	if err != nil {
		log.Println("Dotenv load fail")
		panic(err)
	}

	for key := range Env {
		val, exist := env[key]
		if exist {
			Env[key] = val
		} else if Env[key] == "" {
			msg := fmt.Sprintf("Can't find key '%s'", key)
			log.Println(msg)
			// panic(errors.New(msg))
		}
	}
}
