package controlador

import (
	"errors"
	"example/fleetwise/fuente/conectorbd/constantes"
	servicioVehicularModelos "example/fleetwise/modelos/capturaserviciovehicular"
	vehiculosModelos "example/fleetwise/modelos/capturavehiculos"
	conectorModelos "example/fleetwise/modelos/conectorbd"
	sesionModelos "example/fleetwise/modelos/iniciosesion"
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Controlador struct {
}

func (c *Controlador) ObtenerVehiculoPorPlacas(solicitud *conectorModelos.ObtenerVehiculoPorPlacasSolicitud) *vehiculosModelos.Vehiculo {

	database, errConectarBD := gorm.Open("mysql", c.ConexionBd())

	if errConectarBD != nil {
		log.Fatal(constantes.ERROR_CONECTAR_BD)
	}

	vehiculos := []vehiculosModelos.Vehiculo{}

	resultadoBusqueda := database.Where(&vehiculosModelos.Vehiculo{Placas: solicitud.ObtenerPlacas()}).Find(&vehiculos)

	if resultadoBusqueda.Error != nil {
		log.Fatal(constantes.ERROR_PLACAS_INEXISTENTES_EN_BD)
	}

	var vehiculoEncontrado *vehiculosModelos.Vehiculo

	if len(vehiculos) != 0 {
		vehiculoEncontrado = &vehiculos[0]
	}

	return vehiculoEncontrado
}

func (c *Controlador) ObtenerVehiculoPorSerie(solicitud *conectorModelos.ObtenerVehiculoPorSerieSolicitud) *vehiculosModelos.Vehiculo {

	database, errConectarBD := gorm.Open("mysql", c.ConexionBd())

	if errConectarBD != nil {
		log.Fatal(constantes.ERROR_CONECTAR_BD)
	}

	vehiculos := []vehiculosModelos.Vehiculo{}

	resultadoBusqueda := database.Where(&vehiculosModelos.Vehiculo{Serie: solicitud.ObtenerSerie()}).Find(&vehiculos)

	if resultadoBusqueda.Error != nil {
		log.Fatal(constantes.ERROR_BUSQUEDA_EN_BD)
	}

	var vehiculoEncontrado *vehiculosModelos.Vehiculo

	if len(vehiculos) != 0 {
		vehiculoEncontrado = &vehiculos[0]
	}

	return vehiculoEncontrado
}

func (c *Controlador) GuardarVehiculo(vehiculo *vehiculosModelos.Vehiculo) *conectorModelos.GuardarVehiculoRespuesta {

	database, errConectarBD := gorm.Open("mysql", c.ConexionBd())

	if errConectarBD != nil {
		log.Fatal(constantes.ERROR_CONECTAR_BD)
	}

	database.AutoMigrate(&vehiculo)

	resultGuardarVehiculo := database.Create(
		&vehiculosModelos.Vehiculo{
			FechaLanzamiento: vehiculo.ObtenerFechaLanzamiento(),
			Marca:            vehiculo.ObtenerMarca(),
			Modelo:           vehiculo.ObtenerModelo(),
			Placas:           vehiculo.ObtenerPlacas(),
			Serie:            vehiculo.ObtenerSerie(),
		})
	respuesta := &conectorModelos.GuardarVehiculoRespuesta{}
	if resultGuardarVehiculo.Error != nil {
		respuesta.AsignarOk(false)
		respuesta.AsignarErr(errors.New(constantes.ERROR_GUARDAR_FILA_BD))
		return respuesta
	}
	respuesta.AsignarOk(true)
	respuesta.AsignarErr(nil)

	return respuesta
}

func (c *Controlador) ObtenerServicioVehicularPorNombre(solicitud *conectorModelos.ObtenerServicioVehicularPorNombreSolicitud) *servicioVehicularModelos.ServicioVehicular {

	database, errConectarBD := gorm.Open("mysql", c.ConexionBd())

	if errConectarBD != nil {
		log.Fatal(constantes.ERROR_CONECTAR_BD)
	}

	serviciosVehiculares := []servicioVehicularModelos.ServicioVehicular{}

	resultadoBusqueda := database.Where(&servicioVehicularModelos.ServicioVehicular{Nombre: solicitud.ObtenerNombre()}).Find(&serviciosVehiculares)

	if resultadoBusqueda.Error != nil {
		log.Fatal(constantes.ERROR_BUSQUEDA_EN_BD)
	}

	var servicioVehicularEncontrado *servicioVehicularModelos.ServicioVehicular

	if len(serviciosVehiculares) != 0 {
		servicioVehicularEncontrado = &serviciosVehiculares[0]
	}

	return servicioVehicularEncontrado
}

func (c *Controlador) GuardarServicioVehicular(servicioVehicular *servicioVehicularModelos.ServicioVehicular) *conectorModelos.GuardarServicioVehicularRespuesta {

	database, errConectarBD := gorm.Open("mysql", c.ConexionBd())

	if errConectarBD != nil {
		log.Fatal(constantes.ERROR_CONECTAR_BD)
	}

	database.AutoMigrate(&servicioVehicular)

	resultadoGuardarServicioVehicular := database.Create(
		&servicioVehicularModelos.ServicioVehicular{
			Nombre: servicioVehicular.ObtenerNombre(),
		})
	respuesta := &conectorModelos.GuardarServicioVehicularRespuesta{}
	if resultadoGuardarServicioVehicular.Error != nil {
		respuesta.AsignarOk(false)
		respuesta.AsignarErr(errors.New(constantes.ERROR_GUARDAR_FILA_BD))
		return respuesta
	}
	respuesta.AsignarOk(true)
	respuesta.AsignarErr(nil)

	return respuesta
}

func (c *Controlador) ObtenerServiciosVehiculares() []servicioVehicularModelos.ServicioVehicular {

	database, errConectarBD := gorm.Open("mysql", c.ConexionBd())

	if errConectarBD != nil {
		log.Fatal(constantes.ERROR_CONECTAR_BD)
	}

	serviciosVehiculares := []servicioVehicularModelos.ServicioVehicular{}

	resultadoBusqueda := database.Find(&serviciosVehiculares)

	if resultadoBusqueda.Error != nil {
		log.Fatal(constantes.ERROR_BUSQUEDA_EN_BD)
	}

	return serviciosVehiculares
}

func (c *Controlador) EliminarServicioVehicular(servicioVehicular *servicioVehicularModelos.ServicioVehicular) {

	database, errConectarBD := gorm.Open("mysql", c.ConexionBd())

	database.AutoMigrate(&servicioVehicular)

	if errConectarBD != nil {
		log.Fatal(constantes.ERROR_CONECTAR_BD)
	}

	database.Where("Nombre = ?", servicioVehicular.ObtenerNombre()).Delete(servicioVehicular)

	return
}

func (c *Controlador) EditarServicioVehicular(solicitud *conectorModelos.EditarServicioVehicularSolicitud) error {

	database, errConectarBD := gorm.Open("mysql", c.ConexionBd())

	servicioVehicularNuevo := &servicioVehicularModelos.ServicioVehicular{}
	servicioVehicularNuevo.AsignarNombre(solicitud.ObtenerNuevoNombre())

	if errConectarBD != nil {
		log.Fatal(constantes.ERROR_CONECTAR_BD)
	}

	database.Model(servicioVehicularNuevo).Where("Nombre = ?", solicitud.ObtenerNombreActual()).Update("nombre", servicioVehicularNuevo.ObtenerNombre())

	return nil
}

func (c *Controlador) GuardarUsuario(usuario *sesionModelos.Usuario) *conectorModelos.GuardarUsuarioRespuesta {

	database, errConectarBD := gorm.Open("mysql", c.ConexionBd())

	if errConectarBD != nil {
		log.Fatal(constantes.ERROR_CONECTAR_BD)
	}

	database.AutoMigrate(&usuario)

	resultadoGuardarUsuario := database.Create(usuario)
	respuesta := &conectorModelos.GuardarUsuarioRespuesta{}
	if resultadoGuardarUsuario.Error != nil {
		respuesta.AsignarOk(false)
		respuesta.AsignarErr(errors.New(constantes.ERROR_GUARDAR_FILA_BD))
		return respuesta
	}
	respuesta.AsignarOk(true)
	respuesta.AsignarErr(nil)

	return respuesta
}

func (c *Controlador) ObtenerUsuarioPorNombreUsuario(solicitud *conectorModelos.ObtenerUsuarioPorNombreUsuarioSolicitud) *sesionModelos.Usuario {

	database, errConectarBD := gorm.Open("mysql", c.ConexionBd())

	if errConectarBD != nil {
		log.Fatal(constantes.ERROR_CONECTAR_BD)
	}

	usuario := sesionModelos.Usuario{}

	resultadoBusqueda := database.Where(&sesionModelos.Usuario{NombreUsuario: solicitud.ObtenerNombre()}).Find(&usuario)

	if resultadoBusqueda.Error != nil {
		log.Fatal(constantes.ERROR_BUSQUEDA_EN_BD)
	}

	if (usuario == sesionModelos.Usuario{}) {
		return nil
	}

	return &usuario
}

func (c *Controlador) ConexionBd() string {
	dbHost := os.Getenv("dbHost")
	dbNombre := os.Getenv("dbNombre")
	dbUsuario := os.Getenv("dbUsuario")
	dbContrasena := os.Getenv("dbContrasena")
	dbPuerto := os.Getenv("dbPuerto")
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUsuario, dbContrasena, dbHost, dbPuerto, dbNombre)
}
