package mapeador

import (
	vehiculosModelos "example/fleetwise/modelos/vehiculos"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Mapeador struct {
}

func (m *Mapeador) GinContextAAgregarVehiculoSolicitud(context *gin.Context) (o vehiculosModelos.AgregarVehiculosSolicitud) {
	if anualidad, err := strconv.Atoi(context.Param("id")); err == nil {
		o.AsignarAnualidad(anualidad)
	}
	marca := context.Param("marca")
	o.AsignarMarca(marca)
	modelo := context.Param("modelo")
	o.AsignarModelo(modelo)
	placas := context.Param("placas")
	o.AsignarPlacas(placas)
	serie := context.Param("serie")
	o.AsignarSerie(serie)

	return
}
