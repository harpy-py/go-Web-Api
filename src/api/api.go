package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/harpy-py/go-Web-Api/api/middlewares"
	"github.com/harpy-py/go-Web-Api/api/routers"
	validation "github.com/harpy-py/go-Web-Api/api/validations"
	"github.com/harpy-py/go-Web-Api/config"
	"github.com/harpy-py/go-Web-Api/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Initserver(conf *config.Config)  {
	r := gin.New()

	RegisterSwagger(r, conf)

	val, ok := binding.Validator.Engine().(*validator.Validate)
	if ok {
		val.RegisterValidation("mobile", validation.IranianMobileNumberValidator, true)
		val.RegisterValidation("password", validation.PasswdStrengthValidator)
	}
	
	r.Use(gin.Logger(), gin.Recovery(), middlewares.LimitByRequest(), middlewares.DefaultStructureLogger(conf))

	api := r.Group("/api")
	v1 := api.Group("/v1")
	{
		health := v1.Group("/health")
		test_router := v1.Group("/test")
		routers.Health(health)
		routers.TestRouter(test_router)
	}
	v2 := api.Group("/v2")
	{
		health := v2.Group("/health")
		routers.Health(health)
	}
	r.Run(fmt.Sprintf(":%s", conf.Server.Port))
}


func RegisterSwagger(r *gin.Engine, conf *config.Config){
	docs.SwaggerInfo.Title = "Golang web api"
	docs.SwaggerInfo.Description = "Golang web api"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.BasePath = "/api"
	docs.SwaggerInfo.Host = fmt.Sprintf("localhost:%s", conf.Server.Port)
	docs.SwaggerInfo.Schemes = []string{"http"}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}