package db

import (
	"context"
	"strings"
	"sync"
)

type Address struct {
	city *City
}

var addressOnce sync.Once
var addressInstance *Address

func NewAddress(ctx context.Context) *Address {
	addressOnce.Do(func() {
		addressInstance = &Address{
			city: &City{},
		}
	})

	return addressInstance
}

func (s *Address) Store(ctx context.Context, address string) {
	s.city.RunQuery(ctx, extractCity(ctx, address))
}

func extractCity(ctx context.Context, address string) string {
	if strings.Contains(address, "Tel Aviv") {
		return "Tel Aviv"
	}

	return "N/Person"
}
