package api

import (
	"net/http"

	"github.com/Thewalkers2012/DOJ/model"
	"github.com/Thewalkers2012/DOJ/response"
	"github.com/Thewalkers2012/DOJ/server"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

const (
	createProbleSuccessful = "创建题目成功"
)

func CreateProblem(ctx *gin.Context) {
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
	}, createProbleSuccessful)
}
