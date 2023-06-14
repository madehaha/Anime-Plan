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
    PRIMARY KEY (`uid`),
    UNIQUE INDEX `email` (`email`)
) CHARSET utf8mb4 COLLATE utf8mb4_bin AUTO_INCREMENT 3;
-- Create "subjects" table
CREATE TABLE `subjects`
(
    `uid`          int unsigned NOT NULL AUTO_INCREMENT,
    `image`        varchar(255) NOT NULL DEFAULT "https://lain.bgm.tv/pic/user/l/icon.jpg",
    `summary`      varchar(300) NOT NULL,
    `name`         varchar(255) NOT NULL,
    `date`         varchar(255) NOT NULL,
    `name_cn`      varchar(255) NOT NULL,
    `on_hold`      int unsigned NOT NULL,
    `wish`         int unsigned NOT NULL,
    `doing`        int unsigned NOT NULL,
    `subject_type` tinyint unsigned NOT NULL DEFAULT 0,
    `collect`      int unsigned NOT NULL DEFAULT 0,
    PRIMARY KEY (`uid`)
) CHARSET utf8mb4 COLLATE utf8mb4_bin;
