package controlador

import (
	"errors"
	"example/fleetwise/fuente/conectorbdcapturaregistromantenimientovehiculo/constantes"
	capturaRegistroMantenimientoVehiculoModelos "example/fleetwise/modelos/capturaregistromantenimientovehiculo"
	capturaVehiculosModelos "example/fleetwise/modelos/capturavehiculos"
	conectorBDModelos "example/fleetwise/modelos/conectorbdcapturaregistromantenimientovehiculo"
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


func (c *Controlador) GuardarRegistro(registro *capturaRegistroMantenimientoVehiculoModelos.RegistroDeMantenimientoDeVehiculo) *conectorBDModelos.GuardarRegistroMantenimientoVehiculoRespuesta {
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


func (c *Controlador) ObtenerRegistroYVehiculoAsociadoPorNumeroDeRegistro(solicitud *capturaRegistroMantenimientoVehiculoModelos.ObtenerRegistroYVehiculoAsociadoPorNumeroDeRegistroSolicitud) (*capturaRegistroMantenimientoVehiculoModelos.RegistroDeMantenimientoDeVehiculo, *capturaVehiculosModelos.Vehiculo) {
	baseDeDatos, errConectarBD := gorm.Open("mysql", c.obtenerConexionABd())

	if errConectarBD != nil {
		log.Fatal(constantes.ERROR_CONECTAR_BD)
	}

	registro := &capturaRegistroMantenimientoVehiculoModelos.RegistroDeMantenimientoDeVehiculo{NumeroDeRegistro: solicitud.ObtenerNumDeRegistro()}
	vehiculoAsociado := &capturaVehiculosModelos.Vehiculo{}

	consultaRegistroConNumeroDeRegistro := baseDeDatos.Select("*").Table(registro.TableName()).Joins("JOIN vehiculos on vehiculos.placas = registros_de_mantenimiento_de_vehiculo.placas_de_vehiculo").Where(registro)

	if consultaRegistroConNumeroDeRegistro.Error != nil {
		log.Fatal(constantes.ERROR_BUSQUEDA_EN_BD)
	}

	consultaRegistroConNumeroDeRegistro.Find(&registro)
	consultaRegistroConNumeroDeRegistro.Find(&vehiculoAsociado)

	return registro, vehiculoAsociado
}


func (c *Controlador) EditarRegistroDeMantenimientoDeVehiculo(solicitud *conectorBDModelos.EditarRegistroDeMantenimientoDelVehiculoSolicitud) *conectorBDModelos.ActualizarRegistroMantenimientoVehiculoRespuesta {

	baseDeDatos, errConectarBD := gorm.Open("mysql", c.obtenerConexionABd())

	registroMantenimientoVehiculoActualizado := &capturaRegistroMantenimientoVehiculoModelos.RegistroDeMantenimientoDeVehiculo{
		TipoDeRegistro:   solicitud.ObtenerTipo(),
		Fecha:            solicitud.ObtenerFecha(),
		LitrosDeGasolina: solicitud.ObtenerLitrosDeGasolina(),
		Kilometraje:      solicitud.ObtenerKilometraje(),
		Importe:          solicitud.ObtenerImporte(),
		Observaciones:    solicitud.ObtenerObservaciones(),
		Concepto:         solicitud.ObtenerConcepto(),
		PlacasDeVehiculo: solicitud.ObtenerPlacasVehiculo(),
	}

	respuesta := &conectorBDModelos.ActualizarRegistroMantenimientoVehiculoRespuesta{}

	if errConectarBD != nil {
		respuesta.AsignarOk(false)
		respuesta.AsignarErr(errConectarBD)
		return respuesta
	}

	respuestaActualizacionRegistro := baseDeDatos.Table(registroMantenimientoVehiculoActualizado.TableName()).Where("numero_de_registro = ?", solicitud.ObtenerNumeroDeRegistro()).Update(registroMantenimientoVehiculoActualizado)

	if respuestaActualizacionRegistro.Error != nil {
		respuesta.AsignarOk(false)
		respuesta.AsignarErr(respuestaActualizacionRegistro.Error)
		return respuesta
	}

	respuesta.AsignarOk(true)
	respuesta.AsignarErr(nil)

	return respuesta
}

