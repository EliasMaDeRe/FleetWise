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

	function respuestaVacia(registro, vehiculoAsociado) {
		return registro.TipoRegistro === "" || vehiculoAsociado.Placas === "";
	}

	async function mostrarDatosDelRegistroEnElFormulario(registro, vehiculoAsociado) {
		cargarCampoEnElFormulario("placa", vehiculoAsociado.Placas);
        cargarCampoEnElFormulario("serie", vehiculoAsociado.Serie);
		cargarCampoEnElFormulario("marca", vehiculoAsociado.Marca);
		cargarCampoEnElFormulario("modelo", vehiculoAsociado.Modelo);
		cargarCampoEnElFormulario("fechaDeLanzamiento", vehiculoAsociado.FechaLanzamiento);
	}

	function cargarCampoEnElFormulario(idCampo, valorCampo) {
		document.getElementById(idCampo).value = valorCampo;
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

	async function ObtenerServiciosVehiculares() {
		const urlPeticion = "/ObtenerServiciosVehicularesParaNuevoRegistro";

		const respuesta = await fetch(urlPeticion, { method: "POST" });

		const registros = await respuesta.json();

		return await registros.ServiciosVehiculares;
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
