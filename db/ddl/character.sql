CREATE TABLE `ca_tech_dojo`.`character`
(
    `id`   VARCHAR(128) NOT NULL,
    `name` VARCHAR(64)  NULL,
    `power` INT NOT NULL DEFAULT 0,
    PRIMARY KEY (`id`)
);
