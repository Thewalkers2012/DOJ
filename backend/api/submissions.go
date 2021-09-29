package api

import (
	"fmt"
	"net/http"

	"github.com/Thewalkers2012/DOJ/middleware"
	"github.com/Thewalkers2012/DOJ/model"
	"github.com/Thewalkers2012/DOJ/repository/mysql"
	"github.com/Thewalkers2012/DOJ/response"
	"github.com/Thewalkers2012/DOJ/server"
	"github.com/Thewalkers2012/DOJ/util/judge"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

const (
	GetAllSubmissionSuccess = "获取所有的提交记录成功"
)

// 返给给前端的结果
type SubmissionResponse struct {
	Score   int    `json:"score"`
	Result  string `json:"result"`
	Message string `json:"message"`
}

func CreateSubmissionHandler(ctx *gin.Context) {
	// 首先获取提交代码的用户
	studentID, _ := ctx.Get(middleware.ContextStudentIDKey)
	// 根据学号来获取 用户的信息
	user, err := mysql.GetUser(studentID.(string))
	if err != nil {
		zap.L().Error("mysql.GetUser failed", zap.Error(err))

		response.Response(ctx, http.StatusInternalServerError, http.StatusInternalServerError, gin.H{}, busy)
		return
	}

	// 然后进行获取参数
	req := new(model.CreateSubmissionRequest)
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

	req.UserID = user.ID

	// TODO: 通过评测机来获取用户的运行结果
	// sub, err := server.CreateSubmission(req)
	if err != nil {
		response.Response(ctx, http.StatusInternalServerError, http.StatusInternalServerError, gin.H{}, busy)
		return
	}

}

// func printSliceData(slice []*judge.Data) {
// 	fmt.Print("[\n")
// 	for _, item := range slice {
// 		fmt.Printf("\t%#v,\n", item)
// 	}
// 	fmt.Print("]\n")
// }

func RunCodeHandler(ctx *gin.Context) {
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

	// 将数据传送给 Client 然后 Client 再将结果放回给 后端
	client := judge.NewClient(judge.WithTimeout(0))
	client.SetOptions(judge.WithEndpointURL("http://127.0.0.1:12358"), judge.WithToken("YOUR_TOKEN_HERE"))
	resp, err := client.JudgeWithRequest(&judge.JudgeRequest{
		Src:            req.Code,
		LanguageConfig: judge.CPPLangConfig,
		MaxCpuTime:     1000,
		MaxMemory:      128 * 1024 * 1024,
		TestCaseId:     "normal",
	})

	fmt.Println("运行时间：", resp.SliceData()[0].CpuTime)
	fmt.Println("运行结果：", resp.SliceData()[0].ExitCode)
	fmt.Println("使用内存：", resp.SliceData()[0].Memory)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	response.Response(ctx, http.StatusOK, http.StatusOK, gin.H{
		"response": req,
	}, "测试结果")
}

type GetAllSubmissionsByIdAndProblemParams struct {
	ProblemID int64 `json:"problem_id"`
}

func GetAllSubmissionsByIdAndProblem(ctx *gin.Context) {
	// 首先获取提交代码的用户
	studentID, _ := ctx.Get(middleware.ContextStudentIDKey)

	req := new(GetAllSubmissionsByIdAndProblemParams)
	if err := ctx.ShouldBindJSON(req); err != nil {
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
