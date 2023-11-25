package mapeador

import (
	modelosBd "example/fleetwise/modelos/conectorbd"
	historialRegistrosModelo "example/fleetwise/modelos/visualizacionhistorialregistrosmantenimientovehicular"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Mapeador struct{
}

func (m *Mapeador) GinContextAObtenerRegistrosYVehiculosFiltradosSolicitud(contexto *gin.Context) *historialRegistrosModelo.ObtenerRegistrosYVehiculosFiltradosSolicitud{
	solicitud := &historialRegistrosModelo.ObtenerRegistrosYVehiculosFiltradosSolicitud{}
	contexto.BindJSON(solicitud)
	return solicitud
}

func (m *Mapeador) ObtenerRegistrosYVehiculosFiltradosSolicitudAObtenerRegistrosConVehiculosFiltradosSolicitud (solicitud *historialRegistrosModelo.ObtenerRegistrosYVehiculosFiltradosSolicitud) *modelosBd.ObtenerRegistrosConVehiculosFiltradosSolicitud{
	
	solicitudBd:= &modelosBd.ObtenerRegistrosConVehiculosFiltradosSolicitud{}

	solicitudBd.AsignarFiltroPlaca(solicitud.ObtenerPlacas())
	solicitudBd.AsignarFiltroMarca(solicitud.ObtenerMarca())
	solicitudBd.AsignarFiltroModelo(solicitud.ObtenerModelo())

	if fechaLanzamiento, err := strconv.Atoi(solicitud.ObtenerFechaDeLanzamiento()); err == nil {
		solicitudBd.AsignarFiltroFechaDeLanzamiento(fechaLanzamiento)
	} else {
		solicitudBd.AsignarFiltroFechaDeLanzamiento(0)
	}	

	return solicitudBd
}