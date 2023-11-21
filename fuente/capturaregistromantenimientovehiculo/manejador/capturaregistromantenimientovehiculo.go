package manejador

import (
	registroMantenimientoVehiculoControlador "example/fleetwise/fuente/capturaregistromantenimientovehiculo/controlador"
	registroMantenimientoVehiculoMapeador "example/fleetwise/fuente/capturaregistromantenimientovehiculo/mapeador"

	//"net/http"

	"github.com/gin-gonic/gin"
)

type Manejador struct {
	RegistroMantenimientoVehiculoControlador *registroMantenimientoVehiculoControlador.Controlador
}

func NuevoManejador() (c *Manejador) {
	return &Manejador{
		RegistroMantenimientoVehiculoControlador: &registroMantenimientoVehiculoControlador.Controlador{
			RegistroMantenimientoVehiculoMapeador: &registroMantenimientoVehiculoMapeador.Mapeador{},
		},
	}
}

func (m *Manejador) SeleccionarVehiculo(contexto *gin.Context) {
	//implementacipon futura por Pablo
}

func (m *Manejador) AgregarRegistroMantemientoVehiculo(contexto *gin.Context) {
	//implementacipon futura por Pablo
}
