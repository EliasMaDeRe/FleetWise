package visualizacionresumenflotilla

type ExportarResumenFlotillaSolicitud struct {
	Formato Formato `json:"formato"`
}

func (solicitud *ExportarResumenFlotillaSolicitud) ObtenerFormato() (formato Formato) {
	if solicitud != nil {
		formato = solicitud.Formato
	}
	return
}

func (solicitud *ExportarResumenFlotillaSolicitud) AsignarFormato(formato Formato) {
	if solicitud != nil {
		solicitud.Formato = formato
	}
}

type Formato string

const (
	FormatoExcel Formato = "excel"
)
