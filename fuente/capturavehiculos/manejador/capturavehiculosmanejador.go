package manejador

import (
	capturaVehiculosControlador "example/fleetwise/fuente/capturavehiculos/controlador"
	capturaVehiculosMapeador "example/fleetwise/fuente/capturavehiculos/mapeador"
	conectorBDControlador "example/fleetwise/fuente/conectorbd/controlador"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Manejador struct {
	CapturaVehiculosControlador *capturaVehiculosControlador.Controlador
}

func NuevoManejador() (c *Manejador) {
	return &Manejador{
		CapturaVehiculosControlador: &capturaVehiculosControlador.Controlador{
			CapturaVehiculosMapeador: &capturaVehiculosMapeador.Mapeador{},
			ConectorBDControlador: &conectorBDControlador.Controlador{},
		},
	}
}

func (m *Manejador) AgregarVehiculo(contexto *gin.Context) {
	solicitud := m.CapturaVehiculosControlador.CapturaVehiculosMapeador.GinContextAAgregarVehiculoSolicitud(contexto)
	respuesta := m.CapturaVehiculosControlador.AgregarVehiculo(solicitud)
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


