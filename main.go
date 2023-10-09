package main

import (
	"log"
	"smart-clothesline-http/database"
	"smart-clothesline-http/handlers"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error al cargar el archivo .env")
	}
}

func main() {
	database.Init()
	handlers.Init()
}
