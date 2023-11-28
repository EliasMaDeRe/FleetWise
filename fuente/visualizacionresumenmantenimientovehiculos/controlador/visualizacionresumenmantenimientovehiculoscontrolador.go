package controlador

import (
	conectorBDControlador "example/fleetwise/fuente/conectorbd/controlador"
	visualizacionHistorialRegistrosControlador "example/fleetwise/fuente/visualizacionhistorialregistrosmantenimientovehiculo/controlador"
	capturaRegistroMantenimientoVehiculoModelos "example/fleetwise/modelos/capturaregistromantenimientovehiculo"
	capturaVehiculosModelos "example/fleetwise/modelos/capturavehiculos"
	visualizacionHistorialRegistroMantenimientoVehiculosModelos "example/fleetwise/modelos/visualizacionhistorialregistrosmantenimientovehiculo"
	"sort"
	"time"
)

type Controlador struct {
	ConectorBDControlador     *conectorBDControlador.Controlador
	VisualizacionHistorialRegistrosControlador *visualizacionHistorialRegistrosControlador.Controlador
}

func (c *Controlador) ObtenerMetricasVehiculares() ([]capturaVehiculosModelos.Vehiculo, []float64, []float64, []float64, []int, []int, []int, []float64){
	registrosMantenimientoVehiculo, vehiculos := c.VisualizacionHistorialRegistrosControlador.ObtenerRegistrosFiltradosConVehiculos(&visualizacionHistorialRegistroMantenimientoVehiculosModelos.ObtenerRegistrosFiltradosConVehiculosSolicitud{});

	vehiculosSinRepetirse := c.filtrarVehiculosRepetidos(vehiculos)

	gastosDeCombustiblePorVehiculo := []float64{}
	gastosEnMantenimientoPorVehiculo := []float64{}
	rendimientoKilometroLitroPorVehiculo := []float64{}
	kilometrajeInicialPorVehiculo := []int{}
	ultimoKilometrajePorVehiculo := []int{}
	kilometrosTotalesRecorridosPorVehiculo := []int{}
	kilometrosPromedioDiariosRecorridosPorVehiculo := []float64{}
	
	
	for _, vehiculo := range vehiculosSinRepetirse{
		registrosDelVehiculo := c.obtenerRegistrosPorPlacaOrdenadosPorFecha(vehiculo.ObtenerPlacas(), registrosMantenimientoVehiculo)
		
		gastosDeCombustiblePorVehiculo = append(gastosDeCombustiblePorVehiculo, c.calcularGastoDeCombustible(*registrosDelVehiculo))

		gastosEnMantenimientoPorVehiculo = append(gastosEnMantenimientoPorVehiculo, c.calcularGastoEnMantenimiento(*registrosDelVehiculo))

		rendimientoKilometroLitroPorVehiculo = append(rendimientoKilometroLitroPorVehiculo, c.calcularRendimientoKilometrosLitros(*registrosDelVehiculo))

		kilometrajeInicialPorVehiculo = append(kilometrajeInicialPorVehiculo, c.obtenerKilometrajeInicial(*registrosDelVehiculo))

		ultimoKilometrajePorVehiculo = append(ultimoKilometrajePorVehiculo, c.obtenerUltimoKilometraje(*registrosDelVehiculo))

		kilometrosTotalesRecorridosPorVehiculo = append(kilometrosTotalesRecorridosPorVehiculo, c.calcularKilometrosTotalesRecorridos(*registrosDelVehiculo))
		
		kilometrosPromedioDiariosRecorridosPorVehiculo = append(kilometrosPromedioDiariosRecorridosPorVehiculo, c.calcularKilometrosPromedioDiariosRecorridos(*registrosDelVehiculo))
	}

	return vehiculosSinRepetirse, gastosDeCombustiblePorVehiculo, gastosEnMantenimientoPorVehiculo, rendimientoKilometroLitroPorVehiculo, kilometrajeInicialPorVehiculo, ultimoKilometrajePorVehiculo, kilometrosTotalesRecorridosPorVehiculo, kilometrosPromedioDiariosRecorridosPorVehiculo
}
 
func (c *Controlador) obtenerRegistrosPorPlacaOrdenadosPorFecha(placa string, registros []capturaRegistroMantenimientoVehiculoModelos.RegistroMantenimientoVehiculo) *[]capturaRegistroMantenimientoVehiculoModelos.RegistroMantenimientoVehiculo{
	registrosFiltrados := []capturaRegistroMantenimientoVehiculoModelos.RegistroMantenimientoVehiculo{}

	for _,registroMantenimientoVehiculo := range registros{
		if(registroMantenimientoVehiculo.ObtenerPlacasVehiculo() == placa){
			registrosFiltrados = append(registrosFiltrados, registroMantenimientoVehiculo)
		}
	}

	sort.Slice(registrosFiltrados, func(i,j int) bool{
		return registrosFiltrados[i].ObtenerFecha() < registrosFiltrados[j].ObtenerFecha()
	})
	
	return &registrosFiltrados
}

func (c *Controlador) filtrarVehiculosRepetidos(vehiculos []capturaVehiculosModelos.Vehiculo) []capturaVehiculosModelos.Vehiculo{
	vehiculosEncontrados := map[capturaVehiculosModelos.Vehiculo]bool{}
	vehiculosNoRepetidos := []capturaVehiculosModelos.Vehiculo{}

	for _, vehiculo := range vehiculos {
		if !vehiculosEncontrados[vehiculo] {
			vehiculosEncontrados[vehiculo] = true
			vehiculosNoRepetidos = append(vehiculosNoRepetidos, vehiculo)
		}
	}

	return vehiculosNoRepetidos
}

func (c *Controlador) calcularGastoDeCombustible(registros []capturaRegistroMantenimientoVehiculoModelos.RegistroMantenimientoVehiculo) float64{
	var gastoCombustibleTotal float64 = 0

	for _, registro := range registros{
		if(registro.ObtenerTipo() == "carga de combustible"){
			gastoCombustibleTotal += registro.ObtenerImporte()
		}
	}

	return gastoCombustibleTotal
}

func (c *Controlador) calcularGastoEnMantenimiento(registros []capturaRegistroMantenimientoVehiculoModelos.RegistroMantenimientoVehiculo) float64{
	var gastoMantenimientoTotal float64 = 0

	for _, registro := range registros{
		if(registro.ObtenerTipo() != "carga de combustible"){
			gastoMantenimientoTotal += registro.ObtenerImporte()
		}
	}

	return gastoMantenimientoTotal
}

func (c *Controlador) calcularRendimientoKilometrosLitros(registros []capturaRegistroMantenimientoVehiculoModelos.RegistroMantenimientoVehiculo) float64{
	var litrosTotales float64 = 0

	for _, registro := range registros {
		if(registro.ObtenerTipo() == "carga de combustible"){
			litrosTotales += registro.ObtenerLitrosDeGasolina()
		}
	}
	
	rendimientoKilometrosLitros := float64(c.calcularKilometrosTotalesRecorridos(registros)) / litrosTotales

	return rendimientoKilometrosLitros
}

func (c *Controlador) obtenerKilometrajeInicial(registros []capturaRegistroMantenimientoVehiculoModelos.RegistroMantenimientoVehiculo) int{
	var kilometrajeInicial int = 0

	primerRegistro := registros[0]
	kilometrajeInicial = primerRegistro.ObtenerKilometraje()

	return kilometrajeInicial
}

func (c *Controlador) obtenerUltimoKilometraje(registros []capturaRegistroMantenimientoVehiculoModelos.RegistroMantenimientoVehiculo) int{
	var ultimoKilometraje int = 0

	ultimoIndice := len(registros) - 1
	ultimoRegistro := registros[ultimoIndice]
	ultimoKilometraje = ultimoRegistro.ObtenerKilometraje()

	return ultimoKilometraje
}

func (c *Controlador) calcularKilometrosTotalesRecorridos(registros []capturaRegistroMantenimientoVehiculoModelos.RegistroMantenimientoVehiculo) int{
	var kilometrosTotalesRecorridos int = 0

	var kilometrajeInicial = c.obtenerKilometrajeInicial(registros)
	var ultimoKilometraje = c.obtenerUltimoKilometraje(registros)

	kilometrosTotalesRecorridos = ultimoKilometraje - kilometrajeInicial

	return kilometrosTotalesRecorridos
}

func (c *Controlador) calcularKilometrosPromedioDiariosRecorridos(registros []capturaRegistroMantenimientoVehiculoModelos.RegistroMantenimientoVehiculo) float64{
	formatoFecha:= "2006-01-02"

	primerRegistro := registros[0]
	ultimoRegistro := registros[len(registros)-1]

	fechaInicial, _ := time.Parse(formatoFecha, primerRegistro.ObtenerFecha())
	fechaFinal, _ := time.Parse(formatoFecha, ultimoRegistro.ObtenerFecha())

	diasTotales := fechaFinal.Sub(fechaInicial).Hours() / 24;

	kilometroTotalesRecorridos := ultimoRegistro.ObtenerKilometraje() - primerRegistro.ObtenerKilometraje()

	promedioKilometrosDiariosRecorridos := float64(kilometroTotalesRecorridos) / diasTotales

	return promedioKilometrosDiariosRecorridos

}