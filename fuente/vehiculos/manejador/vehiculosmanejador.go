package manejador

import (
	vehiculosControlador "example/fleetwise/fuente/vehiculos/controlador"
	vehiculosMapeador "example/fleetwise/fuente/vehiculos/mapeador"
	vehiculosModelos "example/fleetwise/modelos/vehiculos"

	"github.com/gin-gonic/gin"
)

type Manejador struct {
	VehiculosControlador *vehiculosControlador.Controlador
	VehiculosMapeador    *vehiculosMapeador.Mapeador
}

func NuevoManejador() (c *Manejador) {
	return &Manejador{
		VehiculosControlador: &vehiculosControlador.Controlador{},
		VehiculosMapeador:    &vehiculosMapeador.Mapeador{},
	}
}

func (m *Manejador) AgregarVehiculo(context *gin.Context) *vehiculosModelos.AgregarVehiculoRespuesta {
	solicitud := m.VehiculosMapeador.GinContextAAgregarVehiculoSolicitud(context)
	return m.VehiculosControlador.AgregarVehiculo(&solicitud)
}
