package controlador

import (
	"errors"
	conectorBDControlador "example/fleetwise/fuente/conectorbd/controlador"
	"strconv"

	"example/fleetwise/fuente/capturavehiculos/constantes"
	capturaVehiculosMapeador "example/fleetwise/fuente/capturavehiculos/mapeador"
	capturaVehiculosModelos "example/fleetwise/modelos/capturavehiculos"
	conectorBDModelos "example/fleetwise/modelos/conectorbd"
)

type Controlador struct {
	ConectorBDControlador *conectorBDControlador.Controlador
	CapturaVehiculosMapeador     *capturaVehiculosMapeador.Mapeador
}

func (c *Controlador) AgregarVehiculo(solicitud *capturaVehiculosModelos.AgregarVehiculosSolicitud) *capturaVehiculosModelos.AgregarVehiculoRespuesta {
	respuesta := &capturaVehiculosModelos.AgregarVehiculoRespuesta{}

	if solicitud == nil {
		respuesta.AsignarOk(false)
		respuesta.AsignarErr(errors.New(constantes.ERROR_SOLICITUD_NULA))
		return respuesta
	}

	if err := c.validarAgregarVehiculosSolicitud(solicitud); err != nil {
		respuesta.AsignarOk(false)
		respuesta.AsignarErr(err)
		return respuesta
	}

	vehiculo := c.CapturaVehiculosMapeador.AgregarVehiculosSolicitudAVehiculo(solicitud)

	if guardarVehiculoRespuesta := c.ConectorBDControlador.GuardarVehiculo(vehiculo); guardarVehiculoRespuesta.ObtenerErr() != nil {
		respuesta.AsignarOk(false)
		respuesta.AsignarErr(guardarVehiculoRespuesta.ObtenerErr())
		return respuesta
	}

	respuesta.AsignarOk(true)
	respuesta.AsignarErr(nil)

	return respuesta
}

func (c *Controlador) validarAgregarVehiculosSolicitud(solicitud *capturaVehiculosModelos.AgregarVehiculosSolicitud) error {
	if fechalanzamiento, err := strconv.Atoi(solicitud.ObtenerFechaLanzamiento()); err != nil || fechalanzamiento <= 0 {
		return errors.New(constantes.ERROR_FECHALANZAMIENTO_NO_NATURAL)
	}
	obtenerVehiculoPorPlacasSolicitud := &conectorBDModelos.ObtenerVehiculoPorPlacasSolicitud{}
	obtenerVehiculoPorPlacasSolicitud.AsignarPlacas(solicitud.ObtenerPlacas())
	obtenerVehiculoPorSerieSolicitud := &conectorBDModelos.ObtenerVehiculoPorSerieSolicitud{}
	obtenerVehiculoPorSerieSolicitud.AsignarSerie(solicitud.ObtenerSerie())
	if c.ConectorBDControlador.ObtenerVehiculoPorPlacas(obtenerVehiculoPorPlacasSolicitud) != nil {
		return errors.New(constantes.ERROR_PLACAS_EXISTENTES_EN_BD)
	}
	if c.ConectorBDControlador.ObtenerVehiculoPorSerie(obtenerVehiculoPorSerieSolicitud) != nil {
		return errors.New(constantes.ERROR_SERIE_EXISTENTE_EN_BD)
	}
	return nil
}
