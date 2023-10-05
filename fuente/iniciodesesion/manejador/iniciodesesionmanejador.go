package manejadores

import (
	"errors"
	InicioDeSesionControlador "example/fleetwise/fuente/iniciodesesion/controlador"
	"example/fleetwise/modelos/iniciodesesion"
)

type Manejador struct {
	Controlador *controladores.InicioDeSesionControlador
}

func NuevoInicioDeSesionManejador(controlador *controladores.InicioDeSesionControlador) *Manejador {
	return &Manejador{Controlador: controlador}
}

var err error
func (h *Manejador) IniciarSesion(contexto *gin.Context) {
	var solicitud := modelos.InicioDeSesionSolicitud
	if err := c.ShouldBindJSON(&solicitud); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{err.Error()})
		return
	}

	h.InicioDeSesionControlador.IniciarSesion(c, solicitud)
}

func (h *Manejador) CerrarSesion(c *gin.Context) {
	claveUsuario := c.Param("claveUsuario")
	h.InicioDeSesionControlador.CerrarSesion(c, claveUsuario)
}
