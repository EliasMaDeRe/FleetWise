(function () {
    const botonEnviarFormularioSeleccionarVehiculo = document.querySelector(
      ".seleccionar_vehiculo__formulario-contenedor .formulario__submit"
    );
  
    if (botonEnviarFormularioSeleccionarVehiculo) {
      botonEnviarFormularioSeleccionarVehiculo.addEventListener("click", (e) => {
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
        const seleccionarVehiculoParaNuevoRegistroURL = "/SeleccionarVehiculoParaNuevoRegistro";
        const peticion = await fetch(seleccionarVehiculoParaNuevoRegistroURL, {
          method: "POST",
          body: datosFormularioJSON,
        });
        const respuesta = await peticion.json();
        if (!respuesta.OK) {
          desplegarAlerta("error", respuesta.err);
        } else {
          desplegarAlerta("exito", "Vehículo encontrado exitosamente");
          setTimeout(() => {
            window.location.href = '/AgregarRegistroMantenimientoVehiculo?placas='+datosFormulario.get("placas");
          },1000);
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
          ".seleccionar_vehiculo__contenedor"
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
  