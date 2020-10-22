package service

import (
	"context"
	"test.com/wire/nowire/db"
)

// Bad practice 1 - Not using interfaces
type Person struct {
}

func (s *Person) GoSomewhere(ctx context.Context, address string) {
	// Bad practice 2 - Initialize services directly in code
	(&Car{}).Drive(ctx, address)

	// Bad practice 3 - Use sync once
	db.NewAddress(ctx).Store(ctx, address)
}
