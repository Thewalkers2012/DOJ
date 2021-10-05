package testcase

import (
	"testing"

	"github.com/Thewalkers2012/DOJ/model"
	"github.com/Thewalkers2012/DOJ/util/random"
	"github.com/stretchr/testify/assert"
)

func TestCreateTestcase(t *testing.T) {
	args := &model.Info{
		TestCaseNum: 2,
		Spj:         false,
		TestCases: []model.TestCase{
			{
				Input:  random.RandomStringWithNumber(20),
				Output: random.RandomStringWithNumber(20),
			},
			{
				Input:  random.RandomStringWithNumber(20),
				Output: random.RandomStringWithNumber(20),
			},
		},
	}
	pid := random.RandomInt(1, 6)
	err := CreateTestcase(args, pid)
	assert.NoError(t, err)
}