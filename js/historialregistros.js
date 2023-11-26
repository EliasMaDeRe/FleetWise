(function () {
	const tablaRegistros = new DataTable("#tabla-registros", {
		responsive: true,
		language: {
			url: "//cdn.datatables.net/plug-ins/1.13.7/i18n/es-ES.json",
			emptyTable: "Ningun registro coincide con los filtros seleccionados",
		},
	});
	const filtros = {
		placas: "",
		marca: "",
		modelo: "",
		fechaDeLanzamiento: "",
		tipoDeRegistro: "",
	};

	document.addEventListener("DOMContentLoaded", (e) => {
		cargarOpcionesDeFiltros();
		obtenerRegistrosConFiltros(filtros);
	});

	const botonEnviarFiltros = document.getElementById("botonEnviarFiltros");

	botonEnviarFiltros.addEventListener("click", (e) => {
		e.preventDefault();

		obtenerFiltros();

		obtenerRegistrosConFiltros(filtros);
	});

	function cargarPluginDataTable() {
		let table = new DataTable("#tabla-registros", {
			responsive: true,
			language: {
				url: "//cdn.datatables.net/plug-ins/1.13.7/i18n/es-ES.json",
				emptyTable: "Seleccione los filtros deseados y de click en enviar para obtener los registros.",
			},
		});
		return table;
	}

	async function cargarOpcionesDeFiltros() {
		const filtrosDeCamposDeVehiculo = {
			placas: [],
			marcas: [],
			modelos: [],
			fechasLanzamiento: [],
		};

		const vehiculosExistentes = await obtenerVehiculosExistentes();

		Array.from(vehiculosExistentes).forEach((vehiculo) => {
			insertarFiltroNuevo(filtrosDeCamposDeVehiculo.placas, vehiculo.Placas);
			insertarFiltroNuevo(filtrosDeCamposDeVehiculo.marcas, vehiculo.Marca);
			insertarFiltroNuevo(filtrosDeCamposDeVehiculo.modelos, vehiculo.Modelo);
			insertarFiltroNuevo(filtrosDeCamposDeVehiculo.fechasLanzamiento, vehiculo.FechaLanzamiento);
		});

		cargarFiltrosEnLaVista(filtrosDeCamposDeVehiculo);
	}

	async function obtenerVehiculosExistentes() {
		const urlPeticion = "/HistorialRegistrosMantenimientoVehicular";

		const respuesta = await fetch(urlPeticion, {
			method: "POST",
			body: JSON.stringify(filtros),
		});

		const registrosYVehiculos = await respuesta.json();

		return registrosYVehiculos.vehiculosFiltrados;
	}

	function insertarFiltroNuevo(arregloFiltros, valorExistente) {
		if (!arregloFiltros.includes(valorExistente)) {
			arregloFiltros.push(valorExistente);
		}
	}

	function cargarFiltrosEnLaVista(filtrosDeCamposDeVehiculo) {
		cargarFiltro("placa", filtrosDeCamposDeVehiculo.placas);
		cargarFiltro("marca", filtrosDeCamposDeVehiculo.marcas);
		cargarFiltro("modelo", filtrosDeCamposDeVehiculo.modelos);
		cargarFiltro("fechaDeLanzamiento", filtrosDeCamposDeVehiculo.fechasLanzamiento);
	}

	function cargarFiltro(nombreFiltro, valoresFiltro) {
		const selectFiltro = document.getElementById(nombreFiltro);
		selectFiltro.innerHTML = "";

		const optionVacio = document.createElement("OPTION");
		optionVacio.value = "";
		optionVacio.innerText = "No Filtrar";

		selectFiltro.appendChild(optionVacio);

		valoresFiltro.forEach((valorFiltro) => {
			const optionValorFiltro = document.createElement("OPTION");
			optionValorFiltro.value = valorFiltro;
			optionValorFiltro.innerText = valorFiltro;
			selectFiltro.appendChild(optionValorFiltro);
		});
	}

	function obtenerFiltros() {
		filtros.placas = document.getElementById("placa").value;
		filtros.marca = document.getElementById("marca").value;
		filtros.modelo = document.getElementById("modelo").value;
		filtros.fechaDeLanzamiento = document.getElementById("fechaDeLanzamiento").value;
		filtros.tipoDeRegistro = document.getElementById("tipoDeRegistro").value;
	}

	async function obtenerRegistrosConFiltros(filtros) {
		const peticionUrl = "/HistorialRegistrosMantenimientoVehicular";

		const resultado = await fetch(peticionUrl, {
			method: "POST",
			body: JSON.stringify(filtros),
		});

		const respuesta = await resultado.json();

		if (respuesta.registrosFiltrados.length > 0) {
			desplegarRegistrosConVehiculosAsociadosFiltrados(respuesta.registrosFiltrados, respuesta.vehiculosFiltrados);
		} else {
			desplegarMensajeVacioTablaRegistros();
		}
	}

	function desplegarRegistrosConVehiculosAsociadosFiltrados(registrosFiltrados, vehiculosFiltrados) {
		tablaRegistros.clear();

		let cantidadRegistros = registrosFiltrados.length;
		if (cantidadRegistros > 0) {
			for (let i = 0; i < cantidadRegistros; i++) {
				let registro = registrosFiltrados[i];
				let vehiculo = vehiculosFiltrados[i];

				let enlaceEditarRegistro = document.createElement("A");
				enlaceEditarRegistro.innerText = "Editar";
				enlaceEditarRegistro.href = `/EditarRegistroMantenimientoVehicular?numRegistro=${registro.NumeroDeRegistro}`;
				console.log(enlaceEditarRegistro.innerHTML);
				let filaNueva = [
					registro.Fecha,
					vehiculo.Placas,
					vehiculo.Marca,
					vehiculo.Modelo,
					vehiculo.FechaLanzamiento,
					registro.TipoRegistro,
					registro.Kilometraje,
					registro.TipoRegistro == "Carga de Combustible" ? registro.LitrosGasolina : "N/A",
					registro.Importe,
					registro.Concepto,
					enlaceEditarRegistro.outerHTML,
				];

				tablaRegistros.row.add(filaNueva).draw();
			}
		}
	}

	function desplegarMensajeVacioTablaRegistros() {
		tablaRegistros.clear().draw();
	}
})();
