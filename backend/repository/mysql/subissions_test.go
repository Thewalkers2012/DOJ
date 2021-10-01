package mysql

import (
	"testing"

	"github.com/Thewalkers2012/DOJ/model"
	"github.com/Thewalkers2012/DOJ/util/judge"
	"github.com/Thewalkers2012/DOJ/util/random"
	"github.com/stretchr/testify/assert"
)

func createRandomSubmission(t *testing.T) {
	user := CreateUser(createRandomUser(t))
	problem := createRandomProblem(t)

	params := &model.CreateSubmissionRequest{
		UserID:    user.ID,
		ProblemID: problem.ID,
		Language:  random.RandomLanguage(),
		Code:      random.RandomStringWithLetter(50),
	}

	score, result := int(random.RandomInt(0, 100)), random.RandomAnswer()

	submission, err := CreateSubmission(params, score, result)
	assert.NoError(t, err)
	assert.Equal(t, submission.Code, params.Code)
	assert.Equal(t, submission.Language, params.Language)
	assert.Equal(t, submission.ProblemID, params.ProblemID)
	assert.Equal(t, submission.UserID, params.UserID)
	assert.Equal(t, submission.Score, score)
	assert.Equal(t, submission.Result, judge.GetAnswerMsg(result))
	assert.NotZero(t, submission.ID)
}

func TestCreateSubmission(t *testing.T) {
	createRandomSubmission(t)
}
