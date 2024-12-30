package commands

import (
	"context"
	"graphql-quiz/cmd/application/domain"
	"graphql-quiz/cmd/application/ports/driven"
	"graphql-quiz/cmd/graphql/dtos"
)

type QuestionsCommand struct {
	questionsRepository driven.IQuestionsRepository
}

type UpdateQuestionParams struct {
	QuestionID string
	Title      string
}

func NewQuestionsCommands(questionsRepository driven.IQuestionsRepository) QuestionsCommand {
	return QuestionsCommand{
		questionsRepository: questionsRepository,
	}
}

// CreateQuestion implements driver.IQuestionsCommands.
func (c *QuestionsCommand) CreateQuestion(ctx context.Context, cmd *dtos.QuestionInput) (*dtos.Question, *domain.ErrResp) {
	question, errResp := domain.CreateQuestion(cmd.Title)
	if errResp != nil {
		return nil, errResp
	}

	question, errResp = c.questionsRepository.CreateQuestion(ctx, question)
	if errResp != nil {
		return nil, errResp
	}

	return question, nil
}

// DeleteQuestion implements driver.IQuestionsCommands.
func (c *QuestionsCommand) DeleteQuestion(ctx context.Context, questionID string) (*dtos.Question, *domain.ErrResp) {
	panic("unimplemented")
}

// UpdateQuestion implements driver.IQuestionsCommands.
func (c *QuestionsCommand) UpdateQuestion(ctx context.Context, cmd *UpdateQuestionParams) (*dtos.Question, *domain.ErrResp) {
	question, err := c.questionsRepository.GetQuestionByID(ctx, cmd.QuestionID)
	if err != nil {
		return nil, err
	}

	updatedQuestion, err := domain.UpdateQuestion(cmd.Title, question)
	if err != nil {
		return nil, err
	}

	err = c.questionsRepository.UpdateQuestion(ctx, updatedQuestion)
	if err != nil {
		return nil, err
	}

	return updatedQuestion, nil
}
