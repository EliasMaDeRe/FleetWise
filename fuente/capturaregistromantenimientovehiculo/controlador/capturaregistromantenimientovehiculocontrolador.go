package controlador

import (
	"errors"
	conectorBDControlador "example/fleetwise/fuente/conectorbd/controlador"
	"strconv"

	"example/fleetwise/fuente/capturaregistromantenimientovehiculo/constantes"
	registroMantenimientoVehiculoMapeador "example/fleetwise/fuente/capturaregistromantenimientovehiculo/mapeador"
	registroMantenimientoVehiculoModelos "example/fleetwise/modelos/capturaregistromantenimientovehiculo"
	"example/fleetwise/modelos/conectorbd"
)

type Controlador struct {
	ConectorBDControlador                 *conectorBDControlador.Controlador
	RegistroMantenimientoVehiculoMapeador *registroMantenimientoVehiculoMapeador.Mapeador
}

func (c *Controlador) SeleccionarVehiculoDeNuevoRegistro(solicitud *registroMantenimientoVehiculoModelos.SeleccionarVehiculoParaNuevoRegistroSolicitud) *registroMantenimientoVehiculoModelos.SeleccionarVehiculoParaNuevoRegistroRespuesta {
	respuesta := &registroMantenimientoVehiculoModelos.SeleccionarVehiculoParaNuevoRegistroRespuesta{}

	if solicitud == nil {
		respuesta.AsignarOk(false)
		respuesta.AsignarErr(errors.New(constantes.ERROR_SOLICITUD_NULA))
	}

	var err error
	if err = c.validarSeleccionarVehiculoParaRegistro(solicitud); err != nil {
		respuesta.AsignarOk(false)
		respuesta.AsignarErr(err)
		return respuesta
	}

	respuesta.AsignarOk(true)
	respuesta.AsignarErr(nil)

	return respuesta
}

func (c *Controlador) validarSeleccionarVehiculoParaRegistro(solicitud *registroMantenimientoVehiculoModelos.SeleccionarVehiculoParaNuevoRegistroSolicitud) error {
	obtenerVehiculoPorPlacasSolicitud := &conectorbd.ObtenerVehiculoPorPlacasSolicitud{}
	obtenerVehiculoPorPlacasSolicitud.AsignarPlacas(solicitud.ObtenerPlacas())
	if c.ConectorBDControlador.ObtenerVehiculoPorPlacas(obtenerVehiculoPorPlacasSolicitud) != nil {
		return errors.New(constantes.ERROR_PLACAS_INEXISTENTES_EN_BD)
	}
	return nil
}

func (c *Controlador) AgregarRegistroMantemientoVehiculo(solicitud *registroMantenimientoVehiculoModelos.AgregarRegistroMantenimientoVehiculoSolicitud) *registroMantenimientoVehiculoModelos.AgregarRegistroMantenimientoVehiculoRespuesta {
	respuesta := &registroMantenimientoVehiculoModelos.AgregarRegistroMantenimientoVehiculoRespuesta{}

	if solicitud == nil {
		respuesta.AsignarOk(false)
		respuesta.AsignarErr(errors.New(constantes.ERROR_SOLICITUD_NULA))
	}

	var err error
	if err = c.validarAgregarRegistroMantemientoVehiculo(solicitud); err != nil {
		respuesta.AsignarOk(false)
		respuesta.AsignarErr(err)
		return respuesta
	}

	registro := c.RegistroMantenimientoVehiculoMapeador.AgregaRegistroMantemientoVehiculoSolicitudARegistroMantemientoVehiculo(solicitud)

	if c.ConectorBDControlador.GuardarRegistro(registro).ObtenerErr() != nil {
		respuesta.AsignarOk(false)
		respuesta.AsignarErr(err)
		return respuesta
	}

	respuesta.AsignarOk(true)
	respuesta.AsignarErr(nil)

	return respuesta
}

func (c *Controlador) validarAgregarRegistroMantemientoVehiculo(solicitud *registroMantenimientoVehiculoModelos.AgregarRegistroMantenimientoVehiculoSolicitud) error {
	if fecha, err := strconv.Atoi(solicitud.ObtenerFecha()); err != nil || fecha <= 0 {
		return errors.New(constantes.ERROR_NO_NATURAL)
	}
	if kilometraje, err := strconv.Atoi(solicitud.ObtenerKilometraje()); err != nil || kilometraje <= 0 {
		return errors.New(constantes.ERROR_NO_NATURAL)
	}
	if gasolina, err := strconv.Atoi(solicitud.ObtenerLitrosGasolina()); err != nil || gasolina <= 0 {
		return errors.New(constantes.ERROR_NO_NATURAL)
	}
	if importe, err := strconv.Atoi(solicitud.ObtenerImporte()); err != nil || importe <= 0 {
		return errors.New(constantes.ERROR_NO_NATURAL)
	}
	return nil
}

/*
func (c *Controlador) EditarRegistroMantemientoVehiculo(solicitud *visualizacionHistorialRegistrosControlador.EditaregistroMantenimientoVehiculo) *registroMantenimientoVehiculoModelos.EditarRegistroMantenimientoVehiculoRespuesta {
	//implementacipon futura por Pablo
	return nil
}

func (c *Controlador) validarEditarRegistroMantemientoVehiculo(solicitud *visualizacionHistorialRegistrosControlador.EditaregistroMantenimientoVehiculo) error {
	//implementacipon futura por Pablo
	return nil
}

func (c *Controlador) BorrarRegistroMantemientoVehiculo(solicitud *visualizacionHistorialRegistrosControlador.BorrarRegistroMantenimientoVehiculo) *registroMantenimientoVehiculoModelos.BorrarRegistroMantenimientoVehiculoRespuesta {
	//implementacipon futura por Pablo
	return nil
}

func (c *Controlador) validarBorrarRegistroMantemientoVehiculo(solicitud *visualizacionHistorialRegistrosControlador.BorrarRegistroMantenimientoVehiculo) error {
	//implementacipon futura por Pablo
	return nil
}*/
