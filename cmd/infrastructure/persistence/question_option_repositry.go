package persistence

import (
	"context"
	"database/sql"
	"graphql-quiz/cmd/application/domain"
	"graphql-quiz/cmd/application/ports/driven"
	"graphql-quiz/cmd/graphql/dtos"
	"net/http"
	"strings"
)

type questionOptionsRepository struct {
	db *sql.DB
}

func NewQuestionOptionsRepository(db *sql.DB) *questionOptionsRepository {
	return &questionOptionsRepository{
		db: db,
	}
}

var _ driven.IQuestionOptionsRepository = &questionOptionsRepository{}

func (r *questionOptionsRepository) CreateQuestionOption(ctx context.Context, questOpt *dtos.QuestionOption) (*dtos.QuestionOption, *domain.ErrResp) {
	query := `
		INSERT INTO question_options (question_id, title, is_correct, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id
	`
	row := r.db.QueryRowContext(ctx, query, questOpt.QuestionID, questOpt.Title, questOpt.IsCorrect, questOpt.CreatedAt, questOpt.UpdatedAt)

	var id string
	err := row.Scan(&id)
	if err != nil {
		if strings.Contains(err.Error(), "unique constraint") {
			return nil, &domain.ErrResp{
				Message: "",
				Status:  http.StatusConflict,
			}
		}
		return nil, &domain.ErrResp{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	questOpt.ID = id
	return questOpt, nil
}

func (r *questionOptionsRepository) UpdateQuestionOption(ctx context.Context, questOpt *dtos.QuestionOption) *domain.ErrResp {
	query := `
		UPDATE question_options
		SET title = $1, is_correct = $2, updated_at = $3
		WHERE id = $4
	`
	_, err := r.db.ExecContext(ctx, query, questOpt.Title, questOpt.IsCorrect, questOpt.UpdatedAt, questOpt.ID)
	if err != nil {
		return &domain.ErrResp{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	return nil
}

func (r *questionOptionsRepository) DeleteQuestionOption(ctx context.Context, id string) *domain.ErrResp {
	query := `
		DELETE FROM question_options
		WHERE id = $1
	`
	_, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return &domain.ErrResp{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	return nil
}

func (r *questionOptionsRepository) DeleteQuestionOptionByQuestionID(ctx context.Context, questId string) *domain.ErrResp {
	query := `
		DELETE FROM question_options
		WHERE question_id = &1
	`
	_, err := r.db.ExecContext(ctx, query, questId)
	if err != nil {
		return &domain.ErrResp{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	return nil
}

func (r *questionOptionsRepository) GetQuestionOptionByID(ctx context.Context, id string) (*dtos.QuestionOption, *domain.ErrResp) {
	query := `
		SELECT *
		FROM question_options
		WHERE id = $1
	`
	row := r.db.QueryRowContext(ctx, query, id)

	var questionOption dtos.QuestionOption
	err := row.Scan(&questionOption.ID, &questionOption.Title, &questionOption.IsCorrect, &questionOption.CreatedAt, &questionOption.UpdatedAt)
	if err != nil {
		return nil, &domain.ErrResp{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	return &questionOption, nil
}

func (r *questionOptionsRepository) GetQuestionOptionByQuestionID(ctx context.Context, id string) ([]*dtos.QuestionOption, *domain.ErrResp) {
	query := `
		SELECT *
		FROM question_options
		WHERE question_id = $1
	`
	rows, err := r.db.QueryContext(ctx, query, id)
	if err != nil {
		return nil, &domain.ErrResp{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}
	defer rows.Close()

	questionOptions := make([]*dtos.QuestionOption, 0)
	for rows.Next() {
		var questionOption dtos.QuestionOption
		err := rows.Scan(&questionOption.ID, &questionOption.Title, &questionOption.IsCorrect, &questionOption.CreatedAt, &questionOption.UpdatedAt)
		if err != nil {
			return nil, &domain.ErrResp{
				Message: err.Error(),
				Status:  http.StatusInternalServerError,
			}
		}
		questionOptions = append(questionOptions, &questionOption)
	}
	return questionOptions, nil
}
