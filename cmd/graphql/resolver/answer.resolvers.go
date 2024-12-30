package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.60

import (
	"context"
	"graphql-quiz/cmd/graphql/dtos"
	"net/http"
)

// SubmitAnswer is the resolver for the SubmitAnswer field.
func (r *mutationResolver) SubmitAnswer(ctx context.Context, questionID string, optionID string) (*dtos.AnswerResponse, error) {
	result, err := r.app.SubmitAnswer(ctx, optionID)
	if err != nil {
		return answerErrResp(err)
	}

	return &dtos.AnswerResponse{
		Message: "",
		Status:  http.StatusOK,
		Data: &dtos.Answer{
			QuestionID: questionID,
			OptionID:   optionID,
			IsCorrect:  result,
		},
	}, nil
}