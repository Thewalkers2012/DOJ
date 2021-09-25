package mysql

import (
	"github.com/Thewalkers2012/DOJ/model"
)

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

func GetProblemByID(pid int64) (*model.Problem, error) {
	problem := new(model.Problem)
	err := DB.Where("id", pid).First(problem).Error

	return problem, err
}

func GetProblemList(offset, limit int) ([]*model.Problem, error) {
	problems := []*model.Problem{}
	err := DB.Offset(offset).Limit(limit).Find(&problems).Error
	return problems, err
}

func GetProblemSize() int64 {
	var count int64
	DB.Model(&model.Problem{}).Count(&count)
	return count
}
