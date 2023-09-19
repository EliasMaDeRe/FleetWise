package conectorbd

import "example/fleetwise/modelos/vehiculos"

type GuardarVehiculoSolicitud struct {
	vehiculo vehiculos.Vehiculo
}

func (v *GuardarVehiculoSolicitud) ObtenerVehiculo() (o vehiculos.Vehiculo) {
	if v != nil {
		o = v.vehiculo
	}
	return
}

func (v *GuardarVehiculoSolicitud) AsignarVehiculo(vehiculo vehiculos.Vehiculo) {
	if v != nil {
		v.vehiculo = vehiculo
	}
}

type ObtenerCargoSoicitud struct {
	claveUsuario  string
	nombreUsuario string
}

func (v *ObtenerCargoSoicitud) ObtenerClaveUsuario() (o string) {
	if v != nil {
		o = v.claveUsuario
	}
	return
}

func (v *ObtenerCargoSoicitud) AsignarClaveUsuario(claveUsuario string) {
	if v != nil {
		v.claveUsuario = claveUsuario
	}
}

func (v *ObtenerCargoSoicitud) ObtenerNombreUsuario() (o string) {
	if v != nil {
		o = v.nombreUsuario
	}
	return
}

func (v *ObtenerCargoSoicitud) AsignarNombreUsuario(nombreUsuario string) {
	if v != nil {
		v.nombreUsuario = nombreUsuario
	}
}

type ObtenerVehiculoPorPlacasSolicitud struct {
	placa string
}

func (v *ObtenerVehiculoPorPlacasSolicitud) ObtenerPlacas() (o string) {
	if v != nil {
		o = v.placa
	}
	return
}

func (v *ObtenerVehiculoPorPlacasSolicitud) AsignarPlacas(placa string) {
	if v != nil {
		v.placa = placa
	}
}
