package controlador

import (
	"errors"
	"example/fleetwise/fuente/conectorbd/constantes"

	servicioVehicularModelos "example/fleetwise/modelos/capturaserviciovehicular"
	vehiculosModelos "example/fleetwise/modelos/capturavehiculos"
	conectorModelos "example/fleetwise/modelos/conectorbd"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Controlador struct {
}

func (c *Controlador) ObtenerVehiculoPorPlacas(solicitud *conectorModelos.ObtenerVehiculoPorPlacasSolicitud) *vehiculosModelos.Vehiculo {
	var errConectarBD error

	database, errConectarBD := gorm.Open("mysql", constantes.CONEXION_BD)

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
	var errConectarBD error

	database, errConectarBD := gorm.Open("mysql", constantes.CONEXION_BD)

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
	var errConectarBD error

	database, errConectarBD := gorm.Open("mysql", constantes.CONEXION_BD)

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
	var errConectarBD error

	database, errConectarBD := gorm.Open("mysql", constantes.CONEXION_BD)

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
	var errConectarBD error

	database, errConectarBD := gorm.Open("mysql", constantes.CONEXION_BD)

	if errConectarBD != nil {
		log.Fatal(constantes.ERROR_CONECTAR_BD)
	}

	database.AutoMigrate(&servicioVehicular)

	resultGuardarServicioVehicular := database.Create(
		&servicioVehicularModelos.ServicioVehicular{
			Nombre: servicioVehicular.ObtenerNombre(),
		})
	respuesta := &conectorModelos.GuardarServicioVehicularRespuesta{}
	if resultGuardarServicioVehicular.Error != nil {
		respuesta.AsignarOk(false)
		respuesta.AsignarErr(errors.New(constantes.ERROR_GUARDAR_FILA_BD))
		return respuesta
	}
	respuesta.AsignarOk(true)
	respuesta.AsignarErr(nil)

	return respuesta
}

func (c *Controlador) ObtenerServiciosVehiculares() []servicioVehicularModelos.ServicioVehicular {
	var errConectarBD error

	database, errConectarBD := gorm.Open("mysql", constantes.CONEXION_BD)

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
