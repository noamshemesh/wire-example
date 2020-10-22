package service

import (
	"context"
	"github.com/golang/mock/gomock"
	"test.com/wire/db"
	mock_service "test.com/wire/service/mock"
	"testing"
)

func Test_person_GoSomewhere(t *testing.T) {
	type args struct {
		address string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "drive nowhere with fake car",
			args: args{
				address: "nowhere, Tel Aviv",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockCar := mock_service.NewMockCar(ctrl)
			s := &person{
				car:       mockCar,
				dbAddress: db.NewAddress(context.Background(), db.NewCity(context.Background())),
			}

			mockCar.EXPECT().Drive(context.Background(), tt.args.address)

			s.GoSomewhere(context.Background(), tt.args.address)
		})
	}
}
