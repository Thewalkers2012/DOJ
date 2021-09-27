package model

type CreateUserParams struct {
	Username       string `json:"username"`
	HashedPassword string `json:"hashed_password"`
	StudentID      string `json:"student_id"`
}

type LoginParams struct {
	StudentID string `json:"student_id"`
	Password  string `json:"password"`
}

type CreateProblemRequest struct {
	Name            string `json:"problem_name" binding:"required"`
	Description     string `json:"description" binding:"required"`
	TestCase        string `json:"test_case" binding:"required"`
	Author          string `json:"author"`
	TimeLimit       int64  `json:"time_limit" binding:"required,min=1"`
	MemoryLimit     int64  `json:"memory_limit" binding:"required,min=1"`
	DifficultyLevel string `json:"difficulty_level" binding:"required"`
}

type RunCodeParams struct {
	ProblemID int64  `json:"problemID"`
	Language  string `json:"language"`
	Code      string `json:"code"`
}

// 从前端获取的用户提交的参数
type CreateSubmissionRequest struct {
	UserID    int64  `json:"user_id"`
	ProblemID int64  `json:"problemID"`
	Language  string `json:"language"`
	Code      string `json:"code"`
}

// TODO: 从评测机来获取运行结果

type CreateCategoryParams struct {
	ProblemID int64  `json:"problemID"`
	Content   string `json:"content" binding:"required"`
}

type GetCategoryByProblemParams struct {
	ProblemID int64 `form:"problemID"`
}
