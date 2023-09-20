package main

import (
	"net/http"

	conectorControlador "example/fleetwise/ConectorBD/controlador"
	rutas "example/fleetwise/Rutas"

	"github.com/gin-gonic/gin"
)

func init() {
	conectorControlador.GuardarVehiculo()
}

func main() {

	r := gin.Default()

	rutas.VehiculoRouter(r)

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello world from server Go.",
		})
	})
	r.Run()
}
