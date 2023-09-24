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

type ObtenerVehiculoPorSerieSolicitud struct {
	serie string
}

func (v *ObtenerVehiculoPorSerieSolicitud) ObtenerSerie() (o string) {
	if v != nil {
		o = v.serie
	}
	return
}

func (v *ObtenerVehiculoPorSerieSolicitud) AsignarSerie(serie string) {
	if v != nil {
		v.serie = serie
	}
}

type AgregarConectorBDRespuesta struct {
	ok  bool
	err error
}

func (v *AgregarConectorBDRespuesta) ObtenerOk() (o bool) {
	if v != nil {
		o = v.ok
	}
	return
}

func (v *AgregarConectorBDRespuesta) AsignarOk(ok bool) {
	if v != nil {
		v.ok = ok
	}
}

func (v *AgregarConectorBDRespuesta) ObtenerErr() (o error) {
	if v != nil {
		o = v.err
	}
	return
}

func (v *AgregarConectorBDRespuesta) AsignarErr(err error) {
	if v != nil {
		v.err = err
	}
}
