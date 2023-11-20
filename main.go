package main

import (
	vehiculosManejador "example/fleetwise/fuente/capturavehiculos/manejador"
	"log"

	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

type ControladorMain struct {
	VehiculosManejador *vehiculosManejador.Manejador
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
		VehiculosManejador: vehiculosManejador.NuevoManejador(),
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
	router.Run("localhost:8080")
}
