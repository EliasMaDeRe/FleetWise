package controladores

import (
	"errors"
	"strconv"

	"example/fleetwise/fuente/iniciodesesion/constantes"
)

type Controlador struct {
	Usuarios map[string]Usuario // Mapa para almacenar usuarios
	Sesion   map[string]bool   // Mapa para gestionar sesiones de usuario
}

func NuevoInicioDeSesionControlador() *Controlador {
	return &Controlador{
		Usuarios: make(map[string]Usuario),
		Sesion:   make(map[string]bool),
	}
}

func (c *Controlador) IniciarSesion(c *gin.Context, solicitud modelos.InicioDeSesionSolicitud) {
	usuario, encontrado := c.Usuarios[solicitud.ClaveUsuario]

	if !encontrado || usuario.ClaveUsuario != solicitud.ClaveUsuario {
		c.JSON(http.StatusUnauthorized, gin.H{constantes.ERROR_CREDENCIALES_INVALIDAS})
		return
	}

	// Inicio de sesión
	c.Sesion[solicitud.ClaveUsuario] = true
	c.JSON(http.StatusOK, gin.H{"mensaje": "Inicio de sesión exitoso"})
}

func (c *Controlador) CerrarSesion(c *gin.Context, claveUsuario string) {
	_, sesionIniciada := c.Sesion[claveUsuario]
	if !sesionIniciada {
		c.JSON(http.StatusUnauthorized, gin.H{constantes.ERROR_SESION_NO_INICIADA})
		return
	}

	// Cerrar sesión
	delete(c.Sesion, claveUsuario)
	c.JSON(http.StatusOK, gin.H{"mensaje": "Sesión cerrada"})
}
