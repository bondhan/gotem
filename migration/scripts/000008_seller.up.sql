CREATE TABLE `m_seller` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  `name` varchar(255) NOT NULL,
  `address` varchar(255) DEFAULT NULL,
  `zipcode_id` int(10) unsigned NULL,
  PRIMARY KEY (`id`),
  KEY `m_seller_deleted_at` (`deleted_at`),
  CONSTRAINT `fk_m_seller_zip_code` FOREIGN KEY (`zipcode_id`) REFERENCES `m_zipcode` (`id`) ON DELETE RESTRICT ON UPDATE CASCADE
) ENGINE = InnoDB DEFAULT CHARSET = latin1;
