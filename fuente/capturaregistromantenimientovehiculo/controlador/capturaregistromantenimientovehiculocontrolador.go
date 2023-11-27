package controlador

import (
	"errors"
	servicioVehicularControlador "example/fleetwise/fuente/capturaserviciovehicular/controlador"
	"fmt"
	"strconv"
	"time"

	"example/fleetwise/fuente/capturaregistromantenimientovehiculo/constantes"
	registroMantenimientoVehiculoMapeador "example/fleetwise/fuente/capturaregistromantenimientovehiculo/mapeador"
	registroMantenimientoVehiculoModelos "example/fleetwise/modelos/capturaregistromantenimientovehiculo"
	"example/fleetwise/modelos/capturaserviciovehicular"
	"example/fleetwise/modelos/conectorbd"
)

type Controlador struct {
	RegistroMantenimientoVehiculoMapeador *registroMantenimientoVehiculoMapeador.Mapeador
	ServicioVehicularControlador          *servicioVehicularControlador.Controlador
}

func (c *Controlador) SeleccionarVehiculoParaNuevoRegistro(solicitud *registroMantenimientoVehiculoModelos.SeleccionarVehiculoParaNuevoRegistroSolicitud) *registroMantenimientoVehiculoModelos.SeleccionarVehiculoParaNuevoRegistroRespuesta {
	respuesta := &registroMantenimientoVehiculoModelos.SeleccionarVehiculoParaNuevoRegistroRespuesta{}

	if solicitud == nil {
		respuesta.AsignarOk(false)
		respuesta.AsignarErr(errors.New(constantes.ERROR_SOLICITUD_NULA))
		return respuesta
	}

	if err := c.validarSeleccionarVehiculoParaRegistro(solicitud); err != nil {
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
	if c.ServicioVehicularControlador.ConectorBDControlador.ObtenerVehiculoPorPlacas(obtenerVehiculoPorPlacasSolicitud) == nil {
		return errors.New(constantes.ERROR_PLACAS_INEXISTENTES_EN_BD)
	}
	return nil
}

func (c *Controlador) AgregarRegistroMantemientoVehiculo(solicitud *registroMantenimientoVehiculoModelos.AgregarRegistroMantenimientoVehiculoSolicitud) *registroMantenimientoVehiculoModelos.AgregarRegistroMantenimientoVehiculoRespuesta {
	respuesta := &registroMantenimientoVehiculoModelos.AgregarRegistroMantenimientoVehiculoRespuesta{}

	if solicitud == nil {
		respuesta.AsignarOk(false)
		respuesta.AsignarErr(errors.New(constantes.ERROR_PLACAS_INEXISTENTES_EN_BD))
		return respuesta
	}

	if err := c.validarAgregarRegistroMantemientoVehiculo(solicitud); err != nil {
		respuesta.AsignarOk(false)
		respuesta.AsignarErr(err)
		return respuesta
	}

	registro := c.RegistroMantenimientoVehiculoMapeador.AgregarRegistroMantemientoVehiculoSolicitudARegistroMantemientoVehiculo(solicitud)

	if guardarRegistroMantenimientoVehiculoRespuesta := c.ServicioVehicularControlador.ConectorBDControlador.GuardarRegistro(registro); guardarRegistroMantenimientoVehiculoRespuesta.ObtenerErr() != nil {
		respuesta.AsignarOk(false)
		respuesta.AsignarErr(guardarRegistroMantenimientoVehiculoRespuesta.ObtenerErr())
		return respuesta
	}

	respuesta.AsignarOk(true)
	respuesta.AsignarErr(nil)

	return respuesta
}

func (c *Controlador) validarAgregarRegistroMantemientoVehiculo(solicitud *registroMantenimientoVehiculoModelos.AgregarRegistroMantenimientoVehiculoSolicitud) error {

	_, err := time.Parse("02/01/2006", solicitud.ObtenerFecha()) // DD/MM/AAAA
	if err != nil {
		return errors.New(constantes.ERROR_FECHA_FORMATO_INVALIDO)
	}

	tipo := solicitud.ObtenerTipo()

	if tipo == "carga_combustible" {
		if gasolina, err := strconv.ParseFloat(solicitud.ObtenerLitrosDeGasolina(), 64); err != nil || gasolina <= 0 {
			return errors.New(constantes.ERROR_LITROS_GASOLINA_NO_ES_NUMERO)
		}
	}

	if kilometraje, err := strconv.Atoi(solicitud.ObtenerKilometraje()); err != nil || kilometraje <= 0 {
		fmt.Println(kilometraje)
		return errors.New(constantes.ERROR_KILOMETRAJE_NO_VALIDO)
	}

	if importe, err := strconv.ParseFloat(solicitud.ObtenerImporte(), 64); err != nil || importe <= 0 {
		fmt.Println(importe)
		return errors.New(constantes.ERROR_IMPORTE_NO_VALIDO)
	}

	return nil
}

func (c *Controlador) ObtenerServiciosVehicularesParaNuevoRegistro() []capturaserviciovehicular.ServicioVehicular {
	return c.ServicioVehicularControlador.ObtenerServiciosVehiculares()
}
