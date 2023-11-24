(function () {
    /*window.addEventListener('DOMContentLoaded',() => {
      //pedir los tiposDeRegistro al servidor

      //recibir los tipos de registro
      
      //construir los tipos de registro

    
      // const optionTipoDeRegistro = document.createElement('OPTION');

      // optionTipoDeRegistro.value = tipoDeServicio;

      // optionTipoDeRegistro.innerText = tipoDeServicio;

      // selectConcepto.appendChild(optionTipoDeRegistro);
    
        
    })*/

    const selectTipoServicioVehicular = document.getElementById('select_tipo_registro');
    selectTipoServicioVehicular.addEventListener('change', function (){
        const gasolina_label = document.getElementById('litros_gasolina_label');
        const gasolina_input = document.getElementById('litros_gasolina_input');
        const concepto = document.getElementById('concepto');
        if(selectTipoServicioVehicular.value == 'carga_combustible'){
            gasolina_label.removeAttribute("hidden"); 
            gasolina_input.removeAttribute("hidden"); 
            concepto.setAttribute('hidden', 'hidden'); 
        }else{
            gasolina_label.setAttribute('hidden', 'hidden'); 
            gasolina_input.setAttribute('hidden', 'hidden'); 
            concepto.removeAttribute('hidden');  
        }
    });

    const botonEnviarFormularioCapturarVehiculo = document.querySelector(
      ".captura-registro__formulario-contenedor .formulario__submit"
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
        const datosFormularioJSON = JSON.stringify(Object.fromEntries(datosFormulario));
        const agregarRegistroURL = "/AgregarRegistroMantenimientoVehiculo";
        const peticion = await fetch(agregarRegistroURL, {
          method: "POST",
          body: datosFormularioJSON,
        });
        const respuesta = await peticion.json();
        if (!respuesta.OK) {
          desplegarAlerta("error", respuesta.err);
        } else {
          desplegarAlerta("exito", "Registro agregado exitosamente");
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
  