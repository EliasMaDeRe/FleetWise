package controlador

import (
	"errors"
	servicioVehicularControlador "example/fleetwise/fuente/capturaserviciovehicular/controlador"
	conectorBDControlador "example/fleetwise/fuente/conectorbd/controlador"
	"time"

	"example/fleetwise/fuente/capturaregistromantenimientovehiculo/constantes"
	registroMantenimientoVehiculoMapeador "example/fleetwise/fuente/capturaregistromantenimientovehiculo/mapeador"
	registroMantenimientoVehiculoModelos "example/fleetwise/modelos/capturaregistromantenimientovehiculo"
	"example/fleetwise/modelos/capturaserviciovehicular"
	"example/fleetwise/modelos/conectorbd"
)

type Controlador struct {
	ConectorBDControlador                 *conectorBDControlador.Controlador
	RegistroMantenimientoVehiculoMapeador *registroMantenimientoVehiculoMapeador.Mapeador
	ServicioVehicularControlador          *servicioVehicularControlador.Controlador
}

func (c *Controlador) SeleccionarVehiculoParaNuevoRegistro(solicitud *registroMantenimientoVehiculoModelos.SeleccionarVehiculoParaNuevoRegistroSolicitud) *registroMantenimientoVehiculoModelos.SeleccionarVehiculoParaNuevoRegistroRespuesta {
	respuesta := &registroMantenimientoVehiculoModelos.SeleccionarVehiculoParaNuevoRegistroRespuesta{}

	if solicitud == nil {
		respuesta.AsignarOk(false)
		respuesta.AsignarErr(errors.New(constantes.ERROR_PLACAS_INEXISTENTES_EN_BD))
		return respuesta
	}

	var err error
	if err = c.SeleccionarVehiculoParaNuevoRegistro(solicitud); err != nil {
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
	if c.ConectorBDControlador.ObtenerVehiculoPorPlacas(obtenerVehiculoPorPlacasSolicitud) == nil {
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

	_, err := time.Parse("02/01/2006", solicitud.ObtenerFecha()) // DD/MM/AAAA
	if err != nil {
		return errors.New(constantes.ERROR_FECHA_FORMATO_INVALIDO)
	}

	tipo := solicitud.ObtenerTipo()
	if tipo == "cargadecombustible" { // verificar como llega del json
		gasolina := solicitud.ObtenerLitrosDeGasolina()
		if gasolina <= 0 {
			return errors.New(constantes.ERROR_LITROS_GASOLINA_NO_ES_NUMERO)
		}
	}

	kilometraje := solicitud.ObtenerKilometraje()
	if kilometraje <= 0 {
		return errors.New(constantes.ERROR_KILOMETRAJE_NO_VALIDO)
	}

	importe := solicitud.ObtenerImporte()
	if importe <= 0 {
		return errors.New(constantes.ERROR_IMPORTE_NO_VALIDO)
	}

	return nil
}

func (c *Controlador) ObtenerServiciosVehicularesParaNuevoRegistro() []capturaserviciovehicular.ServicioVehicular {
	return c.ServicioVehicularControlador.ObtenerServiciosVehiculares()
}

/*
func (c *Controlador) EditarRegistroMantenimientoVehiculo(solicitud *visualizacionHistorialRegistrosControlador.EditarRegistroMantenimientoVehiculo) *registroMantenimientoVehiculoModelos.EditarRegistroMantenimientoVehiculoRespuesta {
	respuesta := &registroMantenimientoVehiculoModelos.EditarRegistroMantenimientoVehiculoRespuesta{}

	if solicitud == nil {
		respuesta.AsignarOk(false)
		respuesta.AsignarErr(errors.New(constantes.ERROR_SOLICITUD_NULA))
		return respuesta
	}

	if err := c.validarEditarRegistroMantenimientoVehiculoSolicitud(solicitud); err != nil {
		respuesta.AsignarOk(false)
		respuesta.AsignarErr(err)
		return respuesta
	}

	// Implementación para editar el registro de mantenimiento vehicular en la bd
	// Usar el mapeador para transformar la solicitud a la estructura correspondiente

	respuesta.AsignarOk(true)
	respuesta.AsignarErr(nil)

	return respuesta
}

func (c *Controlador) validarEditarRegistroMantenimientoVehiculoSolicitud(solicitud *visualizacionHistorialRegistrosControlador.EditarRegistroMantenimientoVehiculo) error {
	// Por hacer: Validaciones para asegurar la edición exitosa de registros de mantenimiento vehicular

	// ...

	return nil
}

func (c *Controlador) BorrarRegistroMantenimientoVehiculo(solicitud *visualizacionHistorialRegistrosControlador.BorrarRegistroMantenimientoVehiculo) *registroMantenimientoVehiculoModelos.BorrarRegistroMantenimientoVehiculoRespuesta {
	respuesta := &registroMantenimientoVehiculoModelos.BorrarRegistroMantenimientoVehiculoRespuesta{}

	if solicitud == nil {
		respuesta.AsignarOk(false)
		respuesta.AsignarErr(errors.New(constantes.ERROR_SOLICITUD_NULA))
		return respuesta
	}

	if err := c.validarBorrarRegistroMantenimientoVehiculoSolicitud(solicitud); err != nil {
		respuesta.AsignarOk(false)
		respuesta.AsignarErr(err)
		return respuesta
	}

	// Por hacer: Implementación para borrar el registro de mantenimiento vehicular en la base de datos
	// Usar el mapeador para transformar la solicitud a la estructura correspondiente

	respuesta.AsignarOk(true)
	respuesta.AsignarErr(nil)

	return respuesta
}

func (c *Controlador) validarBorrarRegistroMantenimientoVehiculoSolicitud(solicitud *visualizacionHistorialRegistrosControlador.BorrarRegistroMantenimientoVehiculo) error {

	// Por hacer: Validaciones específicas para el borrado exitoso de registros de mantenimiento vehicular

	return nil
}
*/
