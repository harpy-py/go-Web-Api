package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type header struct {
	UserId  string
	Browser string
}

type personData struct {
	FirstName string `json:"first_name" binding:"required,alpha,min=3,max=10"`
	LastName  string `json:"last_name" binding:"required,alpha,min=5,max=10"`
	MobileNumber string `json:"mobile_number" binding:"required,mobile,min=11,max=11"`
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
}

func (h *TestHandler) Users(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"result": "Users",
	})
}

func (h *TestHandler) UserById(ctx *gin.Context) {
	id := ctx.Param("id")
	ctx.JSON(http.StatusOK, gin.H{
		"result": "UserById",
		"id":     id,
	})
}

func (h *TestHandler) UserByUsername(ctx *gin.Context) {
	usrName := ctx.Param("username")
	ctx.JSON(http.StatusOK, gin.H{
		"result":   "UserByUsername",
		"username": usrName,
	})
}

func (h *TestHandler) Accounts(ctx *gin.Context) {
	id := ctx.Param("id")
	ctx.JSON(http.StatusOK, gin.H{
		"result": "Accounts",
		"id":     id,
	})
}

func (h *TestHandler) AddUser(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"result": "AddUser",
	})
}

func (h *TestHandler) HeaderBinder1(ctx *gin.Context) {
	userId := ctx.GetHeader("UserId")

	ctx.JSON(http.StatusOK, gin.H{
		"result": "HeaderBinder1",
		"userId": userId,
	})
}

func (h *TestHandler) HeaderBinder2(ctx *gin.Context) {
	header := header{}
	ctx.BindHeader(&header)

	ctx.JSON(http.StatusOK, gin.H{
		"result": "HeaderBinder2",
		"header": header,
	})
}

func (h *TestHandler) QueryBinder1(ctx *gin.Context) {
	id := ctx.Query("id")
	name := ctx.Query("name")

	ctx.JSON(http.StatusOK, gin.H{
		"result": "HeaderBinder2",
		"id":     id,
		"name":   name,
	})
}

func (h *TestHandler) QueryBinder2(ctx *gin.Context) {
	ids := ctx.QueryArray("id")
	name := ctx.Query("name")

	ctx.JSON(http.StatusOK, gin.H{
		"result": "HeaderBinder2",
		"ids":    ids,
		"name":   name,
	})
}

func (h *TestHandler) UriBinder(ctx *gin.Context) {
	id := ctx.Param("id")
	name := ctx.Param("name")

	ctx.JSON(http.StatusOK, gin.H{
		"result": "UriBinder",
		"id":     id,
		"name":   name,
	})
}

func (h *TestHandler) BodyBinder(ctx *gin.Context) {
	p := personData{}
	err := ctx.ShouldBindJSON(&p)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"validationError": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"result": "BodyBinder",
		"person": p,
	})
}

func (h *TestHandler) FormBinder(ctx *gin.Context) {
	p := personData{}
	ctx.ShouldBind(&p)
	ctx.JSON(http.StatusOK, gin.H{
		"result": "FormBinder",
		"person": p,
	})
}

func (h *TestHandler) FileBinder(ctx *gin.Context) {
	file, _ := ctx.FormFile("file")
	err := ctx.SaveUploadedFile(file, "file")
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"result": "FileBinder",
		"file":   file.Filename,
	})
}
