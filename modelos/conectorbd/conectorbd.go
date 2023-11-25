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

type ObtenerRegistrosConVehiculosFiltradosSolicitud struct{
	filtroPlaca string
	filtroMarca string
	filtroModelo string
	filtroFechaDeLanzamiento int
	filtroTipoDeRegistro string
}

func (registrosFiltradosSolicitud *ObtenerRegistrosConVehiculosFiltradosSolicitud) ObtenerFiltroPlaca() (filtroPlaca string){
	if registrosFiltradosSolicitud != nil{
		filtroPlaca = registrosFiltradosSolicitud.filtroPlaca
	}
	return
}

func (registrosFiltradosSolicitud *ObtenerRegistrosConVehiculosFiltradosSolicitud) ObtenerFiltroMarca() (filtroMarca string){
	if registrosFiltradosSolicitud != nil{
		filtroMarca = registrosFiltradosSolicitud.filtroMarca
	}
	return
}

func (registrosFiltradosSolicitud *ObtenerRegistrosConVehiculosFiltradosSolicitud) ObtenerFiltroModelo() (filtroModelo string){
	if registrosFiltradosSolicitud != nil{
		filtroModelo = registrosFiltradosSolicitud.filtroModelo
	}
	return
}

func (registrosFiltradosSolicitud *ObtenerRegistrosConVehiculosFiltradosSolicitud) ObtenerFiltroFechaDeLanzamiento() (filtroFechaDeLanzamiento int){
	if registrosFiltradosSolicitud != nil{
		filtroFechaDeLanzamiento = registrosFiltradosSolicitud.filtroFechaDeLanzamiento
	}
	return
}

func (registrosFiltradosSolicitud *ObtenerRegistrosConVehiculosFiltradosSolicitud) ObtenerFiltroTipoDeRegistro() (tipoDeRegistro string){
	if registrosFiltradosSolicitud != nil{
		tipoDeRegistro = registrosFiltradosSolicitud.filtroTipoDeRegistro
	}
	return
}


func (registrosFiltradosSolicitud *ObtenerRegistrosConVehiculosFiltradosSolicitud) AsignarFiltroPlaca(filtroPlaca string) {
	if registrosFiltradosSolicitud != nil{
		registrosFiltradosSolicitud.filtroPlaca = filtroPlaca 
	}
}

func (registrosFiltradosSolicitud *ObtenerRegistrosConVehiculosFiltradosSolicitud) AsignarFiltroMarca(filtroMarca string) {
	if registrosFiltradosSolicitud != nil{
		registrosFiltradosSolicitud.filtroMarca = filtroMarca 
	}
}

func (registrosFiltradosSolicitud *ObtenerRegistrosConVehiculosFiltradosSolicitud) AsignarFiltroModelo(filtroModelo string) {
	if registrosFiltradosSolicitud != nil{
		registrosFiltradosSolicitud.filtroModelo = filtroModelo 
	}
}

func (registrosFiltradosSolicitud *ObtenerRegistrosConVehiculosFiltradosSolicitud) AsignarFiltroFechaDeLanzamiento(filtroFechaDeLanzamiento int) {
	if registrosFiltradosSolicitud != nil{
		registrosFiltradosSolicitud.filtroFechaDeLanzamiento = filtroFechaDeLanzamiento 
	}
}
