package capturavehiculos

type Vehiculo struct {
	FechaLanzamiento int
	Marca            string
	Modelo           string
	Placas           string
	Serie            string
}

func (Vehiculo) TableName() string {

	return "vehiculos"
}

func (v *Vehiculo) ObtenerFechaLanzamiento() (o int) {
	if v != nil {
		o = v.FechaLanzamiento
	}
	return
}

func (v *Vehiculo) AsignarFechaLanzamiento(fechaLanzamiento int) {
	if v != nil {
		v.FechaLanzamiento = fechaLanzamiento
	}
}

func (v *Vehiculo) ObtenerMarca() (o string) {
	if v != nil {
		o = v.Marca
	}
	return
}

func (v *Vehiculo) AsignarMarca(marca string) {
	if v != nil {
		v.Marca = marca
	}
}

func (v *Vehiculo) ObtenerModelo() (o string) {
	if v != nil {
		o = v.Modelo
	}
	return
}

func (v *Vehiculo) AsignarModelo(modelo string) {
	if v != nil {
		v.Modelo = modelo
	}
}

func (v *Vehiculo) ObtenerPlacas() (o string) {
	if v != nil {
		o = v.Placas
	}
	return
}

func (v *Vehiculo) AsignarPlacas(placas string) {
	if v != nil {
		v.Placas = placas
	}
}

func (v *Vehiculo) ObtenerSerie() (o string) {
	if v != nil {
		o = v.Serie
	}
	return
}

func (v *Vehiculo) AsignarSerie(serie string) {
	if v != nil {
		v.Serie = serie
	}
}

type AgregarVehiculoSolicitud struct {
	FechaLanzamiento string `json:"fechalanzamiento"` // año del vehiculo
	Marca            string `json:"marca"`
	Modelo           string `json:"modelo"`
	Placas           string `json:"placas"`
	Serie            string `json:"serie"`
}

func (v *AgregarVehiculoSolicitud) ObtenerFechaLanzamiento() (o string) {
	if v != nil {
		o = v.FechaLanzamiento
	}
	return
}

func (v *AgregarVehiculoSolicitud) AsignarFechaLanzamiento(fechaLanzamiento string) {
	if v != nil {
		v.FechaLanzamiento = fechaLanzamiento
	}
}

func (v *AgregarVehiculoSolicitud) ObtenerMarca() (o string) {
	if v != nil {
		o = v.Marca
	}
	return
}

func (v *AgregarVehiculoSolicitud) AsignarMarca(marca string) {
	if v != nil {
		v.Marca = marca
	}
}

func (v *AgregarVehiculoSolicitud) ObtenerModelo() (o string) {
	if v != nil {
		o = v.Modelo
	}
	return
}

func (v *AgregarVehiculoSolicitud) AsignarModelo(modelo string) {
	if v != nil {
		v.Modelo = modelo
	}
}

func (v *AgregarVehiculoSolicitud) ObtenerPlacas() (o string) {
	if v != nil {
		o = v.Placas
	}
	return
}

func (v *AgregarVehiculoSolicitud) AsignarPlacas(placas string) {
	if v != nil {
		v.Placas = placas
	}
}

func (v *AgregarVehiculoSolicitud) ObtenerSerie() (o string) {
	if v != nil {
		o = v.Serie
	}
	return
}

func (v *AgregarVehiculoSolicitud) AsignarSerie(serie string) {
	if v != nil {
		v.Serie = serie
	}
}

type AgregarVehiculoRespuesta struct {
	ok  bool
	err error
}

func (v *AgregarVehiculoRespuesta) ObtenerOk() (o bool) {
	if v != nil {
		o = v.ok
	}
	return
}

func (v *AgregarVehiculoRespuesta) AsignarOk(ok bool) {
	if v != nil {
		v.ok = ok
	}
}

func (v *AgregarVehiculoRespuesta) ObtenerErr() (o error) {
	if v != nil {
		o = v.err
	}
	return
}

func (v *AgregarVehiculoRespuesta) AsignarErr(err error) {
	if v != nil {
		v.err = err
	}
}

type EditarVehiculoSolicitud struct {
	PlacasActuales        string `json:"placasactuales"`
	PlacasNuevas          string `json:"placasnuevas"`
	FechaLanzamientoNueva string `json:"fechalanzamientonueva"`
	MarcaNueva            string `json:"marcanueva"`
	ModeloNuevo           string `json:"modelonuevo"`
	SerieNuevo            string `json:"serienuevo"`
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

func (v *EditarVehiculoSolicitud) ObtenerFechaDeLanzamientoNueva() (o string) {
	if v != nil {
		o = v.FechaLanzamientoNueva
	}
	return
}

func (v *EditarVehiculoSolicitud) AsignarFechaDeLanzamientoNueva(fechaDeLanzamientoNueva string) {
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

type EditarVehiculoRespuesta struct {
	ok  bool
	err error
}

func (v *EditarVehiculoRespuesta) ObtenerOk() (o bool) {
	if v != nil {
		o = v.ok
	}
	return
}

func (v *EditarVehiculoRespuesta) AsignarOk(ok bool) {
	if v != nil {
		v.ok = ok
	}
}

func (v *EditarVehiculoRespuesta) ObtenerErr() (o error) {
	if v != nil {
		o = v.err
	}
	return
}

func (v *EditarVehiculoRespuesta) AsignarErr(err error) {
	if v != nil {
		v.err = err
	}
}

type ObtenerVehiculoPorPlacasSolicitud struct {
	Placas string `json:"placas"`
}

func (v *ObtenerVehiculoPorPlacasSolicitud) ObtenerPlacas() (o string) {
	if v != nil {
		o = v.Placas
	}
	return
}

func (v *ObtenerVehiculoPorPlacasSolicitud) AsignarPlacas(placas string) {
	if v != nil {
		v.Placas = placas
	}
}

type EliminarVehiculoSolicitud struct {
	Placas string `json:"placas"`
}

func (v *EliminarVehiculoSolicitud) ObtenerPlacas() (o string) {
	if v != nil {
		o = v.Placas
	}
	return
}

func (v *EliminarVehiculoSolicitud) AsignarPlacas(placas string) {
	if v != nil {
		v.Placas = placas
	}
}

type EliminarVehiculoRespuesta struct {
	ok  bool
	err error
}

func (v *EliminarVehiculoRespuesta) ObtenerOk() (o bool) {
	if v != nil {
		o = v.ok
	}
	return
}

func (v *EliminarVehiculoRespuesta) AsignarOk(ok bool) {
	if v != nil {
		v.ok = ok
	}
}

func (v *EliminarVehiculoRespuesta) ObtenerErr() (o error) {
	if v != nil {
		o = v.err
	}
	return
}

func (v *EliminarVehiculoRespuesta) AsignarErr(err error) {
	if v != nil {
		v.err = err
	}
}
