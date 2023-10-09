-- phpMyAdmin SQL Dump
-- version 5.2.1
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Generation Time: Oct 09, 2023 at 11:51 AM
-- Server version: 10.4.28-MariaDB
-- PHP Version: 8.2.4

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `crud_go`
--

-- --------------------------------------------------------

--
-- Table structure for table `book`
--

CREATE TABLE `book` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `title` longtext DEFAULT NULL,
  `author` longtext DEFAULT NULL,
  `publisher` longtext DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `book`
--

INSERT INTO `book` (`id`, `created_at`, `updated_at`, `deleted_at`, `title`, `author`, `publisher`) VALUES
(1, '2023-10-09 16:49:59.887', '2023-10-09 16:49:59.887', NULL, 'Bandung Ketika Dini Hari', 'Gigi Suliadi', 'Bandung Media'),
(2, '2023-10-09 16:50:12.684', '2023-10-09 16:50:12.684', NULL, 'jakarta Ketika Dini Hari', 'Gigi Suliadi', 'Bandung Media'),
(3, '2023-10-09 16:50:21.389', '2023-10-09 16:50:21.389', NULL, 'malam minggu', 'Gigi Suliadi', 'Bandung Media'),
(4, '2023-10-09 16:50:26.828', '2023-10-09 16:50:26.828', NULL, 'Minggu Pagi', 'Gigi Suliadi', 'Bandung Media'),
(5, '2023-10-09 16:50:33.025', '2023-10-09 16:50:33.025', NULL, 'Pagi Hari', 'Gigi Suliadi', 'Bandung Media');

-- --------------------------------------------------------

--
-- Table structure for table `user`
--

CREATE TABLE `user` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `name` longtext DEFAULT NULL,
  `email` longtext DEFAULT NULL,
  `password` longtext DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `user`
--

INSERT INTO `user` (`id`, `created_at`, `updated_at`, `deleted_at`, `name`, `email`, `password`) VALUES
(1, '2023-10-09 16:48:07.790', '2023-10-09 16:48:07.790', NULL, 'Muhammad FFandia', 'muhfandioo87@gmail.com', 'fandia01'),
(2, '2023-10-09 16:48:25.146', '2023-10-09 16:48:25.146', NULL, 'Haji Memashite', 'hajimema23@gmail.com', 'mamah12'),
(3, '2023-10-09 16:48:40.021', '2023-10-09 16:48:40.021', NULL, 'Janji Jiwa', 'matos13@gmail.com', 'tosca12'),
(4, '2023-10-09 16:48:56.577', '2023-10-09 16:48:56.577', NULL, 'Maharani Putri', 'maput@gmail.com', 'puput65'),
(5, '2023-10-09 16:49:13.935', '2023-10-09 16:49:13.935', NULL, 'Hajah Juhaeri', 'matot12@gmail.com', 'ninggalkeun');

--
-- Indexes for dumped tables
--

--
-- Indexes for table `book`
--
ALTER TABLE `book`
  ADD PRIMARY KEY (`id`),
  ADD KEY `idx_book_deleted_at` (`deleted_at`);

--
-- Indexes for table `user`
--
ALTER TABLE `user`
  ADD PRIMARY KEY (`id`),
  ADD KEY `idx_user_deleted_at` (`deleted_at`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `book`
--
ALTER TABLE `book`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=6;

--
-- AUTO_INCREMENT for table `user`
--
ALTER TABLE `user`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=6;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
