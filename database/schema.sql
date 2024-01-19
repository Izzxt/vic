-- Users
CREATE TABLE users (
  id int NOT NULL AUTO_INCREMENT,
  username varchar(255) NOT NULL,
  password text NOT NULL,
  auth_ticket varchar(255) NOT NULL DEFAULT '',
  email varchar(255) NOT NULL,
  rank_id int NOT NULL,
  account_created_date timestamp NOT NULL,
  last_online_date timestamp NOT NULL,
  is_online tinyint(1) NOT NULL DEFAULT '0',
  motto varchar(255) NOT NULL,
  look varchar(255) NOT NULL DEFAULT 'hr-115-42.hd-195-19.ch-3030-82.lg-275-1408.fa-1201.ca-1804-64',
  gender enum('M','F') NOT NULL DEFAULT 'M',
  ip_register varchar(255) NOT NULL,
  ip_current varchar(255) NOT NULL,
  home_room int NOT NULL DEFAULT '0',
  PRIMARY KEY (id)
);

CREATE TABLE rooms (
  id int NOT NULL AUTO_INCREMENT,
  owner_id int NOT NULL DEFAULT 0,
  name varchar(50) NOT NULL DEFAULT '',
  description varchar(512) NOT NULL DEFAULT '',
  model_id int NOT NULL DEFAULT 1,
  password varchar(20) NOT NULL DEFAULT '',
  state ENUM('open','locked','password','invisible') NOT NULL DEFAULT 'open',
  users int NOT NULL DEFAULT 0,
  max_users int NOT NULL DEFAULT 25,
  flat_category_id int NOT NULL DEFAULT 1,
  score int NOT NULL DEFAULT 0,
  floorpaper varchar(5) NOT NULL DEFAULT '0.0',
  wallpaper varchar(5) NOT NULL DEFAULT '0.0',
  landscape varchar(5) NOT NULL DEFAULT '0.0',
  wall_thickness int NOT NULL DEFAULT 0,
  wall_height int NOT NULL DEFAULT -1,
  floor_thickness int NOT NULL DEFAULT 0,
  tags varchar(500) NOT NULL DEFAULT '',
  is_public BOOLEAN NOT NULL DEFAULT FALSE,
  is_staff_picked BOOLEAN NOT NULL DEFAULT FALSE,
  allow_other_pets BOOLEAN NOT NULL DEFAULT FALSE,
  allow_other_pets_eat BOOLEAN NOT NULL DEFAULT FALSE,
  allow_walkthrough BOOLEAN NOT NULL DEFAULT TRUE,
  is_wall_hidden BOOLEAN NOT NULL DEFAULT FALSE,
  chat_mode int NOT NULL DEFAULT 0,
  chat_weight int NOT NULL DEFAULT 1,
  chat_scrolling_speed int NOT NULL DEFAULT 1,
  chat_hearing_distance int NOT NULL DEFAULT 50,
  chat_protection int NOT NULL DEFAULT 2,
  who_can_mute int NOT NULL DEFAULT 0,
  who_can_kick int NOT NULL DEFAULT 0,
  who_can_ban int NOT NULL DEFAULT 0,
  roller_speed int NOT NULL DEFAULT 4,
  is_promoted BOOLEAN NOT NULL DEFAULT FALSE,
  trade_mode int NOT NULL DEFAULT 2,
  move_diagonal BOOLEAN NOT NULL DEFAULT TRUE,
  is_wired_hidden BOOLEAN NOT NULL DEFAULT FALSE,
  is_forsale BOOLEAN NOT NULL DEFAULT FALSE,
  PRIMARY KEY (id),
  FOREIGN KEY (owner_id) REFERENCES users (id),
  FOREIGN KEY (model_id) REFERENCES room_models (id),
  FOREIGN KEY (flat_category_id) REFERENCES navigator_flat_cats (id)
);

CREATE TABLE room_models (
    id INT AUTO_INCREMENT PRIMARY KEY NOT NULL UNIQUE,
    name VARCHAR(255) NOT NULL,
    heightmap VARCHAR(10000) DEFAULT '' NOT NULL,
    is_club BOOLEAN DEFAULT false NOT NULL,
    is_custom BOOLEAN DEFAULT false NOT NULL,
    x INT DEFAULT 0 NOT NULL,
    y INT DEFAULT 0 NOT NULL,
    dir INT DEFAULT 0 NOT NULL
);

CREATE TABLE navigator_flat_cats  (
  id int NOT NULL AUTO_INCREMENT,
  min_rank int NOT NULL DEFAULT 0,
  caption_save varchar(32) NOT NULL DEFAULT 'caption_save',
  caption varchar(100) NOT NULL,
  allow_trade BOOLEAN NOT NULL DEFAULT true,
  max_users int NOT NULL DEFAULT 100,
  is_public BOOLEAN NOT NULL DEFAULT false,
  order_num int NOT NULL DEFAULT 0,
  PRIMARY KEY (id) USING BTREE
);

CREATE TABLE navigator_public_cats  (
  id int NOT NULL AUTO_INCREMENT,
  name varchar(32) NOT NULL DEFAULT 'Staff Picks',
  has_image BOOLEAN NOT NULL DEFAULT '0',
  visible BOOLEAN NOT NULL DEFAULT '1',
  order_num int NOT NULL DEFAULT 0,
  PRIMARY KEY (id) USING BTREE
);