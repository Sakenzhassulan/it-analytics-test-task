package handler

import "github.com/gin-gonic/gin"

func (h *Handler) DeleteTournament(ctx *gin.Context) {
	ok, err := h.Service.DeleteTournament(ctx)
	if err != nil {
		ctx.JSON(400, gin.H{
			"err": err.Error(),
		})
		return
	}

	ctx.JSON(200, ok)
}
