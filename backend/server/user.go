package server

import (
	"errors"

	"github.com/Thewalkers2012/DOJ/model"
	"github.com/Thewalkers2012/DOJ/repository/mysql"
	"github.com/Thewalkers2012/DOJ/util"
	"github.com/Thewalkers2012/DOJ/util/jwt"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	ErrorUserHasExists   = errors.New("该用户已经存在")
	ErrorUserNotExists   = errors.New("该用户不存在")
	ErrorInValidPassword = errors.New("密码错误")
)

func CreateUser(arg *model.CreateUserParams) (user *model.User, err error) {
	// 1. 判断用户是否存在
	if ok := mysql.CheckUserExists(arg.StudentID); !ok {
		return nil, ErrorUserHasExists
	}

	u := &model.User{
		Username:       arg.Username,
		HashedPassword: arg.HashedPassword,
		StudentID:      arg.StudentID,
	}

	// 2. dao 层创建角色
	user = mysql.CreateUser(u)

	return user, nil
}

func Login(arg *model.LoginParams) (token string, user *model.User, err error) {
	user, err = mysql.GetUser(arg.StudentID)

	// 1. 能否找到该用户
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return "", nil, ErrorUserNotExists
		}
		return "", nil, err
	}

	// 2. 用户的密码是否正确
	err = util.CheckPassword(arg.Password, user.HashedPassword)
	if err != nil {
		return "", nil, ErrorInValidPassword
	}

	// 3. 发放 token
	token, err = jwt.ReleaseToken(user)
	if err != nil {
		return "", nil, err
	}

	return token, user, nil
}

// GetUserList
func GetUserList(offset, limit int) ([]*model.User, int64, error) {
	users, err := mysql.GetUserList(offset, limit)
	total := mysql.GetUserSize()

	return users, total, err
}

// UpdateUser
func UpdateUser(req *model.UpdateUserParams) (*model.User, error) {
	user, err := mysql.GetUserByID(req.UserID)
	if err != nil {
		zap.L().Error("mysql.GetUserByID failed", zap.Error(err))
		return nil, err
	}

	user.StudentID = req.StudentID
	user.Username = req.Username

	return mysql.UpdateUser(user)
}

// GetUserDetails
func GetUserDetails(userID int64) (*model.UserDetailResponse, error) {
	user, err := mysql.GetUserByID(userID)
	if err != nil {
		zap.L().Error("mysql.GetUserByID failed", zap.Error(err))
		return nil, err
	}

	accept, err := mysql.GetPersonSolved(userID)
	if err != nil {
		zap.L().Error("mysql.GetPersonSolved failed", zap.Error(err))
		return nil, err
	}

	submissionCount, err := mysql.GetPersonSubmission(userID)
	if err != nil {
		zap.L().Error("mysql.GetPersonSubmission failed", zap.Error(err))
		return nil, err
	}

	return &model.UserDetailResponse{
		UserID:          userID,
		Username:        user.Username,
		StudentID:       user.StudentID,
		AcceptCount:     accept,
		SubmissionCount: submissionCount,
	}, nil
}
