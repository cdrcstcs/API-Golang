package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"inititaryplanner/common/gin_ctx"
	"inititaryplanner/controllers/inf"
	"inititaryplanner/models"
)

type MainController interface {
	CreateAttraction(c *gin.Context)
}

func NewMainController(
	ac inf.AttractionController,
) MainController {
	return &mainController{
		AttractionController: ac,
	}
}

type mainController struct {
	inf.AttractionController
}

func (m mainController) CreateAttraction(c *gin.Context) {
	ctx, req, err := gin_ctx.GetCtxAndReqFromGinCtx[models.CreateAttractionReq](c, models.CreateAttractionReq{})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Error decoding request body: %s", err.Error())})
		return
	}
	resp, err := m.AttractionController.CreateAttraction(ctx, &req)
	// we usually just use ok status for client
	c.JSON(http.StatusOK, resp)
}
