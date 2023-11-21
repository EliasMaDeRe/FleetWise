package manejador

import (
	servicioVehicularControlador "example/fleetwise/fuente/capturaserviciovehicular/controlador"
	servicioVehicularMapeador "example/fleetwise/fuente/capturaserviciovehicular/mapeador"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Manejador struct {
	ServicioVehicularControlador *servicioVehicularControlador.Controlador
}

func NuevoManejador() (c *Manejador) {
	return &Manejador{
		ServicioVehicularControlador: &servicioVehicularControlador.Controlador{
			ServicioVehicularMapeador: &servicioVehicularMapeador.Mapeador{},
		},
	}
}

func (m *Manejador) AgregarServicioVehicular(contexto *gin.Context) {
	solicitud := m.ServicioVehicularControlador.ServicioVehicularMapeador.GinContextAAgregarServicioVehicularSolicitud(contexto)
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
	solicitud := m.ServicioVehicularControlador.ServicioVehicularMapeador.GinContextAEliminarServicioVehicularSolicitud(contexto)
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
	solicitud := m.ServicioVehicularControlador.ServicioVehicularMapeador.GinContextAEditarServicioVehicularSolicitud(contexto)
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
