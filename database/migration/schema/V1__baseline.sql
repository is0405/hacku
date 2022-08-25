CREATE TABLE `user` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `name` varchar(32) NOT NULL,
  `mail` varchar(64) NOT NULL,
  `age` int NOT NULL DEFAULT 18,
  `faculty` int  NOT NULL,
  `password` varchar(255) NOT NULL,
  `gender` int NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE `sub_user` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `name` varchar(32) NOT NULL,
  `mail` varchar(64) NOT NULL,
  `age` int NOT NULL DEFAULT 18,
  `faculty` int  NOT NULL,
  `password` varchar(255) NOT NULL,
  `gender` int NOT NULL,
  `code` char(6) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;


CREATE TABLE `recruitment` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `conditions` varchar(255) NOT NULL,
  `contents` varchar(255) NOT NULL,
  `max_participation` int  NOT NULL DEFAULT 1,
  `reward` varchar(255) NOT NULL,
  `title` varchar(64) NOT NULL,
  `period` varchar(64) NOT NULL,
  `gender` int NOT NULL DEFAULT 2,
  `min_age` int NOT NULL DEFAULT 18,
  `max_age` int NOT NULL DEFAULT 60,
  `submit_id` int NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE `participation` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `recruitment_id` int NOT NULL,
  `user_id` int NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
