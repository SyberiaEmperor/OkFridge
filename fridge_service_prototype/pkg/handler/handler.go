package handler

import (
	"github.com/SyberiaEmperor/OkFridge/fridge_service_prototype/pkg/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *service.Service	
}

func NewHandler(s *service.Service) *Handler {
	return &Handler{service: s}
}

func(h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	api := router.Group("/account")
	{
		api.POST("/products",h.addProduct)
		api.PUT("/products",h.changeProduct)
	}

	return router
}