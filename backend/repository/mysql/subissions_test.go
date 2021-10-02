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

	params := &model.RunCodeParams{
		ProblemID: problem.ID,
		Language:  random.RandomLanguage(),
		Code:      random.RandomStringWithLetter(50),
	}

	sub := &model.SubmitResult{
		AnswerCode: random.RandomAnswer(),
		Score:      int(random.RandomInt(0, 100)),
		Time:       int(random.RandomInt(0, 1000)),
		Memory:     random.RandomInt(0, 1000),
	}

	submission, err := CreateSubmission(params, user.ID, sub)
	assert.NoError(t, err)
	assert.Equal(t, submission.Code, params.Code)
	assert.Equal(t, submission.Language, params.Language)
	assert.Equal(t, submission.ProblemID, params.ProblemID)
	assert.Equal(t, submission.UserID, user.ID)
	assert.Equal(t, submission.Score, sub.Score)
	assert.Equal(t, submission.Result, judge.GetAnswerMsg(sub.AnswerCode))
	assert.Equal(t, submission.Time, sub.Time)
	assert.Equal(t, submission.Memory, sub.Memory)
	assert.NotZero(t, submission.ID)
}

func TestCreateSubmission(t *testing.T) {
	createRandomSubmission(t)
}
