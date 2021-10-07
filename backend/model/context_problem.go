package model

type ContextProblem struct {
	ID        int64  `json:"context_problem_id"`
	ProblemID int64  `json:"problem_id"`
	ContextID int64  `json:"context_id"`
	Title     string `json:"title"`
}
