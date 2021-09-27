package api

import (
	"net/http"

	"github.com/Thewalkers2012/DOJ/middleware"
	"github.com/Thewalkers2012/DOJ/model"
	"github.com/Thewalkers2012/DOJ/response"
	"github.com/Thewalkers2012/DOJ/server"
	_ "github.com/Thewalkers2012/DOJ/swagger/docs"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

const (
	GetCategoryByProblemSuccess = "获取题目讨论成功"
	CreateCategorySuccess       = "创建讨论成功"
)

// Create Category
func CreateCategory(ctx *gin.Context) {
	studentID, _ := ctx.Get(middleware.ContextStudentIDKey)
	req := new(model.CreateCategoryParams)
	if err := ctx.ShouldBindJSON(req); err != nil {
		zap.L().Error("CreateCategory failed", zap.Error(err))
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			response.Response(ctx, http.StatusBadRequest, http.StatusBadRequest, gin.H{}, err.Error())
		} else {
			response.Response(ctx, http.StatusBadRequest, http.StatusBadRequest, gin.H{}, removeTopStruct(errs.Translate(trans)))
		}

		return
	}

	category, err := server.CreateCategory(req, studentID.(string))
	if err != nil {
		zap.L().Error("server.CreateCategory failed", zap.Error(err))
		response.Response(ctx, http.StatusInternalServerError, http.StatusInternalServerError, gin.H{}, busy)
		return
	}

	response.Response(ctx, http.StatusOK, http.StatusOK, gin.H{
		"category": category,
	}, CreateCategorySuccess)
}

// Get Category By Problem ID
func GetCategoryByProblem(ctx *gin.Context) {
	req := new(model.GetCategoryByProblemParams)
	if err := ctx.ShouldBindQuery(req); err != nil {
		zap.L().Error("CreateCategory failed", zap.Error(err))
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			response.Response(ctx, http.StatusBadRequest, http.StatusBadRequest, gin.H{}, err.Error())
		} else {
			response.Response(ctx, http.StatusBadRequest, http.StatusBadRequest, gin.H{}, removeTopStruct(errs.Translate(trans)))
		}

		return
	}

	categories, err := server.GetCategoryByProblem(req.ProblemID)
	if err != nil {
		zap.L().Error("server.GetCategoryByProblem failed", zap.Error(err))

		response.Response(ctx, http.StatusInternalServerError, http.StatusInternalServerError, gin.H{}, busy)
		return
	}

	response.Response(ctx, http.StatusOK, http.StatusOK, gin.H{
		"categories": categories,
	}, GetCategoryByProblemSuccess)
}
