package application

import (
	"graphql-quiz/cmd/application/commands"
	"graphql-quiz/cmd/application/ports/driven"
	"graphql-quiz/cmd/application/ports/driver"
	"graphql-quiz/cmd/application/queries"
)

type Application struct {
	Commands
	Queries
}

type Commands struct {
	commands.QuestionsCommand
	commands.QuestionOptionsCommand
	commands.AnswersCommand
}

type Queries struct {
	queries.QuestionsQuery
	queries.QuestionOptionsQuery
	queries.AnswersQuery
}

var _ driver.IApplication = (*Application)(nil)

func New(
	questionOptionsRepository driven.IQuestionOptionsRepository,
	questionsRepository driven.IQuestionsRepository,
) Application {
	return Application{
		Commands: Commands{
			AnswersCommand:         commands.NewAnswersCommand(questionOptionsRepository),
			QuestionsCommand:       commands.NewQuestionsCommands(questionsRepository),
			QuestionOptionsCommand: commands.NewQuestionOptionsCommand(questionOptionsRepository),
		},
		Queries: Queries{
			AnswersQuery:         queries.NewAnswersQuery(),
			QuestionsQuery:       queries.NewQuestionsQuery(questionsRepository),
			QuestionOptionsQuery: queries.NewQuestionOptionsQuery(questionOptionsRepository),
		},
	}
}
