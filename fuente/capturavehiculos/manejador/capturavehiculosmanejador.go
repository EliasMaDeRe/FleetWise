package manejador

import (
	capturaVehiculosControlador "example/fleetwise/fuente/capturavehiculos/controlador"
	capturaVehiculosMapeador "example/fleetwise/fuente/capturavehiculos/mapeador"
	conectorBDControlador "example/fleetwise/fuente/conectorbd/controlador"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Manejador struct {
	CapturaVehiculosControlador *capturaVehiculosControlador.Controlador
}

func NuevoManejador() (c *Manejador) {
	return &Manejador{
		CapturaVehiculosControlador: &capturaVehiculosControlador.Controlador{
			CapturaVehiculosMapeador: &capturaVehiculosMapeador.Mapeador{},
			ConectorBDControlador:    &conectorBDControlador.Controlador{},
		},
	}
}

func (m *Manejador) AgregarVehiculo(contexto *gin.Context) {
	solicitud := m.CapturaVehiculosControlador.CapturaVehiculosMapeador.GinContextAAgregarVehiculoSolicitud(contexto)
	respuesta := m.CapturaVehiculosControlador.AgregarVehiculo(solicitud)
	status := http.StatusOK
	if respuesta.ObtenerOk() == false {
		status = http.StatusBadRequest
	}
	var mensajeError string
	if respuesta.ObtenerErr() != nil {
		mensajeError = respuesta.ObtenerErr().Error()
	}
	contexto.IndentedJSON(status, gin.H{"OK": respuesta.ObtenerOk(), "err": mensajeError})
}

func (m *Manejador) AgregarImagenVehiculo(contexto *gin.Context) {
	archivo, err := contexto.FormFile("archivo")

	if err != nil {
		contexto.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"err": "No se recupero el archivo",
		})
		return
	}

	extension := filepath.Ext(archivo.Filename)
	nuevoNombreArchivo := uuid.New().String() + extension

	if err := contexto.SaveUploadedFile(archivo, "./img/uploads/"+nuevoNombreArchivo); err != nil {
		contexto.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"err": "No se guardo el archivo",
		})
		return
	}

	contexto.JSON(http.StatusOK, gin.H{
		"OK": true,
	})
}

func (m *Manejador) EditarVehiculo(contexto *gin.Context) {
	solicitud := m.CapturaVehiculosControlador.CapturaVehiculosMapeador.GinContextAEditarVehiculoSolicitud(contexto)
	respuesta := m.CapturaVehiculosControlador.EditarVehiculo(solicitud)
	status := http.StatusOK
	if respuesta.ObtenerOk() == false {
		status = http.StatusBadRequest
	}
	var mensajeError string
	if respuesta.ObtenerErr() != nil {
		mensajeError = respuesta.ObtenerErr().Error()
	}
	contexto.IndentedJSON(status, gin.H{"OK": respuesta.ObtenerOk(), "err": mensajeError})
}

func (m *Manejador) ObtenerVehiculoPorPlacas(contexto *gin.Context) {
	solicitud := m.CapturaVehiculosControlador.CapturaVehiculosMapeador.GinContextAObtenerVehiculoPorPlacasSolicitud(contexto)
	respuesta := m.CapturaVehiculosControlador.ObtenerVehiculoPorPlacas(solicitud)
	contexto.IndentedJSON(http.StatusOK, gin.H{"Vehiculo": respuesta})
}

func (m *Manejador) EliminarVehiculo(contexto *gin.Context) {
	solicitud := m.CapturaVehiculosControlador.CapturaVehiculosMapeador.GinContextAEliminarVehiculoSolicitud(contexto)
	respuesta := m.CapturaVehiculosControlador.EliminarVehiculo(solicitud)
	status := http.StatusOK
	if respuesta.ObtenerOk() == false {
		status = http.StatusBadRequest
	}
	var mensajeError string
	if respuesta.ObtenerErr() != nil {
		mensajeError = respuesta.ObtenerErr().Error()
	}
	contexto.IndentedJSON(status, gin.H{"OK": respuesta.ObtenerOk(), "err": mensajeError})
}
