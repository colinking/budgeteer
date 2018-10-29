ALTER DATABASE CHARACTER SET utf8 COLLATE utf8_general_ci;

-- users table

DROP TABLE IF EXISTS `users`;
CREATE TABLE `users` (
  `auth_id` varbinary(64) NOT NULL,
  `first_name` varbinary(255),
  `last_name` varbinary(255),
  `email` varbinary(255),
  `picture` varbinary(1024),
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`auth_id`),
  UNIQUE KEY `users_auth_id_index` (`auth_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
