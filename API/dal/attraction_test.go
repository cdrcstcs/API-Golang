package dal

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"inititaryplanner/common/custom_errs"
	"inititaryplanner/constant"
	"inititaryplanner/dal/db"
	"inititaryplanner/models"
)

func TestCreateAttraction(t *testing.T) {
	ctx := context.Background()

	type arg struct {
		attraction *models.Attraction
		ctx        context.Context
	}

	tests := []struct {
		name   string
		before func(t *testing.T)
		arg
		wantErr        error
		wantAttraction *models.Attraction
	}{
		{
			name: "success",
			arg: arg{
				ctx: ctx,
				attraction: &models.Attraction{
					Name: "test",
				},
			},
			wantAttraction: &models.Attraction{
				Name: "test",
			},
		},
		{
			name: "with id err",
			arg: arg{
				ctx: ctx,
				attraction: &models.Attraction{
					Id:   "1",
					Name: "test",
				},
			},
			wantErr: custom_errs.DBErrCreateWithID,
		},
	}

	attractionDal := attractionDal{mainDB: db.GetMemoMongo(constant.MainMongoDB)}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotAttraction, err := attractionDal.CreateAttraction(tt.ctx, tt.attraction)
			assert.Equal(t, tt.wantErr, err)
			if err != nil {
				return
			}
			assert.NotEmpty(t, gotAttraction.Id)
			gotAttraction.Id = ""
			assert.Equal(t, tt.wantAttraction, gotAttraction)
		})
	}
}
