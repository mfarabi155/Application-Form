-- phpMyAdmin SQL Dump
-- version 5.2.1
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Waktu pembuatan: 22 Jun 2024 pada 17.58
-- Versi server: 10.4.28-MariaDB
-- Versi PHP: 8.2.4

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `application_form`
--

-- --------------------------------------------------------

--
-- Struktur dari tabel `log_errors`
--

CREATE TABLE `log_errors` (
  `id` int(11) NOT NULL,
  `timestamp` datetime(3) DEFAULT current_timestamp(3),
  `error_message` text DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data untuk tabel `log_errors`
--

INSERT INTO `log_errors` (`id`, `timestamp`, `error_message`) VALUES
(1, '2024-06-22 15:49:16.159', 'Record not found!: record not found'),
(2, '2024-06-22 15:50:59.863', 'Record not found!: record not found'),
(3, '2024-06-22 15:51:00.543', 'Record not found!: record not found');

-- --------------------------------------------------------

--
-- Struktur dari tabel `persons`
--

CREATE TABLE `persons` (
  `id` int(11) NOT NULL,
  `name` longtext DEFAULT NULL,
  `identity_number` longtext DEFAULT NULL,
  `email` longtext DEFAULT NULL,
  `date_of_birth` date NOT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data untuk tabel `persons`
--

INSERT INTO `persons` (`id`, `name`, `identity_number`, `email`, `date_of_birth`, `created_at`, `updated_at`) VALUES
(8, 'Muhammad Farabi Ismail', '293759823752', 'mfarabi155@gmail.com', '1998-03-11', '2024-06-22 22:51:19.549', '2024-06-22 22:51:19.549');

--
-- Indexes for dumped tables
--

--
-- Indeks untuk tabel `log_errors`
--
ALTER TABLE `log_errors`
  ADD PRIMARY KEY (`id`);

--
-- Indeks untuk tabel `persons`
--
ALTER TABLE `persons`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `uni_persons_identity_number` (`identity_number`) USING HASH;

--
-- AUTO_INCREMENT untuk tabel yang dibuang
--

--
-- AUTO_INCREMENT untuk tabel `log_errors`
--
ALTER TABLE `log_errors`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=4;

--
-- AUTO_INCREMENT untuk tabel `persons`
--
ALTER TABLE `persons`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=9;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
