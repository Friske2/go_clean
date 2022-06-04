CREATE DATABASE  IF NOT EXISTS `labs`;
USE `labs`;

DROP TABLE IF EXISTS `todolist`;

create table `todolist` (
    `id` int(11) NOT NULL AUTO_INCREMENT,
    `name` varchar(45) COLLATE utf8_unicode_ci NOT NULL,
    `description` longtext COLLATE utf8_unicode_ci NOT NULL,
    `updated_at` datetime DEFAULT NULL,
    `created_at` datetime DEFAULT NULL,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;
