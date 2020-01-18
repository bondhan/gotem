CREATE TABLE `m_country` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  `country_name` varchar(6) NOT NULL,
  `country_code` varchar(6) NOT NULL,
  PRIMARY KEY (`id`),
  KEY `m_country_deleted_at` (`deleted_at`)
) ENGINE = InnoDB DEFAULT CHARSET = latin1;
