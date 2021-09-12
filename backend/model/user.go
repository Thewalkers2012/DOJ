package model

import "time"

type User struct {
	ID             int64     `json:"user_id"`
	Username       string    `json:"username"`
	HashedPassword string    `json:"hashed_password"`
	StudentID      string    `json:"student_id"`
	CreateAt       time.Time `json:"create_at"`
}
