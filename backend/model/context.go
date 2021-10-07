package model

type Context struct {
	ID        int64  `json:"context_id"`
	Title     string `json:"title"`
	StartTime Time   `json:"start_time"`
	EndTime   Time   `json:"end_time"`
	Defunct   string `json:"defunct"` // 判断该比赛是否还有效
	Author    string `json:"author"`
}
