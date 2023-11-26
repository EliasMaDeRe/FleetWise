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
