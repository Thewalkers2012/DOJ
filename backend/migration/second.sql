CREATE TABLE `contexts` (
  `id` bigint PRIMARY KEY AUTO_INCREMENT,
  `title` varchar(255) NOT NULL,
  `start_time` timestamp NOT NULL,
  `end_time` timestamp NOT NULL,
  `defunct` char(1) NOT NULL,
  `author` varchar(255) NOT NULL
);
CREATE TABLE `context_problems` (
  `id` bigint PRIMARY KEY AUTO_INCREMENT,
  `problem_id` bigint NOT NULL,
  `context_id` bigint NOT NULL,
  `title` varchar(255) NOT NULL
);
CREATE TABLE `context_submissions` (
  `id` bigint PRIMARY KEY AUTO_INCREMENT,
  `problem_id` bigint NOT NULL,
  `context_id` bigint NOT NULL,
  `result` bigint NOT NULL
);
ALTER TABLE `context_problems`
ADD FOREIGN KEY (`problem_id`) REFERENCES `problems` (`id`);
ALTER TABLE `context_problems`
ADD FOREIGN KEY (`context_id`) REFERENCES `contexts` (`id`);
CREATE INDEX `context_problem_index_7` ON `context_problems` (`problem_id`);
CREATE INDEX `context_problem_index_8` ON `context_problems` (`context_id`);
ALTER TABLE `context_submissions`
ADD FOREIGN KEY (`problem_id`) REFERENCES `problems` (`id`);
ALTER TABLE `context_submissions`
ADD FOREIGN KEY (`context_id`) REFERENCES `contexts` (`id`);
CREATE INDEX `context_problem_index_7` ON `context_submissions` (`problem_id`);
CREATE INDEX `context_problem_index_8` ON `context_submissions` (`context_id`);