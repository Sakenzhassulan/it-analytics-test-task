package handler

import (
	"github.com/Sakenzhassulan/it-analytics-test-task/config"
	"github.com/Sakenzhassulan/it-analytics-test-task/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	Service *service.Service
	Config  *config.Config
	Router  *gin.Engine
}

func New(service *service.Service, config *config.Config, router *gin.Engine) *Handler {
	handler := &Handler{
		Service: service,
		Config:  config,
		Router:  router,
	}
	handler.registerRoutes()
	return handler
}

func (h *Handler) Run(port string) {
	h.Router.Run(port)
}

func (h *Handler) registerRoutes() {
	r := h.Router.Group("api")

	r.POST("/teams", h.CreateTeams)
	r.POST("/generate/play-off", h.GeneratePlayOffResults)
	r.POST("/generate/:divisionName", h.GenerateResults)
	r.DELETE("/delete", h.DeleteTournament)
}
