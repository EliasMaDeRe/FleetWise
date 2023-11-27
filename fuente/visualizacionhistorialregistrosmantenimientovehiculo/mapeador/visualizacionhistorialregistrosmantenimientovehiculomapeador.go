package mapeador

import (
	modelosBd "example/fleetwise/modelos/conectorbd"
	historialRegistrosModelo "example/fleetwise/modelos/visualizacionhistorialregistrosmantenimientovehiculo"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Mapeador struct{
}

func (m *Mapeador) GinContextAObtenerRegistrosYVehiculosFiltradosSolicitud(contexto *gin.Context) *historialRegistrosModelo.ObtenerRegistrosFiltradosConVehiculosSolicitud{
	solicitud := &historialRegistrosModelo.ObtenerRegistrosFiltradosConVehiculosSolicitud{}
	contexto.BindJSON(solicitud)
	return solicitud
}

func (m *Mapeador) ObtenerRegistrosFiltradosConVehiculosSolicitudAObtenerRegistrosYVehiculosAsociadosFiltradosSolicitud (solicitud *historialRegistrosModelo.ObtenerRegistrosFiltradosConVehiculosSolicitud) *modelosBd.ObtenerRegistrosYVehiculosAsociadosFiltradosSolicitud{
	
	obtenerRegistrosConVehiculosFiltradosSolicitud:= &modelosBd.ObtenerRegistrosYVehiculosAsociadosFiltradosSolicitud{}

	obtenerRegistrosConVehiculosFiltradosSolicitud.AsignarFiltroPlaca(solicitud.ObtenerPlacas())
	obtenerRegistrosConVehiculosFiltradosSolicitud.AsignarFiltroMarca(solicitud.ObtenerMarca())
	obtenerRegistrosConVehiculosFiltradosSolicitud.AsignarFiltroModelo(solicitud.ObtenerModelo())

	if fechaLanzamiento, err := strconv.Atoi(solicitud.ObtenerFechaDeLanzamiento()); err == nil {
		obtenerRegistrosConVehiculosFiltradosSolicitud.AsignarFiltroFechaDeLanzamiento(fechaLanzamiento)
	} else {
		obtenerRegistrosConVehiculosFiltradosSolicitud.AsignarFiltroFechaDeLanzamiento(0)
	}	

	return obtenerRegistrosConVehiculosFiltradosSolicitud
}