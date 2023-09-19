package fleetwise

import (
	vehiculosManejador "example/fleetwise/fuente/vehiculos/manejador"

	"github.com/gin-gonic/gin"
	"net/http"
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
	router.POST("/AgregarVehiculo", controlador.VehiculosManejador.AgregarVehiculo)
	router.GET("/AgregarVehiculo", func(c *gin.Context){
		c.HTML(http.StatusOK,"index.html",gin.H{"title":"Agregar Vehículo"})
	})
	router.Run("localhost:8080")
}
