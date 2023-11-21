# smart-clothesline-http

Servidor http para control del dispositivo Smart Clothesline

Por defecto se utilizó la IP dada según una VPN, por medio del puerto 8080
Utilizando Go, Go Fiber para la gestión del servidor HTTP y gorm para comunicación con la base de datos en Postgres

# Rutas

POST /login - permite la autenticación del usuario por contraseña, devuelve un objeto JSON con la llave JWT correspondiente

POST /signup - permite el registro de usuario

GET /tendederos - obtiene la información de todos los dispositivos a nombre del usuario

GET /tendedero/:device_id - obtiene la información de un dispositivo según el id correspondiente a nombre del usuario

PATCH /tendedero/:device_id/:modo/:estado - actualiza el estado (0 o 1) y modo (0 o 1) del dispositivo para cambiarlo según:
    modo automático (0) o manual (1)
    estado adentro (0) o afuera (1)


