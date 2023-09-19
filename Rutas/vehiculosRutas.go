package routes

import (
	"github.com/gin-gonic/gin"

	controlador "example/fleetwise/ConectorBD/Controlador"
)

func VehiculoRouter(router *gin.Engine) {

	routes := router.Group("api/v1/vehiculos")
	routes.POST("", controlador.GuardarVehiculo)

}
