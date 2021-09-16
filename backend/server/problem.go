package server

import (
	"github.com/Thewalkers2012/DOJ/model"
	"github.com/Thewalkers2012/DOJ/repository/mysql"
)

func CreateProblem(req *model.CreateProblemRequest) *model.Problem {
	return mysql.CreateProblem(req)
}
