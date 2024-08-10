package dataservices

import (
	"exo-planets/dataservices/model"
	"testing"

	"github.com/golang/mock/gomock"
	"gorm.io/gorm"
)

func TestDBClient_CreateExoplanet(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()


	type fields struct {
		DB *gorm.DB
	}
	type args struct {
		exoplanet model.Exoplanet
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db := &DBClient{
				DB: tt.fields.DB,
			}
			if err := db.CreateExoplanet(tt.args.exoplanet); (err != nil) != tt.wantErr {
				t.Errorf("DBClient.CreateExoplanet() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
