CREATE DATABASE anime_data;

USE anime_data;

show tables ;
drop table anime_info;

select * from anime_info;

CREATE TABLE anime_info (
                            anime_id INT PRIMARY KEY,
                            name VARCHAR(255),
                            english_name VARCHAR(255),
                            other_name VARCHAR(255),
                            score FLOAT,
                            genres TEXT,
                            synopsis TEXT,
                            type VARCHAR(50),
                            episodes VARCHAR(50),
                            aired VARCHAR(100),
                            premiered VARCHAR(100),
                            status VARCHAR(50),
                            producers TEXT,
                            licensors TEXT,
                            studios TEXT,
                            source VARCHAR(100),
                            duration VARCHAR(50),
                            rating VARCHAR(50),
                            `rank` FLOAT,
                            popularity INT,
                            favorites INT,
                            scored_by VARCHAR(50),
                            members INT,
                            image_url TEXT
);

CREATE TABLE users (
                       user_id INT AUTO_INCREMENT PRIMARY KEY,
                       username VARCHAR(100) UNIQUE NOT NULL,
                       password VARCHAR(255) NOT NULL
);

desc users;

CREATE TABLE user_favorites (
                                favorite_id INT AUTO_INCREMENT PRIMARY KEY,
                                user_id INT NOT NULL,
                                anime_id INT NOT NULL,
                                FOREIGN KEY (user_id) REFERENCES users(user_id) ON DELETE CASCADE,
                                FOREIGN KEY (anime_id) REFERENCES anime_info(anime_id) ON DELETE CASCADE
);

select * from users;