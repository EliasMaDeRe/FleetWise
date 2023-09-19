package manejador

import (
	vehiculos "example/fleetwise/fuente/vehiculos/controlador"
	vehiculosModelos "example/fleetwise/modelos/vehiculos"
)

type Manejador struct {
	vehiculosControlador *vehiculos.Controlador
}

func NuevoManejador() (c *vehiculos.Controlador) {
	return &vehiculos.Controlador{}
}

func (m *Manejador) AgregarVehiculo(solicitud vehiculosModelos.AgregarVehiculosSolicitud) {
	return
}
