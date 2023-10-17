package main

import (
	"smart-clothesline-http/handlers"
	"smart-clothesline-http/helpers"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	helpers.MigrateDB()
	handlers.InitRouter()
}
