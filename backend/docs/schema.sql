-- Create "members" table
CREATE TABLE `members`
(
    `uid`           int unsigned NOT NULL AUTO_INCREMENT,
    `username`      varchar(30)  NOT NULL DEFAULT "",
    `email`         varchar(50)  NOT NULL,
    `password`      varchar(255) NOT NULL,
    `nickname`      varchar(30)  NOT NULL,
    `avatar`        varchar(255) NOT NULL DEFAULT "https://lain.bgm.tv/pic/user/l/icon.jpg",
    `gid`           tinyint unsigned NOT NULL DEFAULT 0,
    `register_time` varchar(255) NOT NULL,
    PRIMARY KEY (`uid`)
) CHARSET utf8mb4 COLLATE utf8mb4_bin AUTO_INCREMENT 2;
