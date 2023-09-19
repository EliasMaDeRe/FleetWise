package manejador

import (
	vehiculosControlador "example/fleetwise/fuente/vehiculos/controlador"
	vehiculosMapeador "example/fleetwise/fuente/vehiculos/mapeador"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Manejador struct {
	VehiculosControlador *vehiculosControlador.Controlador
}

func NuevoManejador() (c *Manejador) {
	return &Manejador{
		VehiculosControlador: &vehiculosControlador.Controlador{
			VehiculosMapeador: &vehiculosMapeador.Mapeador{},
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
	contexto.IndentedJSON(status, gin.H{"OK": respuesta.ObtenerOk()})
}
