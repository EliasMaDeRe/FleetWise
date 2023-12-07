(function () {
	const botonEnviarFormularioCapturarServicioVehicular = document.querySelector(".captura-servicio-vehicular__formulario-contenedor .formulario__submit");

	if (botonEnviarFormularioCapturarServicioVehicular) {
		botonEnviarFormularioCapturarServicioVehicular.addEventListener("click", (e) => {
			e.preventDefault();
			enviarFormulario();
		});

		function obtenerInputsDelFormulario() {
			return document.querySelectorAll(".formulario__input");
		}

		async function enviarFormulario() {
			const URLAEnviar = detectarURL();
			const datosFormulario = construirPeticionFormulario(URLAEnviar);
			const datosFormularioJSON = JSON.stringify(Object.fromEntries(datosFormulario));
			const peticion = await fetch(URLAEnviar, {
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

		function obtenerMensaje() {
			var mensaje;
			if (document.querySelector(".captura-servicio-vehicular__card-editor")) {
				mensaje = "Servicio editado exitosamente";
			} else {
				mensaje = "Servicio agregado exitosamente";
			}
			return mensaje;
		}

		function detectarURL() {
			var urlDetectado;
			if (document.querySelector(".captura-servicio-vehicular__card-editor")) {
				urlDetectado = "/EditarServicioVehicular";
			} else {
				urlDetectado = "/AgregarServicioVehicular";
			}
			return urlDetectado;
		}

		function construirPeticionFormulario(URLAEnviar) {
			const nombreServicio = document.getElementById("nombre").value;
			if (nombreServicio === "") {
				desplegarAlerta("error", "Campo de nombre vacío");
				return;
			}
			const datosFormulario = new FormData();
			if(URLAEnviar == "/EditarServicioVehicular"){
				const nombreActual = document.querySelector(".captura-servicio-vehicular__card-editor").name;
				datosFormulario.append("nombreactual", nombreActual);	
				datosFormulario.append("nuevonombre", nombreServicio);	
			}else{
				datosFormulario.append("nombre", nombreServicio);				
			}
			return datosFormulario;
		}

		function desplegarAlerta(tipoDeAlerta, mensajeAlerta) {
			Swal.fire({
				title: tipoDeAlerta === "error" ? "¡Error!" : "¡Éxito!",
				text: mensajeAlerta,
				icon: tipoDeAlerta === "error" ? "error" : "success",
			});
		}
	}

	const contenedorTarjetas = document.querySelector(".captura-servicio-vehicular__card-container");
	obtenerTarjetas();

	function obtenerTarjetas() {
		obtenerDatosTarjetas().then(function (response) {
			datosTarjetasRespuesta = response.ServiciosVehiculares;
			crearTarjetas(datosTarjetasRespuesta);
		});
	}
	async function obtenerDatosTarjetas() {
		const obtenerServiciosVehicularesURL = "/ObtenerServiciosVehiculares";
		const peticion = await fetch(obtenerServiciosVehicularesURL, {
			method: "GET",
		});
		const respuesta = await peticion.json();
		return respuesta;
	}

	function crearTarjetas(datosTarjetas) {
		contenedorTarjetas.replaceChildren();

		datosTarjetas.map((datosTarjeta) => {
			const elementoTarjeta = document.createElement("div");
			elementoTarjeta.classList.add("card");
			elementoTarjeta.innerHTML = `
        <img class="card__edit" name="${datosTarjeta.Nombre}"src="https://upload.wikimedia.org/wikipedia/commons/6/64/Edit_icon_%28the_Noun_Project_30184%29.svg">
        <h3 class="card__heading">${datosTarjeta.Nombre}</h3>
        <img class="card__delete" name="${datosTarjeta.Nombre}"src="https://cdn-icons-png.flaticon.com/512/3389/3389152.png">
      `;
			elementoTarjeta.lastElementChild.onclick = function () {
				borrarTarjeta(this.name);
			};
			elementoTarjeta.firstElementChild.onclick = function () {
				crearDivEditorServicio(this.name);
			};
			contenedorTarjetas.appendChild(elementoTarjeta);
		});
	}

	async function borrarTarjeta(nombreTarjeta) {
		borrarEditorTarjetas();

		const nombreTarjetaJSON = JSON.stringify({ nombre: nombreTarjeta });
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

	function borrarEditorTarjetas() {
		const editorServicioVehicular = document.querySelector(".captura-servicio-vehicular__card-editor");
		if (editorServicioVehicular) {
			editorServicioVehicular.remove();
		}
	}

	function crearDivEditorServicio(nombreServicioAEditar) {
		var elementoEditarServicio = document.querySelector(".captura-servicio-vehicular__card-editor");

		if (elementoEditarServicio) {
			elementoEditarServicio.firstChild.textContent = 'Editando el servicio "' + nombreServicioAEditar + '"';
			elementoEditarServicio.name = nombreServicioAEditar;
			return;
		}

		elementoEditarServicio = document.createElement("div");
		elementoEditarServicio.classList.add("captura-servicio-vehicular__card-editor");
		elementoEditarServicio.name = nombreServicioAEditar;
		const Formulario = document.querySelector(".captura-servicio-vehicular__formulario-contenedor");

		const labelEditarServicio = document.createElement("h3");
		labelEditarServicio.textContent = 'Editando el servicio "' + nombreServicioAEditar + '"';
		elementoEditarServicio.appendChild(labelEditarServicio);

		const botonCerrar = document.createElement("img");
		botonCerrar.classList.add("card-editor__delete");
		botonCerrar.src = "https://cdn-icons-png.flaticon.com/512/3389/3389152.png";
		botonCerrar.onclick = borrarEditorTarjetas;
		elementoEditarServicio.appendChild(botonCerrar);

		Formulario.before(elementoEditarServicio);
	}
})();
