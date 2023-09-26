package manejadores

import (
	"net/http"
	"tuaplicacion/controladores"
	"github.com/gin-gonic/gin"
)

type InicioDeSesionManejador struct {
	Controlador *controladores.InicioDeSesionControlador
}

func NuevoInicioDeSesionManejador(controlador *controladores.InicioDeSesionControlador) *InicioDeSesionManejador {
	return &InicioDeSesionManejador{Controlador: controlador}
}

func (h *InicioDeSesionManejador) IniciarSesion(c *gin.Context) {
	var solicitud modelos.InicioDeSesionSolicitud
	if err := c.ShouldBindJSON(&solicitud); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	h.Controlador.IniciarSesion(c, solicitud)
}

func (h *InicioDeSesionManejador) CerrarSesion(c *gin.Context) {
	claveUsuario := c.Param("claveUsuario")
	h.Controlador.CerrarSesion(c, claveUsuario)
}
