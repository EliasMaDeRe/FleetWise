package main

import (
	vehiculosManejador "example/fleetwise/fuente/capturavehiculos/manejador"

	"net/http"

	"github.com/gin-gonic/gin"
)

type ControladorMain struct {
	VehiculosManejador *vehiculosManejador.Manejador
}

func main() {
	controlador := &ControladorMain{
		VehiculosManejador: vehiculosManejador.NuevoManejador(),
	}
	router := gin.Default()
	router.LoadHTMLGlob("*.html")
	router.Static("/styles", "./styles/")
	router.Static("/js", "./js/")
	router.POST("/AgregarVehiculo", func(ctx *gin.Context) {
		controlador.VehiculosManejador.AgregarVehiculo(ctx)
	})
	router.GET("/AgregarVehiculo", func(c *gin.Context){
		c.HTML(http.StatusOK,"index.html",gin.H{})
	})
	router.Run("localhost:8080")
}
