CREATE DATABASE IS NOT EXISTS users;
USE users;
CREATE TABLE IF NOT EXISTS user (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    uuid VARCHAR(64) NOT NULL,
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
delimiter $$ CREATE PROCEDURE insertLoop() BEGIN
DECLARE i INT DEFAULT 1;
WHILE (i <= 100000) DO
INSERT INTO user
VALUES (
        NULL,
        "martin",
        "f8d1c837-719f-42a9-9a37-0e2ed7c0e458",
        "5'9",
        "9",
        15,
        "Student and Developer",
        100,
        "horse",
        "red",
        "apple",
        "555-555-5555"
    );
SET i = i + 1;
END WHILE;
END $$

/* 
    Notes: use transaction for insertion then commit once it is done for more optimal performance (MySQL)
*/ 