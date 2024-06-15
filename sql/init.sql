DROP TABLE IF EXISTS `rooms`;
CREATE TABLE `rooms` (
  `room_id` VARCHAR(36) NOT NULL,
  `room_name` varchar(30) NOT NULL,
  `is_public` BOOLEAN NOT NULL,
  PRIMARY KEY (`room_id`)
);

DROP TABLE IF EXISTS `words`;
CREATE TABLE `words` (
  `word_id` INTEGER NOT NULL AUTO_INCREMENT,
  `room_id` VARCHAR(36) NOT NULL,
  `word` varchar(30) NOT NULL,
  `reading` varchar(30) NOT NULL,
  `basic_score` INTEGER NOT NULL,
  PRIMARY KEY (`word_id`),
  FOREIGN KEY (`room_id`) REFERENCES `rooms`(`room_id`)
)
