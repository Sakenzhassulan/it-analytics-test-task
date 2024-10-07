package handler

import "github.com/gin-gonic/gin"

func (h *Handler) GenerateResults(ctx *gin.Context) {
	division := ctx.Param("divisionName")

	teams, err := h.Service.GenerateResults(ctx, division)
	if err != nil {
		ctx.JSON(400, err.Error())
		return
	}

	ctx.JSON(200, teams)
}

func (h *Handler) GeneratePlayOffResults(ctx *gin.Context) {
	results, err := h.Service.GeneratePlayOffResults(ctx)
	if err != nil {
		ctx.JSON(400, err.Error())
		return
	}

	ctx.JSON(200, results)
}
