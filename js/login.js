(function () {
  const botonEnviarFormularioLogin = document.querySelector(
    ".login__formulario-contenedor .formulario__submit"
  );

  const logoHeader = document.querySelector(".header__logo");

  if(logoHeader){
      crearImagenLogo();
  }

  function crearImagenLogo(){
      const imagen = document.createElement("img");
      imagen.classList.add("header__logo-image");
      imagen.src = "img/Logo.png"
      imagen.addEventListener("click", (e) =>{
          e.preventDefault();
          solicitarWeb("/")
      })
      logoHeader.appendChild(imagen);
  }

  if (botonEnviarFormularioLogin) {
    botonEnviarFormularioLogin.addEventListener("click", (e) => {
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
      const peticion = await fetch("/Login", {
        method: "POST",
        body: datosFormularioJSON,
      });
      const respuesta = await peticion.json();
      if (!respuesta.OK) {
        desplegarAlerta("error", "Acceso no autorizado");
      } else {
        desplegarAlerta("exito", "Sesion iniciada con exito");
        setTimeout(async() => {
          window.location.href = "/";
        },1500);


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
        ".formulario"
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

  function solicitarWeb(url){
    window.location.href = url;
  }
  
})();