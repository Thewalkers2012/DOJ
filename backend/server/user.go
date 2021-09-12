package server

import (
	"errors"

	"github.com/Thewalkers2012/DOJ/model"
	"github.com/Thewalkers2012/DOJ/repository/mysql"
)

var (
	ErrorUserHasExists = errors.New("该用户已经存在")
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
