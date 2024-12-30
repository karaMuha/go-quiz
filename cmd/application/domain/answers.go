package domain

import "graphql-quiz/cmd/graphql/dtos"

func CheckAnswer(questionOption *dtos.QuestionOption) bool {
	return questionOption.IsCorrect
}
