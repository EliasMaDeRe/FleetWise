(function () {
	const tablaRegistros = cargarPluginDataTable();

	function cargarPluginDataTable() {
		let table = new DataTable("#tabla-flotillas", {
			responsive: true,
			language: {
				url: "//cdn.datatables.net/plug-ins/1.13.7/i18n/es-ES.json",
				emptyTable: "Ninguna flotilla registrada",
			},
			columnDefs: [{ width: "20%", targets: 0 }],
			fixedColumns: true,
			paging: false,
			scrollCollapse: true,
			scrollX: true,
			scrollY: 300,
		});
		return table;
	}
})();
