package controlador

import (
	"errors"
	"example/fleetwise/fuente/conectorbdcapturavehiculos/constantes"
	capturaVehiculosModelos "example/fleetwise/modelos/capturavehiculos"
	conectorBDModelos "example/fleetwise/modelos/conectorbdcapturavehiculos"
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
)

type Controlador struct{

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

func (c *Controlador) ObtenerVehiculoPorPlacas(solicitud *conectorBDModelos.ObtenerVehiculoPorPlacasSolicitud) *capturaVehiculosModelos.Vehiculo {
	var errConectarBD error

	baseDeDatos, errConectarBD := gorm.Open("mysql", c.obtenerConexionABd())

	if errConectarBD != nil {
		log.Fatal(constantes.ERROR_CONECTAR_BD)
	}

	vehiculos := []capturaVehiculosModelos.Vehiculo{}
	fmt.Println(solicitud)
	// Llamada al procedimiento almacenado
	resultadoBusqueda := baseDeDatos.Raw("CALL obtenerVehiculoPorPlacas(?)", solicitud.ObtenerPlacas()).Scan(&vehiculos)

	if resultadoBusqueda.Error != nil {
		return nil
	}

	var vehiculoEncontrado *capturaVehiculosModelos.Vehiculo
	vehiculoCeros := &capturaVehiculosModelos.Vehiculo{}

	if len(vehiculos) != 0 {
		vehiculoEncontrado = &vehiculos[0]
		if vehiculoEncontrado == vehiculoCeros {
			vehiculoEncontrado = nil
		}
	}
	fmt.Println(vehiculoEncontrado)
	return vehiculoEncontrado
}

func (c *Controlador) GuardarVehiculo(vehiculo *capturaVehiculosModelos.Vehiculo) *conectorBDModelos.GuardarVehiculoRespuesta {
	var errConectarBD error

	baseDeDatos, errConectarBD := gorm.Open("mysql", c.obtenerConexionABd())

	if errConectarBD != nil {
		log.Fatal(constantes.ERROR_CONECTAR_BD)
	}

	baseDeDatos.AutoMigrate(&vehiculo)

	resultGuardarVehiculo := baseDeDatos.Create(
		&capturaVehiculosModelos.Vehiculo{
			FechaLanzamiento: vehiculo.ObtenerFechaLanzamiento(),
			Marca:            vehiculo.ObtenerMarca(),
			Modelo:           vehiculo.ObtenerModelo(),
			Placas:           vehiculo.ObtenerPlacas(),
			Serie:            vehiculo.ObtenerSerie(),
		})
	respuesta := &conectorBDModelos.GuardarVehiculoRespuesta{}
	if resultGuardarVehiculo.Error != nil {
		respuesta.AsignarOk(false)
		respuesta.AsignarErr(errors.New(constantes.ERROR_GUARDAR_FILA_BD))
		return respuesta
	}
	respuesta.AsignarOk(true)
	respuesta.AsignarErr(nil)

	return respuesta
}

func (c *Controlador) ObtenerVehiculoPorSerie(solicitud *conectorBDModelos.ObtenerVehiculoPorSerieSolicitud) *capturaVehiculosModelos.Vehiculo {
	var errConectarBD error

	baseDeDatos, errConectarBD := gorm.Open("mysql", c.obtenerConexionABd())

	if errConectarBD != nil {
		log.Fatal(constantes.ERROR_CONECTAR_BD)
	}

	vehiculos := []capturaVehiculosModelos.Vehiculo{}

	resultadoBusqueda := baseDeDatos.Raw("CALL obtenerVehiculoPorSerie(?)", solicitud.ObtenerSerie()).Scan(&vehiculos)

	if resultadoBusqueda.Error != nil {
		return nil
	}

	var vehiculoEncontrado *capturaVehiculosModelos.Vehiculo
	vehiculoCeros := &capturaVehiculosModelos.Vehiculo{}

	if len(vehiculos) != 0 {
		vehiculoEncontrado = &vehiculos[0]
		if vehiculoEncontrado == vehiculoCeros {
			vehiculoEncontrado = nil
		}
	}

	return vehiculoEncontrado
}

func (c *Controlador) EditarVehiculo(solicitud *conectorBDModelos.EditarVehiculoSolicitud) error {
	baseDeDatos, errConectarBD := gorm.Open("mysql", c.obtenerConexionABd())

	vehiculoNuevo := &capturaVehiculosModelos.Vehiculo{}
	vehiculoNuevo.AsignarFechaLanzamiento(solicitud.ObtenerFechaDeLanzamientoNueva())
	vehiculoNuevo.AsignarMarca(solicitud.ObtenerMarcaNueva())
	vehiculoNuevo.AsignarModelo(solicitud.ObtenerModeloNuevo())
	vehiculoNuevo.AsignarPlacas(solicitud.ObtenerPlacasNuevas())
	vehiculoNuevo.AsignarSerie(solicitud.ObtenerSerieNuevo())

	if errConectarBD != nil {
		log.Fatal(constantes.ERROR_CONECTAR_BD)
	}

	respuestaEditarvehiculo := baseDeDatos.Table(vehiculoNuevo.TableName()).Where("placas = ?", solicitud.ObtenerPlacasActuales()).Update(vehiculoNuevo)

	return respuestaEditarvehiculo.Error
}

func (c *Controlador) EliminarVehiculo(vehiculo *capturaVehiculosModelos.Vehiculo) {

	baseDeDatos, errConectarBD := gorm.Open("mysql", c.obtenerConexionABd())

	baseDeDatos.AutoMigrate(&vehiculo)

	if errConectarBD != nil {
		log.Fatal(constantes.ERROR_CONECTAR_BD)
	}

	baseDeDatos.Where("Placas = ?", vehiculo.ObtenerPlacas()).Delete(vehiculo)
}