package handlers

import (
	"errors"
	"net/http"
	"smart-clothesline-http/helpers"
	"smart-clothesline-http/models"

	"github.com/gin-gonic/gin"
)

func GetUsuarios(c *gin.Context) {
	// Variable declaration
	var usuarios []models.Usuario

	// Query handling
	helpers.OpenDBConnection()
	defer helpers.CloseDBConnection()

	// users with their tendederos, hide the password
	result := helpers.DB.Preload("Tendederos").Select("id, nombres, apellidos, email").Find(&usuarios)
	helpers.CheckFatal(result.Error, http.StatusInternalServerError, errors.New("no se pudo obtener los usuarios"))

	// Response handling
	c.JSON(http.StatusOK, usuarios)
}
