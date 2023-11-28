package capturaregistromantenimientovehiculo

type RegistroDeMantenimientoDeVehiculo struct {
	NumeroDeRegistro int
	TipoDeRegistro   string
	Fecha            string
	LitrosDeGasolina float64
	Kilometraje      int
	Importe          float64
	Observaciones    string
	Concepto         string
	PlacasDeVehiculo string
}

func (RegistroDeMantenimientoDeVehiculo) TableName() string {
	return "registros_de_mantenimiento_de_vehiculo"
}

func (r *RegistroDeMantenimientoDeVehiculo) ObtenerNumeroDeRegistro() (o int) {
	if r != nil {
		o = r.NumeroDeRegistro
	}
	return
}

func (r *RegistroDeMantenimientoDeVehiculo) AsignarNumeroDeRegistro(NumeroDeRegistro int) {
	if r != nil {
		r.NumeroDeRegistro = NumeroDeRegistro
	}
}

func (r *RegistroDeMantenimientoDeVehiculo) ObtenerTipo() (o string) {
	if r != nil {
		o = r.TipoDeRegistro
	}
	return
}

func (r *RegistroDeMantenimientoDeVehiculo) AsignarTipo(TipoDeRegistro string) {
	if r != nil {
		r.TipoDeRegistro = TipoDeRegistro
	}
}

func (r *RegistroDeMantenimientoDeVehiculo) ObtenerFecha() (o string) {
	if r != nil {
		o = r.Fecha
	}
	return
}

func (r *RegistroDeMantenimientoDeVehiculo) AsignarFecha(fecha string) {
	if r != nil {
		r.Fecha = fecha
	}
}

func (r *RegistroDeMantenimientoDeVehiculo) ObtenerLitrosDeGasolina() (o float64) {
	if r != nil {
		o = r.LitrosDeGasolina
	}
	return
}

func (r *RegistroDeMantenimientoDeVehiculo) AsignarLitrosDeGasolina(LitrosDeGasolina float64) {
	if r != nil {
		r.LitrosDeGasolina = LitrosDeGasolina
	}
}

func (r *RegistroDeMantenimientoDeVehiculo) ObtenerKilometraje() (o int) {
	if r != nil {
		o = r.Kilometraje
	}
	return
}

func (r *RegistroDeMantenimientoDeVehiculo) AsignarKilometraje(kilometraje int) {
	if r != nil {
		r.Kilometraje = kilometraje
	}
}

func (r *RegistroDeMantenimientoDeVehiculo) ObtenerImporte() (o float64) {
	if r != nil {
		o = r.Importe
	}
	return
}

func (r *RegistroDeMantenimientoDeVehiculo) AsignarImporte(importe float64) {
	if r != nil {
		r.Importe = importe
	}
}

func (r *RegistroDeMantenimientoDeVehiculo) ObtenerObservaciones() (o string) {
	if r != nil {
		o = r.Observaciones
	}
	return
}

func (r *RegistroDeMantenimientoDeVehiculo) AsignarObservaciones(observaciones string) {
	if r != nil {
		r.Observaciones = observaciones
	}
}

func (r *RegistroDeMantenimientoDeVehiculo) ObtenerConcepto() (o string) {
	if r != nil {
		o = r.Concepto
	}
	return
}

func (r *RegistroDeMantenimientoDeVehiculo) AsignarConcepto(concepto string) {
	if r != nil {
		r.Concepto = concepto
	}
}

func (r *RegistroDeMantenimientoDeVehiculo) ObtenerPlacasVehiculo() (placasVehiculo string) {
	if r != nil {
		placasVehiculo = r.PlacasDeVehiculo
	}
	return
}

func (r *RegistroDeMantenimientoDeVehiculo) AsignarPlacasVehiculo(placasVehiculo string) {
	if r != nil {
		r.PlacasDeVehiculo = placasVehiculo
	}
}

type AgregarRegistroDeMantenimientoDeVehiculoSolicitud struct {
	NumeroDeRegistro string `json:"numeroderegistro"`
	TipoDeRegistro   string `json:"tiporegistro"`
	Fecha            string `json:"fecha"`
	LitrosDeGasolina string `json:"litrosdegasolina"`
	Kilometraje      string `json:"kilometraje"`
	Importe          string `json:"importe"`
	Observaciones    string `json:"observaciones"`
	Concepto         string `json:"conceptoregistro"`
	PlacasVehiculo   string `json:"placasVehiculo"`
}

func (r *AgregarRegistroDeMantenimientoDeVehiculoSolicitud) ObtenerNumeroDeRegistro() (o string) {
	if r != nil {
		o = r.NumeroDeRegistro
	}
	return
}

func (r *AgregarRegistroDeMantenimientoDeVehiculoSolicitud) AsignarNumeroDeRegistro(NumeroDeRegistro string) {
	if r != nil {
		r.NumeroDeRegistro = NumeroDeRegistro
	}
}

func (r *AgregarRegistroDeMantenimientoDeVehiculoSolicitud) ObtenerTipo() (o string) {
	if r != nil {
		o = r.TipoDeRegistro
	}
	return
}

func (r *AgregarRegistroDeMantenimientoDeVehiculoSolicitud) AsignarTipo(TipoDeRegistro string) {
	if r != nil {
		r.TipoDeRegistro = TipoDeRegistro
	}
}

func (r *AgregarRegistroDeMantenimientoDeVehiculoSolicitud) ObtenerFecha() (o string) {
	if r != nil {
		o = r.Fecha
	}
	return
}

func (r *AgregarRegistroDeMantenimientoDeVehiculoSolicitud) AsignarFecha(fecha string) {
	if r != nil {
		r.Fecha = fecha
	}
}

func (r *AgregarRegistroDeMantenimientoDeVehiculoSolicitud) ObtenerLitrosDeGasolina() (o string) {
	if r != nil {
		o = r.LitrosDeGasolina
	}
	return
}

func (r *AgregarRegistroDeMantenimientoDeVehiculoSolicitud) AsignarLitrosDeGasolina(LitrosDeGasolina string) {
	if r != nil {
		r.LitrosDeGasolina = LitrosDeGasolina
	}
}

func (r *AgregarRegistroDeMantenimientoDeVehiculoSolicitud) ObtenerKilometraje() (o string) {
	if r != nil {
		o = r.Kilometraje
	}
	return
}

func (r *AgregarRegistroDeMantenimientoDeVehiculoSolicitud) AsignarKilometraje(kilometraje string) {
	if r != nil {
		r.Kilometraje = kilometraje
	}
}

func (r *AgregarRegistroDeMantenimientoDeVehiculoSolicitud) ObtenerImporte() (o string) {
	if r != nil {
		o = r.Importe
	}
	return
}

func (r *AgregarRegistroDeMantenimientoDeVehiculoSolicitud) AsignarImporte(importe string) {
	if r != nil {
		r.Importe = importe
	}
}

func (r *AgregarRegistroDeMantenimientoDeVehiculoSolicitud) ObtenerObservaciones() (o string) {
	if r != nil {
		o = r.Observaciones
	}
	return
}

func (r *AgregarRegistroDeMantenimientoDeVehiculoSolicitud) AsignarObservaciones(observaciones string) {
	if r != nil {
		r.Observaciones = observaciones
	}
}

func (r *AgregarRegistroDeMantenimientoDeVehiculoSolicitud) ObtenerConcepto() (o string) {
	if r != nil {
		o = r.Concepto
	}
	return
}

func (r *AgregarRegistroDeMantenimientoDeVehiculoSolicitud) AsignarConcepto(concepto string) {
	if r != nil {
		r.Concepto = concepto
	}
}

func (agregarRegistroSolicitud *AgregarRegistroDeMantenimientoDeVehiculoSolicitud) ObtenerPlacasVehiculo() (placasVehiculo string) {
	if agregarRegistroSolicitud != nil {
		placasVehiculo = agregarRegistroSolicitud.PlacasVehiculo
	}
	return
}

func (r *AgregarRegistroDeMantenimientoDeVehiculoSolicitud) AsignarPlacasVehiculo(placasVehiculo string) {
	if r != nil {
		r.PlacasVehiculo = placasVehiculo
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

type EditarRegistroMantenimientoVehiculoSolicitud struct {
	PlacaNueva             string
	TipoDeRegistroNuevo    string
	ImporteNuevo           float32
	ServicioVehicularNuevo string
}

func (editarSolicitud *EditarRegistroMantenimientoVehiculoSolicitud) ObtenerPlacaNueva() (placa string) {
	if editarSolicitud != nil {
		placa = editarSolicitud.PlacaNueva
	}
	return
}

func (editarSolicitud *EditarRegistroMantenimientoVehiculoSolicitud) ObtenerTipoDeRegistro() (tipoDeRegistro string) {
	if editarSolicitud != nil {
		tipoDeRegistro = editarSolicitud.TipoDeRegistroNuevo
	}
	return
}

func (editarSolicitud *EditarRegistroMantenimientoVehiculoSolicitud) ObtenerImporteNuevo() (importe float32) {
	if editarSolicitud != nil {
		importe = editarSolicitud.ImporteNuevo
	}
	return
}

func (editarSolicitud *EditarRegistroMantenimientoVehiculoSolicitud) ObtenerServicioVehicularNuevo() (servicioVehicular string) {
	if editarSolicitud != nil {
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

type SeleccionarRegistroMantenimientoVehiculoSolicitud struct {
	NumeroDeRegistro int
}

func (seleccionarRegistroSolicitud *SeleccionarRegistroMantenimientoVehiculoSolicitud) ObtenerNumRegistro() (numRegistro int) {
	if seleccionarRegistroSolicitud != nil {
		numRegistro = seleccionarRegistroSolicitud.NumeroDeRegistro
	}
	return
}

type SeleccionarRegistroMantenimientoVehiculoRespuesta struct {
	Ok       bool
	Registro map[string]interface{}
}

type SeleccionarVehiculoSolicitud struct {
	Placas string
}

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

type ObtenerRegistroDeMantenimientoDelVehiculoPorNumeroDeRegistroSolicitud struct {
	NumDeRegistro int `json:"NumDeRegistro"`
}

func (obtenerRegistroSolicitud *ObtenerRegistroDeMantenimientoDelVehiculoPorNumeroDeRegistroSolicitud) ObtenerNumDeRegistro() (numDeRegistro int) {
	if obtenerRegistroSolicitud != nil {
		numDeRegistro = obtenerRegistroSolicitud.NumDeRegistro
	}
	return
}

type ObtenerRegistroYVehiculoAsociadoPorNumeroDeRegistroSolicitud struct {
	NumDeRegistro int
}

func (obtenerRegistroYVehiculoAsociadoSolicitud *ObtenerRegistroYVehiculoAsociadoPorNumeroDeRegistroSolicitud) ObtenerNumDeRegistro() (numDeRegistro int) {
	if obtenerRegistroYVehiculoAsociadoSolicitud != nil {
		numDeRegistro = obtenerRegistroYVehiculoAsociadoSolicitud.NumDeRegistro
	}
	return
}

func (obtenerRegistroYVehiculoAsociadoSolicitud *ObtenerRegistroYVehiculoAsociadoPorNumeroDeRegistroSolicitud) AsignarNumDeRegistro(numDeRegistro int) {
	if obtenerRegistroYVehiculoAsociadoSolicitud != nil {
		obtenerRegistroYVehiculoAsociadoSolicitud.NumDeRegistro = numDeRegistro
	}
}

type EditarRegistroDeMantenimientoDeVehiculoSolicitud struct {
	NumeroDeRegistro int     `json:"numeroDeRegistro"`
	Tipo             string  `json:"tipoDeRegistro"`
	Fecha            string  `json:"fecha"`
	LitrosDeGasolina float64 `json:"litrosDeGasolina"`
	Kilometraje      int     `json:"kilometraje"`
	Importe          float64 `json:"importe"`
	Observaciones    string  `json:"observaciones"`
	Concepto         string  `json:"concepto"`
	PlacasVehiculo   string  `json:"placaVehiculo"`
}

func (editarRegistroSolicitud *EditarRegistroDeMantenimientoDeVehiculoSolicitud) ObtenerNumeroDeRegistro() (numeroDeRegistro int) {
	if editarRegistroSolicitud != nil {
		numeroDeRegistro = editarRegistroSolicitud.NumeroDeRegistro
	}
	return
}

func (editarRegistroSolicitud *EditarRegistroDeMantenimientoDeVehiculoSolicitud) AsignarNumeroDeRegistro(NumeroDeRegistro int) {
	if editarRegistroSolicitud != nil {
		editarRegistroSolicitud.NumeroDeRegistro = NumeroDeRegistro
	}
}

func (editarRegistroSolicitud *EditarRegistroDeMantenimientoDeVehiculoSolicitud) ObtenerTipo() (tipo string) {
	if editarRegistroSolicitud != nil {
		tipo = editarRegistroSolicitud.Tipo
	}
	return
}

func (editarRegistroSolicitud *EditarRegistroDeMantenimientoDeVehiculoSolicitud) AsignarTipo(tipo string) {
	if editarRegistroSolicitud != nil {
		editarRegistroSolicitud.Tipo = tipo
	}
}

func (editarRegistroSolicitud *EditarRegistroDeMantenimientoDeVehiculoSolicitud) ObtenerFecha() (fecha string) {
	if editarRegistroSolicitud != nil {
		fecha = editarRegistroSolicitud.Fecha
	}
	return
}

func (editarRegistroSolicitud *EditarRegistroDeMantenimientoDeVehiculoSolicitud) AsignarFecha(fecha string) {
	if editarRegistroSolicitud != nil {
		editarRegistroSolicitud.Fecha = fecha
	}
}

func (editarRegistroSolicitud *EditarRegistroDeMantenimientoDeVehiculoSolicitud) ObtenerLitrosDeGasolina() (litrosGasolina float64) {
	if editarRegistroSolicitud != nil {
		litrosGasolina = editarRegistroSolicitud.LitrosDeGasolina
	}
	return
}

func (editarRegistroSolicitud *EditarRegistroDeMantenimientoDeVehiculoSolicitud) AsignarLitrosDeGasolina(LitrosDeGasolina float64) {
	if editarRegistroSolicitud != nil {
		editarRegistroSolicitud.LitrosDeGasolina = LitrosDeGasolina
	}
}

func (editarRegistroSolicitud *EditarRegistroDeMantenimientoDeVehiculoSolicitud) ObtenerKilometraje() (kilometraje int) {
	if editarRegistroSolicitud != nil {
		kilometraje = editarRegistroSolicitud.Kilometraje
	}
	return
}

func (editarRegistroSolicitud *EditarRegistroDeMantenimientoDeVehiculoSolicitud) AsignarKilometraje(kilometraje int) {
	if editarRegistroSolicitud != nil {
		editarRegistroSolicitud.Kilometraje = kilometraje
	}
}

func (editarRegistroSolicitud *EditarRegistroDeMantenimientoDeVehiculoSolicitud) ObtenerImporte() (importe float64) {
	if editarRegistroSolicitud != nil {
		importe = editarRegistroSolicitud.Importe
	}
	return
}

func (editarRegistroSolicitud *EditarRegistroDeMantenimientoDeVehiculoSolicitud) AsignarImporte(importe float64) {
	if editarRegistroSolicitud != nil {
		editarRegistroSolicitud.Importe = importe
	}
}

func (editarRegistroSolicitud *EditarRegistroDeMantenimientoDeVehiculoSolicitud) ObtenerObservaciones() (observaciones string) {
	if editarRegistroSolicitud != nil {
		observaciones = editarRegistroSolicitud.Observaciones
	}
	return
}

func (editarRegistroSolicitud *EditarRegistroDeMantenimientoDeVehiculoSolicitud) AsignarObservaciones(observaciones string) {
	if editarRegistroSolicitud != nil {
		editarRegistroSolicitud.Observaciones = observaciones
	}
}

func (editarRegistroSolicitud *EditarRegistroDeMantenimientoDeVehiculoSolicitud) ObtenerConcepto() (concepto string) {
	if editarRegistroSolicitud != nil {
		concepto = editarRegistroSolicitud.Concepto
	}
	return
}

func (editarRegistroSolicitud *EditarRegistroDeMantenimientoDeVehiculoSolicitud) AsignarConcepto(concepto string) {
	if editarRegistroSolicitud != nil {
		editarRegistroSolicitud.Concepto = concepto
	}
}

func (editarRegistroSolicitud *EditarRegistroDeMantenimientoDeVehiculoSolicitud) ObtenerPlacasVehiculo() (placasVehiculo string) {
	if editarRegistroSolicitud != nil {
		placasVehiculo = editarRegistroSolicitud.PlacasVehiculo
	}
	return
}

func (editarRegistroSolicitud *EditarRegistroDeMantenimientoDeVehiculoSolicitud) AsignarPlacasVehiculo(placasVehiculo string) {
	if editarRegistroSolicitud != nil {
		editarRegistroSolicitud.PlacasVehiculo = placasVehiculo
	}
}
