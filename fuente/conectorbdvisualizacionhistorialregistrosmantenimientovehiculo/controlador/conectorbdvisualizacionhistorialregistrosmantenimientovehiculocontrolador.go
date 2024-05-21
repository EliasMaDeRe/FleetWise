package controlador

import (
	"example/fleetwise/fuente/conectorbdvisualizacionhistorialregistrosmantenimientovehiculo/constantes"
	capturaRegistroMantenimientoVehiculoModelos "example/fleetwise/modelos/capturaregistromantenimientovehiculo"
	capturaVehiculosModelos "example/fleetwise/modelos/capturavehiculos"
	conectorBDModelos "example/fleetwise/modelos/conectorbdvisualizacionhistorialregistrosmantenimientovehiculo"
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
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

func (c *Controlador) ObtenerRegistrosYVehiculosAsociados() ([]capturaRegistroMantenimientoVehiculoModelos.RegistroDeMantenimientoDeVehiculo, []capturaVehiculosModelos.Vehiculo) {
	baseDeDatos, errConectarBD := gorm.Open("mysql", c.obtenerConexionABd())

	if errConectarBD != nil {
		log.Fatal(constantes.ERROR_CONECTAR_BD)
	}

	registro := &capturaRegistroMantenimientoVehiculoModelos.RegistroDeMantenimientoDeVehiculo{}

	registros := []capturaRegistroMantenimientoVehiculoModelos.RegistroDeMantenimientoDeVehiculo{}
	vehiculos := []capturaVehiculosModelos.Vehiculo{}

	tablaDeRegistrosConVehiculosFiltrados := baseDeDatos.Select("*").Table(registro.TableName()).Joins(constantes.CONSULTA_REGISTROS_CON_VEHICULOS)

	if tablaDeRegistrosConVehiculosFiltrados.Error != nil {
		log.Fatal(constantes.ERROR_BUSQUEDA_EN_BD)
	}

	tablaDeRegistrosConVehiculosFiltrados.Find(&registros)
	tablaDeRegistrosConVehiculosFiltrados.Find(&vehiculos)

	return registros, vehiculos
}

func (c *Controlador) ObtenerRegistrosYVehiculosAsociadosFiltrados(solicitud *conectorBDModelos.ObtenerRegistrosYVehiculosAsociadosFiltradosSolicitud) ([]capturaRegistroMantenimientoVehiculoModelos.RegistroDeMantenimientoDeVehiculo, []capturaVehiculosModelos.Vehiculo) {

	baseDeDatos, errConectarBD := gorm.Open("mysql", c.obtenerConexionABd())

	if errConectarBD != nil {
		log.Fatal(constantes.ERROR_CONECTAR_BD)
	}

	registrosFiltrados := []capturaRegistroMantenimientoVehiculoModelos.RegistroDeMantenimientoDeVehiculo{}
	vehiculosFiltrados := []capturaVehiculosModelos.Vehiculo{}

	filtroRegistro := &capturaRegistroMantenimientoVehiculoModelos.RegistroDeMantenimientoDeVehiculo{TipoDeRegistro: solicitud.ObtenerFiltroTipoDeRegistro(), PlacasDeVehiculo: solicitud.ObtenerFiltroPlaca()}
	filtroVehiculo := &capturaVehiculosModelos.Vehiculo{FechaLanzamiento: solicitud.ObtenerFiltroFechaDeLanzamiento(), Marca: solicitud.ObtenerFiltroMarca(), Modelo: solicitud.ObtenerFiltroModelo()}

	tablaDeRegistrosConVehiculosFiltrados := baseDeDatos.Select("*").Table(filtroRegistro.TableName()).Joins(constantes.CONSULTA_REGISTROS_CON_VEHICULOS).Where(&filtroRegistro).Where(&filtroVehiculo)

	if tablaDeRegistrosConVehiculosFiltrados.Error != nil {
		log.Fatal(constantes.ERROR_BUSQUEDA_EN_BD)
	}

	tablaDeRegistrosConVehiculosFiltrados.Find(&registrosFiltrados)
	tablaDeRegistrosConVehiculosFiltrados.Find(&vehiculosFiltrados)

	return registrosFiltrados, vehiculosFiltrados
}
