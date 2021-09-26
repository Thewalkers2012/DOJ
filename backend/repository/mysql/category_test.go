package mysql

import (
	"testing"

	"github.com/Thewalkers2012/DOJ/model"
	"github.com/Thewalkers2012/DOJ/util/random"
	"github.com/stretchr/testify/assert"
)

func createRandomCategory(t *testing.T) (*model.Category, error) {
	user := CreateUser(createRandomUser(t))
	problem := createRandomProblem(t)

	params := &model.CreateCategoryParams{
		ProblemID: problem.ID,
		Title:     random.RandomStringWithLetter(10),
		Content:   random.RandomStringWithLetter(50),
	}

	category, err := CreateCategory(params, user.ID)
	assert.NoError(t, err)
	assert.Equal(t, category.UserID, user.ID)
	assert.Equal(t, category.ProblemID, problem.ID)
	assert.Equal(t, params.Title, category.Title)
	assert.Equal(t, params.Content, category.Content)

	return category, err
}

func TestCreateCategory(t *testing.T) {
	createRandomCategory(t)
}
