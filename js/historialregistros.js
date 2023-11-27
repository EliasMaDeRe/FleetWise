(function () {
	const tablaRegistros = cargarPluginDataTable();
	const filtros = {
		placas: "",
		marca: "",
		modelo: "",
		fechaDeLanzamiento: "",
		tipoDeRegistro: "",
	};

	document.addEventListener("DOMContentLoaded", (e) => {
		cargarPaginaHistorialDeRegistros();
	});

	const botonEnviarFiltros = document.getElementById("botonEnviarFiltros");

	botonEnviarFiltros.addEventListener("click", (e) => {
		e.preventDefault();

		cargarRegistrosConFiltros();
	});

	function cargarPluginDataTable() {
		let table = new DataTable("#tabla-registros", {
			responsive: true,
			language: {
				url: "//cdn.datatables.net/plug-ins/1.13.7/i18n/es-ES.json",
				emptyTable: "Ningún registro coincide con los filtros establecidos.",
			},
		});
		return table;
	}

	async function cargarPaginaHistorialDeRegistros() {
		const filtrosDeCamposDeVehiculo = {
			placas: [],
			marcas: [],
			modelos: [],
			fechasLanzamiento: [],
		};

		const registrosConVehiculosAsociadosExistentes = await obtenerRegistrosConVehiculosAsociadosFiltrados();

		const registrosExistentes = registrosConVehiculosAsociadosExistentes.registrosFiltrados;
		const vehiculosAsociados = registrosConVehiculosAsociadosExistentes.vehiculosFiltrados;

		Array.from(vehiculosAsociados).forEach((vehiculo) => {
			insertarOpcionDeFiltro(filtrosDeCamposDeVehiculo.placas, vehiculo.Placas);
			insertarOpcionDeFiltro(filtrosDeCamposDeVehiculo.marcas, vehiculo.Marca);
			insertarOpcionDeFiltro(filtrosDeCamposDeVehiculo.modelos, vehiculo.Modelo);
			insertarOpcionDeFiltro(filtrosDeCamposDeVehiculo.fechasLanzamiento, vehiculo.FechaLanzamiento);
		});

		cargarOpcionesDeFiltros(filtrosDeCamposDeVehiculo);
		cargarInformacionRegistros(registrosExistentes, vehiculosAsociados);
	}

	async function obtenerRegistrosConVehiculosAsociadosFiltrados() {
		const peticionUrl = "/HistorialRegistrosMantenimientoVehicular";

		const resultado = await fetch(peticionUrl, {
			method: "POST",
			body: JSON.stringify(filtros),
		});

		const respuesta = await resultado.json();

		return respuesta;
	}

	function insertarOpcionDeFiltro(arregloFiltros, valorExistente) {
		if (!arregloFiltros.includes(valorExistente)) {
			arregloFiltros.push(valorExistente);
		}
	}

	function cargarOpcionesDeFiltros(filtrosDeCamposDeVehiculo) {
		desplegarFiltroEnLaVista("placa", filtrosDeCamposDeVehiculo.placas);
		desplegarFiltroEnLaVista("marca", filtrosDeCamposDeVehiculo.marcas);
		desplegarFiltroEnLaVista("modelo", filtrosDeCamposDeVehiculo.modelos);
		desplegarFiltroEnLaVista("fechaDeLanzamiento", filtrosDeCamposDeVehiculo.fechasLanzamiento);
	}

	function desplegarFiltroEnLaVista(nombreFiltro, valoresFiltro) {
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

	function cargarInformacionRegistros(registros, vehiculos) {
		if (registros.length > 0) {
			desplegarRegistrosConVehiculosAsociadosFiltrados(registros, vehiculos);
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
				enlaceEditarRegistro.target = "_blank";

				let filaNueva = [
					registro.Fecha,
					vehiculo.Placas,
					vehiculo.Marca,
					vehiculo.Modelo,
					vehiculo.FechaLanzamiento,
					registro.TipoRegistro,
					registro.Kilometraje,
					registro.TipoRegistro == "carga de combustible" ? registro.LitrosGasolina : "N/A",
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

	async function cargarRegistrosConFiltros() {
		obtenerFiltros();

		const registrosConVehiculosAsociadosFiltrados = await obtenerRegistrosConVehiculosAsociadosFiltrados();

		const registrosFiltrados = registrosConVehiculosAsociadosFiltrados.registrosFiltrados;
		const vehiculosAsociadosFiltrados = registrosConVehiculosAsociadosFiltrados.vehiculosFiltrados;
		cargarInformacionRegistros(registrosFiltrados, vehiculosAsociadosFiltrados);
	}

	function obtenerFiltros() {
		filtros.placas = document.getElementById("placa").value;
		filtros.marca = document.getElementById("marca").value;
		filtros.modelo = document.getElementById("modelo").value;
		filtros.fechaDeLanzamiento = document.getElementById("fechaDeLanzamiento").value;
		filtros.tipoDeRegistro = document.getElementById("tipoDeRegistro").value;
	}
})();
