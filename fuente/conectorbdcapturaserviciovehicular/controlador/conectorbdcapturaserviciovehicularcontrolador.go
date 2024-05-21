package controlador

import (
	"errors"
	"example/fleetwise/fuente/conectorbdcapturaserviciovehicular/constantes"
	capturaServicioVehicularModelos "example/fleetwise/modelos/capturaserviciovehicular"

	conectorBDModelos "example/fleetwise/modelos/conectorbdcapturaserviciovehicular"
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
)

type Controlador struct {
}

func (c *Controlador) obtenerConexionABd() string {
	dbHost := os.Getenv("dbHost")
	dbNombre := os.Getenv("dbNombre")
	dbUsuario := os.Getenv("dbUsuario")
	dbContrasena := os.Getenv("dbContrasena")
	dbPuerto := os.Getenv("dbPuerto")
	conexionDb := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUsuario, dbContrasena, dbHost, dbPuerto, dbNombre)
	return conexionDb
}


func (c *Controlador) ObtenerServicioVehicularPorNombre(solicitud *conectorBDModelos.ObtenerServicioVehicularPorNombreSolicitud) *capturaServicioVehicularModelos.ServicioVehicular {

	baseDeDatos, errConectarBD := gorm.Open("mysql", c.obtenerConexionABd())

	if errConectarBD != nil {
		log.Fatal(constantes.ERROR_CONECTAR_BD)
	}

	serviciosVehiculares := []capturaServicioVehicularModelos.ServicioVehicular{}

	resultadoBusqueda := baseDeDatos.Where(&capturaServicioVehicularModelos.ServicioVehicular{Nombre: solicitud.ObtenerNombre()}).Find(&serviciosVehiculares)

	if resultadoBusqueda.Error != nil {
		log.Fatal(constantes.ERROR_BUSQUEDA_EN_BD)
	}

	var servicioVehicularEncontrado *capturaServicioVehicularModelos.ServicioVehicular

	if len(serviciosVehiculares) != 0 {
		servicioVehicularEncontrado = &serviciosVehiculares[0]
	}

	return servicioVehicularEncontrado
}

func (c *Controlador) ObtenerServiciosVehiculares() []capturaServicioVehicularModelos.ServicioVehicular {

	baseDeDatos, errConectarBD := gorm.Open("mysql", c.obtenerConexionABd())

	if errConectarBD != nil {
		log.Fatal(constantes.ERROR_CONECTAR_BD)
	}

	serviciosVehiculares := []capturaServicioVehicularModelos.ServicioVehicular{}

	resultadoBusqueda := baseDeDatos.Find(&serviciosVehiculares)

	if resultadoBusqueda.Error != nil {
		log.Fatal(constantes.ERROR_BUSQUEDA_EN_BD)
	}

	return serviciosVehiculares
}

func (c *Controlador) GuardarServicioVehicular(servicioVehicular *capturaServicioVehicularModelos.ServicioVehicular) *conectorBDModelos.GuardarServicioVehicularRespuesta {

	baseDeDatos, errConectarBD := gorm.Open("mysql", c.obtenerConexionABd())

	if errConectarBD != nil {
		log.Fatal(constantes.ERROR_CONECTAR_BD)
	}

	baseDeDatos.AutoMigrate(&servicioVehicular)

	resultadoGuardarServicioVehicular := baseDeDatos.Create(
		&capturaServicioVehicularModelos.ServicioVehicular{
			Nombre: servicioVehicular.ObtenerNombre(),
		})
	respuesta := &conectorBDModelos.GuardarServicioVehicularRespuesta{}
	if resultadoGuardarServicioVehicular.Error != nil {
		respuesta.AsignarOk(false)
		respuesta.AsignarErr(errors.New(constantes.ERROR_GUARDAR_FILA_BD))
		return respuesta
	}
	respuesta.AsignarOk(true)
	respuesta.AsignarErr(nil)

	return respuesta
}


func (c *Controlador) EditarServicioVehicular(solicitud *conectorBDModelos.EditarServicioVehicularSolicitud) error {

	baseDeDatos, errConectarBD := gorm.Open("mysql", c.obtenerConexionABd())

	servicioVehicularNuevo := &capturaServicioVehicularModelos.ServicioVehicular{}
	servicioVehicularNuevo.AsignarNombre(solicitud.ObtenerNuevoNombre())

	if errConectarBD != nil {
		log.Fatal(constantes.ERROR_CONECTAR_BD)
	}

	baseDeDatos.Model(servicioVehicularNuevo).Where("Nombre = ?", solicitud.ObtenerNombreActual()).Update("nombre", servicioVehicularNuevo.ObtenerNombre())

	return nil
}

func (c *Controlador) EliminarServicioVehicular(servicioVehicular *capturaServicioVehicularModelos.ServicioVehicular) {

	baseDeDatos, errConectarBD := gorm.Open("mysql", c.obtenerConexionABd())

	baseDeDatos.AutoMigrate(&servicioVehicular)

	if errConectarBD != nil {
		log.Fatal(constantes.ERROR_CONECTAR_BD)
	}

	baseDeDatos.Where("Nombre = ?", servicioVehicular.ObtenerNombre()).Delete(servicioVehicular)

}