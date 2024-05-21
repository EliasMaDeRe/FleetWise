package manejador

import (
	capturaRegistroMantenimientoVehiculoControlador "example/fleetwise/fuente/capturaregistromantenimientovehiculo/controlador"
	capturaRegistroMantenimientoVehiculoMapeador "example/fleetwise/fuente/capturaregistromantenimientovehiculo/mapeador"
	capturaServicioVehicularControlador "example/fleetwise/fuente/capturaserviciovehicular/controlador"
	capturaVehiculosControlador "example/fleetwise/fuente/capturavehiculos/controlador"
	conectorBDControlador "example/fleetwise/fuente/conectorbdcapturaregistromantenimientovehiculo/controlador"

	"strings"

	"net/http"

	"github.com/gin-gonic/gin"
)

type Manejador struct {
	CapturaRegistroMantenimientoVehiculoControlador *capturaRegistroMantenimientoVehiculoControlador.Controlador
}

func NuevoManejador() (c *Manejador) {
	return &Manejador{
		CapturaRegistroMantenimientoVehiculoControlador: &capturaRegistroMantenimientoVehiculoControlador.Controlador{
			CapturaRegistroMantenimientoVehiculoMapeador: &capturaRegistroMantenimientoVehiculoMapeador.Mapeador{},
			ConectorBDControlador:                        &conectorBDControlador.Controlador{},
			CapturaServicioVehicularControlador:          &capturaServicioVehicularControlador.Controlador{},
			CapturaVehiculosControlador:          &capturaVehiculosControlador.Controlador{},
		},
	}
}

func (m *Manejador) SeleccionarVehiculoParaNuevoRegistro(contexto *gin.Context) {
	solicitud := m.CapturaRegistroMantenimientoVehiculoControlador.CapturaRegistroMantenimientoVehiculoMapeador.GinContextASeleccionarVehiculoParaNuevoRegistroSolicitud(contexto)
	respuesta := m.CapturaRegistroMantenimientoVehiculoControlador.SeleccionarVehiculoParaNuevoRegistro(solicitud)
	status := http.StatusOK
	if !respuesta.ObtenerOk() {
		status = http.StatusBadRequest
	}
	if strings.Contains(contexto.Request.URL.String(), "AgregarRegistroMantenimientoVehiculo") {
		if respuesta.ObtenerOk() {
			contexto.Next()
			return
		} else {
			contexto.HTML(http.StatusOK, "seleccionarvehiculo.html", gin.H{})
			contexto.Abort()
			return
		}
	}
	var mensajeError string
	if respuesta.ObtenerErr() != nil {
		mensajeError = respuesta.ObtenerErr().Error()
	}
	contexto.IndentedJSON(status, gin.H{"OK": respuesta.ObtenerOk(), "err": mensajeError})
}

func (m *Manejador) AgregarRegistroMantemientoVehiculo(contexto *gin.Context) {
	solicitud := m.CapturaRegistroMantenimientoVehiculoControlador.CapturaRegistroMantenimientoVehiculoMapeador.GinContextAAgregarRegistroMantemientoVehiculoSolicitud(contexto)
	respuesta := m.CapturaRegistroMantenimientoVehiculoControlador.AgregarRegistroMantemientoVehiculo(solicitud)
	status := http.StatusOK
	if !respuesta.ObtenerOk() {
		status = http.StatusBadRequest
	}
	var mensajeError string
	if respuesta.ObtenerErr() != nil {
		mensajeError = respuesta.ObtenerErr().Error()
	}
	contexto.IndentedJSON(status, gin.H{"OK": respuesta.ObtenerOk(), "err": mensajeError})
}

func (m *Manejador) ObtenerServiciosVehiculares(contexto *gin.Context) {
	respuesta := m.CapturaRegistroMantenimientoVehiculoControlador.ObtenerServiciosVehicularesParaNuevoRegistro()
	contexto.IndentedJSON(http.StatusOK, gin.H{"ServiciosVehiculares": respuesta})
}

func (m *Manejador) ObtenerRegistroMantenimientoVehiculoPorNumeroDeRegistro(contexto *gin.Context) {
	solicitud := m.CapturaRegistroMantenimientoVehiculoControlador.CapturaRegistroMantenimientoVehiculoMapeador.GinContextAObtenerRegistroMantenimientoVehiculoPorNumeroDeRegistroSolicitud(contexto)
	registro, vehiculo := m.CapturaRegistroMantenimientoVehiculoControlador.ObtenerRegistroMantenimientoVehiculo(solicitud)

	contexto.IndentedJSON(http.StatusOK, gin.H{"registro": registro, "vehiculo": vehiculo})
}

func (m *Manejador) EditarRegistroDeMantenimientoDelVehiculo(contexto *gin.Context) {
	solicitud := m.CapturaRegistroMantenimientoVehiculoControlador.CapturaRegistroMantenimientoVehiculoMapeador.GinContextAEditarRegistroDeMantenimientoDelVehiculoSolicitud(contexto)

	respuesta := m.CapturaRegistroMantenimientoVehiculoControlador.EditarRegistroDeMantenimientoDeVehiculo(solicitud)

	status := http.StatusOK
	if !respuesta.ObtenerOk() {
		status = http.StatusBadRequest
	}

	mensajeError := ""
	if respuesta.ObtenerErr() != nil {
		mensajeError = respuesta.ObtenerErr().Error()
	}

	contexto.IndentedJSON(status, gin.H{"OK": respuesta.ObtenerOk(), "mensajeError": mensajeError})
}
