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
-- AUTO_INCREMENT de las tablas volcadas
--

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;



