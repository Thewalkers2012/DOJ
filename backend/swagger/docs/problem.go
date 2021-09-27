package docs

import "github.com/Thewalkers2012/DOJ/model"

// swagger:route POST /api/v1/problem problem CreateProblemHandler
// 创建题目
// responses:
//	200: CreateProblemHandler

// swagger:route POST /api/v1/problem/{id} problem GetProblemByIDHandler
// 通过题目的编号来获取题目的信息
// responses:
// 	200: GetProblemByIDHandler

// swagger:route GET /api/v1/problem problem GetProblemList
// 获取题目列表
// responses:
//	200: GetProblemList

// swagger:parameters CreateProblemHandler
type createProblemParamsWrapper struct {
	// This text will appear as description of your request body.
	// in:body
	Body model.CreateProblemRequest
}

// swagger:parameters GetProblemByIDHandler
type getProblemByIDParamsWrapper struct {
	// This text will appear as description of your request body.
	// in:path
	Id string `json:"id"`
}

// swagger:parameters GetProblemList
type getProblemListParamsWrapper struct {
	// This text will appear as description of your request body.
	// in:body
	Body model.ListProblemRequest
}

// swagger:response CreateProblemHandler
type createProblemResponseWrapper struct {
	// This text will appear as description of your response body.
	// in:body
	Body model.Problem
}

// swagger:response GetProblemByIDHandler
type getProblemByIDResponseWrapper struct {
	// This text will appear as description of your response body.
	// in:body
	Body model.Problem
}

// swagger:response GetProblemList
type getProblemListResponseWrapper struct {
	// This text will appear as description of your response body.
	// in:body
	Body model.Problem
}
