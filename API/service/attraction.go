package service

import (
	"context"

	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"

	"inititaryplanner/common/custom_errs"
	dal_inf "inititaryplanner/dal/inf"
	"inititaryplanner/models"
	"inititaryplanner/service/inf"
)

func NewAttractionService(dal dal_inf.AttractionDal) inf.AttractionService {
	return &attractionService{
		dal: dal,
	}
}

type attractionService struct {
	dal dal_inf.AttractionDal
}

func (a *attractionService) CreateAttraction(ctx context.Context, req *models.CreateAttractionReq) (*models.CreateAttractionResp, error) {
	// TODO verify list of tag if they are valid in db

	attraction := &models.Attraction{}
	err := copier.Copy(attraction, req)
	if err != nil {
		log.Info().Ctx(ctx).Msgf("copier fails %v", err)
		return nil, errors.Wrap(custom_errs.ServerError, err.Error())
	}

	attraction, err = a.dal.CreateAttraction(ctx, attraction)
	if err != nil {
		// TODO logging
		return nil, err
	}

	dto, err := a.convertDBOToDTOAttraction(ctx, attraction)
	if err != nil {
		// TODO logging
		return nil, err
	}

	return &models.CreateAttractionResp{Attraction: dto}, nil
}

func (a *attractionService) convertDBOToDTOAttraction(ctx context.Context, att *models.Attraction) (*models.AttractionDTO, error) {
	if att == nil {
		return nil, custom_errs.ServerError
	}

	attraction := &models.AttractionDTO{}
	err := copier.Copy(attraction, att)
	if err != nil {
		// TODO logging
		return nil, errors.Wrap(custom_errs.ServerError, err.Error())
	}

	// TODO overfill the attraction tag here with value from db
	var tags []models.Tag
	for _, tID := range att.TagIDs {
		tags = append(tags, models.Tag{Id: tID, Value: "todo fill"})
	}

	attraction.Tags = tags
	return attraction, nil
}
