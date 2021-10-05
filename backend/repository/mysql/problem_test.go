package mysql

import (
	"testing"

	"github.com/Thewalkers2012/DOJ/model"
	"github.com/Thewalkers2012/DOJ/util/random"
	"github.com/stretchr/testify/assert"
)

func createRandomProblem(t *testing.T) *model.Problem {
	problem := &model.CreateProblemRequest{
		Name:            random.RandomUserName(),
		Description:     random.RandomStringWithLetter(100),
		TestCase:        random.RandomStringWithNumber(20),
		Author:          random.RandomUserName(),
		TimeLimit:       int(random.RandomInt(1, 100)),
		MemoryLimit:     int(random.RandomInt(1, 100)),
		DifficultyLevel: random.RandomDiffcultyLevel(),
	}

	p := CreateProblem(problem)
	assert.NotZero(t, p.ID)
	assert.Equal(t, problem.Author, p.Author)
	assert.Equal(t, problem.Description, p.Description)
	assert.Equal(t, problem.DifficultyLevel, p.DifficultyLevel)
	assert.Equal(t, problem.MemoryLimit, p.MemoryLimit)
	assert.Equal(t, problem.TimeLimit, p.TimeLimit)
	assert.Equal(t, problem.Name, p.Name)
	assert.Equal(t, problem.TestCase, p.TestCase)

	return p
}

func TestCreateProblem(t *testing.T) {
	createRandomProblem(t)
}

func TestGetProblem(t *testing.T) {
	p1 := createRandomProblem(t)
	p2, err := GetProblemByID(p1.ID)
	assert.NoError(t, err)

	assert.Equal(t, p1.ID, p2.ID)
	assert.Equal(t, p1.Author, p2.Author)
	assert.Equal(t, p1.Description, p2.Description)
	assert.Equal(t, p1.DifficultyLevel, p2.DifficultyLevel)
	assert.Equal(t, p1.MemoryLimit, p2.MemoryLimit)
	assert.Equal(t, p1.TimeLimit, p2.TimeLimit)
	assert.Equal(t, p1.TestCase, p2.TestCase)
}
