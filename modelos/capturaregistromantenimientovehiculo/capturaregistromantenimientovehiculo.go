package capturavehiculos

type RegistroMantenimientoVehiculo struct {
	NumRegistro   int
	Tipo          string
	Fecha         string
	Litros        string
	Kilometraje   string
	Importe       string
	Observaciones string
	Concepto      string
}

func (r *RegistroMantenimientoVehiculo) ObtenerNumRegistro() (o int) {
	if r != nil {
		o = r.NumRegistro
	}
	return
}

func (r *RegistroMantenimientoVehiculo) AsignarNumRegistro(numRegistro int) {
	if r != nil {
		r.NumRegistro = numRegistro
	}
}

func (r *RegistroMantenimientoVehiculo) ObtenerTipo() (o string) {
	if r != nil {
		o = r.Tipo
	}
	return
}

func (r *RegistroMantenimientoVehiculo) AsignarTipo(tipo string) {
	if r != nil {
		r.Tipo = tipo
	}
}

func (r *RegistroMantenimientoVehiculo) ObtenerFecha() (o string) {
	if r != nil {
		o = r.Fecha
	}
	return
}

func (r *RegistroMantenimientoVehiculo) AsignarFecha(fecha string) {
	if r != nil {
		r.Fecha = fecha
	}
}

func (r *RegistroMantenimientoVehiculo) ObtenerLitros() (o string) {
	if r != nil {
		o = r.Litros
	}
	return
}

func (r *RegistroMantenimientoVehiculo) AsignarLitros(litros string) {
	if r != nil {
		r.Litros = litros
	}
}

func (r *RegistroMantenimientoVehiculo) ObtenerKilometraje() (o string) {
	if r != nil {
		o = r.Kilometraje
	}
	return
}

func (r *RegistroMantenimientoVehiculo) AsignarKilometraje(kilometraje string) {
	if r != nil {
		r.Kilometraje = kilometraje
	}
}

func (r *RegistroMantenimientoVehiculo) ObtenerImporte() (o string) {
	if r != nil {
		o = r.Importe
	}
	return
}

func (r *RegistroMantenimientoVehiculo) AsignarImporte(importe string) {
	if r != nil {
		r.Importe = importe
	}
}

func (r *RegistroMantenimientoVehiculo) ObtenerObservaciones() (o string) {
	if r != nil {
		o = r.Observaciones
	}
	return
}

func (r *RegistroMantenimientoVehiculo) AsignarObservaciones(observaciones string) {
	if r != nil {
		r.Observaciones = observaciones
	}
}

func (r *RegistroMantenimientoVehiculo) ObtenerConcepto() (o string) {
	if r != nil {
		o = r.Concepto
	}
	return
}

func (r *RegistroMantenimientoVehiculo) AsignarConcepto(concepto string) {
	if r != nil {
		r.Concepto = concepto
	}
}

type AgregarRegistroMantenimientoVehiculoSolicitud struct {
	NumRegistro   int
	Tipo          string
	Fecha         string
	Litros        string
	Kilometraje   string
	Importe       string
	Observaciones string
	Concepto      string
}

func (r *AgregarRegistroMantenimientoVehiculoSolicitud) ObtenerNumRegistro() (o int) {
	if r != nil {
		o = r.NumRegistro
	}
	return
}

func (r *AgregarRegistroMantenimientoVehiculoSolicitud) AsignarNumRegistro(numRegistro int) {
	if r != nil {
		r.NumRegistro = numRegistro
	}
}

func (r *AgregarRegistroMantenimientoVehiculoSolicitud) ObtenerTipo() (o string) {
	if r != nil {
		o = r.Tipo
	}
	return
}

func (r *AgregarRegistroMantenimientoVehiculoSolicitud) AsignarTipo(tipo string) {
	if r != nil {
		r.Tipo = tipo
	}
}

func (r *AgregarRegistroMantenimientoVehiculoSolicitud) ObtenerFecha() (o string) {
	if r != nil {
		o = r.Fecha
	}
	return
}

func (r *AgregarRegistroMantenimientoVehiculoSolicitud) AsignarFecha(fecha string) {
	if r != nil {
		r.Fecha = fecha
	}
}

func (r *AgregarRegistroMantenimientoVehiculoSolicitud) ObtenerLitros() (o string) {
	if r != nil {
		o = r.Litros
	}
	return
}

func (r *AgregarRegistroMantenimientoVehiculoSolicitud) AsignarLitros(litros string) {
	if r != nil {
		r.Litros = litros
	}
}

func (r *AgregarRegistroMantenimientoVehiculoSolicitud) ObtenerKilometraje() (o string) {
	if r != nil {
		o = r.Kilometraje
	}
	return
}

func (r *AgregarRegistroMantenimientoVehiculoSolicitud) AsignarKilometraje(kilometraje string) {
	if r != nil {
		r.Kilometraje = kilometraje
	}
}

func (r *AgregarRegistroMantenimientoVehiculoSolicitud) ObtenerImporte() (o string) {
	if r != nil {
		o = r.Importe
	}
	return
}

func (r *AgregarRegistroMantenimientoVehiculoSolicitud) AsignarImporte(importe string) {
	if r != nil {
		r.Importe = importe
	}
}

func (r *AgregarRegistroMantenimientoVehiculoSolicitud) ObtenerObservaciones() (o string) {
	if r != nil {
		o = r.Observaciones
	}
	return
}

func (r *AgregarRegistroMantenimientoVehiculoSolicitud) AsignarObservaciones(observaciones string) {
	if r != nil {
		r.Observaciones = observaciones
	}
}

func (r *AgregarRegistroMantenimientoVehiculoSolicitud) ObtenerConcepto() (o string) {
	if r != nil {
		o = r.Concepto
	}
	return
}

func (r *AgregarRegistroMantenimientoVehiculoSolicitud) AsignarConcepto(concepto string) {
	if r != nil {
		r.Concepto = concepto
	}
}

type AgregarRegistroMantenimientoVehiculoRespuesta struct {
	ok  bool
	err error
}

func (r *AgregarRegistroMantenimientoVehiculoRespuesta) ObtenerOk() (o bool) {
	if r != nil {
		o = r.ok
	}
	return
}

func (r *AgregarRegistroMantenimientoVehiculoRespuesta) AsignarOk(ok bool) {
	if r != nil {
		r.ok = ok
	}
}

func (r *AgregarRegistroMantenimientoVehiculoRespuesta) ObtenerErr() (o error) {
	if r != nil {
		o = r.err
	}
	return
}

func (r *AgregarRegistroMantenimientoVehiculoRespuesta) AsignarErr(err error) {
	if r != nil {
		r.err = err
	}
}

type EditarRegistroMantenimientoVehiculoRespuesta struct {
	ok  bool
	err error
}

func (e *EditarRegistroMantenimientoVehiculoRespuesta) ObtenerOk() (o bool) {
	if e != nil {
		o = e.ok
	}
	return
}

func (e *EditarRegistroMantenimientoVehiculoRespuesta) AsignarOk(ok bool) {
	if e != nil {
		e.ok = ok
	}
}

func (e *EditarRegistroMantenimientoVehiculoRespuesta) ObtenerErr() (o error) {
	if e != nil {
		o = e.err
	}
	return
}

func (e *EditarRegistroMantenimientoVehiculoRespuesta) AsignarErr(err error) {
	if e != nil {
		e.err = err
	}
}

type BorrarRegistroMantenimientoVehiculoRespuesta struct {
	ok  bool
	err error
}

func (b *BorrarRegistroMantenimientoVehiculoRespuesta) ObtenerOk() (o bool) {
	if b != nil {
		o = b.ok
	}
	return
}

func (b *BorrarRegistroMantenimientoVehiculoRespuesta) AsignarOk(ok bool) {
	if b != nil {
		b.ok = ok
	}
}

func (b *BorrarRegistroMantenimientoVehiculoRespuesta) ObtenerErr() (o error) {
	if b != nil {
		o = b.err
	}
	return
}

func (b *BorrarRegistroMantenimientoVehiculoRespuesta) AsignarErr(err error) {
	if b != nil {
		b.err = err
	}
}

type SeleccionarVehiculoSolicitud struct {
	Placas string
}

func (s *SeleccionarVehiculoSolicitud) ObtenerPlacas() (o string) {
	if s != nil {
		o = s.Placas
	}
	return
}

func (s *SeleccionarVehiculoSolicitud) AsignarPlacas(placas string) {
	if s != nil {
		s.Placas = placas
	}
}

type SeleccionarVehiculoRespuesta struct {
	ok  bool
	err error
}

func (s *SeleccionarVehiculoRespuesta) ObtenerOk() (o bool) {
	if s != nil {
		o = s.ok
	}
	return
}

func (s *SeleccionarVehiculoRespuesta) AsignarOk(ok bool) {
	if s != nil {
		s.ok = ok
	}
}

func (s *SeleccionarVehiculoRespuesta) ObtenerErr() (o error) {
	if s != nil {
		o = s.err
	}
	return
}

func (s *SeleccionarVehiculoRespuesta) AsignarErr(err error) {
	if s != nil {
		s.err = err
	}
}