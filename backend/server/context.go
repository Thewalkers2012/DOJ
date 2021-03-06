package server

import (
	"errors"
	"time"

	"github.com/Thewalkers2012/DOJ/model"
	"github.com/Thewalkers2012/DOJ/repository/mysql"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func CreateContext(req *model.CreateContextParams) (*model.Context, error) {
	return mysql.CreateContext(req)
}

func GetContextList(offset, limit int) ([]*model.Context, int64, error) {
	contexts, err := mysql.GetContextList(offset, limit)
	total := mysql.GetContextSize()

	for i := 0; i < len(contexts); i++ {
		if time.Time(contexts[i].StartTime).After(time.Now()) {
			// 比赛尚未开始
			if contexts[i].Defunct != "1" {
				_, err := mysql.UpdateContext(contexts[i])
				if err != nil {
					zap.L().Error("change defunct failed", zap.Error(err))
					return nil, 0, err
				}

				contexts[i].Defunct = "1"
			}
			continue
		}

		// 比赛已经进行
		if time.Time(contexts[i].EndTime).After(time.Now()) {
			if contexts[i].Defunct != "2" {
				contexts[i].Defunct = "2"
				_, err := mysql.UpdateContext(contexts[i])
				if err != nil {
					zap.L().Error("change defunct failed", zap.Error(err))
					return nil, 0, err
				}
			}

			continue
		}

		// 比赛已经结束
		if contexts[i].Defunct != "3" {
			contexts[i].Defunct = "3"
			_, err := mysql.UpdateContext(contexts[i])
			if err != nil {
				zap.L().Error("change defunct failed", zap.Error(err))
				return nil, 0, err
			}

		}
	}

	return contexts, total, err
}

func GetContext(id int64) (*model.Context, error) {
	return mysql.GetContextByID(id)
}

// DeleteContext 删除题目
func DeleteContext(id int64) error {
	return mysql.DeleteContext(id)
}

func UpdateContext(req *model.UpdateContextParams) (*model.Context, error) {

	context, err := mysql.GetContextByID(req.ID)
	if err != nil {
		zap.L().Error("mysql.GetContextByID failed", zap.Error(err))
		return nil, err
	}

	context.Title = req.Title
	context.Author = req.Author
	context.StartTime = req.StartTime
	context.EndTime = req.EndTime

	return mysql.UpdateContext(context)
}

func AddProblemToContext(req *model.AddProblemParams) error {
	return mysql.CreateContextProblem(req)
}

func ProblemInContext(problemID, contextID int64) (bool, error) {
	_, err := mysql.GetProblemByID(problemID)
	if err != nil {
		zap.L().Error("mysql GetProblemByID failed", zap.Error(err))

		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}

		return false, err
	}

	_, err = mysql.GetContextByID(contextID)
	if err != nil {
		zap.L().Error("mysql GetContextByID failed", zap.Error(err))

		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}

		return false, err
	}

	err = mysql.ProblemInContext(problemID, contextID)
	if err != nil {
		zap.L().Error("mysql ProblemInContext failed", zap.Error(err))

		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}

		return false, err
	}

	return true, nil
}

func DeleteProblemInContext(contextID, problemID int64) error {
	_, err := mysql.GetProblemByID(problemID)
	if err != nil {
		zap.L().Error("mysql GetProblemByID failed", zap.Error(err))
		return err
	}

	_, err = mysql.GetContextByID(contextID)
	if err != nil {
		zap.L().Error("mysql GetContextByID failed", zap.Error(err))
		return err
	}

	return mysql.DeleteProblemInContext(contextID, problemID)
}

func ContextProblemList(contextID int64, offset, limit int) ([]*model.ContextProblemResponse, int64, error) {
	problems, err := mysql.GetProblemList(offset, limit)
	total := mysql.GetProblemSize()

	n := len(problems)
	res := make([]*model.ContextProblemResponse, n)

	for i := 0; i < n; i++ {
		inThere, _ := ProblemInContext(problems[i].ID, contextID)

		res[i] = &model.ContextProblemResponse{
			ProblemID:       problems[i].ID,
			TimeLimit:       problems[i].TimeLimit,
			MemoryLimit:     problems[i].MemoryLimit,
			ProblemName:     problems[i].Name,
			Author:          problems[i].Author,
			DifficultyLevel: problems[i].DifficultyLevel,
			InContext:       inThere,
		}
	}

	return res, total, err
}

func GetAllContextProblem(contextID int64) ([]*model.Problem, error) {
	contextProblems, err := mysql.GetAllContextProblemIDs(contextID)
	if err != nil {
		zap.L().Error("mysql GetAllContextProblemIDs failed", zap.Error(err))
		return nil, err
	}

	n := len(contextProblems)

	problems := make([]*model.Problem, n)
	for i := 0; i < n; i++ {
		problems[i], err = mysql.GetProblemByID(contextProblems[i].ProblemID)
		if err != nil {
			zap.L().Error("mysql GetProblemByID failed", zap.Error(err))
			return nil, err
		}
	}

	return problems, err
}
