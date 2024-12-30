package commands

import (
	"context"
	"graphql-quiz/cmd/application/domain"
	"graphql-quiz/cmd/application/ports/driven"
	"graphql-quiz/cmd/graphql/dtos"
)

type QuestionOptionsCommand struct {
	questionOptionsRepository driven.IQuestionOptionsRepository
}

func NewQuestionOptionsCommand(questionOptionsRepository driven.IQuestionOptionsRepository) QuestionOptionsCommand {
	return QuestionOptionsCommand{
		questionOptionsRepository: questionOptionsRepository,
	}
}

type UpdateQuestionOptionParams struct {
	QuestionOptionID string
	Title            string
	IsCorrect        bool
}

// CreateQuestionOption implements driver.IQuestionOptionsCommand.
func (c *QuestionOptionsCommand) CreateQuestionOption(ctx context.Context, questionID string, cmd *dtos.QuestionOptionInput) (*dtos.QuestionOption, *domain.ErrResp) {
	questionOption, errResp := domain.CreateQuestionOption(questionID, cmd.Title, cmd.IsCorrect)
	if errResp != nil {
		return nil, errResp
	}

	questionOption, errResp = c.questionOptionsRepository.CreateQuestionOption(ctx, questionOption)
	if errResp != nil {
		return nil, errResp
	}

	return questionOption, nil
}

// DeleteQuestionOption implements driver.IQuestionOptionsCommand.
func (c *QuestionOptionsCommand) DeleteQuestionOption(ctx context.Context, id string) *domain.ErrResp {
	panic("unimplemented")
}

// DeleteQuestionOptionByQuestionID implements driver.IQuestionOptionsCommand.
func (c *QuestionOptionsCommand) DeleteQuestionOptionByQuestionID(ctx context.Context, questionId string) *domain.ErrResp {
	panic("unimplemented")
}

// UpdateQuestionOption implements driver.IQuestionOptionsCommand.
func (c *QuestionOptionsCommand) UpdateQuestionOption(ctx context.Context, cmd *UpdateQuestionOptionParams) (*dtos.QuestionOption, *domain.ErrResp) {
	panic("unimplemented")
}
