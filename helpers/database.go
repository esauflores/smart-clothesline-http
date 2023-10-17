package helpers

import (
	"fmt"
	"os"
	"smart-clothesline-http/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func OpenDBConnection() {
	// Get the dsn string
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s TimeZone=America/El_Salvador",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	// Try to connect to the database
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	DB = db
	if err != nil {
		panic("no se pudo establecer la conexión con la base de datos")
	}
}

func CloseDBConnection() {
	// Close the database connection
	db, err := DB.DB()
	if err != nil {
		panic("no se pudo cerrar la conexión con la base de datos")
	}

	err = db.Close()
	if err != nil {
		panic("no se pudo cerrar la conexión con la base de datos")
	}
}

func MigrateDB() {
	OpenDBConnection()
	defer CloseDBConnection()

	DB.AutoMigrate(&models.Usuario{})
	DB.AutoMigrate(&models.Tendedero{})
	DB.AutoMigrate(&models.Evento{})
}
