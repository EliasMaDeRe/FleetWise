package controlador

import (
	"example/fleetwise/fuente/conectorbd/constantes"

	"example/fleetwise/modelos/conectorbd"
	vehiculosModelos "example/fleetwise/modelos/vehiculos"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB

type Controlador struct {
	conectorModelos conectorbd.GuardarVehiculoSolicitud
}

func (c *Controlador) ObtenerVehiculoPorPlacas(solicitud *conectorbd.ObtenerVehiculoPorPlacasSolicitud) *vehiculosModelos.Vehiculo {
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

	return nil
}

func (c *Controlador) ObtenerVehiculoPorSerie(solicitud *conectorbd.ObtenerVehiculoPorSerieSolicitud) *vehiculosModelos.Vehiculo {
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

	return nil
}

func (c *Controlador) AgregarVehiculo(vehiculo *vehiculosModelos.Vehiculo) error {
	var errConectarBD error

	database, errConectarBD := gorm.Open("mysql", constantes.CONEXION_BD)

	if errConectarBD != nil {
		log.Fatal(constantes.ERROR_CONECTAR_BD)
	}

	database.AutoMigrate(&vehiculo)

	//c.conectorModelos.AsignarVehiculo(*vehiculo)

	//resultCrearCelda := database.Create(c.conectorModelos.ObtenerVehiculo())

	resultCrearCelda := database.Create(
		&vehiculosModelos.Vehiculo{Anualidad: vehiculo.ObtenerAnualidad(),
			ID:     vehiculo.ObtenerID(),
			Marca:  vehiculo.ObtenerMarca(),
			Modelo: vehiculo.ObtenerModelo(),
			Placas: vehiculo.ObtenerPlacas(),
			Serie:  vehiculo.ObtenerSerie(),
		})

	if resultCrearCelda.Error != nil {
		return resultCrearCelda.Error
	}

	return nil
}
