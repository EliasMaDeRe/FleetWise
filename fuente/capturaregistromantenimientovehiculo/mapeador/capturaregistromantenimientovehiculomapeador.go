package mapeador

import (
	capturaRegistroMantenimientoVehiculoModelos "example/fleetwise/modelos/capturaregistromantenimientovehiculo"
	conectorBDModelos "example/fleetwise/modelos/conectorbd"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type Mapeador struct {
}

func (m *Mapeador) GinContextASeleccionarVehiculoParaNuevoRegistroSolicitud(contexto *gin.Context) *capturaRegistroMantenimientoVehiculoModelos.SeleccionarVehiculoParaNuevoRegistroSolicitud {
	solicitud := &capturaRegistroMantenimientoVehiculoModelos.SeleccionarVehiculoParaNuevoRegistroSolicitud{}
	if strings.Contains(contexto.Request.URL.String(), "?placas=") {
		divisionesURL := strings.SplitAfter(contexto.Request.URL.String(), "?placas=")
		solicitud.AsignarPlacas(divisionesURL[1])
		return solicitud
	}
	contexto.BindJSON(solicitud)
	return solicitud
}

func (m *Mapeador) GinContextAAgregarRegistroMantemientoVehiculoSolicitud(contexto *gin.Context) *capturaRegistroMantenimientoVehiculoModelos.AgregarRegistroMantenimientoVehiculoSolicitud {
	solicitud := &capturaRegistroMantenimientoVehiculoModelos.AgregarRegistroMantenimientoVehiculoSolicitud{}
	contexto.BindJSON(solicitud)
	return solicitud
}

func (m *Mapeador) AgregarRegistroMantemientoVehiculoSolicitudARegistroMantemientoVehiculo(solicitud *capturaRegistroMantenimientoVehiculoModelos.AgregarRegistroMantenimientoVehiculoSolicitud) *capturaRegistroMantenimientoVehiculoModelos.RegistroMantenimientoVehiculo {
	registro := &capturaRegistroMantenimientoVehiculoModelos.RegistroMantenimientoVehiculo{}

	tipoDeRegistro := solicitud.ObtenerTipo()
	registro.AsignarTipo(tipoDeRegistro)

	fecha := solicitud.ObtenerFecha()
	registro.AsignarFecha(fecha)

	if kilometraje, err := strconv.Atoi(solicitud.ObtenerKilometraje()); err == nil {
		registro.AsignarKilometraje(kilometraje)
	} else {
		registro.AsignarKilometraje(0)
	}

	if litrosDeGasolina, err := strconv.ParseFloat(solicitud.ObtenerLitrosDeGasolina(), 64); err == nil {
		registro.AsignarLitrosDeGasolina(litrosDeGasolina)
	} else {
		registro.AsignarLitrosDeGasolina(0)
	}

	if importe, err := strconv.ParseFloat(solicitud.ObtenerImporte(), 64); err == nil {
		registro.AsignarImporte(importe)
	} else {
		registro.AsignarImporte(0)
	}

	concepto := solicitud.ObtenerConcepto()
	registro.AsignarConcepto(concepto)
	registro.AsignarPlacasVehiculo(solicitud.ObtenerPlacasVehiculo())

	return registro
}

func (m *Mapeador) GinContextAObtenerRegistroMantenimientoVehiculoPorNumeroDeRegistroSolicitud(contexto *gin.Context) *capturaRegistroMantenimientoVehiculoModelos.ObtenerRegistroDeMantenimientoDelVehiculoPorNumeroDeRegistroSolicitud {
	solicitud := &capturaRegistroMantenimientoVehiculoModelos.ObtenerRegistroDeMantenimientoDelVehiculoPorNumeroDeRegistroSolicitud{}

	contexto.BindJSON(solicitud)

	return solicitud
}

func (m *Mapeador) ObtenerRegistroMantenimientoVehiculoSolicitudAObtenerRegistroYVehiculoAsociadoPorNumeroDeRegistroSolicitud(solicitud *capturaRegistroMantenimientoVehiculoModelos.ObtenerRegistroDeMantenimientoDelVehiculoPorNumeroDeRegistroSolicitud) *capturaRegistroMantenimientoVehiculoModelos.ObtenerRegistroYVehiculoAsociadoPorNumeroDeRegistroSolicitud {
	solicitudBD := &capturaRegistroMantenimientoVehiculoModelos.ObtenerRegistroYVehiculoAsociadoPorNumeroDeRegistroSolicitud{NumDeRegistro: solicitud.ObtenerNumDeRegistro()}
	return solicitudBD
}

func (m *Mapeador) GinContextAEditarRegistroDeMantenimientoDelVehiculoSolicitud(contexto *gin.Context) *capturaRegistroMantenimientoVehiculoModelos.EditarRegistroDeMantenimientoDelVehiculoSolicitud {
	solicitud := &capturaRegistroMantenimientoVehiculoModelos.EditarRegistroDeMantenimientoDelVehiculoSolicitud{}
	contexto.BindJSON(solicitud)
	return solicitud
}

func (m *Mapeador) EditarRegistroDeMantenimientoDelVehiculoSolicitudAActualizarRegistroDeMantenimientoDelVehiculoSolicitud(solicitud *capturaRegistroMantenimientoVehiculoModelos.EditarRegistroDeMantenimientoDelVehiculoSolicitud) *conectorBDModelos.ActualizarRegistroDeMantenimientoDelVehiculoSolicitud {
	actualizarRegistroSolicitud := &conectorBDModelos.ActualizarRegistroDeMantenimientoDelVehiculoSolicitud{
		NumeroDeRegistro: solicitud.ObtenerNumeroDeRegistro(),
		Tipo:             solicitud.ObtenerTipo(),
		Fecha:            solicitud.ObtenerFecha(),
		LitrosDeGasolina: solicitud.ObtenerLitrosDeGasolina(),
		Kilometraje:      solicitud.ObtenerKilometraje(),
		Importe:          solicitud.ObtenerImporte(),
		Observaciones:    solicitud.ObtenerObservaciones(),
		Concepto:         solicitud.ObtenerConcepto(),
		PlacasVehiculo:   solicitud.ObtenerPlacasVehiculo(),
	}
	return actualizarRegistroSolicitud
}

func (m *Mapeador) ActualizarRegistroMantenimientoVehiculoRespuestaAEditarRegistroMantenimientoVehiculoRespuesta(actualizarRegistroRespuesta *conectorBDModelos.ActualizarRegistroMantenimientoVehiculoRespuesta) *capturaRegistroMantenimientoVehiculoModelos.EditarRegistroMantenimientoVehiculoRespuesta {
	editarRegistroRespuesta := &capturaRegistroMantenimientoVehiculoModelos.EditarRegistroMantenimientoVehiculoRespuesta{}
	editarRegistroRespuesta.AsignarErr(actualizarRegistroRespuesta.ObtenerErr())
	editarRegistroRespuesta.AsignarOk(actualizarRegistroRespuesta.ObtenerOk())
	return editarRegistroRespuesta
}
