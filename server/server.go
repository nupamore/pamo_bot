package main

import (
	"log"

	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load("../.env")
}

func main() {
	log.Print("server")
}
