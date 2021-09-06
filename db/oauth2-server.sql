/*
 Navicat Premium Data Transfer

 Source Server         : 10.100.0.88
 Source Server Type    : MySQL
 Source Server Version : 50726
 Source Host           : 10.100.0.88:30113
 Source Schema         : oauth2-server

 Target Server Type    : MySQL
 Target Server Version : 50726
 File Encoding         : 65001

 Date: 06/09/2021 13:50:09
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for oauth_access_token
-- ----------------------------
DROP TABLE IF EXISTS `oauth_access_token`;
CREATE TABLE `oauth_access_token` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `token` varchar(64) NOT NULL COMMENT 'token',
  `client_id` varchar(255) NOT NULL COMMENT '客户端Id',
  `user_id` bigint(20) NOT NULL COMMENT '用户ID',
  `expired_at` bigint(20) NOT NULL COMMENT '过期时间戳',
  `scope` varchar(255) NOT NULL COMMENT 'scope',
  `del` tinyint(4) NOT NULL DEFAULT '0' COMMENT '是否删除  0：否  1：是',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY `token` (`token`) USING HASH
) ENGINE=InnoDB AUTO_INCREMENT=33 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of oauth_access_token
-- ----------------------------
BEGIN;
INSERT INTO `oauth_access_token` VALUES (10, 'b5787504c0de4ca7bbde9435a4cca74a', 'fOTQB8es4GBQwcsy', 1, 1627203295127, 'all', 0, '2021-07-25 08:54:48', '2021-07-25 12:27:37');
INSERT INTO `oauth_access_token` VALUES (11, '4ba800985f5a46758b7e1800f6e84f02', 'fOTQB8es4GBQwcsy', 1, 1627210798137, 'all', 1, '2021-07-25 08:59:58', '2021-07-25 12:27:40');
INSERT INTO `oauth_access_token` VALUES (12, '0e13aa4ce5ee47aa8a78cfa52782715c', 'fOTQB8es4GBQwcsy', 1, 1627211052599, 'all', 0, '2021-07-25 09:04:12', '2021-07-25 12:27:44');
INSERT INTO `oauth_access_token` VALUES (15, '774c6622c1d142d2bfccb572802160d6', 'fOTQB8es4GBQwcsy', 1, 1627224892789, 'all', 1, '2021-07-25 12:54:57', '2021-07-25 12:56:51');
INSERT INTO `oauth_access_token` VALUES (16, 'a55119e2cc40404dbfe2f2eb73315ad0', 'fOTQB8es4GBQwcsy', 1, 1627225011314, 'all', 1, '2021-07-25 12:56:51', '2021-07-25 12:58:29');
INSERT INTO `oauth_access_token` VALUES (17, '92e9dcbd5e8e499195c68bbcc0f364b0', 'fOTQB8es4GBQwcsy', 1, 1627232208194, 'all', 1, '2021-07-25 12:58:29', '2021-07-25 14:57:41');
INSERT INTO `oauth_access_token` VALUES (18, 'e407ff5a61e14cea9ebbc4bc1ed91ad6', 'fOTQB8es4GBQwcsy', 1, 1627232830344, 'all', 0, '2021-07-25 14:57:54', '2021-07-25 15:07:10');
INSERT INTO `oauth_access_token` VALUES (19, '5af6e4f6fd674f6caa225510c1825e2f', 'fOTQB8es4GBQwcsy', 1, 1627281399208, 'all', 0, '2021-07-26 04:36:39', '2021-07-26 04:36:39');
INSERT INTO `oauth_access_token` VALUES (20, 'fd491517ad544922823eb1127178890c', 'fOTQB8es4GBQwcsy', 1, 1627482037792, 'all', 1, '2021-07-28 12:20:37', '2021-07-28 12:23:36');
INSERT INTO `oauth_access_token` VALUES (21, 'e54bf765a5ce415a97157e5bc2b3ba3d', 'fOTQB8es4GBQwcsy', 1, 1627482139553, 'all', 1, '2021-07-28 12:23:54', '2021-07-28 12:25:23');
INSERT INTO `oauth_access_token` VALUES (22, 'febf169fc2024c4488d7350c4d41fa2b', 'fOTQB8es4GBQwcsy', 1, 1627482321411, 'all', 1, '2021-07-28 12:25:23', '2021-07-28 12:26:53');
INSERT INTO `oauth_access_token` VALUES (23, '939cde9c5c994daaa08b743eaad8030b', 'fOTQB8es4GBQwcsy', 1, 1627482413330, 'all', 1, '2021-07-28 12:26:53', '2021-07-28 12:26:55');
INSERT INTO `oauth_access_token` VALUES (24, 'e927792c0c8e4a929d47d6c320f1ed0a', 'fOTQB8es4GBQwcsy', 1, 1627482473550, 'all', 1, '2021-07-28 12:26:55', '2021-07-28 12:30:27');
INSERT INTO `oauth_access_token` VALUES (25, 'c24c14fd8dfa4898b8e8322d0c06548a', 'fOTQB8es4GBQwcsy', 1, 1627482627927, 'all', 1, '2021-07-28 12:30:28', '2021-07-28 12:33:15');
INSERT INTO `oauth_access_token` VALUES (26, '0335f80b199c4e19b782392ca05c6fd1', 'fOTQB8es4GBQwcsy', 1, 1627482795630, 'all', 1, '2021-07-28 12:33:15', '2021-07-28 12:40:05');
INSERT INTO `oauth_access_token` VALUES (27, '6cfedbda872a4371ae123cf66436e979', 'fOTQB8es4GBQwcsy', 1, 1627483211918, 'all', 0, '2021-07-28 12:40:05', '2021-07-28 12:40:11');
INSERT INTO `oauth_access_token` VALUES (28, 'ae85c9bc5e2341099223e0f5e5db0526', 'fOTQB8es4GBQwcsy', 0, 1627484431905, 'all', 1, '2021-07-28 13:00:31', '2021-07-28 13:01:53');
INSERT INTO `oauth_access_token` VALUES (29, '5439c979100b4d0bb66e605720105fe9', 'fOTQB8es4GBQwcsy', 0, 1627484513087, 'all', 1, '2021-07-28 13:01:53', '2021-07-28 13:01:56');
INSERT INTO `oauth_access_token` VALUES (30, '96e64e66a471424aa0f5fcd54c4e0d1d', 'fOTQB8es4GBQwcsy', 0, 1627484516206, 'all', 1, '2021-07-28 13:01:56', '2021-07-28 13:01:57');
INSERT INTO `oauth_access_token` VALUES (31, '1d6b19f283754e348a05cf9ac4627a07', 'fOTQB8es4GBQwcsy', 0, 1627484517105, 'all', 1, '2021-07-28 13:01:57', '2021-07-28 13:01:58');
INSERT INTO `oauth_access_token` VALUES (32, '88272d80c43c4cbc8216dfb591bdd8f7', 'fOTQB8es4GBQwcsy', 0, 1627484517990, 'all', 0, '2021-07-28 13:01:58', '2021-07-28 13:01:58');
COMMIT;

-- ----------------------------
-- Table structure for oauth_client_detail
-- ----------------------------
DROP TABLE IF EXISTS `oauth_client_detail`;
CREATE TABLE `oauth_client_detail` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `client_id` varchar(255) NOT NULL COMMENT '客户端ID',
  `client_secret` varchar(255) NOT NULL COMMENT '客户端访问密匙\r\n\r\n',
  `scope` varchar(255) NOT NULL COMMENT '客户端申请的权限范围，可选值包括read,write,trust;若有多个权限范围用逗号(,)分隔',
  `authorized_grant_types` varchar(255) NOT NULL COMMENT '客户端支持的授权许可类型(grant_type)，可选值包括authorization_code,password,refresh_token,implicit,client_credentials,若支持多个授权许可类型用逗号(,)分隔',
  `web_server_redirect_uri` varchar(1024) NOT NULL COMMENT '客户端重定向URI，当grant_type为authorization_code或implicit时, 在Oauth的流程中会使用并检查与数据库内的redirect_uri是否一致',
  `access_token_validity` int(255) NOT NULL DEFAULT '7200' COMMENT '设定客户端的access_token的有效时间值(单位:秒)，若不设定值则使用默认的有效时间值：2 * 60 * 60， 2小时',
  `refresh_token_validity` int(255) NOT NULL DEFAULT '2592000' COMMENT '设定客户端的refresh_token的有效时间值(单位:秒)，若不设定值则使用默认的有效时间值：30 * 24 * 60 * 60 ，30天',
  `additional_information` json DEFAULT NULL COMMENT '预留字段',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `client_id` (`client_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of oauth_client_detail
-- ----------------------------
BEGIN;
INSERT INTO `oauth_client_detail` VALUES (7, 'fOTQB8es4GBQwcsy', 'd4f9b92d0c9a9395cf81d819afebd74a', 'all', 'authorization_code,implicit,refresh_token,password,client_credentials', 'https://www.baidu.com', 7200, 2592000, '{}', '2021-07-24 03:57:14', '2021-07-28 13:00:30');
COMMIT;

-- ----------------------------
-- Table structure for oauth_code
-- ----------------------------
DROP TABLE IF EXISTS `oauth_code`;
CREATE TABLE `oauth_code` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `code` varchar(10) NOT NULL COMMENT '授权码',
  `user_id` bigint(20) NOT NULL COMMENT '用户ID',
  `client_id` varchar(64) NOT NULL COMMENT '客户端id',
  `expired_at` bigint(20) NOT NULL COMMENT '过期时间戳',
  `scope` varchar(255) NOT NULL COMMENT 'scope',
  `del` tinyint(4) NOT NULL DEFAULT '0' COMMENT '是否删除 0：未删除 1：已删除',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `code` (`code`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=20 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of oauth_code
-- ----------------------------
BEGIN;
INSERT INTO `oauth_code` VALUES (6, 'IX7yR5', 1, 'fOTQB8es4GBQwcsy', 1627133599139, 'all', 1, '2021-07-24 13:28:19', '2021-07-25 08:03:22');
INSERT INTO `oauth_code` VALUES (7, 'uawiLL', 1, 'fOTQB8es4GBQwcsy', 1627134902194, 'all', 0, '2021-07-24 13:50:02', '2021-07-24 14:07:06');
INSERT INTO `oauth_code` VALUES (8, '4NViHp', 1, 'fOTQB8es4GBQwcsy', 1627201611031, 'all', 0, '2021-07-25 08:21:51', '2021-07-25 08:21:51');
INSERT INTO `oauth_code` VALUES (9, 'O180YM', 1, 'fOTQB8es4GBQwcsy', 1627201978017, 'all', 0, '2021-07-25 08:27:58', '2021-07-25 08:27:58');
INSERT INTO `oauth_code` VALUES (10, 'bI2EDX', 1, 'fOTQB8es4GBQwcsy', 1627202348251, 'all', 0, '2021-07-25 08:35:17', '2021-07-25 08:35:17');
INSERT INTO `oauth_code` VALUES (11, 'eK9J5j', 1, 'fOTQB8es4GBQwcsy', 1627203060904, 'all', 0, '2021-07-25 08:46:00', '2021-07-25 08:46:00');
INSERT INTO `oauth_code` VALUES (12, 'e7xs0T', 1, 'fOTQB8es4GBQwcsy', 1627203500057, 'all', 1, '2021-07-25 08:53:20', '2021-07-25 08:54:48');
INSERT INTO `oauth_code` VALUES (13, 'hPi2vu', 1, 'fOTQB8es4GBQwcsy', 1627203889401, 'all', 1, '2021-07-25 08:59:49', '2021-07-25 08:59:58');
INSERT INTO `oauth_code` VALUES (14, 'Lea1yQ', 1, 'fOTQB8es4GBQwcsy', 1627204145908, 'all', 1, '2021-07-25 09:04:05', '2021-07-25 09:04:13');
INSERT INTO `oauth_code` VALUES (15, 'Sdk9kP', 1, 'fOTQB8es4GBQwcsy', 1627218175418, 'all', 1, '2021-07-25 12:57:55', '2021-07-25 12:58:29');
INSERT INTO `oauth_code` VALUES (16, 'X6bjrU', 1, 'fOTQB8es4GBQwcsy', 1627274469729, 'all', 0, '2021-07-26 04:36:09', '2021-07-26 04:36:09');
INSERT INTO `oauth_code` VALUES (17, 'L7zzZJ', 1, 'fOTQB8es4GBQwcsy', 1627289456072, 'all', 0, '2021-07-26 08:45:56', '2021-07-26 08:45:56');
INSERT INTO `oauth_code` VALUES (18, 'eo4Ty7', 1, 'fOTQB8es4GBQwcsy', 1627475720238, 'all', 1, '2021-07-28 12:30:20', '2021-07-28 12:30:28');
INSERT INTO `oauth_code` VALUES (19, 'qto9w3', 1, 'fOTQB8es4GBQwcsy', 1627476288606, 'all', 1, '2021-07-28 12:39:48', '2021-07-28 12:40:05');
COMMIT;

-- ----------------------------
-- Table structure for oauth_refresh_token
-- ----------------------------
DROP TABLE IF EXISTS `oauth_refresh_token`;
CREATE TABLE `oauth_refresh_token` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `refresh_token` varchar(64) NOT NULL COMMENT 'refresh_token',
  `client_id` varchar(255) NOT NULL COMMENT '客户端Id',
  `user_id` bigint(20) NOT NULL COMMENT '用户ID',
  `expired_at` bigint(20) NOT NULL COMMENT '过期时间戳',
  `del` tinyint(4) NOT NULL DEFAULT '0' COMMENT '是否删除  0：否  1：是',
  `scope` varchar(255) NOT NULL COMMENT 'scope',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=12 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of oauth_refresh_token
-- ----------------------------
BEGIN;
INSERT INTO `oauth_refresh_token` VALUES (1, '1c4abeed2cb34e7a835266bc4036cf92', 'fOTQB8es4GBQwcsy', 1, 1627205879927, 1, 'all', '2021-07-25 08:54:48', '2021-07-25 12:29:31');
INSERT INTO `oauth_refresh_token` VALUES (2, 'b2bae71a923f4c31ab16dc3562fb49d2', 'fOTQB8es4GBQwcsy', 1, 1629795598137, 1, 'all', '2021-07-25 08:59:58', '2021-07-25 12:29:34');
INSERT INTO `oauth_refresh_token` VALUES (3, '32f596682ea44e1ca913326277b16a61', 'fOTQB8es4GBQwcsy', 1, 1629795852599, 1, 'all', '2021-07-25 09:04:12', '2021-07-25 12:54:57');
INSERT INTO `oauth_refresh_token` VALUES (4, '30f01b07c49e4feb926388d5fe7fb352', 'fOTQB8es4GBQwcsy', 1, 1629809909025, 1, 'all', '2021-07-25 12:58:29', '2021-07-28 12:20:37');
INSERT INTO `oauth_refresh_token` VALUES (5, 'c783a6b3eeda4f51b2ff4249b49cb956', 'fOTQB8es4GBQwcsy', 1, 1630066837792, 1, 'all', '2021-07-28 12:20:37', '2021-07-28 12:23:50');
INSERT INTO `oauth_refresh_token` VALUES (6, '706e5d0219344a3b83a5f7c9bfce1014', 'fOTQB8es4GBQwcsy', 1, 1630066956286, 1, 'all', '2021-07-28 12:23:59', '2021-07-28 12:25:23');
INSERT INTO `oauth_refresh_token` VALUES (7, '0afda625dc0047bfb7c3087d20d00499', 'fOTQB8es4GBQwcsy', 1, 1630067121411, 1, 'all', '2021-07-28 12:25:24', '2021-07-28 12:26:53');
INSERT INTO `oauth_refresh_token` VALUES (8, 'd68e1c13698142cea5af6236a05fb142', 'fOTQB8es4GBQwcsy', 1, 1630067213330, 1, 'all', '2021-07-28 12:26:53', '2021-07-28 12:26:55');
INSERT INTO `oauth_refresh_token` VALUES (9, 'af096bc081544535bb9d63a45bb57f13', 'fOTQB8es4GBQwcsy', 1, 1630067215235, 1, 'all', '2021-07-28 12:26:55', '2021-07-28 12:30:28');
INSERT INTO `oauth_refresh_token` VALUES (10, '20d9744aaf6741a084be5e3260804eaf', 'fOTQB8es4GBQwcsy', 1, 1630067427927, 1, 'all', '2021-07-28 12:30:28', '2021-07-28 12:33:15');
INSERT INTO `oauth_refresh_token` VALUES (11, 'c53f9feb9dc343dd8c5b539c46cd1211', 'fOTQB8es4GBQwcsy', 1, 1630068005530, 0, 'all', '2021-07-28 12:40:05', '2021-07-28 12:40:05');
COMMIT;

-- ----------------------------
-- Table structure for oauth_user
-- ----------------------------
DROP TABLE IF EXISTS `oauth_user`;
CREATE TABLE `oauth_user` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL COMMENT '用户名',
  `nick_name` varchar(255) DEFAULT NULL COMMENT '昵称',
  `nation` varchar(255) DEFAULT NULL COMMENT '国家',
  `province` varchar(255) DEFAULT NULL COMMENT '省份',
  `city` varchar(255) DEFAULT NULL COMMENT '城市',
  `password` varchar(255) NOT NULL COMMENT '密码，md5加密',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `name` (`name`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of oauth_user
-- ----------------------------
BEGIN;
INSERT INTO `oauth_user` VALUES (1, 'admin', '正是那朵玫瑰', '中国', '广东', '深圳', '21218cca77804d2ba1922c33e0151105', '2021-07-24 07:18:59', '2021-07-25 14:57:25');
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
