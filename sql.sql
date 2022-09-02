# ************************************************************
# Sequel Pro SQL dump
# Version 5446
#
# https://www.sequelpro.com/
# https://github.com/sequelpro/sequelpro
#
# Host: 192.168.0.121 (MySQL 8.0.28)
# Database: go-zero
# Generation Time: 2022-09-02 05:51:18 +0000
# ************************************************************


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
SET NAMES utf8mb4;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;


# Dump of table crm_roles
# ------------------------------------------------------------

DROP TABLE IF EXISTS `crm_roles`;

CREATE TABLE `crm_roles` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(255) DEFAULT NULL,
  `is_enable` tinyint DEFAULT NULL,
  `level` smallint DEFAULT NULL,
  `parent_id` int DEFAULT NULL,
  `create_at` timestamp NULL DEFAULT NULL,
  `update_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

LOCK TABLES `crm_roles` WRITE;
/*!40000 ALTER TABLE `crm_roles` DISABLE KEYS */;

INSERT INTO `crm_roles` (`id`, `name`, `is_enable`, `level`, `parent_id`, `create_at`, `update_at`)
VALUES
	(1,'角色一',1,1,0,'2022-09-01 16:45:42','2022-09-01 16:45:42'),
	(2,'角色一1',1,2,1,'2022-09-01 16:47:47','2022-09-01 16:47:47'),
	(3,'角色一2',1,2,1,'2022-09-01 16:47:55','2022-09-01 16:47:55'),
	(4,'角色二',1,1,0,'2022-09-01 16:48:25','2022-09-01 16:48:25'),
	(5,'角色二1',1,2,4,'2022-09-01 16:48:41','2022-09-01 16:48:41'),
	(6,'角色三',1,1,0,'2022-09-01 16:48:57','2022-09-01 16:48:57');

/*!40000 ALTER TABLE `crm_roles` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table crm_users
# ------------------------------------------------------------

DROP TABLE IF EXISTS `crm_users`;

CREATE TABLE `crm_users` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '姓名',
  `email` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '邮箱',
  `phone` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '积分',
  `level` smallint DEFAULT NULL COMMENT '用户等级',
  `role_ids` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '角色',
  `create_time` timestamp NULL DEFAULT NULL,
  `update_time` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;



# Dump of table users
# ------------------------------------------------------------

DROP TABLE IF EXISTS `users`;

CREATE TABLE `users` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL DEFAULT '' COMMENT '用户姓名',
  `gender` tinyint unsigned NOT NULL DEFAULT '0' COMMENT '用户性别',
  `mobile` varchar(255) NOT NULL DEFAULT '' COMMENT '用户电话',
  `password` varchar(255) NOT NULL DEFAULT '' COMMENT '用户密码',
  `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `update_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_mobile_unique` (`mobile`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

LOCK TABLES `users` WRITE;
/*!40000 ALTER TABLE `users` DISABLE KEYS */;

INSERT INTO `users` (`id`, `name`, `gender`, `mobile`, `password`, `create_time`, `update_time`)
VALUES
	(1,'测试-zo-zero',1,'18212121114','ef8364020560a703cfc7aebebcd0b62d1bfea7d4a841eb8964cfbcda2ba85dd5','2022-08-25 13:41:59','2022-08-29 17:07:34');

/*!40000 ALTER TABLE `users` ENABLE KEYS */;
UNLOCK TABLES;



/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;
/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
