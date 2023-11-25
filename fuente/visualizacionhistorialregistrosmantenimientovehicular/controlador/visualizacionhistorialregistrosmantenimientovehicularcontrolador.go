package controlador

import (
	conectorBDControlador "example/fleetwise/fuente/conectorbd/controlador"
	historialRegistrosMapeador "example/fleetwise/fuente/visualizacionhistorialregistrosmantenimientovehicular/mapeador"
	capturaRegistrosModelo "example/fleetwise/modelos/capturaregistromantenimientovehiculo"
	capturaVehiculosModelo "example/fleetwise/modelos/capturavehiculos"
	historialRegistrosModelo "example/fleetwise/modelos/visualizacionhistorialregistrosmantenimientovehicular"
)

type Controlador struct{
	HistorialRegistrosMapeador *historialRegistrosMapeador.Mapeador
	conectorBDControlador *conectorBDControlador.Controlador 
}

func (c *Controlador) ObtenerRegistrosYVehiculosFiltrados(solicitud *historialRegistrosModelo.ObtenerRegistrosYVehiculosFiltradosSolicitud) ([]capturaRegistrosModelo.RegistroMantenimientoVehiculo, []capturaVehiculosModelo.Vehiculo){
	
	if (solicitud == nil || validarSolicitudSinFiltros(solicitud)){
		return c.conectorBDControlador.ObtenerRegistrosConVehiculos()
	}

	solicitudDb := c.HistorialRegistrosMapeador.ObtenerRegistrosYVehiculosFiltradosSolicitudAObtenerRegistrosConVehiculosFiltradosSolicitud(solicitud)


	return c.conectorBDControlador.ObtenerRegistrosConVehiculosFiltrados(solicitudDb)
} 

func validarSolicitudSinFiltros(solicitud *historialRegistrosModelo.ObtenerRegistrosYVehiculosFiltradosSolicitud) bool{
	return solicitud == &historialRegistrosModelo.ObtenerRegistrosYVehiculosFiltradosSolicitud{}
}
