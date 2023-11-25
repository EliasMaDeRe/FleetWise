package main

import (
	servicioVehicularManejador "example/fleetwise/fuente/capturaserviciovehicular/manejador"
	vehiculosManejador "example/fleetwise/fuente/capturavehiculos/manejador"
	sesionManejador "example/fleetwise/fuente/iniciosesion/manejador"
	historialRegistrosManejador "example/fleetwise/fuente/visualizacionhistorialregistrosmantenimientovehicular/manejador"
	"log"

	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

type ControladorMain struct {
	VehiculosManejador         *vehiculosManejador.Manejador
	ServicioVehicularManejador *servicioVehicularManejador.Manejador
	SesionManejador            *sesionManejador.Manejador
	HistorialRegistrosManejador *historialRegistrosManejador.Manejador
}

func loadEnvFile() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file:", err)
	}
}

func main() {
	godotenv.Load()

	controlador := &ControladorMain{
		VehiculosManejador:         vehiculosManejador.NuevoManejador(),
		ServicioVehicularManejador: servicioVehicularManejador.NuevoManejador(),
		SesionManejador:            sesionManejador.NuevoManejador(),
		HistorialRegistrosManejador: historialRegistrosManejador.NuevoManejador(),
	}

	router := gin.Default()

	router.LoadHTMLGlob("*.html")
	router.Static("/styles", "./styles/")
	router.Static("/js", "./js/")
	router.Static("/img", "./img/")
	router.POST("/AgregarVehiculo", func(ctx *gin.Context) {
		controlador.SesionManejador.ValidarSesion(ctx)
	}, func(ctx *gin.Context) {
		controlador.VehiculosManejador.AgregarVehiculo(ctx)
	})
	router.GET("/AgregarVehiculo", func(ctx *gin.Context) {
		controlador.SesionManejador.ValidarSesion(ctx)
	}, func(c *gin.Context) {
		c.HTML(http.StatusOK, "vehiculos.html", gin.H{})
	})
	router.POST("/AgregarServicioVehicular", func(ctx *gin.Context) {
		controlador.SesionManejador.ValidarSesion(ctx)
	}, func(ctx *gin.Context) {
		controlador.ServicioVehicularManejador.AgregarServicioVehicular(ctx)
	})
	router.GET("/AgregarServicioVehicular", func(ctx *gin.Context) {
		controlador.SesionManejador.ValidarSesion(ctx)
	}, func(c *gin.Context) {
		c.HTML(http.StatusOK, "serviciovehicular.html", gin.H{})
	})
	router.POST("/EliminarServicioVehicular", func(ctx *gin.Context) {
		controlador.SesionManejador.ValidarSesion(ctx)
	}, func(ctx *gin.Context) {
		controlador.ServicioVehicularManejador.EliminarServicioVehicular(ctx)
	})

	router.GET("/ObtenerServiciosVehiculares", func(ctx *gin.Context) {
		controlador.SesionManejador.ValidarSesion(ctx)
	}, func(ctx *gin.Context) {
		controlador.ServicioVehicularManejador.ObtenerServiciosVehiculares(ctx)
	})

	router.POST("/EditarServicioVehicular", func(ctx *gin.Context) {
		controlador.SesionManejador.ValidarSesion(ctx)
	}, func(ctx *gin.Context) {
		controlador.ServicioVehicularManejador.EditarServicioVehicular(ctx)
	})
	router.GET("/Login", func(ctx *gin.Context) {
		controlador.SesionManejador.ValidarSesion(ctx)
	}, func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})
	router.POST("/Login", func(ctx *gin.Context) {
		controlador.SesionManejador.IniciarSesion(ctx)
	})
	router.GET("/", func(ctx *gin.Context) {
		controlador.SesionManejador.ValidarSesion(ctx)
	}, func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})
	router.GET("/Logout", func(ctx *gin.Context) {
		controlador.SesionManejador.ValidarSesion(ctx)
	}, func(ctx *gin.Context) {
		controlador.SesionManejador.CerrarSesion(ctx)
	}, func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", gin.H{})
	})
	
	router.POST("/HistorialRegistrosServicioVehicular", func(ctx *gin.Context) {
		controlador.HistorialRegistrosManejador.ObtenerRegistrosYVehiculosFiltrados(ctx);
	})

	router.Run("localhost:8080")
}
