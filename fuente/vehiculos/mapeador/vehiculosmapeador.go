package mapeador

import (
	vehiculosModelos "example/fleetwise/modelos/vehiculos"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Mapeador struct {
}

func (m *Mapeador) GinContextAAgregarVehiculoSolicitud(contexto *gin.Context) (o vehiculosModelos.AgregarVehiculosSolicitud) {
	if anualidad, err := strconv.Atoi(contexto.Param("id")); err == nil {
		o.AsignarAnualidad(anualidad)
	}
	marca := contexto.Param("marca")
	o.AsignarMarca(marca)
	modelo := contexto.Param("modelo")
	o.AsignarModelo(modelo)
	placas := contexto.Param("placas")
	o.AsignarPlacas(placas)
	serie := contexto.Param("serie")
	o.AsignarSerie(serie)

	return
}
