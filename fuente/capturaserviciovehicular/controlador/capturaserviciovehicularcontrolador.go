package controlador

import (
	"errors"
	"example/fleetwise/fuente/capturaserviciovehicular/constantes"
	servicioVehicularMapeador "example/fleetwise/fuente/capturaserviciovehicular/mapeador"
	conectorBDControlador "example/fleetwise/fuente/conectorbd/controlador"
	"example/fleetwise/modelos/capturaserviciovehicular"
	conectorBDmodelos "example/fleetwise/modelos/conectorbd"
)

type Controlador struct {
	ConectorBDControlador     *conectorBDControlador.Controlador
	ServicioVehicularMapeador *servicioVehicularMapeador.Mapeador
}

func (c *Controlador) AgregarServicioVehicular(solicitud *capturaserviciovehicular.AgregarServicioVehicularSolicitud) *capturaserviciovehicular.AgregarServicioVehicularRespuesta {
	respuesta := &capturaserviciovehicular.AgregarServicioVehicularRespuesta{}

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

	servicioVehicular := c.ServicioVehicularMapeador.AgregarServicioVehicularSolicitudAServicioVehicular(solicitud)

	if guardarServicioVehicularRespuesta := c.ConectorBDControlador.GuardarServicioVehicular(servicioVehicular); guardarServicioVehicularRespuesta.ObtenerErr() != nil {
		respuesta.AsignarOk(false)
		respuesta.AsignarErr(guardarServicioVehicularRespuesta.ObtenerErr())
		return respuesta
	}

	respuesta.AsignarOk(true)
	respuesta.AsignarErr(nil)

	return respuesta
}

func (c *Controlador) validarAgregarServicioVehicularSolicitud(solicitud *capturaserviciovehicular.AgregarServicioVehicularSolicitud) error {

	obtenerServicioVehicularPorNombreSolicitud := &conectorBDmodelos.ObtenerServicioVehicularPorNombreSolicitud{}
	obtenerServicioVehicularPorNombreSolicitud.AsignarNombre(solicitud.ObtenerNombre())

	if c.ConectorBDControlador.ObtenerServicioVehicularPorNombre(obtenerServicioVehicularPorNombreSolicitud) != nil {
		return errors.New(constantes.ERROR_NOMBRE_DE_SERVICIO_VEHICULAR_EXISTE_EN_BD)
	}

	return nil
}

func (c *Controlador) EliminarServicioVehicular(solicitud *capturaserviciovehicular.EliminarServicioVehicularSolicitud) *capturaserviciovehicular.EliminarServicioVehicularRespuesta {
	respuesta := &capturaserviciovehicular.EliminarServicioVehicularRespuesta{}

	if solicitud == nil {
		respuesta.AsignarOk(false)
		respuesta.AsignarErr(errors.New(constantes.ERROR_SOLICITUD_NULA))
		return respuesta
	}

	servicioVehicular := c.ServicioVehicularMapeador.EliminarServicioVehicularSolicitudAServicioVehicular(solicitud)

	c.ConectorBDControlador.EliminarServicioVehicular(servicioVehicular)

	respuesta.AsignarOk(true)
	respuesta.AsignarErr(nil)

	return respuesta
}

func (c *Controlador) EditarServicioVehicular(solicitud *capturaserviciovehicular.EditarServicioVehicularSolicitud) *capturaserviciovehicular.EditarServicioVehicularRespuesta {
	respuesta := &capturaserviciovehicular.EditarServicioVehicularRespuesta{}

	if solicitud == nil {
		respuesta.AsignarOk(false)
		respuesta.AsignarErr(errors.New(constantes.ERROR_SOLICITUD_NULA))
		return respuesta
	}

	editarServicioVehicularSolicitudDeConectorBD := c.ServicioVehicularMapeador.EditarServicioVehicularSolicitudAServicioVehicularDeConectorBD(solicitud)

	c.ConectorBDControlador.EditarServicioVehicular(editarServicioVehicularSolicitudDeConectorBD)

	respuesta.AsignarOk(true)
	respuesta.AsignarErr(nil)

	return respuesta
}

func (c *Controlador) ObtenerServiciosVehiculares() []capturaserviciovehicular.ServicioVehicular {
	return c.ConectorBDControlador.ObtenerServiciosVehiculares()
}
