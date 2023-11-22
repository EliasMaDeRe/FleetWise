package main

import (
	servicioVehicularManejador "example/fleetwise/fuente/capturaserviciovehicular/manejador"
	vehiculosManejador "example/fleetwise/fuente/capturavehiculos/manejador"
	"log"

	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

type ControladorMain struct {
	VehiculosManejador         *vehiculosManejador.Manejador
	ServicioVehicularManejador *servicioVehicularManejador.Manejador
}

func loadEnvFile(){
	err := godotenv.Load();
	if err != nil {
		log.Fatal("Error loading .env file:", err)
	}
}

func main() {
	godotenv.Load()
	

	controlador := &ControladorMain{
		VehiculosManejador:         vehiculosManejador.NuevoManejador(),
		ServicioVehicularManejador: servicioVehicularManejador.NuevoManejador(),
	}
	router := gin.Default()
	
	router.LoadHTMLGlob("*.html")
	router.Static("/styles", "./styles/")
	router.Static("/js", "./js/")
	router.POST("/AgregarVehiculo", func(ctx *gin.Context) {
		controlador.VehiculosManejador.AgregarVehiculo(ctx)
	})
	router.GET("/AgregarVehiculo", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})
	router.POST("/AgregarServicioVehicular", func(ctx *gin.Context) {
		controlador.ServicioVehicularManejador.AgregarServicioVehicular(ctx)
	})
	router.GET("/AgregarServicioVehicular", func(c *gin.Context) {
		c.HTML(http.StatusOK, "serviciovehicular.html", gin.H{})
	})
	router.POST("/EliminarServicioVehicular", func(ctx *gin.Context) {
		controlador.ServicioVehicularManejador.EliminarServicioVehicular(ctx)
	})
	router.GET("/ObtenerServiciosVehiculares", func(ctx *gin.Context) {
		controlador.ServicioVehicularManejador.ObtenerServiciosVehiculares(ctx)
	})
	router.POST("/EditarServicioVehicular", func(ctx *gin.Context) {
		controlador.ServicioVehicularManejador.EditarServicioVehicular(ctx)
	})
	router.Run("localhost:8080")
}
