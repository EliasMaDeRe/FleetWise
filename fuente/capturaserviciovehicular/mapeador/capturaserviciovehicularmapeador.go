package mapeador

import (
	servicioVehicularModelos "example/fleetwise/modelos/capturaserviciovehicular"
	conectorBDModelos "example/fleetwise/modelos/conectorbd"

	"github.com/gin-gonic/gin"
)

type Mapeador struct {
}

func (m *Mapeador) GinContextAAgregarServicioVehicularSolicitud(contexto *gin.Context) *servicioVehicularModelos.AgregarServicioVehicularSolicitud {
	solicitud := &servicioVehicularModelos.AgregarServicioVehicularSolicitud{}
	contexto.BindJSON(solicitud)
	return solicitud
}

func (m *Mapeador) GinContextAEliminarServicioVehicularSolicitud(contexto *gin.Context) *servicioVehicularModelos.EliminarServicioVehicularSolicitud {
	solicitud := &servicioVehicularModelos.EliminarServicioVehicularSolicitud{}
	contexto.BindJSON(solicitud)
	return solicitud
}

func (m *Mapeador) GinContextAEditarServicioVehicularSolicitud(contexto *gin.Context) *servicioVehicularModelos.EditarServicioVehicularSolicitud {
	solicitud := &servicioVehicularModelos.EditarServicioVehicularSolicitud{}
	contexto.BindJSON(solicitud)
	return solicitud
}

func (m *Mapeador) AgregarServicioVehicularSolicitudAServicioVehicular(solicitud *servicioVehicularModelos.AgregarServicioVehicularSolicitud) *servicioVehicularModelos.ServicioVehicular {
	servicioVehicular := &servicioVehicularModelos.ServicioVehicular{}
	servicioVehicular.AsignarNombre(solicitud.ObtenerNombre())
	return servicioVehicular
}

func (m *Mapeador) EliminarServicioVehicularSolicitudAServicioVehicular(solicitud *servicioVehicularModelos.EliminarServicioVehicularSolicitud) *servicioVehicularModelos.ServicioVehicular {
	servicioVehicular := &servicioVehicularModelos.ServicioVehicular{}
	servicioVehicular.AsignarNombre(solicitud.ObtenerNombre())
	return servicioVehicular
}

func (m *Mapeador) EditarServicioVehicularSolicitudAServicioVehicularDeConectorBD(solicitud *servicioVehicularModelos.EditarServicioVehicularSolicitud) *conectorBDModelos.EditarServicioVehicularSolicitud {
	editarServicioVehicularSolicitud := conectorBDModelos.EditarServicioVehicularSolicitud{}
	editarServicioVehicularSolicitud.AsignarNombreActual(solicitud.ObtenerNombreActual())
	editarServicioVehicularSolicitud.AsignarNuevoNombre(solicitud.ObtenerNuevoNombre())
	return &editarServicioVehicularSolicitud
}
