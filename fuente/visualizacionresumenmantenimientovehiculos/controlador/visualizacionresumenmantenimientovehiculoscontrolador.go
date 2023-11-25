package controlador

import (
	conectorBDControlador "example/fleetwise/fuente/conectorbd/controlador"
)

type Controlador struct {
	ConectorBDControlador     *conectorBDControlador.Controlador
}

func (c *Controlador) CalcularMetricasVehiculares(){

}

func (c *Controlador) ObtenerResumenMantenimientoVehiculos(){
	
}