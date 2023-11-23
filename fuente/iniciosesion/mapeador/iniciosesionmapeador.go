package mapeador

import (
	inicioSesionModelos "example/fleetwise/modelos/iniciosesion"

	"github.com/gin-gonic/gin"
)

type Mapeador struct {
}

func (m *Mapeador) GinContextAIniciarSesionSolicitud(contexto *gin.Context) *inicioSesionModelos.IniciarSesionSolicitud {
	solicitud := &inicioSesionModelos.IniciarSesionSolicitud{}
	contexto.BindJSON(solicitud)
	return solicitud
}
