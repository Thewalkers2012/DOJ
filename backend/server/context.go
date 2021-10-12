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
