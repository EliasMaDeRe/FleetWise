package Manejador

import (
	conectorBDControlador "example/fleetwise/fuente/conectorbd/controlador"

	"github.com/gin-gonic/gin"
)

type Manejador struct {
	ConectorControlador *conectorBDControlador.Controlador
}

func NuevoManejador() (c *Manejador) {
	return &Manejador{
		ConectorControlador: &conectorBDControlador.Controlador{},
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

func (m *Manejador) GuardarServicioVehicular(contexto *gin.Context) {
	//Se dejan vacíos para su futura implementación
}

func (m *Manejador) ObtenerServicioVehicularPorNombre(contexto *gin.Context) {
	//Se dejan vacíos para su futura implementación
}

func (m *Manejador) ObtenerServiciosVehiculares(contexto *gin.Context) {
	//Se dejan vacíos para su futura implementación
}
