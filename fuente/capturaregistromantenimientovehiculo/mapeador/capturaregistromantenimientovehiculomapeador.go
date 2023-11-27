package mapeador

import (
	registroMantenimientoVehiculo "example/fleetwise/modelos/capturaregistromantenimientovehiculo"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type Mapeador struct {
}

func (m *Mapeador) GinContextASeleccionarVehiculoParaNuevoRegistroSolicitud(contexto *gin.Context) *registroMantenimientoVehiculo.SeleccionarVehiculoParaNuevoRegistroSolicitud {
	solicitud := &registroMantenimientoVehiculo.SeleccionarVehiculoParaNuevoRegistroSolicitud{}
	if strings.Contains(contexto.Request.URL.String(), "?placas=") {
		divisionesURL := strings.SplitAfter(contexto.Request.URL.String(), "?placas=")
		solicitud.AsignarPlacas(divisionesURL[1])
		return solicitud
	}
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

	return registro
}
