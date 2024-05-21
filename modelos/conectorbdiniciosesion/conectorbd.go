package conectorbd

type ObtenerCargoSolicitud struct {
	claveUsuario  string
	nombreUsuario string
}

func (v *ObtenerCargoSolicitud) ObtenerClaveUsuario() (o string) {
	if v != nil {
		o = v.claveUsuario
	}
	return
}

func (v *ObtenerCargoSolicitud) AsignarClaveUsuario(claveUsuario string) {
	if v != nil {
		v.claveUsuario = claveUsuario
	}
}

func (v *ObtenerCargoSolicitud) ObtenerNombreUsuario() (o string) {
	if v != nil {
		o = v.nombreUsuario
	}
	return
}

func (v *ObtenerCargoSolicitud) AsignarNombreUsuario(nombreUsuario string) {
	if v != nil {
		v.nombreUsuario = nombreUsuario
	}
}

type GuardarUsuarioRespuesta struct {
	ok  bool
	err error
}

func (v *GuardarUsuarioRespuesta) ObtenerOk() (o bool) {
	if v != nil {
		o = v.ok
	}
	return
}

func (v *GuardarUsuarioRespuesta) AsignarOk(ok bool) {
	if v != nil {
		v.ok = ok
	}
}

func (v *GuardarUsuarioRespuesta) ObtenerErr() (o error) {
	if v != nil {
		o = v.err
	}
	return
}

func (v *GuardarUsuarioRespuesta) AsignarErr(err error) {
	if v != nil {
		v.err = err
	}
}

type ObtenerUsuarioPorNombreUsuarioSolicitud struct {
	nombre string
}

func (v *ObtenerUsuarioPorNombreUsuarioSolicitud) ObtenerNombre() (o string) {
	if v != nil {
		o = v.nombre
	}
	return
}

func (v *ObtenerUsuarioPorNombreUsuarioSolicitud) AsignarNombre(nombre string) {
	if v != nil {
		v.nombre = nombre
	}
}
