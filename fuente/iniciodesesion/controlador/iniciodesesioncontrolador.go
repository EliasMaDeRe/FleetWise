package controladores

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"example/fleetwise/fuente/iniciodesesion/constantes"
	iniciodesesionmodelos "example/fleetwise/modelos/iniciodesesion"
)

type Controlador struct {
	Usuarios map[string]iniciodesesionmodelos.Usuario // Mapa para almacenar usuarios
	Sesion   map[string]bool                           // Mapa para gestionar sesiones de usuario
}

func NuevoInicioDeSesionControlador() *Controlador {
	return &Controlador{
		Usuarios: make(map[string]iniciodesesionmodelos.Usuario),
		Sesion:   make(map[string]bool),
	}
}

func (c *Controlador) IniciarSesion(ctx *gin.Context, solicitud iniciodesesionmodelos.InicioDeSesionSolicitud) *iniciodesesionmodelos.InicioDeSesionRespuesta {
	respuesta := iniciodesesionmodelos.NuevaInicioDeSesionRespuesta() // Crear una nueva respuesta predeterminada

	usuario, encontrado := c.Usuarios[solicitud.ClaveUsuario]

	if !encontrado || usuario.ClaveUsuario != solicitud.ClaveUsuario {
		// En caso de error, asigna un mensaje de error a la respuesta
		respuesta.AsignarError(errors.New(constantes.ERROR_CREDENCIALES_INVALIDAS))
		return respuesta
	}

	// Enviar la respuesta al manejador
	return respuesta
}

func (c *Controlador) CerrarSesion(ctx *gin.Context, claveUsuario string) *iniciodesesionmodelos.InicioDeSesionRespuesta {
	respuesta := iniciodesesionmodelos.NuevaInicioDeSesionRespuesta() // Crear una nueva respuesta predeterminada

	sesionIniciada, encontrado := c.Sesion[claveUsuario]
	if !encontrado || !sesionIniciada {
		// En caso de error, asigna un mensaje de error a la respuesta
		respuesta.AsignarError(errors.New(constantes.ERROR_SESION_NO_INICIADA))
		return respuesta
	}

	// Cerrar sesión
	delete(c.Sesion, claveUsuario)
	return respuesta
}
