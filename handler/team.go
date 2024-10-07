package handler

import (
	"fmt"

	"github.com/Sakenzhassulan/it-analytics-test-task/models"
	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateTeams(ctx *gin.Context) {
	var input models.TeamsInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(400, err.Error())
		return
	}

	teams, err := h.Service.CreateTeams(ctx, input.Teams)
	if err != nil {
		ctx.JSON(400, err.Error())
		fmt.Println(err)
		return
	}

	ctx.JSON(200, teams)
}
