package mapeador

import (
	servicioVehicularModelos "example/fleetwise/modelos/capturaserviciovehicular"

	"github.com/gin-gonic/gin"
)

type Mapeador struct {
}

func (m *Mapeador) GinContextAAgregarServicioVehicularSolicitud(contexto *gin.Context) *servicioVehicularModelos.AgregarServicioVehicularSolicitud {
	solicitud := &servicioVehicularModelos.AgregarServicioVehicularSolicitud{}
	contexto.BindJSON(solicitud)
	return solicitud
}

func (m *Mapeador) AgregarServicioVehicularSolicitudAServicioVehicular(solicitud *servicioVehicularModelos.AgregarServicioVehicularSolicitud) *servicioVehicularModelos.ServicioVehicular {
	servicioVehicular := &servicioVehicularModelos.ServicioVehicular{}
	servicioVehicular.AsignarNombre(solicitud.ObtenerNombre())
	return servicioVehicular
}
