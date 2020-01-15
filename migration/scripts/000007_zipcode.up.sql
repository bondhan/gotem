CREATE TABLE `m_zipcode` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  `zip_code` varchar(6) NOT NULL,
  `city_id` int(10) unsigned NOT NULL,
  PRIMARY KEY (`id`),
  KEY `m_zipcode_deleted_at` (`deleted_at`),
  CONSTRAINT `fk_m_zipcode_city_id` FOREIGN KEY (`city_id`) REFERENCES `m_city` (`id`) ON DELETE RESTRICT ON UPDATE CASCADE
) ENGINE = InnoDB DEFAULT CHARSET = latin1;
