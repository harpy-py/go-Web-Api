package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type HealthHandler struct {
}

func NewHealthHandler() *HealthHandler {
	return &HealthHandler{}
}

func (h *HealthHandler) Health(ctx *gin.Context){
	ctx.JSON(http.StatusOK, "Working!")
	return
}

func (h *HealthHandler) HealthPost(ctx *gin.Context){
	ctx.JSON(http.StatusOK, "Post Working!")
	return
}

func (h *HealthHandler) HealthPostById(ctx *gin.Context){
	id := ctx.Params.ByName("id")
	ctx.JSON(http.StatusOK, fmt.Sprintf("Post ID: %s Working!", id))
	return
}