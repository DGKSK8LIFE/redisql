CREATE TABLE IF NOT EXISTS users (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    uuid VARCHAR(32) NOT NULL,
    height VARCHAR(5) NOT NULL,
    shoesize TINYINT NOT NULL,
    age TINYINT NOT NULL,
    bio TEXT NOT NULL,
    friends_count TINYINT NOT NULL,
    favorite_animal VARCHAR(20) NOT NULL,
    favorite_color VARCHAR(10) NOT NULL,
    favorite_food VARCHAR(20) NOT NULL,
    mobile_phone VARCHAR(50) NOT NULL
);