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

// GetUserListParams
type GetUserListParams struct {
	PageNum  int `form:"pageNum" binding:"required,min=1"`
	PageSize int `form:"pageSize" binding:"required,min=5,max=10"`
}

type UpdateUserParams struct {
	UserID    int64  `json:"userID"`
	Username  string `json:"username"`
	StudentID string `json:"studentID"`
}

type UserParams struct {
	UserID int64 `form:"userID"`
}

type DeleteUserParams struct {
	UserID int64 `form:"userID"`
}

type UserDetailResponse struct {
	UserID          int64  `json:"userID"`
	AcceptCount     int64  `json:"acceptCount"`
	SubmissionCount int64  `json:"submissionCount"`
	Username        string `json:"username"`
	StudentID       string `json:"studentID"`
}

/*
* 与 Problem 相关的请求参数
 */

type Cases struct {
	Input  string `json:"input"`
	Output string `json:"output"`
}

type CreateProblemRequest struct {
	TimeLimit       int      `json:"timeLimit" binding:"required,min=1"`
	MemoryLimit     int      `json:"memoryLimit" binding:"required,min=1"`
	Name            string   `json:"problemName" binding:"required"`
	Description     string   `json:"description" binding:"required"`
	TestCase        string   `json:"testCase" binding:"required"`
	Author          string   `json:"author"`
	DifficultyLevel string   `json:"difficultyLevel" binding:"required"`
	Cases           []*Cases `json:"cases"`
}

type ListProblemRequest struct {
	PageNum  int `form:"pageNum" binding:"required,min=1"`
	PageSize int `form:"pageSize" binding:"required,min=5,max=10"`
}

// DeleteProblem params
type DeleteProblemParams struct {
	ProblemID int64 `form:"problemID"`
}

/*
* 与 Submission 相关的请求参数
 */

type RunCodeParams struct {
	ProblemID int64  `json:"problemID"`
	Language  string `json:"language"`
	Code      string `json:"code"`
}

type SubmitResult struct {
	AnswerCode int   `json:"answer_code"`
	Time       int   `json:"time"`
	Score      int   `json:"score"`
	Memory     int64 `json:"memory"`
}

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

// GetAllCategories Params
type GetAllCategoriesParams struct {
	PageNum  int `form:"pageNum" binding:"required,min=1"`
	PageSize int `form:"pageSize" binding:"required,min=5,max=10"`
}

// DeleteCategory And GetCategoryDetails
type CategoryParams struct {
	CategoryID int64 `form:"categoryID"`
}

/*
*	与比赛相关的请求参数
 */
type CreateContextParams struct {
	Title     string `json:"title"`
	StartTime Time   `json:"startTime"`
	EndTime   Time   `json:"endTime"`
	Author    string `json:"author"`
}

type GetContextListParams struct {
	PageNum  int `form:"pageNum" binding:"required,min=1"`
	PageSize int `form:"pageSize" binding:"required,min=5,max=10"`
}

type GetContextParams struct {
	ContextID int64 `params:"contextID"`
}

type DeleteContextParams struct {
	ContextID int64 `form:"contextID"`
}

type UpdateContextParams struct {
	ID        int64  `json:"contextID"`
	Title     string `json:"title"`
	StartTime Time   `json:"startTime"`
	EndTime   Time   `json:"endTime"`
	Author    string `json:"author"`
}

type AddProblemParams struct {
	ProblemID int64  `json:"problemID" binding:"required"`
	ContextID int64  `json:"contextID" binding:"required"`
	Title     string `json:"title"`
}
