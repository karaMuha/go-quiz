package domain

import (
	"graphql-quiz/cmd/graphql/dtos"
	"net/http"
	"time"
)

type QuestionOption struct {
	ID         string
	QuestionID string
	Title      string `validate:"required"`
	IsCorrect  bool   `validate:"required,boolean"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

func CreateQuestionOption(questionID string, title string, isCorret bool) (*dtos.QuestionOption, *ErrResp) {
	questionOption := QuestionOption{
		QuestionID: questionID,
		Title:      title,
		IsCorrect:  isCorret,
	}

	err := validate.Struct(&questionOption)
	if err != nil {
		return nil, &ErrResp{
			Message: err.Error(),
			Status:  http.StatusBadRequest,
		}
	}

	return &dtos.QuestionOption{
		QuestionID: questionOption.QuestionID,
		Title:      questionOption.Title,
		IsCorrect:  questionOption.IsCorrect,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}, nil
}

func UpdateQuestionOption(title string, isCorrect bool, questionOption *dtos.QuestionOption) (*dtos.QuestionOption, *ErrResp) {
	questionOption.Title = title
	questionOption.IsCorrect = isCorrect
	questionOption.UpdatedAt = time.Now()

	err := validate.Struct(questionOption)
	if err != nil {
		return nil, &ErrResp{
			Message: err.Error(),
			Status:  http.StatusBadRequest,
		}
	}

	return questionOption, nil
}
