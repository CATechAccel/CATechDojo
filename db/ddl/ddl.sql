# ユーザーテーブル
CREATE TABLE `users`
(
    `user_id`    varchar(128) NOT NULL,
    `auth_token` varchar(128) DEFAULT NULL,
    `name`       varchar(64)  DEFAULT NULL,
    PRIMARY KEY (`user_id`)
);

# キャラクターテーブル
CREATE TABLE `ca_tech_dojo`.`characters`
(
    `id`   VARCHAR(128) NOT NULL,
    `name` VARCHAR(64)  NULL,
    `power` INT NOT NULL DEFAULT 0,
    PRIMARY KEY (`id`)
);

# ユーザーキャラクターテーブル
CREATE TABLE `ca_tech_dojo`.`user_characters`
(
    `user_character_id` VARCHAR(128) NOT NULL,
    `user_id`           VARCHAR(128) NOT NULL,
    `character_id`      VARCHAR(128) NOT NULL,
    PRIMARY KEY (`user_character_id`),
    INDEX `user_id_idx` (`user_id` ASC),
    INDEX `character_id_idx` (`character_id` ASC),
    CONSTRAINT `user_id`
        FOREIGN KEY (`user_id`)
            REFERENCES `ca_tech_dojo`.`users` (`user_id`)
            ON DELETE CASCADE
            ON UPDATE CASCADE,
    CONSTRAINT `character_id`
        FOREIGN KEY (`character_id`)
            REFERENCES `ca_tech_dojo`.`characters` (`id`)
            ON DELETE CASCADE
            ON UPDATE CASCADE
);
