package manejador

import (
	capturaServicioVehicularControlador "example/fleetwise/fuente/capturaserviciovehicular/controlador"
	capturaServicioVehicularMapeador "example/fleetwise/fuente/capturaserviciovehicular/mapeador"
	conectorBDControlador "example/fleetwise/fuente/conectorbd/controlador"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Manejador struct {
	ServicioVehicularControlador *capturaServicioVehicularControlador.Controlador
}

func NuevoManejador() (c *Manejador) {
	return &Manejador{
		ServicioVehicularControlador: &capturaServicioVehicularControlador.Controlador{
			CapturaServicioVehicularMapeador: &capturaServicioVehicularMapeador.Mapeador{},
			ConectorBDControlador:     &conectorBDControlador.Controlador{},
		},
	}
}

func (m *Manejador) AgregarServicioVehicular(contexto *gin.Context) {
	solicitud := m.ServicioVehicularControlador.CapturaServicioVehicularMapeador.GinContextAAgregarServicioVehicularSolicitud(contexto)
	respuesta := m.ServicioVehicularControlador.AgregarServicioVehicular(solicitud)
	status := http.StatusOK
	if respuesta.ObtenerOk() == false {
		status = http.StatusBadRequest
	}
	var mensajeError string
	if respuesta.ObtenerErr() != nil {
		mensajeError = respuesta.ObtenerErr().Error()
	}
	contexto.IndentedJSON(status, gin.H{"OK": respuesta.ObtenerOk(), "err": mensajeError})
}

func (m *Manejador) EliminarServicioVehicular(contexto *gin.Context) {
	solicitud := m.ServicioVehicularControlador.CapturaServicioVehicularMapeador.GinContextAEliminarServicioVehicularSolicitud(contexto)
	respuesta := m.ServicioVehicularControlador.EliminarServicioVehicular(solicitud)
	status := http.StatusOK
	if respuesta.ObtenerOk() == false {
		status = http.StatusBadRequest
	}
	var mensajeError string
	if respuesta.ObtenerErr() != nil {
		mensajeError = respuesta.ObtenerErr().Error()
	}
	contexto.IndentedJSON(status, gin.H{"OK": respuesta.ObtenerOk(), "err": mensajeError})
}

func (m *Manejador) EditarServicioVehicular(contexto *gin.Context) {
	solicitud := m.ServicioVehicularControlador.CapturaServicioVehicularMapeador.GinContextAEditarServicioVehicularSolicitud(contexto)
	respuesta := m.ServicioVehicularControlador.EditarServicioVehicular(solicitud)
	status := http.StatusOK
	if respuesta.ObtenerOk() == false {
		status = http.StatusBadRequest
	}
	var mensajeError string
	if respuesta.ObtenerErr() != nil {
		mensajeError = respuesta.ObtenerErr().Error()
	}
	contexto.IndentedJSON(status, gin.H{"OK": respuesta.ObtenerOk(), "err": mensajeError})
}

func (m *Manejador) ObtenerServiciosVehiculares(contexto *gin.Context) {
	respuesta := m.ServicioVehicularControlador.ObtenerServiciosVehiculares()
	contexto.IndentedJSON(http.StatusOK, gin.H{"ServiciosVehiculares": respuesta})
}
