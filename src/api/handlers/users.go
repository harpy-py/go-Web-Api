package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	dto "github.com/harpy-py/go-Web-Api/api/DTO"
	"github.com/harpy-py/go-Web-Api/api/helper"
	"github.com/harpy-py/go-Web-Api/config"
	"github.com/harpy-py/go-Web-Api/services"
)

type UserHandler struct {
	service *services.UserService
}

func NewUserHandler(conf *config.Config) *UserHandler{
	service := services.NewUserService(conf)
	return &UserHandler{service: service}
}

// SendOtp godoc
// @Summary Send otp to user
// @Description Send otp to user
// @Tags Users
// @Accept  json
// @Produce  json
// @Param Request body dto.GetOtpRequest true "GetOtpRequest"
// @Success 201 {object} helper.BaseHttpResponse "Success"
// @Failure 400 {object} helper.BaseHttpResponse "failure"
// @Failure 409 {object} helper.BaseHttpResponse "failure"
// @Router /v1/users/send-otp [post]
func (h *UserHandler) SendOtp(c *gin.Context){
	req := new(dto.GetOtpRequest)
	err := c.ShouldBindJSON(&req)
	if err != nil{
		c.AbortWithStatusJSON(http.StatusBadRequest,
		helper.GenerateBaseResponseWithValidationError(nil, false, -1, err))
		return
	}
	err = h.service.SendOtp(req)
	if err != nil{
		c.AbortWithStatusJSON(helper.TranslateErrorToStatusCode(err),
		helper.GenerateBaseResponseWithError(nil, false, -1, err))
		return
	}
	c.JSON(http.StatusCreated, helper.GenerateBaseResponse(nil, true, 0))
}