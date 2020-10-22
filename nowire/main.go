package main

import (
	"context"
	"test.com/wire/nowire/service"
)

func main() {
	ctx := context.Background()
	(&service.Person{}).GoSomewhere(ctx, "Menachem Begin 156, Tel Aviv")
}
