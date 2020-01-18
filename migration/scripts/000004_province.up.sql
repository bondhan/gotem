CREATE TABLE `m_province` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  `province_name` varchar(255) NOT NULL,
  `province_code` varchar(3) NOT NULL,
  `country_id` int(10) unsigned NOT NULL,
  PRIMARY KEY (`id`),
  KEY `m_province_deleted_at` (`deleted_at`),
  INDEX `m_province_index_country_id` (`country_id` ASC),
  CONSTRAINT `fk_m_province` FOREIGN KEY (`country_id`) REFERENCES `m_country` (`id`) ON DELETE RESTRICT ON UPDATE CASCADE
) ENGINE = InnoDB DEFAULT CHARSET = latin1;
