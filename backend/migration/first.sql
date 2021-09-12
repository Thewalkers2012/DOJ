CREATE TABLE `users` (
  `id` bigint PRIMARY KEY AUTO_INCREMENT,
  `username` varchar(255) UNIQUE NOT NULL,
  `email` varchar(255) UNIQUE NOT NULL,
  `hashed_password` varchar(255) NOT NULL,
  `personal_signature` varchar(255),
  `password_changed_at` timestamp NOT NULL DEFAULT (now()),
  `created_at` timestamp NOT NULL DEFAULT (now())
);
CREATE TABLE `problems` (
  `id` bigint PRIMARY KEY AUTO_INCREMENT,
  `name` varchar(255) NOT NULL,
  `description` text NOT NULL,
  `testcases` varchar(255) NOT NULL,
  `author` varchar(255) NOT NULL DEFAULT "nobody",
  `timelimit` bigint NOT NULL,
  `memorylimit` bigint NOT NULL,
  `difficulty_level` varchar(255) NOT NULL
);
CREATE TABLE `submissions` (
  `id` bigint PRIMARY KEY AUTO_INCREMENT,
  `user_id` bigint NOT NULL,
  `question_id` bigint NOT NULL,
  `language` varchar(255) NOT NULL,
  `score` bigint NOT NULL,
  `result` varchar(255) NOT NULL
);
ALTER TABLE `submissions`
ADD FOREIGN KEY (`user_id`) REFERENCES `users` (`id`);
ALTER TABLE `submissions`
ADD FOREIGN KEY (`question_id`) REFERENCES `problems` (`id`);
CREATE INDEX `users_index_0` ON `users` (`username`);
CREATE INDEX `users_index_1` ON `users` (`email`);
CREATE INDEX `problems_index_2` ON `problems` (`id`);
CREATE INDEX `problems_index_3` ON `problems` (`name`);
CREATE INDEX `submissions_index_4` ON `submissions` (`user_id`);
CREATE INDEX `submissions_index_5` ON `submissions` (`question_id`);