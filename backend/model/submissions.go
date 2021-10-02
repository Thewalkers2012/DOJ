package model

type Submission struct {
	ID        int64  `json:"submission_id"`
	UserID    int64  `json:"user_id"`
	ProblemID int64  `json:"problem_id"`
	Memory    int64  `json:"memory"`
	Time      int    `json:"time"`
	Score     int    `json:"score"`
	Language  string `json:"language"`
	Result    string `json:"result"`
	Code      string `json:"code"`
	CreatedAt Time   `json:"created_at"`
}
