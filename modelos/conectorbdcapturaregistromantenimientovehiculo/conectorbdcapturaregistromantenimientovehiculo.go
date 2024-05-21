package conectorbdcapturaregistromantenimientovehiculo


type GuardarRegistroMantenimientoVehiculoRespuesta struct {
	ok  bool
	err error
}

func (v *GuardarRegistroMantenimientoVehiculoRespuesta) ObtenerOk() (o bool) {
	if v != nil {
		o = v.ok
	}
	return
}

func (v *GuardarRegistroMantenimientoVehiculoRespuesta) AsignarOk(ok bool) {
	if v != nil {
		v.ok = ok
	}
}

func (v *GuardarRegistroMantenimientoVehiculoRespuesta) ObtenerErr() (o error) {
	if v != nil {
		o = v.err
	}
	return
}

func (v *GuardarRegistroMantenimientoVehiculoRespuesta) AsignarErr(err error) {
	if v != nil {
		v.err = err
	}
}


type ActualizarRegistroMantenimientoVehiculoRespuesta struct {
	Ok  bool
	Err error
}

func (editarRegistroRespuesta *ActualizarRegistroMantenimientoVehiculoRespuesta) ObtenerOk() (ok bool) {
	if editarRegistroRespuesta != nil {
		ok = editarRegistroRespuesta.Ok
	}
	return
}

func (editarRegistroRespuesta *ActualizarRegistroMantenimientoVehiculoRespuesta) AsignarOk(ok bool) {
	if editarRegistroRespuesta != nil {
		editarRegistroRespuesta.Ok = ok
	}
}

func (editarRegistroRespuesta *ActualizarRegistroMantenimientoVehiculoRespuesta) ObtenerErr() (err error) {
	if editarRegistroRespuesta != nil {
		err = editarRegistroRespuesta.Err
	}
	return
}

func (editarRegistroRespuesta *ActualizarRegistroMantenimientoVehiculoRespuesta) AsignarErr(err error) {
	if editarRegistroRespuesta != nil {
		editarRegistroRespuesta.Err = err
	}
}


type EditarRegistroDeMantenimientoDelVehiculoSolicitud struct {
	NumeroDeRegistro int
	Tipo             string
	Fecha            string
	LitrosDeGasolina float64
	Kilometraje      int
	Importe          float64
	Observaciones    string
	Concepto         string
	PlacasVehiculo   string
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