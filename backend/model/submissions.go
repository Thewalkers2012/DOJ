package model

type Submissions struct {
	ID          int64  `json:"submissionID"`
	UserID      int64  `json:"userID"`
	QuestionsID int64  `json:"questionID"`
	Language    string `json:"langueage"`
	Score       int    `json:"score"`
	Result      string `json:"result"`
	Code        string `json:"code"`
}
