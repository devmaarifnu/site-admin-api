-- Migration: Add token_blacklist table for logout functionality
-- Created: 2026-01-31

CREATE TABLE IF NOT EXISTS `token_blacklist` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `token` varchar(500) COLLATE utf8mb4_unicode_ci NOT NULL,
  `expires_at` timestamp NOT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_token` (`token`),
  KEY `idx_expires_at` (`expires_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Cleanup job (optional - to be run periodically via cron)
-- DELETE FROM token_blacklist WHERE expires_at < NOW();
