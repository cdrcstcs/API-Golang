package inf

import (
	"context"

	"inititaryplanner/models"
)

//go:generate mockgen -source=./inf.go -destination=../mock/service_inf_mock.go -package=mock .
type AttractionService interface {
	CreateAttraction(ctx context.Context, req *models.CreateAttractionReq) (*models.CreateAttractionResp, error)
}

