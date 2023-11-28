package mapeador

import (
	CapturaVehiculosModelos "example/fleetwise/modelos/capturavehiculos"
	ConectorBDModelos "example/fleetwise/modelos/conectorbd"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Mapeador struct {
}

func (m *Mapeador) GinContextAAgregarVehiculoSolicitud(contexto *gin.Context) *CapturaVehiculosModelos.AgregarVehiculoSolicitud {
	solicitud := &CapturaVehiculosModelos.AgregarVehiculoSolicitud{}
	contexto.BindJSON(solicitud)
	return solicitud
}

func (m *Mapeador) AgregarVehiculosSolicitudAVehiculo(solicitud *CapturaVehiculosModelos.AgregarVehiculoSolicitud) *CapturaVehiculosModelos.Vehiculo {
	vehiculo := &CapturaVehiculosModelos.Vehiculo{}
	if fechaLanzamiento, err := strconv.Atoi(solicitud.ObtenerFechaLanzamiento()); err == nil {
		vehiculo.AsignarFechaLanzamiento(fechaLanzamiento)
	} else {
		vehiculo.AsignarFechaLanzamiento(0)
	}
	vehiculo.AsignarMarca(solicitud.ObtenerMarca())
	vehiculo.AsignarModelo(solicitud.ObtenerModelo())
	vehiculo.AsignarPlacas(solicitud.ObtenerPlacas())
	vehiculo.AsignarSerie(solicitud.ObtenerSerie())
	return vehiculo
}

func (m *Mapeador) GinContextAEditarVehiculoSolicitud(contexto *gin.Context) *CapturaVehiculosModelos.EditarVehiculoSolicitud {
	solicitud := &CapturaVehiculosModelos.EditarVehiculoSolicitud{}
	contexto.BindJSON(solicitud)
	return solicitud
}

func (m *Mapeador) EditarVehiculosSolicitudAVehiculo(solicitud *CapturaVehiculosModelos.EditarVehiculoSolicitud) *ConectorBDModelos.EditarVehiculoSolicitud {
	solicitudConectorBD := &ConectorBDModelos.EditarVehiculoSolicitud{}
	solicitudConectorBD.AsignarPlacasActuales(solicitud.ObtenerPlacasActuales())
	solicitudConectorBD.AsignarPlacasNuevas(solicitud.ObtenerPlacasNuevas())
	if fechaLanzamiento, err := strconv.Atoi(solicitud.ObtenerFechaDeLanzamientoNueva()); err == nil {
		solicitudConectorBD.AsignarFechaDeLanzamientoNueva(fechaLanzamiento)
	} else {
		solicitudConectorBD.AsignarFechaDeLanzamientoNueva(0)
	}
	solicitudConectorBD.AsignarMarcaNueva(solicitud.ObtenerMarcaNueva())
	solicitudConectorBD.AsignarModeloNuevo(solicitud.ObtenerModeloNuevo())
	solicitudConectorBD.AsignarSerieNuevo(solicitud.ObtenerSerieNuevo())
	return solicitudConectorBD
}

func (m *Mapeador) GinContextAObtenerVehiculoPorPlacasSolicitud(contexto *gin.Context) *CapturaVehiculosModelos.ObtenerVehiculoPorPlacasSolicitud {
	solicitud := &CapturaVehiculosModelos.ObtenerVehiculoPorPlacasSolicitud{}
	contexto.BindJSON(solicitud)
	return solicitud
}
