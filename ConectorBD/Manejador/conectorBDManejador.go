package Manejador

import (
	"github.com/gin-gonic/gin"

	conectorControlador "example/fleetwise/ConectorBD/Controlador"
	conectorMapeador "example/fleetwise/ConectorBD/Mapeador"
	conectorModelos "example/fleetwise/modelos/conectorbd"
)

type Manejador struct {
	ConectorControlador *conectorControlador.Controlador
	ConectorMapeador    *conectorMapeador.Mapeador
}

func NuevoManejador() (c *Manejador) {
	return &Manejador{
		ConectorControlador: &conectorControlador.Controlador{},
	}
}

func (m *Manejador) GuardarVehiculo(contexto *gin.Context) {
	//no se me ocurrió que poner pipipi
}

func (m *Manejador) ObtenerVehiculo(contexto *gin.Context) *conectorModelos.ObtenerCargoSoicitud {

	respuesta := &conectorModelos.ObtenerCargoSoicitud{}
	return respuesta
}
