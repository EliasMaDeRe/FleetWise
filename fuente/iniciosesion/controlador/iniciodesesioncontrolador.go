package controladores

import (
	"errors"
	conectorBDControlador "example/fleetwise/fuente/conectorbd/controlador"
	"example/fleetwise/fuente/iniciosesion/constantes"
	inicioSesionMapeador "example/fleetwise/fuente/iniciosesion/mapeador"
	conectorBDModelos "example/fleetwise/modelos/conectorbd"
	inicioSesionModelos "example/fleetwise/modelos/iniciosesion"
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
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

	if usuarioEncontrado == nil {
		return tokenCadena, errors.New(constantes.ERROR_CREDENCIALES_INVALIDAS)
	}

	if err := bcrypt.CompareHashAndPassword([]byte(usuarioEncontrado.ClaveUsuario), []byte(solicitud.ClaveUsuario)); err != nil {
		return tokenCadena, errors.New(constantes.ERROR_CREDENCIALES_INVALIDAS)
	}

	tokenCadena, _ = c.crearJSONWebToken(usuarioEncontrado)

	return tokenCadena, nil
}

func (c *Controlador) RegistrarUsuario(nombreUsuario string, contrasena string) bool {

	contrasenaEncriptada, err := bcrypt.GenerateFromPassword([]byte(contrasena), 10)
	if err != nil {
		return false
	}

	usuarioARegistrar := &inicioSesionModelos.Usuario{
		Cargo:         "administrador",
		NombreUsuario: nombreUsuario,
		ClaveUsuario:  string(contrasenaEncriptada),
	}

	guardarUsuarioRespuesta := c.ConectorBDControlador.GuardarUsuario(usuarioARegistrar)

	return guardarUsuarioRespuesta.ObtenerOk()
}

func (c *Controlador) crearJSONWebToken(usuario *inicioSesionModelos.Usuario) (string, error) {
	token := jwt.New(jwt.SigningMethodES256)
	claims := token.Claims.(jwt.MapClaims)

	claims["expiracion"] = time.Now().Add(time.Hour).Unix()
	claims["nombreUsuario"] = usuario.NombreUsuario
	claims["cargo"] = usuario.Cargo

	tokenCadena, err := token.SignedString(os.Getenv("llaveSecretaDeAutenticacion"))

	if err != nil {
		return "", err
	}

	return tokenCadena, nil
}

func (c *Controlador) validarJSONWebToken(siguienteManejador gin.HandlerFunc) gin.HandlerFunc {
	return func(contexto *gin.Context) {
		if contexto.Request.Header.Get("Token") != "" {
			token, err := jwt.Parse(contexto.Request.Header.Get("Token"), func(t *jwt.Token) (interface{}, error) {
				_, ok := t.Method.(*jwt.SigningMethodHMAC)
				if !ok {
					contexto.IndentedJSON(http.StatusUnauthorized, gin.H{"OK": false, "err": "Usuario no autorizado"})
				}
				return os.Getenv("llaveSecretaDeAutenticacion"), nil
			})

			if err != nil {
				contexto.IndentedJSON(http.StatusUnauthorized, gin.H{"OK": false, "err": err.Error()})
			}

			if token.Valid {
				siguienteManejador(contexto)
			}
		} else {
			contexto.IndentedJSON(http.StatusUnauthorized, gin.H{"OK": false, "err": "Usuario no autorizado"})
		}
	}
}
