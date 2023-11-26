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

func (m *Manejador) SeleccionarVehiculoParaNuevoRegistro(contexto *gin.Context) {
	solicitud := m.RegistroMantenimientoVehiculoControlador.RegistroMantenimientoVehiculoMapeador.GinContextASeleccionarVehiculoParaNuevoRegistroSolicitud(contexto)
	respuesta := m.RegistroMantenimientoVehiculoControlador.SeleccionarVehiculoParaNuevoRegistro(solicitud)
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
	solicitud := m.RegistroMantenimientoVehiculoControlador.RegistroMantenimientoVehiculoMapeador.GinContextAAgregarRegistroMantemientoVehiculoSolicitud(contexto)
	respuesta := m.RegistroMantenimientoVehiculoControlador.AgregarRegistroMantemientoVehiculo(solicitud)
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
	respuesta := m.RegistroMantenimientoVehiculoControlador.ObtenerServiciosVehicularesParaNuevoRegistro()
	contexto.IndentedJSON(http.StatusOK, gin.H{"ServiciosVehiculares": respuesta})
}

func (m *Manejador) ObtenerRegistroMantenimientoVehicular(contexto *gin.Context){
	solicitud := m.RegistroMantenimientoVehiculoControlador.RegistroMantenimientoVehiculoMapeador.GinContextAObtenerRegistroMantenimientoVehiculoSolicitud(contexto)
	registro, vehiculo := m.RegistroMantenimientoVehiculoControlador.ObtenerRegistroMantenimientoVehicular(solicitud)
	
	contexto.IndentedJSON(http.StatusOK,gin.H{"registro": registro, "vehiculo": vehiculo})
}
