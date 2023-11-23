package manejadores

import (
	conectorBDControlador "example/fleetwise/fuente/conectorbd/controlador"
	inicioSesionControlador "example/fleetwise/fuente/iniciosesion/controlador"
	inicioSesionMapeador "example/fleetwise/fuente/iniciosesion/mapeador"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Manejador struct {
	InicioSesionControlador *inicioSesionControlador.Controlador
}

func NuevoManejador() (c *Manejador) {
	return &Manejador{
		InicioSesionControlador: &inicioSesionControlador.Controlador{
			InicioSesionMapeador:  &inicioSesionMapeador.Mapeador{},
			ConectorBDControlador: &conectorBDControlador.Controlador{},
		},
	}
}

func (m *Manejador) IniciarSesion(contexto *gin.Context) {
	solicitud := m.InicioSesionControlador.InicioSesionMapeador.GinContextAIniciarSesionSolicitud(contexto)
	tokenCadena, err := m.InicioSesionControlador.IniciarSesion(solicitud)
	status := http.StatusOK
	ok := true
	if tokenCadena == "" {
		status = http.StatusUnauthorized
		ok = false
	}
	var mensajeError string
	if err != nil {
		mensajeError = err.Error()
	}
	contexto.IndentedJSON(status, gin.H{"OK": ok, "err": mensajeError, "token": tokenCadena})
}
