package mapeador

import (
	registroMantenimientoVehiculo "example/fleetwise/modelos/capturaregistromantenimientovehiculo"
	//"strconv"

	"github.com/gin-gonic/gin"
)

type Mapeador struct {
}

func (m *Mapeador) GinContextASeleccionarVehiculoSolicitud(contexto *gin.Context) *registroMantenimientoVehiculo.SeleccionarVehiculoSolicitud {
	//implementacipon futura por Pablo
	return nil
}

func (m *Mapeador) GinContextAAgregarRegistroMantemientoVehiculoSolicitud(contexto *gin.Context) *registroMantenimientoVehiculo.AgregarRegistroMantenimientoVehiculoSolicitud{
	//implementacipon futura por Pablo
	return nil
}

func (m *Mapeador) AgregaRegistroMantemientoVehiculoSolicitudARegistroMantemientoVehiculo(solicitud *registroMantenimientoVehiculo.AgregarRegistroMantenimientoVehiculoSolicitud) *registroMantenimientoVehiculo.RegistroMantenimientoVehiculo {
	//implementacipon futura por Pablo
	return nil
}

