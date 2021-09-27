package docs

import "github.com/Thewalkers2012/DOJ/model"

// swagger:route POST /api/v1/category category CreateCategory
// 创建题目
// responses:
//	200: CreateCategory

// swagger:route Get /api/v1/category category GetCategoryByProblem
// 通过题目的编号来获取题目的信息
// responses:
// 	200: GetCategoryByProblem

// swagger:parameters CreateCategory
type createCategoryParamsWrapper struct {
	// This text will appear as description of your request body.
	// in:body
	Body model.CreateCategoryParams
}

// swagger:parameters GetCategoryByProblem
type getCategoryByProblemParamsWrapper struct {
	// This text will appear as description of your request body.
	// in:body
	Body model.GetCategoryByProblemParams
}

// swagger:response CreateCategory
type createCategoryResponseWrapper struct {
	// This text will appear as description of your response body.
	// in:body
	Body model.Category
}

// swagger:response GetCategoryByProblem
type getCategoryByProblemResponseWrapper struct {
	// This text will appear as description of your response body.
	// in:body
	Body model.Category
}
