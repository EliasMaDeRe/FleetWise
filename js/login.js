(function () {
	const botonEnviarFormularioLogin = document.querySelector(".login__submit");

	if (botonEnviarFormularioLogin) {
		botonEnviarFormularioLogin.addEventListener("click", (e) => {
			e.preventDefault();
			const inputsFormulario = obtenerInputsDelFormulario();
			enviarFormulario(inputsFormulario);
		});

		function obtenerInputsDelFormulario() {
			return document.querySelectorAll(".login__input");
		}

		async function enviarFormulario(inputsFormulario) {
			const datosFormulario = construirPeticionFormulario(inputsFormulario);
			const datosFormularioJSON = JSON.stringify(Object.fromEntries(datosFormulario));
			const peticion = await fetch("/Login", {
				method: "POST",
				body: datosFormularioJSON,
			});
			const respuesta = await peticion.json();
			if (!respuesta.OK) {
				desplegarAlerta("error", "Acceso no autorizado");
			} else {
				desplegarAlerta("exito", "Sesion iniciada con exito");
				setTimeout(async () => {
					window.location.href = "/";
				}, 1500);
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
				title: tipoDeAlerta === "error" ? "Error!" : "Éxito",
				text: mensajeAlerta,
				icon: tipoDeAlerta === "error" ? "error" : "success",
				showConfirmButton: tipoDeAlerta === "error",
			});
		}
	}
})();
