package mysql

import (
	"github.com/Thewalkers2012/DOJ/model"
)

func CreateContext(req *model.CreateContextParams) (*model.Context, error) {

	context := &model.Context{
		Title:     req.Title,
		Author:    req.Author,
		StartTime: req.StartTime,
		EndTime:   req.EndTime,
		Defunct:   "1",
	}

	err := DB.Create(context).Error

	return context, err
}

func GetContextList(offset, limit int) ([]*model.Context, error) {
	contexts := []*model.Context{}
	err := DB.Offset(offset).Limit(limit).Find(&contexts).Error
	return contexts, err
}

func GetContextSize() int64 {
	var count int64
	DB.Model(&model.Context{}).Count(&count)
	return count
}

func GetContextByID(id int64) (*model.Context, error) {
	context := new(model.Context)
	err := DB.Where("id = ?", id).First(context).Error
	return context, err
}

func DeleteContext(id int64) error {
	err := DB.Delete(&model.Context{}, id).Error
	return err
}

func UpdateContext(con *model.Context) (*model.Context, error) {
	err := DB.Save(&con).Error
	return con, err
}

func CreateContextProblem(req *model.AddProblemParams) error {
	con := &model.ContextProblem{
		ProblemID: req.ProblemID,
		ContextID: req.ContextID,
		Title:     req.Title,
	}

	err := DB.Create(con).Error

	return err
}

func ProblemInContext(problemID int64, contextID int64) error {
	contextProblem := new(model.ContextProblem)
	err := DB.Where("problem_id = ? and context_id = ?", problemID, contextID).First(contextProblem).Error
	return err
}

func DeleteProblemInContext(contextID, problemID int64) error {
	err := DB.Where("context_id = ? and problem_id = ?", contextID, problemID).Delete(model.ContextProblem{}).Error
	return err
}

func GetAllContextProblemIDs(contextID int64) ([]*model.ContextProblem, error) {
	contextProblems := []*model.ContextProblem{}
	err := DB.Model(&model.ContextProblem{}).Where("context_id", contextID).Find(&contextProblems).Error
	return contextProblems, err
}
