package controlador

import (
	"errors"
	conectorBDControlador "example/fleetwise/fuente/conectorbd/controlador"
	"example/fleetwise/fuente/vehiculos/constantes"
	"example/fleetwise/modelos/conectorbd"
	vehiculosModelos "example/fleetwise/modelos/vehiculos"
)

type Controlador struct {
	ConectorBDControlador conectorBDControlador.Controlador
}

func (c *Controlador) AgregarVehiculo(solicitud *vehiculosModelos.AgregarVehiculosSolicitud) *vehiculosModelos.AgregarVehiculoRespuesta {

	respuesta := &vehiculosModelos.AgregarVehiculoRespuesta{}

	if solicitud == nil {
		respuesta.AsignarOk(false)
		respuesta.AsignarErr(errors.New(constantes.ERROR_SOLICITUD_NULA))
	}

	if err := c.validarAgregarVehiculosSolicitud(solicitud); err != nil {
		respuesta.AsignarOk(false)
		respuesta.AsignarErr(err)
	}

	return respuesta
}

func (c *Controlador) validarAgregarVehiculosSolicitud(solicitud *vehiculosModelos.AgregarVehiculosSolicitud) error {
	if solicitud.ObtenerAnualidad() <= 0 {
		return errors.New(constantes.ERROR_ANUALIDAD_NO_NATURAL)
	}
	obtenerVehiculoPorPlacasSolicitud := &conectorbd.ObtenerVehiculoPorPlacasSolicitud{}
	obtenerVehiculoPorPlacasSolicitud.AsignarPlacas(solicitud.ObtenerPlacas())
	if c.ConectorBDControlador.ObtenerVehiculoPorPlacas(obtenerVehiculoPorPlacasSolicitud) == nil {

	}
	return nil
}
