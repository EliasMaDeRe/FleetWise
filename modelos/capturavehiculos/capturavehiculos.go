package capturavehiculos

type Vehiculo struct {
	Anualidad int
	ID        string
	Marca     string
	Modelo    string
	Placas    string
	Serie     string
}

func (v *Vehiculo) ObtenerAnualidad() (o int) {
	if v != nil {
		o = v.Anualidad
	}
	return
}

func (v *Vehiculo) AsignarAnualidad(anualidad int) {
	if v != nil {
		v.Anualidad = anualidad
	}
}

func (v *Vehiculo) ObtenerID() (o string) {
	if v != nil {
		o = v.ID
	}
	return
}

func (v *Vehiculo) AsignarID(ID string) {
	if v != nil {
		v.ID = ID
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

type AgregarVehiculosSolicitud struct {
	Anualidad string `json:"anualidad"` // año del vehiculo
	Marca     string
	Modelo    string
	Placas    string
	Serie     string
}

func (v *AgregarVehiculosSolicitud) ObtenerAnualidad() (o string) {
	if v != nil {
		o = v.Anualidad
	}
	return
}

func (v *AgregarVehiculosSolicitud) AsignarAnualidad(anualidad string) {
	if v != nil {
		v.Anualidad = anualidad
	}
}

func (v *AgregarVehiculosSolicitud) ObtenerMarca() (o string) {
	if v != nil {
		o = v.Marca
	}
	return
}

func (v *AgregarVehiculosSolicitud) AsignarMarca(marca string) {
	if v != nil {
		v.Marca = marca
	}
}

func (v *AgregarVehiculosSolicitud) ObtenerModelo() (o string) {
	if v != nil {
		o = v.Modelo
	}
	return
}

func (v *AgregarVehiculosSolicitud) AsignarModelo(modelo string) {
	if v != nil {
		v.Modelo = modelo
	}
}

func (v *AgregarVehiculosSolicitud) ObtenerPlacas() (o string) {
	if v != nil {
		o = v.Placas
	}
	return
}

func (v *AgregarVehiculosSolicitud) AsignarPlacas(placas string) {
	if v != nil {
		v.Placas = placas
	}
}

func (v *AgregarVehiculosSolicitud) ObtenerSerie() (o string) {
	if v != nil {
		o = v.Serie
	}
	return
}

func (v *AgregarVehiculosSolicitud) AsignarSerie(serie string) {
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
