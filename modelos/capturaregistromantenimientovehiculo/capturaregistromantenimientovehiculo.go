package capturaregistromantenimientovehiculo

type RegistroMantenimientoVehiculo struct {
	NumeroDeRegistro int
	Tipo             string
	Fecha            string
	LitrosDeGasolina float64
	Kilometraje      int
	Importe          float64
	Observaciones    string
	Concepto         string
}

func (RegistroMantenimientoVehiculo) TableName() string {
	return "registros_mantenimiento_vehicular"
}

func (r *RegistroMantenimientoVehiculo) ObtenerNumeroDeRegistro() (o int) {
	if r != nil {
		o = r.NumeroDeRegistro
	}
	return
}

func (r *RegistroMantenimientoVehiculo) AsignarNumeroDeRegistro(NumeroDeRegistro int) {
	if r != nil {
		r.NumeroDeRegistro = NumeroDeRegistro
	}
}

func (r *RegistroMantenimientoVehiculo) ObtenerTipo() (o string) {
	if r != nil {
		o = r.TipoRegistro
	}
	return
}

func (r *RegistroMantenimientoVehiculo) AsignarTipo(tipo string) {
	if r != nil {
		r.TipoRegistro = tipo
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

func (r *RegistroMantenimientoVehiculo) ObtenerLitrosDeGasolina() (o float64) {
	if r != nil {
		o = r.LitrosDeGasolina
	}
	return
}

func (r *RegistroMantenimientoVehiculo) AsignarLitrosDeGasolina(LitrosDeGasolina float64) {
	if r != nil {
		r.LitrosDeGasolina = LitrosDeGasolina
	}
}

func (r *RegistroMantenimientoVehiculo) ObtenerKilometraje() (o int) {
	if r != nil {
		o = r.Kilometraje
	}
	return
}

func (r *RegistroMantenimientoVehiculo) AsignarKilometraje(kilometraje int) {
	if r != nil {
		r.Kilometraje = kilometraje
	}
}

func (r *RegistroMantenimientoVehiculo) ObtenerImporte() (o float64) {
	if r != nil {
		o = r.Importe
	}
	return
}

func (r *RegistroMantenimientoVehiculo) AsignarImporte(importe float64) {
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
	NumeroDeRegistro int     `json:"numeroderegistro"`
	Tipo             string  `json:"tipo"`
	Fecha            string  `json:"fecha"`
	LitrosDeGasolina float64 `json:"litrosdegasolina"`
	Kilometraje      int     `json:"kilometraje"`
	Importe          float64 `json:"importe"`
	Observaciones    string  `json:"observaciones"`
	Concepto         string  `json:"concepto"`
}

func (r *AgregarRegistroMantenimientoVehiculoSolicitud) ObtenerNumeroDeRegistro() (o int) {
	if r != nil {
		o = r.NumeroDeRegistro
	}
	return
}

func (r *AgregarRegistroMantenimientoVehiculoSolicitud) AsignarNumeroDeRegistro(NumeroDeRegistro int) {
	if r != nil {
		r.NumeroDeRegistro = NumeroDeRegistro
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

func (r *AgregarRegistroMantenimientoVehiculoSolicitud) ObtenerLitrosDeGasolina() (o float64) {
	if r != nil {
		o = r.LitrosDeGasolina
	}
	return
}

func (r *AgregarRegistroMantenimientoVehiculoSolicitud) AsignarLitrosDeGasolina(LitrosDeGasolina float64) {
	if r != nil {
		r.LitrosDeGasolina = LitrosDeGasolina
	}
}

func (r *AgregarRegistroMantenimientoVehiculoSolicitud) ObtenerKilometraje() (o int) {
	if r != nil {
		o = r.Kilometraje
	}
	return
}

func (r *AgregarRegistroMantenimientoVehiculoSolicitud) AsignarKilometraje(kilometraje int) {
	if r != nil {
		r.Kilometraje = kilometraje
	}
}

func (r *AgregarRegistroMantenimientoVehiculoSolicitud) ObtenerImporte() (o float64) {
	if r != nil {
		o = r.Importe
	}
	return
}

func (r *AgregarRegistroMantenimientoVehiculoSolicitud) AsignarImporte(importe float64) {
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

type EditarRegistroMantenimientoVehiculoSolicitud struct{
	PlacaNueva string
	TipoDeRegistroNuevo string
	ImporteNuevo float32
	ServicioVehicularNuevo string
}

func (editarSolicitud *EditarRegistroMantenimientoVehiculoSolicitud) ObtenerPlacaNueva() (placa string) {
	if(editarSolicitud != nil){
		placa = editarSolicitud.PlacaNueva
	}
	return
} 

func (editarSolicitud *EditarRegistroMantenimientoVehiculoSolicitud) ObtenerTipoDeRegistro() (tipoDeRegistro string) {
	if(editarSolicitud != nil){
		tipoDeRegistro = editarSolicitud.TipoDeRegistroNuevo
	}
	return
} 

func (editarSolicitud *EditarRegistroMantenimientoVehiculoSolicitud) ObtenerImporteNuevo() (importe float32) {
	if(editarSolicitud != nil){
		importe = editarSolicitud.ImporteNuevo
	}
	return
} 

func (editarSolicitud *EditarRegistroMantenimientoVehiculoSolicitud) ObtenerServicioVehicularNuevo() (servicioVehicular string) {
	if(editarSolicitud != nil){
		servicioVehicular = editarSolicitud.ServicioVehicularNuevo
	}
	return
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

type SeleccionarRegistroMantenimientoVehicularSolicitud struct{
	NumeroDeRegistro int
}

func (seleccionarRegistroSolicitud *SeleccionarRegistroMantenimientoVehicularSolicitud) ObtenerNumRegistro() (numRegistro int){
	if(seleccionarRegistroSolicitud != nil){
		numRegistro = seleccionarRegistroSolicitud.NumeroDeRegistro
	}
	return
}

type SeleccionarRegistroMantenimientoVehicularRespuesta struct{
	Ok bool
	Registro map[string]interface{}
}

type SeleccionarVehiculoSolicitud struct {
	Placas string
type SeleccionarVehiculoParaNuevoRegistroSolicitud struct {
	Placas string `json:"placas"`
}



func (s *SeleccionarVehiculoParaNuevoRegistroSolicitud) ObtenerPlacas() (o string) {
	if s != nil {
		o = s.Placas
	}
	return
}

func (s *SeleccionarVehiculoParaNuevoRegistroSolicitud) AsignarPlacas(placas string) {
	if s != nil {
		s.Placas = placas
	}
}

type SeleccionarVehiculoParaNuevoRegistroRespuesta struct {
	ok  bool
	err error
}

// Error implements error.
func (*SeleccionarVehiculoParaNuevoRegistroRespuesta) Error() string {
	panic("unimplemented")
}

func (s *SeleccionarVehiculoParaNuevoRegistroRespuesta) ObtenerOk() (o bool) {
	if s != nil {
		o = s.ok
	}
	return
}

func (s *SeleccionarVehiculoParaNuevoRegistroRespuesta) AsignarOk(ok bool) {
	if s != nil {
		s.ok = ok
	}
}

func (s *SeleccionarVehiculoParaNuevoRegistroRespuesta) ObtenerErr() (o error) {
	if s != nil {
		o = s.err
	}
	return
}

func (s *SeleccionarVehiculoParaNuevoRegistroRespuesta) AsignarErr(err error) {
	if s != nil {
		s.err = err
	}
}
