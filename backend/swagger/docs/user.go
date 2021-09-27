package docs

import (
	"github.com/Thewalkers2012/DOJ/model"
)

// swagger:route POST /api/v1/login user LoginHandler
// 用户登录
// responses:
//	200: LoginHandler

// swagger:route POST /api/v1/register user SignUpHandler
// 用户注册
// responses:
// 	200: SignUpHandler

// swagger:route GET /api/v1/info user InfoHandler
// 用户信息
// responses:
//	200: InfoHandler

// swagger:parameters LoginHandler
type loginParamsWrapper struct {
	// This text will appear as description of your request body.
	// in:body
	Body model.LoginRequest
}

// swagger:parameters SignUpHandler
type signupParamsWrapper struct {
	// This text will appear as description of your request body.
	// in:body
	Body model.CreateUserRequest
}

// swagger:response LoginHandler
type loginResponseWrapper struct {
	// This text will appear as description of your response body.
	// in:body
	Body model.LoginResponse
}

// swagger:response SignUpHandler
type signupResponseWrapper struct {
	// This text will appear as description of your response body.
	// in:body
	Body model.CreateUserResponse
}

// swagger:response InfoHandler
type infoResponseWrapper struct {
	// This text will appear as description of your response body.
	// in:body
	Body model.User
}
