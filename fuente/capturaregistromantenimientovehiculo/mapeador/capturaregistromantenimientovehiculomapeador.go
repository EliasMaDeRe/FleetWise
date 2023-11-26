package mapeador

import (
	registroMantenimientoVehiculo "example/fleetwise/modelos/capturaregistromantenimientovehiculo"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Mapeador struct {
}

func (m *Mapeador) GinContextASeleccionarVehiculoParaNuevoRegistroSolicitud(contexto *gin.Context) *registroMantenimientoVehiculo.SeleccionarVehiculoParaNuevoRegistroSolicitud {
	solicitud := &registroMantenimientoVehiculo.SeleccionarVehiculoParaNuevoRegistroSolicitud{}
	contexto.BindJSON(solicitud)
	return solicitud
}

func (m *Mapeador) GinContextAAgregarRegistroMantemientoVehiculoSolicitud(contexto *gin.Context) *registroMantenimientoVehiculo.AgregarRegistroMantenimientoVehiculoSolicitud {
	solicitud := &registroMantenimientoVehiculo.AgregarRegistroMantenimientoVehiculoSolicitud{}
	contexto.BindJSON(solicitud)
	return solicitud
}

func (m *Mapeador) AgregarRegistroMantemientoVehiculoSolicitudARegistroMantemientoVehiculo(solicitud *registroMantenimientoVehiculo.AgregarRegistroMantenimientoVehiculoSolicitud) *registroMantenimientoVehiculo.RegistroMantenimientoVehiculo {
	registro := &registroMantenimientoVehiculo.RegistroMantenimientoVehiculo{}

	tipo := solicitud.ObtenerTipo()
	registro.AsignarTipo(tipo)

	if fecha, err := strconv.Atoi(solicitud.ObtenerFecha()); err == nil {
		registro.AsignarFecha(strconv.Itoa(fecha))
	} else {
		registro.AsignarFecha("0")
	}

	kilometraje := solicitud.ObtenerKilometraje()
	registro.AsignarKilometraje(kilometraje)

	litrosGasolina := solicitud.ObtenerLitrosDeGasolina()
	registro.AsignarLitrosDeGasolina(litrosGasolina)

	importe := solicitud.ObtenerImporte()
	registro.AsignarImporte(importe)

	concepto := solicitud.ObtenerConcepto()
	registro.AsignarConcepto(concepto)

	return registro
}

func (m *Mapeador) GinContextAObtenerRegistroMantenimientoVehiculoSolicitud(contexto *gin.Context) *registroMantenimientoVehiculo.ObtenerRegistroMantenimientoVehicularSolicitud{
	solicitud := &registroMantenimientoVehiculo.ObtenerRegistroMantenimientoVehicularSolicitud{}

	contexto.BindJSON(solicitud)

	return solicitud
}

func(m *Mapeador) ObtenerRegistroMantenimientoVehiculoSolicitudAObtenerRegistroYVehiculoAsociadoPorNumeroDeRegistroSolicitud(solicitud *registroMantenimientoVehiculo.ObtenerRegistroMantenimientoVehicularSolicitud) *registroMantenimientoVehiculo.ObtenerRegistroYVehiculoAsociadoPorNumeroDeRegistro{
	solicitudBD := &registroMantenimientoVehiculo.ObtenerRegistroYVehiculoAsociadoPorNumeroDeRegistro{NumDeRegistro: solicitud.ObtenerNumDeRegistro()}
	return solicitudBD
}