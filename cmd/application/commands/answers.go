package commands

import (
	"context"
	"graphql-quiz/cmd/application/domain"
	"graphql-quiz/cmd/application/ports/driven"
)

type AnswersCommand struct {
	questionOptionsRepository driven.IQuestionOptionsRepository
}

func NewAnswersCommand(questionOptionsRepository driven.IQuestionOptionsRepository) AnswersCommand {
	return AnswersCommand{
		questionOptionsRepository: questionOptionsRepository,
	}
}

// SubmitAnswer implements driver.IAnswersCommands.
func (c *AnswersCommand) SubmitAnswer(ctx context.Context, questionOptionID string) (bool, *domain.ErrResp) {
	questionOption, respErr := c.questionOptionsRepository.GetQuestionOptionByID(ctx, questionOptionID)
	if respErr != nil {
		return false, respErr
	}

	return domain.CheckAnswer(questionOption), nil
}
