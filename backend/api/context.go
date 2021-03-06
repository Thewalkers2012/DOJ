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
	CreateContextSuccess          = "创建比赛成功"
	GetContextSuccess             = "获取比赛成功"
	getContextFiled               = "获取比赛失败"
	AddProblemToContextSuccess    = "添加题目成功"
	QuerySuccess                  = "查询成功"
	DeteleProblemInContextSuccess = "从比赛中删除题目成功"
	NotFound                      = "好像有什么木有找到呢"
	GetContextProblemSuccess      = "获取比赛题目列表成功"
	GetAllContextProblemSuccess   = "获取比赛中的所有题目成功"
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

func AddProblemToContext(ctx *gin.Context) {
	req := new(model.AddProblemParams)
	if err := ctx.ShouldBindJSON(req); err != nil {
		zap.L().Error("api add problem to context failed", zap.Error(err))

		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			response.Response(ctx, http.StatusBadRequest, http.StatusBadRequest, gin.H{}, err.Error())
		} else {
			response.Response(ctx, http.StatusBadRequest, http.StatusBadRequest, gin.H{}, removeTopStruct(errs.Translate(trans)))
		}
		return
	}

	err := server.AddProblemToContext(req)
	if err != nil {
		zap.L().Error("server add problem failed", zap.Error(err))

		response.Response(ctx, http.StatusInternalServerError, http.StatusInternalServerError, gin.H{}, busy)
		return
	}

	response.Response(ctx, http.StatusOK, http.StatusOK, gin.H{}, AddProblemToContextSuccess)
}

func ProblemInContext(ctx *gin.Context) {
	req := new(model.PorblemInContextParams)
	if err := ctx.ShouldBindQuery(req); err != nil {
		zap.L().Error("api problem in context failed", zap.Error(err))

		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			response.Response(ctx, http.StatusBadRequest, http.StatusBadRequest, gin.H{}, err.Error())
		} else {
			response.Response(ctx, http.StatusBadRequest, http.StatusBadRequest, gin.H{}, removeTopStruct(errs.Translate(trans)))
		}
		return
	}

	inThere, err := server.ProblemInContext(req.ProblemID, req.ContextID)
	if err != nil {
		response.Response(ctx, http.StatusInternalServerError, http.StatusInternalServerError, gin.H{}, busy)
		return
	}

	response.Response(ctx, http.StatusOK, http.StatusOK, gin.H{
		"inThere": inThere,
	}, QuerySuccess)
}

func DeletePorblemInContext(ctx *gin.Context) {
	req := new(model.DeleteProblemInContext)
	if err := ctx.ShouldBindQuery(req); err != nil {
		zap.L().Error("api delete problem in context", zap.Error(err))

		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			response.Response(ctx, http.StatusBadRequest, http.StatusBadRequest, gin.H{}, err.Error())
		} else {
			response.Response(ctx, http.StatusBadRequest, http.StatusBadRequest, gin.H{}, removeTopStruct(errs.Translate(trans)))
		}
		return
	}

	err := server.DeleteProblemInContext(req.ContextID, req.ProblemID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			response.Response(ctx, http.StatusNotFound, http.StatusNotFound, gin.H{}, NotFound)
			return
		}

		response.Response(ctx, http.StatusInternalServerError, http.StatusInternalServerError, gin.H{}, busy)
		return
	}

	response.Response(ctx, http.StatusOK, http.StatusOK, gin.H{}, DeteleProblemInContextSuccess)
}

func ContextProblemList(ctx *gin.Context) {
	req := new(model.ContextProblemParams)
	if err := ctx.ShouldBindQuery(req); err != nil {
		zap.L().Error("api context problem list failed", zap.Error(err))

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

	contextProblems, total, err := server.ContextProblemList(req.ContextID, (offset-1)*limit, limit)
	if err != nil {
		response.Response(ctx, http.StatusInternalServerError, http.StatusInternalServerError, gin.H{}, busy)
		return
	}

	response.Response(ctx, http.StatusOK, http.StatusOK, gin.H{
		"contextProblems": contextProblems,
		"total":           total,
	}, GetContextProblemSuccess)
}

func GetAllContextProblem(ctx *gin.Context) {
	req := new(model.GetContextProblemParams)
	if err := ctx.ShouldBindQuery(req); err != nil {
		zap.L().Error("api get all context problems failed", zap.Error(err))

		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			response.Response(ctx, http.StatusBadRequest, http.StatusBadRequest, gin.H{}, err.Error())
		} else {
			response.Response(ctx, http.StatusBadRequest, http.StatusBadRequest, gin.H{}, removeTopStruct(errs.Translate(trans)))
		}
		return
	}

	problems, err := server.GetAllContextProblem(req.ContextID)
	if err != nil {
		response.Response(ctx, http.StatusInternalServerError, http.StatusInternalServerError, gin.H{}, busy)
		return
	}

	response.Response(ctx, http.StatusOK, http.StatusOK, gin.H{
		"problems": problems,
	}, GetAllContextProblemSuccess)
}
