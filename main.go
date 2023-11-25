package main

import (
	registroMantenimientoVehicularManejador "example/fleetwise/fuente/capturaregistromantenimientovehiculo/manejador"
	servicioVehicularManejador "example/fleetwise/fuente/capturaserviciovehicular/manejador"
	vehiculosManejador "example/fleetwise/fuente/capturavehiculos/manejador"
	sesionManejador "example/fleetwise/fuente/iniciosesion/manejador"
	"log"

	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

type ControladorMain struct {
	VehiculosManejador                      *vehiculosManejador.Manejador
	ServicioVehicularManejador              *servicioVehicularManejador.Manejador
	RegistroMantenimientoVehicularManejador *registroMantenimientoVehicularManejador.Manejador
	SesionManejador                         *sesionManejador.Manejador
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
		VehiculosManejador:                      vehiculosManejador.NuevoManejador(),
		ServicioVehicularManejador:              servicioVehicularManejador.NuevoManejador(),
		RegistroMantenimientoVehicularManejador: registroMantenimientoVehicularManejador.NuevoManejador(),
		SesionManejador:                         sesionManejador.NuevoManejador(),
	}

	router := gin.Default()

	controlador.SesionManejador.InicioSesionControlador.RegistrarUsuario("Fer", "123")

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

	router.GET("/HistorialRegistrosMantenimientoVehicular", func(c *gin.Context) {
		c.HTML(http.StatusOK, "historialregistros.html", gin.H{})
	})

	router.GET("/EditarRegistroMantenimientoVehicular", func(c *gin.Context) {
		c.HTML(http.StatusOK, "editarRegistro.html", gin.H{})
	})
	router.GET("/AgregarRegistroMantenimientoVehicular", func(ctx *gin.Context) {
		controlador.RegistroMantenimientoVehicularManejador.AgregarRegistroMantemientoVehiculo(ctx)
	})
	router.GET("/SeleccionarVehiculoParaNuevoRegistro", func(ctx *gin.Context) {
		controlador.RegistroMantenimientoVehicularManejador.SeleccionarVehiculoParaNuevoRegistro(ctx)
	})
	router.POST("/ObtenerServiciosVehiculareParaNuevoRegistro", func(ctx *gin.Context) {
		controlador.RegistroMantenimientoVehicularManejador.ObtenerServiciosVehiculares(ctx)
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
	router.Run("localhost:8080")
}
