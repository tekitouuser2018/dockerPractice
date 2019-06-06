DROP DATABASE IF EXISTS test_db;
CREATE DATABASE test_db;
use test_db;
CREATE USER subuser@localhost IDENTIFIED BY 'password2';
GRANT ALL PRIVILEGES ON test_db.* TO 'subuser'@'localhost' WITH GRANT OPTION;
FLUSH PRIVILEGES;

CREATE TABLE `users` (
    `id` INT(10) NOT NULL PRIMARY KEY AUTO_INCREMENT,
    `name` varchar(20) not null,
    created_time datetime not null default current_timestamp,
    updated_time datetime not null default current_timestamp on update current_timestamp
);


INSERT INTO users(name) VALUES('user1'),('second');