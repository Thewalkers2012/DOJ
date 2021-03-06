package model

type Problem struct {
	ID              int64  `json:"problem_id"`
	TimeLimit       int    `json:"time_limit"`
	MemoryLimit     int    `json:"memory_limit"`
	Name            string `json:"problem_name"`
	Description     string `json:"description"`
	TestCase        string `json:"test_case"`
	Author          string `json:"author"`
	DifficultyLevel string `json:"difficulty_level"`
}
