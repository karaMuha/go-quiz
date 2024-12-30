package resolver

import (
	"errors"
	"graphql-quiz/cmd/application/domain"
	"graphql-quiz/cmd/graphql/dtos"
	"net/http"
)

func questionErrResp(errResp *domain.ErrResp) (*dtos.QuestionResponse, error) {
	if errResp.Status == http.StatusInternalServerError {
		return nil, errors.New(errResp.Message)
	}
	return &dtos.QuestionResponse{
		Message: errResp.Message,
		Status:  int32(errResp.Status),
	}, nil
}

func answerErrResp(errResp *domain.ErrResp) (*dtos.AnswerResponse, error) {
	if errResp.Status == http.StatusInternalServerError {
		return nil, errors.New(errResp.Message)
	}
	return &dtos.AnswerResponse{
		Message: errResp.Message,
		Status:  int32(errResp.Status),
	}, nil
}
