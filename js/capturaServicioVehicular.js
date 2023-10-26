(function () {
  const botonEnviarFormularioCapturarServicioVehicular = document.querySelector(
    ".captura-servicio-vehicular__formulario-contenedor .formulario__submit"
  );

  if (botonEnviarFormularioCapturarServicioVehicular) {
    botonEnviarFormularioCapturarServicioVehicular.addEventListener("click", (e) => {
      e.preventDefault();
      const inputsFormulario = obtenerInputsDelFormulario();
      enviarFormulario(inputsFormulario);
    });

    function obtenerInputsDelFormulario() {
      return document.querySelectorAll(".formulario__input");
    }

    async function enviarFormulario(inputsFormulario) {
      const datosFormulario = construirPeticionFormulario(inputsFormulario);
      const datosFormularioJSON = JSON.stringify(Object.fromEntries(datosFormulario));
      const agregarServicioVehicularURL = "/AgregarServicioVehicular";
      const peticion = await fetch(agregarServicioVehicularURL, {
        method: "POST",
        body: datosFormularioJSON,
      });
      const respuesta = await peticion.json();
      if (!respuesta.OK) {
        desplegarAlerta("error", respuesta.err);
      } else {
        desplegarAlerta("exito", "Servicio agregado exitosamente");
      }
    }

    function construirPeticionFormulario(inputsFormulario) {
      const datosFormulario = new FormData();
      inputsFormulario.forEach((inputFormulario) => {
        datosFormulario.append(inputFormulario.id, inputFormulario.value);
      });
      return datosFormulario;
    }

    function desplegarAlerta(tipoDeAlerta, mensajeAlerta) {
      const contenedorPrincipal = document.querySelector(
        ".captura-servicio-vehicular__contenedor"
      );
      const alertaExistente = contenedorPrincipal.querySelector(".alerta");
      if (alertaExistente) {
        contenedorPrincipal.removeChild(alertaExistente);
      }
      contenedorPrincipal.appendChild(
        construirAlerta(tipoDeAlerta, mensajeAlerta)
      );
    }

    function construirAlerta(tipoDeAlerta, mensajeDeAlerta) {
      const contenedorAlerta = crearContenedorAlerta(tipoDeAlerta);
      contenedorAlerta.append(crearTextoAlerta(mensajeDeAlerta));
      return contenedorAlerta;
    }

    function crearContenedorAlerta(tipoDeAlerta) {
      const contenedorAlerta = document.createElement("DIV");
      contenedorAlerta.classList.add("alerta");
      contenedorAlerta.classList.add(`alerta--${tipoDeAlerta}`);
      return contenedorAlerta;
    }

    function crearTextoAlerta(mensajeDeAlerta) {
      const textoAlerta = document.createElement("P");
      textoAlerta.textContent = mensajeDeAlerta;
      return textoAlerta;
    }
  }

  const contenedorTarjetas = document.querySelector(".captura-servicio-vehicular__card-container");

  
  obtenerDatosTarjetas().then(function(response){
    datosTarjetasRespuesta = response.ServiciosVehiculares;
    crearTarjetas(datosTarjetasRespuesta);
  });
  async function obtenerDatosTarjetas(){
    const obtenerServiciosVehicularesURL = "/ObtenerServiciosVehiculares";
    const peticion = await fetch(obtenerServiciosVehicularesURL, {
      method: "GET",
    });
    const respuesta = await peticion.json();
    return respuesta;
  }

  function crearTarjetas(datosTarjetas){
    datosTarjetas.map((datosTarjeta)=>{
      const elementoTarjeta = document.createElement('div');
      elementoTarjeta.classList.add("card");
      elementoTarjeta.innerHTML= `
        <h3 class="card__heading">${datosTarjeta.Nombre}</h3>
      `
      contenedorTarjetas.appendChild(elementoTarjeta)
    })
  }
  
})();
