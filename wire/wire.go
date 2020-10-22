//+build wireinject

package wire

import (
	"context"
	"github.com/google/wire"
	"test.com/wire/db"
	"test.com/wire/service"
)

func initApp(ctx context.Context) *App {
	wire.Build(NewApp, service.NewPerson, service.NewCar, db.NewAddress, db.NewCity)

	return nil
}