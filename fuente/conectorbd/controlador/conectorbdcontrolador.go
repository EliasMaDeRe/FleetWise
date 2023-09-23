package controlador

import (
	"errors"
	"example/fleetwise/fuente/conectorbd/constantes"

	conectorModelos "example/fleetwise/modelos/conectorbd"
	vehiculosModelos "example/fleetwise/modelos/vehiculos"
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

	vehiculoEncontrado := &vehiculos[0]

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
		log.Fatal(constantes.ERROR_SERIE_INEXISTENTES_EN_BD)
	}

	vehiculoEncontrado := &vehiculos[0]

	return vehiculoEncontrado
}

func (c *Controlador) GuardarVehiculo(vehiculo *vehiculosModelos.Vehiculo) *conectorModelos.AgregarConectorBDRespuesta {
	var errConectarBD error

	database, errConectarBD := gorm.Open("mysql", constantes.CONEXION_BD)

	if errConectarBD != nil {
		log.Fatal(constantes.ERROR_CONECTAR_BD)
	}

	database.AutoMigrate(&vehiculo)

	resultCrearCelda := database.Create(
		&vehiculosModelos.Vehiculo{Anualidad: vehiculo.ObtenerAnualidad(),
			ID:     vehiculo.ObtenerID(),
			Marca:  vehiculo.ObtenerMarca(),
			Modelo: vehiculo.ObtenerModelo(),
			Placas: vehiculo.ObtenerPlacas(),
			Serie:  vehiculo.ObtenerSerie(),
		})
	respuesta := &conectorModelos.AgregarConectorBDRespuesta{}
	if resultCrearCelda.Error != nil {
		respuesta.AsignarOk(false)
		respuesta.AsignarErr(errors.New(constantes.ERROR_GUARDAR_FILA_BD))
		return respuesta
	}
	respuesta.AsignarOk(true)
	respuesta.AsignarErr(nil)

	return respuesta
}
