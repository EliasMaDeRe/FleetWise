package controlador

import (
	vehiculosModelos "example/fleetwise/modelos/vehiculos"
)

type Controlador struct {
}

func (c *Controlador) AgregarVehiculo(solicitud *vehiculosModelos.AgregarVehiculosSolicitud) *vehiculosModelos.AgregarVehiculoRespuesta {

	respuesta := &vehiculosModelos.AgregarVehiculoRespuesta{}
	return respuesta
}
