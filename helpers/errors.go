package helpers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type HTTPError struct {
	StatusCode int
	Message    string
}

func (ce HTTPError) Error() string {
	return ce.Message
}

func Fatal(code int, err error) {
	panic(HTTPError{code, err.Error()})
}

func CheckFatal(checked error, code int, err error) {
	if checked != nil {
		panic(HTTPError{code, err.Error()})
	}
}

func HTTPRecoveryHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				if httpError, ok := err.(HTTPError); ok {
					// Log the HTTP error message
					log.Println("HTTP Error:", httpError)

					// Return the HTTP error as the response
					c.JSON(httpError.StatusCode, gin.H{"error": httpError.Message})
				} else {
					// Log the panic message
					log.Println("Panic:", err)

					// For other panics, return a generic error response
					c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
				}

				// Abort the current request
				c.Abort()
			}
		}()

		// Continue processing the request
		c.Next()
	}
}
