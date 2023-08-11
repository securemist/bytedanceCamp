/*
 Navicat Premium Data Transfer

 Source Server         : mysql
 Source Server Type    : MySQL
 Source Server Version : 80016 (8.0.16)
 Source Host           : localhost:3306
 Source Schema         : camp

 Target Server Type    : MySQL
 Target Server Version : 80016 (8.0.16)
 File Encoding         : 65001

 Date: 11/08/2023 18:13:10
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for comment
-- ----------------------------
DROP TABLE IF EXISTS `comment`;
CREATE TABLE `comment` (
  `user_id` bigint(20) NOT NULL COMMENT '评论者id',
  `video_id` bigint(20) NOT NULL COMMENT '视频id',
  `content` varchar(200) NOT NULL COMMENT '评论内容',
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_user_video` (`user_id`,`video_id`),
  KEY `idx_comment_deleted_at` (`deleted_at`),
  KEY `fk_comment_video` (`video_id`),
  CONSTRAINT `fk_comment_user` FOREIGN KEY (`user_id`) REFERENCES `user` (`uuid`),
  CONSTRAINT `fk_comment_video` FOREIGN KEY (`video_id`) REFERENCES `video` (`uuid`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Table structure for favorite
-- ----------------------------
DROP TABLE IF EXISTS `favorite`;
CREATE TABLE `favorite` (
  `user_id` bigint(20) NOT NULL COMMENT '评论者id',
  `video_id` bigint(20) NOT NULL COMMENT '视频id',
  `is_favorite` tinyint(1) NOT NULL COMMENT '用户是否点赞视频',
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_user_video` (`user_id`,`video_id`),
  KEY `idx_favorite_deleted_at` (`deleted_at`),
  KEY `fk_favorite_video` (`video_id`),
  CONSTRAINT `fk_favorite_user` FOREIGN KEY (`user_id`) REFERENCES `user` (`uuid`),
  CONSTRAINT `fk_favorite_video` FOREIGN KEY (`video_id`) REFERENCES `video` (`uuid`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Table structure for friend
-- ----------------------------
DROP TABLE IF EXISTS `friend`;
CREATE TABLE `friend` (
  `user_id` bigint(20) NOT NULL COMMENT '用户id',
  `friend_id` bigint(20) NOT NULL COMMENT '好友id',
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_friend_deleted_at` (`deleted_at`),
  KEY `idx_user_friend` (`user_id`,`friend_id`),
  KEY `fk_friend_friend` (`friend_id`),
  CONSTRAINT `fk_friend_friend` FOREIGN KEY (`friend_id`) REFERENCES `user` (`uuid`),
  CONSTRAINT `fk_friend_user` FOREIGN KEY (`user_id`) REFERENCES `user` (`uuid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Table structure for relation
-- ----------------------------
DROP TABLE IF EXISTS `relation`;
CREATE TABLE `relation` (
  `user_id` bigint(20) NOT NULL COMMENT '关注者id',
  `to_user_id` bigint(20) NOT NULL COMMENT '被关注者id',
  `is_relation` tinyint(1) NOT NULL COMMENT '用户是否关注',
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_user_to_user` (`user_id`,`to_user_id`),
  KEY `idx_relation_deleted_at` (`deleted_at`),
  KEY `fk_relation_to_user` (`to_user_id`),
  CONSTRAINT `fk_relation_to_user` FOREIGN KEY (`to_user_id`) REFERENCES `user` (`uuid`),
  CONSTRAINT `fk_relation_user` FOREIGN KEY (`user_id`) REFERENCES `user` (`uuid`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` (
  `uuid` bigint(20) NOT NULL COMMENT '用户id',
  `username` varchar(32) NOT NULL COMMENT '用户名',
  `password` varchar(100) NOT NULL COMMENT '登录密码',
  `jwt_token` varchar(250) NOT NULL COMMENT '用户鉴权token',
  `avatar` varchar(200) DEFAULT NULL COMMENT '用户头像',
  `followers` bigint(20) DEFAULT '0' COMMENT '粉丝数',
  `followings` bigint(20) DEFAULT '0' COMMENT '关注数',
  `signature` varchar(200) DEFAULT NULL COMMENT '个人简介',
  `total_favorite` bigint(20) DEFAULT '0' COMMENT '获赞数量',
  `work_count` bigint(20) DEFAULT '0' COMMENT '作品数量',
  `favorite_count` bigint(20) DEFAULT '0' COMMENT '喜欢数量',
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uuid` (`uuid`),
  UNIQUE KEY `username` (`username`),
  KEY `idx_user_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Table structure for video
-- ----------------------------
DROP TABLE IF EXISTS `video`;
CREATE TABLE `video` (
  `uuid` bigint(20) NOT NULL COMMENT '视频id',
  `author_id` bigint(20) NOT NULL COMMENT '视频作者id',
  `favorite_count` bigint(20) DEFAULT '0' COMMENT '视频的点赞总数',
  `comment_count` bigint(20) DEFAULT '0' COMMENT '视频的评论总数',
  `title` varchar(200) NOT NULL COMMENT '视频标题',
  `play_url` varchar(200) NOT NULL COMMENT '视频播放地址',
  `cover_url` varchar(200) NOT NULL COMMENT '视频封面地址',
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uuid` (`uuid`),
  KEY `idx_video_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

SET FOREIGN_KEY_CHECKS = 1;
