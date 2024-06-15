DROP TABLE IF EXISTS `rooms`;
CREATE TABLE `rooms` (
  `room_id` VARCHAR(36) NOT NULL,
  `room_name` varchar(30) NOT NULL,
  `is_public` BOOLEAN NOT NULL,
  PRIMARY KEY (`room_id`)
);

CREATE TABLE `words` (
  `room_id` VARCHAR(36) NOT NULL,
  `word_id` INTEGER NOT NULL AUTO_INCREMENT,
  `word` varchar(30) NOT NULL,
  `reading` varchar(30) NOT NULL,
  `basic_score` INTEGER NOT NULL,
  PRIMARY KEY (`room_id`, `word_id`)
)
