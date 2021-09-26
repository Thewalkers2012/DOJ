package server

import (
	"github.com/Thewalkers2012/DOJ/model"
	"github.com/Thewalkers2012/DOJ/repository/mysql"
)

func CreateSubmission(req *model.CreateSubmissionRequest) (*model.Submissions, error) {
	return nil, nil
}

func GetAllSubmissionsByIdAndProblem(studentID string, problemID int64) ([]*model.Submissions, error) {
	user, err := mysql.GetUser(studentID)
	if err != nil {
		return nil, err
	}

	return mysql.GetAllSubmissionsByIdAndProblem(user.ID, problemID)
}
