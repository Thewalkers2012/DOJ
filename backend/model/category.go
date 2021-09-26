package model

type Category struct {
	ID        int64  `json:"category_id"`
	UserID    int64  `json:"user_id"`
	ProblemID int64  `json:"problem_id"`
	Title     string `json:"title"`
	Content   string `json:"content"`
}
