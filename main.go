package main

import (
	"smart-clothesline-http/handlers"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	// helpers.MigrateDB()
	handlers.InitRouter()
}
