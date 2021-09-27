package server

import (
	"github.com/Thewalkers2012/DOJ/model"
	"github.com/Thewalkers2012/DOJ/repository/mysql"
)

func CreateCategory(req *model.CreateCategoryParams, studentID string) (*model.Category, error) {
	user, err := mysql.GetUser(studentID)
	if err != nil {
		return nil, err
	}

	return mysql.CreateCategory(req, user.ID, user.Username)
}

func GetCategoryByProblem(problemID int64) ([]*model.Category, error) {
	return mysql.GetCategoryByProblem(problemID)
}
