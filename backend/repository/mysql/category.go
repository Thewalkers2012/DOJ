package mysql

import (
	"github.com/Thewalkers2012/DOJ/model"
)

func CreateCategory(req *model.CreateCategoryParams, userID int64, username string) (*model.Category, error) {
	category := &model.Category{
		UserID:    userID,
		ProblemID: req.ProblemID,
		Content:   req.Content,
		Author:    username,
	}

	err := DB.Select("UserID", "ProblemID", "Content", "Author").Create(category).Error

	return category, err
}

func GetCategoryByProblem(problem int64) ([]*model.Category, error) {
	categories := []*model.Category{}
	err := DB.Where("problem_id = ?", problem).Find(&categories).Error
	return categories, err
}
