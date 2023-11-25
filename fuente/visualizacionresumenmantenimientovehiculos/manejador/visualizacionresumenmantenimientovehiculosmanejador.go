package manejador

import (
	visualizacionResumenMantenimientoVehiculosControlador "example/fleetwise/fuente/visualizacionresumenmantenimientovehiculos/controlador"

	"github.com/gin-gonic/gin"
)

type Manejador struct {
	visualizacionResumenMantenimientoVehiculosControlador *visualizacionResumenMantenimientoVehiculosControlador.Controlador
}

func NuevoManejador() (c *Manejador) {
	return &Manejador{
		visualizacionResumenMantenimientoVehiculosControlador: &visualizacionResumenMantenimientoVehiculosControlador.Controlador{},
	}
}

func (m *Manejador) CalcularMetricasVehiculares(contexto *gin.Context) {

}

func (m *Manejador) ObtenerResumenMantenimientoVehiculos(contexto *gin.Context) {
	
}