# テーブルを作成するSQL文
CREATE TABLE `user`
(
    `user_id`    varchar(128) NOT NULL,
    `auth_token` varchar(128) DEFAULT NULL,
    `name`       varchar(64)  DEFAULT NULL,
    PRIMARY KEY (`user_id`)
);

# テストデータの登録を行うSQL文
INSERT INTO `ca_tech_dojo`.`user` (`user_id`, `auth_token`, `name`) VALUES ('user_1', 'srnkrakm', 'sanae');
INSERT INTO `ca_tech_dojo`.`user` (`user_id`, `auth_token`, `name`) VALUES ('user_2', 'gjoaurajk', 'kaneda');
INSERT INTO `ca_tech_dojo`.`user` (`user_id`, `auth_token`, `name`) VALUES ('user_3', 'gaoenrklk', 'takuma');
INSERT INTO `ca_tech_dojo`.`user` (`user_id`, `auth_token`, `name`) VALUES ('user_4', 'oaernom', 'sasa');
INSERT INTO `ca_tech_dojo`.`user` (`user_id`, `auth_token`, `name`) VALUES ('user_5', 'rnaomrt', 'kishida');
