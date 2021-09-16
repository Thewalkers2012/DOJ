package model

type Problem struct {
	ID              int64  `json:"problem_id"`
	Name            string `json:"problem_name"`
	Description     string `json:"description"`
	TestCase        string `json:"test_case"`
	Author          string `json:"author"`
	TimeLimit       int64  `json:"time_limit"`
	MemoryLimit     int64  `json:"memory_limit"`
	DifficultyLevel string `json:"difficulty_level"`
}
