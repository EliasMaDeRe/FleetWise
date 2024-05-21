package conectorbdcapturaserviciovehicular

type ObtenerServicioVehicularPorNombreSolicitud struct {
	nombre string
}

func (v *ObtenerServicioVehicularPorNombreSolicitud) ObtenerNombre() (o string) {
	if v != nil {
		o = v.nombre
	}
	return
}

func (v *ObtenerServicioVehicularPorNombreSolicitud) AsignarNombre(nombre string) {
	if v != nil {
		v.nombre = nombre
	}
}


type EditarServicioVehicularSolicitud struct {
	nombreActual string
	nuevoNombre  string
}

func (v *EditarServicioVehicularSolicitud) ObtenerNombreActual() (o string) {
	if v != nil {
		o = v.nombreActual
	}
	return
}

func (v *EditarServicioVehicularSolicitud) AsignarNombreActual(nombreActual string) {
	if v != nil {
		v.nombreActual = nombreActual
	}
}

func (v *EditarServicioVehicularSolicitud) ObtenerNuevoNombre() (o string) {
	if v != nil {
		o = v.nuevoNombre
	}
	return
}

func (v *EditarServicioVehicularSolicitud) AsignarNuevoNombre(nuevoNombre string) {
	if v != nil {
		v.nuevoNombre = nuevoNombre
	}
}

type GuardarServicioVehicularRespuesta struct {
	ok  bool
	err error
}

func (v *GuardarServicioVehicularRespuesta) ObtenerOk() (o bool) {
	if v != nil {
		o = v.ok
	}
	return
}

func (v *GuardarServicioVehicularRespuesta) AsignarOk(ok bool) {
	if v != nil {
		v.ok = ok
	}
}

func (v *GuardarServicioVehicularRespuesta) ObtenerErr() (o error) {
	if v != nil {
		o = v.err
	}
	return
}

func (v *GuardarServicioVehicularRespuesta) AsignarErr(err error) {
	if v != nil {
		v.err = err
	}
}