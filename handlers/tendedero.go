package handlers

import (
	"errors"
	"net/http"
	"os"
	"smart-clothesline-http/helpers"
	"smart-clothesline-http/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetTendederos(c *gin.Context) {
	// Variable declaration
	var tendederos []models.Tendedero

	// Query handling
	helpers.OpenDBConnection()

	result := helpers.DB.Find(&tendederos)
	helpers.CheckFatal(result.Error, http.StatusInternalServerError, errors.New("no se pudo obtener los tendederos"))

	helpers.CloseDBConnection()

	// Response handling
	c.JSON(http.StatusOK, tendederos)
}

func GetTendedero(c *gin.Context) {
	// Variable declaration
	var tendedero models.Tendedero

	// Param handling
	device_id, err := helpers.GetURLParam(c, "device_id")
	helpers.CheckFatal(err, http.StatusBadRequest, err)

	// Query handling
	helpers.OpenDBConnection()

	result := helpers.DB.Where("id = ?", device_id).First(&tendedero)
	helpers.CheckFatal(result.Error, http.StatusInternalServerError, errors.New("no se pudo obtener los tendederos"))

	helpers.CloseDBConnection()

	// Response handling
	c.JSON(http.StatusOK, tendedero)
}

func PatchTendedero(c *gin.Context) {
	// Param handling
	modo, err := helpers.GetURLParam(c, "modo")
	helpers.CheckFatal(err, http.StatusBadRequest, err)

	estado, err := helpers.GetURLParam(c, "estado")
	helpers.CheckFatal(err, http.StatusBadRequest, err)

	// Validation
	if modo != MODO_AUTO && modo != MODO_MANUAL {
		helpers.Fatal(http.StatusBadRequest, errors.New("modo invalido"))
	}

	if estado != ESTADO_AFUERA && estado != ESTADO_ADENTRO {
		helpers.Fatal(http.StatusBadRequest, errors.New("estado invalido"))
	}

	// File handling
	modo_int, _ := strconv.Atoi(modo)
	estado_int, _ := strconv.Atoi(estado)

	message := strconv.Itoa(modo_int*2 + estado_int)
	helpers.WriteToFile(os.Getenv("DEVICE_PATH"), message)

	// Query handling
	helpers.OpenDBConnection()

	result := helpers.DB.Model(&models.Tendedero{}).Where("id = ?", os.Getenv("DEVICE_ID")).Update("estado", estado).Update("modo", modo)
	helpers.CheckFatal(result.Error, http.StatusInternalServerError, errors.New("no se pudo actualizar el tendedero"))

	helpers.CloseDBConnection()

	// Response handling
	c.JSON(http.StatusOK, gin.H{"mensaje": "Actualizado exitoso"})
}
