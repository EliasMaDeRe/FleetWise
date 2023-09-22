package Mapeador

import (
	conectorModelos "example/fleetwise/modelos/conectorbd"
)

type Mapeador struct {
}

func (m *Mapeador) GinContextoGuardarVehiculoSolicitud() *conectorModelos.GuardarVehiculoSolicitud {
	nuevoVehiculo := &conectorModelos.GuardarVehiculoSolicitud{}
	return nuevoVehiculo
}
