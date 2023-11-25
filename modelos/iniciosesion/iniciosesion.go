package iniciosesion

type Usuario struct {
	Cargo         string
	ClaveUsuario  string
	NombreUsuario string
}

func (v *Usuario) ObtenerCargo() (o string) {
	if v != nil {
		o = v.Cargo
	}
	return
}

func (v *Usuario) AsignarCargo(cargo string) {
	if v != nil {
		v.Cargo = cargo
	}
}

func (v *Usuario) ObtenerClaveUsuario() (o string) {
	if v != nil {
		o = v.ClaveUsuario
	}
	return
}

func (v *Usuario) AsignarClaveUsuario(claveUsuario string) {
	if v != nil {
		v.ClaveUsuario = claveUsuario
	}
}

func (v *Usuario) ObtenerNombreUsuario() (o string) {
	if v != nil {
		o = v.NombreUsuario
	}
	return
}

func (v *Usuario) AsignarNombreUsuario(nombreUsuario string) {
	if v != nil {
		v.NombreUsuario = nombreUsuario
	}
}

type IniciarSesionSolicitud struct {
	ClaveUsuario  string `json:"claveusuario"`
	NombreUsuario string `json:"nombreusuario"`
}

func (v *IniciarSesionSolicitud) ObtenerClaveUsuario() (o string) {
	if v != nil {
		o = v.ClaveUsuario
	}
	return
}

func (v *IniciarSesionSolicitud) AsignarClaveUsuario(claveUsuario string) {
	if v != nil {
		v.ClaveUsuario = claveUsuario
	}
}

func (v *IniciarSesionSolicitud) ObtenerNombreUsuario() (o string) {
	if v != nil {
		o = v.NombreUsuario
	}
	return
}

func (v *IniciarSesionSolicitud) AsignarNombreUsuario(nombreUsuario string) {
	if v != nil {
		v.NombreUsuario = nombreUsuario
	}
}

type IniciarSesionRespuesta struct {
	ok  bool
	err error
}

func (v *IniciarSesionRespuesta) ObtenerOk() (o bool) {
	if v != nil {
		o = v.ok
	}
	return
}

func (v *IniciarSesionRespuesta) AsignarOk(ok bool) {
	if v != nil {
		v.ok = ok
	}
}

func (v *IniciarSesionRespuesta) ObtenerErr() (o error) {
	if v != nil {
		o = v.err
	}
	return
}

func (v *IniciarSesionRespuesta) AsignarErr(err error) {
	if v != nil {
		v.err = err
	}
}

type ObtenerUsuarioPorNombreUsuario struct {
	NombreUsuario string `json:"nombreusuario"`
}

func (v *ObtenerUsuarioPorNombreUsuario) ObtenerNombreUsuario() (o string) {
	if v != nil {
		o = v.NombreUsuario
	}
	return
}

func (v *ObtenerUsuarioPorNombreUsuario) AsignarNombreUsuario(nombreUsuario string) {
	if v != nil {
		v.NombreUsuario = nombreUsuario
	}
}
