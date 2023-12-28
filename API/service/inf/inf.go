package inf

import (
	"context"

	"inititaryplanner/models"
)

//go:generate mockgen -source=./inf.go -destination=../mock/service_inf_mock.go -package=mock .
type AttractionDal interface {
	CreateAttraction(ctx context.Context, attraction *models.CreateAttractionReq) (*models.CreateAttractionResp, error)
}

