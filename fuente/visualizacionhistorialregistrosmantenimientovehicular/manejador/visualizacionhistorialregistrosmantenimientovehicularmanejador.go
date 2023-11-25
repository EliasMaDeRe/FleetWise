package manejador

import (
	historialRegistrosControlador "example/fleetwise/fuente/visualizacionhistorialregistrosmantenimientovehicular/controlador"
	historialRegistrosMapeador "example/fleetwise/fuente/visualizacionhistorialregistrosmantenimientovehicular/mapeador"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Manejador struct{
	HistorialRegistrosControlador *historialRegistrosControlador.Controlador

}

func NuevoManejador() (m *Manejador){
	return &Manejador{
		HistorialRegistrosControlador: &historialRegistrosControlador.Controlador{
			HistorialRegistrosMapeador: &historialRegistrosMapeador.Mapeador{},
		},
	}
}

func (m *Manejador) ObtenerRegistrosYVehiculosFiltrados(contexto *gin.Context){
	solicitud:= m.HistorialRegistrosControlador.HistorialRegistrosMapeador.GinContextAObtenerRegistrosYVehiculosFiltradosSolicitud(contexto)
	registrosFiltrados, vehiculosFiltrados:= m.HistorialRegistrosControlador.ObtenerRegistrosYVehiculosFiltrados(solicitud);

	contexto.IndentedJSON(http.StatusOK, gin.H{"registrosFiltrados":registrosFiltrados, "vehiculosFiltrados":vehiculosFiltrados})

}