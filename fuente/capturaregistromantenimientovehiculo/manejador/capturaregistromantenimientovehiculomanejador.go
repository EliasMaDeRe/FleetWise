package manejador

import (
	registroMantenimientoVehiculoControlador "example/fleetwise/fuente/capturaregistromantenimientovehiculo/controlador"
	registroMantenimientoVehiculoMapeador "example/fleetwise/fuente/capturaregistromantenimientovehiculo/mapeador"

	"net/http"

	"github.com/gin-gonic/gin"
)

type Manejador struct {
	RegistroMantenimientoVehiculoControlador *registroMantenimientoVehiculoControlador.Controlador
}

func NuevoManejador() (c *Manejador) {
	return &Manejador{
		RegistroMantenimientoVehiculoControlador: &registroMantenimientoVehiculoControlador.Controlador{
			RegistroMantenimientoVehiculoMapeador: &registroMantenimientoVehiculoMapeador.Mapeador{},
		},
	}
}

func (m *Manejador) SeleccionarVehiculoDeNuevoRegistro(contexto *gin.Context) {
	solicitud := m.RegistroMantenimientoVehiculoControlador.RegistroMantenimientoVehiculoMapeador.GinContextASeleccionarVehiculoDeNuevoRegistroSolicitud(contexto)
	respuesta := m.RegistroMantenimientoVehiculoControlador.SeleccionarVehiculoDeNuevoRegistro(solicitud)
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

func (m *Manejador) AgregarRegistroMantemientoVehiculo(contexto *gin.Context) {
	solicitud := m.RegistroMantenimientoVehiculoControlador.RegistroMantenimientoVehiculoMapeador.GinContextASeleccionarVehiculoDeNuevoRegistroSolicitud(contexto)
	respuesta := m.RegistroMantenimientoVehiculoControlador.SeleccionarVehiculoDeNuevoRegistro(solicitud)
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
