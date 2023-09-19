package controlador

import (
	conectorConfig "example/fleetwise/ConectorBD/config"
	conectorModelos "example/fleetwise/modelos/conectorbd"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Controlador struct {
	flag bool
}

func GuardarVehiculo(contexto *gin.Context) {
	nuevoVehiculo := &conectorModelos.GuardarVehiculoSolicitud{}
	result := conectorConfig.DB.Create(nuevoVehiculo)
	if result.Error != nil {
		contexto.JSON(500, gin.H{"Error": "Failed to insert"})
		return
	}
	contexto.BindJSON(&nuevoVehiculo)
	contexto.IndentedJSON(http.StatusCreated, &nuevoVehiculo)
}

func (c *Controlador) RespuestaGuardarVehiculo() {
	//tampoco supe que ponerle pipipi
}

func (c *Controlador) ObtenerVehiculo(solicitud *conectorModelos.ObtenerCargoSoicitud) *conectorModelos.ObtenerCargoSoicitud {

	respuesta := &conectorModelos.ObtenerCargoSoicitud{}
	return respuesta
}
