package controlador

import (
	"errors"
	"example/fleetwise/fuente/capturaserviciovehicular/constantes"
	capturaServicioVehicularMapeador "example/fleetwise/fuente/capturaserviciovehicular/mapeador"
	conectorBDControlador "example/fleetwise/fuente/conectorbd/controlador"
	capturaServicioVehicularModelos "example/fleetwise/modelos/capturaserviciovehicular"
	conectorBDModelos "example/fleetwise/modelos/conectorbd"

	"github.com/microcosm-cc/bluemonday"
)

type Controlador struct {
	ConectorBDControlador            *conectorBDControlador.Controlador
	CapturaServicioVehicularMapeador *capturaServicioVehicularMapeador.Mapeador
}

func (c *Controlador) AgregarServicioVehicular(solicitud *capturaServicioVehicularModelos.AgregarServicioVehicularSolicitud) *capturaServicioVehicularModelos.AgregarServicioVehicularRespuesta {
	respuesta := &capturaServicioVehicularModelos.AgregarServicioVehicularRespuesta{}

	if solicitud == nil {
		respuesta.AsignarOk(false)
		respuesta.AsignarErr(errors.New(constantes.ERROR_SOLICITUD_NULA))
		return respuesta
	}

	if err := c.validarAgregarServicioVehicularSolicitud(solicitud); err != nil {
		respuesta.AsignarOk(false)
		respuesta.AsignarErr(err)
		return respuesta
	}

	servicioVehicular := c.CapturaServicioVehicularMapeador.AgregarServicioVehicularSolicitudAServicioVehicular(solicitud)

	if guardarServicioVehicularRespuesta := c.ConectorBDControlador.GuardarServicioVehicular(servicioVehicular); guardarServicioVehicularRespuesta.ObtenerErr() != nil {
		respuesta.AsignarOk(false)
		respuesta.AsignarErr(guardarServicioVehicularRespuesta.ObtenerErr())
		return respuesta
	}

	respuesta.AsignarOk(true)
	respuesta.AsignarErr(nil)

	return respuesta
}

func (c *Controlador) validarAgregarServicioVehicularSolicitud(solicitud *capturaServicioVehicularModelos.AgregarServicioVehicularSolicitud) error {

	solicitud.AsignarNombre(bluemonday.StrictPolicy().Sanitize(solicitud.ObtenerNombre()))

	obtenerServicioVehicularPorNombreSolicitud := &conectorBDModelos.ObtenerServicioVehicularPorNombreSolicitud{}
	obtenerServicioVehicularPorNombreSolicitud.AsignarNombre(solicitud.ObtenerNombre())

	if c.ConectorBDControlador.ObtenerServicioVehicularPorNombre(obtenerServicioVehicularPorNombreSolicitud) != nil {
		return errors.New(constantes.ERROR_NOMBRE_DE_SERVICIO_VEHICULAR_EXISTE_EN_BD)
	}

	return nil
}

func (c *Controlador) EliminarServicioVehicular(solicitud *capturaServicioVehicularModelos.EliminarServicioVehicularSolicitud) *capturaServicioVehicularModelos.EliminarServicioVehicularRespuesta {
	respuesta := &capturaServicioVehicularModelos.EliminarServicioVehicularRespuesta{}

	if solicitud == nil {
		respuesta.AsignarOk(false)
		respuesta.AsignarErr(errors.New(constantes.ERROR_SOLICITUD_NULA))
		return respuesta
	}

	servicioVehicular := c.CapturaServicioVehicularMapeador.EliminarServicioVehicularSolicitudAServicioVehicular(solicitud)

	c.ConectorBDControlador.EliminarServicioVehicular(servicioVehicular)

	respuesta.AsignarOk(true)
	respuesta.AsignarErr(nil)

	return respuesta
}

func (c *Controlador) EditarServicioVehicular(solicitud *capturaServicioVehicularModelos.EditarServicioVehicularSolicitud) *capturaServicioVehicularModelos.EditarServicioVehicularRespuesta {
	respuesta := &capturaServicioVehicularModelos.EditarServicioVehicularRespuesta{}

	if solicitud == nil {
		respuesta.AsignarOk(false)
		respuesta.AsignarErr(errors.New(constantes.ERROR_SOLICITUD_NULA))
		return respuesta
	}

	if err := c.validarEditarServicioVehicularSolicitud(solicitud); err != nil {
		respuesta.AsignarOk(false)
		respuesta.AsignarErr(err)
		return respuesta
	}

	editarServicioVehicularSolicitudDeConectorBD := c.CapturaServicioVehicularMapeador.EditarServicioVehicularSolicitudAServicioVehicularDeConectorBD(solicitud)

	c.ConectorBDControlador.EditarServicioVehicular(editarServicioVehicularSolicitudDeConectorBD)

	respuesta.AsignarOk(true)
	respuesta.AsignarErr(nil)

	return respuesta
}

func (c *Controlador) validarEditarServicioVehicularSolicitud(solicitud *capturaServicioVehicularModelos.EditarServicioVehicularSolicitud) error {

	solicitud.AsignarNombreActual(bluemonday.StrictPolicy().Sanitize(solicitud.ObtenerNombreActual()))
	solicitud.AsignarNuevoNombre(bluemonday.StrictPolicy().Sanitize(solicitud.ObtenerNuevoNombre()))

	obtenerServicioVehicularPorNombreSolicitud := &conectorBDModelos.ObtenerServicioVehicularPorNombreSolicitud{}
	obtenerServicioVehicularPorNombreSolicitud.AsignarNombre(solicitud.ObtenerNuevoNombre())

	if c.ConectorBDControlador.ObtenerServicioVehicularPorNombre(obtenerServicioVehicularPorNombreSolicitud) != nil {
		return errors.New(constantes.ERROR_NOMBRE_DE_SERVICIO_VEHICULAR_EXISTE_EN_BD)
	}

	return nil
}

func (c *Controlador) ObtenerServiciosVehiculares() []capturaServicioVehicularModelos.ServicioVehicular {
	return c.ConectorBDControlador.ObtenerServiciosVehiculares()
}
