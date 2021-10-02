package server

import (
	"github.com/Thewalkers2012/DOJ/model"
	"github.com/Thewalkers2012/DOJ/repository/mysql"
)

func CreateSubmission(req *model.RunCodeParams, studentID string, sub *model.SubmitResult) (*model.Submission, error) {
	user, err := mysql.GetUser(studentID)
	if err != nil {
		return nil, err
	}

	return mysql.CreateSubmission(req, user.ID, sub)
}

func GetAllSubmissionsByIdAndProblem(studentID string, problemID int64) ([]*model.Submission, error) {
	user, err := mysql.GetUser(studentID)
	if err != nil {
		return nil, err
	}

	return mysql.GetAllSubmissionsByIdAndProblem(user.ID, problemID)
}

func GetPersonSolved(studentID string) (int64, error) {
	user, err := mysql.GetUser(studentID)
	if err != nil {
		return -1, err
	}

	return mysql.GetPersonSolved(user.ID)
}

func GetPersonSubmission(studentID string) (int64, error) {
	user, err := mysql.GetUser(studentID)
	if err != nil {
		return -1, err
	}

	return mysql.GetPersonSubmission(user.ID)
}
