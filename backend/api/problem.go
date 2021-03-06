package api

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/Thewalkers2012/DOJ/model"
	"github.com/Thewalkers2012/DOJ/response"
	"github.com/Thewalkers2012/DOJ/server"
	_ "github.com/Thewalkers2012/DOJ/swagger/docs"
	"github.com/Thewalkers2012/DOJ/util/testcase"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

const (
	createProblemSuccessful = "创建题目成功"
	getProblemIdFailed      = "获取题目ID失败"
	getProblemSuccess       = "获取题目成功"
	invalidParams           = "请求参数错误"
	getProblemListSuccess   = "获取问题列表成功"
	deleteProblemSuccess    = "删除题目成功"
)

// Create Problem
func CreateProblemHandler(ctx *gin.Context) {
	req := new(model.CreateProblemRequest)
	if err := ctx.ShouldBindJSON(req); err != nil {
		zap.L().Error("api.CreateProblem failed", zap.Error(err))
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			response.Response(ctx, http.StatusBadRequest, http.StatusBadRequest, gin.H{}, err.Error())
		} else {
			response.Response(ctx, http.StatusBadRequest, http.StatusBadRequest, gin.H{}, removeTopStruct(errs.Translate(trans)))
		}

		return
	}

	// 1. 创建 info 变量
	info := &model.Info{
		TestCaseNum: len(req.Cases),
		Spj:         false,
		TestCases:   []model.TestCase{},
	}

	for i := 0; i < len(req.Cases); i++ {
		info.TestCases = append(info.TestCases, model.TestCase{
			Input:  req.Cases[i].Input,
			Output: req.Cases[i].Output,
		})
	}

	// 2. 在数据库中建立该题目
	problem := server.CreateProblem(req)

	// 3. 将 info 以及 testcase 存储到相应的位置
	err := testcase.CreateTestcase(info, problem.ID)
	if err != nil {
		response.Response(ctx, http.StatusInternalServerError, http.StatusInternalServerError, gin.H{}, busy)
		return
	}

	response.Response(ctx, http.StatusOK, http.StatusOK, gin.H{
		"problem": problem,
	}, createProblemSuccessful)
}

// Get Problem By ID
func GetProblemByIDHandler(ctx *gin.Context) {
	// 获取参数
	pidStr := ctx.Param("id")
	pid, err := strconv.ParseInt(pidStr, 10, 64)
	if err != nil {
		response.Response(ctx, http.StatusBadRequest, http.StatusBadRequest, gin.H{}, getProblemIdFailed)
		return
	}

	problem, err := server.GetProblemByID(pid)
	if err != nil {
		zap.L().Error("server.GetProblemById failed", zap.Error(err))
		if errors.Is(err, server.ErrorProblemNotExist) {
			response.Response(ctx, http.StatusNotFound, http.StatusNotFound, gin.H{}, getProblemIdFailed)
		} else {
			response.Response(ctx, http.StatusInternalServerError, http.StatusInternalServerError, gin.H{}, busy)
		}
		return
	}

	response.Response(ctx, http.StatusOK, http.StatusOK, gin.H{
		"problem": problem,
	}, getProblemSuccess)
}

func GetProblemList(ctx *gin.Context) {
	req := new(model.ListProblemRequest)
	if err := ctx.ShouldBindQuery(req); err != nil {
		zap.L().Error("api.GetProblem failed", zap.Error(err))

		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			response.Response(ctx, http.StatusBadRequest, http.StatusBadRequest, gin.H{}, err.Error())
		} else {
			response.Response(ctx, http.StatusBadRequest, http.StatusBadRequest, gin.H{}, removeTopStruct(errs.Translate(trans)))
		}
		return
	}

	offset := req.PageNum
	limit := req.PageSize

	problems, total, err := server.GetProblemList((offset-1)*limit, limit)
	if err != nil {
		response.Response(ctx, http.StatusInternalServerError, http.StatusInternalServerError, gin.H{}, busy)
		return
	}

	response.Response(ctx, http.StatusOK, http.StatusOK, gin.H{
		"problems": problems,
		"total":    total,
	}, getProblemListSuccess)
}

// delete problem
func DeleteProblem(ctx *gin.Context) {
	req := new(model.DeleteProblemParams)
	if err := ctx.ShouldBindQuery(req); err != nil {
		zap.L().Error("api.DeleteProblem failed", zap.Error(err))

		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			response.Response(ctx, http.StatusBadRequest, http.StatusBadRequest, gin.H{}, err.Error())
		} else {
			response.Response(ctx, http.StatusBadRequest, http.StatusBadRequest, gin.H{}, removeTopStruct(errs.Translate(trans)))
		}
		return
	}

	err := server.DeleteProblem(req.ProblemID)
	if err != nil {
		response.Response(ctx, http.StatusInternalServerError, http.StatusInternalServerError, gin.H{}, busy)
		return
	}

	response.Response(ctx, http.StatusOK, http.StatusOK, gin.H{}, deleteProblemSuccess)
}
