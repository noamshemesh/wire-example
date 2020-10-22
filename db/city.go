package db

import (
	"context"
	"log"
)

type City interface {
	RunQuery(ctx context.Context, cityName string)
}

type city struct {
}

func NewCity(ctx context.Context) City {
	return &city{
	}
}

func (s *city) RunQuery(ctx context.Context, cityName string) {
	log.Printf("City Db: Storing city %v...\n", cityName)
}
