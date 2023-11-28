package controlador

import (
	"errors"
	capturaServicioVehicularControlador "example/fleetwise/fuente/capturaserviciovehicular/controlador"
	conectorBDControlador "example/fleetwise/fuente/conectorbd/controlador"
	"strconv"
	"time"

	"example/fleetwise/fuente/capturaregistromantenimientovehiculo/constantes"
	capturaRegistroMantenimientoVehiculoMapeador "example/fleetwise/fuente/capturaregistromantenimientovehiculo/mapeador"
	capturaRegistroMantenimientoVehiculoModelos "example/fleetwise/modelos/capturaregistromantenimientovehiculo"
	capturaServicioVehicularModelos "example/fleetwise/modelos/capturaserviciovehicular"
	capturaVehiculoModelos "example/fleetwise/modelos/capturavehiculos"
	conectorBDModelos "example/fleetwise/modelos/conectorbd"
)

type Controlador struct {
	ConectorBDControlador                        *conectorBDControlador.Controlador
	CapturaRegistroMantenimientoVehiculoMapeador *capturaRegistroMantenimientoVehiculoMapeador.Mapeador
	CapturaServicioVehicularControlador          *capturaServicioVehicularControlador.Controlador
}

func (c *Controlador) SeleccionarVehiculoParaNuevoRegistro(solicitud *capturaRegistroMantenimientoVehiculoModelos.SeleccionarVehiculoParaNuevoRegistroSolicitud) *capturaRegistroMantenimientoVehiculoModelos.SeleccionarVehiculoParaNuevoRegistroRespuesta {
	respuesta := &capturaRegistroMantenimientoVehiculoModelos.SeleccionarVehiculoParaNuevoRegistroRespuesta{}

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

func (c *Controlador) validarSeleccionarVehiculoParaRegistro(solicitud *capturaRegistroMantenimientoVehiculoModelos.SeleccionarVehiculoParaNuevoRegistroSolicitud) error {
	obtenerVehiculoPorPlacasSolicitud := &conectorBDModelos.ObtenerVehiculoPorPlacasSolicitud{}
	obtenerVehiculoPorPlacasSolicitud.AsignarPlacas(solicitud.ObtenerPlacas())
	if c.CapturaServicioVehicularControlador.ConectorBDControlador.ObtenerVehiculoPorPlacas(obtenerVehiculoPorPlacasSolicitud) == nil {
		return errors.New(constantes.ERROR_PLACAS_INEXISTENTES_EN_BD)
	}
	return nil
}

func (c *Controlador) AgregarRegistroMantemientoVehiculo(solicitud *capturaRegistroMantenimientoVehiculoModelos.AgregarRegistroMantenimientoVehiculoSolicitud) *capturaRegistroMantenimientoVehiculoModelos.AgregarRegistroMantenimientoVehiculoRespuesta {
	respuesta := &capturaRegistroMantenimientoVehiculoModelos.AgregarRegistroMantenimientoVehiculoRespuesta{}

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

	registro := c.CapturaRegistroMantenimientoVehiculoMapeador.AgregarRegistroMantemientoVehiculoSolicitudARegistroMantemientoVehiculo(solicitud)

	if guardarRegistroMantenimientoVehiculoRespuesta := c.CapturaServicioVehicularControlador.ConectorBDControlador.GuardarRegistro(registro); guardarRegistroMantenimientoVehiculoRespuesta.ObtenerErr() != nil {
		respuesta.AsignarOk(false)
		respuesta.AsignarErr(guardarRegistroMantenimientoVehiculoRespuesta.ObtenerErr())
		return respuesta
	}

	respuesta.AsignarOk(true)
	respuesta.AsignarErr(nil)

	return respuesta
}

func (c *Controlador) validarAgregarRegistroMantemientoVehiculo(solicitud *capturaRegistroMantenimientoVehiculoModelos.AgregarRegistroMantenimientoVehiculoSolicitud) error {

	_, err := time.Parse("2006-01-02", solicitud.ObtenerFecha()) // AAAA/MM/DD
	if err != nil {
		return errors.New(constantes.ERROR_FECHA_FORMATO_INVALIDO)
	}

	tipoDeRegistro := solicitud.ObtenerTipo()

	if tipoDeRegistro == "carga_combustible" {
		if litrosDeGasolina, err := strconv.ParseFloat(solicitud.ObtenerLitrosDeGasolina(), 64); err != nil || litrosDeGasolina <= 0 {
			return errors.New(constantes.ERROR_LITROS_GASOLINA_NO_ES_NUMERO)
		}
	}

	if kilometraje, err := strconv.Atoi(solicitud.ObtenerKilometraje()); err != nil || kilometraje <= 0 {
		return errors.New(constantes.ERROR_KILOMETRAJE_NO_VALIDO)
	}

	if importe, err := strconv.ParseFloat(solicitud.ObtenerImporte(), 64); err != nil || importe <= 0 {
		return errors.New(constantes.ERROR_IMPORTE_NO_VALIDO)
	}

	return nil
}

func (c *Controlador) ObtenerServiciosVehicularesParaNuevoRegistro() []capturaServicioVehicularModelos.ServicioVehicular {
	return c.CapturaServicioVehicularControlador.ObtenerServiciosVehiculares()
}

func (c *Controlador) ObtenerRegistroMantenimientoVehiculo(solicitud *capturaRegistroMantenimientoVehiculoModelos.ObtenerRegistroDeMantenimientoDelVehiculoPorNumeroDeRegistroSolicitud) (*capturaRegistroMantenimientoVehiculoModelos.RegistroMantenimientoVehiculo, *capturaVehiculoModelos.Vehiculo) {
	if solicitud == nil || c.esSolicitudVacia(solicitud) {
		return &capturaRegistroMantenimientoVehiculoModelos.RegistroMantenimientoVehiculo{}, &capturaVehiculoModelos.Vehiculo{}
	}

	ObtenerRegistroYVehiculoAsociadoPorNumeroDeRegistroSolicitud := c.CapturaRegistroMantenimientoVehiculoMapeador.ObtenerRegistroMantenimientoVehiculoSolicitudAObtenerRegistroYVehiculoAsociadoPorNumeroDeRegistroSolicitud(solicitud)

	return c.ConectorBDControlador.ObtenerRegistroYVehiculoAsociadoPorNumeroDeRegistro(ObtenerRegistroYVehiculoAsociadoPorNumeroDeRegistroSolicitud)
}

func (c *Controlador) esSolicitudVacia(solicitud *capturaRegistroMantenimientoVehiculoModelos.ObtenerRegistroDeMantenimientoDelVehiculoPorNumeroDeRegistroSolicitud) bool {
	return solicitud == &capturaRegistroMantenimientoVehiculoModelos.ObtenerRegistroDeMantenimientoDelVehiculoPorNumeroDeRegistroSolicitud{}
}

func (c *Controlador) EditarRegistroDeMantenimientoDelVehiculo(solicitud *capturaRegistroMantenimientoVehiculoModelos.EditarRegistroDeMantenimientoDelVehiculoSolicitud) *capturaRegistroMantenimientoVehiculoModelos.EditarRegistroMantenimientoVehiculoRespuesta {

	if c.solicitudEditarRegistroVacia(solicitud) {
		respuesta := &capturaRegistroMantenimientoVehiculoModelos.EditarRegistroMantenimientoVehiculoRespuesta{}
		respuesta.AsignarOk(false)
		respuesta.AsignarErr(errors.New(constantes.ERROR_SOLICITUD_NULA))
		return respuesta
	}

	if !c.existePlaca(solicitud.ObtenerPlacasVehiculo()) {
		respuesta := &capturaRegistroMantenimientoVehiculoModelos.EditarRegistroMantenimientoVehiculoRespuesta{}
		respuesta.AsignarOk(false)
		respuesta.AsignarErr(errors.New(constantes.ERROR_PLACAS_INEXISTENTES_EN_BD))
		return respuesta
	}

	if !c.numeroDeRegistroExistente(solicitud.ObtenerNumeroDeRegistro()) {
		respuesta := &capturaRegistroMantenimientoVehiculoModelos.EditarRegistroMantenimientoVehiculoRespuesta{}
		respuesta.AsignarOk(false)
		respuesta.AsignarErr(errors.New(constantes.ERROR_NUMERO_REGISTRO_NO_EXISTENTE))
		return respuesta
	}

	mensajeError, solicitudValida := c.validarSolicitudEditarRegistro(solicitud)

	if !solicitudValida {
		respuesta := &capturaRegistroMantenimientoVehiculoModelos.EditarRegistroMantenimientoVehiculoRespuesta{}
		respuesta.AsignarOk(false)
		respuesta.AsignarErr(errors.New(mensajeError))
		return respuesta
	}

	actualizarRegistroSolicitud := c.CapturaRegistroMantenimientoVehiculoMapeador.EditarRegistroDeMantenimientoDelVehiculoSolicitudAActualizarRegistroDeMantenimientoDelVehiculoSolicitud(solicitud)

	actualizarRegistroRespuesta := c.ConectorBDControlador.ActualizarRegistroMantenimientoVehiculo(actualizarRegistroSolicitud)

	editarRegistroRespuesta := c.CapturaRegistroMantenimientoVehiculoMapeador.ActualizarRegistroMantenimientoVehiculoRespuestaAEditarRegistroMantenimientoVehiculoRespuesta(actualizarRegistroRespuesta)

	return editarRegistroRespuesta
}

func (c *Controlador) solicitudEditarRegistroVacia(solicitud *capturaRegistroMantenimientoVehiculoModelos.EditarRegistroDeMantenimientoDelVehiculoSolicitud) bool {
	return solicitud == nil || solicitud == &capturaRegistroMantenimientoVehiculoModelos.EditarRegistroDeMantenimientoDelVehiculoSolicitud{}
}

func (c *Controlador) existePlaca(placa string) bool {
	solicitudObtenerVehiculoPorPlaca := &conectorBDModelos.ObtenerVehiculoPorPlacasSolicitud{}
	solicitudObtenerVehiculoPorPlaca.AsignarPlacas(placa)

	vehiculo := c.ConectorBDControlador.ObtenerVehiculoPorPlacas(solicitudObtenerVehiculoPorPlaca)

	return vehiculo != nil && vehiculo != &capturaVehiculoModelos.Vehiculo{}
}

func (c *Controlador) numeroDeRegistroExistente(numeroDeRegistro int) bool {
	solicitud := &capturaRegistroMantenimientoVehiculoModelos.ObtenerRegistroYVehiculoAsociadoPorNumeroDeRegistroSolicitud{
		NumDeRegistro: numeroDeRegistro,
	}

	registro, vehiculo := c.ConectorBDControlador.ObtenerRegistroYVehiculoAsociadoPorNumeroDeRegistro(solicitud)

	return registro != &capturaRegistroMantenimientoVehiculoModelos.RegistroMantenimientoVehiculo{} && vehiculo != &capturaVehiculoModelos.Vehiculo{}
}

func (c *Controlador) validarSolicitudEditarRegistro(solicitud *capturaRegistroMantenimientoVehiculoModelos.EditarRegistroDeMantenimientoDelVehiculoSolicitud) (string, bool) {

	if solicitud.ObtenerFecha() == "" {
		return constantes.ERROR_FECHA_VACIA, false
	}

	if solicitud.ObtenerTipo() == "" {
		return constantes.ERROR_TIPO_REGISTRO_VACIO, false
	}

	if solicitud.ObtenerKilometraje() == 0 {
		return constantes.ERROR_KILOMETRAJE_NO_VALIDO, false
	}

	if solicitud.ObtenerImporte() == 0 {
		return constantes.ERROR_IMPORTE_NO_VALIDO, false
	}

	if solicitud.ObtenerObservaciones() == "" {
		return constantes.ERROR_OBSERVACIONES_VACIO, false
	}

	if solicitud.ObtenerTipo() == "carga de combustible" && solicitud.ObtenerLitrosDeGasolina() == 0 {

		return constantes.ERROR_LIBROS_GASOLINA_VACIO, false
	}

	if solicitud.ObtenerTipo() != "carga de combustible" && solicitud.ObtenerConcepto() == "" {
		return constantes.ERROR_CONCEPTO_VACIO, false
	}

	return "", true
}
