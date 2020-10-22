package db

import (
	"context"
	"log"
)

type City struct {
}

func (s *City) RunQuery(ctx context.Context, cityName string) {
	log.Printf("City Db: Storing city %v...\n", cityName)
}
