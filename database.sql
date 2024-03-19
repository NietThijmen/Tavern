CREATE TABLE IF NOT EXISTS`tavern_tokens`(
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `vault_id` int(11) NOT NULL,
  `token` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
