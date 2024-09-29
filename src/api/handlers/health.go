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

// HealthChecker godoc
// @Summary Health Check
// @Description Health Check
// @Tags health
// @Accept  json
// @Produce  json
// @Success 200 {object} helper.BaseHttpResponse "Success"
// @Failure 400 {object} helper.BaseHttpResponse "failure"
// @Router /v1/health/ [get]
func (h *HealthHandler) Health(ctx *gin.Context){
	ctx.JSON(http.StatusOK, helper.GenerateBaseResponse("working", true, 0))
	return
}