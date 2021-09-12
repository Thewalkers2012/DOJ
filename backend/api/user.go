package api

import (
	"errors"
	"net/http"

	"github.com/Thewalkers2012/DOJ/model"
	"github.com/Thewalkers2012/DOJ/server"
	"github.com/Thewalkers2012/DOJ/util"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

// create user request
type createUserRequest struct {
	Username   string `json:"username" binding:"required"`
	Password   string `json:"password" binding:"required"`
	StudentID  string `json:"student_id" binding:"required"`
	RePassword string `json:"re_password" binding:"required,eqfield=Password"`
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

	res := createUserResponse{
		Username:  user.Username,
		StudentID: user.StudentID,
	}

	ctx.JSON(http.StatusOK, gin.H{
		"user": res,
	})
}
