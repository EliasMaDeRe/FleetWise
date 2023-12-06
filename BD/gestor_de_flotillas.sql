-- phpMyAdmin SQL Dump
-- version 5.3.0-dev+20220510.314f251104
-- https://www.phpmyadmin.net/
--
-- Servidor: localhost:3307
-- Tiempo de generación: 06-12-2023 a las 22:06:02
-- Versión del servidor: 10.4.24-MariaDB
-- Versión de PHP: 8.1.4

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Base de datos: `gestor_de_flotillas`
--

DELIMITER $$
--
-- Procedimientos
--
CREATE DEFINER=`root`@`localhost` PROCEDURE `obtenerRegistrosYVehiculosAsociadoPorNumeroDeRegistro()` (IN `numRegistro` INT(11) UNSIGNED)   SELECT * FROM v_registros_y_vehiculos_aociados WHERE numero_de_registro LIKE numRegistro$$

CREATE DEFINER=`root`@`localhost` PROCEDURE `obtenerUsuariosPorNombreUsuario()` (IN `nombre` VARCHAR(255))   SELECT * FROM usuarios WHERE nombre_usuario LIKE nombre$$

CREATE DEFINER=`root`@`localhost` PROCEDURE `obtenerVehiculosPorPlacas()` (IN `placas` VARCHAR(255))  NO SQL SELECT * FROM vehiculos WHERE placas like placas$$

CREATE DEFINER=`root`@`localhost` PROCEDURE `obtenerVehiculosPorSerie()` (IN `serie` VARCHAR(255))  NO SQL SELECT * FROM  vehiculos WHERE serie LIKE serie$$

DELIMITER ;

-- --------------------------------------------------------

--
-- Estructura de tabla para la tabla `log`
--

CREATE TABLE `log` (
  `id` int(11) NOT NULL,
  `accion` varchar(255) NOT NULL,
  `fecha` datetime NOT NULL DEFAULT current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Volcado de datos para la tabla `log`
--

INSERT INTO `log` (`id`, `accion`, `fecha`) VALUES
(1, 'Se\r\ninserto un nuevo vehiculo con fecha_lanzamiento: 2021 marca: Nissan modelo: Versa placas: MD1232 y serie: SE1235', '2023-12-06 14:45:00'),
(2, 'Se inserto un nuevo registro con numero de registro nuevo: 6, tipo de registro nuevo: servicio, fecha nueva: 2023-12-06, litros de gasolina nuevo: 0, kilometraje nuevo: 12, importe nuevo: 12, observacion nuevo: a y concepto nuevo: Motor', '2023-12-06 14:55:46'),
(3, 'Se inserto el servicio de vehiculo con el nombre: LLantas', '2023-12-06 14:56:01'),
(4, 'Se elimino el servicio de vehiculo con el nombre: LLantas', '2023-12-06 14:56:38');

-- --------------------------------------------------------

--
-- Estructura de tabla para la tabla `registros_de_mantenimiento_de_vehiculo`
--

CREATE TABLE `registros_de_mantenimiento_de_vehiculo` (
  `numero_de_registro` int(11) NOT NULL,
  `placas_de_vehiculo` varchar(255) NOT NULL,
  `tipo_de_registro` varchar(255) NOT NULL,
  `fecha` varchar(255) NOT NULL,
  `litros_de_gasolina` float DEFAULT NULL,
  `kilometraje` int(11) NOT NULL,
  `importe` float NOT NULL,
  `observaciones` varchar(255) DEFAULT NULL,
  `concepto` varchar(255) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Volcado de datos para la tabla `registros_de_mantenimiento_de_vehiculo`
--

INSERT INTO `registros_de_mantenimiento_de_vehiculo` (`numero_de_registro`, `placas_de_vehiculo`, `tipo_de_registro`, `fecha`, `litros_de_gasolina`, `kilometraje`, `importe`, `observaciones`, `concepto`) VALUES
(4, 'MD123', 'servicio', '2023-12-06', 0, 12, 12, 'a', 'Motor'),
(6, 'MD123', 'servicio', '2023-12-06', 0, 12, 12, 'a', 'Motor');

--
-- Disparadores `registros_de_mantenimiento_de_vehiculo`
--
DELIMITER $$
CREATE TRIGGER `log_accion_registro_AI` AFTER INSERT ON `registros_de_mantenimiento_de_vehiculo` FOR EACH ROW INSERT into log(accion) VALUES(concat('Se inserto un nuevo registro con numero de registro nuevo: ',new.numero_de_registro,', tipo de registro nuevo: ',new.tipo_de_registro,
', fecha nueva: ',new.fecha,', litros de gasolina nuevo: ',new.litros_de_gasolina,
', kilometraje nuevo: ',new.kilometraje,', importe nuevo: ',
new.importe,', observacion nuevo: ',new.observaciones,' y concepto nuevo: ',new.concepto))
$$
DELIMITER ;
DELIMITER $$
CREATE TRIGGER `log_accion_registro_AU` AFTER UPDATE ON `registros_de_mantenimiento_de_vehiculo` FOR EACH ROW INSERT into log(accion) VALUES(concat('Se\r\nactualizo el vehiculo con numero de registro viejo: ',old.numero_de_registro,', tipo de registro viejo: ',old.tipo_de_registro,
', fecha viejo: ',old.fecha,', litros de gasolina vieja: ',old.litros_de_gasolina,
', kilometraje viejo: ',old.kilometraje,', importe viejo: ',
old.importe,', observacion vieja: ',old.observaciones,' y concepto viejo: ',old.concepto,' a numero de registro nuevo: ',new.numero_de_registro,', tipo de registro nuevo: ',new.tipo_de_registro,
', fecha nueva: ',new.fecha,', litros de gasolina nuevo: ',new.litros_de_gasolina,
', kilometraje nuevo: ',new.kilometraje,', importe nuevo: ',
new.importe,', observacion nuevo: ',new.observaciones,' y concepto nuevo: ',new.concepto))
$$
DELIMITER ;
DELIMITER $$
CREATE TRIGGER `log_accion_registro_BD` BEFORE DELETE ON `registros_de_mantenimiento_de_vehiculo` FOR EACH ROW INSERT into log(accion) VALUES(concat('Se elimino el registro con el numero de registro: ',old.numero_de_registro))
$$
DELIMITER ;

-- --------------------------------------------------------

--
-- Estructura de tabla para la tabla `servicios_vehiculares`
--

CREATE TABLE `servicios_vehiculares` (
  `nombre` varchar(255) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Volcado de datos para la tabla `servicios_vehiculares`
--

INSERT INTO `servicios_vehiculares` (`nombre`) VALUES
('Motor');

--
-- Disparadores `servicios_vehiculares`
--
DELIMITER $$
CREATE TRIGGER `log_accion_servicio_vehiculo_AI` AFTER INSERT ON `servicios_vehiculares` FOR EACH ROW INSERT into log(accion) VALUES(concat('Se inserto el servicio de vehiculo con el nombre: ',new.nombre))
$$
DELIMITER ;
DELIMITER $$
CREATE TRIGGER `log_accion_servicio_vehiculo_AU` AFTER UPDATE ON `servicios_vehiculares` FOR EACH ROW INSERT into log(accion) VALUES(concat('Se inserto un nuevo servicio de vehiculo con nombre viejo: ',old.nombre,'por nombre nuevo',new.nombre))
$$
DELIMITER ;
DELIMITER $$
CREATE TRIGGER `log_accion_servicio_vehiculo_BD` BEFORE DELETE ON `servicios_vehiculares` FOR EACH ROW INSERT into log(accion) VALUES(concat('Se elimino el servicio de vehiculo con el nombre: ',old.nombre))
$$
DELIMITER ;

-- --------------------------------------------------------

--
-- Estructura de tabla para la tabla `usuarios`
--

CREATE TABLE `usuarios` (
  `cargo` varchar(255) DEFAULT NULL,
  `clave_usuario` varchar(255) DEFAULT NULL,
  `nombre_usuario` varchar(255) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Volcado de datos para la tabla `usuarios`
--

INSERT INTO `usuarios` (`cargo`, `clave_usuario`, `nombre_usuario`) VALUES
('administrador', 'admin', 'admin'),
('administrador', '$2a$10$VfbUnFgz5VoUtmW.AEiZneOUVqCUQiPCWIZB1k2Wzo0ujC4lkqW7C', 'Fer');

--
-- Disparadores `usuarios`
--
DELIMITER $$
CREATE TRIGGER `log_accion_usuario_AI` AFTER INSERT ON `usuarios` FOR EACH ROW INSERT into log(accion) VALUES(concat('Se inserto un nuevo usuario con cargo: ',new.cargo,', clave usuario: ',new.clave_usuario,
'y nombre: ',new.nombre_usuario))
$$
DELIMITER ;
DELIMITER $$
CREATE TRIGGER `log_accion_usuario_AU` AFTER UPDATE ON `usuarios` FOR EACH ROW INSERT into log(accion) VALUES(concat('Se actualizo un usuario con cargo viejo: ',old.cargo,', clave usuario viejo: ',old.clave_usuario,
'y nombre viejo: ',old.nombre_usuario,'por un nuevo usuario con cargo: ',new.cargo,', clave usuario: ',new.clave_usuario,
'y nombre: ',new.nombre_usuario))
$$
DELIMITER ;
DELIMITER $$
CREATE TRIGGER `log_accion_usuario_BD` BEFORE DELETE ON `usuarios` FOR EACH ROW INSERT into log(accion) VALUES(concat('Se elimino un usuario con clave: ',old.clave_usuario))
$$
DELIMITER ;

-- --------------------------------------------------------

--
-- Estructura de tabla para la tabla `vehiculos`
--

CREATE TABLE `vehiculos` (
  `fecha_lanzamiento` int(11) DEFAULT NULL,
  `marca` varchar(255) DEFAULT NULL,
  `modelo` varchar(255) DEFAULT NULL,
  `placas` varchar(255) NOT NULL,
  `serie` varchar(255) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Volcado de datos para la tabla `vehiculos`
--

INSERT INTO `vehiculos` (`fecha_lanzamiento`, `marca`, `modelo`, `placas`, `serie`) VALUES
(2021, 'Nissan', 'Versa', 'MD123', 'SE123'),
(2021, 'Nissan', 'Versa', 'MD1232', 'SE1235');

--
-- Disparadores `vehiculos`
--
DELIMITER $$
CREATE TRIGGER `log_accion_vehiculo_AI` AFTER INSERT ON `vehiculos` FOR EACH ROW INSERT into log(accion) VALUES(concat('Se\r\ninserto un nuevo vehiculo con fecha_lanzamiento: ',new.fecha_lanzamiento,' marca: ',new.marca,
' modelo: ',new.modelo,' placas: ',new.placas,
' y serie: ',new.serie))
$$
DELIMITER ;
DELIMITER $$
CREATE TRIGGER `log_accion_vehiculo_AU` AFTER UPDATE ON `vehiculos` FOR EACH ROW INSERT into log(accion) VALUES(concat('Se actualizo el vehiculo con fecha_lanzamiento viejo: ',old.fecha_lanzamiento,', marca vieja: ',old.marca,
', modelo viejo: ',old.modelo,', placas viejas: ',old.placas,
'y serie viejo: ',old.serie,'por fecha_lanzamiento nueva: ',new.fecha_lanzamiento,' marca nueva: ',new.marca,
' modelo nuevo: ',new.modelo,' placas nueva: ',new.placas,
' serie nuevo: ',new.serie))
$$
DELIMITER ;
DELIMITER $$
CREATE TRIGGER `log_accion_vehiculo_BD` BEFORE DELETE ON `vehiculos` FOR EACH ROW INSERT into log(accion) VALUES(concat('Se\r\nelimino el vehiculo con las placas viejas: ',old.placas,
' y serie vieja: ',old.serie))
$$
DELIMITER ;

-- --------------------------------------------------------

--
-- Estructura Stand-in para la vista `v_registros_y_vehiculos_aociados`
-- (Véase abajo para la vista actual)
--
CREATE TABLE `v_registros_y_vehiculos_aociados` (
`fecha_lanzamiento` int(11)
,`marca` varchar(255)
,`modelo` varchar(255)
,`placas` varchar(255)
,`serie` varchar(255)
,`numero_de_registro` int(11)
,`tipo_de_registro` varchar(255)
,`fecha` varchar(255)
,`litros_de_gasolina` float
,`kilometraje` int(11)
,`importe` float
,`observaciones` varchar(255)
,`concepto` varchar(255)
);

-- --------------------------------------------------------

--
-- Estructura para la vista `v_registros_y_vehiculos_aociados`
--
DROP TABLE IF EXISTS `v_registros_y_vehiculos_aociados`;

CREATE ALGORITHM=UNDEFINED DEFINER=`root`@`localhost` SQL SECURITY DEFINER VIEW `v_registros_y_vehiculos_aociados`  AS SELECT `v`.`fecha_lanzamiento` AS `fecha_lanzamiento`, `v`.`marca` AS `marca`, `v`.`modelo` AS `modelo`, `v`.`placas` AS `placas`, `v`.`serie` AS `serie`, `r`.`numero_de_registro` AS `numero_de_registro`, `r`.`tipo_de_registro` AS `tipo_de_registro`, `r`.`fecha` AS `fecha`, `r`.`litros_de_gasolina` AS `litros_de_gasolina`, `r`.`kilometraje` AS `kilometraje`, `r`.`importe` AS `importe`, `r`.`observaciones` AS `observaciones`, `r`.`concepto` AS `concepto` FROM (`vehiculos` `v` join `registros_de_mantenimiento_de_vehiculo` `r` on(`v`.`placas` = `r`.`placas_de_vehiculo`))  ;

--
-- Índices para tablas volcadas
--

--
-- Indices de la tabla `log`
--
ALTER TABLE `log`
  ADD PRIMARY KEY (`id`);

--
-- Indices de la tabla `registros_de_mantenimiento_de_vehiculo`
--
ALTER TABLE `registros_de_mantenimiento_de_vehiculo`
  ADD PRIMARY KEY (`numero_de_registro`),
  ADD KEY `placas` (`placas_de_vehiculo`);

--
-- Indices de la tabla `servicios_vehiculares`
--
ALTER TABLE `servicios_vehiculares`
  ADD PRIMARY KEY (`nombre`);

--
-- Indices de la tabla `usuarios`
--
ALTER TABLE `usuarios`
  ADD PRIMARY KEY (`nombre_usuario`);

--
-- Indices de la tabla `vehiculos`
--
ALTER TABLE `vehiculos`
  ADD PRIMARY KEY (`placas`);

--
-- AUTO_INCREMENT de las tablas volcadas
--

--
-- AUTO_INCREMENT de la tabla `log`
--
ALTER TABLE `log`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=5;

--
-- AUTO_INCREMENT de la tabla `registros_de_mantenimiento_de_vehiculo`
--
ALTER TABLE `registros_de_mantenimiento_de_vehiculo`
  MODIFY `numero_de_registro` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=7;

--
-- Restricciones para tablas volcadas
--

--
-- Filtros para la tabla `registros_de_mantenimiento_de_vehiculo`
--
ALTER TABLE `registros_de_mantenimiento_de_vehiculo`
  ADD CONSTRAINT `registros_de_mantenimiento_de_vehiculo_ibfk_1` FOREIGN KEY (`placas_de_vehiculo`) REFERENCES `vehiculos` (`placas`) ON DELETE CASCADE;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;



