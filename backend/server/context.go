package server

import (
	"time"

	"github.com/Thewalkers2012/DOJ/model"
	"github.com/Thewalkers2012/DOJ/repository/mysql"
	"go.uber.org/zap"
)

func CreateContext(req *model.CreateContextParams) (*model.Context, error) {
	return mysql.CreateContext(req)
}

func GetContextList(offset, limit int) ([]*model.Context, int64, error) {
	contexts, err := mysql.GetContextList(offset, limit)
	total := mysql.GetContextSize()

	for i := 0; i < len(contexts); i++ {
		if time.Time(contexts[i].StartTime).After(time.Now()) {
			contexts[i].Defunct = "1"
			continue
		}
		if time.Time(contexts[i].EndTime).After(time.Now()) {
			contexts[i].Defunct = "2"
			continue
		}
		contexts[i].Defunct = "3"
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
