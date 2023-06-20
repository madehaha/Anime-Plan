-- Create "members" table
CREATE TABLE `members`
(
    `uid`           int unsigned     NOT NULL AUTO_INCREMENT,
    `username`      varchar(30)      NOT NULL,
    `email`         varchar(50)      NOT NULL,
    `password`      varchar(255)     NOT NULL,
    `nickname`      varchar(30)      NOT NULL,
    `avatar`        varchar(255)     NOT NULL DEFAULT "https://lain.bgm.tv/pic/user/l/icon.jpg",
    `gid`           tinyint unsigned NOT NULL DEFAULT 0,
    `register_time` varchar(255)     NOT NULL,
    PRIMARY KEY (`uid`),
    UNIQUE INDEX `email` (`email`),
    UNIQUE INDEX `username` (`username`)
) CHARSET utf8mb4
  COLLATE utf8mb4_bin
  AUTO_INCREMENT 2;
-- Create "subjects" table
CREATE TABLE `subjects`
(
    `id`       int unsigned     NOT NULL AUTO_INCREMENT,
    `image`    varchar(255)     NOT NULL DEFAULT "https://lain.bgm.tv/pic/user/l/icon.jpg",
    `summary`  varchar(300)     NOT NULL DEFAULT "No summary.",
    `name`     varchar(255)     NOT NULL,
    `name_cn`  varchar(255)     NOT NULL,
    `episodes` tinyint unsigned NOT NULL,
    `wish`     int unsigned     NOT NULL DEFAULT 0,
    `doing`    int unsigned     NOT NULL DEFAULT 0,
    `watched`  int unsigned     NOT NULL DEFAULT 0,
    `on_hold`  int unsigned     NOT NULL DEFAULT 0,
    `dropped`  int unsigned     NOT NULL DEFAULT 0,
    PRIMARY KEY (`id`)
) CHARSET utf8mb4
  COLLATE utf8mb4_bin
  AUTO_INCREMENT 2;
-- Create "collections" table
CREATE TABLE `collections`
(
    `id`          int unsigned     NOT NULL AUTO_INCREMENT,
    `type`        tinyint unsigned NOT NULL,
    `has_comment` bool             NOT NULL DEFAULT 0,
    `comment`     varchar(100)     NOT NULL DEFAULT "",
    `score`       tinyint unsigned NOT NULL DEFAULT 0,
    `add_time`    varchar(255)     NOT NULL DEFAULT "2000-01-01",
    `ep_status`   tinyint unsigned NOT NULL DEFAULT 0,
    `member_id`   int unsigned     NULL,
    `subject_id`  int unsigned     NULL,
    PRIMARY KEY (`id`),
    INDEX `collections_members_collections` (`member_id`),
    INDEX `collections_subjects_collections` (`subject_id`),
    CONSTRAINT `collections_members_collections` FOREIGN KEY (`member_id`) REFERENCES `members` (`uid`) ON UPDATE NO ACTION ON DELETE SET NULL,
    CONSTRAINT `collections_subjects_collections` FOREIGN KEY (`subject_id`) REFERENCES `subjects` (`id`) ON UPDATE NO ACTION ON DELETE SET NULL
) CHARSET utf8mb4
  COLLATE utf8mb4_bin
  AUTO_INCREMENT 2;
-- Create "subject_fields" table
CREATE TABLE `subject_fields`
(
    `id`                    bigint           NOT NULL AUTO_INCREMENT,
    `rate_1`                int unsigned     NOT NULL DEFAULT 0,
    `rate_2`                int unsigned     NOT NULL DEFAULT 0,
    `rate_3`                int unsigned     NOT NULL DEFAULT 0,
    `rate_4`                int unsigned     NOT NULL DEFAULT 0,
    `rate_5`                int unsigned     NOT NULL DEFAULT 0,
    `rate_6`                int unsigned     NOT NULL DEFAULT 0,
    `rate_7`                int unsigned     NOT NULL DEFAULT 0,
    `rate_8`                int unsigned     NOT NULL DEFAULT 0,
    `rate_9`                int unsigned     NOT NULL DEFAULT 0,
    `rate_10`               int unsigned     NOT NULL DEFAULT 0,
    `average_score`         double           NOT NULL DEFAULT 0,
    `rank`                  int unsigned     NOT NULL DEFAULT 0,
    `year`                  int unsigned     NOT NULL,
    `month`                 tinyint unsigned NOT NULL,
    `date`                  tinyint unsigned NOT NULL,
    `weekday`               tinyint unsigned NOT NULL,
    `subject_subject_field` int unsigned     NULL,
    PRIMARY KEY (`id`),
    UNIQUE INDEX `subject_subject_field` (`subject_subject_field`),
    CONSTRAINT `subject_fields_subjects_subject_field` FOREIGN KEY (`subject_subject_field`) REFERENCES `subjects` (`id`) ON UPDATE NO ACTION ON DELETE SET NULL
) CHARSET utf8mb4
  COLLATE utf8mb4_bin
  AUTO_INCREMENT 2;
