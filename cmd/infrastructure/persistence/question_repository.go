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

type questionsRepository struct {
	db *sql.DB
}

func NewQuestionsRepository(db *sql.DB) *questionsRepository {
	return &questionsRepository{
		db: db,
	}
}

var _ driven.IQuestionsRepository = &questionsRepository{}

func (r *questionsRepository) CreateQuestion(ctx context.Context, question *dtos.Question) (*dtos.Question, *domain.ErrResp) {
	query := `
		INSERT INTO questions (title, created_at, updated_at)
		VALUES ($1, $2, $3)
		RETURNING id
	`
	row := r.db.QueryRowContext(ctx, query, question.Title, question.CreatedAt, question.UpdatedAt)

	var id string
	err := row.Scan(&id)
	if err != nil {
		if strings.Contains(err.Error(), "unique constraint") {
			return nil, &domain.ErrResp{
				Message: "question title already taken",
				Status:  http.StatusConflict,
			}
		}
		return nil, &domain.ErrResp{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	question.ID = id

	return question, nil
}

func (r *questionsRepository) UpdateQuestion(ctx context.Context, question *dtos.Question) *domain.ErrResp {
	query := `
		UPDATE questions
		SET title = $1, updated_at = $1
		WHERE id = $3
	`
	_, err := r.db.ExecContext(ctx, query, question.Title, question.UpdatedAt, question.ID)
	if err != nil {
		if strings.Contains(err.Error(), "unique constraint") {
			return &domain.ErrResp{
				Message: "question title already taken",
				Status:  http.StatusConflict,
			}
		}
		return &domain.ErrResp{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	return nil

}

func (r *questionsRepository) DeleteQuestion(ctx context.Context, id string) *domain.ErrResp {
	query := `
		DELETE FROM questions
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

func (r *questionsRepository) GetQuestionByID(ctx context.Context, id string) (*dtos.Question, *domain.ErrResp) {
	query := `
		SELECT *
		FROM questions
		WHERE id = $1
	`
	row := r.db.QueryRowContext(ctx, query, id)
	var question dtos.Question
	err := row.Scan(&question.ID, &question.Title, &question.CreatedAt, &question.UpdatedAt)
	if err != nil {
		return nil, &domain.ErrResp{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	return &question, nil
}

func (r *questionsRepository) GetAllQuestions(ctx context.Context) ([]*dtos.Question, *domain.ErrResp) {
	query := `
		SELECT *
		FROM questions
	`
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, &domain.ErrResp{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	defer rows.Close()

	questionList := make([]*dtos.Question, 0)
	for rows.Next() {
		var question dtos.Question
		err := rows.Scan(&question.ID, &question.Title, &question.CreatedAt, &question.UpdatedAt)
		if err != nil {
			return nil, &domain.ErrResp{
				Message: err.Error(),
				Status:  http.StatusInternalServerError,
			}
		}

		questionList = append(questionList, &question)
	}

	return questionList, nil
}
