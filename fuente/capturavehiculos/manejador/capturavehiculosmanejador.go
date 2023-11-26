package manejador

import (
	vehiculosControlador "example/fleetwise/fuente/capturavehiculos/controlador"
	vehiculosMapeador "example/fleetwise/fuente/capturavehiculos/mapeador"
	conectorBDControlador "example/fleetwise/fuente/conectorbd/controlador"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Manejador struct {
	VehiculosControlador *vehiculosControlador.Controlador
}

func NuevoManejador() (c *Manejador) {
	return &Manejador{
		VehiculosControlador: &vehiculosControlador.Controlador{
			VehiculosMapeador:     &vehiculosMapeador.Mapeador{},
			ConectorBDControlador: &conectorBDControlador.Controlador{},
		},
	}
}

func (m *Manejador) AgregarVehiculo(contexto *gin.Context) {
	solicitud := m.VehiculosControlador.VehiculosMapeador.GinContextAAgregarVehiculoSolicitud(contexto)
	respuesta := m.VehiculosControlador.AgregarVehiculo(solicitud)
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
