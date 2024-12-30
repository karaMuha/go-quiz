package queries

import (
	"context"
	"graphql-quiz/cmd/application/domain"
	"graphql-quiz/cmd/application/ports/driven"
	"graphql-quiz/cmd/graphql/dtos"
)

type QuestionOptionsQuery struct {
	questionOptionsRepository driven.IQuestionOptionsRepository
}

func NewQuestionOptionsQuery(questionOptionsRepository driven.IQuestionOptionsRepository) QuestionOptionsQuery {
	return QuestionOptionsQuery{
		questionOptionsRepository: questionOptionsRepository,
	}
}

// GetQuestionOptionByID implements driver.IQuestionOptionsQueries.
func (q *QuestionOptionsQuery) GetQuestionOptionByID(ctx context.Context, id string) (*dtos.QuestionOption, *domain.ErrResp) {
	panic("unimplemented")
}

// GetQuestionOptionByQuestionID implements driver.IQuestionOptionsQueries.
func (q *QuestionOptionsQuery) GetQuestionOptionByQuestionID(ctx context.Context, questionId string) ([]*dtos.QuestionOption, *domain.ErrResp) {
	panic("unimplemented")
}
