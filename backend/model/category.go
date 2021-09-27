package model

import "time"

type Category struct {
	ID        int64     `json:"category_id"`
	UserID    int64     `json:"user_id"`
	ProblemID int64     `json:"problem_id"`
	Author    string    `json:"author"`
	Content   string    `json:"content"`
	CreateAt  time.Time `json:"create_at"`
}
