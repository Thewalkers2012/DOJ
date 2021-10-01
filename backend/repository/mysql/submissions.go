package mysql

import (
	"github.com/Thewalkers2012/DOJ/model"
	"github.com/Thewalkers2012/DOJ/util/judge"
)

func GetAllSubmissionsByIdAndProblem(userID, problemID int64) ([]*model.Submission, error) {
	submissions := []*model.Submission{}
	err := DB.Where("user_id = ? && question_id = ?").Find(&submissions).Error
	return submissions, err
}

type Submissions struct {
	ID          int64  `json:"submissionID"`
	UserID      int64  `json:"userID"`
	QuestionsID int64  `json:"questionID"`
	Language    string `json:"langueage"`
	Score       int    `json:"score"`
	Result      string `json:"result"`
	Code        string `json:"code"`
}

func CreateSubmission(sub *model.CreateSubmissionRequest, score int, result int) (*model.Submission, error) {
	submission := &model.Submission{
		UserID:    sub.UserID,
		ProblemID: sub.ProblemID,
		Language:  sub.Language,
		Code:      sub.Code,
		Result:    judge.GetAnswerMsg(result),
		Score:     score,
	}

	err := DB.Select("UserID", "ProblemID", "Language", "Code", "Score", "Result").Create(submission).Error

	return submission, err
}
