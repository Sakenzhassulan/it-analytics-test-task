package main

import (
	"github.com/Sakenzhassulan/it-analytics-test-task/config"
	"github.com/Sakenzhassulan/it-analytics-test-task/handler"
	"github.com/Sakenzhassulan/it-analytics-test-task/repo"
	"github.com/Sakenzhassulan/it-analytics-test-task/service"
	"github.com/gin-gonic/gin"
)

func main() {
	config := config.LoadConfig()
	repo := repo.New(config)
	service := service.New(repo)
	router := gin.Default()
	handler := handler.New(service, config, router)
	handler.Run(config.Port)
}
