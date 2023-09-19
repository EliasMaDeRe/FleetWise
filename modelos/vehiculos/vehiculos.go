package vehiculos

type Vehiculo struct {
	anualidad int
	ID        string
	marca     string
	modelo    string
	placas    string
	serie     string
}

func (v *Vehiculo) ObtenerAnualidad() (o int) {
	if v != nil {
		o = v.anualidad
	}
	return
}

func (v *Vehiculo) AsignarAnualidad(anualidad int) {
	if v != nil {
		v.anualidad = anualidad
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
		o = v.marca
	}
	return
}

func (v *Vehiculo) AsignarMarca(marca string) {
	if v != nil {
		v.marca = marca
	}
}

func (v *Vehiculo) ObtenerModelo() (o string) {
	if v != nil {
		o = v.modelo
	}
	return
}

func (v *Vehiculo) AsignarModelo(modelo string) {
	if v != nil {
		v.modelo = modelo
	}
}

func (v *Vehiculo) ObtenerPlacas() (o string) {
	if v != nil {
		o = v.placas
	}
	return
}

func (v *Vehiculo) AsignarPlacas(placas string) {
	if v != nil {
		v.placas = placas
	}
}

func (v *Vehiculo) ObtenerSerie() (o string) {
	if v != nil {
		o = v.serie
	}
	return
}

func (v *Vehiculo) AsignarSerie(serie string) {
	if v != nil {
		v.serie = serie
	}
}

type AgregarVehiculosSolicitud struct {
	anualidad int
	marca     string
	modelo    string
	placas    string
	serie     string
}

func (v *AgregarVehiculosSolicitud) ObtenerAnualidad() (o int) {
	if v != nil {
		o = v.anualidad
	}
	return
}

func (v *AgregarVehiculosSolicitud) AsignarAnualidad(anualidad int) {
	if v != nil {
		v.anualidad = anualidad
	}
}

func (v *AgregarVehiculosSolicitud) ObtenerMarca() (o string) {
	if v != nil {
		o = v.marca
	}
	return
}

func (v *AgregarVehiculosSolicitud) AsignarMarca(marca string) {
	if v != nil {
		v.marca = marca
	}
}

func (v *AgregarVehiculosSolicitud) ObtenerModelo() (o string) {
	if v != nil {
		o = v.modelo
	}
	return
}

func (v *AgregarVehiculosSolicitud) AsignarModelo(modelo string) {
	if v != nil {
		v.modelo = modelo
	}
}

func (v *AgregarVehiculosSolicitud) ObtenerPlacas() (o string) {
	if v != nil {
		o = v.placas
	}
	return
}

func (v *AgregarVehiculosSolicitud) AsignarPlacas(placas string) {
	if v != nil {
		v.placas = placas
	}
}

func (v *AgregarVehiculosSolicitud) ObtenerSerie() (o string) {
	if v != nil {
		o = v.serie
	}
	return
}

func (v *AgregarVehiculosSolicitud) AsignarSerie(serie string) {
	if v != nil {
		v.serie = serie
	}
}
