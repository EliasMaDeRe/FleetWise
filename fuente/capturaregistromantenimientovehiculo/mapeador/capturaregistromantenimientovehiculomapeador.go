package mapeador

import (
	registroMantenimientoVehiculo "example/fleetwise/modelos/capturaregistromantenimientovehiculo"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Mapeador struct {
}

func (m *Mapeador) GinContextASeleccionarVehiculoDeNuevoRegistroSolicitud(contexto *gin.Context) *registroMantenimientoVehiculo.SeleccionarVehiculoParaNuevoRegistroSolicitud {
	solicitud := &registroMantenimientoVehiculo.SeleccionarVehiculoParaNuevoRegistroSolicitud{}
	contexto.BindJSON(solicitud)
	return solicitud
}

func (m *Mapeador) GinContextAAgregarRegistroMantemientoVehiculoSolicitud(contexto *gin.Context) *registroMantenimientoVehiculo.AgregarRegistroMantenimientoVehiculoSolicitud {
	solicitud := &registroMantenimientoVehiculo.AgregarRegistroMantenimientoVehiculoSolicitud{}
	contexto.BindJSON(solicitud)
	return solicitud
}

func (m *Mapeador) AgregaRegistroMantemientoVehiculoSolicitudARegistroMantemientoVehiculo(solicitud *registroMantenimientoVehiculo.AgregarRegistroMantenimientoVehiculoSolicitud) *registroMantenimientoVehiculo.RegistroMantenimientoVehiculo {
	registro := &registroMantenimientoVehiculo.RegistroMantenimientoVehiculo{}
	if fecha, err := strconv.Atoi(solicitud.ObtenerFecha()); err == nil {
		registro.AsignarFecha(fecha)
	} else {
		registro.AsignarFecha(0)
	}
	//implementacion pablo
	//vehiculo.AsignarMarca(solicitud.ObtenerMarca())

	return registro
}
