-- ----------------------------
-- MySQL Database Dump
-- Start Time: 2024-01-20 10:07:05
-- ----------------------------


-- ----------------------------
-- Table structure for bubbles_chat
-- ----------------------------
CREATE TABLE IF NOT EXISTS `bubbles_chat` (
  `id` int NOT NULL AUTO_INCREMENT,
  `key` varchar(255) NOT NULL,
  `is_ambassador` tinyint(1) NOT NULL,
  `is_overrideable` tinyint(1) NOT NULL DEFAULT '0',
  `trigger_talking_furniture` tinyint(1) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`),
  UNIQUE KEY `id` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=40 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;



-- ----------------------------
-- Table structure for navigator_flat_cats
-- ----------------------------
CREATE TABLE IF NOT EXISTS `navigator_flat_cats` (
  `id` int NOT NULL AUTO_INCREMENT,
  `min_rank` int NOT NULL DEFAULT '0',
  `caption_save` varchar(32) NOT NULL DEFAULT 'caption_save',
  `caption` varchar(100) NOT NULL,
  `allow_trade` tinyint(1) NOT NULL DEFAULT '1',
  `max_users` int NOT NULL DEFAULT '100',
  `is_public` tinyint(1) NOT NULL DEFAULT '0',
  `order_num` int NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=13 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;



-- ----------------------------
-- Table structure for navigator_public_cats
-- ----------------------------
CREATE TABLE IF NOT EXISTS `navigator_public_cats` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(32) NOT NULL DEFAULT 'Staff Picks',
  `has_image` tinyint(1) NOT NULL DEFAULT '0',
  `visible` tinyint(1) NOT NULL DEFAULT '1',
  `order_num` int NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;



-- ----------------------------
-- Table structure for room_models
-- ----------------------------
CREATE TABLE IF NOT EXISTS `room_models` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL,
  `heightmap` varchar(10000) NOT NULL DEFAULT '',
  `is_club` tinyint(1) NOT NULL DEFAULT '0',
  `is_custom` tinyint(1) NOT NULL DEFAULT '0',
  `x` int NOT NULL DEFAULT '0',
  `y` int NOT NULL DEFAULT '0',
  `dir` int NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`),
  UNIQUE KEY `id` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=38 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;



-- ----------------------------
-- Table structure for rooms
-- ----------------------------
CREATE TABLE IF NOT EXISTS `rooms` (
  `id` int NOT NULL AUTO_INCREMENT,
  `owner_id` int NOT NULL DEFAULT '0',
  `name` varchar(50) NOT NULL DEFAULT '',
  `description` varchar(512) NOT NULL DEFAULT '',
  `model_id` int NOT NULL DEFAULT '1',
  `password` varchar(20) NOT NULL DEFAULT '',
  `state` enum('open','locked','password','invisible') NOT NULL DEFAULT 'open',
  `users` int NOT NULL DEFAULT '0',
  `max_users` int NOT NULL DEFAULT '25',
  `flat_category_id` int NOT NULL DEFAULT '1',
  `score` int NOT NULL DEFAULT '0',
  `floorpaper` varchar(5) NOT NULL DEFAULT '0.0',
  `wallpaper` varchar(5) NOT NULL DEFAULT '0.0',
  `landscape` varchar(5) NOT NULL DEFAULT '0.0',
  `wall_thickness` int NOT NULL DEFAULT '0',
  `wall_height` int NOT NULL DEFAULT '-1',
  `floor_thickness` int NOT NULL DEFAULT '0',
  `tags` varchar(500) NOT NULL DEFAULT '',
  `is_public` tinyint(1) NOT NULL DEFAULT '0',
  `is_staff_picked` tinyint(1) NOT NULL DEFAULT '0',
  `allow_other_pets` tinyint(1) NOT NULL DEFAULT '0',
  `allow_other_pets_eat` tinyint(1) NOT NULL DEFAULT '0',
  `allow_walkthrough` tinyint(1) NOT NULL DEFAULT '1',
  `is_wall_hidden` tinyint(1) NOT NULL DEFAULT '0',
  `chat_mode` int NOT NULL DEFAULT '0',
  `chat_weight` int NOT NULL DEFAULT '1',
  `chat_scrolling_speed` int NOT NULL DEFAULT '1',
  `chat_hearing_distance` int NOT NULL DEFAULT '50',
  `chat_protection` int NOT NULL DEFAULT '2',
  `who_can_mute` int NOT NULL DEFAULT '0',
  `who_can_kick` int NOT NULL DEFAULT '0',
  `who_can_ban` int NOT NULL DEFAULT '0',
  `roller_speed` int NOT NULL DEFAULT '4',
  `is_promoted` tinyint(1) NOT NULL DEFAULT '0',
  `trade_mode` int NOT NULL DEFAULT '2',
  `move_diagonal` tinyint(1) NOT NULL DEFAULT '1',
  `is_wired_hidden` tinyint(1) NOT NULL DEFAULT '0',
  `is_forsale` tinyint(1) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`),
  KEY `owner_id` (`owner_id`),
  KEY `model_id` (`model_id`),
  KEY `flat_category_id` (`flat_category_id`),
  CONSTRAINT `rooms_ibfk_1` FOREIGN KEY (`owner_id`) REFERENCES `users` (`id`),
  CONSTRAINT `rooms_ibfk_2` FOREIGN KEY (`model_id`) REFERENCES `room_models` (`id`),
  CONSTRAINT `rooms_ibfk_3` FOREIGN KEY (`flat_category_id`) REFERENCES `navigator_flat_cats` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;



-- ----------------------------
-- Table structure for users
-- ----------------------------
CREATE TABLE IF NOT EXISTS `users` (
  `id` int NOT NULL AUTO_INCREMENT,
  `username` varchar(255) NOT NULL,
  `password` text NOT NULL,
  `auth_ticket` varchar(255) NOT NULL DEFAULT '',
  `email` varchar(255) NOT NULL,
  `rank_id` int NOT NULL,
  `account_created_date` timestamp NOT NULL,
  `last_online_date` timestamp NOT NULL,
  `is_online` tinyint(1) NOT NULL DEFAULT '0',
  `motto` varchar(255) NOT NULL,
  `look` varchar(255) NOT NULL DEFAULT 'hr-115-42.hd-195-19.ch-3030-82.lg-275-1408.fa-1201.ca-1804-64',
  `gender` enum('M','F') NOT NULL DEFAULT 'M',
  `ip_register` varchar(255) NOT NULL,
  `ip_current` varchar(255) NOT NULL,
  `home_room` int NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;



-- ----------------------------
-- Table structure for users_stats
-- ----------------------------
CREATE TABLE IF NOT EXISTS `users_stats` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `user_id` int NOT NULL,
  `bubble_chat_id` int DEFAULT NULL,
  UNIQUE KEY `id` (`id`),
  KEY `user_id` (`user_id`),
  KEY `bubble_chat_id` (`bubble_chat_id`),
  CONSTRAINT `users_stats_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`),
  CONSTRAINT `users_stats_ibfk_2` FOREIGN KEY (`bubble_chat_id`) REFERENCES `bubbles_chat` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;



-- ----------------------------
-- Dumped by mysqldump
-- Cost Time: 6.231291ms
-- ----------------------------
