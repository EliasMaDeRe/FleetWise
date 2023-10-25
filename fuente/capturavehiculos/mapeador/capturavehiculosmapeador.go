package mapeador

import (
	vehiculosModelos "example/fleetwise/modelos/capturavehiculos"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Mapeador struct {
}

func (m *Mapeador) GinContextAAgregarVehiculoSolicitud(contexto *gin.Context) *vehiculosModelos.AgregarVehiculosSolicitud {
	solicitud := &vehiculosModelos.AgregarVehiculosSolicitud{}
	contexto.BindJSON(solicitud)
	return solicitud
}

func (m *Mapeador) AgregarVehiculosSolicitudAVehiculo(solicitud *vehiculosModelos.AgregarVehiculosSolicitud) *vehiculosModelos.Vehiculo {
	vehiculo := &vehiculosModelos.Vehiculo{}
	if fechaLanzamiento, err := strconv.Atoi(solicitud.ObtenerFechaLanzamiento()); err == nil {
		vehiculo.AsignarFechaLanzamiento(fechaLanzamiento)
	} else {
		vehiculo.AsignarFechaLanzamiento(0)
	}
	vehiculo.AsignarMarca(solicitud.ObtenerMarca())
	vehiculo.AsignarModelo(solicitud.ObtenerModelo())
	vehiculo.AsignarPlacas(solicitud.ObtenerPlacas())
	vehiculo.AsignarSerie(solicitud.ObtenerSerie())
	return vehiculo
}
