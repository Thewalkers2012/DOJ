package server

import (
	"strconv"

	"github.com/Thewalkers2012/DOJ/model"
	"github.com/Thewalkers2012/DOJ/repository/mysql"
	"github.com/Thewalkers2012/DOJ/util/judge"
	"go.uber.org/zap"
)

func CreateSubmission(req *model.RunCodeParams, studentID string, sub *model.SubmitResult) (*model.Submission, error) {
	user, err := mysql.GetUser(studentID)
	if err != nil {
		return nil, err
	}

	return mysql.CreateSubmission(req, user.ID, sub)
}

func GetAllSubmissionsByIdAndProblem(studentID string, problemID int64) ([]*model.Submission, error) {
	user, err := mysql.GetUser(studentID)
	if err != nil {
		return nil, err
	}

	return mysql.GetAllSubmissionsByIdAndProblem(user.ID, problemID)
}

func GetPersonSolved(studentID string) (int64, error) {
	user, err := mysql.GetUser(studentID)
	if err != nil {
		return -1, err
	}

	return mysql.GetPersonSolved(user.ID)
}

func GetPersonSubmission(studentID string) (int64, error) {
	user, err := mysql.GetUser(studentID)
	if err != nil {
		return -1, err
	}

	return mysql.GetPersonSubmission(user.ID)
}

func RunCode(req *model.RunCodeParams) (*judge.Resp, error) {
	pid := req.ProblemID

	// 从数据库中查找该 problem
	problem, err := mysql.GetProblemByID(pid)
	if err != nil {
		zap.L().Error("In server.RunCode mysql.GetProblemByID failed", zap.Error(err))
		return nil, err
	}

	// 获取 judge server 返回的 msg
	resp := judge.SubmitCode(req.Code, req.Language, int64(problem.TimeLimit), int64(problem.MemoryLimit), strconv.Itoa(int(pid)))

	return resp, nil
}
