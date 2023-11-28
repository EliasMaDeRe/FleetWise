package controlador

import (
	"errors"
	conectorBDControlador "example/fleetwise/fuente/conectorbd/controlador"
	"strconv"

	"example/fleetwise/fuente/capturavehiculos/constantes"
	capturaVehiculosMapeador "example/fleetwise/fuente/capturavehiculos/mapeador"
	"example/fleetwise/modelos/capturavehiculos"
	capturaVehiculosModelos "example/fleetwise/modelos/capturavehiculos"
	conectorBDModelos "example/fleetwise/modelos/conectorbd"
)

type Controlador struct {
	ConectorBDControlador    *conectorBDControlador.Controlador
	CapturaVehiculosMapeador *capturaVehiculosMapeador.Mapeador
}

func (c *Controlador) AgregarVehiculo(solicitud *capturaVehiculosModelos.AgregarVehiculoSolicitud) *capturaVehiculosModelos.AgregarVehiculoRespuesta {
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

func (c *Controlador) validarAgregarVehiculosSolicitud(solicitud *capturaVehiculosModelos.AgregarVehiculoSolicitud) error {
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

func (c *Controlador) EditarVehiculo(solicitud *capturaVehiculosModelos.EditarVehiculoSolicitud) *capturaVehiculosModelos.EditarVehiculoRespuesta {
	respuesta := &capturaVehiculosModelos.EditarVehiculoRespuesta{}

	if solicitud == nil {
		respuesta.AsignarOk(false)
		respuesta.AsignarErr(errors.New(constantes.ERROR_SOLICITUD_NULA))
		return respuesta
	}

	if err := c.validarEditarVehiculosSolicitud(solicitud); err != nil {
		respuesta.AsignarOk(false)
		respuesta.AsignarErr(err)
		return respuesta
	}

	solicitudConectorBD := c.CapturaVehiculosMapeador.EditarVehiculosSolicitudAVehiculo(solicitud)

	c.ConectorBDControlador.EditarVehiculo(solicitudConectorBD)

	respuesta.AsignarOk(true)
	respuesta.AsignarErr(nil)

	return respuesta
}

func (c *Controlador) validarEditarVehiculosSolicitud(solicitud *capturaVehiculosModelos.EditarVehiculoSolicitud) error {
	if fechalanzamiento, err := strconv.Atoi(solicitud.ObtenerFechaDeLanzamientoNueva()); err != nil || fechalanzamiento <= 0 {
		return errors.New(constantes.ERROR_FECHALANZAMIENTO_NO_NATURAL)
	}
	obtenerVehiculoPorPlacasSolicitud := &conectorBDModelos.ObtenerVehiculoPorPlacasSolicitud{}
	obtenerVehiculoPorPlacasSolicitud.AsignarPlacas(solicitud.ObtenerPlacasNuevas())
	obtenerVehiculoPorSerieSolicitud := &conectorBDModelos.ObtenerVehiculoPorSerieSolicitud{}
	obtenerVehiculoPorSerieSolicitud.AsignarSerie(solicitud.ObtenerSerieNuevo())
	if c.ConectorBDControlador.ObtenerVehiculoPorPlacas(obtenerVehiculoPorPlacasSolicitud) != nil {
		return errors.New(constantes.ERROR_PLACAS_EXISTENTES_EN_BD)
	}
	obtenerVehiculoPorPlacasSolicitud = &conectorBDModelos.ObtenerVehiculoPorPlacasSolicitud{}
	obtenerVehiculoPorPlacasSolicitud.AsignarPlacas(solicitud.ObtenerPlacasActuales())
	if c.ConectorBDControlador.ObtenerVehiculoPorPlacas(obtenerVehiculoPorPlacasSolicitud) == nil {
		return errors.New(constantes.ERROR_VEHICULO_A_EDITAR_NO_EXISTE)
	}
	if c.ConectorBDControlador.ObtenerVehiculoPorSerie(obtenerVehiculoPorSerieSolicitud) != nil {
		return errors.New(constantes.ERROR_SERIE_EXISTENTE_EN_BD)
	}
	return nil
}

func (c *Controlador) ObtenerVehiculoPorPlacas(solicitud *capturaVehiculosModelos.ObtenerVehiculoPorPlacasSolicitud) *capturavehiculos.Vehiculo {

	solicitudConectorBD := &conectorBDModelos.ObtenerVehiculoPorPlacasSolicitud{}
	solicitudConectorBD.AsignarPlacas(solicitud.ObtenerPlacas())

	return c.ConectorBDControlador.ObtenerVehiculoPorPlacas(solicitudConectorBD)
}

func (c *Controlador) EliminarVehiculo(solicitud *capturavehiculos.EliminarVehiculoSolicitud) *capturavehiculos.EliminarVehiculoRespuesta {
	respuesta := &capturavehiculos.EliminarVehiculoRespuesta{}

	if solicitud == nil {
		respuesta.AsignarOk(false)
		respuesta.AsignarErr(errors.New(constantes.ERROR_SOLICITUD_NULA))
		return respuesta
	}

	solicitudConectorBD := &conectorBDModelos.ObtenerVehiculoPorPlacasSolicitud{}
	solicitudConectorBD.AsignarPlacas(solicitud.ObtenerPlacas())

	vehiculo := c.ConectorBDControlador.ObtenerVehiculoPorPlacas(solicitudConectorBD)

	if vehiculo != nil {
		c.ConectorBDControlador.EliminarVehiculo(vehiculo)
	}

	respuesta.AsignarOk(true)
	respuesta.AsignarErr(nil)

	return respuesta
}
