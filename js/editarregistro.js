(function () {
	document.getElementById("tipoDeRegistro").addEventListener("change", (e) => {
		if (e.target.value === "carga de combustible") {
			mostrarCampoLitrosGasolina();
		} else {
			mostrarCampoObservaciones();
		}
	});

	function mostrarCampoLitrosGasolina() {
		//TODO
	}

	function mostrarCampoObservaciones() {
		//TODO
	}

	document.addEventListener("DOMContentLoaded", () => {
		cargarDatosDelRegistroEnElFormulario();
	});

	async function cargarDatosDelRegistroEnElFormulario() {
		const registroYVehiculoAsociado = await ObtenerDatosDelRegistro();

		if (respuestaVacia(registroYVehiculoAsociado.registro, registroYVehiculoAsociado.vehiculo)) {
			window.location.href = "/HistorialRegistrosMantenimientoVehicular";
			return;
		}

		mostrarDatosDelRegistroEnElFormulario(registroYVehiculoAsociado.registro, registroYVehiculoAsociado.vehiculo);
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

	function obtenerNumDeRegistro() {
		const params = new URLSearchParams(window.location.search);
		return parseInt(params.get("numRegistro"));
	}

	function respuestaVacia(registro, vehiculoAsociado) {
		return registro.TipoRegistro === "" || vehiculoAsociado.Placas === "";
	}

	function mostrarDatosDelRegistroEnElFormulario(registro, vehiculoAsociado) {
		document.getElementById("placa").value = vehiculoAsociado.Placas;
		document.getElementById("marca").value = vehiculoAsociado.Marca;
		document.getElementById("modelo").value = vehiculoAsociado.Modelo;
		document.getElementById("fechaDeLanzamiento").value = vehiculoAsociado.FechaLanzamiento;

		document.getElementById("fecha").value = registro.Fecha;
		console.log(`option[value="${registro.TipoRegistro}"]`);
		document.querySelector(`option[value="${registro.TipoRegistro}"]`).selected = true;
		document.getElementById("importe").value = registro.Importe;
		document.getElementById("kilometraje").value = registro.Kilometraje;

		if (registro.TipoRegistro === "carga de combustible") {
			document.getElementById("litrosGasolina").value = registro.LitrosGasolina;
		} else {
			document.getElementById("concepto").value = registro.Concepto;
		}
	}
})();
