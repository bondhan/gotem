CREATE TABLE `m_city` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  `city_name` varchar(255) NOT NULL,
  `province_id` int(10) unsigned NOT NULL,
  PRIMARY KEY (`id`),
  KEY `m_city_deleted_at` (`deleted_at`),
  INDEX `m_city_index_province_id` (`province_id` ASC),
  CONSTRAINT `fk_m_city_province_id` FOREIGN KEY (`province_id`) REFERENCES `m_province` (`id`) ON DELETE RESTRICT ON UPDATE CASCADE
) ENGINE = InnoDB DEFAULT CHARSET = latin1;
