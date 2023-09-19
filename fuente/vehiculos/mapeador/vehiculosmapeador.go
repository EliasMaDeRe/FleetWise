package mapeador

import (
	vehiculosModelos "example/fleetwise/modelos/vehiculos"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Mapeador struct {
}

func (m *Mapeador) GinContextAAgregarVehiculoSolicitud(contexto *gin.Context) *vehiculosModelos.AgregarVehiculosSolicitud {
	solicitud := &vehiculosModelos.AgregarVehiculosSolicitud{}
	if anualidad, err := strconv.Atoi(contexto.Param("id")); err == nil {
		solicitud.AsignarAnualidad(anualidad)
	}
	marca := contexto.Param("marca")
	solicitud.AsignarMarca(marca)
	modelo := contexto.Param("modelo")
	solicitud.AsignarModelo(modelo)
	placas := contexto.Param("placas")
	solicitud.AsignarPlacas(placas)
	serie := contexto.Param("serie")
	solicitud.AsignarSerie(serie)

	return solicitud
}

func (m *Mapeador) AgregarVehiculosSolicitudAVehiculo(solicitud *vehiculosModelos.AgregarVehiculosSolicitud) *vehiculosModelos.Vehiculo {
	vehiculo := &vehiculosModelos.Vehiculo{}
	vehiculo.AsignarAnualidad(solicitud.ObtenerAnualidad())
	vehiculo.AsignarMarca(solicitud.ObtenerMarca())
	vehiculo.AsignarModelo(solicitud.ObtenerModelo())
	vehiculo.AsignarPlacas(solicitud.ObtenerPlacas())
	vehiculo.AsignarSerie(solicitud.ObtenerSerie())
	return vehiculo
}
