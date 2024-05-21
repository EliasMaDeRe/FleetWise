package mapeador

import (
	capturaServicioVehicularModelos "example/fleetwise/modelos/capturaserviciovehicular"
	conectorBDModelos "example/fleetwise/modelos/conectorbdcapturaserviciovehicular"

	"github.com/gin-gonic/gin"
)

type Mapeador struct {
}

func (m *Mapeador) GinContextAAgregarServicioVehicularSolicitud(contexto *gin.Context) *capturaServicioVehicularModelos.AgregarServicioVehicularSolicitud {
	solicitud := &capturaServicioVehicularModelos.AgregarServicioVehicularSolicitud{}
	contexto.BindJSON(solicitud)
	return solicitud
}

func (m *Mapeador) GinContextAEliminarServicioVehicularSolicitud(contexto *gin.Context) *capturaServicioVehicularModelos.EliminarServicioVehicularSolicitud {
	solicitud := &capturaServicioVehicularModelos.EliminarServicioVehicularSolicitud{}
	contexto.BindJSON(solicitud)
	return solicitud
}

func (m *Mapeador) GinContextAEditarServicioVehicularSolicitud(contexto *gin.Context) *capturaServicioVehicularModelos.EditarServicioVehicularSolicitud {
	solicitud := &capturaServicioVehicularModelos.EditarServicioVehicularSolicitud{}
	contexto.BindJSON(solicitud)
	return solicitud
}

func (m *Mapeador) AgregarServicioVehicularSolicitudAServicioVehicular(solicitud *capturaServicioVehicularModelos.AgregarServicioVehicularSolicitud) *capturaServicioVehicularModelos.ServicioVehicular {
	servicioVehicular := &capturaServicioVehicularModelos.ServicioVehicular{}
	servicioVehicular.AsignarNombre(solicitud.ObtenerNombre())
	return servicioVehicular
}

func (m *Mapeador) EliminarServicioVehicularSolicitudAServicioVehicular(solicitud *capturaServicioVehicularModelos.EliminarServicioVehicularSolicitud) *capturaServicioVehicularModelos.ServicioVehicular {
	servicioVehicular := &capturaServicioVehicularModelos.ServicioVehicular{}
	servicioVehicular.AsignarNombre(solicitud.ObtenerNombre())
	return servicioVehicular
}

func (m *Mapeador) EditarServicioVehicularSolicitudAServicioVehicularDeConectorBD(solicitud *capturaServicioVehicularModelos.EditarServicioVehicularSolicitud) *conectorBDModelos.EditarServicioVehicularSolicitud {
	editarServicioVehicularSolicitud := conectorBDModelos.EditarServicioVehicularSolicitud{}
	editarServicioVehicularSolicitud.AsignarNombreActual(solicitud.ObtenerNombreActual())
	editarServicioVehicularSolicitud.AsignarNuevoNombre(solicitud.ObtenerNuevoNombre())
	return &editarServicioVehicularSolicitud
}
