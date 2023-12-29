package inf

import (
	"context"

	"inititaryplanner/models"
)

type AttractionController interface {
	CreateAttraction(ctx context.Context, req *models.CreateAttractionReq) (*models.CreateAttractionResp, error)
}
