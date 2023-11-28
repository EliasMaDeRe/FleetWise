package controlador

import (
	"errors"
	"example/fleetwise/fuente/conectorbd/constantes"
	capturaRegistroMantenimientoVehiculoModelos "example/fleetwise/modelos/capturaregistromantenimientovehiculo"
	capturaServicioVehicularModelos "example/fleetwise/modelos/capturaserviciovehicular"
	capturaVehiculosModelos "example/fleetwise/modelos/capturavehiculos"
	conectorBDModelos "example/fleetwise/modelos/conectorbd"
	inicioSesionModelos "example/fleetwise/modelos/iniciosesion"
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Controlador struct {
}

func (c *Controlador) ObtenerVehiculoPorPlacas(solicitud *conectorBDModelos.ObtenerVehiculoPorPlacasSolicitud) *capturaVehiculosModelos.Vehiculo {
	var errConectarBD error

	baseDeDatos, errConectarBD := gorm.Open("mysql", c.obtenerConexionABd())

	if errConectarBD != nil {
		log.Fatal(constantes.ERROR_CONECTAR_BD)
	}

	vehiculos := []capturaVehiculosModelos.Vehiculo{}

	resultadoBusqueda := baseDeDatos.Where(&capturaVehiculosModelos.Vehiculo{Placas: solicitud.ObtenerPlacas()}).Find(&vehiculos)

	if resultadoBusqueda.Error != nil {
		log.Fatal(constantes.ERROR_PLACAS_INEXISTENTES_EN_BD)
	}

	var vehiculoEncontrado *capturaVehiculosModelos.Vehiculo

	if len(vehiculos) != 0 {
		vehiculoEncontrado = &vehiculos[0]
	}

	return vehiculoEncontrado
}

func (c *Controlador) ObtenerVehiculoPorSerie(solicitud *conectorBDModelos.ObtenerVehiculoPorSerieSolicitud) *capturaVehiculosModelos.Vehiculo {
	var errConectarBD error

	baseDeDatos, errConectarBD := gorm.Open("mysql", c.obtenerConexionABd())

	if errConectarBD != nil {
		log.Fatal(constantes.ERROR_CONECTAR_BD)
	}

	vehiculos := []capturaVehiculosModelos.Vehiculo{}

	resultadoBusqueda := baseDeDatos.Where(&capturaVehiculosModelos.Vehiculo{Serie: solicitud.ObtenerSerie()}).Find(&vehiculos)

	if resultadoBusqueda.Error != nil {
		log.Fatal(constantes.ERROR_SERIE_INEXISTENTES_EN_BD)
	}

	var vehiculoEncontrado *capturaVehiculosModelos.Vehiculo

	if len(vehiculos) != 0 {
		vehiculoEncontrado = &vehiculos[0]
	}

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

func (c *Controlador) GuardarRegistro(registro *capturaRegistroMantenimientoVehiculoModelos.RegistroMantenimientoVehiculo) *conectorBDModelos.GuardarRegistroMantenimientoVehiculoRespuesta {
	var errConectarBD error

	baseDeDatos, errConectarBD := gorm.Open("mysql", c.obtenerConexionABd())

	if errConectarBD != nil {
		log.Fatal(constantes.ERROR_CONECTAR_BD)
	}

	baseDeDatos.AutoMigrate(&registro)

	resultadoGuardarRegistro := baseDeDatos.Create(&registro)
	
	respuesta := &conectorBDModelos.GuardarRegistroMantenimientoVehiculoRespuesta{}
	if resultadoGuardarRegistro.Error != nil {
		respuesta.AsignarOk(false)
		respuesta.AsignarErr(errors.New(constantes.ERROR_GUARDAR_FILA_BD))
		return respuesta
	}
	respuesta.AsignarOk(true)
	respuesta.AsignarErr(nil)

	return respuesta
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

func (c *Controlador) EliminarServicioVehicular(servicioVehicular *capturaServicioVehicularModelos.ServicioVehicular) {

	baseDeDatos, errConectarBD := gorm.Open("mysql", c.obtenerConexionABd())

	baseDeDatos.AutoMigrate(&servicioVehicular)

	if errConectarBD != nil {
		log.Fatal(constantes.ERROR_CONECTAR_BD)
	}

	baseDeDatos.Where("Nombre = ?", servicioVehicular.ObtenerNombre()).Delete(servicioVehicular)

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

func (c *Controlador) GuardarUsuario(usuario *inicioSesionModelos.Usuario) *conectorBDModelos.GuardarUsuarioRespuesta {

	baseDeDatos, errConectarBD := gorm.Open("mysql", c.obtenerConexionABd())

	if errConectarBD != nil {
		log.Fatal(constantes.ERROR_CONECTAR_BD)
	}

	baseDeDatos.AutoMigrate(&usuario)

	resultadoGuardarUsuario := baseDeDatos.Create(usuario)
	respuesta := &conectorBDModelos.GuardarUsuarioRespuesta{}
	if resultadoGuardarUsuario.Error != nil {
		respuesta.AsignarOk(false)
		respuesta.AsignarErr(errors.New(constantes.ERROR_GUARDAR_FILA_BD))
		return respuesta
	}
	respuesta.AsignarOk(true)
	respuesta.AsignarErr(nil)

	return respuesta
}

func (c *Controlador) ObtenerUsuarioPorNombreUsuario(solicitud *conectorBDModelos.ObtenerUsuarioPorNombreUsuarioSolicitud) *inicioSesionModelos.Usuario {

	baseDeDatos, errConectarBD := gorm.Open("mysql", c.obtenerConexionABd())

	if errConectarBD != nil {
		log.Fatal(constantes.ERROR_CONECTAR_BD)
	}

	usuario := inicioSesionModelos.Usuario{}

	resultadoBusqueda := baseDeDatos.Where(&inicioSesionModelos.Usuario{NombreUsuario: solicitud.ObtenerNombre()}).Find(&usuario)

	if resultadoBusqueda.Error != nil {
		log.Fatal(constantes.ERROR_BUSQUEDA_EN_BD)
	}

	if (usuario == inicioSesionModelos.Usuario{}) {
		return nil
	}

	return &usuario
}

func (c *Controlador) ObtenerRegistrosYVehiculosAsociados() ([]capturaRegistroMantenimientoVehiculoModelos.RegistroMantenimientoVehiculo, []capturaVehiculosModelos.Vehiculo){
	baseDeDatos, errConectarBD := gorm.Open("mysql", c.obtenerConexionABd())

	if errConectarBD != nil {
		log.Fatal(constantes.ERROR_CONECTAR_BD)
	}

	registros := []capturaRegistroMantenimientoVehiculoModelos.RegistroMantenimientoVehiculo{}
	vehiculos := []capturaVehiculosModelos.Vehiculo{}

	tablaDeRegistrosConVehiculosFiltrados:= baseDeDatos.Select("*").Table("registros_mantenimiento_vehicular").Joins("JOIN vehiculos on vehiculos.placas = placas_vehiculo")

	if(tablaDeRegistrosConVehiculosFiltrados.Error != nil){
		log.Fatal(constantes.ERROR_BUSQUEDA_EN_BD)
	}

	tablaDeRegistrosConVehiculosFiltrados.Find(&registros)
	tablaDeRegistrosConVehiculosFiltrados.Find(&vehiculos)

 	return registros, vehiculos;
}

func (c *Controlador) ObtenerRegistrosYVehiculosAsociadosFiltrados(solicitud *conectorBDModelos.ObtenerRegistrosYVehiculosAsociadosFiltradosSolicitud)  ([]capturaRegistroMantenimientoVehiculoModelos.RegistroMantenimientoVehiculo, []capturaVehiculosModelos.Vehiculo) {

	baseDeDatos, errConectarBD := gorm.Open("mysql", c.obtenerConexionABd())

	if errConectarBD != nil {
		log.Fatal(constantes.ERROR_CONECTAR_BD)
	}

	registrosFiltrados := []capturaRegistroMantenimientoVehiculoModelos.RegistroMantenimientoVehiculo{}
	vehiculosFiltrados := []capturaVehiculosModelos.Vehiculo{}

	filtroRegistro := &capturaRegistroMantenimientoVehiculoModelos.RegistroMantenimientoVehiculo{TipoDeRegistro: solicitud.ObtenerFiltroTipoDeRegistro(),PlacasVehiculo: solicitud.ObtenerFiltroPlaca()}
	filtroVehiculo := &capturaVehiculosModelos.Vehiculo{FechaLanzamiento: solicitud.ObtenerFiltroFechaDeLanzamiento(), Marca: solicitud.ObtenerFiltroMarca(), Modelo: solicitud.ObtenerFiltroModelo()}

	tablaDeRegistrosConVehiculosFiltrados:= baseDeDatos.Select("*").Table("registros_mantenimiento_vehicular").Joins("JOIN vehiculos on vehiculos.placas = registros_mantenimiento_vehicular.placas_vehiculo").Where(&filtroRegistro).Where(&filtroVehiculo)

	if(tablaDeRegistrosConVehiculosFiltrados.Error != nil){
		log.Fatal(constantes.ERROR_BUSQUEDA_EN_BD)
	}

	tablaDeRegistrosConVehiculosFiltrados.Find(&registrosFiltrados)
	tablaDeRegistrosConVehiculosFiltrados.Find(&vehiculosFiltrados)

 	return registrosFiltrados, vehiculosFiltrados;
}

func (c *Controlador) ObtenerRegistroYVehiculoAsociadoPorNumeroDeRegistro(solicitud *capturaRegistroMantenimientoVehiculoModelos.ObtenerRegistroYVehiculoAsociadoPorNumeroDeRegistroSolicitud) (*capturaRegistroMantenimientoVehiculoModelos.RegistroMantenimientoVehiculo, *capturaVehiculosModelos.Vehiculo){
	baseDeDatos, errConectarBD := gorm.Open("mysql", c.obtenerConexionABd())

	if errConectarBD != nil {
		log.Fatal(constantes.ERROR_CONECTAR_BD)
	}

	registro := &capturaRegistroMantenimientoVehiculoModelos.RegistroMantenimientoVehiculo{NumeroDeRegistro: solicitud.ObtenerNumDeRegistro()}
	vehiculoAsociado := &capturaVehiculosModelos.Vehiculo{}

	consultaRegistroConNumeroDeRegistro:= baseDeDatos.Select("*").Table("registros_mantenimiento_vehicular").Joins("JOIN vehiculos on vehiculos.placas = registros_mantenimiento_vehicular.placas_vehiculo").Where(registro)

	if(consultaRegistroConNumeroDeRegistro.Error != nil){
		log.Fatal(constantes.ERROR_BUSQUEDA_EN_BD)
	}

	consultaRegistroConNumeroDeRegistro.Find(&registro)
	consultaRegistroConNumeroDeRegistro.Find(&vehiculoAsociado)

 	return registro, vehiculoAsociado;
}

func (c *Controlador) ActualizarRegistroMantenimientoVehicular(solicitud *conectorBDModelos.ActualizarRegistroDeMantenimientoDelVehiculoSolicitud) *conectorBDModelos.ActualizarRegistroMantenimientoVehicularRespuesta{
	
	baseDeDatos, errConectarBD := gorm.Open("mysql", c.obtenerConexionABd())


	registroMantenimientoVehicularActualizado := &capturaRegistroMantenimientoVehiculoModelos.RegistroMantenimientoVehiculo{
		TipoDeRegistro: 	solicitud.ObtenerTipo(),
		Fecha: 			solicitud.ObtenerFecha(),
		LitrosDeGasolina: solicitud.ObtenerLitrosDeGasolina(),	
		Kilometraje: 	solicitud.ObtenerKilometraje(),
		Importe: 		solicitud.ObtenerImporte(),
		Observaciones:  solicitud.ObtenerObservaciones(),
		Concepto:       solicitud.ObtenerConcepto(),
		PlacasVehiculo: solicitud.ObtenerPlacasVehiculo(),	
	}
	
	respuesta := &conectorBDModelos.ActualizarRegistroMantenimientoVehicularRespuesta{}

	if errConectarBD != nil {
		respuesta.AsignarOk(false)
		respuesta.AsignarError(errConectarBD)
		return respuesta
	}

	respuestaActualizacionRegistro := baseDeDatos.Table("registros_mantenimiento_vehicular").Where("numero_de_registro = ?", solicitud.ObtenerNumeroDeRegistro()).Update(registroMantenimientoVehicularActualizado)

	if(respuestaActualizacionRegistro.Error != nil){
		respuesta.AsignarOk(false)
		respuesta.AsignarError(respuestaActualizacionRegistro.Error)
		return respuesta
	}

	respuesta.AsignarOk(true);
	respuesta.AsignarError(nil)

	return respuesta
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
