package main

import (
	capturaRegistroMantenimientoVehiculoManejador "example/fleetwise/fuente/capturaregistromantenimientovehiculo/manejador"
	capturaServicioVehicularManejador "example/fleetwise/fuente/capturaserviciovehicular/manejador"
	capturaVehiculosManejador "example/fleetwise/fuente/capturavehiculos/manejador"
	inicioSesionManejador "example/fleetwise/fuente/iniciosesion/manejador"
	visualizacionHistorialRegistrosManejador "example/fleetwise/fuente/visualizacionhistorialregistrosmantenimientovehiculo/manejador"
	"log"

	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

type ControladorMain struct {
	CapturaVehiculosManejador         *capturaVehiculosManejador.Manejador
	CapturaServicioVehicularManejador *capturaServicioVehicularManejador.Manejador
	InicioSesionManejador            *inicioSesionManejador.Manejador
	VisualizacionHistorialRegistrosManejador *visualizacionHistorialRegistrosManejador.Manejador
	CapturaRegistroMantenimientoVehiculoManejador *capturaRegistroMantenimientoVehiculoManejador.Manejador
}

func loadEnvFile() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file:", err)
	}
}

func main() {
	loadEnvFile()

	controlador := &ControladorMain{
		CapturaVehiculosManejador:                      capturaVehiculosManejador.NuevoManejador(),
		CapturaServicioVehicularManejador:              capturaServicioVehicularManejador.NuevoManejador(),
		CapturaRegistroMantenimientoVehiculoManejador: capturaRegistroMantenimientoVehiculoManejador.NuevoManejador(),
		InicioSesionManejador:                         inicioSesionManejador.NuevoManejador(),
		VisualizacionHistorialRegistrosManejador: visualizacionHistorialRegistrosManejador.NuevoManejador(),
	}

	router := gin.Default()

	router.LoadHTMLGlob("*.html")
	router.Static("/styles", "./styles/")
	router.Static("/js", "./js/")
	router.Static("/img", "./img/")
	

	// Login
	
	router.GET("/", func(ctx *gin.Context) {
		controlador.InicioSesionManejador.ValidarSesion(ctx, "capturista")
	}, func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})
	
	router.GET("/Login", func(ctx *gin.Context) {
		controlador.InicioSesionManejador.ValidarSesion(ctx, "capturista")
	}, func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})
	router.POST("/Login", func(ctx *gin.Context) {
		controlador.InicioSesionManejador.IniciarSesion(ctx)
	})

	router.POST("/Logout", func(ctx *gin.Context) {
		controlador.InicioSesionManejador.ValidarSesion(ctx, "administrador")
	}, func(ctx *gin.Context) {
		controlador.InicioSesionManejador.CerrarSesion(ctx)
	}, func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", gin.H{})
	})

	router.GET("/Logout", func(ctx *gin.Context) {
		controlador.InicioSesionManejador.ValidarSesion(ctx, "capturista")
	}, func(ctx *gin.Context) {
		controlador.InicioSesionManejador.CerrarSesion(ctx)
	}, func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", gin.H{})
	})

	//Captura vehiculos
	router.GET("/AgregarVehiculo", func(ctx *gin.Context) {
		controlador.InicioSesionManejador.ValidarSesion(ctx, "administrador")
	}, func(c *gin.Context) {
		c.HTML(http.StatusOK, "vehiculos.html", gin.H{})
	})
	router.POST("/AgregarVehiculo", func(ctx *gin.Context) {
		controlador.InicioSesionManejador.ValidarSesion(ctx, "administrador")
	}, func(ctx *gin.Context) {
		controlador.CapturaVehiculosManejador.AgregarVehiculo(ctx)
	})
	
	// Captura servicio vehicular
	router.GET("/AgregarServicioVehicular", func(ctx *gin.Context) {
		controlador.InicioSesionManejador.ValidarSesion(ctx, "administrador")
	}, func(c *gin.Context) {
		c.HTML(http.StatusOK, "serviciovehicular.html", gin.H{})
	})
	router.POST("/AgregarServicioVehicular", func(ctx *gin.Context) {
		controlador.InicioSesionManejador.ValidarSesion(ctx, "administrador")
	}, func(ctx *gin.Context) {
		controlador.CapturaServicioVehicularManejador.AgregarServicioVehicular(ctx)
	})
	router.GET("/ObtenerServiciosVehiculares", func(ctx *gin.Context) {
		controlador.InicioSesionManejador.ValidarSesion(ctx, "administrador")
	}, func(ctx *gin.Context) {
		controlador.CapturaServicioVehicularManejador.ObtenerServiciosVehiculares(ctx)
	})
	router.POST("/EditarServicioVehicular", func(ctx *gin.Context) {
		controlador.InicioSesionManejador.ValidarSesion(ctx, "administrador")
	}, func(ctx *gin.Context) {
		controlador.CapturaServicioVehicularManejador.EditarServicioVehicular(ctx)
	})
	router.POST("/EliminarServicioVehicular", func(ctx *gin.Context) {
		controlador.InicioSesionManejador.ValidarSesion(ctx, "administrador")
	}, func(ctx *gin.Context) {
		controlador.CapturaServicioVehicularManejador.EliminarServicioVehicular(ctx)
	})

	// Captura Registro Mantenimiento Vehiculo

	router.GET("/AgregarRegistroMantenimientoVehiculo", func(ctx *gin.Context) {
		controlador.InicioSesionManejador.ValidarSesion(ctx, "capturista")
	}, func(ctx *gin.Context) {
		controlador.CapturaRegistroMantenimientoVehiculoManejador.SeleccionarVehiculoParaNuevoRegistro(ctx)
	}, func(c *gin.Context) {
		c.HTML(http.StatusOK, "registromantenimientovehiculo.html", gin.H{})
	})

	router.POST("/AgregarRegistroMantenimientoVehicular", func(ctx *gin.Context) {
		controlador.InicioSesionManejador.ValidarSesion(ctx, "capturista")
	}, func(ctx *gin.Context) {
		controlador.CapturaRegistroMantenimientoVehiculoManejador.AgregarRegistroMantemientoVehiculo(ctx)
	})

	router.POST("/ObtenerServiciosVehicularesParaNuevoRegistro", func(ctx *gin.Context) {
		controlador.InicioSesionManejador.ValidarSesion(ctx, "capturista")
	}, func(ctx *gin.Context) {
		controlador.CapturaRegistroMantenimientoVehiculoManejador.ObtenerServiciosVehiculares(ctx)
	})

	router.GET("/SeleccionarVehiculoParaNuevoRegistro", func(ctx *gin.Context) {
		controlador.InicioSesionManejador.ValidarSesion(ctx, "capturista")
	}, func(c *gin.Context) {
		c.HTML(http.StatusOK, "seleccionarvehiculo.html", gin.H{})
	})

	router.POST("/SeleccionarVehiculoParaNuevoRegistro", func(ctx *gin.Context) {
		controlador.InicioSesionManejador.ValidarSesion(ctx, "capturista")
	}, func(ctx *gin.Context) {
		controlador.CapturaRegistroMantenimientoVehiculoManejador.SeleccionarVehiculoParaNuevoRegistro(ctx)
	})

	router.POST("/ObtenerRegistroPorNumeroDeRegistro", func(ctx *gin.Context) {
		controlador.CapturaRegistroMantenimientoVehiculoManejador.ObtenerRegistroMantenimientoVehicularPorNumeroDeRegistro(ctx)
	})	

	// Visualizacion Historial Registro Mantenimiento Vehiculo
	router.GET("/HistorialRegistrosMantenimientoVehicular", func(c *gin.Context) {
		c.HTML(http.StatusOK, "historialregistros.html", gin.H{})
	})

	router.POST("/HistorialRegistrosMantenimientoVehicular", func(ctx *gin.Context) {
		controlador.VisualizacionHistorialRegistrosManejador.ObtenerRegistrosYVehiculosFiltrados(ctx);
	})
	
	router.GET("/EditarRegistroMantenimientoVehicular", func(c *gin.Context) {
		c.HTML(http.StatusOK, "editarRegistro.html", gin.H{})
	})

	router.POST("/EditarRegistroMantenimientoVehicular",func(ctx *gin.Context) {
		controlador.CapturaRegistroMantenimientoVehiculoManejador.EditarRegistroDeMantenimientoDelVehiculo(ctx)
	})

	router.GET("/RegistroEditado", func(c *gin.Context) {
		c.HTML(http.StatusOK, "registroEditado.html", gin.H{})
	})
	
	router.Run("localhost:8080")
}
