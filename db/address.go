package db

import (
	"context"
	"strings"
)

type Address interface {
	Store(ctx context.Context, address string)
}

type address struct {
	city City
}

func NewAddress(ctx context.Context, city City) Address {
	return &address{
		city: city,
	}
}

func (s *address) Store(ctx context.Context, address string) {
	s.city.RunQuery(ctx, extractCity(ctx, address))
}

func extractCity(ctx context.Context, address string) string {
	if strings.Contains(address, "Tel Aviv") {
		return "Tel Aviv"
	}

	return "N/Person"
}
