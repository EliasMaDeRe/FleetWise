package mapeador

import (
	modelosResumen "example/fleetwise/modelos/visualizacionresumenflotilla"

	"github.com/gin-gonic/gin"
)

type Mapeador struct {
}

func (m *Mapeador) GinContextAExportarResumenFlotillaSolicitud(contexto *gin.Context) *modelosResumen.ExportarResumenFlotillaSolicitud {
	solicitud := &modelosResumen.ExportarResumenFlotillaSolicitud{}
	solicitud.AsignarFormato(modelosResumen.Formato(contexto.Request.URL.Query()["Formato"][0]))
	return solicitud
}
