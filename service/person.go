package service

import (
	"context"
	"test.com/wire/db"
)

type Person interface {
	GoSomewhere(ctx context.Context, address string)
}

type person struct {
	car       Car
	dbAddress db.Address
}

func NewPerson(ctx context.Context, car Car, dbAddress db.Address) Person {
	return &person{
		car:       car,
		dbAddress: dbAddress,
	}
}

func (s *person) GoSomewhere(ctx context.Context, address string) {
	s.car.Drive(ctx, address)
	s.dbAddress.Store(ctx, address)
}
