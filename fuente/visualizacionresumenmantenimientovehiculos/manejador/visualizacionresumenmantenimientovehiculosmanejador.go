package manejador

import (
	visualizacionResumenMantenimientoVehiculosControlador "example/fleetwise/fuente/visualizacionresumenmantenimientovehiculos/controlador"
	"net/http"

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

func (m *Manejador) ObtenerMetricasVehiculares(contexto *gin.Context) {
	vehiculos, gastosDeCombustible, gastosEnMantenimiento, rendimientosKilometroLitro, kilometrajesIniciales, ultimosKilometrajes, kilometrosTotalesRecorridos, kilometrosPromedioDiariosRecorridos := m.visualizacionResumenMantenimientoVehiculosControlador.ObtenerMetricasVehiculares()
	contexto.IndentedJSON(http.StatusOK, gin.H{
		"vehiculos":                                      vehiculos,
		"gastosDeCombustiblePorVehiculo":                 gastosDeCombustible,
		"gastosEnMantenimientoPorVehiculo":               gastosEnMantenimiento,
		"rendimientosKilometroLitroPorVehiculo":          rendimientosKilometroLitro,
		"kilometrajesInicialesPorVehiculo":               kilometrajesIniciales,
		"ultimosKilometrajesPorVehiculo":                 ultimosKilometrajes,
		"kilometrosTotalesRecorridosPorVehiculo":         kilometrosTotalesRecorridos,
		"kilometrosPromedioDiariosRecorridosPorVehiculo": kilometrosPromedioDiariosRecorridos})
}
