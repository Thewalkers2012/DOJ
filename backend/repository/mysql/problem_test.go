package mysql

import (
	"testing"

	"github.com/Thewalkers2012/DOJ/model"
	"github.com/Thewalkers2012/DOJ/util/random"
	"github.com/stretchr/testify/assert"
)

func createRandomProblem() *model.CreateProblemRequest {
	problem := &model.CreateProblemRequest{
		Name:            random.RandomUserName(),
		Description:     random.RandomStringWithLetter(100),
		TestCase:        random.RandomStringWithNumber(20),
		Author:          random.RandomUserName(),
		TimeLimit:       random.RandomInt(1, 100),
		MemoryLimit:     random.RandomInt(1, 100),
		DifficultyLevel: random.RandomDiffcultyLevel(),
	}

	return problem
}

func TestCreateProblem(t *testing.T) {
	problem := createRandomProblem()

	p := CreateProblem(problem)

	assert.NotZero(t, p.ID)
	assert.Equal(t, problem.Author, p.Author)
	assert.Equal(t, problem.Description, p.Description)
	assert.Equal(t, problem.DifficultyLevel, p.DifficultyLevel)
	assert.Equal(t, problem.MemoryLimit, p.MemoryLimit)
	assert.Equal(t, problem.TimeLimit, p.TimeLimit)
	assert.Equal(t, problem.Name, p.Name)
	assert.Equal(t, problem.TestCase, p.TestCase)
}
