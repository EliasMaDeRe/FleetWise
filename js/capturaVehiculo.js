(function () {
  const botonEnviarFormularioCapturarVehiculo = document.querySelector(
    ".captura-vehiculo__formulario-contenedor .formulario__submit"
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

  const barraNavegacion = document.querySelector(".navbar");

  const textoBotones = {
      Vehiculos: 'Capturar Vehiculos',
      Servicios: 'Capturar Conceptos',
  }

  const urlBotones = {
      Vehiculos: "/AgregarVehiculo",
      Servicios: "/AgregarServicioVehicular",
  }

  if(barraNavegacion){
      crearBoton("Vehiculos");
      crearBoton("Servicios");
      crearBotonCerrarSesion();
  }

  function crearBoton(servicio){
      const boton = document.createElement("a");
      boton.classList.add("navbar__enlace");
      boton.textContent = textoBotones[servicio];
      boton.addEventListener("click", (e) =>{
          e.preventDefault();
          solicitarWeb(urlBotones[servicio])
      })
      barraNavegacion.appendChild(boton);
  }

  function crearBotonCerrarSesion(){
      const boton = document.createElement("a");
      boton.classList.add("navbar__enlace--logout");
      boton.textContent = "Cerrar Sesión";
      boton.addEventListener("click", (e) =>{
          e.preventDefault();
          solicitarWeb("/Logout")
      })
      barraNavegacion.appendChild(boton);
  }

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
      const agregarVehiculoURL = "/AgregarVehiculo";
      const peticion = await fetch(agregarVehiculoURL, {
        method: "POST",
        body: datosFormularioJSON,
      });
      const respuesta = await peticion.json();
      if (!respuesta.OK) {
        desplegarAlerta("error", respuesta.err);
      } else {
        desplegarAlerta("exito", "Vehículo agregado exitosamente");
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
  function solicitarWeb(url){
    window.location.href = url;
  }

})();
