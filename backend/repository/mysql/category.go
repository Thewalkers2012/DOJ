package mysql

import "github.com/Thewalkers2012/DOJ/model"

func CreateCategory(req *model.CreateCategoryParams, userID int64) (*model.Category, error) {
	category := &model.Category{
		UserID:    userID,
		ProblemID: req.ProblemID,
		Title:     req.Title,
		Content:   req.Content,
	}

	err := DB.Create(&category).Error

	return category, err
}

func GetCategoryByProblem(problem int64) ([]*model.Category, error) {
	categories := []*model.Category{}
	err := DB.Where("problem_id = ?", problem).Find(&categories).Error
	return categories, err
}
