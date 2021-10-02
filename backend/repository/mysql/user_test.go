package mysql

import (
	"testing"

	"github.com/Thewalkers2012/DOJ/model"
	"github.com/Thewalkers2012/DOJ/util"
	"github.com/Thewalkers2012/DOJ/util/random"
	"github.com/stretchr/testify/assert"
)

// createRandomUser create a random user
func createRandomUser(t *testing.T) *model.User {
	hashedPassword, err := util.HashPassword(random.RandomPassword())
	assert.NoError(t, err)

	arg := &model.User{
		StudentID:      random.RandomStudentID(),
		Username:       random.RandomUserName(),
		HashedPassword: hashedPassword,
	}

	return arg
}

func TestCreateUser(t *testing.T) {
	user := createRandomUser(t)
	u := CreateUser(user)

	assert.Equal(t, user.StudentID, u.StudentID)
	assert.Equal(t, user.Username, u.Username)
}

func TestCheckUserExists(t *testing.T) {
	user := createRandomUser(t)
	u := CreateUser(user)

	assert.Equal(t, user.StudentID, u.StudentID)
	assert.Equal(t, user.Username, u.Username)

	ok := CheckUserExists(user.StudentID)
	assert.Equal(t, ok, false)
	ok = CheckUserExists(random.RandomStringWithLetter(8))
	assert.Equal(t, ok, true)
}

func TestGetUser(t *testing.T) {
	user := createRandomUser(t)
	u1 := CreateUser(user)
	u2, err := GetUser(u1.StudentID)

	assert.NoError(t, err)
	assert.Equal(t, u1.HashedPassword, u2.HashedPassword)
	assert.Equal(t, u1.ID, u2.ID)
	assert.Equal(t, u1.StudentID, u2.StudentID)
	assert.Equal(t, u1.Username, u2.Username)
}
