-- Users
CREATE TABLE `users` (
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
)