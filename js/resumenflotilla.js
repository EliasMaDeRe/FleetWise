(function () {
	const tablaResumen = cargarPluginDataTable();

	document.addEventListener("DOMContentLoaded", (e) => {
		cargarPagina();
	});

	window.addEventListener("hashchange", () => {
		desplegarVentanaEliminarVehiculo();
	});

	function cargarPluginDataTable() {
		let table = new DataTable("#tabla-resumen", {
			responsive: true,
			language: {
				url: "//cdn.datatables.net/plug-ins/1.13.7/i18n/es-ES.json",
				emptyTable: "Ningún registro coincide con los filtros establecidos.",
			},
			exportOptions: {
				columns: ":visible",
			},
			dom: "Bfrtip",
			buttons: ["print"],
		});
		return table;
	}

	async function cargarPagina() {
		urlResumenesFlotillas = "/ObtenerMetricasVehiculos";

		const respuesta = await fetch(urlResumenesFlotillas, { method: "POST" });

		const resultado = await respuesta.json();

		await desplegarInformacionEnLaPagina(resultado);

		if (window.location.hash) {
			desplegarVentanaEliminarVehiculo();
		}
	}

	async function desplegarInformacionEnLaPagina(resumenFlotilla) {
		const {
			vehiculos,
			gastosDeCombustiblePorVehiculo,
			gastosEnMantenimientoPorVehiculo,
			kilometrosPromedioDiariosRecorridosPorVehiculo,
			kilometrosTotalesRecorridosPorVehiculo,
			rendimientosKilometroLitroPorVehiculo,
			ultimosKilometrajesPorVehiculo,
			kilometrajesInicialesPorVehiculo,
		} = resumenFlotilla;

		for (let indiceVehiculo = 0; indiceVehiculo < vehiculos.length; indiceVehiculo++) {
			vehiculo = vehiculos[indiceVehiculo];

			const accionesContenedor = document.createElement("DIV");
			accionesContenedor.classList.add("resumen__acciones");

			const enlaceEditarVehiculo = construirEnlaceEditarVehiculo(vehiculo.Placas);

			const enlaceEliminarVehiculo = construirEnlaceEliminarVehiculo(vehiculo.Placas);

			accionesContenedor.appendChild(enlaceEditarVehiculo);
			accionesContenedor.appendChild(enlaceEliminarVehiculo);

			const resumenFlotilla = [
				vehiculo.Placas,
				vehiculo.Marca,
				vehiculo.Modelo,
				vehiculo.FechaLanzamiento,
				vehiculo.Serie,
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

		prepararBotonExportarAExcel();
		tablaResumen.draw();
	}

	function prepararBotonExportarAExcel(){
		document.getElementById('ExportExcelButton').addEventListener('click', () => {
			fetch('/ExportarFlotilla?Formato=excel')
			  .then(response => {
				if (!response.ok) throw new Error('Error en la descarga');
				return response.blob(); 
			  })
			  .then(blob => {
				const url = window.URL.createObjectURL(blob);
				const a = document.createElement('a');
				a.href = url;
				a.download = 'ResumenFlotilla.xlsx'; 
				document.body.appendChild(a);
				a.click();
				a.remove();
				window.URL.revokeObjectURL(url);
			  })
			  .catch(error => {
				console.error('Error:', error);
			  });
		  });
	}

	function construirEnlaceEditarVehiculo(placasVehiculo) {
		const enlaceEditarVehiculo = document.createElement("A");
		enlaceEditarVehiculo.innerText = "Editar";
		enlaceEditarVehiculo.href = `/EditarVehiculo?placas=${placasVehiculo}`;
		enlaceEditarVehiculo.target = "_blank";

		return enlaceEditarVehiculo;
	}

	function construirEnlaceEliminarVehiculo(placasVehiculo) {
		const enlaceEliminarVehiculo = document.createElement("A");
		enlaceEliminarVehiculo.innerText = "Eliminar";
		enlaceEliminarVehiculo.href = "#eliminar:" + placasVehiculo;

		return enlaceEliminarVehiculo;
	}

	function desplegarVentanaEliminarVehiculo() {
		if (!window.location.hash.includes("#eliminar")) {
			return;
		}
		const placaVehiculo = window.location.hash.replace("#eliminar:", "");

		if (!document.querySelector(`a[href='#eliminar:${placaVehiculo}']`)) {
			return;
		}

		Swal.fire({
			title: "¿Deseas eliminar el vehiculo con placa '" + placaVehiculo + "' ?",
			showCancelButton: true,
			confirmButtonText: "Eliminar",
			denyButtonText: `Cancelar`,
		}).then((result) => {
			if (result.isConfirmed) {
				eliminarVehiculo(placaVehiculo);
			} else {
				recargarPagina();
			}
		});
	}

	async function eliminarVehiculo(placa) {
		const urlEliminarVehiculo = "/EliminarVehiculo";
		const peticionEliminarVehiculo = {
			placas: placa,
		};

		const resultado = await fetch(urlEliminarVehiculo, {
			method: "DELETE",
			body: JSON.stringify(peticionEliminarVehiculo),
		});

		const respuesta = await resultado.json();

		if (respuesta.OK) {
			Swal.fire({
				title: "Vehiculo Eliminado",
				text: "Vehiculo eliminado correctamente",
				icon: "success",
			}).then(() => {
				recargarPagina();
			});
		} else {
			Swal.fire({
				title: "Error!",
				text: "Ha ocurrido un error al intentar eliminar el vehiculo, por favor inténtelo mas tarde",
				icon: "error",
			}).then(() => {
				recargarPagina();
			});
		}
	}

	function recargarPagina() {
		window.location.href = window.location.href.replace(window.location.hash, "");
	}
})();
