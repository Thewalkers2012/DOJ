CREATE TABLE `users` (
  `id` bigint PRIMARY KEY AUTO_INCREMENT,
  `username` varchar(255) NOT NULL,
  `student_id` varchar(255) UNIQUE NOT NULL,
  `hashed_password` varchar(255) NOT NULL,
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
CREATE TABLE `categories` (
  `id` bigint PRIMARY KEY AUTO_INCREMENT,
  `user_id` bigint NOT NULL,
  `problem_id` bigint NOT NULL,
  `author` varchar(255) NOT NULL,
  `content` varchar(255) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT (now())
);
ALTER TABLE `submissions`
ADD FOREIGN KEY (`user_id`) REFERENCES `users` (`id`);
ALTER TABLE `submissions`
ADD FOREIGN KEY (`question_id`) REFERENCES `problems` (`id`);
ALTER TABLE `categories`
ADD FOREIGN KEY (`user_id`) REFERENCES `users` (`id`);
ALTER TABLE `categories`
ADD FOREIGN KEY (`problem_id`) REFERENCES `problems` (`id`);
CREATE INDEX `users_index_0` ON `users` (`student_id`);
CREATE INDEX `problems_index_1` ON `problems` (`id`);
CREATE INDEX `problems_index_2` ON `problems` (`name`);
CREATE INDEX `submissions_index_3` ON `submissions` (`user_id`);
CREATE INDEX `submissions_index_4` ON `submissions` (`question_id`);
CREATE INDEX `categorys_index_5` ON `categorys` (`user_id`);
CREATE INDEX `categorys_index_6` ON `categorys` (`question_id`);