package Manejador

import (
	conectorControlador "example/fleetwise/fuente/conectorbd/controlador"

	"github.com/gin-gonic/gin"
)

type Manejador struct {
	ConectorControlador *conectorControlador.Controlador
}

func NuevoManejador() (c *Manejador) {
	return &Manejador{
		ConectorControlador: &conectorControlador.Controlador{},
	}
}

func (m *Manejador) ObtenerVehiculoPorPlacas(contexto *gin.Context) {
	//Se dejan vacíos para su futura implementación
}

func (m *Manejador) ObtenerVehiculoPorSerie(contexto *gin.Context) {
	//Se dejan vacíos para su futura implementación
}

func (m *Manejador) GuardarVehiculo(contexto *gin.Context) {
	//Se dejan vacíos para su futura implementación
}
