package server

import (
	"github.com/Thewalkers2012/DOJ/model"
	"github.com/Thewalkers2012/DOJ/repository/mysql"
)

func CreateContext(req *model.CreateContextParams) (*model.Context, error) {
	return mysql.CreateContext(req)
}

func GetContextList(offset, limit int) ([]*model.Context, int64, error) {
	contexts, err := mysql.GetContextList(offset, limit)
	total := mysql.GetContextSize()

	return contexts, total, err
}

func GetContext(id int64) (*model.Context, error) {
	return mysql.GetContextByID(id)
}
