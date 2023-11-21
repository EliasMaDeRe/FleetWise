package capturaserviciovehicular

type ServicioVehicular struct {
	Nombre string
}

func (v *ServicioVehicular) ObtenerNombre() (o string) {
	if v != nil {
		o = v.Nombre
	}
	return
}

func (v *ServicioVehicular) AsignarNombre(nombre string) {
	if v != nil {
		v.Nombre = nombre
	}
}

func (ServicioVehicular) TableName() string {
	return "servicios_vehiculares"
}

type AgregarServicioVehicularSolicitud struct {
	Nombre string `json:"nombre"` // nombre del servicio vehicular
}

func (v *AgregarServicioVehicularSolicitud) ObtenerNombre() (o string) {
	if v != nil {
		o = v.Nombre
	}
	return
}

func (v *AgregarServicioVehicularSolicitud) AsignarNombre(nombre string) {
	if v != nil {
		v.Nombre = nombre
	}
}

type EditarServicioVehicularSolicitud struct {
	NombreActual string `json:"nombreactual"`
	NuevoNombre  string `json:"nuevonombre"`
}

func (v *EditarServicioVehicularSolicitud) ObtenerNombreActual() (o string) {
	if v != nil {
		o = v.NombreActual
	}
	return
}

func (v *EditarServicioVehicularSolicitud) AsignarNombreActual(nombreActual string) {
	if v != nil {
		v.NombreActual = nombreActual
	}
}

func (v *EditarServicioVehicularSolicitud) ObtenerNuevoNombre() (o string) {
	if v != nil {
		o = v.NuevoNombre
	}
	return
}

func (v *EditarServicioVehicularSolicitud) AsignarNuevoNombre(nuevoNombre string) {
	if v != nil {
		v.NuevoNombre = nuevoNombre
	}
}

type EliminarServicioVehicularSolicitud struct {
	Nombre string `json:"nombre"` // nombre del servicio vehicular
}

func (v *EliminarServicioVehicularSolicitud) ObtenerNombre() (o string) {
	if v != nil {
		o = v.Nombre
	}
	return
}

func (v *EliminarServicioVehicularSolicitud) AsignarNombre(nombre string) {
	if v != nil {
		v.Nombre = nombre
	}
}

type AgregarServicioVehicularRespuesta struct {
	ok  bool
	err error
}

func (v *AgregarServicioVehicularRespuesta) ObtenerOk() (o bool) {
	if v != nil {
		o = v.ok
	}
	return
}

func (v *AgregarServicioVehicularRespuesta) AsignarOk(ok bool) {
	if v != nil {
		v.ok = ok
	}
}

func (v *AgregarServicioVehicularRespuesta) ObtenerErr() (o error) {
	if v != nil {
		o = v.err
	}
	return
}

func (v *AgregarServicioVehicularRespuesta) AsignarErr(err error) {
	if v != nil {
		v.err = err
	}
}

type EliminarServicioVehicularRespuesta struct {
	ok  bool
	err error
}

func (v *EliminarServicioVehicularRespuesta) ObtenerOk() (o bool) {
	if v != nil {
		o = v.ok
	}
	return
}

func (v *EliminarServicioVehicularRespuesta) AsignarOk(ok bool) {
	if v != nil {
		v.ok = ok
	}
}

func (v *EliminarServicioVehicularRespuesta) ObtenerErr() (o error) {
	if v != nil {
		o = v.err
	}
	return
}

func (v *EliminarServicioVehicularRespuesta) AsignarErr(err error) {
	if v != nil {
		v.err = err
	}
}

type EditarServicioVehicularRespuesta struct {
	ok  bool
	err error
}

func (v *EditarServicioVehicularRespuesta) ObtenerOk() (o bool) {
	if v != nil {
		o = v.ok
	}
	return
}

func (v *EditarServicioVehicularRespuesta) AsignarOk(ok bool) {
	if v != nil {
		v.ok = ok
	}
}

func (v *EditarServicioVehicularRespuesta) ObtenerErr() (o error) {
	if v != nil {
		o = v.err
	}
	return
}

func (v *EditarServicioVehicularRespuesta) AsignarErr(err error) {
	if v != nil {
		v.err = err
	}
}
