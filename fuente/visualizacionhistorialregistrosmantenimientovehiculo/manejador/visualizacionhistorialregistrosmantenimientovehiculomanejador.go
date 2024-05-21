package manejador

import (
	conectorBDControlador "example/fleetwise/fuente/conectorbdvisualizacionhistorialregistrosmantenimientovehiculo/controlador"
	visualizacionHistorialRegistrosControlador "example/fleetwise/fuente/visualizacionhistorialregistrosmantenimientovehiculo/controlador"
	visualizacionHistorialRegistrosMapeador "example/fleetwise/fuente/visualizacionhistorialregistrosmantenimientovehiculo/mapeador"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Manejador struct{
	HistorialRegistrosControlador *visualizacionHistorialRegistrosControlador.Controlador

}

func NuevoManejador() (m *Manejador){
	return &Manejador{
		HistorialRegistrosControlador: &visualizacionHistorialRegistrosControlador.Controlador{
			HistorialRegistrosMapeador: &visualizacionHistorialRegistrosMapeador.Mapeador{},
			ConectorBDControlador: &conectorBDControlador.Controlador{},
		},
	}
}

func (m *Manejador) ObtenerRegistrosYVehiculosFiltrados(contexto *gin.Context){
	solicitud := m.HistorialRegistrosControlador.HistorialRegistrosMapeador.GinContextAObtenerRegistrosYVehiculosFiltradosSolicitud(contexto)
	registrosFiltrados, vehiculosFiltrados:= m.HistorialRegistrosControlador.ObtenerRegistrosFiltradosConVehiculos(solicitud);

	contexto.IndentedJSON(http.StatusOK, gin.H{"registrosFiltrados":registrosFiltrados, "vehiculosFiltrados":vehiculosFiltrados})

}