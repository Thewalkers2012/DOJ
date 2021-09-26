package mysql

import "github.com/Thewalkers2012/DOJ/model"

func GetAllSubmissionsByIdAndProblem(userID, problemID int64) ([]*model.Submissions, error) {
	submissions := []*model.Submissions{}
	err := DB.Where("user_id = ? && question_id = ?").Find(&submissions).Error
	return submissions, err
}

// type Submissions struct {
// 	ID          int64  `json:"submissionID"`
// 	UserID      int64  `json:"userID"`
// 	QuestionsID int64  `json:"questionID"`
// 	Language    string `json:"langueage"`
// 	Score       int    `json:"score"`
// 	Result      string `json:"result"`
// 	Code        string `json:"code"`
// }

// func CreateSubmission(sub *model.CreateSubmissionRequest) (*model.Submissions, error) {
// 	submission := new(model.Submissions)
// 	DB.Select("UserID", "QuestionsID", "Language", "Code")
// }
