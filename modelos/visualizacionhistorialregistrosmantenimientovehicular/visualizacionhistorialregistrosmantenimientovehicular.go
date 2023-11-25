package visualizacionhistorialregistrosmantenimientovehicular

type ObtenerRegistrosYVehiculosFiltradosSolicitud struct{
	Placas string `json:"placas"`
	Marca string `json:"marca"`
	Modelo string `json:"modelo"`
	FechaLanzamiento string `json:"fechaDeLanzamiento"`
	TipoDeRegistro string `json:"tipoDeRegistro"`
}

func (registrosSolicitud *ObtenerRegistrosYVehiculosFiltradosSolicitud) ObtenerPlacas() (placa string){
	if(registrosSolicitud != nil){
		placa = registrosSolicitud.Placas;
	}
	return
}

func (registrosSolicitud *ObtenerRegistrosYVehiculosFiltradosSolicitud) ObtenerMarca() (marca string){
	if(registrosSolicitud != nil){
		marca = registrosSolicitud.Marca;
	}
	return
}

func (registrosSolicitud *ObtenerRegistrosYVehiculosFiltradosSolicitud) ObtenerModelo() (modelo string){
	if(registrosSolicitud != nil){
		modelo = registrosSolicitud.Modelo;
	}
	return
}

func (registrosSolicitud *ObtenerRegistrosYVehiculosFiltradosSolicitud) ObtenerFechaDeLanzamiento() (fechaDeLanzamiento string){
	if(registrosSolicitud != nil){
		fechaDeLanzamiento = registrosSolicitud.FechaLanzamiento;
	}
	return
}

func (registrosSolicitud *ObtenerRegistrosYVehiculosFiltradosSolicitud) ObtenerTipoDeRegistro() (tipoDeRegistro string){
	if(registrosSolicitud != nil){
		tipoDeRegistro = registrosSolicitud.TipoDeRegistro;
	}
	return
}


type ObtenerRegistrosYVehiculosFiltradosRespuesta struct {
	ok bool
	err error
}

func(registrosRespuesta *ObtenerRegistrosYVehiculosFiltradosRespuesta) ObtenerOk() (ok bool){
	if registrosRespuesta != nil {
		ok = registrosRespuesta.ok
	}
	return
}

func(registrosRespuesta *ObtenerRegistrosYVehiculosFiltradosRespuesta) ObtenerError() (err error){
	if(registrosRespuesta != nil){
		err = registrosRespuesta.err
	}
	return
}


