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
	ConectorBDControlador *conectorBDControlador.Controlador 
}

func (c *Controlador) ObtenerRegistrosFiltradosConVehiculos(solicitud *historialRegistrosModelo.ObtenerRegistrosFiltradosConVehiculosSolicitud) ([]capturaRegistrosModelo.RegistroMantenimientoVehiculo, []capturaVehiculosModelo.Vehiculo){

	if(solicitud == nil || validarSolicitudSinFiltros(solicitud)){
		return c.ConectorBDControlador.ObtenerRegistrosYVehiculosAsociados()
	}

	obtenerRegistrosConVehiculosFiltradosSolicitud := c.HistorialRegistrosMapeador.ObtenerRegistrosFiltradosConVehiculosSolicitudAObtenerRegistrosYVehiculosAsociadosFiltradosSolicitud(solicitud)

	return c.ConectorBDControlador.ObtenerRegistrosYVehiculosAsociadosFiltrados(obtenerRegistrosConVehiculosFiltradosSolicitud)
} 

func validarSolicitudSinFiltros(solicitud *historialRegistrosModelo.ObtenerRegistrosFiltradosConVehiculosSolicitud) bool{
	return solicitud == &historialRegistrosModelo.ObtenerRegistrosFiltradosConVehiculosSolicitud{}
}
