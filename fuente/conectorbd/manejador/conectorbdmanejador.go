package Manejador

import (
	conectorControlador "example/fleetwise/fuente/conectorbd/controlador"
	conectorMapeador "example/fleetwise/fuente/conectorbd/mapeador"

	"github.com/gin-gonic/gin"
)

type Manejador struct {
	ConectorControlador *conectorControlador.Controlador
	ConectorMapeador    *conectorMapeador.Mapeador
}

func NuevoManejador() (c *Manejador) {
	return &Manejador{
		ConectorControlador: &conectorControlador.Controlador{},
		ConectorMapeador:    &conectorMapeador.Mapeador{},
	}
}

func (m *Manejador) ObtenerVehiculoPorPlacas(contexto *gin.Context) {
	//
}

func (m *Manejador) ObtenerVehiculoPorSerie(contexto *gin.Context) {
	//
}

func (m *Manejador) AgregarVehiculo(contexto *gin.Context) {
	//
}
