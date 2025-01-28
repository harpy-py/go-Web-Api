package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/harpy-py/go-Web-Api/api/handlers"
	"github.com/harpy-py/go-Web-Api/api/middlewares"
	"github.com/harpy-py/go-Web-Api/config"
)

func User(router *gin.RouterGroup, conf *config.Config){
	h := handlers.NewUserHandler(conf)
	router.POST("/send-otp", middlewares.OTPLimiter(conf), h.SendOtp)
}