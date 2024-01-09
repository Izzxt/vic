CREATE TABLE `navigator_public_cats`  (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(32) NOT NULL DEFAULT 'Staff Picks',
  `has_image` BOOLEAN NOT NULL DEFAULT '0',
  `visible` BOOLEAN NOT NULL DEFAULT '1',
  `order_num` int NOT NULL DEFAULT 0,
  PRIMARY KEY (`id`) USING BTREE
);