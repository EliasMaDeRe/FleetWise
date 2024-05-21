package conectorbdvisualizacionhistorialregistrosmantenimientovehiculo


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
