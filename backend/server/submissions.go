package server

import (
	"github.com/Thewalkers2012/DOJ/model"
	"github.com/Thewalkers2012/DOJ/repository/mysql"
)

func CreateSubmission(req *model.CreateSubmissionRequest, score int, result int) (*model.Submission, error) {
	return mysql.CreateSubmission(req, score, result)
}

func GetAllSubmissionsByIdAndProblem(studentID string, problemID int64) ([]*model.Submission, error) {
	user, err := mysql.GetUser(studentID)
	if err != nil {
		return nil, err
	}

	return mysql.GetAllSubmissionsByIdAndProblem(user.ID, problemID)
}
