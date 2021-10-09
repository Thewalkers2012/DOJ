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
	CreateContextSuccess = "创建比赛成功"
)

func CreateContextHandler(ctx *gin.Context) {
	req := new(model.CreateContextParams)
	if err := ctx.ShouldBindJSON(req); err != nil {
		zap.L().Error("CreateContext failed", zap.Error(err))
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			response.Response(ctx, http.StatusBadRequest, http.StatusBadRequest, gin.H{}, err.Error())
		} else {
			response.Response(ctx, http.StatusBadRequest, http.StatusBadRequest, gin.H{}, removeTopStruct(errs.Translate(trans)))
		}

		return
	}

	context, err := server.CreateContext(req)
	if err != nil {
		response.Response(ctx, http.StatusInternalServerError, http.StatusInternalServerError, gin.H{}, busy)
		return
	}

	response.Response(ctx, http.StatusOK, http.StatusOK, gin.H{
		"context": context,
	}, CreateContextSuccess)

}

func GetContextList(ctx *gin.Context) {
	req := new(model.GetContextListParams)
	if err := ctx.ShouldBindQuery(req); err != nil {
		zap.L().Error("api GetContextList failed", zap.Error(err))

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

	contexts, total, err := server.GetContextList((offset-1)*limit, limit)
	if err != nil {
		response.Response(ctx, http.StatusInternalServerError, http.StatusInternalServerError, gin.H{}, busy)
		return
	}

	response.Response(ctx, http.StatusOK, http.StatusOK, gin.H{
		"problems": contexts,
		"total":    total,
	}, getProblemListSuccess)
}
