package helpers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetURLParam(c *gin.Context, param string) (string, error) {
	value, exists := c.Params.Get(param)
	if !exists {
		return "", HTTPError{http.StatusBadRequest, "No se pudo obtener el parametro: " + param}
	}
	return value, nil
}
