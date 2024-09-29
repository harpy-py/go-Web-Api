package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/harpy-py/go-Web-Api/api/helper"
)

type header struct {
	UserId  string
	Browser string
}

type personData struct {
	FirstName string `json:"first_name" binding:"required,alpha,min=3,max=10"`
	LastName  string `json:"last_name" binding:"required,alpha,min=5,max=10"`
	MobileNumber string `json:"mobile_number" binding:"required,mobile,min=11,max=11"`
	Password string `json:"password" binding:"required,password"`
}

type TestHandler struct {
}

func NewTestHandler() *TestHandler {
	return &TestHandler{}
}

func (h *TestHandler) Test(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"result": "Test",
	})

	ctx.JSON(http.StatusOK, helper.GenerateBaseResponse("Test", true, 0))
}

func (h *TestHandler) Users(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"result": "Users",
	})
	ctx.JSON(http.StatusOK, helper.GenerateBaseResponse("Users", true, 0))
}

// UserById godoc
// @Summary UserById
// @Description UserById
// @Tags Test
// @Accept  json
// @Produce  json
// @Param id path int true "user id"
// @Success 200 {object} helper.BaseHttpResponse "Success"
// @Failure 400 {object} helper.BaseHttpResponse "failure"
// @Router /v1/test/user/{id} [get]
func (h *TestHandler) UserById(ctx *gin.Context) {
	id := ctx.Param("id")
	ctx.JSON(http.StatusOK, helper.GenerateBaseResponse(gin.H{
		"result": "UserById",
		"id":     id,
	},true, 0))
}

func (h *TestHandler) UserByUsername(ctx *gin.Context) {
	usrName := ctx.Param("username")
	ctx.JSON(http.StatusOK, helper.GenerateBaseResponse(gin.H{
		"result":   "UserByUsername",
		"username": usrName,
	},true, 0))
}

func (h *TestHandler) Accounts(ctx *gin.Context) {
	id := ctx.Param("id")
	ctx.JSON(http.StatusOK, helper.GenerateBaseResponse(gin.H{
		"result": "Accounts",
		"id":     id,
	},true, 0))
}

func (h *TestHandler) AddUser(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, helper.GenerateBaseResponse("AddUser", true, 0))
}

func (h *TestHandler) HeaderBinder1(ctx *gin.Context) {
	userId := ctx.GetHeader("UserId")

	ctx.JSON(http.StatusOK, helper.GenerateBaseResponse(gin.H{
		"result": "HeaderBinder1",
		"userId": userId,
	},true, 0))
}

func (h *TestHandler) HeaderBinder2(ctx *gin.Context) {
	header := header{}
	ctx.BindHeader(&header)

	ctx.JSON(http.StatusOK, helper.GenerateBaseResponse(gin.H{
		"result": "HeaderBinder2",
		"header": header,
	},true, 0))
}

func (h *TestHandler) QueryBinder1(ctx *gin.Context) {
	id := ctx.Query("id")
	name := ctx.Query("name")

	ctx.JSON(http.StatusOK, helper.GenerateBaseResponse(gin.H{
		"result": "HeaderBinder2",
		"id":     id,
		"name":   name,
	},true, 0))
}

func (h *TestHandler) QueryBinder2(ctx *gin.Context) {
	ids := ctx.QueryArray("id")
	name := ctx.Query("name")

	ctx.JSON(http.StatusOK, helper.GenerateBaseResponse(gin.H{
		"result": "HeaderBinder2",
		"ids":    ids,
		"name":   name,
	},true, 0))
}

func (h *TestHandler) UriBinder(ctx *gin.Context) {
	id := ctx.Param("id")
	name := ctx.Param("name")

	ctx.JSON(http.StatusOK, helper.GenerateBaseResponse(gin.H{
		"result": "UriBinder",
		"id":     id,
		"name":   name,
	},true, 0))
}

// BodyBinder godoc
// @Summary BodyBinder
// @Description BodyBinder
// @Tags Test
// @Accept  json
// @Produce  json
// @Param person body personData true "person data"
// @Success 200 {object} helper.BaseHttpResponse "Success"
// @Failure 400 {object} helper.BaseHttpResponse "failure"
// @Router /v1/test/binder/body [post]
func (h *TestHandler) BodyBinder(ctx *gin.Context) {
	p := personData{}
	err := ctx.ShouldBindJSON(&p)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, 
			helper.GenerateBaseResponseWithValidationError(nil, false, -1, err))
			return
	}
	ctx.JSON(http.StatusOK, helper.GenerateBaseResponse(gin.H{
		"result": "BodyBinder",
		"person": p,
	},true, 0))
}

func (h *TestHandler) FormBinder(ctx *gin.Context) {
	p := personData{}
	ctx.ShouldBind(&p)
	ctx.JSON(http.StatusOK,helper.GenerateBaseResponse(gin.H{
		"result": "FormBinder",
		"person": p,
	},true, 0))
}

func (h *TestHandler) FileBinder(ctx *gin.Context) {
	file, _ := ctx.FormFile("file")
	err := ctx.SaveUploadedFile(file, "file")
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, 
			helper.GenerateBaseResponseWithError(nil, false, -1, err))
			return
	}
	ctx.JSON(http.StatusOK, helper.GenerateBaseResponse(gin.H{
		"result": "FileBinder",
		"file":   file.Filename,
	}, true, 0))
}