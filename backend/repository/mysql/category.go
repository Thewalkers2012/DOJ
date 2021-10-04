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

func GetAllCategories(offset, limit int) ([]*model.Category, error) {
	categories := []*model.Category{}
	err := DB.Offset(offset).Limit(limit).Find(&categories).Error
	return categories, err
}

func GetCategoriySize() (int64, error) {
	var count int64
	err := DB.Model(&model.Category{}).Count(&count).Error
	return count, err
}

func GetCategoryByID(id int64) (*model.Category, error) {
	category := new(model.Category)
	err := DB.Where("id = ?", id).First(&category).Error
	return category, err
}

func DeleteCategory(id int64) error {
	err := DB.Delete(&model.Category{}, id).Error
	return err
}
