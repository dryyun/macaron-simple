
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE IF NOT EXISTS `user` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(255)  NOT NULL,
  `phone` varchar(11)  NOT NULL,
  `password` varchar(128)  NOT NULL,
  `avatar` varchar(256) DEFAULT NULL,
  `created` int(11) DEFAULT '0',
  `updated` int(11) DEFAULT '0',
  PRIMARY KEY (`id`),
  UNIQUE KEY `phone` (`phone`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ;

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE IF EXISTS `user`  ;
