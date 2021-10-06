package api

import (
	"net/http"

	"github.com/Thewalkers2012/DOJ/middleware"
	"github.com/Thewalkers2012/DOJ/model"
	"github.com/Thewalkers2012/DOJ/response"
	"github.com/Thewalkers2012/DOJ/server"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

const (
	OK                      = "ok"
	GetAllSubmissionSuccess = "获取所有的提交记录成功"
	SubmitSuccess           = "提交成功"
	Accept                  = 0
	WrongAnswer             = -1
	TimeLimitExceeded       = 1
	MemoryLimitExceeded     = 3
	RunTimeError            = 4
	CompileError            = 6
)

func RunCodeHandler(ctx *gin.Context) {
	studentID, _ := ctx.Get(middleware.ContextStudentIDKey)
	req := new(model.RunCodeParams)

	if err := ctx.ShouldBindJSON(req); err != nil {
		zap.L().Error("CreateSubmissionHandler with invalid param", zap.Error(err))

		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			response.Response(ctx, http.StatusBadRequest, http.StatusBadRequest, gin.H{}, err.Error())
		} else {
			response.Response(ctx, http.StatusBadRequest, http.StatusBadRequest, gin.H{}, removeTopStruct(errs.Translate(trans)))
		}

		return
	}

	resp, err := server.RunCode(req)
	if err != nil {
		response.Response(ctx, http.StatusInternalServerError, http.StatusInternalServerError, gin.H{}, busy)
		return
	}

	// 编译错误
	if resp.Err() != nil {
		_, err := server.CreateSubmission(req, studentID.(string), &model.SubmitResult{
			AnswerCode: 6,
			Time:       0,
			Memory:     0,
			Score:      0,
		})

		if err != nil {
			zap.L().Error("server.CreateSubmission failed", zap.Error(err))
			response.Response(ctx, http.StatusInternalServerError, http.StatusInternalServerError, gin.H{}, busy)
			return
		}

		response.Response(ctx, http.StatusOK, http.StatusOK, gin.H{
			"data": model.SubmitResult{
				AnswerCode: CompileError,
				Time:       0,
				Memory:     0,
				Score:      0,
			},
		}, resp.Err().Error())
		return
	}

	// 获取所有的测试样例的结果
	res := resp.SliceData()

	total := len(res)     // 样例总数
	score := float64(100) // 最终得分
	avge := float64(100.0 / float64(total))
	time := 0   // 平均时间
	memory := 0 // 平均内存
	result := 0 // 最终结果

	for i := 0; i < total; i++ {
		if res[i].Result != 0 {
			result = res[i].Result
			score -= avge
		}
		time += res[i].RealTime
		memory += int(res[i].Memory)
	}

	time /= total          // 转化为 ms
	memory /= total * 1024 // 转化为 kb

	// 将提交结果保存到数据库中
	_, err = server.CreateSubmission(req, studentID.(string), &model.SubmitResult{
		AnswerCode: result,
		Time:       time,
		Memory:     int64(memory),
		Score:      int(score),
	})

	if err != nil {
		zap.L().Error("server.CreateSubmission failed", zap.Error(err))
		response.Response(ctx, http.StatusInternalServerError, http.StatusInternalServerError, gin.H{}, busy)
		return
	}

	response.Response(ctx, http.StatusOK, http.StatusOK, gin.H{
		"data": model.SubmitResult{
			AnswerCode: result,
			Time:       time,
			Memory:     int64(memory),
			Score:      int(score),
		},
	}, SubmitSuccess)
}

type GetAllSubmissionsByIdAndProblemParams struct {
	ProblemID int64 `form:"problemID"`
}

func GetAllSubmissionsByIdAndProblem(ctx *gin.Context) {
	// 首先获取提交代码的用户
	studentID, _ := ctx.Get(middleware.ContextStudentIDKey)

	req := new(GetAllSubmissionsByIdAndProblemParams)
	if err := ctx.ShouldBindQuery(req); err != nil {
		zap.L().Error("GetAllSubmissionsByIdAndProblem with invalid params", zap.Error(err))

		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			response.Response(ctx, http.StatusBadRequest, http.StatusBadRequest, gin.H{}, err.Error())
		} else {
			response.Response(ctx, http.StatusBadRequest, http.StatusBadRequest, gin.H{}, removeTopStruct(errs.Translate(trans)))
		}
		return
	}

	submisstions, err := server.GetAllSubmissionsByIdAndProblem(studentID.(string), req.ProblemID)
	if err != nil {
		zap.L().Error("server.GetAllSubmittionsByIdAndProblem failed", zap.Error(err))
		response.Response(ctx, http.StatusInternalServerError, http.StatusInternalServerError, gin.H{}, busy)
		return
	}

	response.Response(ctx, http.StatusOK, http.StatusOK, gin.H{
		"submission": submisstions,
	}, GetAllSubmissionSuccess)
}

// 获取用户的解题数目
func GetPersonSolved(ctx *gin.Context) {
	studentID, _ := ctx.Get(middleware.ContextStudentIDKey)

	solvedSum, err := server.GetPersonSolved(studentID.(string))
	if err != nil {
		zap.L().Error("server.GetPersonSolved failed", zap.Error(err))
		response.Response(ctx, http.StatusInternalServerError, http.StatusInternalServerError, gin.H{}, busy)
		return
	}

	response.Response(ctx, http.StatusOK, http.StatusOK, gin.H{
		"total": solvedSum,
	}, OK)
}

// 获取单个用户的提交总数
func GetPersonSubmission(ctx *gin.Context) {
	studentID, _ := ctx.Get(middleware.ContextStudentIDKey)

	solvedSum, err := server.GetPersonSubmission(studentID.(string))
	if err != nil {
		zap.L().Error("server.GetPersonSubmission failed", zap.Error(err))
		response.Response(ctx, http.StatusInternalServerError, http.StatusInternalServerError, gin.H{}, busy)
		return
	}

	response.Response(ctx, http.StatusOK, http.StatusOK, gin.H{
		"total": solvedSum,
	}, OK)
}
