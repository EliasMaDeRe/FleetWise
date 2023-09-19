(function () {
  const botonEnviarFormularioCapturarVehiculo = document.querySelector(
    ".captura-vehiculo__formulario-contenedor .formulario__submit"
  );

  if (botonEnviarFormularioCapturarVehiculo) {
    botonEnviarFormularioCapturarVehiculo.addEventListener("click", (e) => {
      e.preventDefault();
      const inputsFormulario = obtenerInputsDelFormulario();
      enviarFormulario(inputsFormulario);
    });

    function obtenerInputsDelFormulario() {
      return document.querySelectorAll(".formulario__input");
    }

    async function enviarFormulario(inputsFormulario) {
      const datosFormulario = construirPeticionFormulario(inputsFormulario);
      const agregarVehículoUrl = "/AgregarVehiculo";
      const peticion = await fetch(agregarVehículoUrl, {
        method: "POST",
        body: datosFormulario,
      });
      const respuesta = await peticion.json();
      console.log(respuesta);
      if (!respuesta.OK) {
        desplegarAlerta("error", "Error al crear el vehículo");
      } else {
        desplegarAlerta("exito", "Vehículo agregado exitosamente");
      }
    }

    function construirPeticionFormulario(inputsFormulario) {
      const datosFormulario = new FormData();
      inputsFormulario.forEach((inputFormulario) => {
        datosFormulario.append(inputFormulario.name, inputFormulario.value);
      });
      return datosFormulario;
    }

    function desplegarAlerta(tipoDeAlerta, mensajeAlerta) {
      const contenedorPrincipal = document.querySelector(
        ".captura-vehiculo__contenedor"
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
})();
