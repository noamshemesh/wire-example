package service

//go:generate mockgen -source=$GOFILE -destination=mock/car_mock.go

import (
	"context"
	"log"
)

type Car interface {
	Drive(ctx context.Context, address string)
}

type car struct {
}

func NewCar(ctx context.Context) Car {
	return &car{
	}
}

func (s *car) Drive(ctx context.Context, address string) {
	log.Printf("Car Service: Driving to %s\n", address)
}
