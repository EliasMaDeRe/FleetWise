package conectorbd

import (
	capturaVehiculosModelos "example/fleetwise/modelos/capturavehiculos"
)

type GuardarVehiculoSolicitud struct {
	vehiculo capturaVehiculosModelos.Vehiculo
}

func (v *GuardarVehiculoSolicitud) ObtenerVehiculo() (o capturaVehiculosModelos.Vehiculo) {
	if v != nil {
		o = v.vehiculo
	}
	return
}

func (v *GuardarVehiculoSolicitud) AsignarVehiculo(vehiculo capturaVehiculosModelos.Vehiculo) {
	if v != nil {
		v.vehiculo = vehiculo
	}
}

type ObtenerCargoSoicitud struct {
	claveUsuario  string
	nombreUsuario string
}

func (v *ObtenerCargoSoicitud) ObtenerClaveUsuario() (o string) {
	if v != nil {
		o = v.claveUsuario
	}
	return
}

func (v *ObtenerCargoSoicitud) AsignarClaveUsuario(claveUsuario string) {
	if v != nil {
		v.claveUsuario = claveUsuario
	}
}

func (v *ObtenerCargoSoicitud) ObtenerNombreUsuario() (o string) {
	if v != nil {
		o = v.nombreUsuario
	}
	return
}

func (v *ObtenerCargoSoicitud) AsignarNombreUsuario(nombreUsuario string) {
	if v != nil {
		v.nombreUsuario = nombreUsuario
	}
}

type ObtenerVehiculoPorPlacasSolicitud struct {
	placa string
}

func (v *ObtenerVehiculoPorPlacasSolicitud) ObtenerPlacas() (o string) {
	if v != nil {
		o = v.placa
	}
	return
}

func (v *ObtenerVehiculoPorPlacasSolicitud) AsignarPlacas(placa string) {
	if v != nil {
		v.placa = placa
	}
}

type ObtenerVehiculoPorSerieSolicitud struct {
	serie string
}

func (v *ObtenerVehiculoPorSerieSolicitud) ObtenerSerie() (o string) {
	if v != nil {
		o = v.serie
	}
	return
}

func (v *ObtenerVehiculoPorSerieSolicitud) AsignarSerie(serie string) {
	if v != nil {
		v.serie = serie
	}
}

type GuardarVehiculoRespuesta struct {
	ok  bool
	err error
}

func (v *GuardarVehiculoRespuesta) ObtenerOk() (o bool) {
	if v != nil {
		o = v.ok
	}
	return
}

func (v *GuardarVehiculoRespuesta) AsignarOk(ok bool) {
	if v != nil {
		v.ok = ok
	}
}

func (v *GuardarVehiculoRespuesta) ObtenerErr() (o error) {
	if v != nil {
		o = v.err
	}
	return
}

func (v *GuardarVehiculoRespuesta) AsignarErr(err error) {
	if v != nil {
		v.err = err
	}
}

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

type ObtenerRegistrosYVehiculosAsociadosFiltradosSolicitud struct {
	filtroPlaca              string
	filtroMarca              string
	filtroModelo             string
	filtroFechaDeLanzamiento int
	filtroTipoDeRegistro     string
}

func (registrosFiltradosSolicitud *ObtenerRegistrosYVehiculosAsociadosFiltradosSolicitud) ObtenerFiltroPlaca() (filtroPlaca string) {
	if registrosFiltradosSolicitud != nil {
		filtroPlaca = registrosFiltradosSolicitud.filtroPlaca
	}
	return
}

func (registrosFiltradosSolicitud *ObtenerRegistrosYVehiculosAsociadosFiltradosSolicitud) ObtenerFiltroMarca() (filtroMarca string) {
	if registrosFiltradosSolicitud != nil {
		filtroMarca = registrosFiltradosSolicitud.filtroMarca
	}
	return
}

func (registrosFiltradosSolicitud *ObtenerRegistrosYVehiculosAsociadosFiltradosSolicitud) ObtenerFiltroModelo() (filtroModelo string) {
	if registrosFiltradosSolicitud != nil {
		filtroModelo = registrosFiltradosSolicitud.filtroModelo
	}
	return
}

func (registrosFiltradosSolicitud *ObtenerRegistrosYVehiculosAsociadosFiltradosSolicitud) ObtenerFiltroFechaDeLanzamiento() (filtroFechaDeLanzamiento int) {
	if registrosFiltradosSolicitud != nil {
		filtroFechaDeLanzamiento = registrosFiltradosSolicitud.filtroFechaDeLanzamiento
	}
	return
}

func (registrosFiltradosSolicitud *ObtenerRegistrosYVehiculosAsociadosFiltradosSolicitud) ObtenerFiltroTipoDeRegistro() (tipoDeRegistro string) {
	if registrosFiltradosSolicitud != nil {
		tipoDeRegistro = registrosFiltradosSolicitud.filtroTipoDeRegistro
	}
	return
}

func (registrosFiltradosSolicitud *ObtenerRegistrosYVehiculosAsociadosFiltradosSolicitud) AsignarFiltroPlaca(filtroPlaca string) {
	if registrosFiltradosSolicitud != nil {
		registrosFiltradosSolicitud.filtroPlaca = filtroPlaca
	}
}

func (registrosFiltradosSolicitud *ObtenerRegistrosYVehiculosAsociadosFiltradosSolicitud) AsignarFiltroMarca(filtroMarca string) {
	if registrosFiltradosSolicitud != nil {
		registrosFiltradosSolicitud.filtroMarca = filtroMarca
	}
}

func (registrosFiltradosSolicitud *ObtenerRegistrosYVehiculosAsociadosFiltradosSolicitud) AsignarFiltroModelo(filtroModelo string) {
	if registrosFiltradosSolicitud != nil {
		registrosFiltradosSolicitud.filtroModelo = filtroModelo
	}
}

func (registrosFiltradosSolicitud *ObtenerRegistrosYVehiculosAsociadosFiltradosSolicitud) AsignarFiltroFechaDeLanzamiento(filtroFechaDeLanzamiento int) {
	if registrosFiltradosSolicitud != nil {
		registrosFiltradosSolicitud.filtroFechaDeLanzamiento = filtroFechaDeLanzamiento
	}
}

func (registrosFiltradosSolicitud *ObtenerRegistrosYVehiculosAsociadosFiltradosSolicitud) AsignarFiltroTipoDeRegistro(filtroTipoDeRegistro string) {
	if registrosFiltradosSolicitud != nil {
		registrosFiltradosSolicitud.filtroTipoDeRegistro = filtroTipoDeRegistro
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

type EditarVehiculoSolicitud struct {
	PlacasActuales        string
	PlacasNuevas          string
	FechaLanzamientoNueva int
	MarcaNueva            string
	ModeloNuevo           string
	SerieNuevo            string
}

func (v *EditarVehiculoSolicitud) ObtenerPlacasActuales() (o string) {
	if v != nil {
		o = v.PlacasActuales
	}
	return
}

func (v *EditarVehiculoSolicitud) AsignarPlacasActuales(placasActuales string) {
	if v != nil {
		v.PlacasActuales = placasActuales
	}
}

func (v *EditarVehiculoSolicitud) ObtenerPlacasNuevas() (o string) {
	if v != nil {
		o = v.PlacasNuevas
	}
	return
}

func (v *EditarVehiculoSolicitud) AsignarPlacasNuevas(placasNuevas string) {
	if v != nil {
		v.PlacasNuevas = placasNuevas
	}
}

func (v *EditarVehiculoSolicitud) ObtenerFechaDeLanzamientoNueva() (o int) {
	if v != nil {
		o = v.FechaLanzamientoNueva
	}
	return
}

func (v *EditarVehiculoSolicitud) AsignarFechaDeLanzamientoNueva(fechaDeLanzamientoNueva int) {
	if v != nil {
		v.FechaLanzamientoNueva = fechaDeLanzamientoNueva
	}
}

func (v *EditarVehiculoSolicitud) ObtenerMarcaNueva() (o string) {
	if v != nil {
		o = v.MarcaNueva
	}
	return
}

func (v *EditarVehiculoSolicitud) AsignarMarcaNueva(marcaNueva string) {
	if v != nil {
		v.MarcaNueva = marcaNueva
	}
}

func (v *EditarVehiculoSolicitud) ObtenerModeloNuevo() (o string) {
	if v != nil {
		o = v.ModeloNuevo
	}
	return
}

func (v *EditarVehiculoSolicitud) AsignarModeloNuevo(modeloNuevo string) {
	if v != nil {
		v.ModeloNuevo = modeloNuevo
	}
}

func (v *EditarVehiculoSolicitud) ObtenerSerieNuevo() (o string) {
	if v != nil {
		o = v.SerieNuevo
	}
	return
}

func (v *EditarVehiculoSolicitud) AsignarSerieNuevo(serieNuevo string) {
	if v != nil {
		v.SerieNuevo = serieNuevo
	}
}
