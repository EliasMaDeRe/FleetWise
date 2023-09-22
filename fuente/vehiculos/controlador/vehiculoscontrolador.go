package controlador

import (
	"errors"
	conectorBDControlador "example/fleetwise/fuente/conectorbd/controlador"
	"strconv"

	"example/fleetwise/fuente/vehiculos/constantes"
	vehiculosMapeador "example/fleetwise/fuente/vehiculos/mapeador"
	"example/fleetwise/modelos/conectorbd"
	vehiculosModelos "example/fleetwise/modelos/vehiculos"
	//conectorControlador "example/fleetwise/ConectorBD/Controlador"
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
	}

	var err error
	if err = c.validarAgregarVehiculosSolicitud(solicitud); err != nil {
		respuesta.AsignarOk(false)
		respuesta.AsignarErr(err)
		return respuesta
	}

	vehiculo := c.VehiculosMapeador.AgregarVehiculosSolicitudAVehiculo(solicitud)

	if err = c.ConectorBDControlador.AgregarVehiculo(vehiculo); err != nil {
		respuesta.AsignarOk(false)
		respuesta.AsignarErr(err)
		return respuesta
	}

	respuesta.AsignarOk(true)
	respuesta.AsignarErr(nil)

	return respuesta
}

func (c *Controlador) validarAgregarVehiculosSolicitud(solicitud *vehiculosModelos.AgregarVehiculosSolicitud) error {
	if anualidad, err := strconv.Atoi(solicitud.ObtenerAnualidad()); err != nil || anualidad <= 0 {
		return errors.New(constantes.ERROR_ANUALIDAD_NO_NATURAL)
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
