package manejador

import (
	"errors"
	conectorBDControlador "example/fleetwise/fuente/conectorbdiniciosesion/controlador"
	"example/fleetwise/fuente/iniciosesion/constantes"
	inicioSesionControlador "example/fleetwise/fuente/iniciosesion/controlador"
	inicioSesionMapeador "example/fleetwise/fuente/iniciosesion/mapeador"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/microcosm-cc/bluemonday"
)

type Manejador struct {
	InicioSesionControlador *inicioSesionControlador.Controlador
}

func NuevoManejador() (c *Manejador) {
	return &Manejador{
		InicioSesionControlador: &inicioSesionControlador.Controlador{
			InicioSesionMapeador:  &inicioSesionMapeador.Mapeador{},
			ConectorBDControlador: &conectorBDControlador.Controlador{},
		},
	}
}

func (m *Manejador) IniciarSesion(contexto *gin.Context) {
	solicitud := m.InicioSesionControlador.InicioSesionMapeador.GinContextAIniciarSesionSolicitud(contexto)
	fmt.Println(solicitud)
	tokenCadena, err := m.InicioSesionControlador.IniciarSesion(solicitud)
	status := http.StatusOK
	ok := true
	if tokenCadena == "" {
		status = http.StatusUnauthorized
		ok = false
	} else {
		contexto.SetSameSite(http.SameSiteLaxMode)
		contexto.SetCookie("autorizacion", tokenCadena, 3600*2, "", "", false, true)
	}
	var mensajeError string
	if err != nil {
		mensajeError = err.Error()
	}
	fmt.Println(status)
	contexto.IndentedJSON(status, gin.H{"OK": ok, "err": mensajeError})
}

func (m *Manejador) ValidarSesion(contexto *gin.Context, cargoRequerido string) {

	tokenCadena, err := contexto.Cookie("autorizacion")
	
	tokenCadena = (bluemonday.StrictPolicy().Sanitize(tokenCadena))

	if err != nil {
		contexto.HTML(http.StatusUnauthorized, "login", gin.H{})
		contexto.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	token, err := jwt.Parse(tokenCadena, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New(constantes.ERROR_METODO_DE_INICIO_DE_SESION_DESCONOCIDO)
		}
		return []byte(os.Getenv("llaveSecretaDeAutenticacion")), nil
	})

	if err != nil {
		m.RetornarConSesionInvalida(contexto);
		return
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {

		if float64(time.Now().Unix()) > claims["expiracion"].(float64) {
			m.RetornarConSesionInvalida(contexto);
			return
		}

		nombreUsuarioEnClaims := m.InicioSesionControlador.InicioSesionMapeador.TokenClaimsAObtenerUsuarioPorNombreUsuarioSolicitud(claims)
		usuarioEncontrado := m.InicioSesionControlador.ObtenerUsuarioPorNombreUsuario(nombreUsuarioEnClaims)
		if usuarioEncontrado == nil {
			m.RetornarConSesionInvalida(contexto);
			return
		}

		if !m.InicioSesionControlador.UsuarioTieneAutorizacion(usuarioEncontrado, cargoRequerido) {
			contexto.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		contexto.Set("usuario", usuarioEncontrado)
	} else {
		contexto.HTML(http.StatusUnauthorized, "login", gin.H{})
		contexto.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	contexto.Next()
}

func (m *Manejador) RetornarConSesionInvalida(contexto *gin.Context){
	contexto.HTML(http.StatusUnauthorized, "login", gin.H{})
	contexto.AbortWithStatus(http.StatusUnauthorized)
}

func (m *Manejador) CerrarSesion(contexto *gin.Context) {

	contexto.SetSameSite(http.SameSiteLaxMode)
	contexto.SetCookie("autorizacion", "", int(time.Unix(0, 0).Unix()), "", "", false, true)
	contexto.Next()
}
