package domain

import (
	"graphql-quiz/cmd/graphql/dtos"
	"net/http"
	"time"
)

type Question struct {
	ID             string
	Title          string
	QuestionOption []*QuestionOption
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

func CreateQuestion(title string) (*dtos.Question, *ErrResp) {
	if title == "" {
		return nil, &ErrResp{
			Message: "title is required",
			Status:  http.StatusBadRequest,
		}
	}

	return &dtos.Question{
		Title:     title,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}, nil
}

func UpdateQuestion(title string, question *dtos.Question) (*dtos.Question, *ErrResp) {
	if title == "" {
		return nil, &ErrResp{
			Message: "title is required",
			Status:  http.StatusBadRequest,
		}
	}

	question.Title = title
	question.UpdatedAt = time.Now()
	return question, nil
}
