package handlers

import (
	"errors"
	"net/http"
	"os"
	"smart-clothesline-http/database"
	"smart-clothesline-http/models"

	"github.com/gin-gonic/gin"
)

const MODO_AUTO = "0"
const MODO_MANUAL = "1"

const ESTADO_AFUERA = "0"
const ESTADO_ADENTRO = "1"

func GetTendederos(c *gin.Context) {
	// Database handling
	if database.Connection == nil {
		database.Init()
		if database.Connection == nil {
			Fatal(http.StatusInternalServerError, errors.New("no se pudo conectar con la bd"))
		}
	}

	// Query handling
	var tendederos []models.Tendedero

	result := database.Connection.Find(&tendederos)
	CheckFatal(result.Error, http.StatusInternalServerError, errors.New("no se pudo obtener los tendederos"))

	c.JSON(http.StatusOK, tendederos)
}

func PatchTendedero(c *gin.Context) {
	// Param handling
	modo, err := GetParam(c, "modo")
	CheckFatal(err, http.StatusBadRequest, errors.New("no se pudo obtener el modo"))

	if modo != MODO_AUTO && modo != MODO_MANUAL {
		Fatal(http.StatusBadRequest, errors.New("el modo debe ser manual o automatico"))
	}

	estado, err := GetParam(c, "estado")
	CheckFatal(err, http.StatusBadRequest, errors.New("no se pudo obtener el estado"))

	if estado != ESTADO_AFUERA && estado != ESTADO_ADENTRO {
		Fatal(http.StatusBadRequest, errors.New("el estado debe ser adentro o afuera"))
	}

	// File handling
	const devicePath = "/dev/ttyACM0"
	file, err := os.OpenFile(devicePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	CheckFatal(err, http.StatusInternalServerError, errors.New("no se pudo abrir el puerto serial"))
	defer file.Close()

	text := modo + "," + estado
	_, err = file.WriteString(text)
	CheckFatal(err, http.StatusInternalServerError, errors.New("no se pudo escribir en el puerto serial"))

	// Query handling
	result := database.Connection.Model(&models.Tendedero{}).Where("id = ?", os.Getenv("DEVICEID")).Update("estado", estado).Update("modo", modo)
	CheckFatal(result.Error, http.StatusInternalServerError, errors.New("no se pudo actualizar el tendedero"))

	c.JSON(http.StatusOK, gin.H{"mensaje": "Actualizado exitoso"})
}
