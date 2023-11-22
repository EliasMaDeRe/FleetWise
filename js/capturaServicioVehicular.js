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
      const peticion = await fetch(detectarURL(), {
        method: "POST",
        body: datosFormularioJSON,
      });
      const respuesta = await peticion.json();
      if (!respuesta.OK) {
        desplegarAlerta("error", respuesta.err);
      } else {
        desplegarAlerta("exito", obtenerMensaje());
      }
      borrarEditorTarjetas();
      obtenerTarjetas();
    }

    function obtenerMensaje(){
      var mensaje;
      if(document.querySelector(".captura-servicio-vehicular__card-editor")){
        mensaje = "Servicio editado exitosamente";
      }else{
        mensaje = "Servicio agregado exitosamente";
      }
      return mensaje
    }
    
    function detectarURL(){
      var urlDetectado;
      if(document.querySelector(".captura-servicio-vehicular__card-editor")){
        urlDetectado = "/EditarServicioVehicular";
      }else{
        urlDetectado = "/AgregarServicioVehicular";
      }
      return urlDetectado
    }

    function construirPeticionFormulario(inputsFormulario) {
      const datosFormulario = new FormData();

      if(document.querySelector(".captura-servicio-vehicular__card-editor")){
        datosFormulario.append("nombreactual",document.querySelector(".captura-servicio-vehicular__card-editor").name)
        inputsFormulario.forEach((inputFormulario) => {
          datosFormulario.append("nuevonombre", inputFormulario.value);
        });
        
      }else{
        inputsFormulario.forEach((inputFormulario) => {
          datosFormulario.append(inputFormulario.id, inputFormulario.value);
        });
      }
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
  obtenerTarjetas();

  
  function obtenerTarjetas(){
    obtenerDatosTarjetas().then(function(response){
      datosTarjetasRespuesta = response.ServiciosVehiculares;
      crearTarjetas(datosTarjetasRespuesta);
    });
  } 
  async function obtenerDatosTarjetas(){
    const obtenerServiciosVehicularesURL = "/ObtenerServiciosVehiculares";
    const peticion = await fetch(obtenerServiciosVehicularesURL, {
      method: "GET",
    });
    const respuesta = await peticion.json();
    return respuesta;
  }

  function crearTarjetas(datosTarjetas){
    contenedorTarjetas.replaceChildren();

    datosTarjetas.map((datosTarjeta)=>{
      const elementoTarjeta = document.createElement('div');
      elementoTarjeta.classList.add("card");
      elementoTarjeta.innerHTML= `
        <img class="card__edit" name="${datosTarjeta.Nombre}"src="https://upload.wikimedia.org/wikipedia/commons/6/64/Edit_icon_%28the_Noun_Project_30184%29.svg">
        <h3 class="card__heading">${datosTarjeta.Nombre}</h3>
        <img class="card__delete" name="${datosTarjeta.Nombre}"src="https://cdn-icons-png.flaticon.com/512/3389/3389152.png">
      `
      elementoTarjeta.lastElementChild.onclick = function(){
        borrarTarjeta(this.name);
      }
      elementoTarjeta.firstElementChild.onclick = function(){
        crearDivEditorServicio(this.name);
      }
      contenedorTarjetas.appendChild(elementoTarjeta)
    })
  }

  async function borrarTarjeta(nombreTarjeta){

    borrarEditorTarjetas();

    const nombreTarjetaJSON = JSON.stringify({"nombre":nombreTarjeta});
    const eliminarServicioVehicularURL = "/EliminarServicioVehicular";
    const peticion = await fetch(eliminarServicioVehicularURL, {
      method: "POST",
      body: nombreTarjetaJSON,
    });
    const respuesta = await peticion.json();
    if (!respuesta.OK) {
      desplegarAlerta("error", respuesta.err);
    } else {
      desplegarAlerta("exito", "Servicio eliminado exitosamente");
      obtenerTarjetas();
    }
  }

  function borrarEditorTarjetas(){
    const editorServicioVehicular = document.querySelector(".captura-servicio-vehicular__card-editor");
    if(editorServicioVehicular){
      editorServicioVehicular.remove();
    }
  }

  function crearDivEditorServicio(nombreServicioAEditar){

    var elementoEditarServicio = document.querySelector(".captura-servicio-vehicular__card-editor");

    if(elementoEditarServicio){
      elementoEditarServicio.firstChild.textContent = "Editando el servicio \""+nombreServicioAEditar+"\"";
      elementoEditarServicio.name = nombreServicioAEditar
      return;
    }
    
    elementoEditarServicio = document.createElement('div');
    elementoEditarServicio.classList.add("captura-servicio-vehicular__card-editor");
    elementoEditarServicio.name = nombreServicioAEditar
    const Formulario = document.querySelector(".captura-servicio-vehicular__formulario-contenedor");

    const labelEditarServicio = document.createElement('h3');
    labelEditarServicio.textContent = "Editando el servicio \""+nombreServicioAEditar+"\"";
    elementoEditarServicio.appendChild(labelEditarServicio);

    const botonCerrar = document.createElement('img');
    botonCerrar.classList.add("card-editor__delete");
    botonCerrar.src = "https://cdn-icons-png.flaticon.com/512/3389/3389152.png";
    botonCerrar.onclick = borrarEditorTarjetas;
    elementoEditarServicio.appendChild(botonCerrar);

    Formulario.before(elementoEditarServicio);
  }
  
})();
