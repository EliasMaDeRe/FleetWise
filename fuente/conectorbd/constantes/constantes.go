package constantes

const (
	ERROR_CONECTAR_BD               = "Error: No se pudo conectar a la bd"
	ERROR_PLACAS_INEXISTENTES_EN_BD = "Error: Placas no existen en la base de datos"
	ERROR_SERIE_INEXISTENTES_EN_BD  = "Error: La serie no existe en la base de datos"
	ERROR_GUARDAR_FILA_BD           = "Error: No se pudo guardar la fila en la base de datos"
	ERROR_BUSQUEDA_EN_BD            = "Error: Hubo un error con la busqueda en la base de datos"
	CONSULTA_REGISTROS_CON_VEHICULOS = "JOIN vehiculos on vehiculos.placas = registros_de_mantenimiento_de_vehiculo.placas_de_vehiculo"
)
