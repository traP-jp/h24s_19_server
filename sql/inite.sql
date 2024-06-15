DROP TABLE IF EXISTS `rooms`;
CREATE TABLE `rooms` (
  `room_id` VARCHAR(36) NOT NULL,
  `room_name` varchar(30) NOT NULL,
  `is_public` BOOLEAN NOT NULL,
  PRIMARY KEY (`room_id`)
);
