package main

import (
	"context"
	"test.com/wire/wire"
)

func main() {
	ctx := context.Background()
	app := wire.InitApp(ctx)

	app.Person.GoSomewhere(ctx, "Menachem Begin 156, Tel Aviv")
}
