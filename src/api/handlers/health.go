package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/harpy-py/go-Web-Api/api/helper"
)

type HealthHandler struct {
}

func NewHealthHandler() *HealthHandler {
	return &HealthHandler{}
}

func (h *HealthHandler) Health(ctx *gin.Context){
	ctx.JSON(http.StatusOK, helper.GenerateBaseResponse("working", true, 0))
	return
}