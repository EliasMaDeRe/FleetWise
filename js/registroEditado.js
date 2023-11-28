(function () {
	document.addEventListener("DOMContentLoaded", (e) => {
		cargarInformacionAlertaEdicionRegistro();
	});

	async function cargarInformacionAlertaEdicionRegistro() {
		const registroYVehiculoAsociado = await ObtenerDatosDelRegistro();

		if (respuestaVacia(registroYVehiculoAsociado.registro, registroYVehiculoAsociado.vehiculo)) {
			window.location.href = "/HistorialRegistrosMantenimientoVehicular";
			return;
		}

		mostrarDatosDelRegistroEnLaAlerta(registroYVehiculoAsociado.registro, registroYVehiculoAsociado.vehiculo);
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

	function mostrarDatosDelRegistroEnLaAlerta(registro, vehiculoAsociado) {
		mostrarDato("placa", vehiculoAsociado.Placas);
		mostrarDato("marca", vehiculoAsociado.Marca);
		mostrarDato("modelo", vehiculoAsociado.Modelo);
		mostrarDato("fechaDeLanzamiento", vehiculoAsociado.FechaLanzamiento);

		mostrarDato("fecha", registro.Fecha);

		mostrarDato("tipoRegistro", registro.TipoRegistro);
		mostrarDato("kilometraje", registro.Kilometraje);
		mostrarDato("importe", registro.Importe);
		mostrarDato("observaciones", registro.Observaciones);

		if (registro.TipoRegistro === "carga de combustible") {
			document.getElementById("campo-cambiante").innerHTML = `Litros: <span>${registro.LitrosGasolina}</span>`;
		} else {
			document.getElementById("campo-cambiante").innerHTML = `Concepto: <span>${registro.Concepto}</span>`;
		}
	}

	function mostrarDato(idDato, valorDato) {
		document.getElementById(idDato).textContent = valorDato;
	}
})();
