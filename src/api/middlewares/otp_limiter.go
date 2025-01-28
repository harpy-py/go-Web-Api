package middlewares

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/harpy-py/go-Web-Api/api/helper"
	"github.com/harpy-py/go-Web-Api/config"
	"github.com/harpy-py/go-Web-Api/pkg/limiter"
	"golang.org/x/time/rate"
)

func OTPLimiter(conf *config.Config) gin.HandlerFunc {
    var limiter = limiter.NewLimiter(rate.Every(conf.Otp.Limiter*time.Second), 1)
    return func(c *gin.Context) {
        limiter := limiter.GetLimiter(c.Request.RemoteAddr)
        if !limiter.Allow(){
            c.AbortWithStatusJSON(http.StatusTooManyRequests, helper.GenerateBaseResponseWithError(nil, false, -1, nil))
            c.Abort()
        }else {
            c.Next()
        }
    }
}
