package server

import (
	"errors"

	"github.com/Thewalkers2012/DOJ/model"
	"github.com/Thewalkers2012/DOJ/repository/mysql"
	"gorm.io/gorm"
)

// 数据库题目产生的错误
var (
	ErrorProblemNotExist = errors.New("该题目不存在")
)

func CreateProblem(req *model.CreateProblemRequest) *model.Problem {
	return mysql.CreateProblem(req)
}

// 根据 problem 的 id 来获取参数
func GetProblemByID(pid int64) (*model.Problem, error) {
	problem, err := mysql.GetProblemByID(pid)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrorProblemNotExist
		} else {
			return nil, err
		}
	}

	return problem, nil
}

// GetProblemList 获取 problem 的列表
func GetProblemList(offset, limit int) ([]*model.Problem, error) {
	return mysql.GetProblemList(offset, limit)
}
