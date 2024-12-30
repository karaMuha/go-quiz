package resolver

import (
	"graphql-quiz/cmd/application/ports/driver"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	app driver.IApplication
}

func NewResolver(app driver.IApplication) Resolver {
	return Resolver{
		app: app,
	}
}
