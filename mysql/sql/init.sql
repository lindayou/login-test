/*
 Navicat Premium Data Transfer

 Source Server         : mysql
 Source Server Type    : MySQL
 Source Server Version : 50741
 Source Host           : 172.16.15.166:3306
 Source Schema         : test

 Target Server Type    : MySQL
 Target Server Version : 50741
 File Encoding         : 65001

 Date: 31/03/2023 16:11:03
*/

CREATE DATABASE test;

USE test;

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for users
-- ----------------------------
DROP TABLE IF EXISTS `users`;
CREATE TABLE `users`  (
						  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
						  `created_at` datetime(3) NULL DEFAULT NULL,
						  `updated_at` datetime(3) NULL DEFAULT NULL,
						  `deleted_at` datetime(3) NULL DEFAULT NULL,
						  `username` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL,
						  `password` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL,
						  PRIMARY KEY (`id`) USING BTREE,
						  UNIQUE INDEX `username`(`username`) USING BTREE,
						  INDEX `idx_users_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_bin ROW_FORMAT = Dynamic;

SET FOREIGN_KEY_CHECKS = 1;
