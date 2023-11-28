(function () {
	const tablaResumen = cargarPluginDataTable();

	document.addEventListener("DOMContentLoaded", (e) => {
		cargarPagina();
	});

	function cargarPluginDataTable() {
		console.log("Visca Barça");

		let table = new DataTable("#tabla-resumen", {
			responsive: true,
			scrollX: true,
			language: {
				url: "//cdn.datatables.net/plug-ins/1.13.7/i18n/es-ES.json",
				emptyTable: "Ningún registro coincide con los filtros establecidos.",
			},
		});
		return table;
	}

	async function cargarPagina() {
		urlResumenesFlotillas = "/ObtenerMetricasVehiculos";

		const respuesta = await fetch(urlResumenesFlotillas, { method: "POST" });

		const resultado = await respuesta.json();

		desplegarInformacionEnLaPagina(resultado);
	}

	async function desplegarInformacionEnLaPagina(resumenesFlotillas) {
		const {
			vehiculos,
			gastosDeCombustiblePorVehiculo,
			gastosEnMantenimientoPorVehiculo,
			kilometrosPromedioDiariosRecorridosPorVehiculo,
			kilometrosTotalesRecorridosPorVehiculo,
			rendimientosKilometroLitroPorVehiculo,
			ultimosKilometrajesPorVehiculo,
			kilometrajesInicialesPorVehiculo,
		} = resumenesFlotillas;

		for (let indiceVehiculo = 0; indiceVehiculo < vehiculos.length; indiceVehiculo++) {
			vehiculo = vehiculos[indiceVehiculo];

			const accionesContenedor = document.createElement("DIV");
			accionesContenedor.classList.add("resumen__acciones");

			const enlaceEditarVehiculo = document.createElement("A");
			enlaceEditarVehiculo.innerText = "Editar";
			enlaceEditarVehiculo.href = `/EditarVehiculo?placas=${vehiculo.Placas}`;
			enlaceEditarVehiculo.target = "_blank";

			const enlaceEliminarVehiculo = document.createElement("A");
			enlaceEliminarVehiculo.innerText = "Eliminar";
			enlaceEliminarVehiculo.href = `/EliminarVehiculo?placas=${vehiculo.Placas}`;

			accionesContenedor.appendChild(enlaceEditarVehiculo);
			accionesContenedor.appendChild(enlaceEliminarVehiculo);

			const resumenFlotilla = [
				vehiculo.Placas,
				vehiculo.Marca,
				vehiculo.Modelo,
				vehiculo.FechaLanzamiento,
				gastosDeCombustiblePorVehiculo[indiceVehiculo] == 0 ? "N/A" : gastosDeCombustiblePorVehiculo[indiceVehiculo],
				gastosEnMantenimientoPorVehiculo[indiceVehiculo] == 0 ? "N/A" : gastosEnMantenimientoPorVehiculo[indiceVehiculo],
				rendimientosKilometroLitroPorVehiculo[indiceVehiculo] == 0 ? "N/A" : rendimientosKilometroLitroPorVehiculo[indiceVehiculo],
				kilometrajesInicialesPorVehiculo[indiceVehiculo],
				ultimosKilometrajesPorVehiculo[indiceVehiculo],
				kilometrosTotalesRecorridosPorVehiculo[indiceVehiculo] == 0 ? "N/A" : kilometrosTotalesRecorridosPorVehiculo[indiceVehiculo],
				kilometrosPromedioDiariosRecorridosPorVehiculo[indiceVehiculo] == 0 ? "N/A" : kilometrosPromedioDiariosRecorridosPorVehiculo[indiceVehiculo],
				accionesContenedor.outerHTML,
			];

			tablaResumen.row.add(resumenFlotilla);
		}
		tablaResumen.draw();
	}
})();
