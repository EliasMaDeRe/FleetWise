(function () {

	const tablaResumen = cargarPluginDataTable();

	function cargarPluginDataTable() {

        console.log("Visca Barça");
        
		let table = new DataTable("#tabla-resumen", {
			responsive: true,
			language: {
				url: "//cdn.datatables.net/plug-ins/1.13.7/i18n/es-ES.json",
				emptyTable: "Ningún registro coincide con los filtros establecidos.",
			},
		});
		return table;
	}

    // Este es el archivo bueno jajsja Visca Barça

})();
