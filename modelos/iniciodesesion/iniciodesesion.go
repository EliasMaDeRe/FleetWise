package iniciodesesion

type Usuario struct {
	cargo         string
	claveUsuario  string
	ID            string
	nombreUsuario string
}

func (v *Usuario) ObtenerCargo() (o string) {
	if v != nil {
		o = v.cargo
	}
	return
}

func (v *Usuario) AsignarCargo(cargo string) {
	if v != nil {
		v.cargo = cargo
	}
}

func (v *Usuario) ObtenerMarca() (o string) {
	if v != nil {
		o = v.claveUsuario
	}
	return
}

func (v *Usuario) AsignarMarca(claveUsuario string) {
	if v != nil {
		v.claveUsuario = claveUsuario
	}
}

func (v *Usuario) ObtenerID() (o string) {
	if v != nil {
		o = v.ID
	}
	return
}

func (v *Usuario) AsignarID(ID string) {
	if v != nil {
		v.ID = ID
	}
}

func (v *Usuario) ObtenerNombreUsuario() (o string) {
	if v != nil {
		o = v.nombreUsuario
	}
	return
}

func (v *Usuario) AsignarNombreUsuario(nombreUsuario string) {
	if v != nil {
		v.nombreUsuario = nombreUsuario
	}
}

type InicioDeSesionSolicitud struct {
	claveUsuario  string
	nombreUsuario string
}

func (v *InicioDeSesionSolicitud) ObtenerClaveUsuario() (o string) {
	if v != nil {
		o = v.claveUsuario
	}
	return
}

func (v *InicioDeSesionSolicitud) AsignarClaveUsuario(claveUsuario string) {
	if v != nil {
		v.claveUsuario = claveUsuario
	}
}

func (v *InicioDeSesionSolicitud) ObtenerNombreUsuario() (o string) {
	if v != nil {
		o = v.nombreUsuario
	}
	return
}

func (v *InicioDeSesionSolicitud) AsignarNombreUsuario(nombreUsuario string) {
	if v != nil {
		v.nombreUsuario = nombreUsuario
	}
}
