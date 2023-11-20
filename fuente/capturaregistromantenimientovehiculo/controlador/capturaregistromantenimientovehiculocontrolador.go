package controlador

import (
	//"errors"
	conectorBDControlador "example/fleetwise/fuente/conectorbd/controlador"
	//"strconv"

	//"example/fleetwise/fuente/capturaregistromantenimientovehiculo/constantes"
	registroMantenimientoVehiculoMapeador "example/fleetwise/fuente/capturaregistromantenimientovehiculo/mapeador"
	registroMantenimientoVehiculoModelos "example/fleetwise/modelos/capturaregistromantenimientovehiculo"
	//visualizacionHistorialRegistrosControlador "example/fleetwise/modelos/visualizacionhistorialregistros"
	//"example/fleetwise/modelos/conectorbd"
)

type Controlador struct {
	ConectorBDControlador                 *conectorBDControlador.Controlador
	RegistroMantenimientoVehiculoMapeador *registroMantenimientoVehiculoMapeador.Mapeador
}

func (c *Controlador) SeleccionarVehiculo(solicitud *registroMantenimientoVehiculoModelos.SeleccionarVehiculoSolicitud) *registroMantenimientoVehiculoModelos.SeleccionarVehiculoRespuesta {
	//implementacipon futura por Pablo
	return nil
}

func (c *Controlador) validarSeleccionarVehiculo(solicitud *registroMantenimientoVehiculoModelos.SeleccionarVehiculoSolicitud) error {
	//implementacipon futura por Pablo
	return nil
}

func (c *Controlador) AgregarRegistroMantemientoVehiculo(solicitud *registroMantenimientoVehiculoModelos.AgregarRegistroMantenimientoVehiculoSolicitud) *registroMantenimientoVehiculoModelos.AgregarRegistroMantenimientoVehiculoRespuesta {
	//implementacipon futura por Pablo
	return nil
}

func (c *Controlador) validarAgregarRegistroMantemientoVehiculo(solicitud *registroMantenimientoVehiculoModelos.AgregarRegistroMantenimientoVehiculoSolicitud) error {
	//implementacipon futura por Pablo
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
