package configs

import (
	"log"

	"github.com/joho/godotenv"
)

// Env : dotenv
var Env map[string]string

func init() {
	env, err := godotenv.Read("configs/.env")
	if err != nil {
		log.Println("Config error")
	}

	Env = env
}
