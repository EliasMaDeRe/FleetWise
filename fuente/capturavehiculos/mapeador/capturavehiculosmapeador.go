package mapeador

import (
	vehiculosModelos "example/fleetwise/modelos/capturavehiculos"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Mapeador struct {
}

func (m *Mapeador) GinContextAAgregarVehiculoSolicitud(contexto *gin.Context) *vehiculosModelos.AgregarVehiculosSolicitud {
	solicitud := &vehiculosModelos.AgregarVehiculosSolicitud{}
	contexto.BindJSON(solicitud)
	fmt.Println(solicitud.ObtenerAnualidad())
	return solicitud
}

func (m *Mapeador) AgregarVehiculosSolicitudAVehiculo(solicitud *vehiculosModelos.AgregarVehiculosSolicitud) *vehiculosModelos.Vehiculo {
	vehiculo := &vehiculosModelos.Vehiculo{}
	if anualidad, err := strconv.Atoi(solicitud.ObtenerAnualidad()); err != nil {
		vehiculo.AsignarAnualidad(anualidad)
	} else {
		vehiculo.AsignarAnualidad(0)
	}
	vehiculo.AsignarMarca(solicitud.ObtenerMarca())
	vehiculo.AsignarModelo(solicitud.ObtenerModelo())
	vehiculo.AsignarPlacas(solicitud.ObtenerPlacas())
	vehiculo.AsignarSerie(solicitud.ObtenerSerie())
	return vehiculo
}
