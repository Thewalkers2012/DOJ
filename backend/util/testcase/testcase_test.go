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
				Input:  "1 2",
				Output: "3",
			},
			{
				Input:  "4 10",
				Output: "14",
			},
		},
	}
	pid := random.RandomInt(1, 6)
	err := CreateTestcase(args, pid)
	assert.NoError(t, err)
}
