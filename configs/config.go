package configs

import (
	"log"

	"github.com/joho/godotenv"
)

const envFilePath = "configs/.env"

// Env : dotenv
var Env map[string]string

func init() {
	err := godotenv.Load(envFilePath)
	env, err := godotenv.Read(envFilePath)
	if err != nil {
		log.Println("Config error")
	}

	Env = env
}
