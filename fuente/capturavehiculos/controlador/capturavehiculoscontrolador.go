package controlador

import (
	"errors"
	conectorBDControlador "example/fleetwise/fuente/conectorbd/controlador"
	"strconv"

	"example/fleetwise/fuente/capturavehiculos/constantes"
	vehiculosMapeador "example/fleetwise/fuente/capturavehiculos/mapeador"
	vehiculosModelos "example/fleetwise/modelos/capturavehiculos"
	"example/fleetwise/modelos/conectorbd"
)

type Controlador struct {
	ConectorBDControlador *conectorBDControlador.Controlador
	VehiculosMapeador     *vehiculosMapeador.Mapeador
}

func (c *Controlador) AgregarVehiculo(solicitud *vehiculosModelos.AgregarVehiculosSolicitud) *vehiculosModelos.AgregarVehiculoRespuesta {
	respuesta := &vehiculosModelos.AgregarVehiculoRespuesta{}

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

	vehiculo := c.VehiculosMapeador.AgregarVehiculosSolicitudAVehiculo(solicitud)

	if guardarVehiculoRespuesta := c.ConectorBDControlador.GuardarVehiculo(vehiculo); guardarVehiculoRespuesta.ObtenerErr() != nil {
		respuesta.AsignarOk(false)
		respuesta.AsignarErr(guardarVehiculoRespuesta.ObtenerErr())
		return respuesta
	}

	respuesta.AsignarOk(true)
	respuesta.AsignarErr(nil)

	return respuesta
}

func (c *Controlador) validarAgregarVehiculosSolicitud(solicitud *vehiculosModelos.AgregarVehiculosSolicitud) error {
	if fechalanzamiento, err := strconv.Atoi(solicitud.ObtenerFechaLanzamiento()); err != nil || fechalanzamiento <= 0 {
		return errors.New(constantes.ERROR_FECHALANZAMIENTO_NO_NATURAL)
	}
	obtenerVehiculoPorPlacasSolicitud := &conectorbd.ObtenerVehiculoPorPlacasSolicitud{}
	obtenerVehiculoPorPlacasSolicitud.AsignarPlacas(solicitud.ObtenerPlacas())
	obtenerVehiculoPorSerieSolicitud := &conectorbd.ObtenerVehiculoPorSerieSolicitud{}
	obtenerVehiculoPorSerieSolicitud.AsignarSerie(solicitud.ObtenerSerie())
	if c.ConectorBDControlador.ObtenerVehiculoPorPlacas(obtenerVehiculoPorPlacasSolicitud) != nil {
		return errors.New(constantes.ERROR_PLACAS_EXISTENTES_EN_BD)
	}
	if c.ConectorBDControlador.ObtenerVehiculoPorSerie(obtenerVehiculoPorSerieSolicitud) != nil {
		return errors.New(constantes.ERROR_SERIE_EXISTENTE_EN_BD)
	}
	return nil
}
