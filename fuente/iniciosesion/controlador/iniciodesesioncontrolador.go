package controladores

import (
	"errors"
	conectorBDControlador "example/fleetwise/fuente/conectorbdiniciosesion/controlador"
	"example/fleetwise/fuente/iniciosesion/constantes"
	inicioSesionMapeador "example/fleetwise/fuente/iniciosesion/mapeador"
	conectorBDModelos "example/fleetwise/modelos/conectorbdiniciosesion"
	inicioSesionModelos "example/fleetwise/modelos/iniciosesion"
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

type Controlador struct {
	InicioSesionMapeador  *inicioSesionMapeador.Mapeador
	ConectorBDControlador *conectorBDControlador.Controlador
}

func (c *Controlador) IniciarSesion(solicitud *inicioSesionModelos.IniciarSesionSolicitud) (tokenCadena string, err error) {

	if solicitud == nil {
		return tokenCadena, errors.New(constantes.ERROR_SOLICITUD_NULA)
	}

	obtenerUsuarioPorNombreUsuarioSolicitud := &conectorBDModelos.ObtenerUsuarioPorNombreUsuarioSolicitud{}
	obtenerUsuarioPorNombreUsuarioSolicitud.AsignarNombre(solicitud.ObtenerNombreUsuario())

	usuarioEncontrado := c.ConectorBDControlador.ObtenerUsuarioPorNombreUsuario(obtenerUsuarioPorNombreUsuarioSolicitud)
	fmt.Print(usuarioEncontrado)
	if usuarioEncontrado == nil {
		return tokenCadena, errors.New(constantes.ERROR_CREDENCIALES_INVALIDAS)
	}

	if err := bcrypt.CompareHashAndPassword([]byte(usuarioEncontrado.ClaveUsuario), []byte(solicitud.ObtenerClaveUsuario())); err != nil {
		return tokenCadena, errors.New(constantes.ERROR_CREDENCIALES_INVALIDAS)
	}

	tokenCadena, _ = c.crearJSONWebToken(usuarioEncontrado)

	return tokenCadena, nil
}

func (c *Controlador) ObtenerUsuarioPorNombreUsuario(solicitud *inicioSesionModelos.ObtenerUsuarioPorNombreUsuario) *inicioSesionModelos.Usuario {

	if solicitud == nil {
		return nil
	}

	obtenerUsuarioPorNombreUsuarioSolicitud := &conectorBDModelos.ObtenerUsuarioPorNombreUsuarioSolicitud{}
	obtenerUsuarioPorNombreUsuarioSolicitud.AsignarNombre(solicitud.ObtenerNombreUsuario())

	usuarioEncontrado := c.ConectorBDControlador.ObtenerUsuarioPorNombreUsuario(obtenerUsuarioPorNombreUsuarioSolicitud)

	return usuarioEncontrado
}

func (c *Controlador) crearJSONWebToken(usuario *inicioSesionModelos.Usuario) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	
	claims["expiracion"] = time.Now().Add(time.Hour * 2).Unix()
	claims["nombreUsuario"] = usuario.NombreUsuario
	claims["cargo"] = usuario.Cargo
	
	tokenCadena, err := token.SignedString([]byte(os.Getenv("llaveSecretaDeAutenticacion")))
	
	if err != nil {
		return "", err
	}
	
	return tokenCadena, nil
}

func (c *Controlador) UsuarioTieneAutorizacion(usuario *inicioSesionModelos.Usuario, cargoRequerido string) bool {
	
	codigoDeCargos := map[string]int{
		"capturista":    0,
		"supervisor":    1,
		"administrador": 2,
	}
	
	return codigoDeCargos[usuario.Cargo] >= codigoDeCargos[cargoRequerido]
}

func (c *Controlador) RegistrarUsuario(nombreUsuario string, claveUsuario string, cargo string) bool {

	claveUsuarioEncriptada, err := bcrypt.GenerateFromPassword([]byte(claveUsuario), 10)
	if err != nil {
		return false
	}

	usuarioARegistrar := &inicioSesionModelos.Usuario{
		Cargo:         cargo,
		NombreUsuario: nombreUsuario,
		ClaveUsuario:  string(claveUsuarioEncriptada),
	}

	guardarUsuarioRespuesta := c.ConectorBDControlador.GuardarUsuario(usuarioARegistrar)

	return guardarUsuarioRespuesta.ObtenerOk()
}