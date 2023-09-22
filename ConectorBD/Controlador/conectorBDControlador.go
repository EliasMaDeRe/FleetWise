package controlador

import (
	conectorModelos "example/fleetwise/modelos/conectorbd"
	vehiculosModelos "example/fleetwise/modelos/vehiculos"

	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB

type Controlador struct {
}

func GuardarVehiculo() {
	var err error

	database, err := gorm.Open("mysql", "root:@tcp(127.0.0.1:3307)/gestor_de_flotillas")

	if err != nil {
		log.Fatal("Failed to connect DB")
	}

	VehiculoSolicitud := &conectorModelos.GuardarVehiculoSolicitud{}

	nuevoVehiculo := vehiculosModelos.Vehiculo{}

	nuevoVehiculo.AsignarAnualidad(1)
	nuevoVehiculo.AsignarID("14")
	nuevoVehiculo.AsignarMarca("a")
	nuevoVehiculo.AsignarModelo("b")
	nuevoVehiculo.AsignarPlacas("c")
	nuevoVehiculo.AsignarSerie("d")

	VehiculoSolicitud.AsignarVehiculo(nuevoVehiculo)

	database.Create(VehiculoSolicitud.ObtenerVehiculo())

	DB = database
}

func (c *Controlador) RespuestaGuardarVehiculo() {
	//tampoco supe que ponerle pipipi
}

func (c *Controlador) ObtenerVehiculo(solicitud *conectorModelos.ObtenerCargoSoicitud) *conectorModelos.ObtenerCargoSoicitud {

	respuesta := &conectorModelos.ObtenerCargoSoicitud{}
	return respuesta
}
