package controllers

import (
	"context"

	"inititaryplanner/controllers/inf"
	"inititaryplanner/models"
	service_inf "inititaryplanner/service/inf"
)

func NewAttractionController(dal service_inf.AttractionService) inf.AttractionController {
	return &attractionController{
		service: dal,
	}
}

type attractionController struct {
	service service_inf.AttractionService
}

func (a *attractionController) CreateAttraction(ctx context.Context, req *models.CreateAttractionReq) (*models.CreateAttractionResp, error) {
	// we usually do request checking here at this layer, or can even do permission checking.
	// For now, I will leave it empty and just call service
	return a.service.CreateAttraction(ctx, req)
}
