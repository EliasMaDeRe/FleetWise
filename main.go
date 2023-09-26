package main

import (
	vehiculosManejador "example/fleetwise/fuente/capturavehiculos/manejador"

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
	router.POST("/AgregarVehiculo", controlador.VehiculosManejador.AgregarVehiculo)
	router.Run("localhost:8080")
}
