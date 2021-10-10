package server

import (
	"time"

	"github.com/Thewalkers2012/DOJ/model"
	"github.com/Thewalkers2012/DOJ/repository/mysql"
)

func CreateContext(req *model.CreateContextParams) (*model.Context, error) {
	return mysql.CreateContext(req)
}

func GetContextList(offset, limit int) ([]*model.Context, int64, error) {
	contexts, err := mysql.GetContextList(offset, limit)
	total := mysql.GetContextSize()

	for i := 0; i < len(contexts); i++ {
		if time.Time(contexts[i].StartTime).After(time.Now()) {
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
