package api

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/Thewalkers2012/DOJ/model"
	"github.com/Thewalkers2012/DOJ/response"
	"github.com/Thewalkers2012/DOJ/server"
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
)

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

	problem := server.CreateProblem(req)

	response.Response(ctx, http.StatusOK, http.StatusOK, gin.H{
		"problem": problem,
	}, createProblemSuccessful)
}

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

type listProblemRequest struct {
	PageNum  int `form:"pageNum" binding:"required,min=1"`
	PageSize int `form:"pageSize" binding:"required,min=5,max=10"`
}

func GetProblemList(ctx *gin.Context) {
	req := new(listProblemRequest)
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

	problems, err, total := server.GetProblemList((offset-1)*limit, limit)
	if err != nil {
		response.Response(ctx, http.StatusInternalServerError, http.StatusInternalServerError, gin.H{}, busy)
		return
	}

	response.Response(ctx, http.StatusOK, http.StatusOK, gin.H{
		"problems": problems,
		"total":    total,
	}, getProblemListSuccess)
}
