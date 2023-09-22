package Manejador

import (
	conectorControlador "example/fleetwise/fuente/conectorbd/controlador"
	conectorMapeador "example/fleetwise/fuente/conectorbd/mapeador"
	conectorModelos "example/fleetwise/modelos/conectorbd"
	vehiculosModelos "example/fleetwise/modelos/vehiculos"
)

type Manejador struct {
	ConectorControlador *conectorControlador.Controlador
	ConectorMapeador    *conectorMapeador.Mapeador
}

func NuevoManejador() (c *Manejador) {
	return &Manejador{
		ConectorControlador: &conectorControlador.Controlador{},
	}
}
func (m *Manejador) ObtenerVehiculoPorPlacas(solicitud *conectorModelos.GuardarVehiculoSolicitud) {
	//m.ObtenerVehiculoPorPlacas(solicitud)

}

func (m *Manejador) ObtenerVehiculoPorSerie(solicitud *conectorModelos.ObtenerVehiculoPorPlacasSolicitud) {
	//m.ObtenerVehiculoPorSerie(solicitud)
}

func (m *Manejador) AgregarVehiculo(vehiculo *vehiculosModelos.Vehiculo) error {
	respuesta := m.ConectorControlador.AgregarVehiculo(vehiculo)
	return respuesta

}
