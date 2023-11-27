package mapeador

import (
	capturaRegistroModelo "example/fleetwise/modelos/capturaregistromantenimientovehiculo"
	registroMantenimientoVehiculo "example/fleetwise/modelos/capturaregistromantenimientovehiculo"
	conectorBdModelo "example/fleetwise/modelos/conectorbd"
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

func (m *Mapeador) GinContextAObtenerRegistroMantenimientoVehicularPorNumeroDeRegistroSolicitud(contexto *gin.Context) *registroMantenimientoVehiculo.ObtenerRegistroMantenimientoVehicularPorNumeroDeRegistroSolicitud{
	solicitud := &registroMantenimientoVehiculo.ObtenerRegistroMantenimientoVehicularPorNumeroDeRegistroSolicitud{}

	contexto.BindJSON(solicitud)

	return solicitud
}

func(m *Mapeador) ObtenerRegistroMantenimientoVehiculoSolicitudAObtenerRegistroYVehiculoAsociadoPorNumeroDeRegistroSolicitud(solicitud *registroMantenimientoVehiculo.ObtenerRegistroMantenimientoVehicularPorNumeroDeRegistroSolicitud) *registroMantenimientoVehiculo.ObtenerRegistroYVehiculoAsociadoPorNumeroDeRegistroSolicitud{
	solicitudBD := &registroMantenimientoVehiculo.ObtenerRegistroYVehiculoAsociadoPorNumeroDeRegistroSolicitud{NumDeRegistro: solicitud.ObtenerNumDeRegistro()}
	return solicitudBD
}

func (m *Mapeador) GinContextAEditarRegistroMantenimientoVehicularSolicitud(contexto *gin.Context)  *capturaRegistroModelo.EditarRegistroMantenimientoVehicularSolicitud{
	solicitud := &capturaRegistroModelo.EditarRegistroMantenimientoVehicularSolicitud{}
	contexto.BindJSON(solicitud)
	return solicitud
}

func (m *Mapeador) EditarRegistroMantenimientoVehicularSolicitudAActualizarRegistroMantenimientoVehicularSolicitud(solicitud *capturaRegistroModelo.EditarRegistroMantenimientoVehicularSolicitud) (*conectorBdModelo.ActualizarRegistroMantenimientoVehicularSolicitud){
	actualizarRegistroSolicitud := &conectorBdModelo.ActualizarRegistroMantenimientoVehicularSolicitud{
		NumeroDeRegistro : solicitud.ObtenerNumeroDeRegistro(),
		Tipo: solicitud.ObtenerTipo(),
		Fecha: solicitud.ObtenerFecha(),
		LitrosDeGasolina: solicitud.ObtenerLitrosDeGasolina(),
		Kilometraje: solicitud.ObtenerKilometraje(),
		Importe: solicitud.ObtenerImporte(),
		Observaciones: solicitud.ObtenerObservaciones(),
		Concepto: solicitud.ObtenerConcepto(),
		PlacasVehiculo: solicitud.ObtenerPlacasVehiculo(),
	}
	return actualizarRegistroSolicitud
}

func (m *Mapeador) ActualizarRegistroMantenimientoVehicularRespuestaAEditarRegistroMantenimientoVehicularRespuesta(actualizarRegistroRespuesta *conectorBdModelo.ActualizarRegistroMantenimientoVehicularRespuesta) *capturaRegistroModelo.EditarRegistroMantenimientoVehicularRespuesta{
	editarRegistroRespuesta := &capturaRegistroModelo.EditarRegistroMantenimientoVehicularRespuesta{
		Ok: actualizarRegistroRespuesta.ObtenerOk(),
		Err: actualizarRegistroRespuesta.ObtenerError(),
	}
	return editarRegistroRespuesta;
}
