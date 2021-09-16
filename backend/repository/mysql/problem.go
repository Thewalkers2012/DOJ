package mysql

import (
	"github.com/Thewalkers2012/DOJ/model"
)

// type Problem struct {
// 	ID              int64  `json:"problem_id"`
// 	Name            string `json:"problem_name"`
// 	Description     string `json:"description"`
// 	TestCase        string `json:"test_case"`
// 	Author          string `json:"author"`
// 	TimeLimit       int64  `json:"time_limit"`
// 	MemoryLimit     int64  `json:"memory_limit"`
// 	DifficultyLevel string `json:"difficulty_level"`
// }

func CreateProblem(req *model.CreateProblemRequest) *model.Problem {
	problem := &model.Problem{
		Name:            req.Name,
		Description:     req.Description,
		TestCase:        req.TestCase,
		Author:          req.Author,
		TimeLimit:       req.TimeLimit,
		MemoryLimit:     req.MemoryLimit,
		DifficultyLevel: req.DifficultyLevel,
	}

	DB.Select("Name", "Description", "TestCase", "Author", "TimeLimit", "MemoryLimit", "DifficultyLevel").Create(problem)

	return problem
}
