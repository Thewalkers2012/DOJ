package api

import (
	"errors"
	"net/http"

	"github.com/Thewalkers2012/DOJ/middleware"
	"github.com/Thewalkers2012/DOJ/model"
	"github.com/Thewalkers2012/DOJ/repository/mysql"
	"github.com/Thewalkers2012/DOJ/server"
	"github.com/Thewalkers2012/DOJ/util"
	"github.com/Thewalkers2012/DOJ/util/jwt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
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
	StudentID string `json:"student_id"`
}

func SignUpHandler(ctx *gin.Context) {
	req := new(createUserRequest)
	if err := ctx.ShouldBindJSON(req); err != nil {
		zap.L().Error("Sign up Handler with invalid param", zap.Error(err))

		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ctx.JSON(http.StatusBadRequest, responseError(err))
		} else {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"msg": removeTopStruct(errs.Translate(trans)),
			})
		}

		return
	}

	hashedPassword, err := util.HashPassword(req.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, responseError(err))
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
			ctx.JSON(http.StatusForbidden, gin.H{
				"msg": server.ErrorUserHasExists.Error(),
			})
		} else {
			ctx.JSON(http.StatusInternalServerError, responseError(err))
		}
		return
	}

	token, err := jwt.ReleaseToken(user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"data": gin.H{
				"msg": "服务器繁忙",
			},
		})
	}

	res := createUserResponse{
		Username:  user.Username,
		StudentID: user.StudentID,
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": gin.H{
			"user":         res,
			"access_token": token,
		},
	})
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
			ctx.JSON(http.StatusBadRequest, responseError(err))
		} else {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"msg": removeTopStruct(errs.Translate(trans)),
			})
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

		if errors.Is(err, server.ErrorInValidPassword) || errors.Is(err, server.ErrorUserNotExists) {
			ctx.JSON(http.StatusForbidden, responseError(err))
			return
		}

		ctx.JSON(http.StatusInternalServerError, responseBusy(err))
	}

	res := loginResponse{
		User: createUserResponse{
			Username:  user.Username,
			StudentID: user.StudentID,
		},
		AccessToekn: token,
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": res,
	})
}

func InfoHandler(ctx *gin.Context) {
	studentID, _ := ctx.Get(middleware.ContextStudentIDKey)

	user, err := mysql.GetUser(studentID.(string))
	if err != nil {
		zap.L().Error("login failed", zap.Error(err))

		if errors.Is(err, server.ErrorInValidPassword) || errors.Is(err, server.ErrorUserNotExists) {
			ctx.JSON(http.StatusForbidden, responseError(err))
			return
		}

		ctx.JSON(http.StatusInternalServerError, responseBusy(err))
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": gin.H{
			"user": user,
		},
	})
}
