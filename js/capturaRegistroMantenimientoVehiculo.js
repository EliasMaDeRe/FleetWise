(function () {
	window.addEventListener("DOMContentLoaded", () => {
		const selectConcepto = document.getElementById("conceptoregistro");
		obtenerServicios();

		const placasVehiculo = document.querySelector(".captura-registro-mantenimiento-vehiculo__placas-contenedor");
		desplegarPlacas();

		function obtenerServicios() {
			obtenerDatosServicios().then(function (response) {
				const datosServiciosRespuesta = response.ServiciosVehiculares;
				crearSelectorConcepto(datosServiciosRespuesta);
			});
		}

		async function obtenerDatosServicios() {
			const obtenerServiciosVehicularesParNuevoRegistroURL = "/ObtenerServiciosVehicularesParaNuevoRegistro";
			const peticion = await fetch(obtenerServiciosVehicularesParNuevoRegistroURL, {
				method: "POST",
			});
			const respuesta = await peticion.json();
			return respuesta;
		}

		function crearSelectorConcepto(datosServicios) {
			datosServicios.forEach((datoServicio) => {
				const optionTipoDeRegistro = document.createElement("OPTION");
				console.log(datoServicio);
				optionTipoDeRegistro.innerText = datoServicio.Nombre;
				console.log(optionTipoDeRegistro);
				selectConcepto.appendChild(optionTipoDeRegistro);
			});
		}

		function desplegarPlacas() {
			const placas = obtenerPlacasVehiculo();
			const elementoTarjeta = document.createElement("p");
			elementoTarjeta.textContent = "Placas del vehiculo a registrar: " + placas;
			placasVehiculo.appendChild(elementoTarjeta);
		}
	});

	function obtenerPlacasVehiculo() {
		var urlActual = new URL(window.location.href);

		var params = new URLSearchParams(urlActual.search);
		var placas = params.get("placas");
		return placas;
	}

	const selectTipoServicioVehicular = document.getElementById("tiporegistro");
	selectTipoServicioVehicular.addEventListener("change", function () {
		const gasolina_label = document.getElementById("litros_gasolina_label");
		const gasolina_input = document.getElementById("litrosdegasolina");
		const concepto = document.getElementById("concepto");
		if (selectTipoServicioVehicular.value == "carga de combustible") {
			gasolina_label.removeAttribute("hidden");
			gasolina_input.removeAttribute("hidden");
			concepto.setAttribute("hidden", "hidden");
		} else {
			gasolina_label.setAttribute("hidden", "hidden");
			gasolina_input.setAttribute("hidden", "hidden");
			concepto.removeAttribute("hidden");
		}
	});

	const botonEnviarFormulariRegistroMantenimientoVehiculo = document.querySelector(".captura-registro-mantenimiento-vehiculo__formulario-contenedor .formulario__submit");

	if (botonEnviarFormulariRegistroMantenimientoVehiculo) {
		botonEnviarFormulariRegistroMantenimientoVehiculo.addEventListener("click", (e) => {
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
			console.log(datosFormularioJSON);
			agregarRegistroURL = "/AgregarRegistroMantenimientoVehiculo";
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
			datosFormulario.append("placasVehiculo", obtenerPlacasVehiculo());
			return datosFormulario;
		}

		function desplegarAlerta(tipoDeAlerta, mensajeAlerta) {
			const contenedorPrincipal = document.querySelector(".captura-registro-mantenimiento-vehiculo__contenedor");
			const alertaExistente = contenedorPrincipal.querySelector(".alerta");
			if (alertaExistente) {
				contenedorPrincipal.removeChild(alertaExistente);
			}
			contenedorPrincipal.appendChild(construirAlerta(tipoDeAlerta, mensajeAlerta));
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
