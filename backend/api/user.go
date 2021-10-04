package api

import (
	"errors"
	"net/http"

	"github.com/Thewalkers2012/DOJ/middleware"
	"github.com/Thewalkers2012/DOJ/model"
	"github.com/Thewalkers2012/DOJ/repository/mysql"
	"github.com/Thewalkers2012/DOJ/response"
	"github.com/Thewalkers2012/DOJ/server"
	_ "github.com/Thewalkers2012/DOJ/swagger/docs"
	"github.com/Thewalkers2012/DOJ/util"
	"github.com/Thewalkers2012/DOJ/util/jwt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

const (
	busy                  = "服务器繁忙"
	registerSuccessful    = "注册成功"
	loginSuccessful       = "登录成功"
	infoSuccessful        = "获取用户信息成功"
	updateSuccess         = "更新成功"
	getUserDetailsSuccess = "获取用户详情成功"
)

// User SignUp
func SignUpHandler(ctx *gin.Context) {
	req := new(model.CreateUserRequest)
	if err := ctx.ShouldBindJSON(req); err != nil {
		zap.L().Error("Sign up Handler with invalid param", zap.Error(err))

		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			response.Response(ctx, http.StatusBadRequest, http.StatusBadRequest, gin.H{}, err.Error())
		} else {
			response.Response(ctx, http.StatusBadRequest, http.StatusBadRequest, gin.H{}, removeTopStruct(errs.Translate(trans)))
		}

		return
	}

	hashedPassword, err := util.HashPassword(req.Password)
	if err != nil {
		response.Response(ctx, http.StatusInternalServerError, http.StatusInternalServerError, gin.H{}, busy)
		return
	}

	arg := &model.CreateUserParams{
		Username:       req.Username,
		HashedPassword: hashedPassword,
		StudentID:      req.StudentID,
	}

	user, err := server.CreateUser(arg)
	if err != nil {
		zap.L().Error("server CreateUser function failed, ", zap.Error(err))
		if errors.Is(err, server.ErrorUserHasExists) {
			response.Response(ctx, http.StatusForbidden, http.StatusForbidden, gin.H{}, err.Error())
		} else {
			response.Response(ctx, http.StatusInternalServerError, http.StatusInternalServerError, gin.H{}, busy)
		}
		return
	}

	token, err := jwt.ReleaseToken(user)
	if err != nil {
		response.Response(ctx, http.StatusInternalServerError, http.StatusInternalServerError, gin.H{}, busy)
	}

	res := model.CreateUserResponse{
		Username:  user.Username,
		StudentID: user.StudentID,
	}

	response.Response(ctx, http.StatusOK, http.StatusOK, gin.H{
		"user":         res,
		"access_token": token,
	}, registerSuccessful)
}

// User Login
func LoginHandler(ctx *gin.Context) {
	req := new(model.LoginRequest)
	// 1. 校验数据
	if err := ctx.ShouldBindJSON(req); err != nil {
		zap.L().Error("Login Handler with invalid param", zap.Error(err))

		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			response.Response(ctx, http.StatusBadRequest, http.StatusBadRequest, gin.H{}, err.Error())
		} else {
			response.Response(ctx, http.StatusBadRequest, http.StatusBadRequest, gin.H{}, removeTopStruct(errs.Translate(trans)))
		}

		return
	}

	// 2. server 层调用接口
	arg := &model.LoginParams{
		StudentID: req.StudentID,
		Password:  req.Password,
	}

	token, user, err := server.Login(arg)
	if err != nil {
		zap.L().Error("login failed", zap.Error(err))

		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			response.Response(ctx, http.StatusBadRequest, http.StatusBadRequest, gin.H{}, err.Error())
		} else {
			response.Response(ctx, http.StatusBadRequest, http.StatusBadRequest, gin.H{}, errs.Translate(trans))
		}

		return
	}

	res := model.LoginResponse{
		User: model.CreateUserResponse{
			Username:  user.Username,
			StudentID: user.StudentID,
		},
		AccessToekn: token,
	}

	response.Response(ctx, http.StatusOK, http.StatusOK, gin.H{
		"user": res,
	}, loginSuccessful)
}

// User Info
func InfoHandler(ctx *gin.Context) {
	studentID, _ := ctx.Get(middleware.ContextStudentIDKey)

	user, err := mysql.GetUser(studentID.(string))
	if err != nil {
		zap.L().Error("mysql.GetUser failed", zap.Error(err))

		response.Response(ctx, http.StatusInternalServerError, http.StatusInternalServerError, gin.H{}, busy)
		return
	}

	response.Response(ctx, http.StatusOK, http.StatusOK, gin.H{
		"user": user,
	}, infoSuccessful)
}

// User List
func GetUserList(ctx *gin.Context) {
	req := new(model.GetUserListParams)
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

	users, total, err := server.GetUserList((offset-1)*limit, limit)
	if err != nil {
		response.Response(ctx, http.StatusInternalServerError, http.StatusInternalServerError, gin.H{}, busy)
		return
	}

	response.Response(ctx, http.StatusOK, http.StatusOK, gin.H{
		"users": users,
		"total": total,
	}, getProblemListSuccess)
}

// Update User
func UpdateUser(ctx *gin.Context) {
	req := new(model.UpdateUserParams)
	if err := ctx.ShouldBindJSON(req); err != nil {
		zap.L().Error("api.UpdateUser failed", zap.Error(err))

		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			response.Response(ctx, http.StatusBadRequest, http.StatusBadRequest, gin.H{}, err.Error())
		} else {
			response.Response(ctx, http.StatusBadRequest, http.StatusBadRequest, gin.H{}, removeTopStruct(errs.Translate(trans)))
		}
		return
	}

	user, err := server.UpdateUser(req)
	if err != nil {
		response.Response(ctx, http.StatusInternalServerError, http.StatusInternalServerError, gin.H{}, busy)
		return
	}

	response.Response(ctx, http.StatusOK, http.StatusOK, gin.H{
		"user": user,
	}, updateSuccess)
}

func GetUserDetails(ctx *gin.Context) {
	req := new(model.UserParams)
	if err := ctx.ShouldBindQuery(req); err != nil {
		zap.L().Error("api.GetUserDetails failed", zap.Error(err))

		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			response.Response(ctx, http.StatusBadRequest, http.StatusBadRequest, gin.H{}, err.Error())
		} else {
			response.Response(ctx, http.StatusBadRequest, http.StatusBadRequest, gin.H{}, removeTopStruct(errs.Translate(trans)))
		}
		return
	}

	userDetails, err := server.GetUserDetails(req.UserID)
	if err != nil {
		response.Response(ctx, http.StatusInternalServerError, http.StatusInternalServerError, gin.H{}, busy)
		return
	}

	response.Response(ctx, http.StatusOK, http.StatusOK, gin.H{
		"userDetail": userDetails,
	}, getUserDetailsSuccess)
}
