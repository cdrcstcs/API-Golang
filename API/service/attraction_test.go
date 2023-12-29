package service

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"inititaryplanner/common/custom_errs"
	"inititaryplanner/dal/mock"
	"inititaryplanner/models"
)

func TestAttractionService_CreateAttraction(t *testing.T) {
	ctrl := gomock.NewController(t)
	ctx := context.Background()
	mockDal := mock.NewMockAttractionDal(ctrl)
	service := &attractionService{dal: mockDal}

	type arg struct {
		req *models.CreateAttractionReq
		ctx context.Context
	}

	tests := []struct {
		name string
		arg
		before   func(t *testing.T)
		wantResp *models.CreateAttractionResp
		wantErr  error
	}{
		{
			name: "success",
			arg: arg{
				ctx: ctx,
				req: &models.CreateAttractionReq{
					Name:    "test_name",
					Address: "test_address",
					Coordinate: models.Coordinate{
						X: 1,
						Y: 1,
					},
					TagIDs: []string{"1", "2"},
				},
			},
			before: func(t *testing.T) {
				mockDal.EXPECT().CreateAttraction(gomock.Any(), &models.Attraction{
					Id:      "",
					Name:    "test_name",
					Address: "test_address",
					Coordinate: models.Coordinate{
						X: 1,
						Y: 1,
					},
					TagIDs: []string{"1", "2"},
				}).Return(&models.Attraction{
					Id:      "test_id",
					Name:    "test_name",
					Address: "test_address",
					Coordinate: models.Coordinate{
						X: 1,
						Y: 1,
					},
					TagIDs: []string{"1", "2"},
				}, nil)
			},
			wantResp: &models.CreateAttractionResp{Attraction: &models.AttractionDTO{
				Id:      "test_id",
				Name:    "test_name",
				Address: "test_address",
				Coordinate: models.Coordinate{
					X: 1,
					Y: 1,
				},
				Tags: []models.Tag{
					{Id: "1", Value: "todo fill"},
					{Id: "2", Value: "todo fill"},
				},
			}},
		},
		{
			name: "db error",
			arg: arg{
				ctx: ctx,
				req: &models.CreateAttractionReq{
					Name:    "test_name",
					Address: "test_address",
					Coordinate: models.Coordinate{
						X: 1,
						Y: 1,
					},
					TagIDs: []string{"1", "2"},
				},
			},
			before: func(t *testing.T) {
				mockDal.EXPECT().CreateAttraction(gomock.Any(), gomock.Any()).Return(nil, custom_errs.DBErrIDConversion)
			},
			wantResp: nil,
			wantErr:  custom_errs.DBErrIDConversion,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.before(t)
			got, gotErr := service.CreateAttraction(tt.ctx, tt.req)
			assert.Equal(t, tt.wantErr, gotErr)
			if gotErr != nil {
				return
			}
			assert.Equal(t, tt.wantResp, got)
		})
	}
}
