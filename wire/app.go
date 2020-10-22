package wire

import (
	"context"
	"test.com/wire/service"
)

type App struct {
	Person service.Person
}

func NewApp(ctx context.Context, person service.Person) *App {
	return &App{
		Person: person,
	}
}

func InitApp(ctx context.Context) *App {
	return initApp(ctx)
}
