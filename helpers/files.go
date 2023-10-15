package helpers

import (
	"os"
)

func WriteToFile(path string, message string) {
	file, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		panic("no se pudo abrir el puerto serial")
	}
	defer file.Close()

	_, err = file.WriteString(message)
	if err != nil {
		panic("no se pudo escribir en el puerto serial")
	}
}
