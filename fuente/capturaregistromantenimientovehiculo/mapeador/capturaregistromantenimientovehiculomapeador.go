package mapeador

import (
	capturaRegistroMantenimientoVehiculoModelos "example/fleetwise/modelos/capturaregistromantenimientovehiculo"
	conectorBDModelos "example/fleetwise/modelos/conectorbd"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Mapeador struct {
}

func (m *Mapeador) GinContextASeleccionarVehiculoParaNuevoRegistroSolicitud(contexto *gin.Context) *capturaRegistroMantenimientoVehiculoModelos.SeleccionarVehiculoParaNuevoRegistroSolicitud {
	solicitud := &capturaRegistroMantenimientoVehiculoModelos.SeleccionarVehiculoParaNuevoRegistroSolicitud{}
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

func (m *Mapeador) GinContextAObtenerRegistroMantenimientoVehicularPorNumeroDeRegistroSolicitud(contexto *gin.Context) *capturaRegistroMantenimientoVehiculoModelos.ObtenerRegistroDeMantenimientoDelVehiculoPorNumeroDeRegistroSolicitud{
	solicitud := &capturaRegistroMantenimientoVehiculoModelos.ObtenerRegistroDeMantenimientoDelVehiculoPorNumeroDeRegistroSolicitud{}

	contexto.BindJSON(solicitud)

	return solicitud
}

func(m *Mapeador) ObtenerRegistroMantenimientoVehiculoSolicitudAObtenerRegistroYVehiculoAsociadoPorNumeroDeRegistroSolicitud(solicitud *capturaRegistroMantenimientoVehiculoModelos.ObtenerRegistroDeMantenimientoDelVehiculoPorNumeroDeRegistroSolicitud) *capturaRegistroMantenimientoVehiculoModelos.ObtenerRegistroYVehiculoAsociadoPorNumeroDeRegistroSolicitud{
	solicitudBD := &capturaRegistroMantenimientoVehiculoModelos.ObtenerRegistroYVehiculoAsociadoPorNumeroDeRegistroSolicitud{NumDeRegistro: solicitud.ObtenerNumDeRegistro()}
	return solicitudBD
}

func (m *Mapeador) GinContextAEditarRegistroDeMantenimientoDelVehiculoSolicitud(contexto *gin.Context)  *capturaRegistroMantenimientoVehiculoModelos.EditarRegistroDeMantenimientoDelVehiculoSolicitud{
	solicitud := &capturaRegistroMantenimientoVehiculoModelos.EditarRegistroDeMantenimientoDelVehiculoSolicitud{}
	contexto.BindJSON(solicitud)
	return solicitud
}

func (m *Mapeador) EditarRegistroDeMantenimientoDelVehiculoSolicitudAActualizarRegistroDeMantenimientoDelVehiculoSolicitud(solicitud *capturaRegistroMantenimientoVehiculoModelos.EditarRegistroDeMantenimientoDelVehiculoSolicitud) (*conectorBDModelos.ActualizarRegistroDeMantenimientoDelVehiculoSolicitud){
	actualizarRegistroSolicitud := &conectorBDModelos.ActualizarRegistroDeMantenimientoDelVehiculoSolicitud{
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

func (m *Mapeador) ActualizarRegistroMantenimientoVehicularRespuestaAEditarRegistroMantenimientoVehicularRespuesta(actualizarRegistroRespuesta *conectorBDModelos.ActualizarRegistroMantenimientoVehicularRespuesta) *capturaRegistroMantenimientoVehiculoModelos.EditarRegistroMantenimientoVehicularRespuesta{
	editarRegistroRespuesta := &capturaRegistroMantenimientoVehiculoModelos.EditarRegistroMantenimientoVehicularRespuesta{
		Ok: actualizarRegistroRespuesta.ObtenerOk(),
		Err: actualizarRegistroRespuesta.ObtenerError(),
	}
	return editarRegistroRespuesta;
}
