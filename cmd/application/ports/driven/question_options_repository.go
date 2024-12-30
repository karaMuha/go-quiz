package driven

import (
	"context"
	"graphql-quiz/cmd/application/domain"
	"graphql-quiz/cmd/graphql/dtos"
)

type IQuestionOptionsRepository interface {
	CreateQuestionOption(ctx context.Context, question *dtos.QuestionOption) (*dtos.QuestionOption, *domain.ErrResp)
	UpdateQuestionOption(ctx context.Context, question *dtos.QuestionOption) *domain.ErrResp
	DeleteQuestionOption(ctx context.Context, id string) *domain.ErrResp
	DeleteQuestionOptionByQuestionID(ctx context.Context, questionId string) *domain.ErrResp
	GetQuestionOptionByID(ctx context.Context, id string) (*dtos.QuestionOption, *domain.ErrResp)
	GetQuestionOptionByQuestionID(ctx context.Context, questionId string) ([]*dtos.QuestionOption, *domain.ErrResp)
}
