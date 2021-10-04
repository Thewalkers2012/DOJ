package server

import (
	"github.com/Thewalkers2012/DOJ/model"
	"github.com/Thewalkers2012/DOJ/repository/mysql"
	"go.uber.org/zap"
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

func GetAllCategories(offset, limit int) ([]*model.Category, int64, error) {
	categories, err := mysql.GetAllCategories(offset, limit)
	if err != nil {
		zap.L().Error("mysql.GetAllCategories failed", zap.Error(err))
		return nil, 0, err
	}
	total, err := mysql.GetCategoriySize()
	if err != nil {
		zap.L().Error("mysql.GetCategoriySize failed", zap.Error(err))
		return nil, 0, err
	}

	return categories, total, nil
}
