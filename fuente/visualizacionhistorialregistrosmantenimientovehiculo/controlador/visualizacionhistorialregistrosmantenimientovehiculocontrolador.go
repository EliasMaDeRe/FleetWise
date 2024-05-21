package controlador

import (
	conectorBDControlador "example/fleetwise/fuente/conectorbdvisualizacionhistorialregistrosmantenimientovehiculo/controlador"
	visualizacionHistorialRegistrosMapeador "example/fleetwise/fuente/visualizacionhistorialregistrosmantenimientovehiculo/mapeador"
	capturaRegistrosMantenimientoVehiculoModelos "example/fleetwise/modelos/capturaregistromantenimientovehiculo"
	capturaVehiculosModelos "example/fleetwise/modelos/capturavehiculos"
	visualizacionHistorialRegistrosMantenimientoVehiculoModelos "example/fleetwise/modelos/visualizacionhistorialregistrosmantenimientovehiculo"
)

type Controlador struct {
	HistorialRegistrosMapeador *visualizacionHistorialRegistrosMapeador.Mapeador
	ConectorBDControlador      *conectorBDControlador.Controlador
}

func (c *Controlador) ObtenerRegistrosFiltradosConVehiculos(solicitud *visualizacionHistorialRegistrosMantenimientoVehiculoModelos.ObtenerRegistrosFiltradosConVehiculosSolicitud) ([]capturaRegistrosMantenimientoVehiculoModelos.RegistroDeMantenimientoDeVehiculo, []capturaVehiculosModelos.Vehiculo) {

	if solicitud == nil || validarSolicitudSinFiltros(solicitud) {
		return c.ConectorBDControlador.ObtenerRegistrosYVehiculosAsociados()
	}
	
	obtenerRegistrosConVehiculosFiltradosSolicitud := c.HistorialRegistrosMapeador.ObtenerRegistrosFiltradosConVehiculosSolicitudAObtenerRegistrosYVehiculosAsociadosFiltradosSolicitud(solicitud)

	return c.ConectorBDControlador.ObtenerRegistrosYVehiculosAsociadosFiltrados(obtenerRegistrosConVehiculosFiltradosSolicitud)
}

func validarSolicitudSinFiltros(solicitud *visualizacionHistorialRegistrosMantenimientoVehiculoModelos.ObtenerRegistrosFiltradosConVehiculosSolicitud) bool {
	if(solicitud.ObtenerFechaDeLanzamiento() != ""){ return false}
	if(solicitud.ObtenerMarca() != ""){ return false}
	if(solicitud.ObtenerModelo() != ""){ return false}
	if(solicitud.ObtenerPlacas() != ""){ return false}
	if(solicitud.ObtenerTipoDeRegistro() != ""){ return false}
	return true
}
