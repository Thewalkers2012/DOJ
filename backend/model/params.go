package model

/*
* 与 User 结构体相关的请求
 */

// create user request
type CreateUserRequest struct {
	Username  string `json:"username" binding:"required"`
	Password  string `json:"password" binding:"required,min=6"`
	StudentID string `json:"studentID" binding:"required"`
}

// create user response
type CreateUserResponse struct {
	Username  string `json:"username"`
	StudentID string `json:"studentID"`
}

type CreateUserParams struct {
	Username       string `json:"username"`
	HashedPassword string `json:"hashed_password"`
	StudentID      string `json:"student_id"`
}

type LoginParams struct {
	StudentID string `json:"student_id"`
	Password  string `json:"password"`
}

// login request
type LoginRequest struct {
	StudentID string `json:"studentID" binding:"required"`
	Password  string `json:"password" binding:"required"`
}

// login response
type LoginResponse struct {
	AccessToekn string             `json:"access_token"`
	User        CreateUserResponse `json:"user"`
}

/*
* 与 Problem 相关的请求参数
 */

type CreateProblemRequest struct {
	Name            string `json:"problem_name" binding:"required"`
	Description     string `json:"description" binding:"required"`
	TestCase        string `json:"test_case" binding:"required"`
	Author          string `json:"author"`
	TimeLimit       int64  `json:"time_limit" binding:"required,min=1"`
	MemoryLimit     int64  `json:"memory_limit" binding:"required,min=1"`
	DifficultyLevel string `json:"difficulty_level" binding:"required"`
}

type ListProblemRequest struct {
	PageNum  int `form:"pageNum" binding:"required,min=1"`
	PageSize int `form:"pageSize" binding:"required,min=5,max=10"`
}

/*
* 与 Submission 相关的请求参数
 */

type RunCodeParams struct {
	ProblemID int64  `json:"problemID"`
	Language  string `json:"language"`
	Code      string `json:"code"`
}

type ResponseSubmit struct {
	AnswerCode int   `json:"answer_code"`
	Time       int   `json:"time"`
	Memory     int64 `json:"memory"`
}

// 从前端获取的用户提交的参数
type CreateSubmissionRequest struct {
	UserID    int64  `json:"user_id"`
	ProblemID int64  `json:"problemID"`
	Language  string `json:"language"`
	Code      string `json:"code"`
}

// TODO: 从评测机来获取运行结果

/*
*	与文章相关的请求参数
 */

type CreateCategoryParams struct {
	ProblemID int64  `json:"problemID"`
	Content   string `json:"content" binding:"required"`
}

type GetCategoryByProblemParams struct {
	ProblemID int64 `form:"problemID"`
}
