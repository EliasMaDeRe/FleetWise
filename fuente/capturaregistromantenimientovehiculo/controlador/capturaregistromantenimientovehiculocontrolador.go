package controlador

import (
	"errors"
	servicioVehicularControlador "example/fleetwise/fuente/capturaserviciovehicular/controlador"
	conectorBDControlador "example/fleetwise/fuente/conectorbd/controlador"
	"fmt"
	"time"

	"example/fleetwise/fuente/capturaregistromantenimientovehiculo/constantes"
	registroMantenimientoVehiculoMapeador "example/fleetwise/fuente/capturaregistromantenimientovehiculo/mapeador"
	registroMantenimientoVehiculoModelos "example/fleetwise/modelos/capturaregistromantenimientovehiculo"
	"example/fleetwise/modelos/capturaserviciovehicular"
	"example/fleetwise/modelos/capturavehiculos"
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

	if err := c.validarAgregarRegistroMantemientoVehiculo(solicitud); err != nil {
		respuesta.AsignarOk(false)
		respuesta.AsignarErr(err)
		return respuesta
	}

	registro := c.RegistroMantenimientoVehiculoMapeador.AgregarRegistroMantemientoVehiculoSolicitudARegistroMantemientoVehiculo(solicitud)

	if guardarRegistroMantenimientoVehiculoRespuesta := c.ConectorBDControlador.GuardarRegistro(registro); guardarRegistroMantenimientoVehiculoRespuesta.ObtenerErr() != nil {
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

	if tipo == "Carga de Combustible" {
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

func (c *Controlador) ObtenerRegistroMantenimientoVehicular(solicitud *registroMantenimientoVehiculoModelos.ObtenerRegistroMantenimientoVehicularPorNumeroDeRegistroSolicitud) (*registroMantenimientoVehiculoModelos.RegistroMantenimientoVehiculo, *capturavehiculos.Vehiculo){
	if(solicitud == nil || c.solicitudVacia(solicitud)){
		return &registroMantenimientoVehiculoModelos.RegistroMantenimientoVehiculo{}, &capturavehiculos.Vehiculo{}
	}
	
	ObtenerRegistroYVehiculoAsociadoPorNumeroDeRegistroSolicitud := c.RegistroMantenimientoVehiculoMapeador.ObtenerRegistroMantenimientoVehiculoSolicitudAObtenerRegistroYVehiculoAsociadoPorNumeroDeRegistroSolicitud(solicitud)

	return c.ConectorBDControlador.ObtenerRegistroYVehiculoAsociadoPorNumeroDeRegistro(ObtenerRegistroYVehiculoAsociadoPorNumeroDeRegistroSolicitud)
}

func (c *Controlador) solicitudVacia(solicitud *registroMantenimientoVehiculoModelos.ObtenerRegistroMantenimientoVehicularPorNumeroDeRegistroSolicitud) bool {
	return solicitud == &registroMantenimientoVehiculoModelos.ObtenerRegistroMantenimientoVehicularPorNumeroDeRegistroSolicitud{}
}

func (c *Controlador) EditarRegistroMantenimientoVehicular(solicitud *registroMantenimientoVehiculoModelos.EditarRegistroMantenimientoVehicularSolicitud) *registroMantenimientoVehiculoModelos.EditarRegistroMantenimientoVehicularRespuesta{

	if(c.solicitudEditarRegistroVacia(solicitud)){
		respuesta := &registroMantenimientoVehiculoModelos.EditarRegistroMantenimientoVehicularRespuesta{}
		respuesta.AsignarOk(false)
		respuesta.AsignarError(errors.New(constantes.ERROR_SOLICITUD_NULA))
		return respuesta
	}
	
	if(! c.placaExistente(solicitud.ObtenerPlacasVehiculo())){
		respuesta := &registroMantenimientoVehiculoModelos.EditarRegistroMantenimientoVehicularRespuesta{}
		respuesta.AsignarOk(false)
		respuesta.AsignarError(errors.New(constantes.ERROR_PLACAS_INEXISTENTES_EN_BD))
		return respuesta
	}

	if(! c.numeroDeRegistroExistente(solicitud.ObtenerNumeroDeRegistro())){
		respuesta := &registroMantenimientoVehiculoModelos.EditarRegistroMantenimientoVehicularRespuesta{}
		respuesta.AsignarOk(false)
		respuesta.AsignarError(errors.New(constantes.ERROR_NUMERO_REGISTRO_NO_EXISTENTE))
		return respuesta
	}

	mensajeError, solicitudValida := c.validarSolicitudEditarRegistro(solicitud);
	
	if(!solicitudValida){
		respuesta := &registroMantenimientoVehiculoModelos.EditarRegistroMantenimientoVehicularRespuesta{}
		respuesta.AsignarOk(false)
		respuesta.AsignarError(errors.New(mensajeError))
		return respuesta;
	}

	actualizarRegistroSolicitud := c.RegistroMantenimientoVehiculoMapeador.EditarRegistroMantenimientoVehicularSolicitudAActualizarRegistroMantenimientoVehicularSolicitud(solicitud)

	actualizarRegistroRespuesta := c.ConectorBDControlador.ActualizarRegistroMantenimientoVehicular(actualizarRegistroSolicitud)
	
	fmt.Println(actualizarRegistroRespuesta)
	
	editarRegistroRespuesta := c.RegistroMantenimientoVehiculoMapeador.ActualizarRegistroMantenimientoVehicularRespuestaAEditarRegistroMantenimientoVehicularRespuesta(actualizarRegistroRespuesta)

	return editarRegistroRespuesta
}

func (c *Controlador) solicitudEditarRegistroVacia(solicitud *registroMantenimientoVehiculoModelos.EditarRegistroMantenimientoVehicularSolicitud) bool{
	return solicitud == nil || solicitud == &registroMantenimientoVehiculoModelos.EditarRegistroMantenimientoVehicularSolicitud{}
}

func (c *Controlador) placaExistente(placa string) bool{
	solicitudObtenerVehiculoPorPlaca := &conectorbd.ObtenerVehiculoPorPlacasSolicitud{}
	solicitudObtenerVehiculoPorPlaca.AsignarPlacas(placa)

	vehiculo := c.ConectorBDControlador.ObtenerVehiculoPorPlacas(solicitudObtenerVehiculoPorPlaca)

	return vehiculo != nil && vehiculo != &capturavehiculos.Vehiculo{}
}

func (c *Controlador) numeroDeRegistroExistente(numeroDeRegistro int) bool{
	solicitud := &registroMantenimientoVehiculoModelos.ObtenerRegistroYVehiculoAsociadoPorNumeroDeRegistroSolicitud{
		NumDeRegistro: numeroDeRegistro,
	}
	
	registro, vehiculo := c.ConectorBDControlador.ObtenerRegistroYVehiculoAsociadoPorNumeroDeRegistro(solicitud)
	
	return registro != &registroMantenimientoVehiculoModelos.RegistroMantenimientoVehiculo{} && vehiculo != &capturavehiculos.Vehiculo{}
}

func (c *Controlador) validarSolicitudEditarRegistro(solicitud *registroMantenimientoVehiculoModelos.EditarRegistroMantenimientoVehicularSolicitud) (string, bool){
	
	if(solicitud.ObtenerFecha() == "" ){
		return constantes.ERROR_FECHA_VACIA, false
	}

	if(solicitud.ObtenerTipo() == ""){
		return constantes.ERROR_TIPO_REGISTRO_VACIO, false
	}

	if(solicitud.ObtenerKilometraje() == 0){
		return constantes.ERROR_KILOMETRAJE_NO_VALIDO, false
	}

	if(solicitud.ObtenerImporte() == 0){
		return constantes.ERROR_IMPORTE_NO_VALIDO, false
	}

	if(solicitud.ObtenerObservaciones() == ""){
		return constantes.ERROR_OBSERVACIONES_VACIO, false
	}

	if(solicitud.ObtenerTipo() == "carga de combustible" && solicitud.ObtenerLitrosDeGasolina() == 0){
		
		return constantes.ERROR_LIBROS_GASOLINA_VACIO, false
	}

	if(solicitud.ObtenerTipo() != "carga de combustible" && solicitud.ObtenerConcepto() == ""){
		return constantes.ERROR_CONCEPTO_VACIO, false
	}

	return "", true
}