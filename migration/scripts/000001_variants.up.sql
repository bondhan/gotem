CREATE TABLE `m_variants` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  `size` ENUM('S', 'M', 'L', 'XL', 'XXL', 'LL', 'LLL') NOT NULL,  
  PRIMARY KEY (`id`),
  KEY `m_products_deleted_at` (`deleted_at`)
) ENGINE = InnoDB DEFAULT CHARSET = latin1;
