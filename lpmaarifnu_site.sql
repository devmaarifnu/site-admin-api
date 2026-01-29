-- --------------------------------------------------------
-- Host:                         127.0.0.1
-- Server version:               8.0.30 - MySQL Community Server - GPL
-- Server OS:                    Win64
-- HeidiSQL Version:             12.1.0.6537
-- --------------------------------------------------------

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET NAMES utf8 */;
/*!50503 SET NAMES utf8mb4 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;


-- Dumping database structure for lpmaarifnu_site
CREATE DATABASE IF NOT EXISTS `lpmaarifnu_site` /*!40100 DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci */ /*!80016 DEFAULT ENCRYPTION='N' */;
USE `lpmaarifnu_site`;

-- Dumping structure for table lpmaarifnu_site.activity_logs
CREATE TABLE IF NOT EXISTS `activity_logs` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `log_name` varchar(100) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `description` text COLLATE utf8mb4_unicode_ci,
  `subject_type` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `subject_id` bigint unsigned DEFAULT NULL,
  `causer_type` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `causer_id` bigint unsigned DEFAULT NULL,
  `properties` json DEFAULT NULL,
  `ip_address` varchar(45) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `user_agent` text COLLATE utf8mb4_unicode_ci,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `idx_log_name` (`log_name`),
  KEY `idx_subject` (`subject_type`,`subject_id`),
  KEY `idx_causer` (`causer_type`,`causer_id`),
  KEY `idx_created_at` (`created_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Dumping data for table lpmaarifnu_site.activity_logs: ~0 rows (approximately)
DELETE FROM `activity_logs`;

-- Dumping structure for table lpmaarifnu_site.board_members
CREATE TABLE IF NOT EXISTS `board_members` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `position_id` int unsigned NOT NULL,
  `name` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `title` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT 'Gelar akademik',
  `photo` varchar(500) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `bio` text COLLATE utf8mb4_unicode_ci,
  `email` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `phone` varchar(20) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `social_media` json DEFAULT NULL COMMENT '{"facebook": "", "twitter": "", "linkedin": ""}',
  `period_start` year NOT NULL,
  `period_end` year NOT NULL,
  `is_active` tinyint(1) DEFAULT '1',
  `order_number` int DEFAULT '0',
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `position_id` (`position_id`),
  KEY `idx_period` (`period_start`,`period_end`),
  KEY `idx_is_active` (`is_active`),
  KEY `idx_order_number` (`order_number`),
  CONSTRAINT `board_members_ibfk_1` FOREIGN KEY (`position_id`) REFERENCES `organization_positions` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Dumping data for table lpmaarifnu_site.board_members: ~5 rows (approximately)
DELETE FROM `board_members`;
INSERT INTO `board_members` (`id`, `position_id`, `name`, `title`, `photo`, `bio`, `email`, `phone`, `social_media`, `period_start`, `period_end`, `is_active`, `order_number`, `created_at`, `updated_at`) VALUES
	(1, 1, 'Prof. Dr. KH. Said Aqil Siradj, MA', 'Prof. Dr. KH.', 'https://ui-avatars.com/api/?name=Said+Aqil+Siradj&background=1976D2&color=fff&size=400', 'Ketua Umum LP Ma\'arif NU periode 2024-2029. Beliau adalah ulama dan akademisi terkemuka yang memiliki dedikasi tinggi dalam pengembangan pendidikan Islam di Indonesia.', 'ketua@lpmaarifnu.or.id', '021-12345678', '{"twitter": "https://twitter.com/saidaqil", "facebook": "https://facebook.com/saidaqilsiradj", "instagram": "https://instagram.com/saidaqilsiradj"}', '2024', '2029', 1, 1, '2026-01-14 07:22:18', '2026-01-14 07:22:18'),
	(2, 2, 'Dr. H. Ahmad Lutfi, M.Pd', 'Dr. H.', 'https://ui-avatars.com/api/?name=Ahmad+Lutfi&background=388E3C&color=fff&size=400', 'Wakil Ketua I LP Ma\'arif NU yang membidangi Pendidikan Dasar dan Menengah. Berpengalaman dalam manajemen pendidikan lebih dari 20 tahun.', 'wakilketua1@lpmaarifnu.or.id', '021-12345679', '{"linkedin": "https://linkedin.com/in/ahmadlutfi"}', '2024', '2029', 1, 2, '2026-01-14 07:22:18', '2026-01-14 07:22:18'),
	(3, 3, 'Dr. H. Abdurrahman Wahid, M.Ag', 'Dr. H.', 'https://ui-avatars.com/api/?name=Abdurrahman+Wahid&background=F57C00&color=fff&size=400', 'Wakil Ketua II yang fokus pada pengembangan kurikulum dan peningkatan kualitas pembelajaran di sekolah-sekolah Ma\'arif.', 'wakilketua2@lpmaarifnu.or.id', '021-12345680', '{}', '2024', '2029', 1, 3, '2026-01-14 07:22:18', '2026-01-14 07:22:18'),
	(4, 4, 'Dr. Hj. Siti Aisyah, M.Pd', 'Dr. Hj.', 'https://ui-avatars.com/api/?name=Siti+Aisyah&background=C62828&color=fff&size=400', 'Sekretaris Umum LP Ma\'arif NU yang bertanggung jawab atas administrasi dan koordinasi kegiatan organisasi.', 'sekretaris@lpmaarifnu.or.id', '021-12345681', '{"instagram": "https://instagram.com/sitiaisyah"}', '2024', '2029', 1, 4, '2026-01-14 07:22:18', '2026-01-14 07:22:18'),
	(5, 5, 'H. Abdul Rahman, SE, M.Ak', 'H.', 'https://ui-avatars.com/api/?name=Abdul+Rahman&background=7B1FA2&color=fff&size=400', 'Bendahara Umum yang mengelola keuangan dan aset LP Ma\'arif NU dengan prinsip transparansi dan akuntabilitas.', 'bendahara@lpmaarifnu.or.id', '021-12345682', '{}', '2024', '2029', 1, 5, '2026-01-14 07:22:18', '2026-01-14 07:22:18');

-- Dumping structure for table lpmaarifnu_site.cache
CREATE TABLE IF NOT EXISTS `cache` (
  `cache_key` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `cache_value` mediumtext COLLATE utf8mb4_unicode_ci NOT NULL,
  `expiration` int NOT NULL,
  PRIMARY KEY (`cache_key`),
  KEY `idx_expiration` (`expiration`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Dumping data for table lpmaarifnu_site.cache: ~0 rows (approximately)
DELETE FROM `cache`;

-- Dumping structure for table lpmaarifnu_site.cache_locks
CREATE TABLE IF NOT EXISTS `cache_locks` (
  `lock_key` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `owner` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `expiration` int NOT NULL,
  PRIMARY KEY (`lock_key`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Dumping data for table lpmaarifnu_site.cache_locks: ~0 rows (approximately)
DELETE FROM `cache_locks`;

-- Dumping structure for table lpmaarifnu_site.categories
CREATE TABLE IF NOT EXISTS `categories` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(100) COLLATE utf8mb4_unicode_ci NOT NULL,
  `slug` varchar(100) COLLATE utf8mb4_unicode_ci NOT NULL,
  `description` text COLLATE utf8mb4_unicode_ci,
  `type` enum('news','opinion','document') COLLATE utf8mb4_unicode_ci NOT NULL,
  `icon` varchar(100) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `color` varchar(7) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `is_active` tinyint(1) DEFAULT '1',
  `order_number` int DEFAULT '0',
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `slug` (`slug`),
  KEY `idx_slug` (`slug`),
  KEY `idx_type` (`type`),
  KEY `idx_is_active` (`is_active`)
) ENGINE=InnoDB AUTO_INCREMENT=10 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Dumping data for table lpmaarifnu_site.categories: ~9 rows (approximately)
DELETE FROM `categories`;
INSERT INTO `categories` (`id`, `name`, `slug`, `description`, `type`, `icon`, `color`, `is_active`, `order_number`, `created_at`, `updated_at`) VALUES
	(1, 'Nasional', 'nasional', 'Berita tingkat nasional', 'news', NULL, NULL, 1, 1, '2026-01-14 07:21:51', '2026-01-14 07:21:51'),
	(2, 'Daerah', 'daerah', 'Berita tingkat daerah', 'news', NULL, NULL, 1, 2, '2026-01-14 07:21:51', '2026-01-14 07:21:51'),
	(3, 'Program', 'program', 'Program dan kegiatan', 'news', NULL, NULL, 1, 3, '2026-01-14 07:21:51', '2026-01-14 07:21:51'),
	(4, 'Pengumuman', 'pengumuman', 'Pengumuman resmi', 'news', NULL, NULL, 1, 4, '2026-01-14 07:21:51', '2026-01-14 07:21:51'),
	(5, 'Pedoman', 'pedoman', 'Pedoman dan panduan', 'document', NULL, NULL, 1, 1, '2026-01-14 07:21:51', '2026-01-14 07:21:51'),
	(6, 'Kurikulum', 'kurikulum', 'Dokumen kurikulum', 'document', NULL, NULL, 1, 2, '2026-01-14 07:21:51', '2026-01-14 07:21:51'),
	(7, 'Regulasi', 'regulasi', 'Peraturan dan regulasi', 'document', NULL, NULL, 1, 3, '2026-01-14 07:21:51', '2026-01-14 07:21:51'),
	(8, 'Panduan', 'panduan', 'Panduan teknis', 'document', NULL, NULL, 1, 4, '2026-01-14 07:21:51', '2026-01-14 07:21:51'),
	(9, 'Formulir', 'formulir', 'Formulir dan template', 'document', NULL, NULL, 1, 5, '2026-01-14 07:21:51', '2026-01-14 07:21:51');

-- Dumping structure for table lpmaarifnu_site.contact_messages
CREATE TABLE IF NOT EXISTS `contact_messages` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `ticket_id` varchar(50) COLLATE utf8mb4_unicode_ci NOT NULL,
  `name` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `email` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `phone` varchar(20) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `subject` varchar(500) COLLATE utf8mb4_unicode_ci NOT NULL,
  `message` text COLLATE utf8mb4_unicode_ci NOT NULL,
  `status` enum('new','read','in_progress','resolved','closed') COLLATE utf8mb4_unicode_ci DEFAULT 'new',
  `priority` enum('low','medium','high','urgent') COLLATE utf8mb4_unicode_ci DEFAULT 'medium',
  `ip_address` varchar(45) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `user_agent` text COLLATE utf8mb4_unicode_ci,
  `assigned_to` bigint unsigned DEFAULT NULL,
  `replied_at` timestamp NULL DEFAULT NULL,
  `resolved_at` timestamp NULL DEFAULT NULL,
  `notes` text COLLATE utf8mb4_unicode_ci COMMENT 'Internal notes',
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `ticket_id` (`ticket_id`),
  KEY `assigned_to` (`assigned_to`),
  KEY `idx_ticket_id` (`ticket_id`),
  KEY `idx_status` (`status`),
  KEY `idx_priority` (`priority`),
  KEY `idx_email` (`email`),
  KEY `idx_created_at` (`created_at`),
  CONSTRAINT `contact_messages_ibfk_1` FOREIGN KEY (`assigned_to`) REFERENCES `users` (`id`) ON DELETE SET NULL
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Dumping data for table lpmaarifnu_site.contact_messages: ~4 rows (approximately)
DELETE FROM `contact_messages`;
INSERT INTO `contact_messages` (`id`, `ticket_id`, `name`, `email`, `phone`, `subject`, `message`, `status`, `priority`, `ip_address`, `user_agent`, `assigned_to`, `replied_at`, `resolved_at`, `notes`, `created_at`, `updated_at`) VALUES
	(1, 'CTK-2024-0001', 'Budi Santoso', 'budi.santoso@email.com', '081234567890', 'Pertanyaan tentang Program Beasiswa', 'Saya ingin menanyakan persyaratan dan cara pendaftaran program beasiswa unggulan 2024. Mohon informasinya.', 'read', 'medium', '192.168.1.100', 'Mozilla/5.0', NULL, NULL, NULL, NULL, '2026-01-12 07:22:18', '2026-01-14 07:22:18'),
	(2, 'CTK-2024-0002', 'Siti Rahmawati', 'siti.rahmawati@email.com', '082345678901', 'Informasi Pendaftaran Siswa Baru', 'Apakah ada informasi tentang pendaftaran siswa baru untuk tahun ajaran 2024/2025?', 'new', 'high', '192.168.1.101', 'Mozilla/5.0', NULL, NULL, NULL, NULL, '2026-01-13 07:22:18', '2026-01-14 07:22:18'),
	(3, 'CTK-2024-0003', 'Ahmad Fauzi', 'ahmad.fauzi@email.com', '083456789012', 'Kerjasama Kelembagaan', 'Kami dari Yayasan XYZ ingin menjalin kerjasama dengan LP Ma\'arif NU. Bagaimana prosedurnya?', 'in_progress', 'high', '192.168.1.102', 'Mozilla/5.0', NULL, NULL, NULL, NULL, '2026-01-11 07:22:18', '2026-01-14 07:22:18'),
	(4, 'CTK-2024-0004', 'Dewi Lestari', 'dewi.lestari@email.com', NULL, 'Pertanyaan Umum', 'Dimana saya bisa mendapatkan informasi lengkap tentang program-program LP Ma\'arif NU?', 'resolved', 'low', '192.168.1.103', 'Mozilla/5.0', NULL, NULL, NULL, NULL, '2026-01-09 07:22:18', '2026-01-14 07:22:18'),
	(5, 'CTK-2024-0005', 'John Doe', 'john@example.com', '081234567890', 'Pertanyaan tentang pendaftaran', 'Saya ingin menanyakan prosedur pendaftaran untuk siswa baru tahun 2024/2025. Mohon informasinya.', 'new', 'medium', '::1', 'PostmanRuntime/7.51.0', NULL, NULL, NULL, NULL, '2026-01-14 07:52:16', '2026-01-14 07:52:16');

-- Dumping structure for table lpmaarifnu_site.departments
CREATE TABLE IF NOT EXISTS `departments` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `description` text COLLATE utf8mb4_unicode_ci,
  `head_name` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `order_number` int DEFAULT '0',
  `is_active` tinyint(1) DEFAULT '1',
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `idx_order_number` (`order_number`),
  KEY `idx_is_active` (`is_active`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Dumping data for table lpmaarifnu_site.departments: ~6 rows (approximately)
DELETE FROM `departments`;
INSERT INTO `departments` (`id`, `name`, `description`, `head_name`, `order_number`, `is_active`, `created_at`, `updated_at`) VALUES
	(1, 'Bidang Pendidikan Dasar', 'Mengelola pengembangan pendidikan tingkat SD/MI', NULL, 1, 1, '2026-01-14 07:21:51', '2026-01-14 07:21:51'),
	(2, 'Bidang Pendidikan Menengah', 'Mengelola pengembangan pendidikan tingkat SMP/MTs dan SMA/MA', NULL, 2, 1, '2026-01-14 07:21:51', '2026-01-14 07:21:51'),
	(3, 'Bidang Pendidikan Tinggi', 'Mengelola perguruan tinggi di bawah LP Ma\'arif NU', NULL, 3, 1, '2026-01-14 07:21:51', '2026-01-14 07:21:51'),
	(4, 'Bidang Kurikulum', 'Mengembangkan kurikulum berbasis Ma\'arif NU', NULL, 4, 1, '2026-01-14 07:21:51', '2026-01-14 07:21:51'),
	(5, 'Bidang SDM dan Kemitraan', 'Mengembangkan SDM dan kerjasama kelembagaan', NULL, 5, 1, '2026-01-14 07:21:51', '2026-01-14 07:21:51'),
	(6, 'Bidang Penelitian dan Pengembangan', 'Melakukan penelitian dan pengembangan pendidikan', NULL, 6, 1, '2026-01-14 07:21:51', '2026-01-14 07:21:51');

-- Dumping structure for table lpmaarifnu_site.documents
CREATE TABLE IF NOT EXISTS `documents` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `title` varchar(500) COLLATE utf8mb4_unicode_ci NOT NULL,
  `description` text COLLATE utf8mb4_unicode_ci,
  `category_id` int unsigned DEFAULT NULL,
  `file_name` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `file_path` varchar(500) COLLATE utf8mb4_unicode_ci NOT NULL,
  `file_type` varchar(50) COLLATE utf8mb4_unicode_ci NOT NULL,
  `file_size` bigint unsigned NOT NULL COMMENT 'in bytes',
  `mime_type` varchar(100) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `download_count` int unsigned DEFAULT '0',
  `is_public` tinyint(1) DEFAULT '1',
  `uploaded_by` bigint unsigned DEFAULT NULL,
  `status` enum('active','archived') COLLATE utf8mb4_unicode_ci DEFAULT 'active',
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `uploaded_by` (`uploaded_by`),
  KEY `idx_file_type` (`file_type`),
  KEY `idx_is_public` (`is_public`),
  KEY `idx_status` (`status`),
  KEY `idx_download_count` (`download_count`),
  KEY `idx_documents_category_type` (`category_id`,`file_type`,`status`),
  FULLTEXT KEY `idx_search` (`title`,`description`),
  CONSTRAINT `documents_ibfk_1` FOREIGN KEY (`category_id`) REFERENCES `categories` (`id`) ON DELETE SET NULL,
  CONSTRAINT `documents_ibfk_2` FOREIGN KEY (`uploaded_by`) REFERENCES `users` (`id`) ON DELETE SET NULL
) ENGINE=InnoDB AUTO_INCREMENT=14 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Dumping data for table lpmaarifnu_site.documents: ~12 rows (approximately)
DELETE FROM `documents`;
INSERT INTO `documents` (`id`, `title`, `description`, `category_id`, `file_name`, `file_path`, `file_type`, `file_size`, `mime_type`, `download_count`, `is_public`, `uploaded_by`, `status`, `created_at`, `updated_at`, `deleted_at`) VALUES
	(1, 'Pedoman Penyelenggaraan Pendidikan Ma\'arif NU 2024', 'Pedoman lengkap penyelenggaraan pendidikan di lingkungan Ma\'arif NU yang mencakup standar operasional, kurikulum, dan tata kelola sekolah.', 5, 'pedoman-pendidikan-maarif-2024.pdf', 'https://cdn.lpmaarifnu.or.id/documents/pedoman-pendidikan-maarif-2024.pdf', 'PDF', 2621440, 'application/pdf', 1536, 1, 1, 'active', '2026-01-14 07:22:18', '2026-01-14 07:22:18', NULL),
	(2, 'Kurikulum Integratif Ma\'arif NU Jenjang SMP/MTs', 'Dokumen kurikulum integratif yang mengintegrasikan nilai-nilai keislaman dengan kurikulum nasional untuk jenjang SMP/MTs.', 6, 'kurikulum-integratif-smp-mts.pdf', 'https://cdn.lpmaarifnu.or.id/documents/kurikulum-integratif-smp-mts.pdf', 'PDF', 3145728, 'application/pdf', 2342, 1, 1, 'active', '2026-01-14 07:22:18', '2026-01-14 07:22:18', NULL),
	(3, 'Kurikulum Integratif Ma\'arif NU Jenjang SMA/MA', 'Dokumen kurikulum integratif untuk jenjang SMA/MA dengan pendekatan saintifik dan nilai-nilai keislaman.', 6, 'kurikulum-integratif-sma-ma.pdf', 'https://cdn.lpmaarifnu.or.id/documents/kurikulum-integratif-sma-ma.pdf', 'PDF', 3670016, 'application/pdf', 1987, 1, 1, 'active', '2026-01-14 07:22:18', '2026-01-14 07:22:18', NULL),
	(4, 'Peraturan Standar Akreditasi Sekolah Ma\'arif', 'Peraturan dan standar akreditasi yang harus dipenuhi oleh sekolah-sekolah di bawah naungan LP Ma\'arif NU.', 7, 'standar-akreditasi-sekolah.pdf', 'https://cdn.lpmaarifnu.or.id/documents/standar-akreditasi-sekolah.pdf', 'PDF', 1572864, 'application/pdf', 876, 1, 3, 'active', '2026-01-14 07:22:18', '2026-01-14 07:22:18', NULL),
	(5, 'Panduan Implementasi Kurikulum Merdeka di Sekolah Ma\'arif', 'Panduan praktis implementasi Kurikulum Merdeka dengan tetap mempertahankan nilai-nilai Ma\'arif NU.', 8, 'panduan-kurikulum-merdeka.pdf', 'https://cdn.lpmaarifnu.or.id/documents/panduan-kurikulum-merdeka.pdf', 'PDF', 2097152, 'application/pdf', 3457, 1, 2, 'active', '2026-01-14 07:22:18', '2026-01-14 07:22:18', NULL),
	(6, 'Panduan Teknis Penilaian Pembelajaran Berbasis Proyek', 'Panduan teknis untuk guru dalam melakukan penilaian pembelajaran berbasis proyek sesuai Kurikulum Merdeka.', 8, 'panduan-penilaian-pbp.pdf', 'https://cdn.lpmaarifnu.or.id/documents/panduan-penilaian-pbp.pdf', 'PDF', 1048576, 'application/pdf', 1234, 1, 2, 'active', '2026-01-14 07:22:18', '2026-01-14 07:22:18', NULL),
	(7, 'Formulir Pendaftaran Beasiswa Unggulan 2024', 'Formulir pendaftaran untuk program beasiswa unggulan LP Ma\'arif NU tahun 2024.', 9, 'formulir-beasiswa-2024.pdf', 'https://cdn.lpmaarifnu.or.id/documents/formulir-beasiswa-2024.pdf', 'PDF', 524288, 'application/pdf', 4567, 1, 1, 'active', '2026-01-14 07:22:18', '2026-01-14 07:22:18', NULL),
	(8, 'Formulir Usulan Kegiatan Sekolah', 'Template formulir untuk mengajukan usulan kegiatan sekolah kepada LP Ma\'arif NU.', 9, 'formulir-usulan-kegiatan.docx', 'https://cdn.lpmaarifnu.or.id/documents/formulir-usulan-kegiatan.docx', 'DOCX', 204800, 'application/vnd.openxmlformats-officedocument.wordprocessingml.document', 789, 1, 3, 'active', '2026-01-14 07:22:18', '2026-01-14 07:22:18', NULL),
	(9, 'Juknis Penyelenggaraan Ujian Akhir Sekolah', 'Petunjuk teknis penyelenggaraan ujian akhir sekolah di lingkungan Ma\'arif NU.', 8, 'juknis-ujian-akhir.pdf', 'https://cdn.lpmaarifnu.or.id/documents/juknis-ujian-akhir.pdf', 'PDF', 1310720, 'application/pdf', 2134, 1, 2, 'active', '2026-01-14 07:22:18', '2026-01-14 07:22:18', NULL),
	(10, 'Modul Pelatihan Guru Digital', 'Modul lengkap untuk pelatihan kompetensi digital bagi guru-guru Ma\'arif.', 8, 'modul-pelatihan-guru-digital.pdf', 'https://cdn.lpmaarifnu.or.id/documents/modul-pelatihan-guru-digital.pdf', 'PDF', 5242880, 'application/pdf', 3890, 1, 1, 'active', '2026-01-14 07:22:18', '2026-01-14 07:22:18', NULL),
	(11, 'Modul Pelatihan Guru Digital', 'Modul lengkap untuk pelatihan kompetensi digital bagi guru-guru Ma\'arif.9', 8, 'modul-pelatihan-guru-digital.pdf', 'https://cdn.lpmaarifnu.or.id/documents/modul-pelatihan-guru-digital.pdf', 'PDF', 5242880, 'application/pdf', 3890, 1, 1, 'active', '2026-01-14 07:22:18', '2026-01-14 07:22:18', NULL),
	(12, 'Modul Pelatihan Guru Digital', 'Modul lengkap untuk pelatihan kompetensi digital bagi guru-guru Ma\'arif.0', 8, 'modul-pelatihan-guru-digital.pdf', 'https://cdn.lpmaarifnu.or.id/documents/modul-pelatihan-guru-digital.pdf', 'PDF', 5242880, 'application/pdf', 3890, 1, 1, 'active', '2026-01-14 07:22:18', '2026-01-14 08:05:56', NULL),
	(13, 'Modul Pelatihan Guru Digital', 'Modul lengkap untuk pelatihan kompetensi digital bagi guru-guru Ma\'arif.1', 8, 'modul-pelatihan-guru-digital.pdf', 'https://cdn.lpmaarifnu.or.id/documents/modul-pelatihan-guru-digital.pdf', 'PDF', 5242880, 'application/pdf', 3890, 1, 1, 'active', '2026-01-14 07:22:18', '2026-01-14 08:05:56', NULL);

-- Dumping structure for table lpmaarifnu_site.download_logs
CREATE TABLE IF NOT EXISTS `download_logs` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `document_id` bigint unsigned NOT NULL,
  `ip_address` varchar(45) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `user_agent` text COLLATE utf8mb4_unicode_ci,
  `downloaded_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `idx_document_id` (`document_id`),
  KEY `idx_downloaded_at` (`downloaded_at`),
  CONSTRAINT `download_logs_ibfk_1` FOREIGN KEY (`document_id`) REFERENCES `documents` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Dumping data for table lpmaarifnu_site.download_logs: ~4 rows (approximately)
DELETE FROM `download_logs`;
INSERT INTO `download_logs` (`id`, `document_id`, `ip_address`, `user_agent`, `downloaded_at`) VALUES
	(1, 1, '192.168.1.104', 'Mozilla/5.0', '2026-01-13 07:22:18'),
	(2, 1, '192.168.1.105', 'Mozilla/5.0', '2026-01-12 07:22:18'),
	(3, 2, '192.168.1.106', 'Mozilla/5.0', '2026-01-13 07:22:18'),
	(4, 5, '192.168.1.107', 'Mozilla/5.0', '2026-01-11 07:22:18');

-- Dumping structure for table lpmaarifnu_site.editorial_council
CREATE TABLE IF NOT EXISTS `editorial_council` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `institution` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `expertise` varchar(500) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `photo` varchar(500) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `bio` text COLLATE utf8mb4_unicode_ci,
  `email` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `order_number` int DEFAULT '0',
  `is_active` tinyint(1) DEFAULT '1',
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `idx_order_number` (`order_number`),
  KEY `idx_is_active` (`is_active`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Dumping data for table lpmaarifnu_site.editorial_council: ~4 rows (approximately)
DELETE FROM `editorial_council`;
INSERT INTO `editorial_council` (`id`, `name`, `institution`, `expertise`, `photo`, `bio`, `email`, `order_number`, `is_active`, `created_at`, `updated_at`) VALUES
	(1, 'Prof. Dr. KH. Abdullah Shiddiq, MA', 'UIN Syarif Hidayatullah Jakarta', 'Pendidikan Islam & Budaya', 'https://ui-avatars.com/api/?name=Abdullah+Shiddiq&background=0891B2&color=fff&size=400', NULL, NULL, 1, 1, '2026-01-14 07:22:18', '2026-01-14 07:22:18'),
	(2, 'Dr. Hj. Fatimah Zahra, M.Pd', 'Universitas Nahdlatul Ulama', 'Kurikulum & Pembelajaran', 'https://ui-avatars.com/api/?name=Fatimah+Zahra&background=DB2777&color=fff&size=400', NULL, NULL, 2, 1, '2026-01-14 07:22:18', '2026-01-14 07:22:18'),
	(3, 'Dr. Muhammad Ridwan, M.A', 'IAIN Surakarta', 'Media & Komunikasi Islam', 'https://ui-avatars.com/api/?name=Muhammad+Ridwan&background=EA580C&color=fff&size=400', NULL, NULL, 3, 1, '2026-01-14 07:22:18', '2026-01-14 07:22:18'),
	(4, 'Dr. Hj. Siti Nurjanah, M.Si', 'UIN Maulana Malik Ibrahim Malang', 'Manajemen Pendidikan', 'https://ui-avatars.com/api/?name=Siti+Nurjanah&background=16A34A&color=fff&size=400', NULL, NULL, 4, 1, '2026-01-14 07:22:18', '2026-01-14 07:22:18');

-- Dumping structure for table lpmaarifnu_site.editorial_team
CREATE TABLE IF NOT EXISTS `editorial_team` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `position` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `role_type` enum('pemimpin_redaksi','wakil_pemimpin_redaksi','redaktur_pelaksana','tim_redaksi') COLLATE utf8mb4_unicode_ci NOT NULL,
  `photo` varchar(500) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `bio` text COLLATE utf8mb4_unicode_ci,
  `email` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `phone` varchar(20) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `order_number` int DEFAULT '0',
  `is_active` tinyint(1) DEFAULT '1',
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `idx_role_type` (`role_type`),
  KEY `idx_order_number` (`order_number`),
  KEY `idx_is_active` (`is_active`)
) ENGINE=InnoDB AUTO_INCREMENT=11 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Dumping data for table lpmaarifnu_site.editorial_team: ~10 rows (approximately)
DELETE FROM `editorial_team`;
INSERT INTO `editorial_team` (`id`, `name`, `position`, `role_type`, `photo`, `bio`, `email`, `phone`, `order_number`, `is_active`, `created_at`, `updated_at`) VALUES
	(1, 'Dr. H. Muhammad Fadhil, M.Pd', 'Pemimpin Redaksi', 'pemimpin_redaksi', 'https://ui-avatars.com/api/?name=Muhammad+Fadhil&background=059669&color=fff&size=400', 'Pakar pendidikan Islam dengan pengalaman lebih dari 15 tahun di bidang jurnalistik pendidikan.', 'fadhil@lpmaarifnu.or.id', '021-12345678', 1, 1, '2026-01-14 07:22:18', '2026-01-14 07:22:18'),
	(2, 'Dra. Hj. Nur Azizah, M.Si', 'Wakil Pemimpin Redaksi I', 'wakil_pemimpin_redaksi', 'https://ui-avatars.com/api/?name=Nur+Azizah&background=7C3AED&color=fff&size=400', 'Spesialis media dan komunikasi publik.', 'azizah@lpmaarifnu.or.id', NULL, 2, 1, '2026-01-14 07:22:18', '2026-01-14 07:22:18'),
	(3, 'H. Abdul Malik, S.Pd.I, M.M', 'Wakil Pemimpin Redaksi II', 'wakil_pemimpin_redaksi', 'https://ui-avatars.com/api/?name=Abdul+Malik&background=DC2626&color=fff&size=400', 'Praktisi media dengan fokus pada jurnalisme pendidikan.', 'malik@lpmaarifnu.or.id', NULL, 3, 1, '2026-01-14 07:22:18', '2026-01-14 07:22:18'),
	(4, 'Ahmad Syarif, S.Sos, M.I.Kom', 'Redaktur Pelaksana', 'redaktur_pelaksana', 'https://ui-avatars.com/api/?name=Ahmad+Syarif&background=2563EB&color=fff&size=400', 'Koordinator harian tim redaksi dengan pengalaman di berbagai media nasional.', 'syarif@lpmaarifnu.or.id', NULL, 4, 1, '2026-01-14 07:22:18', '2026-01-14 07:22:18'),
	(5, 'Rizki Aulia Rahman, S.Pd', 'Editor Berita', 'tim_redaksi', 'https://ui-avatars.com/api/?name=Rizki+Rahman&background=6366F1&color=fff&size=400', NULL, NULL, NULL, 5, 1, '2026-01-14 07:22:18', '2026-01-14 07:22:18'),
	(6, 'Dewi Kusuma Wardani, S.Sos', 'Editor Opini', 'tim_redaksi', 'https://ui-avatars.com/api/?name=Dewi+Wardani&background=EC4899&color=fff&size=400', NULL, NULL, NULL, 6, 1, '2026-01-14 07:22:18', '2026-01-14 07:22:18'),
	(7, 'Faisal Akbar, S.Kom', 'Web Administrator', 'tim_redaksi', 'https://ui-avatars.com/api/?name=Faisal+Akbar&background=8B5CF6&color=fff&size=400', NULL, NULL, NULL, 7, 1, '2026-01-14 07:22:18', '2026-01-14 07:22:18'),
	(8, 'Rina Melati, S.Ds', 'Desainer Grafis', 'tim_redaksi', 'https://ui-avatars.com/api/?name=Rina+Melati&background=F59E0B&color=fff&size=400', NULL, NULL, NULL, 8, 1, '2026-01-14 07:22:18', '2026-01-14 07:22:18'),
	(9, 'Hendra Gunawan, S.I.Kom', 'Reporter', 'tim_redaksi', 'https://ui-avatars.com/api/?name=Hendra+Gunawan&background=10B981&color=fff&size=400', NULL, NULL, NULL, 9, 1, '2026-01-14 07:22:18', '2026-01-14 07:22:18'),
	(10, 'Laila Nur Hidayah, S.Pd', 'Content Writer', 'tim_redaksi', 'https://ui-avatars.com/api/?name=Laila+Hidayah&background=F43F5E&color=fff&size=400', NULL, NULL, NULL, 10, 1, '2026-01-14 07:22:18', '2026-01-14 07:22:18');

-- Dumping structure for table lpmaarifnu_site.event_flayers
CREATE TABLE IF NOT EXISTS `event_flayers` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `title` varchar(500) COLLATE utf8mb4_unicode_ci NOT NULL,
  `description` text COLLATE utf8mb4_unicode_ci,
  `image` varchar(500) COLLATE utf8mb4_unicode_ci NOT NULL,
  `event_date` date DEFAULT NULL,
  `event_location` varchar(500) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `registration_url` varchar(500) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `contact_person` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `contact_phone` varchar(20) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `contact_email` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `order_number` int DEFAULT '0',
  `is_active` tinyint(1) DEFAULT '1',
  `start_display_date` timestamp NULL DEFAULT NULL,
  `end_display_date` timestamp NULL DEFAULT NULL,
  `created_by` bigint unsigned DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `created_by` (`created_by`),
  KEY `idx_event_date` (`event_date`),
  KEY `idx_order_number` (`order_number`),
  KEY `idx_is_active` (`is_active`),
  KEY `idx_display_dates` (`start_display_date`,`end_display_date`),
  CONSTRAINT `event_flayers_ibfk_1` FOREIGN KEY (`created_by`) REFERENCES `users` (`id`) ON DELETE SET NULL
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Dumping data for table lpmaarifnu_site.event_flayers: ~3 rows (approximately)
DELETE FROM `event_flayers`;
INSERT INTO `event_flayers` (`id`, `title`, `description`, `image`, `event_date`, `event_location`, `registration_url`, `contact_person`, `contact_phone`, `contact_email`, `order_number`, `is_active`, `start_display_date`, `end_display_date`, `created_by`, `created_at`, `updated_at`) VALUES
	(1, 'Seminar Nasional Pendidikan Islam 2024', 'Seminar nasional dengan tema "Transformasi Pendidikan Islam di Era Digital" yang menghadirkan narasumber terkemuka.', 'https://images.unsplash.com/photo-1540575467063-178a50c2df87?w=800', '2024-03-15', 'Hotel Borobudur Jakarta', 'https://forms.lpmaarifnu.or.id/seminar2024', 'Panitia Seminar', '021-3920677', 'seminar@lpmaarifnu.or.id', 1, 1, '2026-01-07 07:22:18', '2024-03-14 17:00:00', 1, '2026-01-14 07:22:18', '2026-01-14 07:22:18'),
	(2, 'Workshop Kurikulum Merdeka', 'Workshop implementasi Kurikulum Merdeka untuk guru-guru Ma\'arif se-Indonesia.', 'https://images.unsplash.com/photo-1524178232363-1fb2b075b655?w=800', '2024-02-20', 'Gedung LP Ma\'arif NU Jakarta', 'https://forms.lpmaarifnu.or.id/workshop-kurikulum', 'Tim Kurikulum', '021-3920678', 'kurikulum@lpmaarifnu.or.id', 2, 1, '2026-01-04 07:22:18', '2024-02-19 17:00:00', 1, '2026-01-14 07:22:18', '2026-01-14 07:22:18'),
	(3, 'Jambore Pramuka Ma\'arif Nasional', 'Jambore Pramuka tingkat nasional untuk siswa-siswa Ma\'arif seluruh Indonesia.', 'https://images.unsplash.com/photo-1519995451813-39e29e054914?w=800', '2024-04-10', 'Cibubur, Jakarta Timur', 'https://forms.lpmaarifnu.or.id/jambore2024', 'Kwartir Pramuka Ma\'arif', '021-8093789', 'pramuka@lpmaarifnu.or.id', 3, 1, '2026-01-09 07:22:18', '2024-04-09 17:00:00', 2, '2026-01-14 07:22:18', '2026-01-14 07:22:18');

-- Dumping structure for procedure lpmaarifnu_site.get_popular_articles
DELIMITER //
CREATE PROCEDURE `get_popular_articles`(
    IN p_limit INT,
    IN p_days INT
)
BEGIN
    SELECT
        id, title, slug, views, published_at
    FROM news_articles
    WHERE
        status = 'published'
        AND deleted_at IS NULL
        AND published_at >= DATE_SUB(NOW(), INTERVAL p_days DAY)
    ORDER BY views DESC
    LIMIT p_limit;
END//
DELIMITER ;

-- Dumping structure for table lpmaarifnu_site.hero_slides
CREATE TABLE IF NOT EXISTS `hero_slides` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `title` varchar(500) COLLATE utf8mb4_unicode_ci NOT NULL,
  `description` text COLLATE utf8mb4_unicode_ci,
  `image` varchar(500) COLLATE utf8mb4_unicode_ci NOT NULL,
  `cta_label` varchar(100) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `cta_href` varchar(500) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `cta_secondary_label` varchar(100) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `cta_secondary_href` varchar(500) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `order_number` int DEFAULT '0',
  `is_active` tinyint(1) DEFAULT '1',
  `start_date` timestamp NULL DEFAULT NULL,
  `end_date` timestamp NULL DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `idx_order_number` (`order_number`),
  KEY `idx_is_active` (`is_active`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Dumping data for table lpmaarifnu_site.hero_slides: ~3 rows (approximately)
DELETE FROM `hero_slides`;
INSERT INTO `hero_slides` (`id`, `title`, `description`, `image`, `cta_label`, `cta_href`, `cta_secondary_label`, `cta_secondary_href`, `order_number`, `is_active`, `start_date`, `end_date`, `created_at`, `updated_at`) VALUES
	(1, 'Membangun Pendidikan Islam Berkualitas dan Berkarakter', 'LP Ma\'arif NU berkomitmen mengembangkan pendidikan Islam yang berkualitas dengan mengintegrasikan nilai-nilai keislaman, kearifan lokal, dan kemajuan teknologi untuk membentuk generasi yang berakhlak mulia dan berdaya saing tinggi.', 'https://images.unsplash.com/photo-1427504494785-3a9ca7044f45?w=1920', 'Pelajari Lebih Lanjut', '/tentang/visi-misi', 'Hubungi Kami', '/kontak', 1, 1, '2025-12-15 07:22:18', NULL, '2026-01-14 07:22:18', '2026-01-14 07:22:18'),
	(2, 'Program Beasiswa Unggulan 2024', 'Raih kesempatan emas untuk mendapatkan beasiswa pendidikan penuh! Daftar sekarang dan wujudkan impian pendidikanmu bersama LP Ma\'arif NU.', 'https://images.unsplash.com/photo-1523050854058-8df90110c9f1?w=1920', 'Daftar Sekarang', '/beasiswa', 'Info Lengkap', '/berita/peluncuran-program-beasiswa-unggulan-2024', 2, 1, '2025-12-25 07:22:18', '2026-03-15 07:22:18', '2026-01-14 07:22:18', '2026-01-14 07:22:18'),
	(3, 'Digitalisasi Pendidikan Ma\'arif', 'Bergabunglah dalam transformasi digital pendidikan! Pelatihan guru digital, pembelajaran online, dan inovasi teknologi pendidikan untuk masa depan yang lebih baik.', 'https://images.unsplash.com/photo-1501504905252-473c47e087f8?w=1920', 'Ikuti Program', '/program/digital', 'Lihat Dokumentasi', '/dokumen', 3, 1, '2026-01-04 07:22:18', NULL, '2026-01-14 07:22:18', '2026-01-14 07:22:18');

-- Dumping structure for procedure lpmaarifnu_site.increment_view_count
DELIMITER //
CREATE PROCEDURE `increment_view_count`(
    IN p_viewable_type VARCHAR(255),
    IN p_viewable_id BIGINT UNSIGNED
)
BEGIN
    IF p_viewable_type = 'news_articles' THEN
        UPDATE news_articles SET views = views + 1 WHERE id = p_viewable_id;
    ELSEIF p_viewable_type = 'opinion_articles' THEN
        UPDATE opinion_articles SET views = views + 1 WHERE id = p_viewable_id;
    END IF;
END//
DELIMITER ;

-- Dumping structure for table lpmaarifnu_site.media
CREATE TABLE IF NOT EXISTS `media` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `file_name` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `original_name` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `file_path` varchar(500) COLLATE utf8mb4_unicode_ci NOT NULL,
  `file_url` varchar(500) COLLATE utf8mb4_unicode_ci NOT NULL,
  `file_type` varchar(50) COLLATE utf8mb4_unicode_ci NOT NULL,
  `mime_type` varchar(100) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `file_size` bigint unsigned NOT NULL COMMENT 'in bytes',
  `width` int unsigned DEFAULT NULL COMMENT 'for images',
  `height` int unsigned DEFAULT NULL COMMENT 'for images',
  `folder` varchar(100) COLLATE utf8mb4_unicode_ci DEFAULT 'general',
  `alt_text` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `caption` text COLLATE utf8mb4_unicode_ci,
  `uploaded_by` bigint unsigned DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_file_type` (`file_type`),
  KEY `idx_folder` (`folder`),
  KEY `idx_uploaded_by` (`uploaded_by`),
  CONSTRAINT `media_ibfk_1` FOREIGN KEY (`uploaded_by`) REFERENCES `users` (`id`) ON DELETE SET NULL
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Dumping data for table lpmaarifnu_site.media: ~3 rows (approximately)
DELETE FROM `media`;
INSERT INTO `media` (`id`, `file_name`, `original_name`, `file_path`, `file_url`, `file_type`, `mime_type`, `file_size`, `width`, `height`, `folder`, `alt_text`, `caption`, `uploaded_by`, `created_at`, `updated_at`, `deleted_at`) VALUES
	(1, 'hero-education.jpg', 'hero-education.jpg', '/media/images/hero-education.jpg', 'https://cdn.lpmaarifnu.or.id/images/hero-education.jpg', 'image', 'image/jpeg', 2048000, 1920, 1080, 'hero', 'Pendidikan Berkualitas', 'Siswa sedang belajar di kelas', 1, '2026-01-14 07:22:18', '2026-01-14 07:22:18', NULL),
	(2, 'beasiswa-2024.jpg', 'beasiswa-2024.jpg', '/media/images/beasiswa-2024.jpg', 'https://cdn.lpmaarifnu.or.id/images/beasiswa-2024.jpg', 'image', 'image/jpeg', 1536000, 1280, 720, 'news', 'Program Beasiswa 2024', 'Banner program beasiswa unggulan', 1, '2026-01-14 07:22:18', '2026-01-14 07:22:18', NULL),
	(3, 'seminar-pendidikan.jpg', 'seminar-pendidikan.jpg', '/media/images/seminar-pendidikan.jpg', 'https://cdn.lpmaarifnu.or.id/images/seminar-pendidikan.jpg', 'image', 'image/jpeg', 1843200, 1600, 900, 'events', 'Seminar Pendidikan', 'Peserta seminar pendidikan Islam', 2, '2026-01-14 07:22:18', '2026-01-14 07:22:18', NULL);

-- Dumping structure for table lpmaarifnu_site.news_articles
CREATE TABLE IF NOT EXISTS `news_articles` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `title` varchar(500) COLLATE utf8mb4_unicode_ci NOT NULL,
  `slug` varchar(500) COLLATE utf8mb4_unicode_ci NOT NULL,
  `excerpt` text COLLATE utf8mb4_unicode_ci NOT NULL,
  `content` longtext COLLATE utf8mb4_unicode_ci NOT NULL,
  `image` varchar(500) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `category_id` int unsigned DEFAULT NULL,
  `author_id` bigint unsigned DEFAULT NULL,
  `status` enum('draft','published','archived') COLLATE utf8mb4_unicode_ci DEFAULT 'draft',
  `published_at` timestamp NULL DEFAULT NULL,
  `views` int unsigned DEFAULT '0',
  `is_featured` tinyint(1) DEFAULT '0',
  `meta_title` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `meta_description` text COLLATE utf8mb4_unicode_ci,
  `meta_keywords` varchar(500) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `slug` (`slug`),
  KEY `category_id` (`category_id`),
  KEY `author_id` (`author_id`),
  KEY `idx_slug` (`slug`),
  KEY `idx_status` (`status`),
  KEY `idx_published_at` (`published_at`),
  KEY `idx_views` (`views`),
  KEY `idx_is_featured` (`is_featured`),
  KEY `idx_news_published` (`status`,`published_at`,`is_featured`),
  FULLTEXT KEY `idx_search` (`title`,`excerpt`,`content`),
  CONSTRAINT `news_articles_ibfk_1` FOREIGN KEY (`category_id`) REFERENCES `categories` (`id`) ON DELETE SET NULL,
  CONSTRAINT `news_articles_ibfk_2` FOREIGN KEY (`author_id`) REFERENCES `users` (`id`) ON DELETE SET NULL
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Dumping data for table lpmaarifnu_site.news_articles: ~7 rows (approximately)
DELETE FROM `news_articles`;
INSERT INTO `news_articles` (`id`, `title`, `slug`, `excerpt`, `content`, `image`, `category_id`, `author_id`, `status`, `published_at`, `views`, `is_featured`, `meta_title`, `meta_description`, `meta_keywords`, `created_at`, `updated_at`, `deleted_at`) VALUES
	(1, 'Peluncuran Program Beasiswa Unggulan 2024 untuk Siswa Berprestasi', 'peluncuran-program-beasiswa-unggulan-2024', 'LP Ma\'arif NU meluncurkan program beasiswa unggulan tahun 2024 yang ditujukan untuk siswa berprestasi dari keluarga kurang mampu di seluruh Indonesia.', '<h2>Program Beasiswa Unggulan 2024</h2><p>LP Ma\'arif NU dengan bangga mengumumkan peluncuran Program Beasiswa Unggulan 2024 yang merupakan komitmen kami dalam meningkatkan akses pendidikan berkualitas bagi siswa berprestasi dari berbagai latar belakang ekonomi.</p><h3>Latar Belakang</h3><p>Program ini diluncurkan sebagai respons terhadap kebutuhan akan dukungan pendidikan yang lebih luas dan merata di Indonesia. Dengan mempertimbangkan kondisi ekonomi yang masih menantang, LP Ma\'arif NU berupaya memberikan kesempatan kepada siswa-siswa berbakat untuk melanjutkan pendidikan mereka tanpa terkendala masalah finansial.</p><h3>Kriteria Penerima</h3><ul><li>Siswa aktif jenjang SMP/MTs dan SMA/MA</li><li>Memiliki prestasi akademik minimal rata-rata 8.0</li><li>Berasal dari keluarga kurang mampu (dibuktikan dengan SKTM)</li><li>Aktif dalam kegiatan ekstrakurikuler</li><li>Memiliki sikap dan karakter yang baik</li></ul><h3>Fasilitas yang Diberikan</h3><ol><li>Biaya pendidikan penuh selama 1 tahun ajaran</li><li>Uang saku bulanan</li><li>Bantuan buku dan alat tulis</li><li>Program mentoring dan pembinaan karakter</li><li>Kesempatan mengikuti pelatihan kepemimpinan</li></ol><h3>Cara Pendaftaran</h3><p>Pendaftaran dapat dilakukan secara online melalui website resmi LP Ma\'arif NU mulai tanggal 1 Februari 2024 hingga 30 Maret 2024. Berkas yang diperlukan meliputi:</p><ul><li>Formulir pendaftaran</li><li>Fotokopi rapor 2 semester terakhir</li><li>Surat Keterangan Tidak Mampu (SKTM)</li><li>Surat rekomendasi dari sekolah</li><li>Essay motivasi</li><li>Foto terbaru</li></ul><p>Untuk informasi lebih lanjut, silakan hubungi sekretariat LP Ma\'arif NU atau kunjungi website resmi kami.</p>', 'https://images.unsplash.com/photo-1523050854058-8df90110c9f1?w=800', 1, 1, 'published', '2026-01-12 07:22:18', 1520, 1, 'Program Beasiswa Unggulan 2024 LP Ma\'arif NU', 'LP Ma\'arif NU meluncurkan program beasiswa unggulan 2024 untuk siswa berprestasi. Daftar sekarang!', 'beasiswa, pendidikan, ma\'arif nu, siswa berprestasi, 2024', '2026-01-14 07:22:18', '2026-01-14 07:22:18', NULL),
	(2, 'Implementasi Kurikulum Merdeka di Sekolah Ma\'arif Se-Indonesia', 'implementasi-kurikulum-merdeka-sekolah-maarif', 'Lebih dari 5000 sekolah Ma\'arif di seluruh Indonesia telah berhasil mengimplementasikan Kurikulum Merdeka dengan pendampingan intensif dari LP Ma\'arif NU.', '<h2>Kurikulum Merdeka di Sekolah Ma\'arif</h2><p>Dalam upaya meningkatkan kualitas pendidikan di lingkungan Ma\'arif, LP Ma\'arif NU telah melakukan pendampingan implementasi Kurikulum Merdeka kepada lebih dari 5000 sekolah di seluruh Indonesia.</p><h3>Tahapan Implementasi</h3><p>Proses implementasi dilakukan secara bertahap dengan mempertimbangkan kesiapan setiap sekolah:</p><ol><li><strong>Fase Persiapan:</strong> Sosialisasi dan pelatihan guru</li><li><strong>Fase Pilot Project:</strong> Uji coba di sekolah percontohan</li><li><strong>Fase Implementasi:</strong> Pelaksanaan di seluruh sekolah</li><li><strong>Fase Evaluasi:</strong> Monitoring dan evaluasi berkelanjutan</li></ol><h3>Program Pendampingan</h3><p>LP Ma\'arif NU menyediakan program pendampingan komprehensif yang meliputi:</p><ul><li>Pelatihan guru tentang Kurikulum Merdeka</li><li>Workshop pengembangan modul ajar</li><li>Asesmen pembelajaran berbasis proyek</li><li>Pendampingan implementasi di kelas</li><li>Sharing session antar sekolah</li></ul><h3>Hasil dan Dampak</h3><p>Setelah satu semester implementasi, terlihat dampak positif yang signifikan:</p><ul><li>Peningkatan motivasi belajar siswa sebesar 40%</li><li>Guru lebih kreatif dalam mengembangkan pembelajaran</li><li>Siswa lebih aktif dan mandiri dalam belajar</li><li>Pembelajaran lebih kontekstual dan bermakna</li></ul>', 'https://images.unsplash.com/photo-1503676260728-1c00da094a0b?w=800', 1, 1, 'published', '2026-01-09 07:22:18', 2348, 1, 'Implementasi Kurikulum Merdeka Sekolah Ma\'arif', 'Lebih dari 5000 sekolah Ma\'arif telah sukses implementasi Kurikulum Merdeka dengan pendampingan LP Ma\'arif NU.', 'kurikulum merdeka, sekolah maarif, pendidikan, implementasi', '2026-01-14 07:22:18', '2026-01-15 15:51:46', NULL),
	(3, 'Pelatihan Guru Digital untuk Meningkatkan Kualitas Pembelajaran', 'pelatihan-guru-digital-kualitas-pembelajaran', 'LP Ma\'arif NU menyelenggarakan pelatihan guru digital yang diikuti oleh 2000 guru dari berbagai daerah untuk meningkatkan kompetensi dalam pembelajaran berbasis teknologi.', '<h2>Pelatihan Guru Digital</h2><p>Dalam era digital ini, kompetensi guru dalam menggunakan teknologi pembelajaran menjadi sangat penting. LP Ma\'arif NU menyelenggarakan program pelatihan guru digital yang komprehensif.</p><h3>Materi Pelatihan</h3><ul><li>Penggunaan Learning Management System (LMS)</li><li>Pembuatan konten pembelajaran digital</li><li>Penggunaan aplikasi pembelajaran interaktif</li><li>Asesmen online dan feedback digital</li><li>Keamanan digital dan etika online</li></ul><h3>Metodologi</h3><p>Pelatihan dilakukan dengan pendekatan blended learning:</p><ul><li>Sesi online melalui platform Zoom</li><li>Praktek langsung dengan pendampingan</li><li>Diskusi dan sharing session</li><li>Tugas proyek pembuatan konten</li><li>Presentasi hasil karya</li></ul>', 'https://images.unsplash.com/photo-1516321318423-f06f85e504b3?w=800', 2, 2, 'published', '2026-01-07 07:22:18', 1891, 1, 'Pelatihan Guru Digital LP Ma\'arif NU', 'Program pelatihan guru digital untuk meningkatkan kompetensi pembelajaran berbasis teknologi di era digital.', 'pelatihan guru, digital, teknologi, pembelajaran, ma\'arif', '2026-01-14 07:22:18', '2026-01-15 15:45:26', NULL),
	(4, 'Workshop Pengembangan Karakter Siswa Melalui Pendidikan Kepramukaan', 'workshop-pengembangan-karakter-pramuka', 'Workshop ini bertujuan mengintegrasikan nilai-nilai kepramukaan dalam pembentukan karakter siswa di lingkungan sekolah Ma\'arif.', '<h2>Workshop Karakter melalui Pramuka</h2><p>Pendidikan karakter merupakan aspek penting dalam pembentukan generasi muda yang berkualitas. Workshop ini dirancang untuk mengintegrasikan nilai-nilai kepramukaan dalam pendidikan karakter di sekolah.</p><h3>Nilai-Nilai Pramuka</h3><ul><li>Kedisiplinan dan tanggung jawab</li><li>Kerjasama dan kepemimpinan</li><li>Kemandirian dan kreativitas</li><li>Kepedulian terhadap lingkungan</li></ul>', 'https://images.unsplash.com/photo-1529390079861-591de354faf5?w=800', 3, 2, 'published', '2026-01-04 07:22:18', 890, 0, 'Workshop Karakter Siswa Melalui Pramuka', 'Workshop pengembangan karakter siswa melalui pendidikan kepramukaan di sekolah Ma\'arif.', 'pramuka, karakter, pendidikan, workshop, siswa', '2026-01-14 07:22:18', '2026-01-14 07:22:18', NULL),
	(5, 'Pelaksanaan Ujian Nasional Sekolah Ma\'arif Berjalan Lancar', 'pelaksanaan-ujian-nasional-maarif-lancar', 'Ujian Nasional di seluruh sekolah Ma\'arif berjalan dengan lancar dan tertib mengikuti protokol kesehatan yang ketat.', '<h2>Ujian Nasional Sekolah Ma\'arif</h2><p>Pelaksanaan Ujian Nasional di lingkungan sekolah Ma\'arif se-Indonesia berjalan dengan lancar dan tertib. Seluruh sekolah telah menerapkan protokol yang ketat untuk memastikan kelancaran ujian.</p><h3>Persiapan</h3><p>Persiapan yang matang dilakukan meliputi:</p><ul><li>Koordinasi dengan dinas pendidikan</li><li>Pengecekan sistem dan jaringan</li><li>Briefing pengawas</li><li>Persiapan ruang ujian</li></ul>', 'https://images.unsplash.com/photo-1434030216411-0b793f4b4173?w=800', 4, 1, 'published', '2025-12-30 07:22:18', 1234, 0, 'Ujian Nasional Sekolah Ma\'arif', 'Pelaksanaan Ujian Nasional di sekolah Ma\'arif berjalan lancar dengan protokol ketat.', 'ujian nasional, sekolah, ma\'arif, pendidikan', '2026-01-14 07:22:18', '2026-01-14 07:22:18', NULL),
	(6, 'Seminar Nasional Pendidikan Islam Modern di Era Digital', 'seminar-nasional-pendidikan-islam-modern', 'Seminar yang dihadiri 500 peserta ini membahas tantangan dan peluang pendidikan Islam di era digital dengan narasumber kompeten.', '<h2>Seminar Pendidikan Islam Modern</h2><p>LP Ma\'arif NU menyelenggarakan Seminar Nasional bertema "Pendidikan Islam Modern di Era Digital" yang dihadiri oleh 500 peserta dari berbagai daerah.</p><h3>Narasumber</h3><ul><li>Prof. Dr. KH. Ahmad Syafi\'i, MA - Pakar Pendidikan Islam</li><li>Dr. Miftahul Huda, M.Pd - Ahli Teknologi Pendidikan</li><li>Dr. Siti Aisyah, M.Pd - Praktisi Pendidikan</li></ul>', 'https://images.unsplash.com/photo-1540575467063-178a50c2df87?w=800', 3, 2, 'published', '2025-12-25 07:22:18', 2100, 0, 'Seminar Pendidikan Islam Modern', 'Seminar Nasional membahas tantangan dan peluang pendidikan Islam di era digital.', 'seminar, pendidikan islam, modern, digital, ma\'arif', '2026-01-14 07:22:18', '2026-01-14 07:22:18', NULL),
	(7, 'Kerjasama Internasional LP Ma\'arif NU dengan Universitas Malaysia', 'kerjasama-internasional-universitas-malaysia', 'Penandatanganan MoU antara LP Ma\'arif NU dengan Universitas Islam Malaysia membuka peluang pertukaran pelajar dan dosen.', '<h2>Kerjasama Internasional</h2><p>LP Ma\'arif NU menjalin kerjasama internasional dengan Universitas Islam Malaysia dalam bidang pendidikan dan penelitian.</p><h3>Program Kerjasama</h3><ul><li>Pertukaran pelajar</li><li>Pertukaran dosen</li><li>Penelitian bersama</li><li>Joint conference</li><li>Pengembangan kurikulum</li></ul>', 'https://images.unsplash.com/photo-1521737711867-e3b97375f902?w=800', 1, 1, 'published', '2025-12-20 07:22:18', 1567, 0, 'Kerjasama LP Ma\'arif NU dengan Malaysia', 'MoU LP Ma\'arif NU dengan Universitas Islam Malaysia untuk pertukaran pelajar dan dosen.', 'kerjasama, internasional, malaysia, ma\'arif, pendidikan', '2026-01-14 07:22:18', '2026-01-14 07:22:18', NULL);

-- Dumping structure for table lpmaarifnu_site.news_tags
CREATE TABLE IF NOT EXISTS `news_tags` (
  `news_article_id` bigint unsigned NOT NULL,
  `tag_id` int unsigned NOT NULL,
  PRIMARY KEY (`news_article_id`,`tag_id`),
  KEY `tag_id` (`tag_id`),
  CONSTRAINT `news_tags_ibfk_1` FOREIGN KEY (`news_article_id`) REFERENCES `news_articles` (`id`) ON DELETE CASCADE,
  CONSTRAINT `news_tags_ibfk_2` FOREIGN KEY (`tag_id`) REFERENCES `tags` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Dumping data for table lpmaarifnu_site.news_tags: ~20 rows (approximately)
DELETE FROM `news_tags`;
INSERT INTO `news_tags` (`news_article_id`, `tag_id`) VALUES
	(1, 1),
	(2, 1),
	(3, 1),
	(4, 1),
	(5, 1),
	(6, 1),
	(7, 1),
	(1, 2),
	(2, 3),
	(3, 4),
	(6, 5),
	(4, 7),
	(5, 7),
	(1, 8),
	(3, 9),
	(6, 9),
	(4, 10),
	(2, 12),
	(7, 13),
	(7, 15);

-- Dumping structure for table lpmaarifnu_site.notifications
CREATE TABLE IF NOT EXISTS `notifications` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `user_id` bigint unsigned NOT NULL,
  `type` varchar(100) COLLATE utf8mb4_unicode_ci NOT NULL,
  `title` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `message` text COLLATE utf8mb4_unicode_ci,
  `data` json DEFAULT NULL,
  `read_at` timestamp NULL DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `idx_user_id` (`user_id`),
  KEY `idx_read_at` (`read_at`),
  KEY `idx_created_at` (`created_at`),
  CONSTRAINT `notifications_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Dumping data for table lpmaarifnu_site.notifications: ~0 rows (approximately)
DELETE FROM `notifications`;

-- Dumping structure for table lpmaarifnu_site.opinion_articles
CREATE TABLE IF NOT EXISTS `opinion_articles` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `title` varchar(500) COLLATE utf8mb4_unicode_ci NOT NULL,
  `slug` varchar(500) COLLATE utf8mb4_unicode_ci NOT NULL,
  `excerpt` text COLLATE utf8mb4_unicode_ci NOT NULL,
  `content` longtext COLLATE utf8mb4_unicode_ci NOT NULL,
  `image` varchar(500) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `author_name` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `author_title` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `author_image` varchar(500) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `author_bio` text COLLATE utf8mb4_unicode_ci,
  `status` enum('draft','published','archived') COLLATE utf8mb4_unicode_ci DEFAULT 'draft',
  `published_at` timestamp NULL DEFAULT NULL,
  `views` int unsigned DEFAULT '0',
  `is_featured` tinyint(1) DEFAULT '0',
  `meta_title` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `meta_description` text COLLATE utf8mb4_unicode_ci,
  `meta_keywords` varchar(500) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `created_by` bigint unsigned DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `slug` (`slug`),
  KEY `created_by` (`created_by`),
  KEY `idx_slug` (`slug`),
  KEY `idx_status` (`status`),
  KEY `idx_published_at` (`published_at`),
  KEY `idx_views` (`views`),
  KEY `idx_is_featured` (`is_featured`),
  KEY `idx_opinion_published` (`status`,`published_at`,`is_featured`),
  FULLTEXT KEY `idx_search` (`title`,`excerpt`,`content`),
  CONSTRAINT `opinion_articles_ibfk_1` FOREIGN KEY (`created_by`) REFERENCES `users` (`id`) ON DELETE SET NULL
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Dumping data for table lpmaarifnu_site.opinion_articles: ~3 rows (approximately)
DELETE FROM `opinion_articles`;
INSERT INTO `opinion_articles` (`id`, `title`, `slug`, `excerpt`, `content`, `image`, `author_name`, `author_title`, `author_image`, `author_bio`, `status`, `published_at`, `views`, `is_featured`, `meta_title`, `meta_description`, `meta_keywords`, `created_by`, `created_at`, `updated_at`, `deleted_at`) VALUES
	(1, 'Pendidikan Karakter di Era Digital: Tantangan dan Solusi', 'pendidikan-karakter-era-digital', 'Pentingnya menanamkan nilai-nilai karakter kepada generasi muda di tengah gempuran teknologi digital yang semakin masif.', '<h2>Pendidikan Karakter di Era Digital</h2><p>Era digital membawa perubahan besar dalam cara kita mendidik generasi muda. Kemudahan akses informasi di satu sisi memberikan manfaat, namun di sisi lain juga membawa tantangan tersendiri dalam pembentukan karakter.</p><h3>Tantangan</h3><ul><li>Paparan konten negatif di media sosial</li><li>Kecanduan gadget</li><li>Berkurangnya interaksi sosial langsung</li><li>Instant gratification mindset</li></ul><h3>Solusi</h3><p>Untuk menghadapi tantangan tersebut, diperlukan strategi khusus:</p><ol><li>Pendidikan literasi digital</li><li>Pembentukan filter konten yang baik</li><li>Penanaman nilai-nilai agama dan budaya</li><li>Pendampingan intensif dari orangtua dan guru</li></ol>', 'https://images.unsplash.com/photo-1509062522246-3755977927d7?w=800', 'Prof. Dr. KH. Ahmad Syafi\'i, MA', 'Pakar Pendidikan Islam', 'https://ui-avatars.com/api/?name=Ahmad+Syafii&background=4CAF50&color=fff&size=200', 'Pakar pendidikan Islam dengan pengalaman lebih dari 25 tahun di bidang pendidikan. Aktif menulis dan memberikan ceramah tentang pendidikan karakter.', 'published', '2026-01-11 07:22:18', 897, 1, 'Pendidikan Karakter di Era Digital', 'Pentingnya pendidikan karakter di era digital dengan solusi praktis menghadapi tantangan teknologi.', 'pendidikan, karakter, digital, generasi muda', NULL, '2026-01-14 07:22:18', '2026-01-15 15:49:57', NULL),
	(2, 'Integrasi Nilai-Nilai Keislaman dalam Kurikulum Modern', 'integrasi-nilai-keislaman-kurikulum', 'Bagaimana mengintegrasikan nilai-nilai keislaman dalam kurikulum modern tanpa mengurangi kompetensi akademik siswa.', '<h2>Integrasi Nilai Keislaman</h2><p>Pendidikan Islam modern harus mampu mengintegrasikan nilai-nilai keislaman dengan perkembangan ilmu pengetahuan dan teknologi terkini.</p><h3>Prinsip Integrasi</h3><ul><li>Tidak dikotomi antara ilmu agama dan umum</li><li>Menggunakan pendekatan kontekstual</li><li>Menekankan praktik, bukan hanya teori</li><li>Mengembangkan critical thinking</li></ul>', 'https://images.unsplash.com/photo-1456513080510-7bf3a84b82f8?w=800', 'Dr. H. Muhammad Yusuf, M.Pd', 'Dosen Pendidikan Islam', 'https://ui-avatars.com/api/?name=Muhammad+Yusuf&background=2196F3&color=fff&size=200', 'Dosen di bidang Pendidikan Islam dengan fokus penelitian pada integrasi nilai keislaman dalam pendidikan modern.', 'published', '2026-01-06 07:22:18', 1245, 0, 'Integrasi Nilai Keislaman dalam Kurikulum', 'Strategi mengintegrasikan nilai-nilai keislaman dalam kurikulum pendidikan modern.', 'islam, kurikulum, pendidikan, integrasi', NULL, '2026-01-14 07:22:18', '2026-01-14 07:22:18', NULL),
	(3, 'Peran Guru dalam Membentuk Generasi Literat dan Berakhlak', 'peran-guru-membentuk-generasi-literat', 'Guru memiliki peran sentral dalam membentuk generasi yang tidak hanya literat tetapi juga memiliki akhlak mulia.', '<h2>Peran Guru</h2><p>Guru adalah garda terdepan dalam proses pendidikan. Peran guru sangat strategis dalam membentuk karakter dan kompetensi siswa.</p><h3>Tugas Guru Modern</h3><ul><li>Fasilitator pembelajaran</li><li>Motivator siswa</li><li>Role model</li><li>Inovator pendidikan</li></ul>', 'https://images.unsplash.com/photo-1524178232363-1fb2b075b655?w=800', 'Dr. Hj. Siti Aisyah, M.Pd', 'Praktisi Pendidikan', 'https://ui-avatars.com/api/?name=Siti+Aisyah&background=E91E63&color=fff&size=200', 'Praktisi pendidikan dengan pengalaman 20 tahun. Aktif dalam pengembangan kompetensi guru dan inovasi pembelajaran.', 'published', '2026-01-02 07:22:18', 678, 0, 'Peran Guru Membentuk Generasi Literat', 'Peran strategis guru dalam membentuk generasi yang literat dan berakhlak mulia.', 'guru, literasi, akhlak, pendidikan', NULL, '2026-01-14 07:22:18', '2026-01-14 07:22:18', NULL);

-- Dumping structure for table lpmaarifnu_site.opinion_tags
CREATE TABLE IF NOT EXISTS `opinion_tags` (
  `opinion_article_id` bigint unsigned NOT NULL,
  `tag_id` int unsigned NOT NULL,
  PRIMARY KEY (`opinion_article_id`,`tag_id`),
  KEY `tag_id` (`tag_id`),
  CONSTRAINT `opinion_tags_ibfk_1` FOREIGN KEY (`opinion_article_id`) REFERENCES `opinion_articles` (`id`) ON DELETE CASCADE,
  CONSTRAINT `opinion_tags_ibfk_2` FOREIGN KEY (`tag_id`) REFERENCES `tags` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Dumping data for table lpmaarifnu_site.opinion_tags: ~9 rows (approximately)
DELETE FROM `opinion_tags`;
INSERT INTO `opinion_tags` (`opinion_article_id`, `tag_id`) VALUES
	(1, 1),
	(2, 1),
	(3, 1),
	(2, 3),
	(1, 9),
	(1, 10),
	(3, 10),
	(3, 11),
	(2, 12);

-- Dumping structure for table lpmaarifnu_site.organization_positions
CREATE TABLE IF NOT EXISTS `organization_positions` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `position_name` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `position_level` int NOT NULL COMMENT '1=Ketua, 2=Wakil, 3=Sekretaris, 4=Bendahara, 5=Bidang',
  `position_type` enum('ketua','wakil','sekretaris','bendahara','bidang') COLLATE utf8mb4_unicode_ci NOT NULL,
  `parent_id` int unsigned DEFAULT NULL,
  `order_number` int DEFAULT '0',
  `is_active` tinyint(1) DEFAULT '1',
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `parent_id` (`parent_id`),
  KEY `idx_position_level` (`position_level`),
  KEY `idx_position_type` (`position_type`),
  KEY `idx_order_number` (`order_number`),
  CONSTRAINT `organization_positions_ibfk_1` FOREIGN KEY (`parent_id`) REFERENCES `organization_positions` (`id`) ON DELETE SET NULL
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Dumping data for table lpmaarifnu_site.organization_positions: ~8 rows (approximately)
DELETE FROM `organization_positions`;
INSERT INTO `organization_positions` (`id`, `position_name`, `position_level`, `position_type`, `parent_id`, `order_number`, `is_active`, `created_at`, `updated_at`) VALUES
	(1, 'Ketua Umum', 1, 'ketua', NULL, 1, 1, '2026-01-14 07:21:51', '2026-01-14 07:21:51'),
	(2, 'Wakil Ketua I', 2, 'wakil', NULL, 2, 1, '2026-01-14 07:21:51', '2026-01-14 07:21:51'),
	(3, 'Wakil Ketua II', 2, 'wakil', NULL, 3, 1, '2026-01-14 07:21:51', '2026-01-14 07:21:51'),
	(4, 'Sekretaris Umum', 3, 'sekretaris', NULL, 4, 1, '2026-01-14 07:21:51', '2026-01-14 07:21:51'),
	(5, 'Bendahara Umum', 4, 'bendahara', NULL, 5, 1, '2026-01-14 07:21:51', '2026-01-14 07:21:51'),
	(6, 'Wakil Ketua III', 2, 'wakil', NULL, 4, 1, '2026-01-14 07:22:18', '2026-01-14 07:22:18'),
	(7, 'Wakil Sekretaris', 3, 'sekretaris', 4, 5, 1, '2026-01-14 07:22:18', '2026-01-14 07:22:18'),
	(8, 'Wakil Bendahara', 4, 'bendahara', 5, 6, 1, '2026-01-14 07:22:18', '2026-01-14 07:22:18');

-- Dumping structure for table lpmaarifnu_site.pages
CREATE TABLE IF NOT EXISTS `pages` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `slug` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `title` varchar(500) COLLATE utf8mb4_unicode_ci NOT NULL,
  `content` longtext COLLATE utf8mb4_unicode_ci,
  `metadata` json DEFAULT NULL COMMENT 'Flexible content based on page type',
  `template` varchar(100) COLLATE utf8mb4_unicode_ci DEFAULT 'default',
  `is_active` tinyint(1) DEFAULT '1',
  `meta_title` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `meta_description` text COLLATE utf8mb4_unicode_ci,
  `meta_keywords` varchar(500) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `last_updated_by` bigint unsigned DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `slug` (`slug`),
  KEY `last_updated_by` (`last_updated_by`),
  KEY `idx_slug` (`slug`),
  KEY `idx_is_active` (`is_active`),
  CONSTRAINT `pages_ibfk_1` FOREIGN KEY (`last_updated_by`) REFERENCES `users` (`id`) ON DELETE SET NULL
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Dumping data for table lpmaarifnu_site.pages: ~4 rows (approximately)
DELETE FROM `pages`;
INSERT INTO `pages` (`id`, `slug`, `title`, `content`, `metadata`, `template`, `is_active`, `meta_title`, `meta_description`, `meta_keywords`, `last_updated_by`, `created_at`, `updated_at`) VALUES
	(1, 'visi-misi', 'Visi & Misi', NULL, '{"misi": ["Menyelenggarakan pendidikan berkualitas yang berbasis nilai-nilai keislaman", "Mengembangkan kurikulum yang integratif antara ilmu agama dan ilmu umum", "Memberdayakan sumber daya manusia yang profesional dan berakhlak mulia", "Membangun kerjasama dengan berbagai pihak untuk kemajuan pendidikan", "Mengimplementasikan teknologi dalam pembelajaran"], "visi": "Menjadi lembaga pendidikan Islam yang unggul, modern, dan berkarakter untuk mewujudkan generasi yang berakhlak mulia, cerdas, dan berdaya saing global.", "nilai_nilai": [{"icon": "integrity", "title": "Integritas", "description": "Menjunjung tinggi kejujuran dan tanggung jawab dalam setiap aspek"}, {"icon": "professionalism", "title": "Profesionalisme", "description": "Menjalankan tugas dengan kompeten dan penuh dedikasi"}, {"icon": "innovation", "title": "Inovasi", "description": "Senantiasa berinovasi untuk kemajuan pendidikan"}, {"icon": "collaboration", "title": "Kolaborasi", "description": "Membangun kerjasama yang sinergis dengan berbagai pihak"}]}', 'visi-misi', 1, NULL, NULL, NULL, NULL, '2026-01-14 07:21:51', '2026-01-14 07:22:18'),
	(2, 'sejarah', 'Sejarah LP Ma\'arif NU', NULL, '{"timeline": [{"year": "1916", "title": "Berdirinya Nahdlatul Ulama", "description": "Nahdlatul Ulama (NU) didirikan di Surabaya oleh KH. Hasyim Asy\'ari sebagai organisasi kemasyarakatan Islam."}, {"year": "1926", "title": "Pembentukan Lembaga Pendidikan", "description": "NU mulai fokus pada pengembangan lembaga pendidikan Islam di berbagai daerah."}, {"year": "1960", "title": "Formalisasi LP Ma\'arif", "description": "Pembentukan formal LP Ma\'arif NU sebagai lembaga yang menaungi pendidikan di lingkungan NU."}, {"year": "1990", "title": "Ekspansi Nasional", "description": "Perluasan jangkauan LP Ma\'arif ke seluruh Indonesia dengan ribuan sekolah di bawah naungannya."}, {"year": "2010", "title": "Modernisasi Pendidikan", "description": "Implementasi teknologi dan kurikulum modern dengan tetap mempertahankan nilai-nilai keislaman."}, {"year": "2024", "title": "Transformasi Digital", "description": "Akselerasi digitalisasi pendidikan dan peningkatan kualitas pembelajaran di era digital."}], "introduction": "<p>LP Ma\'arif NU merupakan lembaga pendidikan di bawah naungan Nahdlatul Ulama yang telah berkiprah dalam dunia pendidikan Islam di Indonesia sejak tahun 1916.</p>"}', 'sejarah', 1, NULL, NULL, NULL, NULL, '2026-01-14 07:21:51', '2026-01-14 07:22:18'),
	(3, 'program-strategis', 'Program Strategis', NULL, '{"programs": [{"icon": "teaching", "title": "Peningkatan Kualitas Pembelajaran", "description": "Program pelatihan dan pendampingan guru untuk meningkatkan kualitas pembelajaran di kelas."}, {"icon": "digital", "title": "Digitalisasi Pendidikan", "description": "Implementasi teknologi dalam pembelajaran untuk meningkatkan efektivitas dan efisiensi."}, {"icon": "curriculum", "title": "Pengembangan Kurikulum Integratif", "description": "Pengembangan kurikulum yang mengintegrasikan nilai-nilai keislaman dengan ilmu pengetahuan modern."}, {"icon": "scholarship", "title": "Beasiswa Pendidikan", "description": "Program beasiswa untuk siswa berprestasi dari keluarga kurang mampu."}]}', 'program-strategis', 1, NULL, NULL, NULL, NULL, '2026-01-14 07:21:51', '2026-01-14 07:22:18'),
	(4, 'pramuka', 'Pramuka Ma\'arif', NULL, '{"slug": "pramuka", "title": "Pramuka Ma\'arif", "content": {"cta": {"title": "Bergabung dengan Pramuka Ma\'arif NU", "description": "Mari bersama membangun karakter generasi muda yang berakhlak mulia dan cinta tanah air"}, "hero": {"title": "Gerakan Pramuka Ma\'arif NU", "description": "Membentuk karakter pemuda yang berakhlak mulia, cinta tanah air, dan berwawasan keislaman"}, "about": {"image": "https://images.unsplash.com/photo-1509062522246-3755977927d7?w=800&h=600&fit=crop", "title": "Tentang Pramuka Ma\'arif", "paragraphs": ["Gerakan Pramuka di lingkungan satuan pendidikan LP Ma\'arif NU merupakan wadah pembinaan karakter dan kepribadian siswa yang berlandaskan nilai-nilai Pancasila dan Ahlussunnah Wal Jama\'ah an-Nahdliyyah.", "Melalui berbagai kegiatan kepramukaan, kami membentuk generasi muda yang memiliki jiwa kepemimpinan, tanggung jawab, dan kepedulian terhadap sesama dan lingkungan.", "Dengan jaringan lebih dari 5.000 gugus depan di seluruh Indonesia, Pramuka Ma\'arif aktif dalam berbagai kegiatan nasional dan internasional."]}, "programs": {"list": [{"icon": "target", "title": "Pembinaan Karakter", "description": "Program pembinaan karakter melalui kegiatan kepramukaan yang terintegrasi dengan nilai-nilai Islam"}, {"icon": "users", "title": "Pelatihan Kepemimpinan", "description": "Mengembangkan jiwa kepemimpinan dan kemampuan organisasi siswa"}, {"icon": "award", "title": "Kompetisi & Lomba", "description": "Mengikuti berbagai kompetisi kepramukaan tingkat daerah hingga nasional"}, {"icon": "calendar", "title": "Kegiatan Rutin", "description": "Latihan rutin, perkemahan, dan kegiatan sosial berkelanjutan"}], "title": "Program Unggulan", "description": "Berbagai program pembinaan yang dirancang untuk mengembangkan potensi siswa"}, "achievements": {"list": [{"date": "2023-08-15", "image": "https://images.unsplash.com/photo-1519995451813-39e29e054914?w=800&h=600&fit=crop", "title": "Juara Umum Jambore Nasional 2023", "description": "Kontingen LP Ma\'arif NU meraih juara umum dalam Jambore Nasional Pramuka 2023"}, {"date": "2023-10-20", "image": "https://images.unsplash.com/photo-1529390079861-591de354faf5?w=800&h=600&fit=crop", "title": "Pelatihan Instruktur Nasional", "description": "Mengadakan pelatihan instruktur pramuka tingkat nasional dengan 500 peserta"}, {"date": "2024-03-12", "image": "https://images.unsplash.com/photo-1542601906990-b4d3fb778b09?w=800&h=600&fit=crop", "title": "Bakti Sosial Lingkungan", "description": "Program tanam 10.000 pohon oleh gerakan pramuka Ma\'arif se-Indonesia"}], "title": "Prestasi & Kegiatan", "description": "Pencapaian membanggakan dari gerakan pramuka Ma\'arif NU"}}, "description": "Gerakan Pramuka di lingkungan satuan pendidikan LP Ma\'arif NU"}', 'pramuka', 1, NULL, NULL, NULL, NULL, '2026-01-14 07:21:51', '2026-01-15 15:14:33');

-- Dumping structure for table lpmaarifnu_site.page_views
CREATE TABLE IF NOT EXISTS `page_views` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `viewable_type` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `viewable_id` bigint unsigned NOT NULL,
  `ip_address` varchar(45) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `user_agent` text COLLATE utf8mb4_unicode_ci,
  `referer` varchar(500) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `session_id` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `viewed_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `idx_viewable` (`viewable_type`,`viewable_id`),
  KEY `idx_viewed_at` (`viewed_at`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Dumping data for table lpmaarifnu_site.page_views: ~4 rows (approximately)
DELETE FROM `page_views`;
INSERT INTO `page_views` (`id`, `viewable_type`, `viewable_id`, `ip_address`, `user_agent`, `referer`, `session_id`, `viewed_at`) VALUES
	(1, 'news_articles', 1, '192.168.1.100', 'Mozilla/5.0', NULL, NULL, '2026-01-13 07:22:18'),
	(2, 'news_articles', 1, '192.168.1.101', 'Mozilla/5.0', NULL, NULL, '2026-01-12 07:22:18'),
	(3, 'news_articles', 2, '192.168.1.102', 'Mozilla/5.0', NULL, NULL, '2026-01-13 07:22:18'),
	(4, 'opinion_articles', 1, '192.168.1.103', 'Mozilla/5.0', NULL, NULL, '2026-01-11 07:22:18');

-- Dumping structure for table lpmaarifnu_site.password_resets
CREATE TABLE IF NOT EXISTS `password_resets` (
  `email` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `token` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  KEY `idx_email` (`email`),
  KEY `idx_token` (`token`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Dumping data for table lpmaarifnu_site.password_resets: ~0 rows (approximately)
DELETE FROM `password_resets`;

-- Dumping structure for table lpmaarifnu_site.pengurus
CREATE TABLE IF NOT EXISTS `pengurus` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `nama` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `jabatan` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `kategori` enum('pimpinan_utama','bidang','sekretariat','bendahara') COLLATE utf8mb4_unicode_ci DEFAULT 'bidang',
  `foto` varchar(500) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `bio` text COLLATE utf8mb4_unicode_ci,
  `email` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `phone` varchar(20) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `periode_mulai` year NOT NULL,
  `periode_selesai` year NOT NULL,
  `order_number` int DEFAULT '0',
  `is_active` tinyint(1) DEFAULT '1',
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `idx_kategori` (`kategori`),
  KEY `idx_periode` (`periode_mulai`,`periode_selesai`),
  KEY `idx_order_number` (`order_number`),
  KEY `idx_is_active` (`is_active`)
) ENGINE=InnoDB AUTO_INCREMENT=13 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Dumping data for table lpmaarifnu_site.pengurus: ~12 rows (approximately)
DELETE FROM `pengurus`;
INSERT INTO `pengurus` (`id`, `nama`, `jabatan`, `kategori`, `foto`, `bio`, `email`, `phone`, `periode_mulai`, `periode_selesai`, `order_number`, `is_active`, `created_at`, `updated_at`) VALUES
	(1, 'Prof. Dr. KH. Said Aqil Siradj, MA', 'Ketua Umum', 'pimpinan_utama', 'https://ui-avatars.com/api/?name=Said+Aqil+Siradj&background=059669&color=fff&size=400', 'Ulama dan intelektual muslim Indonesia, Ketua Umum PBNU', 'ketua@lpmaarifnu.or.id', '021-3920677', '2024', '2029', 1, 1, '2026-01-14 07:22:18', '2026-01-14 07:22:18'),
	(2, 'Dr. H. Ahmad Lutfi, M.Pd', 'Wakil Ketua I', 'pimpinan_utama', 'https://ui-avatars.com/api/?name=Ahmad+Lutfi&background=2563EB&color=fff&size=400', 'Pakar pendidikan Islam dengan pengalaman lebih dari 20 tahun', 'wakil1@lpmaarifnu.or.id', '021-3920678', '2024', '2029', 2, 1, '2026-01-14 07:22:18', '2026-01-14 07:22:18'),
	(3, 'Drs. H. Mahfudz Siddiq, M.Si', 'Wakil Ketua II', 'pimpinan_utama', 'https://ui-avatars.com/api/?name=Mahfudz+Siddiq&background=7C3AED&color=fff&size=400', 'Praktisi pendidikan dan manajemen organisasi', 'wakil2@lpmaarifnu.or.id', '021-3920679', '2024', '2029', 3, 1, '2026-01-14 07:22:18', '2026-01-14 07:22:18'),
	(4, 'Dr. Hj. Siti Aisyah, M.Pd', 'Sekretaris Umum', 'sekretariat', 'https://ui-avatars.com/api/?name=Siti+Aisyah&background=DB2777&color=fff&size=400', 'Ahli administrasi dan manajemen pendidikan', 'sekretaris@lpmaarifnu.or.id', '021-3920680', '2024', '2029', 4, 1, '2026-01-14 07:22:18', '2026-01-14 07:22:18'),
	(5, 'H. Abdul Rahman, SE, M.Ak', 'Bendahara Umum', 'bendahara', 'https://ui-avatars.com/api/?name=Abdul+Rahman&background=EA580C&color=fff&size=400', 'Akuntan profesional dan auditor', 'bendahara@lpmaarifnu.or.id', '021-3920681', '2024', '2029', 5, 1, '2026-01-14 07:22:18', '2026-01-14 07:22:18'),
	(6, 'Dr. Muhammad Yusuf, M.Pd', 'Kepala Bidang Pendidikan Dasar', 'bidang', 'https://ui-avatars.com/api/?name=Muhammad+Yusuf&background=16A34A&color=fff&size=400', 'Spesialis pendidikan dasar dan kurikulum', 'pendas@lpmaarifnu.or.id', NULL, '2024', '2029', 6, 1, '2026-01-14 07:22:18', '2026-01-14 07:22:18'),
	(7, 'Dra. Hj. Nur Aini, M.Pd', 'Kepala Bidang Pendidikan Menengah', 'bidang', 'https://ui-avatars.com/api/?name=Nur+Aini&background=8B5CF6&color=fff&size=400', 'Pakar pendidikan menengah dan psikologi remaja', 'penmen@lpmaarifnu.or.id', NULL, '2024', '2029', 7, 1, '2026-01-14 07:22:18', '2026-01-14 07:22:18'),
	(8, 'Prof. Dr. Ahmad Syahid, M.A', 'Kepala Bidang Pendidikan Tinggi', 'bidang', 'https://ui-avatars.com/api/?name=Ahmad+Syahid&background=0891B2&color=fff&size=400', 'Profesor pendidikan tinggi dan peneliti', 'pendikti@lpmaarifnu.or.id', NULL, '2024', '2029', 8, 1, '2026-01-14 07:22:18', '2026-01-14 07:22:18'),
	(9, 'Dr. Hj. Fatimah Azzahra, M.Pd', 'Kepala Bidang Kurikulum', 'bidang', 'https://ui-avatars.com/api/?name=Fatimah+Azzahra&background=DC2626&color=fff&size=400', 'Ahli pengembangan kurikulum dan asesmen', 'kurikulum@lpmaarifnu.or.id', NULL, '2024', '2029', 9, 1, '2026-01-14 07:22:18', '2026-01-14 07:22:18'),
	(10, 'H. Zainuddin Maliki, M.M', 'Kepala Bidang SDM dan Kemitraan', 'bidang', 'https://ui-avatars.com/api/?name=Zainuddin+Maliki&background=F59E0B&color=fff&size=400', 'Praktisi manajemen SDM dan kerjasama kelembagaan', 'sdm@lpmaarifnu.or.id', NULL, '2024', '2029', 10, 1, '2026-01-14 07:22:18', '2026-01-14 07:22:18'),
	(11, 'Dr. Ir. Habiburrahman, M.T', 'Kepala Bidang Litbang', 'bidang', 'https://ui-avatars.com/api/?name=Habiburrahman&background=6366F1&color=fff&size=400', 'Peneliti dan pengembang teknologi pendidikan', 'litbang@lpmaarifnu.or.id', NULL, '2024', '2029', 11, 1, '2026-01-14 07:22:18', '2026-01-14 07:22:18'),
	(12, 'Drs. H. Miftahul Huda, M.Pd.I', 'Kepala Bidang Humas & Publikasi', 'bidang', 'https://ui-avatars.com/api/?name=Miftahul+Huda&background=10B981&color=fff&size=400', 'Ahli komunikasi dan public relations', 'humas@lpmaarifnu.or.id', NULL, '2024', '2029', 12, 1, '2026-01-14 07:22:18', '2026-01-14 07:22:18');

-- Dumping structure for table lpmaarifnu_site.personal_access_tokens
CREATE TABLE IF NOT EXISTS `personal_access_tokens` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `tokenable_type` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `tokenable_id` bigint unsigned NOT NULL,
  `name` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `token` varchar(64) COLLATE utf8mb4_unicode_ci NOT NULL,
  `abilities` text COLLATE utf8mb4_unicode_ci,
  `last_used_at` timestamp NULL DEFAULT NULL,
  `expires_at` timestamp NULL DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `token` (`token`),
  KEY `idx_tokenable` (`tokenable_type`,`tokenable_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Dumping data for table lpmaarifnu_site.personal_access_tokens: ~0 rows (approximately)
DELETE FROM `personal_access_tokens`;

-- Dumping structure for table lpmaarifnu_site.settings
CREATE TABLE IF NOT EXISTS `settings` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `setting_key` varchar(100) COLLATE utf8mb4_unicode_ci NOT NULL,
  `setting_value` text COLLATE utf8mb4_unicode_ci,
  `setting_type` enum('string','text','number','boolean','json') COLLATE utf8mb4_unicode_ci DEFAULT 'string',
  `setting_group` varchar(50) COLLATE utf8mb4_unicode_ci DEFAULT 'general',
  `description` text COLLATE utf8mb4_unicode_ci,
  `is_public` tinyint(1) DEFAULT '0' COMMENT 'Can be accessed without auth',
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `setting_key` (`setting_key`),
  KEY `idx_setting_key` (`setting_key`),
  KEY `idx_setting_group` (`setting_group`),
  KEY `idx_is_public` (`is_public`)
) ENGINE=InnoDB AUTO_INCREMENT=10 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Dumping data for table lpmaarifnu_site.settings: ~9 rows (approximately)
DELETE FROM `settings`;
INSERT INTO `settings` (`id`, `setting_key`, `setting_value`, `setting_type`, `setting_group`, `description`, `is_public`, `created_at`, `updated_at`) VALUES
	(1, 'site_name', 'LP Ma\'arif NU', 'string', 'general', 'Nama website', 1, '2026-01-14 07:21:51', '2026-01-14 07:21:51'),
	(2, 'site_description', 'Lembaga Pendidikan Ma\'arif Nahdlatul Ulama - Membangun Pendidikan Islam Berkualitas', 'text', 'general', 'Deskripsi website', 1, '2026-01-14 07:21:51', '2026-01-14 07:22:18'),
	(3, 'contact_email', 'infoo@lpmaarifnu.or.id', 'string', 'contact', 'Email kontak', 1, '2026-01-14 07:21:51', '2026-01-15 15:50:55'),
	(4, 'contact_phone', '(021) 3920677', 'string', 'contact', 'Nomor telepon', 1, '2026-01-14 07:21:51', '2026-01-14 07:22:18'),
	(5, 'contact_address', 'Jalan Kramat Raya No. 164, Jakarta Pusat 10430', 'text', 'contact', 'Alamat', 1, '2026-01-14 07:21:51', '2026-01-15 15:51:16'),
	(6, 'social_facebook', 'https://facebook.com/lpmaarifnuu', 'string', 'social', 'Facebook URL', 1, '2026-01-14 07:21:51', '2026-01-15 15:51:44'),
	(7, 'social_twitter', 'https://twitter.com/lpmaarifnu', 'string', 'social', 'Twitter URL', 1, '2026-01-14 07:21:51', '2026-01-14 07:22:18'),
	(8, 'social_instagram', 'https://instagram.com/lpmaarifnu', 'string', 'social', 'Instagram URL', 1, '2026-01-14 07:21:51', '2026-01-14 07:22:18'),
	(9, 'social_youtube', 'https://youtube.com/@lpmaarifnu', 'string', 'social', 'YouTube URL', 1, '2026-01-14 07:21:51', '2026-01-14 07:22:18');

-- Dumping structure for table lpmaarifnu_site.tags
CREATE TABLE IF NOT EXISTS `tags` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(100) COLLATE utf8mb4_unicode_ci NOT NULL,
  `slug` varchar(100) COLLATE utf8mb4_unicode_ci NOT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `slug` (`slug`),
  KEY `idx_slug` (`slug`)
) ENGINE=InnoDB AUTO_INCREMENT=16 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Dumping data for table lpmaarifnu_site.tags: ~15 rows (approximately)
DELETE FROM `tags`;
INSERT INTO `tags` (`id`, `name`, `slug`, `created_at`, `updated_at`) VALUES
	(1, 'Pendidikan', 'pendidikan', '2026-01-14 07:22:18', '2026-01-14 07:22:18'),
	(2, 'Beasiswa', 'beasiswa', '2026-01-14 07:22:18', '2026-01-14 07:22:18'),
	(3, 'Kurikulum', 'kurikulum', '2026-01-14 07:22:18', '2026-01-14 07:22:18'),
	(4, 'Pelatihan', 'pelatihan', '2026-01-14 07:22:18', '2026-01-14 07:22:18'),
	(5, 'Seminar', 'seminar', '2026-01-14 07:22:18', '2026-01-14 07:22:18'),
	(6, 'Workshop', 'workshop', '2026-01-14 07:22:18', '2026-01-14 07:22:18'),
	(7, 'Kegiatan', 'kegiatan', '2026-01-14 07:22:18', '2026-01-14 07:22:18'),
	(8, 'Prestasi', 'prestasi', '2026-01-14 07:22:18', '2026-01-14 07:22:18'),
	(9, 'Teknologi', 'teknologi', '2026-01-14 07:22:18', '2026-01-14 07:22:18'),
	(10, 'Karakter', 'karakter', '2026-01-14 07:22:18', '2026-01-14 07:22:18'),
	(11, 'Literasi', 'literasi', '2026-01-14 07:22:18', '2026-01-14 07:22:18'),
	(12, 'Inovasi', 'inovasi', '2026-01-14 07:22:18', '2026-01-14 07:22:18'),
	(13, 'Kerjasama', 'kerjasama', '2026-01-14 07:22:18', '2026-01-14 07:22:18'),
	(14, 'Pengabdian', 'pengabdian', '2026-01-14 07:22:18', '2026-01-14 07:22:18'),
	(15, 'Penelitian', 'penelitian', '2026-01-14 07:22:18', '2026-01-14 07:22:18');

-- Dumping structure for table lpmaarifnu_site.users
CREATE TABLE IF NOT EXISTS `users` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `email` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `password` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `role` enum('super_admin','admin','editor','viewer') COLLATE utf8mb4_unicode_ci DEFAULT 'viewer',
  `avatar` varchar(500) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `phone` varchar(20) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `is_active` tinyint(1) DEFAULT '1',
  `last_login_at` timestamp NULL DEFAULT NULL,
  `email_verified_at` timestamp NULL DEFAULT NULL,
  `remember_token` varchar(100) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `email` (`email`),
  KEY `idx_email` (`email`),
  KEY `idx_role` (`role`),
  KEY `idx_is_active` (`is_active`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Dumping data for table lpmaarifnu_site.users: ~4 rows (approximately)
DELETE FROM `users`;
INSERT INTO `users` (`id`, `name`, `email`, `password`, `role`, `avatar`, `phone`, `is_active`, `last_login_at`, `email_verified_at`, `remember_token`, `created_at`, `updated_at`, `deleted_at`) VALUES
	(1, 'Super Admin', 'admin@lpmaarifnu.or.id', '$2y$10$92IXUNpkjO0rOQ5byMi.Ye4oKoEa3Ro9llC/.og/at2.uheWG/igi', 'super_admin', NULL, NULL, 1, NULL, NULL, NULL, '2026-01-14 07:21:51', '2026-01-14 07:21:51', NULL),
	(2, 'Editor Berita', 'editor@lpmaarifnu.or.id', '$2y$10$92IXUNpkjO0rOQ5byMi.Ye4oKoEa3Ro9llC/.og/at2.uheWG/igi', 'editor', 'https://ui-avatars.com/api/?name=Editor+Berita', '081234567890', 1, NULL, '2026-01-14 07:22:18', NULL, '2026-01-14 07:22:18', '2026-01-14 07:22:18', NULL),
	(3, 'Admin Dokumen', 'dokumen@lpmaarifnu.or.id', '$2y$10$92IXUNpkjO0rOQ5byMi.Ye4oKoEa3Ro9llC/.og/at2.uheWG/igi', 'admin', 'https://ui-avatars.com/api/?name=Admin+Dokumen', '081234567891', 1, NULL, '2026-01-14 07:22:18', NULL, '2026-01-14 07:22:18', '2026-01-14 07:22:18', NULL),
	(4, 'Viewer User', 'viewer@lpmaarifnu.or.id', '$2y$10$92IXUNpkjO0rOQ5byMi.Ye4oKoEa3Ro9llC/.og/at2.uheWG/igi', 'viewer', 'https://ui-avatars.com/api/?name=Viewer+User', '081234567892', 1, NULL, '2026-01-14 07:22:18', NULL, '2026-01-14 07:22:18', '2026-01-14 07:22:18', NULL);

-- Dumping structure for view lpmaarifnu_site.v_published_news
-- Creating temporary table to overcome VIEW dependency errors
CREATE TABLE `v_published_news` (
	`id` BIGINT(20) UNSIGNED NOT NULL,
	`title` VARCHAR(500) NOT NULL COLLATE 'utf8mb4_unicode_ci',
	`slug` VARCHAR(500) NOT NULL COLLATE 'utf8mb4_unicode_ci',
	`excerpt` TEXT NOT NULL COLLATE 'utf8mb4_unicode_ci',
	`content` LONGTEXT NOT NULL COLLATE 'utf8mb4_unicode_ci',
	`image` VARCHAR(500) NULL COLLATE 'utf8mb4_unicode_ci',
	`published_at` TIMESTAMP NULL,
	`views` INT(10) UNSIGNED NULL,
	`is_featured` TINYINT(1) NULL,
	`category_name` VARCHAR(100) NULL COLLATE 'utf8mb4_unicode_ci',
	`category_slug` VARCHAR(100) NULL COLLATE 'utf8mb4_unicode_ci',
	`author_name` VARCHAR(255) NULL COLLATE 'utf8mb4_unicode_ci',
	`tags` TEXT NULL COLLATE 'utf8mb4_unicode_ci'
) ENGINE=MyISAM;

-- Dumping structure for trigger lpmaarifnu_site.after_document_download
SET @OLDTMP_SQL_MODE=@@SQL_MODE, SQL_MODE='ONLY_FULL_GROUP_BY,STRICT_TRANS_TABLES,NO_ZERO_IN_DATE,NO_ZERO_DATE,ERROR_FOR_DIVISION_BY_ZERO,NO_ENGINE_SUBSTITUTION';
DELIMITER //
CREATE TRIGGER `after_document_download` AFTER INSERT ON `download_logs` FOR EACH ROW BEGIN
    UPDATE documents
    SET download_count = download_count + 1
    WHERE id = NEW.document_id;
END//
DELIMITER ;
SET SQL_MODE=@OLDTMP_SQL_MODE;

-- Dumping structure for trigger lpmaarifnu_site.news_articles_before_insert
SET @OLDTMP_SQL_MODE=@@SQL_MODE, SQL_MODE='ONLY_FULL_GROUP_BY,STRICT_TRANS_TABLES,NO_ZERO_IN_DATE,NO_ZERO_DATE,ERROR_FOR_DIVISION_BY_ZERO,NO_ENGINE_SUBSTITUTION';
DELIMITER //
CREATE TRIGGER `news_articles_before_insert` BEFORE INSERT ON `news_articles` FOR EACH ROW BEGIN
    IF NEW.slug IS NULL OR NEW.slug = '' THEN
        SET NEW.slug = LOWER(REPLACE(NEW.title, ' ', '-'));
    END IF;
END//
DELIMITER ;
SET SQL_MODE=@OLDTMP_SQL_MODE;

-- Dumping structure for view lpmaarifnu_site.v_published_news
-- Removing temporary table and create final VIEW structure
DROP TABLE IF EXISTS `v_published_news`;
CREATE ALGORITHM=UNDEFINED SQL SECURITY DEFINER VIEW `v_published_news` AS select `n`.`id` AS `id`,`n`.`title` AS `title`,`n`.`slug` AS `slug`,`n`.`excerpt` AS `excerpt`,`n`.`content` AS `content`,`n`.`image` AS `image`,`n`.`published_at` AS `published_at`,`n`.`views` AS `views`,`n`.`is_featured` AS `is_featured`,`c`.`name` AS `category_name`,`c`.`slug` AS `category_slug`,`u`.`name` AS `author_name`,group_concat(`t`.`name` separator ',') AS `tags` from ((((`news_articles` `n` left join `categories` `c` on((`n`.`category_id` = `c`.`id`))) left join `users` `u` on((`n`.`author_id` = `u`.`id`))) left join `news_tags` `nt` on((`n`.`id` = `nt`.`news_article_id`))) left join `tags` `t` on((`nt`.`tag_id` = `t`.`id`))) where ((`n`.`status` = 'published') and (`n`.`deleted_at` is null)) group by `n`.`id`;

/*!40103 SET TIME_ZONE=IFNULL(@OLD_TIME_ZONE, 'system') */;
/*!40101 SET SQL_MODE=IFNULL(@OLD_SQL_MODE, '') */;
/*!40014 SET FOREIGN_KEY_CHECKS=IFNULL(@OLD_FOREIGN_KEY_CHECKS, 1) */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40111 SET SQL_NOTES=IFNULL(@OLD_SQL_NOTES, 1) */;
