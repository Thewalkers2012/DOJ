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
	"gorm.io/gorm"
)

const (
	CreateContextSuccess = "创建比赛成功"
	GetContextSuccess    = "获取比赛成功"
	getContextFiled      = "获取比赛失败"
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
		"contexts": contexts,
		"total":    total,
	}, getProblemListSuccess)
}

func GetContext(ctx *gin.Context) {
	cidStr := ctx.Param("id")
	cid, err := strconv.ParseInt(cidStr, 10, 64)
	if err != nil {
		response.Response(ctx, http.StatusBadRequest, http.StatusBadRequest, gin.H{}, err.Error())
		return
	}

	context, err := server.GetContext(cid)
	if err != nil {
		zap.L().Error("server.GetContext failed", zap.Error(err))
		if errors.Is(err, gorm.ErrRecordNotFound) {
			response.Response(ctx, http.StatusNotFound, http.StatusNotFound, gin.H{}, getContextFiled)
		} else {
			response.Response(ctx, http.StatusInternalServerError, http.StatusInternalServerError, gin.H{}, busy)
		}
		return
	}

	response.Response(ctx, http.StatusOK, http.StatusOK, gin.H{
		"context": context,
	}, GetContextSuccess)
}

// delete context
func DeleteContext(ctx *gin.Context) {
	req := new(model.DeleteContextParams)
	if err := ctx.ShouldBindQuery(req); err != nil {
		zap.L().Error("Delete Context failed", zap.Error(err))

		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			response.Response(ctx, http.StatusBadRequest, http.StatusBadRequest, gin.H{}, err.Error())
		} else {
			response.Response(ctx, http.StatusBadRequest, http.StatusBadRequest, gin.H{}, removeTopStruct(errs.Translate(trans)))
		}
		return
	}

	err := server.DeleteContext(req.ContextID)
	if err != nil {
		response.Response(ctx, http.StatusInternalServerError, http.StatusInternalServerError, gin.H{}, busy)
		return
	}

	response.Response(ctx, http.StatusOK, http.StatusOK, gin.H{}, deleteProblemSuccess)
}

// update context
func UpdateContext(ctx *gin.Context) {
	req := new(model.UpdateContextParams)
	if err := ctx.ShouldBindJSON(req); err != nil {
		zap.L().Error("api update context failed", zap.Error(err))

		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			response.Response(ctx, http.StatusBadRequest, http.StatusBadRequest, gin.H{}, err.Error())
		} else {
			response.Response(ctx, http.StatusBadRequest, http.StatusBadRequest, gin.H{}, removeTopStruct(errs.Translate(trans)))
		}
		return
	}

	context, err := server.UpdateContext(req)
	if err != nil {
		response.Response(ctx, http.StatusInternalServerError, http.StatusInternalServerError, gin.H{}, busy)
		return
	}

	response.Response(ctx, http.StatusOK, http.StatusOK, gin.H{
		"context": context,
	}, updateSuccess)
}
