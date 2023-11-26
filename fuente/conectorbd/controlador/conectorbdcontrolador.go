package controlador

import (
	"errors"
	"example/fleetwise/fuente/conectorbd/constantes"
	registroMantenimientoVehiculoModelos "example/fleetwise/modelos/capturaregistromantenimientovehiculo"
	servicioVehicularModelos "example/fleetwise/modelos/capturaserviciovehicular"
	"example/fleetwise/modelos/capturavehiculos"
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
	var errConectarBD error

	baseDeDatos, errConectarBD := gorm.Open("mysql", c.obtenerConexionABd())

	if errConectarBD != nil {
		log.Fatal(constantes.ERROR_CONECTAR_BD)
	}

	vehiculos := []vehiculosModelos.Vehiculo{}

	resultadoBusqueda := baseDeDatos.Where(&vehiculosModelos.Vehiculo{Placas: solicitud.ObtenerPlacas()}).Find(&vehiculos)

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

	baseDeDatos, errConectarBD := gorm.Open("mysql", c.obtenerConexionABd())

	if errConectarBD != nil {
		log.Fatal(constantes.ERROR_CONECTAR_BD)
	}

	vehiculos := []vehiculosModelos.Vehiculo{}

	resultadoBusqueda := baseDeDatos.Where(&vehiculosModelos.Vehiculo{Serie: solicitud.ObtenerSerie()}).Find(&vehiculos)

	if resultadoBusqueda.Error != nil {
		log.Fatal(constantes.ERROR_SERIE_INEXISTENTES_EN_BD)
	}

	var vehiculoEncontrado *vehiculosModelos.Vehiculo

	if len(vehiculos) != 0 {
		vehiculoEncontrado = &vehiculos[0]
	}

	return vehiculoEncontrado
}

func (c *Controlador) GuardarVehiculo(vehiculo *vehiculosModelos.Vehiculo) *conectorModelos.GuardarVehiculoRespuesta {
	var errConectarBD error

	baseDeDatos, errConectarBD := gorm.Open("mysql", c.obtenerConexionABd())

	if errConectarBD != nil {
		log.Fatal(constantes.ERROR_CONECTAR_BD)
	}

	baseDeDatos.AutoMigrate(&vehiculo)

	resultGuardarVehiculo := baseDeDatos.Create(
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

func (c *Controlador) GuardarRegistro(registro *registroMantenimientoVehiculoModelos.RegistroMantenimientoVehiculo) *conectorModelos.GuardarRegistroMantenimientoVehiculoRespuesta {
	var errConectarBD error

	baseDeDatos, errConectarBD := gorm.Open("mysql", c.obtenerConexionABd())

	if errConectarBD != nil {
		log.Fatal(constantes.ERROR_CONECTAR_BD)
	}

	baseDeDatos.AutoMigrate(&registro)

	resultadoGuardarRegistro := baseDeDatos.Create(
		&registroMantenimientoVehiculoModelos.RegistroMantenimientoVehiculo{
			NumeroDeRegistro: registro.ObtenerNumeroDeRegistro(),
			TipoRegistro:             registro.ObtenerTipo(),
			Fecha:            registro.ObtenerFecha(),
			LitrosGasolina: registro.ObtenerLitrosDeGasolina(),
			Kilometraje:      registro.ObtenerKilometraje(),
			Importe:          registro.ObtenerImporte(),
			Observaciones:    registro.ObtenerObservaciones(),
			Concepto:         registro.ObtenerConcepto(),
		})
	respuesta := &conectorModelos.GuardarRegistroMantenimientoVehiculoRespuesta{}
	if resultadoGuardarRegistro.Error != nil {
		respuesta.AsignarOk(false)
		respuesta.AsignarErr(errors.New(constantes.ERROR_GUARDAR_FILA_BD))
		return respuesta
	}
	respuesta.AsignarOk(true)
	respuesta.AsignarErr(nil)

	return respuesta
}

func (c *Controlador) ObtenerServicioVehicularPorNombre(solicitud *conectorModelos.ObtenerServicioVehicularPorNombreSolicitud) *servicioVehicularModelos.ServicioVehicular {

	baseDeDatos, errConectarBD := gorm.Open("mysql", c.obtenerConexionABd())

	if errConectarBD != nil {
		log.Fatal(constantes.ERROR_CONECTAR_BD)
	}

	serviciosVehiculares := []servicioVehicularModelos.ServicioVehicular{}

	resultadoBusqueda := baseDeDatos.Where(&servicioVehicularModelos.ServicioVehicular{Nombre: solicitud.ObtenerNombre()}).Find(&serviciosVehiculares)

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

	baseDeDatos, errConectarBD := gorm.Open("mysql", c.obtenerConexionABd())

	if errConectarBD != nil {
		log.Fatal(constantes.ERROR_CONECTAR_BD)
	}

	baseDeDatos.AutoMigrate(&servicioVehicular)

	resultadoGuardarServicioVehicular := baseDeDatos.Create(
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

	baseDeDatos, errConectarBD := gorm.Open("mysql", c.obtenerConexionABd())

	if errConectarBD != nil {
		log.Fatal(constantes.ERROR_CONECTAR_BD)
	}

	serviciosVehiculares := []servicioVehicularModelos.ServicioVehicular{}

	resultadoBusqueda := baseDeDatos.Find(&serviciosVehiculares)

	if resultadoBusqueda.Error != nil {
		log.Fatal(constantes.ERROR_BUSQUEDA_EN_BD)
	}

	return serviciosVehiculares
}

func (c *Controlador) EliminarServicioVehicular(servicioVehicular *servicioVehicularModelos.ServicioVehicular) {

	baseDeDatos, errConectarBD := gorm.Open("mysql", c.obtenerConexionABd())

	baseDeDatos.AutoMigrate(&servicioVehicular)

	if errConectarBD != nil {
		log.Fatal(constantes.ERROR_CONECTAR_BD)
	}

	baseDeDatos.Where("Nombre = ?", servicioVehicular.ObtenerNombre()).Delete(servicioVehicular)

	return
}

func (c *Controlador) EditarServicioVehicular(solicitud *conectorModelos.EditarServicioVehicularSolicitud) error {

	baseDeDatos, errConectarBD := gorm.Open("mysql", c.obtenerConexionABd())

	servicioVehicularNuevo := &servicioVehicularModelos.ServicioVehicular{}
	servicioVehicularNuevo.AsignarNombre(solicitud.ObtenerNuevoNombre())

	if errConectarBD != nil {
		log.Fatal(constantes.ERROR_CONECTAR_BD)
	}

	baseDeDatos.Model(servicioVehicularNuevo).Where("Nombre = ?", solicitud.ObtenerNombreActual()).Update("nombre", servicioVehicularNuevo.ObtenerNombre())

	return nil
}

func (c *Controlador) GuardarUsuario(usuario *sesionModelos.Usuario) *conectorModelos.GuardarUsuarioRespuesta {

	baseDeDatos, errConectarBD := gorm.Open("mysql", c.obtenerConexionABd())

	if errConectarBD != nil {
		log.Fatal(constantes.ERROR_CONECTAR_BD)
	}

	baseDeDatos.AutoMigrate(&usuario)

	resultadoGuardarUsuario := baseDeDatos.Create(usuario)
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

	baseDeDatos, errConectarBD := gorm.Open("mysql", c.obtenerConexionABd())

	if errConectarBD != nil {
		log.Fatal(constantes.ERROR_CONECTAR_BD)
	}

	usuario := sesionModelos.Usuario{}

	resultadoBusqueda := baseDeDatos.Where(&sesionModelos.Usuario{NombreUsuario: solicitud.ObtenerNombre()}).Find(&usuario)

	if resultadoBusqueda.Error != nil {
		log.Fatal(constantes.ERROR_BUSQUEDA_EN_BD)
	}

	if (usuario == sesionModelos.Usuario{}) {
		return nil
	}

	return &usuario
}

func (c *Controlador) ObtenerRegistrosYVehiculosAsociados() ([]registroMantenimientoVehiculoModelos.RegistroMantenimientoVehiculo, []vehiculosModelos.Vehiculo){
	baseDeDatos, errConectarBD := gorm.Open("mysql", c.obtenerConexionABd())

	if errConectarBD != nil {
		log.Fatal(constantes.ERROR_CONECTAR_BD)
	}

	registros := []registroMantenimientoVehiculoModelos.RegistroMantenimientoVehiculo{}
	vehiculos := []vehiculosModelos.Vehiculo{}

	tablaDeRegistrosConVehiculosFiltrados:= baseDeDatos.Select("*").Table("registros_mantenimiento_vehicular").Joins("JOIN vehiculos on vehiculos.placas = registros_mantenimiento_vehicular.placas_vehiculo")

	if(tablaDeRegistrosConVehiculosFiltrados.Error != nil){
		log.Fatal(constantes.ERROR_BUSQUEDA_EN_BD)
	}

	tablaDeRegistrosConVehiculosFiltrados.Find(&registros)
	tablaDeRegistrosConVehiculosFiltrados.Find(&vehiculos)

 	return registros, vehiculos;
}

func (c *Controlador) ObtenerRegistrosYVehiculosAsociadosFiltrados(solicitud *conectorModelos.ObtenerRegistrosYVehiculosAsociadosFiltradosSolicitud)  ([]registroMantenimientoVehiculoModelos.RegistroMantenimientoVehiculo, []vehiculosModelos.Vehiculo) {

	baseDeDatos, errConectarBD := gorm.Open("mysql", c.obtenerConexionABd())

	if errConectarBD != nil {
		log.Fatal(constantes.ERROR_CONECTAR_BD)
	}

	registrosFiltrados := []registroMantenimientoVehiculoModelos.RegistroMantenimientoVehiculo{}
	vehiculosFiltrados := []vehiculosModelos.Vehiculo{}

	filtroRegistro := &registroMantenimientoVehiculoModelos.RegistroMantenimientoVehiculo{TipoRegistro: solicitud.ObtenerFiltroTipoDeRegistro(),PlacasVehiculo: solicitud.ObtenerFiltroPlaca()}
	filtroVehiculo := &vehiculosModelos.Vehiculo{FechaLanzamiento: solicitud.ObtenerFiltroFechaDeLanzamiento(), Marca: solicitud.ObtenerFiltroMarca(), Modelo: solicitud.ObtenerFiltroModelo()}

	tablaDeRegistrosConVehiculosFiltrados:= baseDeDatos.Select("*").Table("registros_mantenimiento_vehicular").Joins("JOIN vehiculos on vehiculos.placas = registros_mantenimiento_vehicular.placas_vehiculo").Where(&filtroRegistro).Where(&filtroVehiculo)

	if(tablaDeRegistrosConVehiculosFiltrados.Error != nil){
		log.Fatal(constantes.ERROR_BUSQUEDA_EN_BD)
	}

	tablaDeRegistrosConVehiculosFiltrados.Find(&registrosFiltrados)
	tablaDeRegistrosConVehiculosFiltrados.Find(&vehiculosFiltrados)

 	return registrosFiltrados, vehiculosFiltrados;
}

func (c *Controlador) ObtenerRegistroYVehiculoAsociadoPorNumeroDeRegistro (solicitud *registroMantenimientoVehiculoModelos.ObtenerRegistroYVehiculoAsociadoPorNumeroDeRegistroSolicitud) (*registroMantenimientoVehiculoModelos.RegistroMantenimientoVehiculo, *capturavehiculos.Vehiculo){
	baseDeDatos, errConectarBD := gorm.Open("mysql", c.obtenerConexionABd())

	if errConectarBD != nil {
		log.Fatal(constantes.ERROR_CONECTAR_BD)
	}

	registro := &registroMantenimientoVehiculoModelos.RegistroMantenimientoVehiculo{NumeroDeRegistro: solicitud.ObtenerNumDeRegistro()}
	vehiculoAsociado := &vehiculosModelos.Vehiculo{}

	consultaRegistroConNumeroDeRegistro:= baseDeDatos.Select("*").Table("registros_mantenimiento_vehicular").Joins("JOIN vehiculos on vehiculos.placas = registros_mantenimiento_vehicular.placas_vehiculo").Where(registro)

	if(consultaRegistroConNumeroDeRegistro.Error != nil){
		log.Fatal(constantes.ERROR_BUSQUEDA_EN_BD)
	}

	consultaRegistroConNumeroDeRegistro.Find(&registro)
	consultaRegistroConNumeroDeRegistro.Find(&vehiculoAsociado)

 	return registro, vehiculoAsociado;
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
