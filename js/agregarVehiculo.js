(function () {
	const botonEnviarFormularioCapturarVehiculo = document.querySelector(".captura-vehiculo__formulario-contenedor .formulario__submit");

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
			Swal.fire({
				title: tipoDeAlerta === "error" ? "¡Error!" : "¡Éxito!",
				text: mensajeAlerta,
				icon: tipoDeAlerta === "error" ? "error" : "success",
			}).then(() => {
				if (tipoDeAlerta !== "error") {
					window.location.reload();
				}
			});
		}
	}
})();
