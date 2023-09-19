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
