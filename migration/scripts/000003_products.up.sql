CREATE TABLE `m_products` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  `name` varchar(255) NOT NULL,
  `variants` int(255) unsigned NOT NULL,
  `price` DECIMAL(14, 2) DEFAULT NULL,
  `desc` text DEFAULT NULL,
  `quantity` int(10) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `m_products_deleted_at` (`deleted_at`)
) ENGINE = InnoDB DEFAULT CHARSET = latin1;
