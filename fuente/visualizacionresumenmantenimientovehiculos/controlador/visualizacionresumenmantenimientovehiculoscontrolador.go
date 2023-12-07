package controlador

import (
	conectorBDControlador "example/fleetwise/fuente/conectorbd/controlador"
	visualizacionHistorialRegistrosControlador "example/fleetwise/fuente/visualizacionhistorialregistrosmantenimientovehiculo/controlador"
	"example/fleetwise/fuente/visualizacionresumenmantenimientovehiculos/constantes"
	capturaRegistroMantenimientoVehiculoModelos "example/fleetwise/modelos/capturaregistromantenimientovehiculo"
	capturaVehiculosModelos "example/fleetwise/modelos/capturavehiculos"
	"fmt"
	"math"
	"sort"
	"strings"
	"time"
)

type Controlador struct {
	ConectorBDControlador                      *conectorBDControlador.Controlador
	VisualizacionHistorialRegistrosControlador *visualizacionHistorialRegistrosControlador.Controlador
}

func (c *Controlador) ObtenerMetricasVehiculares() ([]capturaVehiculosModelos.Vehiculo, []float64, []float64, []float64, []int, []int, []int, []float64) {
	registrosMantenimientoVehiculo, vehiculos := c.ConectorBDControlador.ObtenerRegistrosYVehiculosAsociados()

	vehiculosSinRepetirse := c.filtrarVehiculosRepetidos(vehiculos)

	gastosDeCombustiblePorVehiculo := []float64{}
	gastosEnMantenimientoPorVehiculo := []float64{}
	rendimientoKilometroLitroPorVehiculo := []float64{}
	kilometrajeInicialPorVehiculo := []int{}
	ultimoKilometrajePorVehiculo := []int{}
	kilometrosTotalesRecorridosPorVehiculo := []int{}
	kilometrosPromedioDiariosRecorridosPorVehiculo := []float64{}
	for _, vehiculo := range vehiculosSinRepetirse {

		registrosDelVehiculo := c.obtenerRegistrosPorPlacaOrdenadosPorFecha(vehiculo.ObtenerPlacas(), registrosMantenimientoVehiculo)
		fmt.Println(registrosDelVehiculo)
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

func (c *Controlador) obtenerRegistrosPorPlacaOrdenadosPorFecha(placa string, registros []capturaRegistroMantenimientoVehiculoModelos.RegistroDeMantenimientoDeVehiculo) *[]capturaRegistroMantenimientoVehiculoModelos.RegistroDeMantenimientoDeVehiculo {
	registrosFiltrados := []capturaRegistroMantenimientoVehiculoModelos.RegistroDeMantenimientoDeVehiculo{}

	for _, registroMantenimientoVehiculo := range registros {
		if strings.EqualFold(registroMantenimientoVehiculo.ObtenerPlacasVehiculo(), placa) {

			registrosFiltrados = append(registrosFiltrados, registroMantenimientoVehiculo)
		}
	}

	sort.Slice(registrosFiltrados, func(i, j int) bool {
		return registrosFiltrados[i].ObtenerFecha() < registrosFiltrados[j].ObtenerFecha()
	})
	return &registrosFiltrados
}

func (c *Controlador) filtrarVehiculosRepetidos(vehiculos []capturaVehiculosModelos.Vehiculo) []capturaVehiculosModelos.Vehiculo {
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

func (c *Controlador) calcularGastoDeCombustible(registros []capturaRegistroMantenimientoVehiculoModelos.RegistroDeMantenimientoDeVehiculo) float64 {
	var gastoCombustibleTotal float64 = 0

	for _, registro := range registros {
		fmt.Println(registro.ObtenerTipo())
		if registro.ObtenerTipo() == constantes.TIPO_CARGA_DE_COMBUSTIBLE {
			gastoCombustibleTotal += registro.ObtenerImporte()
		}
	}

	return gastoCombustibleTotal
}

func (c *Controlador) calcularGastoEnMantenimiento(registros []capturaRegistroMantenimientoVehiculoModelos.RegistroDeMantenimientoDeVehiculo) float64 {
	var gastoMantenimientoTotal float64 = 0

	for _, registro := range registros {
		if registro.ObtenerTipo() != constantes.TIPO_CARGA_DE_COMBUSTIBLE {
			gastoMantenimientoTotal += registro.ObtenerImporte()
		}
	}

	return gastoMantenimientoTotal
}

func (c *Controlador) calcularRendimientoKilometrosLitros(registros []capturaRegistroMantenimientoVehiculoModelos.RegistroDeMantenimientoDeVehiculo) float64 {
	var litrosTotales float64 = 0

	if len(registros) == 1 {
		fmt.Println(litrosTotales)
		return litrosTotales
	}

	for _, registro := range registros {
		if registro.ObtenerTipo() == constantes.TIPO_CARGA_DE_COMBUSTIBLE {
			litrosTotales += registro.ObtenerLitrosDeGasolina()
		}
	}

	rendimientoKilometrosLitros := float64(c.calcularKilometrosTotalesRecorridos(registros)) / litrosTotales
	if math.IsNaN(rendimientoKilometrosLitros) {
		rendimientoKilometrosLitros = 0
	}
	fmt.Println(rendimientoKilometrosLitros)
	return rendimientoKilometrosLitros
}

func (c *Controlador) obtenerKilometrajeInicial(registros []capturaRegistroMantenimientoVehiculoModelos.RegistroDeMantenimientoDeVehiculo) int {

	primerRegistro := registros[0]
	return primerRegistro.ObtenerKilometraje()
}

func (c *Controlador) obtenerUltimoKilometraje(registros []capturaRegistroMantenimientoVehiculoModelos.RegistroDeMantenimientoDeVehiculo) int {

	if len(registros) == 1 {
		unicoRegistro := registros[0]
		return unicoRegistro.ObtenerKilometraje()
	}

	indiceUltimoRegistro := len(registros) - 1
	ultimoRegistro := registros[indiceUltimoRegistro]
	ultimoKilometraje := ultimoRegistro.ObtenerKilometraje()

	return ultimoKilometraje
}

func (c *Controlador) calcularKilometrosTotalesRecorridos(registros []capturaRegistroMantenimientoVehiculoModelos.RegistroDeMantenimientoDeVehiculo) int {
	var kilometrosTotalesRecorridos int = 0

	if len(registros) == 1 {
		return kilometrosTotalesRecorridos
	}

	var kilometrajeInicial = c.obtenerKilometrajeInicial(registros)
	var ultimoKilometraje = c.obtenerUltimoKilometraje(registros)

	kilometrosTotalesRecorridos = ultimoKilometraje - kilometrajeInicial
	return kilometrosTotalesRecorridos
}

func (c *Controlador) calcularKilometrosPromedioDiariosRecorridos(registros []capturaRegistroMantenimientoVehiculoModelos.RegistroDeMantenimientoDeVehiculo) float64 {

	if len(registros) == 1 {

		return 0
	}

	primerRegistro := registros[0]
	ultimoRegistro := registros[len(registros)-1]

	fechaInicial, _ := time.Parse(constantes.FORMATO_FECHA, primerRegistro.ObtenerFecha())
	fechaFinal, _ := time.Parse(constantes.FORMATO_FECHA, ultimoRegistro.ObtenerFecha())

	diasTotales := fechaFinal.Sub(fechaInicial).Hours() / 24

	if diasTotales == 0 {
		return 0
	}

	kilometroTotalesRecorridos := ultimoRegistro.ObtenerKilometraje() - primerRegistro.ObtenerKilometraje()
	promedioKilometrosDiariosRecorridos := float64(kilometroTotalesRecorridos) / diasTotales

	return promedioKilometrosDiariosRecorridos

}
