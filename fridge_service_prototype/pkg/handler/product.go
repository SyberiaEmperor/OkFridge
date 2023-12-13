package handler

import (
	"net/http"

	"github.com/SyberiaEmperor/OkFridge/fridge_service_prototype/models"
	"github.com/gin-gonic/gin"
)

const (
	accountId = "id"
)

func (h *Handler) addProduct(c *gin.Context) {
	var prod models.Product

	if err := c.BindJSON(&prod); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	err := h.service.AddProduct(prod)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": prod.Id,
	})
}

func (h *Handler) changeProduct(c *gin.Context) {
	var prod models.Product

	if err := c.BindJSON(&prod); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	err := h.service.changeProduct(prod)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"ok": true,
	})
}