package inf

import (
	"context"

	"inititaryplanner/models"
)

type AttractionDal interface {
	CreateAttraction(ctx context.Context, attraction models.Attraction)
}
