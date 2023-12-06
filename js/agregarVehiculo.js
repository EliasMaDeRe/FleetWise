function upload_file(e){
	e.preventDefault();
	archivoObj = e.dataTransfer.files[0];
	js_file_upload(archivoObj);
}
(function () {
	const botonEnviarFormularioCapturarVehiculo = document.querySelector(".captura-vehiculo__formulario-contenedor .formulario__submit");
	const areaImagen = document.querySelector(".area-imagen__contenedor");

	if(areaImagen){
		const imagenTexto = areaImagen.querySelector(".area-imagen__encabezado");
		const imagenBoton = areaImagen.querySelector(".area-imagen__boton");
		const imagenInput = areaImagen.querySelector(".area-imagen__input");

		let archivo;
		var archivoObj;

		imagenBoton.onclick = () =>{
			imagenInput.click();
			file_browse();
		}

		imagenInput.addEventListener("change",function(){
			archivo = this.files[0];
			areaImagen.classList.add("activo");
			showFile();
		})

		areaImagen.addEventListener("dragover",(event)=>{
			event.preventDefault();
			areaImagen.classList.add("activo");
			imagenTexto.textContent = "Suelta para subir imagen";
		})

		areaImagen.addEventListener("dragleave",()=>{
			areaImagen.classList.remove("activo");
			imagenTexto.textContent = "Arrastra y suelta para subir imagen";
		})

		areaImagen.addEventListener("drop",(event)=>{
			event.preventDefault();
			archivo = event.dataTransfer.files[0];
			showFile();
		})

		function showFile(){
			let tipoArchivo = archivo.type;
			let extensionesValidas = ["image/jpeg","image/jpg","image/png"];
			if(extensionesValidas.includes(tipoArchivo)){
				let fileReader = new FileReader();
				fileReader.onload = ()=>{
					let fileURL = fileReader.result;
					let imgTag = `<img src="${fileURL}" alt ="">`;
					areaImagen.innerHTML = imgTag;
				}
				fileReader.readAsDataURL(archivo);
			}else{
				alert("Tipo de archivo incorrecto");
				areaImagen.classList.remove("activo");
				imagenTexto.textContent = "Arrastra y suelta para subir imagen";
			}

		}

	}

	function file_browse(){
		document.getElementById('file').onchange = function(){
			archivoObj = document.getElementById('file').files[0];
			js_file_upload(archivoObj);
		}
	}

	function js_file_upload(archivo_obj){
		
		if( archivo_obj != undefined ){
			var form_data = new FormData();
			form_data.append('archivo',archivo_obj);
			var xhttp = new XMLHttpRequest();
			xhttp.open("POST","/SubirImagen",true);
			xhttp.onload = function (event) {
				if(xhttp.status == 200){
					desplegarAlerta("exito","Imagen subida con éxito");
				}else{
					desplegarAlerta("error",xhttp.response.err);
				}
			}
			xhttp.send(form_data);
		}
	}

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
