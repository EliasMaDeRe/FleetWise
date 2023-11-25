(function () {
    cargarPluginDataTable();

    document.addEventListener("DOMContentLoaded", (e) => {
        cargarCamposDeFiltrosDeVehiculo();

        const botonEnviarFiltros = document.getElementById("botonEnviarFiltros");
        botonEnviarFiltros.addEventListener("click", (e) => {});
    });

    function cargarPluginDataTable() {
        let table = new DataTable("#tabla-registros", {
            responsive: true,
            language: {
                url: "//cdn.datatables.net/plug-ins/1.13.7/i18n/es-ES.json",
            },
        });
    }

    async function cargarCamposDeFiltrosDeVehiculo() {
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
        const urlPeticion = "/HistorialRegistrosServicioVehicular";

        const respuesta = await fetch(urlPeticion, { method: "POST" });

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
})();
