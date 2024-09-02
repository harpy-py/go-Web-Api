package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type TestHandler struct {
}

func NewTestHandler() *TestHandler {
	return &TestHandler{}
}

func (h *TestHandler) Test(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"result" : "Test",
	})
}

func (h *TestHandler) Users(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"result" : "Users",
	})
}


func (h *TestHandler) UserById(ctx *gin.Context) {
	id := ctx.Param("id")
	ctx.JSON(http.StatusOK, gin.H{
		"result" : "UserById",
		"id" : id,
	})
}

func (h *TestHandler) UserByUsername(ctx *gin.Context) {
	usrName := ctx.Param("username")
	ctx.JSON(http.StatusOK, gin.H{
		"result" : "UserByUsername",
		"username" : usrName,
	})
}

func (h *TestHandler) Accounts(ctx *gin.Context) {
	id := ctx.Param("id")
	ctx.JSON(http.StatusOK, gin.H{
		"result" : "Accounts",
		"id" : id,
	})
}

func (h *TestHandler) AddUser(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"result" : "AddUser",
	})
}