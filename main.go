package main

import (
	registroMantenimientoVehicularManejador "example/fleetwise/fuente/capturaregistromantenimientovehiculo/manejador"
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
	RegistroMantenimientoVehicularManejador *registroMantenimientoVehicularManejador.Manejador
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
		HistorialRegistrosManejador: historialRegistrosManejador.NuevoManejador(),
	}

	router := gin.Default()

	controlador.SesionManejador.InicioSesionControlador.RegistrarUsuario("Fer", "123")

	router.LoadHTMLGlob("*.html")
	router.Static("/styles", "./styles/")
	router.Static("/js", "./js/")
	router.Static("/img", "./img/")

	//Login
	router.GET("/", func(ctx *gin.Context) {
		controlador.SesionManejador.ValidarSesion(ctx)
	}, func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})
	router.GET("/Login", func(ctx *gin.Context) {
		controlador.SesionManejador.ValidarSesion(ctx)
	}, func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})
	router.POST("/Login", func(ctx *gin.Context) {
		controlador.SesionManejador.IniciarSesion(ctx)
	})
	router.GET("/Logout", func(ctx *gin.Context) {
		controlador.SesionManejador.ValidarSesion(ctx)
	}, func(ctx *gin.Context) {
		controlador.SesionManejador.CerrarSesion(ctx)
	}, func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", gin.H{})
	})

	//Captura vehiculos
	router.GET("/AgregarVehiculo", func(ctx *gin.Context) {
		controlador.SesionManejador.ValidarSesion(ctx)
	}, func(c *gin.Context) {
		c.HTML(http.StatusOK, "vehiculos.html", gin.H{})
	})
	router.POST("/AgregarVehiculo", func(ctx *gin.Context) {
		controlador.SesionManejador.ValidarSesion(ctx)
	}, func(ctx *gin.Context) {
		controlador.VehiculosManejador.AgregarVehiculo(ctx)
	})
	
	// Captura servicio vehicular
	router.GET("/AgregarServicioVehicular", func(ctx *gin.Context) {
		controlador.SesionManejador.ValidarSesion(ctx)
	}, func(c *gin.Context) {
		c.HTML(http.StatusOK, "serviciovehicular.html", gin.H{})
	})
	router.POST("/AgregarServicioVehicular", func(ctx *gin.Context) {
		controlador.SesionManejador.ValidarSesion(ctx)
	}, func(ctx *gin.Context) {
		controlador.ServicioVehicularManejador.AgregarServicioVehicular(ctx)
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
	router.POST("/EliminarServicioVehicular", func(ctx *gin.Context) {
		controlador.SesionManejador.ValidarSesion(ctx)
	}, func(ctx *gin.Context) {
		controlador.ServicioVehicularManejador.EliminarServicioVehicular(ctx)
	})

	

	// Captura registro mantenimiento vehicular
	router.GET("/AgregarRegistroMantenimientoVehicular", func(c *gin.Context) {
		c.HTML(http.StatusOK, "registromantenimientovehiculo.html",gin.H{})
	})

	router.POST("/AgregarRegistroMantenimientoVehicular", func(ctx *gin.Context) {
		controlador.RegistroMantenimientoVehicularManejador.AgregarRegistroMantemientoVehiculo(ctx)
	})
	
	router.GET("/SeleccionarVehiculoParaNuevoRegistro", func(ctx *gin.Context) {
		controlador.RegistroMantenimientoVehicularManejador.SeleccionarVehiculoParaNuevoRegistro(ctx)
	})
	
	router.POST("/ObtenerServiciosVehicularesParaNuevoRegistro", func(ctx *gin.Context) {
		controlador.RegistroMantenimientoVehicularManejador.ObtenerServiciosVehiculares(ctx)
	})

	

	// Visualizar historial registros mantenimiento vehicular
	router.GET("/HistorialRegistrosMantenimientoVehicular", func(c *gin.Context) {
		c.HTML(http.StatusOK, "historialregistros.html", gin.H{})
	})

	router.POST("/HistorialRegistrosMantenimientoVehicular", func(ctx *gin.Context) {
		controlador.HistorialRegistrosManejador.ObtenerRegistrosYVehiculosFiltrados(ctx);
	})

	router.GET("/EditarRegistroMantenimientoVehicular", func(c *gin.Context) {
		c.HTML(http.StatusOK, "editarRegistro.html", gin.H{})
	})

	router.POST("/ObtenerRegistroPorNumeroDeRegistro", func(ctx *gin.Context) {
		controlador.RegistroMantenimientoVehicularManejador.ObtenerRegistroMantenimientoVehicularPorNumeroDeRegistro(ctx)
	})


	router.Run("localhost:8080")
}
