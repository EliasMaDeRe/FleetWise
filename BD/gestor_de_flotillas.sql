-- phpMyAdmin SQL Dump
-- version 5.3.0-dev+20220510.314f251104
-- https://www.phpmyadmin.net/
--
-- Servidor: localhost:3307
-- Tiempo de generación: 02-10-2023 a las 05:36:50
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

--
-- Volcado de datos para la tabla `vehiculos`
--

ALTER TABLE `vehiculos`
  ADD PRIMARY KEY (`placas`);

INSERT INTO `vehiculos` (`fecha_lanzamiento`, `marca`, `modelo`, `placas`, `serie`) VALUES
(1, 'a', 'b', 'c', 'd');

CREATE TABLE `servicios_vehiculares` (
  `nombre` varchar(255)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

ALTER TABLE `servicios_vehiculares`
  ADD PRIMARY KEY (`nombre`);

INSERT INTO `servicios_vehiculares` (`nombre`) VALUES
('Llantas');


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
-- Estructura de tabla para la tabla `registros_mantenimiento_vehicular`
--

CREATE TABLE `registros_mantenimiento_vehicular` (
  `numero_registro` int(11) NOT NULL,
  `tipo_registro` varchar(255) NOT NULL,
  `fecha` varchar(255) NOT NULL,
  `litros_gasolina` float DEFAULT NULL,
  `kilometraje` int(11) NOT NULL,
  `importe` float NOT NULL,
  `observaciones` varchar(255) DEFAULT NULL,
  `concepto` varchar(255) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Índices para tablas volcadas
--

--
-- Indices de la tabla `registros_mantenimiento_vehicular`
--
ALTER TABLE `registros_mantenimiento_vehicular`
  ADD PRIMARY KEY (`numero_registro`);

--
-- AUTO_INCREMENT de las tablas volcadas
--

--
-- AUTO_INCREMENT de la tabla `registros_mantenimiento_vehicular`
--
ALTER TABLE `registros_mantenimiento_vehicular`
  MODIFY `numero_registro` int(11) NOT NULL AUTO_INCREMENT;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;






