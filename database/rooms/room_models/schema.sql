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