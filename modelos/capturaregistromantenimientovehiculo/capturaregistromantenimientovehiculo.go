package capturaregistromantenimientovehiculo

type RegistroMantenimientoVehiculo struct {
	NumeroDeRegistro int
	TipoDeRegistro   string
	Fecha            string
	LitrosDeGasolina float64
	Kilometraje      int
	Importe          float64
	Observaciones    string
	Concepto         string
	PlacasVehiculo   string
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
		o = r.TipoDeRegistro
	}
	return
}

func (r *RegistroMantenimientoVehiculo) AsignarTipo(TipoDeRegistro string) {
	if r != nil {
		r.TipoDeRegistro = TipoDeRegistro
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

func (r *RegistroMantenimientoVehiculo) ObtenerPlacasVehiculo() (placasVehiculo string) {
	if r != nil {
		placasVehiculo = r.PlacasVehiculo
	}
	return
}

func (r *RegistroMantenimientoVehiculo) AsignarPlacasVehiculo(placasVehiculo string) {
	if r != nil {
		r.PlacasVehiculo = placasVehiculo
	}
}

type AgregarRegistroMantenimientoVehiculoSolicitud struct {
	NumeroDeRegistro string `json:"numeroderegistro"`
	TipoDeRegistro   string `json:"TipoDeRegistro"`
	Fecha            string `json:"fecha"`
	LitrosDeGasolina string `json:"litrosdegasolina"`
	Kilometraje      string `json:"kilometraje"`
	Importe          string `json:"importe"`
	Observaciones    string `json:"observaciones"`
	Concepto         string `json:"conceptoregistro"`
	PlacasVehiculo   string `json:"placasVehiculo"`
}

func (r *AgregarRegistroMantenimientoVehiculoSolicitud) ObtenerNumeroDeRegistro() (o string) {
	if r != nil {
		o = r.NumeroDeRegistro
	}
	return
}

func (r *AgregarRegistroMantenimientoVehiculoSolicitud) AsignarNumeroDeRegistro(NumeroDeRegistro string) {
	if r != nil {
		r.NumeroDeRegistro = NumeroDeRegistro
	}
}

func (r *AgregarRegistroMantenimientoVehiculoSolicitud) ObtenerTipo() (o string) {
	if r != nil {
		o = r.TipoDeRegistro
	}
	return
}

func (r *AgregarRegistroMantenimientoVehiculoSolicitud) AsignarTipo(TipoDeRegistro string) {
	if r != nil {
		r.TipoDeRegistro = TipoDeRegistro
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

func (r *AgregarRegistroMantenimientoVehiculoSolicitud) ObtenerLitrosDeGasolina() (o string) {
	if r != nil {
		o = r.LitrosDeGasolina
	}
	return
}

func (r *AgregarRegistroMantenimientoVehiculoSolicitud) AsignarLitrosDeGasolina(LitrosDeGasolina string) {
	if r != nil {
		r.LitrosDeGasolina = LitrosDeGasolina
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

func (agregarRegistroSolicitud *AgregarRegistroMantenimientoVehiculoSolicitud) ObtenerPlacasVehiculo() (placasVehiculo string) {
	if agregarRegistroSolicitud != nil {
		placasVehiculo = agregarRegistroSolicitud.PlacasVehiculo
	}
	return
}

func (r *AgregarRegistroMantenimientoVehiculoSolicitud) AsignarPlacasVehiculo(placasVehiculo string) {
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

type EditarRegistroDeMantenimientoDelVehiculoSolicitud struct {
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

func (editarRegistroSolicitud *EditarRegistroDeMantenimientoDelVehiculoSolicitud) ObtenerNumeroDeRegistro() (numeroDeRegistro int) {
	if editarRegistroSolicitud != nil {
		numeroDeRegistro = editarRegistroSolicitud.NumeroDeRegistro
	}
	return
}

func (editarRegistroSolicitud *EditarRegistroDeMantenimientoDelVehiculoSolicitud) AsignarNumeroDeRegistro(NumeroDeRegistro int) {
	if editarRegistroSolicitud != nil {
		editarRegistroSolicitud.NumeroDeRegistro = NumeroDeRegistro
	}
}

func (editarRegistroSolicitud *EditarRegistroDeMantenimientoDelVehiculoSolicitud) ObtenerTipo() (tipo string) {
	if editarRegistroSolicitud != nil {
		tipo = editarRegistroSolicitud.Tipo
	}
	return
}

func (editarRegistroSolicitud *EditarRegistroDeMantenimientoDelVehiculoSolicitud) AsignarTipo(tipo string) {
	if editarRegistroSolicitud != nil {
		editarRegistroSolicitud.Tipo = tipo
	}
}

func (editarRegistroSolicitud *EditarRegistroDeMantenimientoDelVehiculoSolicitud) ObtenerFecha() (fecha string) {
	if editarRegistroSolicitud != nil {
		fecha = editarRegistroSolicitud.Fecha
	}
	return
}

func (editarRegistroSolicitud *EditarRegistroDeMantenimientoDelVehiculoSolicitud) AsignarFecha(fecha string) {
	if editarRegistroSolicitud != nil {
		editarRegistroSolicitud.Fecha = fecha
	}
}

func (editarRegistroSolicitud *EditarRegistroDeMantenimientoDelVehiculoSolicitud) ObtenerLitrosDeGasolina() (litrosGasolina float64) {
	if editarRegistroSolicitud != nil {
		litrosGasolina = editarRegistroSolicitud.LitrosDeGasolina
	}
	return
}

func (editarRegistroSolicitud *EditarRegistroDeMantenimientoDelVehiculoSolicitud) AsignarLitrosDeGasolina(LitrosDeGasolina float64) {
	if editarRegistroSolicitud != nil {
		editarRegistroSolicitud.LitrosDeGasolina = LitrosDeGasolina
	}
}

func (editarRegistroSolicitud *EditarRegistroDeMantenimientoDelVehiculoSolicitud) ObtenerKilometraje() (kilometraje int) {
	if editarRegistroSolicitud != nil {
		kilometraje = editarRegistroSolicitud.Kilometraje
	}
	return
}

func (editarRegistroSolicitud *EditarRegistroDeMantenimientoDelVehiculoSolicitud) AsignarKilometraje(kilometraje int) {
	if editarRegistroSolicitud != nil {
		editarRegistroSolicitud.Kilometraje = kilometraje
	}
}

func (editarRegistroSolicitud *EditarRegistroDeMantenimientoDelVehiculoSolicitud) ObtenerImporte() (importe float64) {
	if editarRegistroSolicitud != nil {
		importe = editarRegistroSolicitud.Importe
	}
	return
}

func (editarRegistroSolicitud *EditarRegistroDeMantenimientoDelVehiculoSolicitud) AsignarImporte(importe float64) {
	if editarRegistroSolicitud != nil {
		editarRegistroSolicitud.Importe = importe
	}
}

func (editarRegistroSolicitud *EditarRegistroDeMantenimientoDelVehiculoSolicitud) ObtenerObservaciones() (observaciones string) {
	if editarRegistroSolicitud != nil {
		observaciones = editarRegistroSolicitud.Observaciones
	}
	return
}

func (editarRegistroSolicitud *EditarRegistroDeMantenimientoDelVehiculoSolicitud) AsignarObservaciones(observaciones string) {
	if editarRegistroSolicitud != nil {
		editarRegistroSolicitud.Observaciones = observaciones
	}
}

func (editarRegistroSolicitud *EditarRegistroDeMantenimientoDelVehiculoSolicitud) ObtenerConcepto() (concepto string) {
	if editarRegistroSolicitud != nil {
		concepto = editarRegistroSolicitud.Concepto
	}
	return
}

func (editarRegistroSolicitud *EditarRegistroDeMantenimientoDelVehiculoSolicitud) AsignarConcepto(concepto string) {
	if editarRegistroSolicitud != nil {
		editarRegistroSolicitud.Concepto = concepto
	}
}

func (editarRegistroSolicitud *EditarRegistroDeMantenimientoDelVehiculoSolicitud) ObtenerPlacasVehiculo() (placasVehiculo string) {
	if editarRegistroSolicitud != nil {
		placasVehiculo = editarRegistroSolicitud.PlacasVehiculo
	}
	return
}

func (editarRegistroSolicitud *EditarRegistroDeMantenimientoDelVehiculoSolicitud) AsignarPlacasVehiculo(placasVehiculo string) {
	if editarRegistroSolicitud != nil {
		editarRegistroSolicitud.PlacasVehiculo = placasVehiculo
	}
}
