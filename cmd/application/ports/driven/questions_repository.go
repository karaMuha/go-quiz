package driven

import (
	"context"
	"graphql-quiz/cmd/application/domain"
	"graphql-quiz/cmd/graphql/dtos"
)

type IQuestionsRepository interface {
	CreateQuestion(ctx context.Context, question *dtos.Question) (*dtos.Question, *domain.ErrResp)
	UpdateQuestion(ctx context.Context, question *dtos.Question) *domain.ErrResp
	DeleteQuestion(ctx context.Context, id string) *domain.ErrResp
	GetQuestionByID(ctx context.Context, id string) (*dtos.Question, *domain.ErrResp)
	GetAllQuestions(ctx context.Context) ([]*dtos.Question, *domain.ErrResp)
}
