package manejador

import (
	"example/fleetwise/fuente/conectorbd/controlador"
	vhc "example/fleetwise/fuente/visualizacionhistorialregistrosmantenimientovehiculo/controlador"
	visualizacionResumenMantenimientoVehiculosControlador "example/fleetwise/fuente/visualizacionresumenmantenimientovehiculos/controlador"
	vrm "example/fleetwise/fuente/visualizacionresumenmantenimientovehiculos/mapeador"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Manejador struct {
	visualizacionResumenMantenimientoVehiculosControlador *visualizacionResumenMantenimientoVehiculosControlador.Controlador
}

func NuevoManejador() (c *Manejador) {
	return &Manejador{
		visualizacionResumenMantenimientoVehiculosControlador: &visualizacionResumenMantenimientoVehiculosControlador.Controlador{
			ConectorBDControlador:                      &controlador.Controlador{},
			VisualizacionHistorialRegistrosControlador: &vhc.Controlador{},
			ResumenFlotillaMapeador:                    &vrm.Mapeador{},
		},
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

func (m *Manejador) ExportarResumenFlotilla(contexto *gin.Context) {

	solicitud := m.visualizacionResumenMantenimientoVehiculosControlador.ResumenFlotillaMapeador.GinContextAExportarResumenFlotillaSolicitud(contexto)

	archivo, _ := m.visualizacionResumenMantenimientoVehiculosControlador.ExportarResumenFlotilla(string(solicitud.ObtenerFormato()))
	contexto.Header("Content-Transfer-Encoding", "binary")
	contexto.Header("Content-Disposition", "attachment; filename=ReporteFlotilla.xlsx")
	contexto.Header("Content-Type", "application/octet-stream")
	archivo.Write(contexto.Writer)
}
