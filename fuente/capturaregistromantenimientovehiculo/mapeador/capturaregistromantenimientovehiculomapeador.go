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

func (m *Mapeador) GinContextAAgregarRegistroMantemientoVehiculoSolicitud(contexto *gin.Context) *capturaRegistroMantenimientoVehiculoModelos.AgregarRegistroDeMantenimientoDeVehiculoSolicitud {
	solicitud := &capturaRegistroMantenimientoVehiculoModelos.AgregarRegistroDeMantenimientoDeVehiculoSolicitud{}
	contexto.BindJSON(solicitud)
	return solicitud
}

func (m *Mapeador) AgregarRegistroMantemientoVehiculoSolicitudARegistroMantemientoVehiculo(solicitud *capturaRegistroMantenimientoVehiculoModelos.AgregarRegistroDeMantenimientoDeVehiculoSolicitud) *capturaRegistroMantenimientoVehiculoModelos.RegistroDeMantenimientoDeVehiculo {
	registro := &capturaRegistroMantenimientoVehiculoModelos.RegistroDeMantenimientoDeVehiculo{}

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

	registro.AsignarObservaciones(solicitud.ObtenerObservaciones())

	if tipoDeRegistro == "carga de combustible" {
		registro.AsignarConcepto("N/A")
	} else {
		registro.AsignarConcepto(solicitud.ObtenerConcepto())
	}

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

func (m *Mapeador) GinContextAEditarRegistroDeMantenimientoDelVehiculoSolicitud(contexto *gin.Context) *capturaRegistroMantenimientoVehiculoModelos.EditarRegistroDeMantenimientoDeVehiculoSolicitud {
	solicitud := &capturaRegistroMantenimientoVehiculoModelos.EditarRegistroDeMantenimientoDeVehiculoSolicitud{}
	contexto.BindJSON(solicitud)
	return solicitud
}

func (m *Mapeador) EditarRegistroDeMantenimientoDeVehiculoSolicitudASolicitudBD(solicitud *capturaRegistroMantenimientoVehiculoModelos.EditarRegistroDeMantenimientoDeVehiculoSolicitud) *conectorBDModelos.EditarRegistroDeMantenimientoDelVehiculoSolicitud {
	actualizarRegistroSolicitud := &conectorBDModelos.EditarRegistroDeMantenimientoDelVehiculoSolicitud{
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
	if actualizarRegistroSolicitud.ObtenerTipo() == "carga de combustible" {
		actualizarRegistroSolicitud.AsignarConcepto("N/A")
	} else {
		actualizarRegistroSolicitud.AsignarLitrosDeGasolina(0)
	}
	return actualizarRegistroSolicitud
}
