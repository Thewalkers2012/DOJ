package api

import (
	"errors"
	"net/http"

	"github.com/Thewalkers2012/DOJ/middleware"
	"github.com/Thewalkers2012/DOJ/model"
	"github.com/Thewalkers2012/DOJ/repository/mysql"
	"github.com/Thewalkers2012/DOJ/response"
	"github.com/Thewalkers2012/DOJ/server"
	"github.com/Thewalkers2012/DOJ/util"
	"github.com/Thewalkers2012/DOJ/util/jwt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

const (
	busy               = "服务器繁忙"
	registerSuccessful = "注册成功"
	loginSuccessful    = "登录成功"
	infoSuccessful     = "获取用户信息成功"
)

// create user request
type createUserRequest struct {
	Username  string `json:"username" binding:"required"`
	Password  string `json:"password" binding:"required,min=6"`
	StudentID string `json:"studentID" binding:"required"`
}

// create user response
type createUserResponse struct {
	Username  string `json:"username"`
	StudentID string `json:"studentID"`
}

func SignUpHandler(ctx *gin.Context) {
	req := new(createUserRequest)
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

	res := createUserResponse{
		Username:  user.Username,
		StudentID: user.StudentID,
	}

	response.Response(ctx, http.StatusOK, http.StatusOK, gin.H{
		"user":         res,
		"access_token": token,
	}, registerSuccessful)
}

// login request
type loginRequest struct {
	StudentID string `json:"studentID" binding:"required"`
	Password  string `json:"password" binding:"required"`
}

// login response
type loginResponse struct {
	AccessToekn string             `json:"access_token"`
	User        createUserResponse `json:"user"`
}

func LoginHandler(ctx *gin.Context) {
	req := new(loginRequest)
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

	res := loginResponse{
		User: createUserResponse{
			Username:  user.Username,
			StudentID: user.StudentID,
		},
		AccessToekn: token,
	}

	response.Response(ctx, http.StatusOK, http.StatusOK, gin.H{
		"user": res,
	}, loginSuccessful)
}

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
