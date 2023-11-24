package mapeador

import (
	inicioSesionModelos "example/fleetwise/modelos/iniciosesion"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type Mapeador struct {
}

func (m *Mapeador) GinContextAIniciarSesionSolicitud(contexto *gin.Context) *inicioSesionModelos.IniciarSesionSolicitud {
	solicitud := &inicioSesionModelos.IniciarSesionSolicitud{}
	contexto.BindJSON(solicitud)
	return solicitud
}

func (m *Mapeador) TokenClaimsAObtenerUsuarioPorNombreUsuarioSolicitud(tokenClaims jwt.MapClaims) *inicioSesionModelos.ObtenerUsuarioPorNombreUsuario {
	solicitud := &inicioSesionModelos.ObtenerUsuarioPorNombreUsuario{}
	solicitud.AsignarNombreUsuario(tokenClaims["nombreUsuario"].(string))
	return solicitud
}
