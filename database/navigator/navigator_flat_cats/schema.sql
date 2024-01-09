CREATE TABLE `navigator_flat_cats`  (
  `id` int NOT NULL AUTO_INCREMENT,
  `min_rank` int NOT NULL DEFAULT 0,
  `caption_save` varchar(32) NOT NULL DEFAULT 'caption_save',
  `caption` varchar(100) NOT NULL,
  `allow_trade` BOOLEAN NOT NULL DEFAULT true,
  `max_users` int NOT NULL DEFAULT 100,
  `is_public` BOOLEAN NOT NULL DEFAULT false,
  `order_num` int NOT NULL DEFAULT 0,
  PRIMARY KEY (`id`) USING BTREE
);