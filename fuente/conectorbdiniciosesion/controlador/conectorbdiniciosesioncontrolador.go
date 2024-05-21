package controlador

import (
	"errors"
	"example/fleetwise/fuente/conectorbdiniciosesion/constantes"
	conectorBDModelos "example/fleetwise/modelos/conectorbdiniciosesion"
	inicioSesionModelos "example/fleetwise/modelos/iniciosesion"
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
	var errConectarBD error

	baseDeDatos, errConectarBD := gorm.Open("mysql", c.obtenerConexionABd())

	if errConectarBD != nil {
		log.Fatal(constantes.ERROR_CONECTAR_BD)
	}

	usuario := inicioSesionModelos.Usuario{}

	// Llamada al procedimiento almacenado
	resultadoBusqueda := baseDeDatos.Raw("CALL obtenerUsuarioPorNombreUsuario(?)", solicitud.ObtenerNombre()).Find(&usuario)

	if resultadoBusqueda.Error != nil {
		fmt.Println(resultadoBusqueda.Error)
		return nil
	}

	if (usuario == inicioSesionModelos.Usuario{}) {
		return nil
	}

	return &usuario
}