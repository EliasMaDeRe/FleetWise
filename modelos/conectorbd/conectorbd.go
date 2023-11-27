package conectorbd

import (
	vehiculosModelos "example/fleetwise/modelos/capturavehiculos"
)

type GuardarVehiculoSolicitud struct {
	vehiculo vehiculosModelos.Vehiculo
}

func (v *GuardarVehiculoSolicitud) ObtenerVehiculo() (o vehiculosModelos.Vehiculo) {
	if v != nil {
		o = v.vehiculo
	}
	return
}

func (v *GuardarVehiculoSolicitud) AsignarVehiculo(vehiculo vehiculosModelos.Vehiculo) {
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

type ObtenerRegistrosYVehiculosAsociadosFiltradosSolicitud struct{
	filtroPlaca string
	filtroMarca string
	filtroModelo string
	filtroFechaDeLanzamiento int
	filtroTipoDeRegistro string
}

func (registrosFiltradosSolicitud *ObtenerRegistrosYVehiculosAsociadosFiltradosSolicitud) ObtenerFiltroPlaca() (filtroPlaca string){
	if registrosFiltradosSolicitud != nil{
		filtroPlaca = registrosFiltradosSolicitud.filtroPlaca
	}
	return
}

func (registrosFiltradosSolicitud *ObtenerRegistrosYVehiculosAsociadosFiltradosSolicitud) ObtenerFiltroMarca() (filtroMarca string){
	if registrosFiltradosSolicitud != nil{
		filtroMarca = registrosFiltradosSolicitud.filtroMarca
	}
	return
}

func (registrosFiltradosSolicitud *ObtenerRegistrosYVehiculosAsociadosFiltradosSolicitud) ObtenerFiltroModelo() (filtroModelo string){
	if registrosFiltradosSolicitud != nil{
		filtroModelo = registrosFiltradosSolicitud.filtroModelo
	}
	return
}

func (registrosFiltradosSolicitud *ObtenerRegistrosYVehiculosAsociadosFiltradosSolicitud) ObtenerFiltroFechaDeLanzamiento() (filtroFechaDeLanzamiento int){
	if registrosFiltradosSolicitud != nil{
		filtroFechaDeLanzamiento = registrosFiltradosSolicitud.filtroFechaDeLanzamiento
	}
	return
}

func (registrosFiltradosSolicitud *ObtenerRegistrosYVehiculosAsociadosFiltradosSolicitud) ObtenerFiltroTipoDeRegistro() (tipoDeRegistro string){
	if registrosFiltradosSolicitud != nil{
		tipoDeRegistro = registrosFiltradosSolicitud.filtroTipoDeRegistro
	}
	return
}

func (registrosFiltradosSolicitud *ObtenerRegistrosYVehiculosAsociadosFiltradosSolicitud) AsignarFiltroPlaca(filtroPlaca string) {
	if registrosFiltradosSolicitud != nil{
		registrosFiltradosSolicitud.filtroPlaca = filtroPlaca 
	}
}

func (registrosFiltradosSolicitud *ObtenerRegistrosYVehiculosAsociadosFiltradosSolicitud) AsignarFiltroMarca(filtroMarca string) {
	if registrosFiltradosSolicitud != nil{
		registrosFiltradosSolicitud.filtroMarca = filtroMarca 
	}
}

func (registrosFiltradosSolicitud *ObtenerRegistrosYVehiculosAsociadosFiltradosSolicitud) AsignarFiltroModelo(filtroModelo string) {
	if registrosFiltradosSolicitud != nil{
		registrosFiltradosSolicitud.filtroModelo = filtroModelo 
	}
}

func (registrosFiltradosSolicitud *ObtenerRegistrosYVehiculosAsociadosFiltradosSolicitud) AsignarFiltroFechaDeLanzamiento(filtroFechaDeLanzamiento int) {
	if registrosFiltradosSolicitud != nil{
		registrosFiltradosSolicitud.filtroFechaDeLanzamiento = filtroFechaDeLanzamiento 
	}
}


type ActualizarRegistroMantenimientoVehicularSolicitud struct {
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


func (actualizarRegistroSolicitud *ActualizarRegistroMantenimientoVehicularSolicitud) ObtenerNumeroDeRegistro() (numeroDeRegistro int) {
	if actualizarRegistroSolicitud != nil {
		numeroDeRegistro = actualizarRegistroSolicitud.NumeroDeRegistro
	}
	return
}

func (actualizarRegistroSolicitud *ActualizarRegistroMantenimientoVehicularSolicitud) AsignarNumeroDeRegistro(NumeroDeRegistro int) {
	if actualizarRegistroSolicitud != nil {
		actualizarRegistroSolicitud.NumeroDeRegistro = NumeroDeRegistro
	}
}

func (actualizarRegistroSolicitud *ActualizarRegistroMantenimientoVehicularSolicitud) ObtenerTipo() (tipo string) {
	if actualizarRegistroSolicitud != nil {
		tipo = actualizarRegistroSolicitud.Tipo
	}
	return
}

func (actualizarRegistroSolicitud *ActualizarRegistroMantenimientoVehicularSolicitud) AsignarTipo(tipo string) {
	if actualizarRegistroSolicitud != nil {
		actualizarRegistroSolicitud.Tipo = tipo
	}
}

func (actualizarRegistroSolicitud *ActualizarRegistroMantenimientoVehicularSolicitud) ObtenerFecha() (fecha string) {
	if actualizarRegistroSolicitud != nil {
		fecha = actualizarRegistroSolicitud.Fecha
	}
	return
}

func (actualizarRegistroSolicitud *ActualizarRegistroMantenimientoVehicularSolicitud) AsignarFecha(fecha string) {
	if actualizarRegistroSolicitud != nil {
		actualizarRegistroSolicitud.Fecha = fecha
	}
}

func (actualizarRegistroSolicitud *ActualizarRegistroMantenimientoVehicularSolicitud) ObtenerLitrosDeGasolina() (litrosGasolina float64) {
	if actualizarRegistroSolicitud != nil {
		litrosGasolina = actualizarRegistroSolicitud.LitrosDeGasolina
	}
	return
}

func (actualizarRegistroSolicitud *ActualizarRegistroMantenimientoVehicularSolicitud) AsignarLitrosDeGasolina(LitrosDeGasolina float64) {
	if actualizarRegistroSolicitud != nil {
		actualizarRegistroSolicitud.LitrosDeGasolina = LitrosDeGasolina
	}
}

func (actualizarRegistroSolicitud *ActualizarRegistroMantenimientoVehicularSolicitud) ObtenerKilometraje() (kilometraje int) {
	if actualizarRegistroSolicitud != nil {
		kilometraje = actualizarRegistroSolicitud.Kilometraje
	}
	return
}

func (actualizarRegistroSolicitud *ActualizarRegistroMantenimientoVehicularSolicitud) AsignarKilometraje(kilometraje int) {
	if actualizarRegistroSolicitud != nil {
		actualizarRegistroSolicitud.Kilometraje = kilometraje
	}
}

func (actualizarRegistroSolicitud *ActualizarRegistroMantenimientoVehicularSolicitud) ObtenerImporte() (importe float64) {
	if actualizarRegistroSolicitud != nil {
		importe = actualizarRegistroSolicitud.Importe
	}
	return
}

func (actualizarRegistroSolicitud *ActualizarRegistroMantenimientoVehicularSolicitud) AsignarImporte(importe float64) {
	if actualizarRegistroSolicitud != nil {
		actualizarRegistroSolicitud.Importe = importe
	}
}

func (actualizarRegistroSolicitud *ActualizarRegistroMantenimientoVehicularSolicitud) ObtenerObservaciones() (observaciones string) {
	if actualizarRegistroSolicitud != nil {
		observaciones = actualizarRegistroSolicitud.Observaciones
	}
	return
}

func (actualizarRegistroSolicitud *ActualizarRegistroMantenimientoVehicularSolicitud) AsignarObservaciones(observaciones string) {
	if actualizarRegistroSolicitud != nil {
		actualizarRegistroSolicitud.Observaciones = observaciones
	}
}

func (actualizarRegistroSolicitud *ActualizarRegistroMantenimientoVehicularSolicitud) ObtenerConcepto() (concepto string) {
	if actualizarRegistroSolicitud != nil {
		concepto = actualizarRegistroSolicitud.Concepto
	}
	return
}

func (actualizarRegistroSolicitud *ActualizarRegistroMantenimientoVehicularSolicitud) AsignarConcepto(concepto string) {
	if actualizarRegistroSolicitud != nil {
		actualizarRegistroSolicitud.Concepto = concepto
	}
}

func (actualizarRegistroSolicitud *ActualizarRegistroMantenimientoVehicularSolicitud) ObtenerPlacasVehiculo() (placasVehiculo string) {
	if actualizarRegistroSolicitud != nil {
		placasVehiculo = actualizarRegistroSolicitud.PlacasVehiculo
	}
	return
}

func (actualizarRegistroSolicitud *ActualizarRegistroMantenimientoVehicularSolicitud) AsignarPlacasVehiculo(placasVehiculo string) {
	if actualizarRegistroSolicitud != nil {
		actualizarRegistroSolicitud.PlacasVehiculo = placasVehiculo
	}
}

type ActualizarRegistroMantenimientoVehicularRespuesta struct{
	Ok bool;
	Err error;
}


func (editarRegistroRespuesta *ActualizarRegistroMantenimientoVehicularRespuesta) ObtenerOk() (ok bool){
	if(editarRegistroRespuesta != nil){
		ok = editarRegistroRespuesta.Ok
	}
	return
}

func (editarRegistroRespuesta *ActualizarRegistroMantenimientoVehicularRespuesta) AsignarOk(ok bool){
	if(editarRegistroRespuesta != nil){
		editarRegistroRespuesta.Ok = ok
	}
}

func (editarRegistroRespuesta *ActualizarRegistroMantenimientoVehicularRespuesta) ObtenerError() (err error){
	if(editarRegistroRespuesta !=nil){
		err = editarRegistroRespuesta.Err
	}
	return
}

func (editarRegistroRespuesta *ActualizarRegistroMantenimientoVehicularRespuesta) AsignarError(err error){
	if(editarRegistroRespuesta != nil){
		editarRegistroRespuesta.Err = err
	}
}
