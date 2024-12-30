package driver

import (
	"context"
	"graphql-quiz/cmd/application/commands"
	"graphql-quiz/cmd/application/domain"
	"graphql-quiz/cmd/graphql/dtos"
)

type IApplication interface {
	IQuestionsCommands
	IQuestionsQueries

	IQuestionOptionsCommand
	IQuestionOptionsQueries

	IAnswersCommands
	IAnswersQueries
}

type IQuestionsCommands interface {
	CreateQuestion(ctx context.Context, cmd *dtos.QuestionInput) (*dtos.Question, *domain.ErrResp)
	UpdateQuestion(ctx context.Context, cmd *commands.UpdateQuestionParams) (*dtos.Question, *domain.ErrResp)
	DeleteQuestion(ctx context.Context, questionID string) (*dtos.Question, *domain.ErrResp)
}

type IQuestionsQueries interface {
	GetQuestion(ctx context.Context, questionID string) (*dtos.Question, *domain.ErrResp)
	GetAllQuestions(ctx context.Context) ([]*dtos.Question, *domain.ErrResp)
}

type IQuestionOptionsCommand interface {
	CreateQuestionOption(ctx context.Context, questionID string, cmd *dtos.QuestionOptionInput) (*dtos.QuestionOption, *domain.ErrResp)
	UpdateQuestionOption(ctx context.Context, cmd *commands.UpdateQuestionOptionParams) (*dtos.QuestionOption, *domain.ErrResp)
	DeleteQuestionOption(ctx context.Context, id string) *domain.ErrResp
	DeleteQuestionOptionByQuestionID(ctx context.Context, questionId string) *domain.ErrResp
}

type IQuestionOptionsQueries interface {
	GetQuestionOptionByID(ctx context.Context, id string) (*dtos.QuestionOption, *domain.ErrResp)
	GetQuestionOptionByQuestionID(ctx context.Context, questionId string) ([]*dtos.QuestionOption, *domain.ErrResp)
}

type IAnswersCommands interface {
	SubmitAnswer(ctx context.Context, questionOptionID string) (bool, *domain.ErrResp)
}

type IAnswersQueries interface {
}
