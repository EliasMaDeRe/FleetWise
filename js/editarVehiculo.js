(function () {
	document.addEventListener("DOMContentLoaded", () => {
		cargarInformacionVehiculoEnElFormulario();
	});

	document.getElementById("editar-vehiculo").addEventListener("click", (e) => {
		e.preventDefault();

		const inputsFormulario = obtenerInputsDelFormulario();
		enviarFormulario(inputsFormulario);
	});

	function obtenerPlacas() {
		const params = new URLSearchParams(window.location.search);
		return params.get("placas");
	}

	async function cargarInformacionVehiculoEnElFormulario() {
		const urlPeticion = "/ObtenerVehiculoPorPlacas";
		const cuerpoPeticion = {
			placas: obtenerPlacas(),
		};

		const respuesta = await fetch(urlPeticion, { method: "POST", body: JSON.stringify(cuerpoPeticion) });

		const resultado = await respuesta.json();

		desplegarDatosVehiculo(resultado.Vehiculo);
	}

	async function desplegarDatosVehiculo(vehiculo) {
		document.getElementById("marca").value = vehiculo.Marca;
		document.getElementById("modelo").value = vehiculo.Modelo;
		document.getElementById("fechalanzamiento").value = vehiculo.FechaLanzamiento;
		document.getElementById("serie").value = vehiculo.Serie;
		document.getElementById("placas").value = obtenerPlacas();
	}

	function obtenerInputsDelFormulario() {
		return document.querySelectorAll(".formulario__input");
	}

	async function enviarFormulario(inputsFormulario) {
		const datosFormulario = construirPeticionFormulario(inputsFormulario);
		const editarVehiculoURL = "/EditarVehiculo";
		const peticion = await fetch(editarVehiculoURL, {
			method: "POST",
			body: datosFormulario,
		});
		const respuesta = await peticion.json();
		if (!respuesta.OK) {
			desplegarAlerta("error", respuesta.err);
		} else {
			desplegarAlerta("exito", "Vehículo agregado exitosamente");
		}
	}

	function construirPeticionFormulario() {
		const datosFormulario = {
			placasactuales: obtenerPlacas(),
			placasnuevas: document.getElementById("placas").value,
			fechalanzamientonueva: document.getElementById("fechalanzamiento").value,
			marcanueva: document.getElementById("marca").value,
			modelonuevo: document.getElementById("modelo").value,
			serienuevo: document.getElementById("serie").value,
		};
		return JSON.stringify(datosFormulario);
	}

	function desplegarAlerta(tipoDeAlerta, mensajeAlerta) {
		Swal.fire({
			title: tipoDeAlerta === "error" ? "Error!" : "Éxito",
			text: mensajeAlerta,
			icon: tipoDeAlerta === "error" ? "error" : "success",
		}).then(() => {
			if (tipoDeAlerta !== "error") {
				window.location.href = "/ResumenMantenimientoVehiculos";
			}
		});
	}
})();
