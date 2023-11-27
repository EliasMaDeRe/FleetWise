package manejador

import (
	registroMantenimientoVehiculoControlador "example/fleetwise/fuente/capturaregistromantenimientovehiculo/controlador"
	registroMantenimientoVehiculoMapeador "example/fleetwise/fuente/capturaregistromantenimientovehiculo/mapeador"
	servicioVehicularControlador "example/fleetwise/fuente/capturaserviciovehicular/controlador"
	"fmt"
	"strings"

	"net/http"

	"github.com/gin-gonic/gin"
)

type Manejador struct {
	RegistroMantenimientoVehiculoControlador *registroMantenimientoVehiculoControlador.Controlador
}

func NuevoManejador() (c *Manejador) {
	return &Manejador{
		RegistroMantenimientoVehiculoControlador: &registroMantenimientoVehiculoControlador.Controlador{
			RegistroMantenimientoVehiculoMapeador: &registroMantenimientoVehiculoMapeador.Mapeador{},
			ServicioVehicularControlador:          &servicioVehicularControlador.Controlador{},
		},
	}
}

func (m *Manejador) SeleccionarVehiculoParaNuevoRegistro(contexto *gin.Context) {
	rawJSON, err := contexto.GetRawData()
	if err != nil {
		contexto.JSON(500, gin.H{"error": "Error al obtener los datos JSON"})
		return
	}
	fmt.Println("Datos JSON sin procesar SeleccionarVehiculo:", string(rawJSON))
	solicitud := m.RegistroMantenimientoVehiculoControlador.RegistroMantenimientoVehiculoMapeador.GinContextASeleccionarVehiculoParaNuevoRegistroSolicitud(contexto)
	respuesta := m.RegistroMantenimientoVehiculoControlador.SeleccionarVehiculoParaNuevoRegistro(solicitud)
	status := http.StatusOK
	if !respuesta.ObtenerOk() {
		status = http.StatusBadRequest
	}
	if strings.Contains(contexto.Request.URL.String(), "AgregarRegistroMantenimientoVehicular") {
		if respuesta.ObtenerOk() {
			contexto.Next()
			return
		} else {
			contexto.HTML(http.StatusOK, "seleccionarvehiculo.html", gin.H{})
			contexto.Abort()
			return
		}
	}
	var mensajeError string
	if respuesta.ObtenerErr() != nil {
		mensajeError = respuesta.ObtenerErr().Error()
	}
	contexto.IndentedJSON(status, gin.H{"OK": respuesta.ObtenerOk(), "err": mensajeError})
}

func (m *Manejador) AgregarRegistroMantemientoVehiculo(contexto *gin.Context) {
	solicitud := m.RegistroMantenimientoVehiculoControlador.RegistroMantenimientoVehiculoMapeador.GinContextAAgregarRegistroMantemientoVehiculoSolicitud(contexto)
	fmt.Println(solicitud)
	respuesta := m.RegistroMantenimientoVehiculoControlador.AgregarRegistroMantemientoVehiculo(solicitud)
	status := http.StatusOK
	if !respuesta.ObtenerOk() {
		status = http.StatusBadRequest
	}
	var mensajeError string
	if respuesta.ObtenerErr() != nil {
		mensajeError = respuesta.ObtenerErr().Error()
	}
	contexto.IndentedJSON(status, gin.H{"OK": respuesta.ObtenerOk(), "err": mensajeError})
}

func (m *Manejador) ObtenerServiciosVehiculares(contexto *gin.Context) {
	respuesta := m.RegistroMantenimientoVehiculoControlador.ObtenerServiciosVehicularesParaNuevoRegistro()
	contexto.IndentedJSON(http.StatusOK, gin.H{"ServiciosVehiculares": respuesta})
}
