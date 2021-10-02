package model

type User struct {
	ID             int64  `json:"user_id"`
	Username       string `json:"username"`
	HashedPassword string `json:"hashed_password"`
	StudentID      string `json:"student_id"`
	CreatedAt      Time   `json:"created_at"`
}
