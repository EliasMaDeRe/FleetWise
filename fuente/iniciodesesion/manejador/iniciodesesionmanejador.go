package manejadores

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"example/fleetwise/fuente/iniciodesesion/controladores"
	"example/fleetwise/modelos/iniciodesesion"
)

type Manejador struct {
	Controlador *controladores.Controlador
}

func NuevoInicioDeSesionManejador(controlador *controladores.Controlador) *Manejador {
	return &Manejador{Controlador: controlador}
}

func (h *Manejador) IniciarSesion(contexto *gin.Context) {
	var solicitud modelos.InicioDeSesionSolicitud
	if err := contexto.ShouldBindJSON(&solicitud); err != nil {
		// Si la solicitud es incorrecta, asigna un mensaje de error a la respuesta
		respuesta := modelos.NuevaInicioDeSesionRespuesta()
		respuesta.AsignarError(err)
		contexto.JSON(http.StatusBadRequest, respuesta)
		return
	}

	// Llamar al controlador y obtener una respuesta
	respuesta := h.Controlador.IniciarSesion(contexto, solicitud)

	// Enviar la respuesta al cliente
	status := http.StatusOK
	if !respuesta.Ok {
		status = http.StatusUnauthorized
	}
	contexto.JSON(status, respuesta)
}

func (h *Manejador) CerrarSesion(contexto *gin.Context) {
	claveUsuario := contexto.Param("claveUsuario")

	// Llamar al controlador y obtener una respuesta
	respuesta := h.Controlador.CerrarSesion(contexto, claveUsuario)

	// Enviar la respuesta al cliente
	status := http.StatusOK
	if !respuesta.Ok {
		status = http.StatusUnauthorized
	}
	contexto.JSON(status, respuesta)
}
