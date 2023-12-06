(function () {
	const botonEnviarFormularioSeleccionarVehiculo = document.querySelector(".seleccionar_vehiculo__formulario-contenedor .formulario__submit");

	if (botonEnviarFormularioSeleccionarVehiculo) {
		botonEnviarFormularioSeleccionarVehiculo.addEventListener("click", (e) => {
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
			const seleccionarVehiculoParaNuevoRegistroURL = "/SeleccionarVehiculoParaNuevoRegistro";
			const peticion = await fetch(seleccionarVehiculoParaNuevoRegistroURL, {
				method: "POST",
				body: datosFormularioJSON,
			});
			const respuesta = await peticion.json();
			if (!respuesta.OK) {
				desplegarAlerta("error", respuesta.err);
			} else {
				desplegarAlerta("exito", "Vehículo encontrado exitosamente");
				setTimeout(() => {
					window.location.href = "/AgregarRegistroMantenimientoVehiculo?placas=" + datosFormulario.get("placas");
				}, 1000);
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
				showConfirmButton: tipoDeAlerta === "error",
			});
		}
	}
})();
