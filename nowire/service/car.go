package service

import (
	"context"
	"log"
)

type Car struct {
}

func (s *Car) Drive(ctx context.Context, address string) {
	log.Printf("Car Service: Driving to %s\n", address)
}
