(function () {
	document.addEventListener("DOMContentLoaded", () => {
		cargarDatosDelRegistroEnElFormulario();
	});

	document.getElementById("tipoDeRegistro").addEventListener("change", (e) => {
		if (e.target.value === "carga de combustible") {
			mostrarCampoLitrosGasolina();
		} else {
			mostrarCampoConcepto();
		}
	});

	document.getElementById("editar-registro").addEventListener("click", async (e) => {
		e.preventDefault();

		const tipoDeRegistro = obtenerCampoDelFormulario("tipoDeRegistro");
		const registro = {
			numeroDeRegistro: obtenerNumDeRegistro(),
			tipoDeRegistro: tipoDeRegistro,
			fecha: obtenerCampoDelFormulario("fecha"),
			kilometraje: parseInt(obtenerCampoDelFormulario("kilometraje")),
			importe: parseFloat(obtenerCampoDelFormulario("importe")),
			observaciones: obtenerCampoDelFormulario("observaciones"),
			placaVehiculo: obtenerCampoDelFormulario("placa"),
		};

		if (tipoDeRegistro === "carga de combustible") {
			registro["litrosDeGasolina"] = parseFloat(obtenerCampoDelFormulario("litros"));
			registro["concepto"] = "";
		} else {
			registro["litrosDeGasolina"] = 0;
			registro["concepto"] = obtenerCampoDelFormulario("concepto");
		}

		const respuestaEditarRegistro = await enviarSolicitudEditarRegistro("/EditarRegistroMantenimientoVehiculo", registro);

		if (respuestaEditarRegistro.OK) {
			window.location.href = "/RegistroEditado?numRegistro=" + obtenerNumDeRegistro();
		} else {
			desplegarAlertaError(respuestaEditarRegistro.mensajeError);
		}
	});

	async function cargarDatosDelRegistroEnElFormulario() {
		const registroYVehiculoAsociado = await ObtenerDatosDelRegistro();

		if (respuestaVacia(registroYVehiculoAsociado.registro, registroYVehiculoAsociado.vehiculo)) {
			window.location.href = "/HistorialRegistrosMantenimientoVehiculo";
			return;
		}

		mostrarDatosDelRegistroEnElFormulario(registroYVehiculoAsociado.registro, registroYVehiculoAsociado.vehiculo);
	}

	function respuestaVacia(registro, vehiculoAsociado) {
		return registro.TipoRegistro === "" || vehiculoAsociado.Placas === "";
	}

	async function mostrarDatosDelRegistroEnElFormulario(registro, vehiculoAsociado) {
		cargarCampoEnElFormulario("placa", vehiculoAsociado.Placas);
		cargarCampoEnElFormulario("marca", vehiculoAsociado.Marca);
		cargarCampoEnElFormulario("modelo", vehiculoAsociado.Modelo);
		cargarCampoEnElFormulario("fechaDeLanzamiento", vehiculoAsociado.FechaLanzamiento);

		cargarCampoEnElFormulario("fecha", registro.Fecha);

		document.querySelector(`option[value="${registro.TipoDeRegistro}"]`).selected = true;

		cargarCampoEnElFormulario("kilometraje", registro.Kilometraje);
		cargarCampoEnElFormulario("importe", registro.Importe);
		cargarCampoEnElFormulario("observaciones", registro.Observaciones);

		if (registro.TipoDeRegistro === "carga de combustible") {
			mostrarCampoLitrosGasolina();
			cargarCampoEnElFormulario("litros", registro.LitrosDeGasolina);
		} else {
			await mostrarCampoConcepto();
			cargarCampoEnElFormulario("concepto", registro.Concepto);
		}
	}

	function cargarCampoEnElFormulario(idCampo, valorCampo) {
		document.getElementById(idCampo).value = valorCampo;
	}

	function mostrarCampoLitrosGasolina() {
		const campoLitros = document.getElementById("litros");

		if (!campoLitros) {
			const divCampoCambiable = document.getElementById("campoCambiable");
			divCampoCambiable.innerHTML = "";

			const labelLitrosGasolina = document.createElement("LABEL");
			labelLitrosGasolina.innerText = "Litros de Gasolina";
			labelLitrosGasolina.setAttribute("for", "litros");
			labelLitrosGasolina.classList.add("registros__etiqueta");
			divCampoCambiable.appendChild(labelLitrosGasolina);

			const inputLitrosGasolina = document.createElement("INPUT");
			inputLitrosGasolina.setAttribute("type", "number");
			inputLitrosGasolina.setAttribute("id", "litros");
			inputLitrosGasolina.classList.add("registros__entrada");
			divCampoCambiable.appendChild(inputLitrosGasolina);
		}
	}

	async function mostrarCampoConcepto() {
		const campoObservaciones = document.getElementById("concepto");

		if (!campoObservaciones) {
			const divCampoCambiable = document.getElementById("campoCambiable");
			divCampoCambiable.innerHTML = "";

			const labelObservaciones = document.createElement("LABEL");
			labelObservaciones.innerText = "Concepto";

			labelObservaciones.setAttribute("for", "concepto");
			labelObservaciones.classList.add("registros__etiqueta");
			divCampoCambiable.appendChild(labelObservaciones);

			const selectConcepto = document.createElement("SELECT");
			selectConcepto.setAttribute("id", "concepto");
			selectConcepto.classList.add("campo__desplegable");

			const serviciosVehiculares = await ObtenerServiciosVehiculares();

			serviciosVehiculares.forEach((servicioVehicular) => {
				const optionServicioVehicular = document.createElement("OPTION");
				optionServicioVehicular.setAttribute("value", servicioVehicular.Nombre);
				optionServicioVehicular.innerText = servicioVehicular.Nombre;

				selectConcepto.appendChild(optionServicioVehicular);
			});

			divCampoCambiable.appendChild(selectConcepto);
		}
	}

	async function ObtenerDatosDelRegistro() {
		const peticionUrl = "/ObtenerRegistroPorNumeroDeRegistro";

		const peticionContenido = {
			NumDeRegistro: obtenerNumDeRegistro(),
		};
		const respuesta = await fetch(peticionUrl, {
			method: "POST",
			body: JSON.stringify(peticionContenido),
		});

		return await respuesta.json();
	}

	async function ObtenerServiciosVehiculares() {
		const urlPeticion = "/ObtenerServiciosVehicularesParaNuevoRegistro";

		const respuesta = await fetch(urlPeticion, { method: "POST" });

		const registros = await respuesta.json();

		return await registros.ServiciosVehiculares;
	}

	function obtenerNumDeRegistro() {
		const params = new URLSearchParams(window.location.search);
		return parseInt(params.get("numRegistro"));
	}

	function obtenerCampoDelFormulario(idCampo) {
		return document.getElementById(idCampo).value;
	}

	async function enviarSolicitudEditarRegistro(url, contenidoPeticion) {
		const respuesta = await fetch(url, {
			method: "POST",
			body: JSON.stringify(contenidoPeticion),
		});

		const resultado = await respuesta.json();

		return resultado;
	}

	function desplegarAlertaError(mensajeError) {
		const contenedorPrincipal = document.querySelector("main.editar-registro");

		const contenedorAlerta = document.querySelector("div.alerta");

		if (contenedorAlerta) {
			const textoAlerta = contenedorAlerta.querySelector("p");
			textoAlerta.textContent = mensajeError;
		} else {
			const contenedorAlerta = document.createElement("DIV");
			contenedorAlerta.classList.add("alerta");
			contenedorAlerta.classList.add("alerta--error");

			const textoAlerta = document.createElement("P");
			textoAlerta.textContent = mensajeError;

			contenedorAlerta.appendChild(textoAlerta);

			contenedorPrincipal.appendChild(contenedorAlerta);
		}
	}
})();
