package conectorbdcapturavehiculos

import (
	capturaVehiculosModelos "example/fleetwise/modelos/capturavehiculos"
)

type GuardarVehiculoSolicitud struct {
	vehiculo capturaVehiculosModelos.Vehiculo
}

func (v *GuardarVehiculoSolicitud) ObtenerVehiculo() (o capturaVehiculosModelos.Vehiculo) {
	if v != nil {
		o = v.vehiculo
	}
	return
}

func (v *GuardarVehiculoSolicitud) AsignarVehiculo(vehiculo capturaVehiculosModelos.Vehiculo) {
	if v != nil {
		v.vehiculo = vehiculo
	}
}

type GuardarVehiculoRespuesta struct {
	ok  bool
	err error
}

func (v *GuardarVehiculoRespuesta) ObtenerOk() (o bool) {
	if v != nil {
		o = v.ok
	}
	return
}

func (v *GuardarVehiculoRespuesta) AsignarOk(ok bool) {
	if v != nil {
		v.ok = ok
	}
}

func (v *GuardarVehiculoRespuesta) ObtenerErr() (o error) {
	if v != nil {
		o = v.err
	}
	return
}

func (v *GuardarVehiculoRespuesta) AsignarErr(err error) {
	if v != nil {
		v.err = err
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

type EditarVehiculoSolicitud struct {
	PlacasActuales        string
	PlacasNuevas          string
	FechaLanzamientoNueva int
	MarcaNueva            string
	ModeloNuevo           string
	SerieNuevo            string
}

func (v *EditarVehiculoSolicitud) ObtenerPlacasActuales() (o string) {
	if v != nil {
		o = v.PlacasActuales
	}
	return
}

func (v *EditarVehiculoSolicitud) AsignarPlacasActuales(placasActuales string) {
	if v != nil {
		v.PlacasActuales = placasActuales
	}
}

func (v *EditarVehiculoSolicitud) ObtenerPlacasNuevas() (o string) {
	if v != nil {
		o = v.PlacasNuevas
	}
	return
}

func (v *EditarVehiculoSolicitud) AsignarPlacasNuevas(placasNuevas string) {
	if v != nil {
		v.PlacasNuevas = placasNuevas
	}
}

func (v *EditarVehiculoSolicitud) ObtenerFechaDeLanzamientoNueva() (o int) {
	if v != nil {
		o = v.FechaLanzamientoNueva
	}
	return
}

func (v *EditarVehiculoSolicitud) AsignarFechaDeLanzamientoNueva(fechaDeLanzamientoNueva int) {
	if v != nil {
		v.FechaLanzamientoNueva = fechaDeLanzamientoNueva
	}
}

func (v *EditarVehiculoSolicitud) ObtenerMarcaNueva() (o string) {
	if v != nil {
		o = v.MarcaNueva
	}
	return
}

func (v *EditarVehiculoSolicitud) AsignarMarcaNueva(marcaNueva string) {
	if v != nil {
		v.MarcaNueva = marcaNueva
	}
}

func (v *EditarVehiculoSolicitud) ObtenerModeloNuevo() (o string) {
	if v != nil {
		o = v.ModeloNuevo
	}
	return
}

func (v *EditarVehiculoSolicitud) AsignarModeloNuevo(modeloNuevo string) {
	if v != nil {
		v.ModeloNuevo = modeloNuevo
	}
}

func (v *EditarVehiculoSolicitud) ObtenerSerieNuevo() (o string) {
	if v != nil {
		o = v.SerieNuevo
	}
	return
}

func (v *EditarVehiculoSolicitud) AsignarSerieNuevo(serieNuevo string) {
	if v != nil {
		v.SerieNuevo = serieNuevo
	}
}

