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

-- --------------------------------------------------------

--
-- Estructura de tabla para la tabla `vehiculos`
--

CREATE TABLE `vehiculos` (
  `fecha_lanzamiento` int(11) DEFAULT NULL,
  `marca` varchar(255) DEFAULT NULL,
  `modelo` varchar(255) DEFAULT NULL,
  `placas` varchar(255),
  `serie` varchar(255) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

ALTER TABLE `vehiculos`
  ADD PRIMARY KEY (`placas`);
--
-- Volcado de datos para la tabla `vehiculos`
--

INSERT INTO `vehiculos` (`fecha_lanzamiento`, `marca`, `modelo`, `placas`, `serie`) VALUES
(2010, 'marca1', 'modelo1', 'placas1', 'numSerie1');

--
-- Estructura de tabla para la tabla `servicios_vehiculares`
--

CREATE TABLE `servicios_vehiculares` (
  `nombre` varchar(255)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

ALTER TABLE `servicios_vehiculares`
  ADD PRIMARY KEY (`nombre`);

INSERT INTO `servicios_vehiculares` (`nombre`) VALUES
('Llantas');

--
-- Estructura de tabla para la tabla `usuarios`
--

CREATE TABLE `usuarios` (
  `cargo` varchar(255) DEFAULT NULL,
  `clave_usuario` varchar(255) DEFAULT NULL,
  `nombre_usuario` varchar(255)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

ALTER TABLE `usuarios`
  ADD PRIMARY KEY (`nombre_usuario`);

INSERT INTO `usuarios` (`cargo`, `clave_usuario`, `nombre_usuario`) VALUES
('administrador','admin','admin');

--
-- Estructura de tabla para la tabla `registros_mantenimiento_vehicular`
--

CREATE TABLE `registros_mantenimiento_vehicular` (
  `numero_de_registro` int(11) NOT NULL AUTO_INCREMENT,
  `placas_vehiculo` varchar(255) NOT NULL,
  `tipo_de_registro` varchar(255) NOT NULL,
  `fecha` varchar(255) NOT NULL,
  `litros_de_gasolina` float DEFAULT NULL,
  `kilometraje` int(11) NOT NULL,
  `importe` float NOT NULL,
  `observaciones` varchar(255) DEFAULT NULL,
  `concepto` varchar(255) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;


--
-- Indices de la tabla `registros_mantenimiento_vehicular`
--
ALTER TABLE `registros_mantenimiento_vehicular`
  ADD PRIMARY KEY (`numero_registro`),
  ADD KEY `placas` (`placas_vehiculo`);

--
-- Filtros para la tabla `registros_mantenimiento_vehicular`
--
ALTER TABLE `registros_mantenimiento_vehicular`
  ADD CONSTRAINT `registros_mantenimiento_vehicular_ibfk_1` FOREIGN KEY (`placas_vehiculo`) REFERENCES `vehiculos` (`placas`);
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;





