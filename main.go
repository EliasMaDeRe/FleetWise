package main

import (
	capturaRegistroMantenimientoVehiculoManejador "example/fleetwise/fuente/capturaregistromantenimientovehiculo/manejador"
	capturaServicioVehicularManejador "example/fleetwise/fuente/capturaserviciovehicular/manejador"
	capturaVehiculosManejador "example/fleetwise/fuente/capturavehiculos/manejador"
	inicioSesionManejador "example/fleetwise/fuente/iniciosesion/manejador"
	visualizacionHistorialRegistrosManejador "example/fleetwise/fuente/visualizacionhistorialregistrosmantenimientovehiculo/manejador"
	visualizacionResumenMantenimientoVehiculosManejador "example/fleetwise/fuente/visualizacionresumenmantenimientovehiculos/manejador"
	"log"

	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

type ControladorMain struct {
	CapturaVehiculosManejador                           *capturaVehiculosManejador.Manejador
	CapturaServicioVehicularManejador                   *capturaServicioVehicularManejador.Manejador
	InicioSesionManejador                               *inicioSesionManejador.Manejador
	VisualizacionHistorialRegistrosManejador            *visualizacionHistorialRegistrosManejador.Manejador
	CapturaRegistroMantenimientoVehiculoManejador       *capturaRegistroMantenimientoVehiculoManejador.Manejador
	VisualizacionResumenMantenimientoVehiculosManejador *visualizacionResumenMantenimientoVehiculosManejador.Manejador
}

var fleetwise *ControladorMain = &ControladorMain{
	CapturaVehiculosManejador:                           capturaVehiculosManejador.NuevoManejador(),
	CapturaServicioVehicularManejador:                   capturaServicioVehicularManejador.NuevoManejador(),
	CapturaRegistroMantenimientoVehiculoManejador:       capturaRegistroMantenimientoVehiculoManejador.NuevoManejador(),
	InicioSesionManejador:                               inicioSesionManejador.NuevoManejador(),
	VisualizacionHistorialRegistrosManejador:            visualizacionHistorialRegistrosManejador.NuevoManejador(),
	VisualizacionResumenMantenimientoVehiculosManejador: visualizacionResumenMantenimientoVehiculosManejador.
	NuevoManejador(),
}

func loadEnvFile() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file:", err)
	}
}

func MiddleWareAuntenticacionCapturista(ctx *gin.Context){
	fleetwise.InicioSesionManejador.ValidarSesion(ctx, "capturista")
}

func MiddleWareAuntenticacionSupervisor(ctx *gin.Context){
	fleetwise.InicioSesionManejador.ValidarSesion(ctx, "supervisor")
}
func MiddleWareAuntenticacionAdministrador(ctx *gin.Context){
	fleetwise.InicioSesionManejador.ValidarSesion(ctx, "administrador")
}

func main() {
	loadEnvFile()

	router := gin.Default()

	// controlador.InicioSesionManejador.InicioSesionControlador.RegistrarUsuario("admin", "admin", "administrador")
	// controlador.InicioSesionManejador.InicioSesionControlador.RegistrarUsuario("supervisor", "supervisor", "supervisor")
	// controlador.InicioSesionManejador.InicioSesionControlador.RegistrarUsuario("capturista", "capturista", "capturista")

	router.LoadHTMLGlob("templates/**/*")
	router.Static("/styles", "./styles/")
	router.Static("/js", "./js/")
	router.Static("/img", "./img/")

	rutasAccesoMinimoCapturista := router.Group("/")
	
	rutasAccesoMinimoCapturista.Use(MiddleWareAuntenticacionCapturista)
	{

		rutasAccesoMinimoCapturista.GET("", func(ctx *gin.Context) {
			ctx.HTML(http.StatusOK, "index", gin.H{"title": "FleetWise | Inicio", "usuario":ctx.Keys["usuario"]})
		})
		
		// ====================================================Login
		rutasAccesoMinimoCapturista.GET("Login", func(ctx *gin.Context) {
			ctx.HTML(http.StatusOK, "index", gin.H{})
		})
		rutasAccesoMinimoCapturista.POST("Login", func(ctx *gin.Context) {
			fleetwise.InicioSesionManejador.IniciarSesion(ctx)
		})

		// ====================================================Logout
		rutasAccesoMinimoCapturista.GET("Logout", func(ctx *gin.Context) {
			fleetwise.InicioSesionManejador.CerrarSesion(ctx)
			ctx.Redirect(http.StatusFound, "/")
		})
		rutasAccesoMinimoCapturista.POST("Logout",func(ctx *gin.Context) {
			fleetwise.InicioSesionManejador.CerrarSesion(ctx)
		})

		// ====================================================Seleccionar vehiculo para crear registro de mantenimiento
		rutasAccesoMinimoCapturista.GET("SeleccionarVehiculoParaNuevoRegistro",func(c *gin.Context) {
			c.HTML(http.StatusOK, "registrosMantenimiento/seleccionarVehiculo", gin.H{"title": "FleetWise | Ingrese la placa de un vehiculo", "usuario":c.Keys["usuario"]})
		})
		rutasAccesoMinimoCapturista.POST("SeleccionarVehiculoParaNuevoRegistro",func(ctx *gin.Context) {
			fleetwise.CapturaRegistroMantenimientoVehiculoManejador.SeleccionarVehiculoParaNuevoRegistro(ctx)
		})
		
		// ==================================================== Captura Registro Mantenimiento Vehiculo
		rutasAccesoMinimoCapturista.GET("AgregarRegistroMantenimientoVehiculo", func(ctx *gin.Context) {
			fleetwise.CapturaRegistroMantenimientoVehiculoManejador.SeleccionarVehiculoParaNuevoRegistro(ctx)
		}, func(c *gin.Context) {
			c.HTML(http.StatusOK, "registrosMantenimiento/agregar", gin.H{"usuario":c.Keys["usuario"]})
		})
		rutasAccesoMinimoCapturista.POST("AgregarRegistroMantenimientoVehiculo", func(ctx *gin.Context) {
			fleetwise.CapturaRegistroMantenimientoVehiculoManejador.AgregarRegistroMantemientoVehiculo(ctx)
		})
		
		// ==================================================== Obtener Servicios Vehiculares existentes
		rutasAccesoMinimoCapturista.POST("ObtenerServiciosVehicularesParaNuevoRegistro", func(ctx *gin.Context) {
			fleetwise.CapturaRegistroMantenimientoVehiculoManejador.ObtenerServiciosVehiculares(ctx)
		})
	}

	rutasAccesoMinimoSupervisor := router.Group("/")
	rutasAccesoMinimoSupervisor.Use(MiddleWareAuntenticacionSupervisor)
	{
		// ==================================================== Visualizacion Historial Registro Mantenimiento Vehiculo
		rutasAccesoMinimoSupervisor.GET("HistorialRegistrosMantenimientoVehiculo",func(c *gin.Context) {
			c.HTML(http.StatusOK, "historialRegistros", gin.H{"title": "FleetWise | Historial de Registros de Mantenimiento", "usuario":c.Keys["usuario"]})
		})

		rutasAccesoMinimoSupervisor.POST("HistorialRegistrosMantenimientoVehiculo", func(ctx *gin.Context) {
			fleetwise.VisualizacionHistorialRegistrosManejador.ObtenerRegistrosYVehiculosFiltrados(ctx)
		})
	}

	rutasAccesoMinimoAdministrador := router.Group("/")
	rutasAccesoMinimoAdministrador.Use(MiddleWareAuntenticacionAdministrador)
	{
		// ==================================================== Captura vehiculos
		rutasAccesoMinimoAdministrador.GET("AgregarVehiculo", func(c *gin.Context) {
			c.HTML(http.StatusOK, "vehiculos/agregar", gin.H{"title": "FleetWise | Capturar Vehículo", "usuario":c.Keys["usuario"]})
		})
		rutasAccesoMinimoAdministrador.POST("AgregarVehiculo", func(ctx *gin.Context) {
			fleetwise.CapturaVehiculosManejador.AgregarVehiculo(ctx)
		})
		rutasAccesoMinimoAdministrador.DELETE("EliminarVehiculo", func(ctx *gin.Context) {
			fleetwise.CapturaVehiculosManejador.EliminarVehiculo(ctx)
		})

		// ==================================================== Captura servicio vehicular
		rutasAccesoMinimoAdministrador.GET("AgregarServicioVehicular", func(c *gin.Context) {
			c.HTML(http.StatusOK, "capturaServicioVehicular", gin.H{"title": "FleetWise | Captura Tipos de Servicios Vehiculares", "usuario":c.Keys["usuario"]})
		})
		rutasAccesoMinimoAdministrador.POST("AgregarServicioVehicular", func(ctx *gin.Context) {
			fleetwise.CapturaServicioVehicularManejador.AgregarServicioVehicular(ctx)
		})

		rutasAccesoMinimoAdministrador.GET("ObtenerServiciosVehiculares", func(ctx *gin.Context) {
			fleetwise.CapturaServicioVehicularManejador.ObtenerServiciosVehiculares(ctx)
		})
		
		rutasAccesoMinimoAdministrador.POST("EditarServicioVehicular", func(ctx *gin.Context) {
			fleetwise.CapturaServicioVehicularManejador.EditarServicioVehicular(ctx)
		})
		
		rutasAccesoMinimoAdministrador.POST("EliminarServicioVehicular", func(ctx *gin.Context) {
			fleetwise.CapturaServicioVehicularManejador.EliminarServicioVehicular(ctx)
		})

		rutasAccesoMinimoAdministrador.POST("ObtenerRegistroPorNumeroDeRegistro", func(ctx *gin.Context) {
			fleetwise.CapturaRegistroMantenimientoVehiculoManejador.ObtenerRegistroMantenimientoVehiculoPorNumeroDeRegistro(ctx)
		})

		rutasAccesoMinimoAdministrador.GET("EditarRegistroMantenimientoVehiculo", func(c *gin.Context) {
			c.HTML(http.StatusOK, "registrosMantenimiento/editar", gin.H{"title": "FleetWise | Editar Registro de Mantenimiento Vehicular", "usuario":c.Keys["usuario"]})
		})

		rutasAccesoMinimoAdministrador.POST("EditarRegistroMantenimientoVehiculo", func(ctx *gin.Context) {
			fleetwise.CapturaRegistroMantenimientoVehiculoManejador.EditarRegistroDeMantenimientoDelVehiculo(ctx)
		})

		rutasAccesoMinimoAdministrador.GET("RegistroEditado", func(c *gin.Context) {
			c.HTML(http.StatusOK, "registrosMantenimiento/registroEditado", gin.H{"title": "FleetWise | Registro Editado Correctamente", "usuario":c.Keys["usuario"]})
		})

		// ==================================================== Visualizacion Resumen Mantenimiento Vehiculos
		rutasAccesoMinimoAdministrador.GET("ResumenMantenimientoVehiculos", func(c *gin.Context) {
			c.HTML(http.StatusOK, "resumenVehiculos", gin.H{"title": "FleetWise | Resumen Mantenimiento Vehiculos", "usuario":c.Keys["usuario"]})
		})

		rutasAccesoMinimoAdministrador.POST("ObtenerMetricasVehiculos", func(ctx *gin.Context) {
			fleetwise.VisualizacionResumenMantenimientoVehiculosManejador.ObtenerMetricasVehiculares(ctx)
		})

		rutasAccesoMinimoAdministrador.POST("ObtenerVehiculoPorPlacas", func(ctx *gin.Context) {
			fleetwise.CapturaVehiculosManejador.ObtenerVehiculoPorPlacas(ctx)
		})

		rutasAccesoMinimoAdministrador.POST("EditarVehiculo", func(ctx *gin.Context) {
			fleetwise.CapturaVehiculosManejador.EditarVehiculo(ctx)
		})

		rutasAccesoMinimoAdministrador.GET("EditarVehiculo", func(c *gin.Context) {
			c.HTML(http.StatusOK, "vehiculos/editar", gin.H{"title": "FleetWise | Editar Vehículo", "usuario":c.Keys["usuario"]})
		})
	}

	router.Run("localhost:8080")
}

