package controlador

import (
	"example/fleetwise/modelos/conectorbd"
	vehiculosModelos "example/fleetwise/modelos/vehiculos"
)

type Controlador struct {
}

func (c *Controlador) ObtenerVehiculoPorPlacas(solicitud *conectorbd.ObtenerVehiculoPorPlacasSolicitud) *vehiculosModelos.Vehiculo {
	return nil
}

func (c *Controlador) ObtenerVehiculoPorSerie(solicitud *conectorbd.ObtenerVehiculoPorSerieSolicitud) *vehiculosModelos.Vehiculo {
	return nil
}

func (c *Controlador) AgregarVehiculo(vehiculo *vehiculosModelos.Vehiculo) error {
	return nil
}
