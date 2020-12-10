CREATE TABLE `ca_tech_dojo`.`user_character`
(
    `user_character_id` VARCHAR(128) NOT NULL,
    `user_id`           VARCHAR(128) NOT NULL,
    `character_id`      VARCHAR(128) NOT NULL,
    PRIMARY KEY (`user_character_id`),
    INDEX `user_id_idx` (`user_id` ASC),
    INDEX `character_id_idx` (`character_id` ASC),
    CONSTRAINT `user_id`
        FOREIGN KEY (`user_id`)
            REFERENCES `ca_tech_dojo`.`user` (`user_id`)
            ON DELETE CASCADE
            ON UPDATE CASCADE,
    CONSTRAINT `character_id`
        FOREIGN KEY (`character_id`)
            REFERENCES `ca_tech_dojo`.`character` (`id`)
            ON DELETE CASCADE
            ON UPDATE CASCADE
);
