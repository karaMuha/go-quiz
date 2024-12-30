package queries

import (
	"context"
	"graphql-quiz/cmd/application/domain"
	"graphql-quiz/cmd/application/ports/driven"
	"graphql-quiz/cmd/graphql/dtos"
)

type QuestionsQuery struct {
	questionsRepository driven.IQuestionsRepository
}

func NewQuestionsQuery(questionsRepository driven.IQuestionsRepository) QuestionsQuery {
	return QuestionsQuery{
		questionsRepository: questionsRepository,
	}
}

// GetAllQuestions implements driver.IQuestionsQueries.
func (q *QuestionsQuery) GetAllQuestions(ctx context.Context) ([]*dtos.Question, *domain.ErrResp) {
	panic("unimplemented")
}

// GetQuestion implements driver.IQuestionsQueries.
func (q *QuestionsQuery) GetQuestion(ctx context.Context, questionID string) (*dtos.Question, *domain.ErrResp) {
	panic("unimplemented")
}
